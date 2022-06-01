package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"time"
)

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		memoryLeaking()
		c.String(http.StatusOK, "ok")
	})

	// Listen and Server in 0.0.0.0:8080
	err := r.Run(":8080")
	if err != nil {
		log.Fatal("failed to start webapp")
	}
}

func memoryLeaking() {
	f, err := os.Open(os.DevNull)
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	max := runtime.NumCPU()
	chosen := rand.Intn(max-1) + 1
	runtime.GOMAXPROCS(chosen)

	for i := 0; i < chosen; i++ {
		go func() {
			for {
				_, _ = fmt.Fprintf(f, ".")
			}
		}()
	}

	time.Sleep(10 * time.Second)
}
