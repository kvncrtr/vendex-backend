package models

import (
	"fmt"
	"time"

	"github.com/kvncrtr/vendex/db"
)

type Part struct {
	ID                   int64     `json:"id"`
	Created_At           time.Time `json:"created_at"`
	Updated_At           time.Time `json:"updated_at"`
	Audited_At           time.Time `json:"audited_at"`
	Part_Number          int64     `json:"part_number" binding:"required"`
	UPC                  int64     `json:"upc" binding:"required"`
	Brand                string    `json:"brand" binding:"required"`
	Name                 string    `json:"name" binding:"required"`
	Category             string    `json:"category" binding:"required"`
	Description          string    `json:"description"`
	Price                float64   `json:"price"`
	Weight               float64   `json:"weight"`
	On_Hand              int64     `json:"on_hand" binding:"required"`
	Reorder_Amount       int64     `json:"reorder_amount" binding:"required"`
	Package_Quantity     int64     `json:"package_quantity" binding:"required"`
	Reinventory_Quantity int64     `json:"reinventory_quantity"`
}

type TempPart struct {
	ID                   int64  `json:"id"`
	Created_At           string `json:"created_at"`
	Updated_At           string `json:"updated_at"`
	Audited_At           string `json:"audited_at"`
	Part_Number          string `json:"part_number" binding:"required"`
	UPC                  string `json:"upc" binding:"required"`
	Brand                string `json:"brand" binding:"required"`
	Name                 string `json:"name" binding:"required"`
	Category             string `json:"category" binding:"required"`
	Description          string `json:"description"`
	Price                string `json:"price"`
	Weight               string `json:"weight"`
	On_Hand              string `json:"on_hand" binding:"required"`
	Reorder_Amount       string `json:"reorder_amount" binding:"required"`
	Package_Quantity     string `json:"package_quantity" binding:"required"`
	Reinventory_Quantity string `json:"reinventory_quantity"`
}

type TempUpdatePart struct {
	ID                   string `json:"id"`
	Created_At           string `json:"created_at"`
	Updated_At           string `json:"updated_at"`
	Audited_At           string `json:"audited_at"`
	Part_Number          string `json:"part_number" binding:"required"`
	UPC                  string `json:"upc" binding:"required"`
	Brand                string `json:"brand" binding:"required"`
	Name                 string `json:"name" binding:"required"`
	Category             string `json:"category" binding:"required"`
	Description          string `json:"description"`
	Price                string `json:"price"`
	Weight               string `json:"weight"`
	On_Hand              string `json:"on_hand" binding:"required"`
	Reorder_Amount       string `json:"reorder_amount" binding:"required"`
	Package_Quantity     string `json:"package_quantity" binding:"required"`
	Reinventory_Quantity string `json:"reinventory_quantity"`
}

func (p *Part) SaveNewPart() error {
	var partID int64
	var err error

	query := `
	INSERT INTO parts(created_at, updated_at, audited_at, part_number, upc, brand, name, category, description, price, weight, on_hand, reorder_amount, package_quantity, reinventory_quantity)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
	RETURNING id
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		p.Created_At,
		p.Updated_At,
		p.Audited_At,
		p.Part_Number,
		p.UPC,
		p.Brand,
		p.Name,
		p.Category,
		p.Description,
		p.Price,
		p.Weight,
		p.On_Hand,
		p.Reorder_Amount,
		p.Package_Quantity,
		p.Reinventory_Quantity,
	).Scan(&partID)

	if err != nil {
		fmt.Println(p.UPC)
		return err
	}

	p.ID = partID
	return nil
}

func GetAllParts() ([]Part, error) {
	var parts []Part

	query := `
	SELECT id, audited_at, part_number, upc, brand, name, category, description, price, weight, on_hand, reorder_amount, package_quantity, reinventory_quantity 
	FROM parts
	`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var part Part
		err := rows.Scan(
			&part.ID,
			&part.Audited_At,
			&part.Part_Number,
			&part.UPC,
			&part.Brand,
			&part.Name,
			&part.Category,
			&part.Description,
			&part.Price,
			&part.Weight,
			&part.On_Hand,
			&part.Reorder_Amount,
			&part.Package_Quantity,
			&part.Reinventory_Quantity,
		)
		if err != nil {
			return nil, err
		}

		parts = append(parts, part)
	}

	return parts, nil
}

func (p *Part) FetchPartById(id int64) (Part, error) {
	query := `
	SELECT id, audited_at, part_number, upc, brand, name, category, description, price, weight, on_hand, reorder_amount, package_quantity, reinventory_quantity 
	FROM parts
	WHERE id = $1
	`

	row := db.DB.QueryRow(query, id)

	err := row.Scan(
		&p.ID,
		&p.Audited_At,
		&p.Part_Number,
		&p.UPC,
		&p.Brand,
		&p.Name,
		&p.Category,
		&p.Description,
		&p.Price,
		&p.Weight,
		&p.On_Hand,
		&p.Reorder_Amount,
		&p.Package_Quantity,
		&p.Reinventory_Quantity,
	)
	if err != nil {
		return Part{}, err
	}

	return *p, nil
}

func (p *Part) ModifyPart() error {
	query := `
	UPDATE parts
	SET created_at = $2, updated_at = $3, audited_at = $4, part_number = $5, upc= $6, brand = $7, name = $8, category = $9, description = $10, price = $11, weight = $12, on_hand = $13, reorder_amount = $14, package_quantity = $15, reinventory_quantity = $16
	WHERE id = $1
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		&p.ID,
		&p.Created_At,
		&p.Updated_At,
		&p.Audited_At,
		&p.Part_Number,
		&p.UPC,
		&p.Brand,
		&p.Name,
		&p.Category,
		&p.Description,
		&p.Price,
		&p.Weight,
		&p.On_Hand,
		&p.Package_Quantity,
		&p.Reinventory_Quantity,
		&p.Reorder_Amount,
	)
	if err != nil {
		return err
	}

	return nil
}

func (p Part) RemovePart() error {
	query := "DELETE FROM parts WHERE id = $1"

	_, err := db.DB.Exec(query, p.ID)
	if err != nil {
		return err
	}

	return nil
}
