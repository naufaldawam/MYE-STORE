package usecase

import (
	"errors"
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

func TestDeleteProduct(t *testing.T) {
	repo := new(mocks.ProductData)

	t.Run("success delete", func(t *testing.T) {
		repo.On("DeleteProductDB", mock.Anything).Return(1, nil)

		srv := New(repo)
		delete, err := srv.DeleteProduct(1)

		assert.NoError(t, err)
		assert.Equal(t, 1, delete)
	})
}
func TestGetAllData(t *testing.T) {
	limit := 1
	offset := 10
	outdata := []domain.Product{
		{
			ID:           1,
			ProductName:  "adidas",
			ProductImage: "ini gambar",
			Price:        10000,
			Stock:        10,
			User: domain.UserAdmin{
				ID:   2,
				Name: "admin",
				Role: "admin",
			},
		},
	}

	repo := new(mocks.ProductData)

	t.Run("data found", func(t *testing.T) {
		repo.On("SelectData", limit, offset).Return(outdata, nil)

		srv := New(repo)
		res, err := srv.GetAllData(limit, offset)
		assert.NoError(t, err)
		assert.Equal(t, outdata, res)
	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("SelectData", 0, 0).Return([]domain.Product{}, errors.New("error"))

		srv := New(repo)
		notFound, err := srv.GetAllData(0, 0)

		assert.NotNil(t, err)
		assert.Equal(t, []domain.Product{}, notFound)
	})
}

func TestGetProductById(t *testing.T) {
	outdata := domain.Product{
		ID:           1,
		ProductName:  "adidas",
		ProductImage: "ini gambar",
		Price:        10000,
		Stock:        10,
		User: domain.UserAdmin{
			ID:   2,
			Name: "admin",
			Role: "admin",
		},
	}

	repo := new(mocks.ProductData)

	t.Run("data found", func(t *testing.T) {
		repo.On("SelectDataById", 1).Return(outdata, nil)

		srv := New(repo)
		res, err := srv.GetProductById(1)
		assert.Nil(t, err)
		assert.Equal(t, outdata, res)
	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("SelectDataById", 0).Return(domain.Product{}, errors.New("error"))
		srv := New(repo)
		notFound, err := srv.GetProductById(0)
		assert.NotNil(t, err)
		assert.Equal(t, domain.Product{}, notFound)
	})
}
func TestUpdateData(t *testing.T) {
	insert := domain.Product{
		ProductName:  "adidas",
		ProductImage: "ini gambar",
		Price:        10000,
		Stock:        10,
	}
	repo := new(mocks.ProductData)

	t.Run("succes update", func(t *testing.T) {
		repo.On("UpdateData", mock.Anything, 1, 1).Return(1, nil).Once()

		srv := New(repo)
		update, err := srv.UpdateData(insert, 1, 1)
		assert.NoError(t, err)
		assert.Equal(t, 1, update)

	})
}
