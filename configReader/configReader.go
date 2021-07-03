package configReader

import (
	"flag"
	"github.com/asaskevich/govalidator"
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

func NewConfig() (*Config, bool) {

	var dataValidErr bool
	var portVal = flag.String("port","8080","Port Number")
	var dbUrlVal = flag.String("dbUrl","postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable","Data Base Url")
	var jaegerUrlVal = flag.String("jaegerUrl","http://jaeger:16686","Tracer tool url")
	var sentryUrlVal = flag.String("sentryUrl","http://sentry:9000","Bug monitoring url")
	var kafkaBrokerVal = flag.String("kafkaBroker","kafka:9092","Msg Broker")
	var AppIdVal = flag.String("AppId","testid","AppId")
	var AppKey = flag.String("AppKey","testkey","AppKey")
	flag.Parse()
	//Data validation
	dataValidErr = false

	var value, _ = strconv.ParseFloat(*portVal, 64)

	if  (value == 0) || !govalidator.IsNatural(value) || !govalidator.IsURL(*dbUrlVal) ||
		!govalidator.IsURL(*jaegerUrlVal) || !govalidator.IsURL(*sentryUrlVal) ||
		!govalidator.IsAlphanumeric(*AppIdVal) ||
		!govalidator.IsAlphanumeric(*AppKey) {
		dataValidErr = true
		return nil, dataValidErr
	}

	config := Config {
		Port: *portVal,
		DbUrl: *dbUrlVal,
		JaegerUrl: *jaegerUrlVal,
		SentryUrl:  *sentryUrlVal,
		KafkaBroker: *kafkaBrokerVal,
		AppId:  *AppIdVal,
		AppKey: *AppKey,
	}

	return &config, dataValidErr
}
