package main

import "fmt"

// func Hello(name string) string {
// 	message := fmt.Sprintf("Ola %v. oia", name)
// 	return message
// }

// func main() {
// 	msg := Hello("Fernanda")
// 	fmt.Println(msg)
// }

func main() {
	var a int = 10
	var b float64 = 12.99

	fmt.Printf(" %d ", a)
	fmt.Printf(" %.2f ", b)

	soma := float64(a) + b

	fmt.Printf("%.2f", soma)
}
