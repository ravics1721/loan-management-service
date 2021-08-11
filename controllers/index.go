package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"loan-management-service.com/lms/models"
)

type Response struct {
	ctx *gin.Context
}

func (c *Response) JSON(status int, message string, data interface{}) {
	c.ctx.JSON(status, gin.H{
		"message": "" + message,
		"data":    data,
	})
}

var response = new(Response)

var val = gin.H{
	"DELETE": "/loans : Cancel Loan",
	"PATCH":  "/loans : Update Loan Status",
	"POST":   "/loans : Create new Loan",
	"GET":    "/loans : To get loans also use query to filter result e.g: /loans?status=new,approved ,   /loans/123132",
}

func DefaultRequest(c *gin.Context) {
	c.JSON(200, gin.H{
		"usages":  val,
		"message": "Welcome 🌈 to Loan Management Service",
	})
}

func AddNewLoan(c *gin.Context) {
	SetHeaders(c)
	var request models.LoanRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	response.JSON(201, "Success", nil)
}
func GetLoans(c *gin.Context) {
	SetHeaders(c)
	response.JSON(200, "Success", nil)
}
func GetLoanById(c *gin.Context) {
	SetHeaders(c)
	response.JSON(200, "Success", nil)
}
func UpdateLoanStatus(c *gin.Context) {
	SetHeaders(c)
	response.JSON(200, "Success", nil)
}
func CancelLoan(c *gin.Context) {
	SetHeaders(c)
	response.JSON(200, "Success", nil)
}
