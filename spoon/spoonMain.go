package spoon

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/url"
	"os"
	"path/filepath"
	"soup/utils"
	"strings"
	"syscall"

	"golang.org/x/sys/windows"
)

func CopyDir(src string, dest string) error {

    if dest == src {
        return fmt.Errorf("Cannot copy a folder into the folder itself!")
    }

    f, err := os.Open(src)
    if err != nil {
        return err
    }

    file, err := f.Stat()
    if err != nil {
        return err
    }
    if !file.IsDir() {
        return fmt.Errorf("Source " + file.Name() + " is not a directory!")
    }

    err = os.Mkdir(dest, 0755)
    if err != nil {
        return err
    }

    files, err := ioutil.ReadDir(src)
    if err != nil {
        return err
    }

    for _, f := range files {

        if f.IsDir() {

            err = CopyDir(src+"/"+f.Name(), dest+"/"+f.Name())
            if err != nil {
                return err
            }

        }

        if !f.IsDir() {

            content, err := ioutil.ReadFile(src + "/" + f.Name())
            if err != nil {
                return err

            }

            err = ioutil.WriteFile(dest+"/"+f.Name(), content, 0755)
            if err != nil {
                return err

            }

        }

    }

    return nil
}

type Dep struct {
	Pkgurl string `pkgurl:"id"`
}

type Data struct {
    Name string `json:"name"`
    Id   string `json:"id"`
	Type   string `json:"type"`
    Description string `description:"id"`
	Dependencies []Dep `dependencies:"id"`
}

func runMeElevated() {
	verb := "runas"
	exe, _ := os.Executable()
	cwd, _ := os.Getwd()
	args := strings.Join(os.Args[1:], " ")

	verbPtr, _ := syscall.UTF16PtrFromString(verb)
	exePtr, _ := syscall.UTF16PtrFromString(exe)
	cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
	argPtr, _ := syscall.UTF16PtrFromString(args)

	var showCmd int32 = 1 //SW_NORMAL

	err := windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd)
	if err != nil {
		fmt.Println(err)
	}
}

func amAdmin() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
		return false
	}
	return true
}

func FilenameFromUrl(urlstr string) string {
	u, err := url.Parse(urlstr)
	if err != nil {
		log.Fatal("Error due to parsing url: ", err)
	}
	x, _ := url.QueryUnescape(u.EscapedPath())
	return filepath.Base(x)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}

func Spoon(){

	ex, _ := os.Executable()
	realEx, _ := filepath.EvalSymlinks(ex)
	ExeLocation := filepath.Dir(realEx)
	if !amAdmin() {
		runMeElevated()
	}else {
		if (len(os.Args) < 2){
			
			fmt.Println(`
			
			Get Package From Github
			spoon.exe get {github_repo_url}

			Get Project Dependencies
			spoon.exe get_dep

			Send project to main PKG folder
			spoon.exe add_pkg

			Create base for project
			spoon.exe init_pkg {name} {desc}

			`)
			
		}else if len(os.Args) > 1{
			if (os.Args[1] == "get" && len(os.Args) == 3){
				if _, err := os.Stat(filepath.Join(ExeLocation, "temp/")); os.IsNotExist(err){
					os.Mkdir(filepath.Join(ExeLocation, "temp/"), 0777)
				}
				q:="temp/"+RandStringBytes(100)
				ss := filepath.Join(ExeLocation, q+".zip")
				sq := filepath.Join(ExeLocation, q)
				downloadZip(os.Args[2]+"/zipball/master", ss)
				Unzip(ss, sq)
				var z string
				p, _ := ioutil.ReadDir(sq)
				for _, v := range p {
					if (v.IsDir()){
						z=filepath.Join(sq, "/"+v.Name()+"/")
						break
					}
				}
				ToPkg(z)
				os.RemoveAll(ss)
				os.Remove(ss)
				os.RemoveAll(sq)
				os.Remove(sq)
			}else if os.Args[1] == "get_dep" && len(os.Args) == 2 {
				Filepath, _ := filepath.Abs("./")
				GetDependencies(Filepath)
			}else if os.Args[1] == "add_pkg" && len(os.Args) == 2 {
				Filepath, _ := filepath.Abs("./")
				ToPkg(Filepath)
			}else if os.Args[1] == "init_pkg" && len(os.Args) == 3 {
				init_pkg(os.Args[2], os.Args[3])
			}else {
				utils.Error("Spoon Command Does Not Exist Or Not All Of The Parameters Were Not Filled")
			}
		}
	}

}