
package usecase

import "context"

type LimitRepo interface {
    Consume(ctx context.Context, userID int64, tenor int, amount float64) error
}

type TransactionUsecase struct {
    repo LimitRepo
}

func NewTransactionUsecase(r LimitRepo) *TransactionUsecase {
    return &TransactionUsecase{repo: r}
}

func (u *TransactionUsecase) Create(ctx context.Context, userID int64, tenor int, amount float64) error {
    return u.repo.Consume(ctx, userID, tenor, amount)
}
