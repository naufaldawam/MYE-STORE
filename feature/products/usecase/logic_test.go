package usecase

import (
	"project3/group3/domain"
	"project3/group3/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInsertProduct(t *testing.T) {
	repo := new(mocks.ProductData)
	insertData := domain.Product{
		ID:           1,
		ProductName:  "adidas",
		ProductImage: "ini gambar",
		Price:        10000,
		Stock:        10,
		UserID:       1,
	}
	outputData := domain.Product{
		ID:           1,
		ProductName:  "adidas",
		ProductImage: "ini gambar",
		Price:        10000,
		Stock:        10,
		UserID:       1,
	}

	t.Run("success insert", func(t *testing.T) {
		repo.On("InsertProductDB", mock.Anything).Return(outputData, nil).Once()
		repo.On("GetUser", 1).Return(outputData, nil)
		srv := New(repo)

		res, err := srv.InsertProduct(insertData)
		assert.NoError(t, err)
		assert.Equal(t, outputData, res)
		repo.AssertExpectations(t)
	})
}
