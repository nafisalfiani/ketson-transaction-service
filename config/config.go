package config

import (
	"github.com/nafisalfiani/ketson-go-lib/auth"
	"github.com/nafisalfiani/ketson-go-lib/broker"
	"github.com/nafisalfiani/ketson-go-lib/cache"
	"github.com/nafisalfiani/ketson-go-lib/log"
	"github.com/nafisalfiani/ketson-go-lib/security"
	"github.com/nafisalfiani/ketson-go-lib/sql"
	"github.com/nafisalfiani/ketson-transaction-service/handler/grpc"
)

type Application struct {
	Auth     auth.Config     `env:"AUTH"`
	Log      log.Config      `env:"LOG"`
	Security security.Config `env:"SECURITY"`
	Sql      sql.Config      `env:"SQL"`
	Cache    cache.Config    `env:"CACHE"`
	Broker   broker.Config   `env:"BROKER"`
	Grpc     grpc.Config     `env:"GRPC"`
	Xendit   Xendit
}

type Xendit struct {
	ApiKey string
}
