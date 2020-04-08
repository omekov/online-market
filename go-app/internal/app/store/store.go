package store

// Store - сборка внешних репозиторий
type Store interface {
	Category() CategoryRepository
}