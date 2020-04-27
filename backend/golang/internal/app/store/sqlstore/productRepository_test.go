package sqlstore_test

import (
	"testing"

	"github.com/omekov/online-market/backend/golang/internal/app/model"
	"github.com/omekov/online-market/backend/golang/internal/app/store"
	"github.com/omekov/online-market/backend/golang/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestCategoryRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("categories")

	s := sqlstore.New(db)
	p := model.TestProduct(t)
	assert.NoError(t, s.Product().Create(p))
	assert.NotNil(t, p)
}

func TestCategoryRepository_GetByID(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("categories")

	s := sqlstore.New(db)
	c := model.TestProduct(t)
	err := s.Product().GetByID(1, c)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())
	s.Product().Create(c)
	err = s.Product().GetByID(1, c)
	assert.NoError(t, err)
	assert.NotNil(t, c)
}

// func TestCategoryRepository_GetAll(t *testing.T) {
// 	db, teardown := sqlstore.TestDB(t, databaseURL)
// 	defer teardown("categories")

// 	s := sqlstore.New(db)
// 	c := model.TestCategory(t)
// 	_, err := s.Product().GetAll()
// 	assert.EqualError(t, err, store.ErrRecordNotFound.Error())
// 	s.Product().Create(c)
// 	_, err = s.Product().GetAll()
// 	assert.NoError(t, err)
// 	assert.NotNil(t, c)
// }
func TestCategoryRepository_Update(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("categories")

	s := sqlstore.New(db)
	c := model.TestProduct(t)
	err := s.Product().Update(1000, c)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())
	s.Product().Create(c)
	err = s.Product().Update(1, c)
	assert.NoError(t, err)
	assert.NotNil(t, c)
}

func TestCategoryRepository_Delete(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("categories")

	s := sqlstore.New(db)
	c := model.TestProduct(t)
	err := s.Product().Delete(1)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())
	s.Product().Create(c)
	err = s.Product().Delete(1)
	assert.NoError(t, err)
	assert.NotNil(t, c)
}
