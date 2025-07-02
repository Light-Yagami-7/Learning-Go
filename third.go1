package main

import (
	"fmt"
	"os"
	"strconv"
)

func cal(num1, num2 int, operand string) {
	if num2 == 0 && operand == "/" {
		fmt.Println("Dont even try to divide by 0")
	} else {

		//		for operand != "+" && operand != "-" && operand != "/" && operand != "*" {

		switch operand {
		case "+":
			fmt.Println("On adding:", num1+num2)

		case "-":
			fmt.Println("On subtracting:", num1-num2)

		case "*":
			fmt.Println("On multipying:", num1*num2)

		case "/":
			fmt.Println("On dividing:", num1/num2)

		default:
			fmt.Println("Pro Tip: Try using some brain")

		}
		//}
	}

}

var looprecholder int = 0

func loop1() {
	if looprecholder == 5 {
		fmt.Println("Error: User is too dumb")
		os.Exit(1)
		return
	}

	var tmp string
	fmt.Println("Your first number:")
	fmt.Scanln(&tmp)

	yn, err := strconv.Atoi(tmp)
	if err != nil {
		fmt.Println("Type a number")
		looprecholder++
		loop1()
		return
	}

	var tmp2 string
	fmt.Println("Your second nubmer")
	fmt.Scanln(&tmp2)

	yn2, err2 := strconv.Atoi(tmp2)
	if err2 != nil {
		looprecholder++
		fmt.Println("Enter a number")
		loop1()
		return
	}

	var operand string
	fmt.Println("What you want to do ?:")
	fmt.Scanln(&operand)

	cal(yn, yn2, operand)
}

func main() {
	for {

		loop1()
		var tryagain string
		fmt.Println("Calculate Again ?(y for yes/anything for no:")
		fmt.Scanln(&tryagain)
		if tryagain != "y" {
			fmt.Println("OkK done calculating.")
			return
		}
	}
}
