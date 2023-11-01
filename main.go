package main

import (
	"fmt"
	"goids/goids"
)

func main() {
	e := goids.CreateEnv(500, 500, 30, 4, 2)

	for i := 0; i < 1000; i++ {
		e.Update()
		fmt.Println(e.Render())
	}
}
