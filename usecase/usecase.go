package usecase

import (
	"github.com/nafisalfiani/ketson-go-lib/broker"
	"github.com/nafisalfiani/ketson-go-lib/log"
	"github.com/nafisalfiani/ketson-transaction-service/domain"
	"github.com/nafisalfiani/ketson-transaction-service/usecase/transaction"
	"github.com/nafisalfiani/ketson-transaction-service/usecase/wallet"
)

type Usecases struct {
	Transaction transaction.Interface
	Wallet      wallet.Interface
}

func Init(logger log.Interface, dom *domain.Domains, broker broker.Interface) *Usecases {
	return &Usecases{
		Transaction: transaction.Init(logger, dom.Transaction, dom.Xendit, broker),
		Wallet:      wallet.Init(dom.Wallet),
	}
}
