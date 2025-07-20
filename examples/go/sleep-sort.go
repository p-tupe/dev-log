package main

import (
"fmt"
"sync"
"time"
)

func main() {
var wg sync.WaitGroup
arr := []int{1, 5, 9, 7, 2, 8, 4, 3, 6}

    for _, n := range arr {
    	wg.Add(1)
    	go (func() {
    		time.Sleep(time.Duration(n) * time.Second)
    		fmt.Println(n)
    		wg.Done()
    	})()
    }

    wg.Wait()
    fmt.Println("Voila!")

}
