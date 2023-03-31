package helpers

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/husnulnawafil/online-learning-platform/global/constants"
)

func GenerateToken(uuid, role string) (string, error) {
	tokenExpired := time.Now().Add(constants.TokenExpiration)
	authClaims := jwt.MapClaims{}
	authClaims["uuid"] = uuid
	authClaims["exp"] = tokenExpired
	authClaims["role"] = role
	authClaims["iat"] = time.Now().Unix()

	return jwt.NewWithClaims(jwt.SigningMethodHS256, authClaims).SignedString([]byte(os.Getenv(os.Getenv("TOKEN_SECRET"))))
}

func ValidateToken(token string) (*jwt.MapClaims, int, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("oops error in validating token, signing method invalid: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})

	fmt.Println(os.Getenv("TOKEN_SECRET"))

	if err != nil {
		if strings.EqualFold(err.Error(), jwt.ErrTokenExpired.Error()) {
			return nil, http.StatusUnauthorized, errors.New("oops token is expired, relog or refresh for new token")
		}
		return nil, http.StatusUnauthorized, fmt.Errorf("oops token is unauthorized with err : %v", err)
	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		return &claims, http.StatusOK, err
	} else {
		return nil, http.StatusUnauthorized, errors.New("token is invalid")
	}
}

func GetToken(authorization string) string {
	splitToken := strings.Split(authorization, "Bearer ")
	if len(splitToken) < 2 {
		return ""
	}
	token := splitToken[1]
	return token
}
