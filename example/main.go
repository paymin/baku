package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"

	"github.com/iDevoid/baku"
)

func main() {
	runtime.GC()
	fmt.Println("starting initialize baku dictionary..")
	fmt.Println(baku.InitializeDictionary())
	fmt.Println("baku dictionary populated")
	PrintMemUsage()

	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Enter text: ")
		scanner.Scan()
		text := scanner.Text()
		fmt.Println("kata baku: ", baku.Formalize(text))
		fmt.Print("\n\n")
		PrintMemUsage()
	}
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
