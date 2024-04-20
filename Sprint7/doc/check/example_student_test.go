package check

import "fmt"

func ExampleStudent_SetName() {

	student := Student{}

	student.SetName("серега")
	fmt.Println(student.GetName())

	student.SetName("васек")
	fmt.Println(student.GetName())

	student.SetName("илюха")
	fmt.Println(student.GetName())

	// Output:
	// Серега
	// Васек
	// Илюха
}
