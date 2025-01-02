package main

import (
	"errors"
	"fmt"
)

	func f(arg int) (int, error) {
		if arg == 42 {
			return -1, errors.New("can't work with 42")
		}

		return arg + 3, nil
	}

var ErrOutOfTea = fmt.Errorf("out of tea")
