package config

import (
	"os"
	"reflect"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var GlobalEnv Env

type Env struct {
	ServiceName    string `mapstructure:"SERVICE_NAME"`
	Debugging      bool   `mapstructure:"DEBUGGING"`
	ContextTimeout int    `mapstructure:"CONTEXT_TIMEOUT"`
	HTTPPort       string `mapstructure:"HTTP_PORT" json:"HTTP_PORT"`
	RPCPort        string `mapstructure:"RPC_PORT" json:"RPC_PORT"`

	RPCPaymentPort string `mapstructure:"GRPC_PAYMENT_PORT"`

	OtelHttpExporter string `mapstructure:"OTEL_HTTP_EXPORTER"`
	OtelGrpcExporter string `mapstructure:"OTEL_GRPC_EXPORTER"`

	DBReadHost   string `mapstructure:"DB_READ_HOST"`
	DBReadPort   string `mapstructure:"DB_READ_PORT"`
	DBReadName   string `mapstructure:"DB_READ_NAME"`
	DBReadUser   string `mapstructure:"DB_READ_USER"`
	DBReadPass   string `mapstructure:"DB_READ_PASS"`
	DBReadSchema string `mapstructure:"DB_READ_SCHEMA"`

	DBWriteHost   string `mapstructure:"DB_WRITE_HOST"`
	DBWritePort   string `mapstructure:"DB_WRITE_PORT"`
	DBWriteName   string `mapstructure:"DB_WRITE_NAME"`
	DBWriteUser   string `mapstructure:"DB_WRITE_USER"`
	DBWritePass   string `mapstructure:"DB_WRITE_PASS"`
	DBWriteSchema string `mapstructure:"DB_WRITE_SCHEMA"`

	RedisHost    string `mapstructure:"REDIS_HOST_PORT"`
	RedisPass    string `mapstructure:"REDIS_PASS"`
	RedisDB      string `mapstructure:"REDIS_DB"`
	RedisTimeOut int64  `mapstructure:"REDIS_TIMEOUT"`

	KafkaHost          string `mapstructure:"KAFKA_HOST"`
	KafkaConsumerGroup string `mapstructure:"KAFKA_CONSUMER_GROUP"`

	ElasticHost string `mapstructure:"ELASTIC_HOST"`
	ElasticUser string `mapstructure:"ELASTIC_USER"`
	ElasticPass string `mapstructure:"ELASTIC_PASS"`
}

func LoadGlobalEnv(path string) (err error) {

	// _ = godotenv.Load(".env")

	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	// root := filepath.Dir(filepath.Dir(wd))

	err = godotenv.Load(wd + "/.env")
	if err != nil {
		return err
	}

	globalEnvType := reflect.TypeOf(Env{})
	for i := 0; i < globalEnvType.NumField(); i++ {
		field := globalEnvType.Field(i)
		fieldTag := field.Tag.Get("mapstructure")
		if fieldTag != "" {
			viper.BindEnv(fieldTag)
		}
	}

	viper.Unmarshal(&GlobalEnv)
	return nil
}
