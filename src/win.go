//go:build windows
// +build windows

package main

import (
	"syscall"
)

func init() {
}

func SetAttr() {
	var p uint32 = 0
	var c uint32 = 0
	p += 0x00000200                                               // CREATE_NEW_PROCESS_GROUP
	p += 0x00000008                                               // DETACHED_PROCESS
	c += 0x08000000                                               // CREATE_NO_WINDOW
	sp = &syscall.SysProcAttr{CreationFlags: p}                   // SysProcAttr_Parent
	sc = &syscall.SysProcAttr{CreationFlags: c, HideWindow: true} // SysProcAttr_Child
}
