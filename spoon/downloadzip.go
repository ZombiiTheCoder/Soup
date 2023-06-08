package spoon

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func unzipSource(source, destination string) error {
    // 1. Open the zip file
    reader, err := zip.OpenReader(source)
    if err != nil {
        return err
    }
    defer reader.Close()

    // 2. Get the absolute destination path
    destination, err = filepath.Abs(destination)
    if err != nil {
        return err
    }

    // 3. Iterate over zip files inside the archive and unzip each of them
    for _, f := range reader.File {
        err := unzipFile(f, destination)
        if err != nil {
            return err
        }
    }

    return nil
}

func unzipFile(f *zip.File, destination string) error {
    // 4. Check if file paths are not vulnerable to Zip Slip
    filePath := filepath.Join(destination, f.Name)
    if !strings.HasPrefix(filePath, filepath.Clean(destination)+string(os.PathSeparator)) {
        return fmt.Errorf("invalid file path: %s", filePath)
    }

    // 5. Create directory tree
    if f.FileInfo().IsDir() {
        if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
            return err
        }
        return nil
    }

    if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
        return err
    }

    // 6. Create a destination file for unzipped content
    destinationFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
    if err != nil {
        return err
    }
    defer destinationFile.Close()

    // 7. Unzip the content of a file and copy it to the destination file
    zippedFile, err := f.Open()
    if err != nil {
        return err
    }
    defer zippedFile.Close()

    if _, err := io.Copy(destinationFile, zippedFile); err != nil {
        return err
    }
    return nil
}

func downloadZip(zipurl, us string){
	specUrl := zipurl
    resp, err := http.Get(specUrl)
    if err != nil {
        fmt.Printf("err: %s", err)
    }


    defer resp.Body.Close()
    fmt.Println("status", resp.Status)
    if resp.StatusCode != 200 {
        return
    }

    // Create the file
    out, err := os.Create(us)
    if err != nil {
        fmt.Printf("err: %s", err)
    }
    defer out.Close()

    // Write the body to file
    _, err = io.Copy(out, resp.Body)
    fmt.Printf("\nerr: %s", err)
}

func Unzip(zipname string, uz string){
	err := unzipSource(zipname, uz)
    if err != nil {
        log.Fatal(err)
    }
}