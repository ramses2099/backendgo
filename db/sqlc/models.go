package db

import "time"

type Account struct {
	ID        int64     `json:"id"`
	Owner     string    `json:"owner"`
	Balance   int64     `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"create_at"`
}

type Entry struct {
	ID        int64     `json:"id"`
	AccountID int64     `json:"account_id"`
	Amount    int64     `json:"Amount"`
	CreatedAt time.Time `json:"create_at"`
}

type Transfer struct {
	ID          int64     `json:"id"`
	FromAccount int64     `json:"from_account_id"`
	ToAccountID int64     `json:"to_account_id"`
	Amount      int64     `json:"amount"`
	CreatedAt   time.Time `json:"created_at"`
}
