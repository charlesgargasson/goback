package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"sync"
	"syscall"
	"time"
)

var sp = &syscall.SysProcAttr{}
var sc = &syscall.SysProcAttr{}
var isLIB bool = false
var idle int = 0
var idlemax int = 50
var keepalivemsg string = "_KEEPALIVE_"

var (
	Info     *log.Logger
	Warning  *log.Logger
	Error    *log.Logger
	Critical *log.Logger
)

func GardnerFileExist(filePath string) bool {
	_, error := os.Stat(filePath)
	return !errors.Is(error, os.ErrNotExist)
}

func init() {
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	Critical = log.New(os.Stderr, "CRITICAL: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// ReadPipe
func RP(input io.ReadCloser, output io.WriteCloser, quit *bool, wg *sync.WaitGroup, keepalive bool) {
	defer wg.Done()
	reader := bufio.NewReader(input)
	inputline := ""
	isquit := make(chan bool)
	go func() {
		for {
			time.Sleep(time.Duration(100 * time.Millisecond))
			if *quit {
				isquit <- true
				break
			}
		}
	}()
	if keepalive {
		go func() {
			for {
				time.Sleep(time.Duration(1000 * time.Millisecond))
				idle++
				if idle > idlemax {
					fmt.Fprint(output, keepalivemsg)
					idle = 0
				}
			}
		}()
	}
	for {
		gotresult := make(chan bool)
		go func() {
			data, err := reader.ReadString('\n')
			if err != nil {
				*quit = true
				return
			} else {
				inputline = data
				gotresult <- true
			}
		}()
		select {
		case <-gotresult:
			idle = 0
			fmt.Fprint(output, inputline)
			time.Sleep(time.Duration(1 * time.Millisecond))
		case <-isquit:
			time.Sleep(time.Duration(100 * time.Millisecond))
			return
		}
	}
}

// WritePipe
func WP(input io.ReadCloser, output io.WriteCloser, quit *bool, wg *sync.WaitGroup, mode *int8) {
	defer wg.Done()
	reader := bufio.NewReader(input)
	inputline := ""
	isquit := make(chan bool)
	go func() {
		for {
			time.Sleep(time.Duration(100 * time.Millisecond))
			if *quit {
				isquit <- true
				break
			}
		}
	}()
	for {
		gotresult := make(chan bool)
		go func() {
			data, err := reader.ReadString('\n')
			if err != nil {
				*quit = true
				return
			} else {
				inputline = data
				gotresult <- true
			}
		}()
		select {
		case <-gotresult:
			inputline += "\n"
			re := regexp.MustCompile(`(\r\n?|\n){2,}`)
			inputline := re.ReplaceAllString(inputline, "$1")

			switch {
			// "gokillme"
			case inputline == `g`+x[0]+`k`+x[4]+`llm`+x[3]+"\n":
				*mode = 1
			// "goreset"
			case inputline == `g`+x[0]+`r`+x[3]+`s`+x[3]+`t`+"\n":
				*mode = 2
			// "gocmd"
			case inputline == `g`+x[0]+`cmd`+"\n":
				*mode = 3
			}

			if *mode != 0 {
				*quit = true
			} else {
				fmt.Fprint(output, inputline)
				time.Sleep(time.Duration(10 * time.Millisecond))
			}
		case <-isquit:
			time.Sleep(time.Duration(100 * time.Millisecond))
			return
		}
	}
}

func StdMapper(origin_out io.ReadCloser, origin_err io.ReadCloser, origin_in io.WriteCloser, dst_out io.WriteCloser, dst_err io.WriteCloser, dst_in io.ReadCloser, quit *bool, wg *sync.WaitGroup, mode *int8, manualkeepalive bool) {
	wg.Add(3)
	go RP(origin_out, dst_out, quit, wg, manualkeepalive)
	go RP(origin_err, dst_err, quit, wg, false)
	go WP(dst_in, origin_in, quit, wg, mode)
}
