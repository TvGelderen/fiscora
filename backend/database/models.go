// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Budget struct {
	ID          string
	UserID      uuid.UUID
	Name        string
	Description string
	Amount      string
	StartDate   time.Time
	EndDate     time.Time
	Created     time.Time
	Updated     time.Time
}

type BudgetExpense struct {
	ID              int32
	BudgetID        string
	Name            string
	AllocatedAmount string
	CurrentAmount   string
	Created         time.Time
	Updated         time.Time
}

type RecurringTransaction struct {
	ID           int32
	UserID       uuid.UUID
	StartDate    time.Time
	EndDate      time.Time
	Interval     string
	DaysInterval sql.NullInt32
	Created      time.Time
	Updated      time.Time
}

type Transaction struct {
	ID                     int32
	UserID                 uuid.UUID
	RecurringTransactionID sql.NullInt32
	Description            string
	Amount                 string
	Type                   string
	Date                   time.Time
	Created                time.Time
	Updated                time.Time
}

type User struct {
	ID         uuid.UUID
	Provider   string
	ProviderID string
	Username   string
	Email      string
	Avatar     sql.NullString
	Created    time.Time
	Updated    time.Time
}
