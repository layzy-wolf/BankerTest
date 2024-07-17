package service

import "errors"

type BankAccount interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	GetBalance() float64
}

type Account struct {
	Id      int     `json:"id"`
	Balance float64 `json:"balance"`
}

func (s *Account) Deposit(amount float64) error {
	if amount > 0 {
		s.Balance += amount
		return nil
	}
	return errors.New("amount to deposit must be greater then 0")
}

func (s *Account) Withdraw(amount float64) error {
	if s.Balance-amount >= 0 {
		s.Balance -= amount
		return nil
	}
	return errors.New("insufficient funds in the account")
}

func (s *Account) GetBalance() float64 {
	return s.Balance
}
