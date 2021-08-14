package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ravics1721/loan-management-service/controllers"
)

func Router(ginMode string) *gin.Engine {
	gin.SetMode(ginMode)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// router.Use() //Todo: Add cors middleware
	router.GET("/", controllers.DefaultRequest)
	loanRoutes := router.Group("/loans")
	{
		loanRoutes.POST("/", controllers.AddNewLoan)           // Create new Loan
		loanRoutes.GET("/", controllers.GetLoans)              // Get all loans or filter loans array
		loanRoutes.GET("/:id", controllers.GetLoanById)        //Get loan by id
		loanRoutes.PATCH("/:id", controllers.UpdateLoanStatus) //Update loan status
		loanRoutes.DELETE("/:id", controllers.CancelLoan)      //Cancel loan
	}

	return router
}
