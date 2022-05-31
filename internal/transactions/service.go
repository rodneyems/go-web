package transactions

import "errors"

type Service interface {
	GetAll() ([]transaction, error)
	Save(currency string, issuer string, receiver string, date string, price float64) (transaction, error)
	Update(id int, currency string, issuer string, receiver string, date string, price float64) (transaction, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}	

func (s *service) GetAll() ([]transaction, error) {
	data, _ := s.repository.GetAll()
	return data, nil
}

func (s *service) Save(currency string, issuer string, receiver string, date string, price float64) (transaction, error) {
	id, _ := s.repository.LastId()
	id++
	t := transaction{id, currency, price, issuer, receiver, date}
	lastId = t.Id
	transactions = append(transactions, t)
	return t, nil
}

func (s *service) Update(id int, currency string, issuer string, receiver string, date string, price float64) (transaction, error) {
	t := transaction{id, currency, price, issuer, receiver, date}
	tUpdated, err := s.repository.Update(t)
	if err != nil {
		return tUpdated, errors.New(err.Error())
	}
	return tUpdated, nil
}