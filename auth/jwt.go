package auth

import (
	"errors"
	"os"
	"time"

	"github.com/ahmedsat/sat-store/models"
	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// structure for claims in JWT
type JWTClaim struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	Privileges string `json:"privileges"`

	jwt.StandardClaims
}

func GenerateJWT(user models.User) (tokenString string, err error) {

	expirationTime := time.Now().Add(24 * 3 * time.Hour) // token expire after 3 days

	// destruct user to claims
	claims := &JWTClaim{
		Email:      user.Email,
		Username:   user.Username,
		Privileges: user.Privileges,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // generate token
	tokenString, err = token.SignedString(jwtKey)              // sign token

	return tokenString, err
}

func ValidateToken(signedToken string) (claims *JWTClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}
