package main

import (
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/nafisalfiani/ketson-transaction-service/config"
	"github.com/nafisalfiani/ketson-transaction-service/domain"
	"github.com/nafisalfiani/ketson-transaction-service/entity"
	"github.com/nafisalfiani/ketson-transaction-service/handler/grpc"
	"github.com/nafisalfiani/ketson-transaction-service/lib/auth"
	"github.com/nafisalfiani/ketson-transaction-service/lib/broker"
	"github.com/nafisalfiani/ketson-transaction-service/lib/configreader"
	"github.com/nafisalfiani/ketson-transaction-service/lib/log"
	"github.com/nafisalfiani/ketson-transaction-service/lib/parser"
	"github.com/nafisalfiani/ketson-transaction-service/lib/security"
	"github.com/nafisalfiani/ketson-transaction-service/lib/sql"
	"github.com/nafisalfiani/ketson-transaction-service/usecase"
	"github.com/xendit/xendit-go/v4"
)

func main() {
	log.DefaultLogger().Info(context.Background(), "starting server...")

	// init config file
	cfg := configreader.Init(configreader.Options{
		Type:       configreader.Viper,
		ConfigFile: "./config.json",
	})

	// read from config
	config := &config.Application{}
	if err := cfg.ReadConfig(config); err != nil {
		log.DefaultLogger().Fatal(context.Background(), err)
	}

	log.DefaultLogger().Info(context.Background(), config)

	// init logger
	logger := log.Init(config.Log, log.Zerolog)

	// init parser
	parser := parser.InitParser(logger, parser.Options{})

	// init validator
	validator := validator.New(validator.WithRequiredStructEnabled())

	// init security
	sec := security.Init(logger, config.Security)

	// init database connection
	db, err := sql.Init(config.Sql, logger)
	if err != nil {
		logger.Fatal(context.Background(), err)
	}

	if err := db.AutoMigrate(
		&entity.Transaction{},
		&entity.Wallet{},
		&entity.WalletHistory{},
	); err != nil {
		logger.Fatal(context.Background(), err)
	}

	// init broker
	broker, err := broker.Init(config.Broker, logger, parser.JSONParser())
	if err != nil {
		logger.Fatal(context.Background(), err)
	}
	defer broker.Close()

	// init auth
	auth := auth.Init(config.Auth, logger, parser.JSONParser(), http.DefaultClient)

	xnd := xendit.NewClient(config.Xendit.ApiKey)

	// init domain
	dom := domain.Init(logger, parser.JSONParser(), db, broker, xnd)

	// init usecase
	uc := usecase.Init(logger, dom, broker)

	// init grpc
	grpc := grpc.Init(config.Grpc, logger, uc, sec, auth, validator)

	// start grpc server
	grpc.Run()
}
