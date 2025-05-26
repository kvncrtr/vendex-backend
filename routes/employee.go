package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kvncrtr/vendex/models"
	"github.com/kvncrtr/vendex/utils"
)

func makeEmployeeProfile(c *gin.Context) {
	var employee models.Employee
	var err error

	err = c.ShouldBindJSON(&employee)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request."})
		return
	}

	err = employee.CreateEmployee()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not make employee's profile."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Profile creation is successful!"})
}

func LoginEmployee(c *gin.Context) {
	var tempEmployee models.TempEmployeeAuth

	err := c.ShouldBindJSON(&tempEmployee)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request."})
		return
	}

	employeeID, err := strconv.ParseInt(tempEmployee.Employee_ID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse int in request."})
		return
	}

	employee := models.EmployeeAuth{
		Employee_ID: employeeID,
		Password:    tempEmployee.Password,
	}

	class, err := employee.ValidateCredentials()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(employee.Employee_ID, class)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate employee. Token failed."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Credentials accepted! Welcome :)", "token": token})
}

func fetchAllEmployees(c *gin.Context) {
	var err error
	var employees []models.Employee

	employees, err = models.ReturnAllEmployees()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch all employees. Try again later."})
		return
	}

	c.JSON(http.StatusOK, employees)
}

func getEmployee(c *gin.Context) {
	var err error
	var employeeID int64
	var employee *models.Employee

	employeeID, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse int for employee's profile. Try again."})
		return
	}

	employee, err = models.GetEmployeeByID(employeeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Employee's profile does not exist. Try again."})
		return
	}

	c.JSON(http.StatusOK, employee)
}

func updateEmployee(c *gin.Context) {
	var employee models.Employee

	err := c.ShouldBindJSON(&employee)
	if err != nil {
		return
	}

	err = employee.UpdateEmployee()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't update profile! "})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee update successfully!"})
}

func terminateEmployee(c *gin.Context) {
	employeeID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse int for employee's profile. Try again."})
		return
	}

	err = models.RemoveEmployeeProfile(employeeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Employee's profile doesn't exist."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}
