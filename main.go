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

	r.GET("/atm", WebATM)

	r.POST("/api/balance", APIBalance)
	r.POST("/api/deposit", APIDeposit)
	r.POST("/api/withdraw", APIWithdraw)

	r.Run() // listen and serve on 0.0.0.0:8080
}

func WebATM(c *gin.Context) {

}

func APIBalance(c *gin.Context) {
	msg := "Your balance: $" + strconv.Itoa(balance)
	c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(msg))
}

func APIDeposit(c *gin.Context) {
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

func APIWithdraw(c *gin.Context) {
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
