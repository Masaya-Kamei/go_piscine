package piscine

const intMax = int(^uint(0) >> 1)
const intMin = -intMax - 1

type overFlowError struct{}
type divZeroError struct{}
type modZeroError struct{}

func (e overFlowError) Error() string { return "" }
func (e divZeroError) Error() string  { return "No division by 0" }
func (e modZeroError) Error() string  { return "No modulo by 0" }

func opePlus(v1, v2 int) (int, error) {
	if v1 > 0 && v2 > 0 && v1+v2 <= 0 ||
		v1 < 0 && v2 < 0 && v1+v2 >= 0 {
		return 0, overFlowError{}
	}
	return v1 + v2, nil
}

func opeMinus(v1, v2 int) (int, error) {
	if v1 == 0 && v2 == intMin ||
		v1 > 0 && v2 < 0 && v1-v2 <= 0 ||
		v1 < 0 && v2 > 0 && v1-v2 >= 0 {
		return 0, overFlowError{}
	}
	return v1 - v2, nil
}

func opeMulti(v1, v2 int) (int, error) {
	if v2 == -1 && v1 == intMin ||
		v2 != 0 && v1*v2/v2 != v1 {
		return 0, overFlowError{}
	}
	return v1 * v2, nil
}

func opeDiv(v1, v2 int) (int, error) {
	if v2 == 0 {
		return 0, divZeroError{}
	} else if v2 == -1 && v1 == intMin {
		return 0, overFlowError{}
	}
	return v1 / v2, nil
}

func opeMod(v1, v2 int) (int, error) {
	if v2 == 0 {
		return 0, modZeroError{}
	} else if v2 == -1 && v1 == intMin {
		return 0, overFlowError{}
	}
	return v1 % v2, nil
}
