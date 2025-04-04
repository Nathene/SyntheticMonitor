package probesql

type Database interface {
	Connect() error
	Query(query string) (interface{}, error)
	Close() error
}
