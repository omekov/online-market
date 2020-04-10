package sqlstore_test

import (
	"testing"

	"github.com/omekov/online-market/go-app/internal/app/model"
	"github.com/omekov/online-market/go-app/internal/app/store"
	"github.com/omekov/online-market/go-app/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestCategoryRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("categories")

	s := sqlstore.New(db)
	c := model.TestCategory(t)
	assert.NoError(t, s.Category().Create(c))
	assert.NotNil(t, c)
}

func TestCategoryRepository_GetByID(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("categories")

	s := sqlstore.New(db)
	c := model.TestCategory(t)
	err := s.Category().GetByID(1, c)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())
	s.Category().Create(c)
	err = s.Category().GetByID(1, c)
	assert.NoError(t, err)
	assert.NotNil(t, c)
}

func TestCategoryRepository_GetAll(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("categories")

	s := sqlstore.New(db)
	c := model.TestCategory(t)
	_, err := s.Category().GetAll()
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())
	s.Category().Create(c)
	_, err = s.Category().GetAll()
	assert.NoError(t, err)
	assert.NotNil(t, c)
}
func TestCategoryRepository_Update(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("categories")

	s := sqlstore.New(db)
	c := model.TestCategory(t)
	err := s.Category().Update(1000, c)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())
	s.Category().Create(c)
	err = s.Category().Update(1, c)
	assert.NoError(t, err)
	assert.NotNil(t, c)
}

func TestCategoryRepository_Delete(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("categories")

	s := sqlstore.New(db)
	c := model.TestCategory(t)
	err := s.Category().Delete(1)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())
	s.Category().Create(c)
	err = s.Category().Delete(1)
	assert.NoError(t, err)
	assert.NotNil(t, c)
}
