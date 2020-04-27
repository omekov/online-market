package store

// Storer - сборка внешних репозиторий
type Store interface {
	Product() ProductRepositorer
	Customer() CustomerRepositorer
	Cart() CartRepositorer
}
