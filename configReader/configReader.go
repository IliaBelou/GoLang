package configReader

import (
	"flag"
	"strconv"
)

type Config struct {
	ProgramNum        int64
}

func NewConfig() (*Config, bool) {

	var dataValidErr bool
	var ProgramNum = flag.String("programNum","1","Program to run")
	flag.Parse()
	//Data validation
	dataValidErr = false

	var value,err = strconv.ParseInt(*ProgramNum,0, 64)
	if err != nil {
		dataValidErr = true
		return nil, dataValidErr
	}
	config := Config {
		ProgramNum: value,
	}

	return &config, dataValidErr
}
