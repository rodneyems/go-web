package transactions

type Service interface {
	GetAll() ([]transaction, error)
	Save(currency string, issuer string, receiver string, date string, price float64) (transaction, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}	

func (s *service) GetAll() ([]transaction, error) {
	return transactions, nil
}

func (s *service) Save(currency string, issuer string, receiver string, date string, price float64) (transaction, error) {
	id, _ := s.repository.LastId()
	t := transaction{id, currency, price, issuer, receiver, date}
	lastId = t.Id
	transactions = append(transactions, t)
	return t, nil
}
