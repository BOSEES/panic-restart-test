package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("panic test start")
	time.Sleep(5 * time.Second)
	panic("panic test!!!!!!!!!")
}
