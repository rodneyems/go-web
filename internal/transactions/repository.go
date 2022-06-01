package transactions

import (
	"errors"
)

type Repository interface {
	GetAll() ([]transaction, error)
	Save(id int, currency string, issuer string, receiver string, date string, price float64) (transaction, error)
	Update(t transaction) (transaction, error)
	LastId() (int, error)
	Delete(id int) error
	UpdateFields(id int, issuer string, price float64) (transaction, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]transaction, error) {
	return transactions, nil
}

func (r *repository) Save(id int, currency string, issuer string, receiver string, date string, price float64) (transaction, error) {
	t := transaction{id, currency, price, issuer, receiver, date}
	lastId = t.Id
	transactions = append(transactions, t)
	return t, nil
}

func (r *repository) LastId() (int, error) {
	return lastId, nil
}

func (r *repository) Update(t transaction) (transaction, error) {

	for i := range transactions {
		if transactions[i].Id == t.Id {
			transactions[i] = t
			return t, nil
		}
	}
	return transaction{}, errors.New("id inválido")
}

func (r *repository) Delete(id int) error {

	for i := range transactions {
		if transactions[i].Id == id {
			transactions = append(transactions[:i], transactions[i+1:]...)
			return nil
		}
	}
	return errors.New("id inválido")
}

func (r *repository) UpdateFields(id int, issuer string, price float64) (transaction, error) {

	for i := range transactions {
		if transactions[i].Id == id {
			transactions[i].Issuer = issuer
			transactions[i].Price = price
			return transactions[i], nil
		}
	}
	return transaction{}, errors.New("id inválido")
}
