package main

import "fmt"

func main() {
	f, err := M2fConsole()
	if err != nil {
		fmt.Print("errorss: ")
		fmt.Println(err)
	} else {
		fmt.Println("fust ", FormatResult(f))
	}
}

const koef = 0.3048

func FormatResult(f float64) string {
	return fmt.Sprintf("%.4f", f)
}

func M2fConsole() (float64, error) {
	fmt.Print("Enter meters: ")

	input, err := getValueFromConsole()
	if err != nil {
		return 0, err
	}
	return M2fConvert(input), nil
}

func M2fConvert(m float64) float64 {
	return m / koef
}

func getValueFromConsole() (float64, error) {
	var input float64
	_, err := fmt.Scanf("%f", &input)
	if err != nil {
		return 0, err
	} else {
		return input, nil
	}
}
