
package usecase

import (
    "context"
    "testing"

    "github.com/stretchr/testify/assert"
)

type mockRepo struct{}

func (m mockRepo) Consume(ctx context.Context, userID int64, tenor int, amount float64) error {
    if amount > 100 {
        return assert.AnError
    }
    return nil
}

func TestCreate(t *testing.T) {
    uc := NewTransactionUsecase(mockRepo{})
    err := uc.Create(context.Background(), 1, 1, 50)
    assert.NoError(t, err)
}
