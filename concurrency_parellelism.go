package main

import (
    "fmt"
    "runtime"
    "time"
    "sync"
    )

func worker1(wg *sync.WaitGroup){
        defer wg.Done()

    for i:=0;i<100000;i++{
        
    }
    fmt.Println("task 1...")
}

func worker2(wg *sync.WaitGroup){
    defer wg.Done()
        for i:=0;i<100000;i++{
        
    }
    fmt.Println("task 2...")
}


func main(){
    var wg sync.WaitGroup
    start := time.Now()
    //Parellism 
    runtime.GOMAXPROCS(runtime.NumCPU())
    for i:=0;i<100;i++{
    wg.Add(1)
    go worker1(&wg)
    wg.Add(1)
    go worker2(&wg)
    }
    
    wg.Wait()
    
    end := time.Since(start)
    fmt.Println(end)
    time.Sleep(time.Millisecond)
}