package handlers

import (
	"database/sql"
	"math/rand"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/tvgelderen/fiscora/auth"
	"github.com/tvgelderen/fiscora/repository"
	"github.com/tvgelderen/fiscora/types"
)

type APIHandler struct {
	UserRepository        repository.IUserRepository
	TransactionRepository repository.ITransactionRepository
	BudgetRepository      repository.IBudgetRepository
	AuthService           *auth.AuthService
}

func NewAPIHandler(db *sql.DB, auth *auth.AuthService) *APIHandler {
	return &APIHandler{
		UserRepository:        repository.CreateUserRepository(db),
		TransactionRepository: repository.CreateTransactionRepository(db),
		BudgetRepository:      repository.CreateBudgetRepository(db),
		AuthService:           auth,
	}
}

func getUserId(c echo.Context) uuid.UUID {
	return c.Get(userIdKey).(uuid.UUID)
}

func getMonth(c echo.Context) int {
	monthParam := c.QueryParam("month")
	month, err := strconv.ParseInt(monthParam, 10, 16)
	if err != nil {
		month = int64(time.Now().Month())
	}

	return int(month)
}

func getYear(c echo.Context) int {
	yearParam := c.QueryParam("year")
	year, err := strconv.ParseInt(yearParam, 10, 16)
	if err != nil {
		year = int64(time.Now().Month())
	}

	return int(year)
}

func getMonthRange(month int, year int) types.DateRange {
	start := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 1, -1)
	return types.DateRange{
		Start: start,
		End:   end,
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func generateRandomString(length int) string {
	str := make([]rune, length)
	for idx := range str {
		str[idx] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(str)
}
