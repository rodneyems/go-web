package transactions

type Repository interface {
	GetAll() ([]transaction, error)
	Save(id int, currency string, issuer string, receiver string, date string, price float64) (transaction, error)
	LastId() (int, error)
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