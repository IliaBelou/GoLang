package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go sysExit(c)
	workersFunc()
}
func sysExit(c <- chan os.Signal) {
	{
		sig := <-signalChannel
		switch sig {
		case syscall.SIGTERM:
			os.Exit(1)
		}
	}
}


func workersFunc() {
	var number int = 0
	jobs := make(chan int)
	results := make(chan int)
	//Запускаем 1000 воркеров
	for w := 1; w <= 1000; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 1000; j++ {
		jobs <- number
		number = <-results
	}
	close(jobs)
	fmt.Println(number)
}

func worker(ID int, jobs <-chan int, results chan<-int) {
	for j := range jobs {
		results <- j +1
	}
}
