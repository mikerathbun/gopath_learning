package main

import (
	"errors"
	"fmt"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Trapped panic: %s (%T)\n", err, err)
		}
	}()
	fmt.Println("Before yikes()")
	yikes()

	fmt.Println("After yikes()")

}
func yikes() {
	panic(errors.New("Something bad happened"))
}
