package pin

func ExamplePin() {
	Pin()
	// Output:
	// 123
	// 132
	// 213
	// 231
	// 312
	// 321
}

func ExampleGenerateAListOfNumbers() {
	Output := GenerateAListOfNumbers([]int{1, 2, 3})
	fmt.Println(Output)
	// Output: [123 132 231 213 312 321]
}
