package golangexamples

import "github.com/ehteshamz/greetings"

func EZGreetings(name string) string {
	return greetings.PrintGreetings(name)
}

func ConcatSlice(sliceToConcat []byte) string {
	var to_return = ""

	for _, value := range sliceToConcat {
		to_return = to_return + string(value) + "-"
	}
	return to_return
}

func Encrypt(sliceToEncrypt []byte, ceaserCount int) {
	for i, value := range sliceToEncrypt {
		sliceToEncrypt[i] = (byte((int(value) + ceaserCount) % 256))
	}

}
