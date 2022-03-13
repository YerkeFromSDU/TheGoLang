package main

import "fmt"

func main(){
	fmt.Println("hello")
	fmt.Println(Fizzbuzz(145))
	fmt.Println(IsPrime(5))
	fmt.Print(IsPalindrome(3445))
}
func Fizzbuzz(i int) string {

		var str = ""
		if i%3 == 0 && i%5 == 0 {
			str = "FizzBuzz"
		}
		if i%3 == 0 {
            // Multiple of 3
            str = "fizz"
        }
        if i%5 == 0 {
            // Multiple of 5
            str = "buzz"
        }

        if i%3 != 0 && i%5 != 0 {
            str = ""
        }
		return str
}
func IsPrime(n int) bool {
	var b = false
	if n%n == 0 && n%1 == 0 {
		b = true
	}
	return b
}
func IsPalindrome(palNum int) bool {
	var remainder int

    reverse := 0

    for temp := palNum; temp > 0; temp = temp / 10 {
        remainder = temp % 10
        reverse = reverse*10 + remainder
    }
	
	var b = false
    if palNum == reverse {
        b = true
    }
	return b
}

