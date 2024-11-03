package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"
)

var n = flag.Int("n", 3*1e6, "Number of goroutines to create")

var ch = make(chan byte)
var counter = 0

func f() {
	counter++
	<-ch// Блокируем эту горутину
}

func main() {
	flag.Parse()
	if *n <= 0 {
		fmt.Fprintf(os.Stderr, "invalid number of goroutines")
		os.Exit(1)
	}

	// Ограничиваем количество свободных потоков ОС до 1
	runtime.GOMAXPROCS(1)

	// Делаем копию MemStats
	var m0 runtime.MemStats
	runtime.ReadMemStats(&m0)

	t0 := time.Now().UnixNano()
	for i := 0; i < *n; i++ {
		go f()
	}
	runtime.Gosched()
	t1 := time.Now().UnixNano()
	runtime.GC()

	// Делаем копию MemStats
	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)

	if counter != *n {
		fmt.Fprintf(os.Stderr, "failed to begin execution of all goroutines")
		os.Exit(1)
	}

	fmt.Printf("Number of goroutines: %d\n", *n)
	fmt.Printf("Per goroutine:\n")
	fmt.Printf("  Memory: %.2f bytes\n", float64(m1.Sys-m0.Sys)/float64(*n))
	fmt.Printf("  Time:   %f µs\n", float64(t1-t0)/float64(*n)/1e3)
}

