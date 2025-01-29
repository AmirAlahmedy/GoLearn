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
var ErrPower = fmt.Errorf("can't boild water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
} 

func main() {
	for _, i := range[]int{7, 42} {
		if r, e := f(i); e != nil {
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("We should buy tea")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Now it's dark")
			} else {
				fmt.Printf("Unkown error: %s\n", err)
			}
			continue
		}

		fmt.Println("Made tea:", i)
	}
}
