package spoon

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"soup/utils"
)

func ToPkg(Filepath string){
	ex, _ := os.Executable()
	realEx, _ := filepath.EvalSymlinks(ex)
	ExeLocation := filepath.Dir(realEx)
	StdPath := filepath.Join(ExeLocation, "/pkg/")
	
	content := utils.ReadFile(filepath.Join(Filepath, "/pkg.json"))
	var payload Data
	json.Unmarshal([]byte(content), &payload)
	
	CopyDir(Filepath, filepath.Join(StdPath, "/"+payload.Id))

	if _, err := os.Stat(filepath.Join(ExeLocation, "temp/")); os.IsNotExist(err){
		os.Mkdir(filepath.Join(ExeLocation, "temp/"), 0777)
	}

	for _, v := range payload.Dependencies {
		q:="temp/"+RandStringBytes(100)
		ss := filepath.Join(ExeLocation, q+".zip")
		sq := filepath.Join(ExeLocation, q)
		downloadZip(v.Pkgurl+"/zipball/master", ss)
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
	}

}

func GetDependencies(Filepath string){
	ex, _ := os.Executable()
	realEx, _ := filepath.EvalSymlinks(ex)
	ExeLocation := filepath.Dir(realEx)
	
	content := utils.ReadFile(filepath.Join(Filepath, "/pkg.json"))
	var payload Data
	json.Unmarshal([]byte(content), &payload)

	for _, v := range payload.Dependencies {
		q:="temp/"+RandStringBytes(100)
		ss := filepath.Join(ExeLocation, q+".zip")
		sq := filepath.Join(ExeLocation, q)
		downloadZip(v.Pkgurl+"/zipball/master", ss)
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
	}

}