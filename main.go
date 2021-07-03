package main

import (
	"./configReader"
	"fmt"
	"errors"
)

var GLOBAL_ConfigParsingErr = errors.New("Config data validation error")


func main()  {
	var cfg,err = configReader.NewConfig()

	if !err {
		fmt.Println(cfg)
	}else{
		errPars := fmt.Errorf("config pars error in main: %w", GLOBAL_ConfigParsingErr)
		fmt.Println(errPars.Error())
	}

}
