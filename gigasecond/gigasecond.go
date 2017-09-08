package main

// import path for the time package from the standard library
import "time"
import "fmt"

const testVersion = 4

//func AddGigasecond(time.Time) time.Time

func main() {
	time.Parse(layout string, value string)
	x := time.Now()
	fmt.Println("x", x)
	y := x.Unix()
	fmt.Println("y", y)

}
