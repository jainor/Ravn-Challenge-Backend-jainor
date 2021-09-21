package db

// DBManager is an interface that must be implemented by each used db manager
type DBManager interface {
	Name() string
	ConnectionStr() string
	QueryStr() string
}
