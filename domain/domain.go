package domain

import (
	"github.com/nafisalfiani/ketson-transaction-service/domain/transaction"
	"github.com/nafisalfiani/ketson-transaction-service/domain/wallet"
	xenditDom "github.com/nafisalfiani/ketson-transaction-service/domain/xendit"
	"github.com/nafisalfiani/ketson-transaction-service/lib/broker"
	"github.com/nafisalfiani/ketson-transaction-service/lib/log"
	"github.com/nafisalfiani/ketson-transaction-service/lib/parser"
	"github.com/xendit/xendit-go/v4"
	"gorm.io/gorm"
)

type Domains struct {
	Transaction transaction.Interface
	Wallet      wallet.Interface
	Xendit      xenditDom.Interface
}

func Init(logger log.Interface, json parser.JSONInterface, db *gorm.DB, broker broker.Interface, xnd *xendit.APIClient) *Domains {
	return &Domains{
		Transaction: transaction.Init(db),
		Wallet:      wallet.Init(db),
		Xendit:      xenditDom.Init(xnd, logger),
	}
}
