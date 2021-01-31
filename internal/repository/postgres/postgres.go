package postgres

import (
	"context"
	"fmt"
	"zamio/internal/domain"

	pgx "github.com/jackc/pgx/v4"
)

// DB struct with pgsql connection field
type DB struct {
	conn *pgx.Conn
}

// New returns DB object with connection to pgsql db
func New(host, port, user, password, dbname string, sslmode bool) (*DB, error) {
	sslmodeStr := "disable"
	if sslmode {
		sslmodeStr = "enable"
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host,
		port,
		user,
		password,
		dbname,
		sslmodeStr,
	)

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	pgDB := &DB{
		conn: conn,
	}

	return pgDB, nil
}

// CreateAccount add account to pg
func (db *DB) CreateAccount(ctx context.Context, account domain.Account) error {
	_, err := db.conn.Exec(ctx, "INSERT INTO accounts (email) VALUES ($1)", account.Email)
	if err != nil {
		return fmt.Errorf("error inserting a account into the database: %+v", err)
	}

	return err
}

// AccountByEmail return account from pg, takes user's email
func (db *DB) AccountByEmail(ctx context.Context, email string) (domain.Account, error) {
	var acc domain.Account

	row := db.conn.QueryRow(context.Background(), "SELECT * FROM accounts WHERE email=$1", email)
	err := row.Scan(
		&acc.ID,
		&acc.Email,
		&acc.Balance,
	)
	if err != nil {
		return acc, fmt.Errorf("finding account with email %s in the database error: %+v", email, err)
	}

	return acc, err
}

// UpdateAccountsBalance updates two account's balance via a transaction
func (db *DB) UpdateAccountsBalance(ctx context.Context, accountFirst, accountSecond domain.Account) error {
	tx, err := db.conn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("the start of the transaction error: %v", err)
	}
	defer tx.Rollback(ctx)

	// TODO: вынести в метод
	ct, err := tx.Exec(ctx, "UPDATE accounts SET balance = $1 WHERE email = $2", accountFirst.Balance, accountFirst.Email)
	if err != nil {
		return fmt.Errorf("updating user %s balance error: %v", accountFirst.Email, err)
	}

	if ct.RowsAffected() != 1 {
		return fmt.Errorf("rows for user %s not updated", accountFirst.Email)
	}

	ct, err = tx.Exec(ctx, "UPDATE accounts SET balance = $1 WHERE email = $2", accountSecond.Balance, accountSecond.Email)
	if err != nil {
		return fmt.Errorf("updating user %s balance error: %v", accountSecond.Email, err)
	}

	if ct.RowsAffected() != 1 {
		return fmt.Errorf("rows for user %s not updated", accountSecond.Email)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("the end of the transaction error: %v", err)
	}

	return err
}

// Close close pg connection
func (db *DB) Close(ctx context.Context) {
	db.conn.Close(ctx)
}
