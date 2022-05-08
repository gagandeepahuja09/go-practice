package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresTransactionLogger struct {
	events chan<- Event // Write only channel for sending events
	errors <-chan error // Read only channel for receiving errors
	db     *sql.DB      // Database access interface
}

type PostgresDBParams struct {
	dbName   string
	host     string
	user     string
	password string
}

func (l *PostgresTransactionLogger) WritePut(key, value string) {
	l.events <- Event{EventType: EventPut, Key: key, Value: value}
}

func (l *PostgresTransactionLogger) WriteDelete(key string) {
	l.events <- Event{EventType: EventDelete, Key: key}
}

func (l *PostgresTransactionLogger) Err() <-chan error {
	return l.errors
}

func (l *PostgresTransactionLogger) createTable() error {
	query := `CREATE TABLE IF NOT EXISTS transactions (
		sequence SERIAL PRIMARY KEY,
		event_type SMALLINT NOT NULL,
		key VARCHAR(255) NOT NULL,
		value VARCHAR(255) DEFAULT NULL
	)`
	_, err := l.db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}
	return nil
}

func NewPostgresTransactionLogger(config PostgresDBParams) (TransactionLogger, error) {
	connStr := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		config.user, config.password, config.host, "5432", config.dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	// Many drivers, including lib/pq don't create a connection to db immediately
	// db.Ping forces the driver to establish and test a connection.
	err = db.Ping() // Test the database connection
	if err != nil {
		return nil, fmt.Errorf("failed to open db connection: %w", err)
	}

	logger := &PostgresTransactionLogger{db: db}
	if err = logger.createTable(); err != nil {
		return nil, fmt.Errorf("failed to create table: %w", err)
	}

	return logger, nil
}

func (l *PostgresTransactionLogger) Run() {
	events := make(chan Event, 16)
	l.events = events

	errors := make(chan error, 1)
	l.errors = errors

	go func() {
		query := `INSERT INTO transactions 
		(event_type, key, value)
		VALUES ($1, $2, $3)`

		for e := range events {
			_, err := l.db.Exec(query, e.EventType, e.Key, e.Value)

			if err != nil {
				errors <- err
			}
		}
	}()
}

func (l *PostgresTransactionLogger) ReadEvents() (<-chan Event, <-chan error) {
	outEvent := make(chan Event)
	outError := make(chan error, 1)

	go func() {
		defer close(outEvent) // close the channels
		defer close(outError) // when the goroutine ends

		query := `SELECT sequence, event_type, key, value FROM transactions
					ORDER BY sequence`

		rows, err := l.db.Query(query)
		if err != nil {
			outError <- fmt.Errorf("sql query error: %w", err)
			return
		}

		defer rows.Close() // this is important!
		e := Event{}

		for rows.Next() {
			err = rows.Scan(&e.Sequence, &e.EventType, &e.Key, &e.Value)

			if err != nil {
				outError <- fmt.Errorf("error reading row: %w", err)
				return
			}

			outEvent <- e
		}

		err = rows.Err()
		if err != nil {
			outError <- fmt.Errorf("transaction log read failure: %w", err)
		}
	}()

	return outEvent, outError
}
