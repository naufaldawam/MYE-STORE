package usecase

import (
	"errors"
	"project3/group3/domain"
	"project3/group3/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllData(t *testing.T) {
	limit := 1
	offset := 10
	idFromToken := 1
	outdata := []domain.Cart{
		{
			ID:         1,
			Stock:      20,
			TotalPrice: 100000,
			UserID:     5,
			Product: domain.ProductCart{
				ID:          1,
				ProductName: "adidas",
				Price:       100,
				Stock:       4,
			},
		},
	}

	repo := new(mocks.CartData)

	t.Run("data found", func(t *testing.T) {
		repo.On("SelectData", limit, offset, idFromToken).Return(outdata, nil)

		srv := New(repo)
		res, err := srv.GetAllData(limit, offset, idFromToken)
		assert.NoError(t, err)
		assert.Equal(t, outdata, res)
	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("SelectData", 0, 0, 0).Return([]domain.Cart{}, errors.New("error"))

		srv := New(repo)
		notFound, err := srv.GetAllData(0, 0, 0)

		assert.NotNil(t, err)
		assert.Equal(t, []domain.Cart{}, notFound)
	})
}

func TestCreateData(t *testing.T) {
	repo := new(mocks.CartData)
	insertData := domain.Cart{
		ID:         1,
		Stock:      7,
		Status:     "Pending",
		TotalPrice: 10000,
		Product: domain.ProductCart{
			ID: 1,
		},
	}
	outputData := 1

	t.Run("success insert", func(t *testing.T) {
		repo.On("CheckCart", 1, 0).Return(false, 1, 1, nil)
		repo.On("InsertData", insertData).Return(outputData, nil)
		srv := New(repo)
		res, err := srv.CreateData(insertData)
		assert.NoError(t, err)
		assert.Equal(t, outputData, res)
		repo.AssertExpectations(t)
	})
}

func TestDeleteData(t *testing.T) {
	idCart := 1
	idToken := 1
	repo := new(mocks.CartData)

	t.Run("success delete", func(t *testing.T) {
		repo.On("DeleteDataDB", idCart, idToken).Return(1, nil)

		srv := New(repo)
		delete, err := srv.DeleteData(1, 1)

		assert.NoError(t, err)
		assert.Equal(t, 1, delete)
	})

}

func TestUpdateData(t *testing.T) {

	repo := new(mocks.CartData)

	t.Run("success update", func(t *testing.T) {
		repo.On("UpdateDataDB", 10, 5, 3).Return(1, nil).Once()

		srv := New(repo)
		update, err := srv.UpdateData(10, 5, 3)
		assert.NoError(t, err)
		assert.Equal(t, 1, update)
	})
}
