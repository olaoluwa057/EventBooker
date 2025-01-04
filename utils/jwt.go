package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "secret"

func GenerateToken (email string, user_id int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"user_id": user_id,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func GenerateAdminToken (email string, user_id int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"user_id": user_id,
		"admin": true,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(tokenString string) (int64, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("could not verify token")
	}
	
	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return 0, errors.New("token is invalid")
	}


	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0,  errors.New("could not parse claims")
	}

	// email := claims["email"].(string)
	user_id := int64(claims["user_id"].(float64))
	


	return user_id, nil
	
}

// func VerifyAdminToken(tokenString string) (bool, error) {
// 	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		_, ok := token.Method.(*jwt.SigningMethodHMAC)

// 		if !ok {
// 			return nil, jwt.ErrSignatureInvalid
// 		}

// 		return []byte(secretKey), nil
// 	})

// 	if err != nil {
// 		return false, errors.New("could not verify token")
// 	}
	
// 	tokenIsValid := parsedToken.Valid

// 	if !tokenIsValid {
// 		return false, errors.New("token is invalid")
// 	}

// 	claims, ok := parsedToken.Claims.(jwt.MapClaims)

// 	if !ok {
// 		return false, errors.New("could not parse claims")
// 	}

// 	admin := claims["admin"].(bool)

// 	return admin, nil

// }