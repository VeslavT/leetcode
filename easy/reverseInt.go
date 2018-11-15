//Given a 32-bit signed integer, reverse digits of an integer.
//
//Example 1:
//
//Input: 123
//Output: 321
//Example 2:
//
//Input: -123
//Output: -321
//Example 3:
//
//Input: 120
//Output: 21
//Note:
//Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1].
// For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
package main

import (
	"fmt"
	"strconv"
	"math"
	"time"
)

func reverse(d int) int{

	if d > -10 && d < 10{
		return d
	}
	stringD := strconv.Itoa(d)
	if d < 0 {
		stringD = stringD[1:]
	}
	var r string
	for i := len(stringD) - 1; i >= 0 ; i-- {
		r += string(stringD[i])
	}
	intD, err := strconv.Atoi(r)
	if err != nil{
		return 0
 	}
	if d > math.MaxInt32 || d < math.MinInt32{
		return 0
	}
	if d < 0 {
		return -intD
	}
	return intD
}

func reverseCopy(d int) int{
	if d > -10 && d < 10{
		return d
	}
	stringD := strconv.Itoa(d)
	if d < 0 {
		stringD = stringD[1:]
	}
	bs := make([]byte, len(stringD))
	bl := 0

	for i := len(stringD) - 1; i >= 0 ; i-- {
		bl += copy(bs[bl:], string(stringD[i]))
	}

	intD, err := strconv.Atoi(string(bs))

	if err != nil{
		return 0
	}
	if d < 0 {
		intD = -intD
	}
	if intD > math.MaxInt32 || intD < math.MinInt32{
		return 0
	}
	return intD
}

func reverseInt(d int) int {
	if d > -10 && d < 10{
		return d
	}
	n := absInt(d)
	newInt := 0
	for n > 0 {
		remainder := n % 10
		newInt *= 10
		newInt += remainder
		n /= 10
	}
	if d < 0 {
		newInt = -newInt
	}
	if newInt > math.MaxInt32 || newInt < math.MinInt32{
		return 0
	}
	return newInt
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}


func main(){
	d := -12345
	start1 := time.Now()
	fmt.Println(reverse(d))
	fmt.Println(time.Since(start1))

	d1 := -12345
	start11 := time.Now()
	fmt.Println(reverseCopy(d1))
	fmt.Println(time.Since(start11))

	d2 := -12345
	start12 := time.Now()
	fmt.Println(reverseInt(d2))
	fmt.Println(time.Since(start12))

}
