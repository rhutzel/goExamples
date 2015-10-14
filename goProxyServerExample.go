package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
	"github.com/rhutzel/goProxyServer"
)

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(20)

	for idx := 0; idx < 20; idx++ {
		go func() {
			body, _ := proxyServerExample.RequestContent("http://localhost")
			fmt.Println(strings.Replace(string(body)[:100], "\n", " ", -1) + "...")
			waitGroup.Done()
		}()
		time.Sleep(25 * time.Millisecond)
	}

	waitGroup.Wait()

	fmt.Println("\nHanging around for a few moments to observe cleanup...\n\n\n\n\n")
	time.Sleep(5 * time.Second)
}
