package piscine

import (
	"os"
)

const BUFSIZE = 32

func catStdin() {
	var buf [BUFSIZE]byte
	line := ""
	for {
		n, err := os.Stdin.Read(buf[:])
		if n == 0 {
			break
		} else if err != nil {
			printStrEndl("ERROR: " + err.Error())
			os.Exit(1)
		}
		line += string(buf[:n])
		if buf[n-1] == '\n' {
			printStr(line)
			line = ""
		}
	}
	os.Exit(0)
}

func isExistArgs() bool {
	for range os.Args[1:] {
		return true
	}
	return false
}

func Cat() {
	if !isExistArgs() {
		catStdin()
		return
	}
	exitStatus := 0
	for _, arg := range os.Args[1:] {
		bytes, err := os.ReadFile(arg)
		if err != nil {
			printStrEndl("ERROR: " + err.Error())
			exitStatus = 1
		}
		printStr(string(bytes))
	}
	os.Exit(exitStatus)
	return
}
