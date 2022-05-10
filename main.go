package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

func main() {

	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		memoryLeaking()
		c.String(http.StatusOK, "pong")
	})

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

func memoryLeaking() {
	var wg sync.WaitGroup
	for {
		// spawn four worker goroutines
		spawnWorkers(4, wg)
		// wait for the workers to finish
		wg.Wait()
	}
}

func spawnWorkers(max int, wg sync.WaitGroup) {
	for n := 0; n < max; n++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			f(n)
			return
		}()
	}
}

func f(n int) {
	for i := 0; i < 1000; i++ {
		fmt.Println(n, ":", i)
	}
}
