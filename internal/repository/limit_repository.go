
package repository

import (
    "context"
    "database/sql"
)

type LimitRepository struct {
    db *sql.DB
}

func NewLimitRepository(db *sql.DB) *LimitRepository {
    return &LimitRepository{db: db}
}

func (r *LimitRepository) Consume(ctx context.Context, userID int64, tenor int, amount float64) error {
    tx, err := r.db.BeginTx(ctx, nil)
    if err != nil {
        return err
    }

    var limit, used float64
    err = tx.QueryRowContext(ctx, `
        SELECT limit_amount, used_amount
        FROM limits
        WHERE user_id=? AND tenor=?
        FOR UPDATE
    `, userID, tenor).Scan(&limit, &used)
    if err != nil {
        tx.Rollback()
        return err
    }

    if used+amount > limit {
        tx.Rollback()
        return sql.ErrNoRows
    }

    _, err = tx.ExecContext(ctx,
        "UPDATE limits SET used_amount=? WHERE user_id=? AND tenor=?",
        used+amount, userID, tenor)
    if err != nil {
        tx.Rollback()
        return err
    }

    return tx.Commit()
}
