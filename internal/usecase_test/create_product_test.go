package usecase_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vitorconti/go-products/internal/dto"
	"github.com/vitorconti/go-products/internal/entity"
	"github.com/vitorconti/go-products/internal/usecase"
	"github.com/vitorconti/go-products/pkg/events"
)

type mockProductRepository struct {
	mock.Mock
}

func (m *mockProductRepository) Save(product *entity.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

type mockEventDispatcher struct {
	mock.Mock
}

func (m *mockEventDispatcher) Dispatch(event events.EventInterface) error {
	args := m.Called(event)
	return args.Error(0)
}

func TestCreateProductUseCase_Execute(t *testing.T) {
	t.Run("Should create product successfully", func(t *testing.T) {
		// Arrange
		product := dto.ProductInputDTO{
			Name:        "Test Product",
			Description: "This is a test product",
			Price:       9.99,
		}
		expectedOutput := dto.ProductOutputDTO{
			Name:        "Test Product",
			Description: "This is a test product",
			Price:       9.99,
		}

		mockProductRepo := new(mockProductRepository)
		mockProductRepo.On("Save", mock.Anything).Return(nil)

		mockEventDispatcher := new(mockEventDispatcher)
		mockEventDispatcher.On("Dispatch", mock.Anything).Return(nil)

		createProductUseCase := usecase.NewCreateProductUseCase(mockProductRepo, events.NewProductCreated(), mockEventDispatcher)

		// Act
		output, err := createProductUseCase.Execute(product)

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, expectedOutput, output)
		mockProductRepo.AssertCalled(t, "Save", mock.Anything)
		mockEventDispatcher.AssertCalled(t, "Dispatch", mock.Anything)
	})

	t.Run("Should fail to create product", func(t *testing.T) {
		// Arrange
		product := dto.ProductInputDTO{
			Name:        "Test Product",
			Description: "This is a test product",
			Price:       9.99,
		}
		expectedErr := errors.New("error saving product")

		mockProductRepo := new(mockProductRepository)
		mockProductRepo.On("Save", mock.Anything).Return(expectedErr)

		mockEventDispatcher := new(mockEventDispatcher)

		createProductUseCase := usecase.NewCreateProductUseCase(mockProductRepo, events.NewProductCreated(), mockEventDispatcher)

		// Act
		output, err := createProductUseCase.Execute(product)

		// Assert
		assert.Equal(t, expectedErr, err)
		assert.Equal(t, dto.ProductOutputDTO{}, output)
		mockProductRepo.AssertCalled(t, "Save", mock.Anything)
		mockEventDispatcher.AssertNotCalled(t, "Dispatch")
	})
}
