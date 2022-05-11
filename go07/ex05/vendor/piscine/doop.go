package piscine

func Doop() {
	v1, v2, opeFunc, ok := readOpeArgs()
	if !ok {
		return
	}

	result, err := opeFunc(v1, v2)
	if err != nil {
		switch err.(type) {
		case divZeroError, modZeroError:
			printStrEndl(err.Error())
		}
		return
	}
	printNbrEndl(result)
}
