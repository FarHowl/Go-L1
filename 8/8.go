package main

import "fmt"

// (1 << pos) returns number 1 in binary number system as 0001. So after OR operation 0+0=0 and 1/0+1=1
func setBit(number int64, pos uint) int64 {
	return number | (1 << pos)
}

// ^(1 << pos) returns 0001 and reverses it to 1110. So after AND operation 0*0/1=0 and 1*1=1
func clearBit(number int64, pos uint) int64 {
	var mask int64 = ^(1 << pos)
	number &= mask
	return number
}

func main() {
	fmt.Println(clearBit(129, 7))
	fmt.Println(setBit(1, 7))
}
