package homework

import (
	"fmt"
	"log"
	"math"
)

// EquationQuadratic 4.2 Giải phương trình bậc 2
func EquationQuadratic() string {
	var a,b,c float64
	fmt.Println("Nhập hệ số bậc 2, a = ")
	if _, err := fmt.Scanln(&a); err != nil {
		log.Fatalln("Error in Inputting First Number: ", err.Error())
	}

	if a <= 0 {
		return fmt.Sprintf("Hệ số bậc 2 phải khác 0")
	}

	fmt.Println("Nhập hệ số bậc 1, b = ")
	if _, err := fmt.Scanln(&b); err != nil {
		log.Fatalln("Error in Inputting Second Number: ", err.Error())
	}

	fmt.Println("Nhập hằng số tự do, c = ")
	if _, err := fmt.Scanln(&c); err != nil {
		log.Fatalln("Error in Inputting Third Number: ", err.Error())
	}

	if a+b+c == 0 {
		return fmt.Sprintf("Phương trình có nghiệm x1 = %d, x2 = %.2f", 1, c/a)
	}

	if a-b+c == 0 {
		return fmt.Sprintf("Phương trình có nghiệm x1 = %d, x2 = %.2f", -1, -c/a)
	}

	delta := b*b - 4*a*c
	if delta > 0 {
		x1 := (-b + math.Sqrt(delta)) / (2*a)
		x2 := (-b - math.Sqrt(delta)) / (2*a)
		return fmt.Sprintf("Phương trình có nghiệm x1 = %.2f, x2 = %.2f", x1, x2)
	} else if delta == 0 {
		return fmt.Sprintf("Phương trình có nghiệm kép x1 = x2 = %.2f", -b/(2*a))
	}

	realNumber := -b / (2 * a)
	imaginaryNumber := math.Sqrt(-delta) / (2 * a)
	return fmt.Sprintf("Phương trình có nghiệm phức x1 = %.2f, x2 = %.2f", realNumber + imaginaryNumber, realNumber - imaginaryNumber)
}
