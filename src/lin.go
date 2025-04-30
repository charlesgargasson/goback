//go:build linux
// +build linux

package main

import "syscall"

func SetAttr() {
	sp = &syscall.SysProcAttr{Setpgid: true}
	sc = &syscall.SysProcAttr{Setpgid: false}
}
