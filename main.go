package main

import (
	"Lesson6/configReader"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
)

const count = 10
var GLOBAL_ConfigParsingErr = errors.New("Config data validation error")

func main() {
	var cfg,err = configReader.NewConfig()

	trace.Start(os.Stderr)
	defer trace.Stop()
	
	if !err {
		switch cfg.ProgramNum {
			case 1:fmt.Println("Program1");protectdataSample()
			case 2: fmt.Println("Program2");implicitSchedulercallSample()
			case 3: fmt.Println("Program3");raceConditionSample()
		}
	}else{
		errPars := fmt.Errorf("config pars error in main: %w", GLOBAL_ConfigParsingErr)
		fmt.Println(errPars.Error())
	}

}

func protectdataSample() {
	var ProtectedData struct {
		mutex        sync.Mutex
		varToProtect int
	}
	for r := 0; r < count; r++ {
		go func() {
			ProtectedData.mutex.Lock()
			ProtectedData.varToProtect = rand.Intn(5)
			ProtectedData.mutex.Unlock()
		}()
		time.Sleep(1 * time.Second)
	}
}

func implicitSchedulercallSample() {
	c1 := make(chan int,100)
	for i := 0; i < 100 ; i++ {
		go func() {
			testvar := rand.Intn(5)
			c1 <-testvar
			if testvar == 5 {
				runtime.Gosched()
			}
		}()
	}

}

func raceConditionSample() {
	value := 0
	for i := 0; i < 1000000; i++ {
		go func() {
			value += 1
		}()
	}
	fmt.Printf("%d\n", value)
}
