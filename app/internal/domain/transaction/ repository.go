package transaction

type Repository interface {
	AddTransaction(trans Transaction) error
	GetBalance() (float64, error)
	GetTransactions() ([]TransactionWithCategory, error)
}
