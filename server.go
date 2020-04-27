package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//----------
// Data definitions
//----------

type (
	transaction struct {
		ID            int       `json:"id"`
		Type          string    `json:"type" validate:"required,oneof='credit' 'debit'"`
		Amount        float64   `json:"amount" validate:"required,gte=0"`
		EffectiveDate time.Time `json:"effectiveDate"`
	}

	customValidator struct {
		validator *validator.Validate
	}
)
type allTransactions []transaction

var (
	transactions = map[int]*transaction{}
	seq          = 1
)

//----------
// Helpers
//----------

func (cv *customValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func transactionStructLevelValidation(sl validator.StructLevel) {

	record := sl.Current().Interface().(transaction)
	balance := getCurrentBalance()

	if record.Type == "debit" && balance < record.Amount {
		sl.ReportError(record.Amount, "amount", "Amount", "lte", strconv.FormatFloat(balance, 'f', 2, 64))
	}
}

func getCurrentBalance() float64 {
	var balance float64 = 0
	for _, record := range transactions {
		if record.Type == "debit" {
			balance -= record.Amount
		} else {
			balance += record.Amount
		}
	}

	return balance
}

//----------
// Handlers
//----------

func homePage(c echo.Context) error {
	return c.File("public/index.html")
}

func getTransactions(c echo.Context) error {
	transactionsArray := make(allTransactions, len(transactions))
	totalTransactions := len(transactions)

	for index, record := range transactions {
		transactionsArray[totalTransactions-index] = *record
	}

	return c.JSON(http.StatusOK, transactionsArray)
}

func createTransaction(c echo.Context) error {
	record := &transaction{
		ID:            seq,
		EffectiveDate: time.Now(),
	}

	if err := c.Bind(record); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(record); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	transactions[record.ID] = record
	seq++
	return c.JSON(http.StatusCreated, record)
}

func getTransaction(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	record := transactions[id]

	if record == nil {
		return c.JSON(http.StatusNotFound, record)
	}
	return c.JSON(http.StatusOK, record)
}

func getBalance(c echo.Context) error {
	return c.JSON(http.StatusOK, getCurrentBalance())
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Client Routes
	e.GET("/", homePage)
	e.Static("/static", "public/static")

	// API Routes
	api := e.Group("/api")
	api.GET("/transactions", getTransactions)
	api.POST("/transactions", createTransaction)
	api.GET("/transactions/:id", getTransaction)
	api.GET("/default", getBalance)

	// Transaction validation
	var validate = validator.New()
	validate.RegisterStructValidation(transactionStructLevelValidation, transaction{})
	e.Validator = &customValidator{validator: validate}

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
