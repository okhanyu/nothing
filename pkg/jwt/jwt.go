package jwt

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var JwtKey []byte

var TokenBox map[string]int64

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func init() {
	if TokenBox == nil {
		TokenBox = make(map[string]int64)
	}
}
func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)

	if err != nil {
		log.Printf("token generate error is:%+v", err)
		return "", err
	}
	TokenBox[tokenString] = expirationTime.Unix()
	return tokenString, nil

}

func Verify(tokenString string) (bool, error) {
	if tokenString == "" {
		log.Printf("token verify error is token empty")
		return false, jwt.NewValidationError("token verify error is token empty", 0)
	}
	if _, ok := TokenBox[tokenString]; !ok {
		log.Printf("token verify error is token empty")
		return false, jwt.NewValidationError("token verify error is token not in TokenBox ", 0)
	}
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			log.Printf("无效的签名")
			delete(TokenBox, tokenString)
			//http.Error(w, "无效的签名", http.StatusUnauthorized)
			return false, err
		}
		//http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("%+v", err)
		delete(TokenBox, tokenString)
		return false, err
	}

	if !token.Valid {
		log.Printf("无效的令牌")
		delete(TokenBox, tokenString)
		//http.Error(w, "无效的令牌", http.StatusUnauthorized)
		return false, err
	}

	return true, nil

}
