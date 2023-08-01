package calculator

var logMessage = "[LOG]"

// Version of the calculator
var Version = "1.0"

// Private function
func internalSum(number int) int {
	return number - 1
}

// Public function
// Sum two integer numbers
func Sum(number1, number2 int) int {
	return number1 + number2
}
