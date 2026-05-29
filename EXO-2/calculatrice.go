package main

import "fmt"

func creerOperation(op string) func(float64, float64) float64 {
	switch op {
	case "+":
		return func(a, b float64) float64 { return a + b }
	case "-":
		return func(a, b float64) float64 { return a - b }
	case "*":
		return func(a, b float64) float64 { return a * b }
	case "/":
		return func(a, b float64) float64 { return a / b }
	default:
		return nil
	}
}

func operer(a, b float64, op string) (float64, error) {
	if op == "/" && b == 0 {
		return 0, fmt.Errorf("division par zero")
	}

	fn := creerOperation(op)
	if fn == nil {
		return 0, fmt.Errorf("operation invalide: %s", op)
	}

	return fn(a, b), nil
}

func main() {
	for {
		var a, b float64
		var op string

		fmt.Scan(&a, &b, &op)

		if op == "quit" {
			break
		}

		result, err := operer(a, b, op)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(result)
	}
}
