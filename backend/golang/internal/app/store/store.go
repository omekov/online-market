package store

// Storer - сборка внешних репозиторий
type Store interface {
	Category() CategoryRepositorer
	Product() ProductRepositorer
	Customer() CustomerRepositorer
}
