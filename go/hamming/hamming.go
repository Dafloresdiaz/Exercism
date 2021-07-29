package hamming

import "errors"

func Distance(a, b string) (int, error) {
	differences := 0
	if len(a) != len(b) {
		return -1, errors.New("the length is not the same")
	}
	for iteration := 0; iteration < len(a); iteration++ {
		if a[iteration] != b[iteration] {
			differences += 1
		}
	}

	return differences, nil
}
