package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var secretKey []byte

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Could not load env file")
	}

	secretKeyEnv := os.Getenv("SECRET_KEY")
	if secretKeyEnv == "" {
		fmt.Println("SECRET_KEY not set in .env")
	} else {
		secretKey = []byte(secretKeyEnv)
	}
}

func GenerateToken(employee_id int64, class string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"employee_id": employee_id,
		"class":       class,
		"expired":     time.Now().Add(time.Hour * 2).Unix(),
	})

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("error signing token: %v", err)
	}

	return signedToken, nil
}

func VerifyToken(token string) (int64, string, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return 0, "", errors.New("trouble parsing token")
	}

	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return 0, "", errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, "", errors.New("invalid token claims")
	}

	employeeID := int64(claims["employee_id"].(float64))
	class := claims["class"].(string)

	return employeeID, class, nil
}
