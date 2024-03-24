package utils

import (
    "archive/zip"
	"fmt"
    "io"
	"os"
	"path/filepath"
    "strings"
)

// create zip file from directory
// TEST: add tests
func ZipDir(source string, target string) error {
    // TODO: needs way of dealing with larger zip files (send one for each event?)

    // delete zip if it already exists
    if _, err := os.Stat(target); err == nil {
        err := os.Remove(target)
        if err != nil {
            fmt.Println("Error deleting zip file:", err)
            return err
        } else {
            fmt.Println("Existing zip file deleted successfully.")
        }
    } else if !os.IsNotExist(err) {
        fmt.Println("Error checking for existing zip file:", err)
        return err
    }

    zipfile, err := os.Create(target)
    if err != nil {
        return err
    }

    defer zipfile.Close() 

    archive := zip.NewWriter(zipfile)
    defer archive.Close()

    // walk through all files/dirs in source dir
    // FIX: error checking on this function?
    // TODO: also change to using filepath.WalkDir which is more efficient
    filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        // disallow zip and png files from being included in the new zip
        if filepath.Ext(path) == ".zip" || filepath.Ext(path) == ".png" {
            return nil
        }

        header, err := zip.FileInfoHeader(info)
        if err != nil {
            return err
        }

        // make file paths relative
        header.Name = filepath.ToSlash(path[len(source):])

        // remove any leading slashes, which seem to cause issues opening zip on windows
        header.Name = strings.TrimPrefix(header.Name, "/")

        // append slash for directories
        if info.IsDir() {
            header.Name += "/"
        } else {
            header.Method = zip.Deflate
        }

        writer, err := archive.CreateHeader(header)
        if err != nil {
            return err
        }

        // write files to zip
        if !info.IsDir() {
            file, err := os.Open(path)
            if err != nil {
                return err
            }
            defer file.Close()
            _, err = io.Copy(writer, file)
        }

        return err
    })

    return nil
}
