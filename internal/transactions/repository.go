package transactions

import (
	"errors"
	"fmt"

	"github.com/rodneyems/go-web/pkg/store"
)

type Repository interface {
	GetAll() ([]transaction, error)
	Save(id int, currency string, issuer string, receiver string, date string, price float64) (transaction, error)
	Update(t transaction) (transaction, error)
	LastId() (int, error)
	Delete(id int) error
	UpdateFields(id int, issuer string, price float64) (transaction, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]transaction, error) {
	ts := []transaction{}
	err := r.db.Read(&ts)
	fmt.Println(ts)
	if err != nil {
		return nil, err
	}
	return ts, nil
}

func (r *repository) Save(id int, currency string, issuer string, receiver string, date string, price float64) (transaction, error) {
	t := transaction{id, currency, price, issuer, receiver, date}
	ts := []transaction{}
	r.db.Read(&ts)
	ts = append(ts, t)
	r.db.Write(ts)
	return t, nil
}

func (r *repository) LastId() (int, error) {
	ts := []transaction{}
	r.db.Read(&ts)
	maxId := ts[0].Id
	for i := 1; i <= len(ts)-1; i++ {
		if ts[i].Id > maxId {
			maxId = ts[i].Id
		}
	}
	return maxId + 1, nil
}

func (r *repository) Update(t transaction) (transaction, error) {
	transactions := []transaction{}
	r.db.Read(&transactions)
	for i := range transactions {
		if transactions[i].Id == t.Id {
			transactions[i] = t
			return t, nil
		}
	}
	return transaction{}, errors.New("id inválido")
}

func (r *repository) Delete(id int) error {
	transactions := []transaction{}
	r.db.Read(&transactions)
	for i := range transactions {
		if transactions[i].Id == id {
			transactions = append(transactions[:i], transactions[i+1:]...)
			r.db.Write(transactions)
			return nil
		}
	}
	return errors.New("id inválido")
}

func (r *repository) UpdateFields(id int, issuer string, price float64) (transaction, error) {
	transactions := []transaction{}
	r.db.Read(&transactions)
	for i := range transactions {
		if transactions[i].Id == id {
			transactions[i].Issuer = issuer
			transactions[i].Price = price
			r.db.Write(transactions)
			return transactions[i], nil
		}
	}
	return transaction{}, errors.New("id inválido")
}
