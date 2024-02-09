package config

import (
	"github.com/nafisalfiani/ketson-transaction-service/handler/grpc"
	"github.com/nafisalfiani/ketson-transaction-service/lib/auth"
	"github.com/nafisalfiani/ketson-transaction-service/lib/broker"
	"github.com/nafisalfiani/ketson-transaction-service/lib/cache"
	"github.com/nafisalfiani/ketson-transaction-service/lib/log"
	"github.com/nafisalfiani/ketson-transaction-service/lib/security"
	"github.com/nafisalfiani/ketson-transaction-service/lib/sql"
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
