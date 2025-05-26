package models

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/kvncrtr/vendex/db"
	"github.com/kvncrtr/vendex/utils"
)

// Models
type Employee struct {
	ID               int64        `json:"id"`
	Class            string       `json:"class"`
	First_Name       string       `json:"first_name" binding:"required"`
	Middle_Name      string       `json:"middle_name" binding:"required"`
	Last_Name        string       `json:"last_name" binding:"required"`
	Sex              string       `json:"sex" binding:"required"`
	Date_Hired       time.Time    `json:"date_hired"`
	Status           string       `json:"status"`
	Termination_Date sql.NullTime `json:"termination_date"`
	Employee_ID      int64        `json:"employee_id"`
	Phone_Number     int64        `json:"phone_number" binding:"required"`
	Email            string       `json:"email" binding:"required"`
	Password         string       `json:"password" binding:"required"`
	Address          string       `json:"address" binding:"required"`
}

type EmployeeAuth struct {
	ID          int64
	Employee_ID int64  `json:"employee_id" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

type TempEmployeeAuth struct {
	Employee_ID string `json:"employee_id" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

// Methods
func (employee *Employee) CreateEmployee() error {
	var employeeID int64
	var err error

	query := `
	INSERT INTO employees(class, first_name, middle_name, last_name, sex, date_hired, phone_number, email, password, address)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	RETURNING id
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	currentDate := utils.CurrentDate()

	hashedPassword, err := utils.HashPassword(employee.Password)
	if err != nil {
		return err
	}

	err = stmt.QueryRow(
		employee.Class,
		employee.First_Name,
		employee.Middle_Name,
		employee.Last_Name,
		employee.Sex,
		currentDate,
		employee.Phone_Number,
		employee.Email,
		hashedPassword,
		employee.Address,
	).Scan(&employeeID)

	if err != nil {
		return err
	}

	employee.ID = employeeID
	return nil
}

func ReturnAllEmployees() ([]Employee, error) {
	var employees []Employee
	query := "SELECT * FROM employees"

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var employee Employee
		err := rows.Scan(
			&employee.ID,
			&employee.Class,
			&employee.First_Name,
			&employee.Middle_Name,
			&employee.Last_Name,
			&employee.Sex,
			&employee.Date_Hired,
			&employee.Status,
			&employee.Termination_Date,
			&employee.Employee_ID,
			&employee.Phone_Number,
			&employee.Email,
			&employee.Password,
			&employee.Address,
		)
		if err != nil {
			return nil, err
		}

		employees = append(employees, employee)
	}

	return employees, nil
}

func GetEmployeeByID(id int64) (*Employee, error) {
	var employee Employee

	query := "SELECT * FROM employees WHERE id = $1"

	row := db.DB.QueryRow(query, id)

	err := row.Scan(
		&employee.ID,
		&employee.Class,
		&employee.First_Name,
		&employee.Middle_Name,
		&employee.Last_Name,
		&employee.Sex,
		&employee.Date_Hired,
		&employee.Status,
		&employee.Termination_Date,
		&employee.Employee_ID,
		&employee.Phone_Number,
		&employee.Email,
		&employee.Password,
		&employee.Address,
	)
	if err != nil {
		return nil, err
	}

	return &employee, nil
}

func (employee *EmployeeAuth) ValidateCredentials() (string, error) {
	var retrievedPassword, class string

	query := `
	SELECT password, class 
	FROM employees 
	WHERE employee_id = $1
	`

	row := db.DB.QueryRow(query, employee.Employee_ID)
	err := row.Scan(&retrievedPassword, &class)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("invalid credentials")
		}
		return "", errors.New("error fetching employee info")
	}

	log.Printf("🔐 Input password: %s", employee.Password)
	log.Printf("🔐 Hashed password from DB: %s", retrievedPassword)

	passwordIsValid := utils.CheckPasswordHash(employee.Password, retrievedPassword)

	if !passwordIsValid {
		return "", errors.New("credentials invalid")
	}

	return class, nil
}

func (employee *Employee) UpdateEmployee() error {
	query := `
	UPDATE employees
	SET id = $1, class = $2, first_name = $3, middle_name = $4, last_name = $5, sex = $6, date_hired = $7, status = $8, termination_date = $9, phone_number = $10, email = $11, password = $12, address = $13    
	WHERE id = $1
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	stmt.Exec(
		employee.ID,
		employee.Class,
		employee.First_Name,
		employee.Middle_Name,
		employee.Last_Name,
		employee.Sex,
		employee.Date_Hired,
		employee.Status,
		employee.Termination_Date,
		employee.Phone_Number,
		employee.Email,
		employee.Password,
		employee.Address,
	)

	return nil
}

func RemoveEmployeeProfile(id int64) error {
	query := `
	DELETE FROM employees
	WHERE id = $1
	`
	result, err := db.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows were deleted")
	}

	return nil
}
