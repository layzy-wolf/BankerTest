package service

import (
	"context"
	"errors"
	"log"
)

type AccountInterface interface {
	CreateAccount() int
	Deposit(id int, amount float64) error
	Withdraw(id int, amount float64) error
	GetBalance(id int) float64
}

type AccountService struct {
	ctx            *context.Context
	accounts       map[int]*Account
	depositChan    chan account
	dChan          chan error
	withdrawChan   chan account
	wChan          chan error
	getBalanceChan chan account
	gChan          chan float64
}

type account struct {
	id     int
	amount float64
}

func NewAccountService(ctx *context.Context) *AccountService {
	accounts := make(map[int]*Account, 100)
	depositChan := make(chan account)
	dChan := make(chan error)
	withdrawChan := make(chan account)
	wChan := make(chan error)
	getBalanceChan := make(chan account)
	gChan := make(chan float64)

	accountService := AccountService{
		ctx:            ctx,
		accounts:       accounts,
		depositChan:    depositChan,
		dChan:          dChan,
		withdrawChan:   withdrawChan,
		wChan:          wChan,
		getBalanceChan: getBalanceChan,
		gChan:          gChan,
	}
	go Deposit(&accountService)
	go Withdraw(&accountService)
	go GetBalance(&accountService)

	return &accountService
}

func (s *AccountService) CreateAccount() int {
	log.Printf("INFO create user")
	id := len(s.accounts) + 1
	s.accounts[id] = &Account{
		Id:      id,
		Balance: 0,
	}
	return id
}

func (s *AccountService) Deposit(id int, amount float64) error {
	log.Printf("INFO deposit user %v amount %v", id, amount)
	s.depositChan <- account{
		id:     id,
		amount: amount,
	}
	err := <-s.dChan
	if err != nil {
		log.Println("INFO ", err)
	}
	return err
}

func (s *AccountService) Withdraw(id int, amount float64) error {
	log.Printf("INFO withdraw user %v amount %v", id, amount)
	s.withdrawChan <- account{
		id:     id,
		amount: amount,
	}
	err := <-s.wChan
	if err != nil {
		log.Println("INFO ", err)
	}
	return err
}

func (s *AccountService) GetBalance(id int) float64 {
	log.Printf("INFO balance user %v", id)
	s.getBalanceChan <- account{
		id:     id,
		amount: 0,
	}
	balance := <-s.gChan
	return balance
}

func Deposit(service *AccountService) {
	ctx := *service.ctx
	for {
		select {
		case ac := <-service.depositChan:
			account, ok := service.accounts[ac.id]
			if !ok {
				service.dChan <- errors.New("undefined user")
				continue
			}
			err := account.Deposit(ac.amount)
			service.dChan <- err
		case <-ctx.Done():
			return
		}
	}
}

func Withdraw(service *AccountService) {
	ctx := *service.ctx
	for {
		select {
		case ac := <-service.withdrawChan:
			account, ok := service.accounts[ac.id]
			if !ok {
				service.wChan <- errors.New("undefined user")
				continue
			}
			err := account.Withdraw(ac.amount)
			service.wChan <- err
		case <-ctx.Done():
			return
		}
	}
}

func GetBalance(service *AccountService) {
	ctx := *service.ctx
	for {
		select {
		case ac := <-service.getBalanceChan:
			account, ok := service.accounts[ac.id]
			if !ok {
				service.gChan <- 0
				continue
			}
			balance := account.GetBalance()
			service.gChan <- balance
		case <-ctx.Done():
			return
		}
	}
}
