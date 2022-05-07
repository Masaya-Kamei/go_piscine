package piscine

import "os"

type ArgcError string

func getArgc() int {
	i := 0
	for range os.Args {
		i++
	}
	return i
}

func (e ArgcError) Error() string {
	return string(e)
}

func validateArgc() error {
	argc := getArgc()
	if argc == 1 {
		return ArgcError("File name missing")
	} else if argc >= 3 {
		return ArgcError("Too many arguments")
	}
	return nil
}
