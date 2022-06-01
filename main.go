package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"os"
	"runtime"
)

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		leakyEndpoint()
		c.String(http.StatusOK, "ok")
	})

	// Listen and Server in 0.0.0.0:8080
	err := r.Run(":8080")
	if err != nil {
		log.Fatal("failed to start webapp")
	}
}

func leakyEndpoint() {
	f, err := os.Open(os.DevNull)
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	max := runtime.NumCPU()

	if max == 1 {
		max = 2
	}
	log.Printf("max: %d\n", max)
	chosen := rand.Intn(max-1) + 1
	runtime.GOMAXPROCS(chosen)
	log.Printf("chosen: %d\n", chosen)

	for i := 0; i < chosen; i++ {
		go func() {
			leakMemory()
			for {
				_, _ = fmt.Fprintf(f, ".")
			}
		}()
	}
}

func leakMemory() {
	type T struct {
		v [1 << 20]int
		t *T
	}

	var finalizer = func(t *T) {
		fmt.Println("finalizer called")
	}

	var x, y T

	// The SetFinalizer call makes x escape to heap.
	runtime.SetFinalizer(&x, finalizer)

	// The following line forms a cyclic reference
	// group with two members, x and y.
	// This causes x and y are not collectable.
	x.t, y.t = &y, &x // y also escapes to heap.
}
