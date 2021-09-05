package homework

import (
	"fmt"
	"log"
)


// IsTriangle 4.1 Kiểm tra 3 cạnh tam giác
func IsTriangle() string {
	var a, b, c int
	fmt.Println("Enter Your First Number: ")
	if _, err := fmt.Scanln(&a); err != nil {
		log.Fatalln("Error in Inputting First Number: ", err.Error())
	}

	if a <= 0 {
		return fmt.Sprintf("Number must not negative or equal 0")
	}

	fmt.Println("Enter Your Second Number: ")
	if _, err := fmt.Scanln(&b); err != nil {
		log.Fatalln("Error in Inputting Second Number: ", err.Error())
	}

	if b <= 0 {
		return fmt.Sprintf("Number must not negative or equal 0")
	}

	fmt.Println("Enter Your Third Number: ")
	if _, err := fmt.Scanln(&c); err != nil {
		log.Fatalln("Error in Inputting Third Number: ", err.Error())
	}

	if c <= 0 {
		return fmt.Sprintf("Number must not negative or equal 0")
	}

	if a+b<c && a+c<b && b+c<a {
		return fmt.Sprintf("Not a Triangle")
	}
	return fmt.Sprintf("This is a Triangle")
}