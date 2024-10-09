package routes

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kvncrtr/vendex/models"
)

func InsertNewPart(c *gin.Context) {
	var part models.Part
	var input models.TempPart

	currentTime := time.Now().Format(time.RFC3339)
	parsedTime, err := time.Parse(time.RFC3339, currentTime)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error occured while parsing time string, Try again later."})
		return
	}

	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Not allowed to bind JSON, Try again later."})
		return
	}

	part_number, _ := strconv.ParseInt(input.Part_Number, 10, 64)
	upc, _ := strconv.ParseInt(input.UPC, 10, 64)
	price, _ := strconv.ParseFloat(input.Price, 64)
	weight, _ := strconv.ParseFloat(input.Weight, 64)
	on_hand, _ := strconv.ParseInt(input.On_Hand, 10, 64)
	reorder_amount, _ := strconv.ParseInt(input.Reorder_Amount, 10, 64)
	package_quantity, _ := strconv.ParseInt(input.Package_Quantity, 10, 64)
	reinventory_quantity, _ := strconv.ParseInt(input.Reinventory_Quantity, 10, 64)

	part = models.Part{
		Created_At:           parsedTime,
		Updated_At:           parsedTime,
		Audited_At:           parsedTime,
		Part_Number:          part_number,
		UPC:                  upc,
		Brand:                input.Brand,
		Name:                 input.Name,
		Category:             input.Category,
		Description:          input.Description,
		Price:                price,
		Weight:               weight,
		On_Hand:              on_hand,
		Reorder_Amount:       reorder_amount,
		Package_Quantity:     package_quantity,
		Reinventory_Quantity: reinventory_quantity,
	}

	err = (&part).SaveNewPart()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error occured when saving part input to database."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Part insertion was successful!"})
}

func FetchAllParts(c *gin.Context) {
	parts, err := models.GetAllParts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch all parts."})
		return
	}

	c.JSON(http.StatusOK, parts)
}

func GetPartById(c *gin.Context) {
	var part models.Part
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong when trying to parse params."})
		return
	}

	part, err = part.FetchPartById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error when fetching part by id, try later."})
		return
	}

	c.JSON(http.StatusOK, part)
}

func UpdatePart(c *gin.Context) {
	var part models.Part
	var input models.TempUpdatePart
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong when trying to parse params."})
		return
	}

	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error while trying to bind JSON. Try again!"})
		return
	}

	parsedTime, err := time.Parse(time.DateTime, input.Updated_At)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error during time parser. Rework format!"})
		return
	}

	created, _ := time.Parse(time.RFC3339, input.Created_At)
	audited, _ := time.Parse(time.RFC3339, input.Audited_At)
	on_hand, _ := strconv.ParseInt(input.On_Hand, 10, 64)
	part_number, _ := strconv.ParseInt(input.Part_Number, 10, 64)
	upc, _ := strconv.ParseInt(input.UPC, 10, 64)
	price, _ := strconv.ParseFloat(input.Price, 64)
	weight, _ := strconv.ParseFloat(input.Weight, 64)
	reorder_amount, _ := strconv.ParseInt(input.Reorder_Amount, 10, 64)
	package_quantity, _ := strconv.ParseInt(input.Package_Quantity, 10, 64)
	reinventory_quantity, _ := strconv.ParseInt(input.Reinventory_Quantity, 10, 64)

	part = models.Part{
		ID:                   id,
		Created_At:           created,
		Updated_At:           parsedTime,
		Audited_At:           audited,
		Part_Number:          part_number,
		UPC:                  upc,
		Brand:                input.Brand,
		Name:                 input.Name,
		Category:             input.Category,
		Description:          input.Description,
		Price:                price,
		Weight:               weight,
		On_Hand:              on_hand,
		Reorder_Amount:       reorder_amount,
		Package_Quantity:     package_quantity,
		Reinventory_Quantity: reinventory_quantity,
	}

	err = (&part).ModifyPart()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Fault from inputing updated detail to database. Sorry!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully updated part's details."})
}

func DeletePart(c *gin.Context) {
	var part models.Part
	partId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong when trying to parse params."})
		return
	}

	part, err = (&part).FetchPartById(partId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Fetching the part's records return an error."})
		return
	}

	err = part.RemovePart()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error occured while removing record."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Part deleted successfully!"})
}
