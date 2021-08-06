package calc

import "errors"

func Add(numbers ...int) (int, error) {
	sum := 0

	if len(numbers) < 2 {
		return sum, errors.New("please provide more than 2 numbers")
	} else {
		for _, num := range numbers {
			sum = sum + num
		}

		return sum, nil

	}
}
