package configReader

import (
	"github.com/asaskevich/govalidator"
	"gopkg.in/yaml.v2"

	"os"
	"strconv"
)

type Config struct {
	Port        string
	DbUrl       string
	JaegerUrl   string
	SentryUrl   string
	KafkaBroker string
	AppId   string
	AppKey  string
}

func NewConfig(configPath string) (*Config, bool) {

	var dataValidErr bool
	config := &Config{}

	file, err := os.Open(configPath)
	if err != nil {
		dataValidErr = true
		return nil, dataValidErr
	}
	defer file.Close()
	dYaml := yaml.NewDecoder(file)

	if err := dYaml.Decode(&config); err != nil {
		dataValidErr = true
		return nil, dataValidErr
	}
	//Data validation
	dataValidErr = false

	var value, _ = strconv.ParseFloat(config.Port, 64)

	if  (value == 0) || !govalidator.IsNatural(value) || !govalidator.IsURL(config.DbUrl) ||
		!govalidator.IsURL(config.JaegerUrl) || !govalidator.IsURL(config.SentryUrl) ||
		!govalidator.IsAlphanumeric(config.AppId) ||
		!govalidator.IsAlphanumeric(config.AppKey) {
		dataValidErr = true


		return nil, dataValidErr
	}



	return config, dataValidErr
}

