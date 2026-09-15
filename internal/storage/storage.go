package storage

type Storage interface {
	CreateStudent(name string, age uint8, class string, email string) (int64, error)
}
