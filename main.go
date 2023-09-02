package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var balance int = 100

const RETURNTYPE = "text/plain; charset=utf-8"

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", entry)
	r.GET("/login", login)
	r.GET("/atm", atm)

	r.POST("/api/balance", getBalance)
	r.POST("/api/deposit", deposit)
	r.POST("/api/withdraw", withdraw)

	r.Run() // listen and serve on 0.0.0.0:8080
}

func entry(c *gin.Context) {

}

func login(c *gin.Context) {

}

func atm(c *gin.Context) {

}

func getBalance(c *gin.Context) {
	msg := "Your balance: $" + strconv.Itoa(balance)
	c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(msg))
}

func deposit(c *gin.Context) {
	input := c.Param("input")
	amount, err := strconv.Atoi(input)

	if err != nil {
		c.Data(http.StatusBadRequest, RETURNTYPE, []byte("Unsupported input"))
		return
	}

	if amount <= 0 {
		c.Data(http.StatusBadRequest, RETURNTYPE, []byte("Amount must be greater than 0"))
		return
	}

	balance += amount
	c.Data(http.StatusAccepted, RETURNTYPE, []byte("Success, current balance: $"+strconv.Itoa(balance)))
}

func withdraw(c *gin.Context) {
	input := c.Param("input")
	amount, err := strconv.Atoi(input)

	if err != nil {
		c.Data(http.StatusBadRequest, RETURNTYPE, []byte("Unsupported input"))
		return
	}

	if amount <= 0 {
		c.Data(http.StatusBadRequest, RETURNTYPE, []byte("Amount must be greater than 0"))
		return
	}

	if balance-amount < 0 {
		c.Data(http.StatusBadRequest, RETURNTYPE, []byte("Insufficient balance: $"+strconv.Itoa(balance)))
		return
	}

	balance -= amount
	c.Data(http.StatusAccepted, RETURNTYPE, []byte("Success, current balance: $"+strconv.Itoa(balance)))
}
