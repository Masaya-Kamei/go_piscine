package piscine

import "os"

type ArgMissingError struct{}
type ArgTooManyError struct{}

func (e ArgMissingError) Error() string {
	return "File name missing"
}

func (e ArgTooManyError) Error() string {
	return "Too many arguments"
}

func getArgc() int {
	i := 0
	for range os.Args {
		i++
	}
	return i
}

func validateArgc() error {
	argc := getArgc()
	if argc == 1 {
		return ArgMissingError{}
	} else if argc >= 3 {
		return ArgTooManyError{}
	}
	return nil
}
