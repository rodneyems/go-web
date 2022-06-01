package transactions

import (
	"errors"
	"fmt"
)

type Service interface {
	GetAll() ([]transaction, error)
	Save(currency string, issuer string, receiver string, date string, price float64) (transaction, error)
	Update(id int, currency string, issuer string, receiver string, date string, price float64) (transaction, error)
	Delete(id int) error
	UpdateFields(id int, issuer string, price float64) (transaction, error)
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
	t, _ := s.repository.Save(id, currency, issuer, receiver, date, price)
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

func (s *service) Delete(id int) error {
	fmt.Println("entrei no service")
	err := s.repository.Delete(id)
	if err != nil {
		return errors.New("id inválido")
	}
	return nil
}

func (s *service) UpdateFields(id int, issuer string, price float64) (transaction, error) {
	tUpdated, err := s.repository.UpdateFields(id, issuer, price)
	if err != nil {
		return tUpdated, errors.New(err.Error())
	}
	return tUpdated, nil
}
