package main

import (
	"fmt"
)

const testVersion = 6

func Distance(x, y string) (int, error) {
	a := []byte(x)
	b := []byte(y)
	lenOfstring := len(a)
	lenOfb := len(b)
	missMatches := make([]int,0)

	for i:=0; i<lenOfstring; i++ {
		if a[i] != b[i]{
			missMatches = append(missMatches, 1)
		}
	}
	numMissMatches := len(missMatches)

	if lenOfstring != lenOfb {
			return error("cool")
		} else if numMissMatches != 0{
			return numMissMatches, nil
		} else {
			return 0, nil
		}
}

func main()  {
	z:=Distance("a", "b")
}