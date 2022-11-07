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
	ID       uint   `gorm:"primarykey"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateJWT(user models.User) (tokenString string, err error) {

	expirationTime := time.Now().Add(24 * 3 * time.Hour) // token expire after 3 days

	// destruct user to claims
	claims := &JWTClaim{
		ID:       user.ID,
		Username: user.Username,
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
			// this function receives receives the parsed, but unverified Token .
			// useful for multi key, you can header of the token (such as `kid`) to identify which key to use.
			return []byte(jwtKey), nil // we are only uses one key for this situation
		},
	)
	if err != nil {
		return
	}

	// parse the claims
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}

	// check token expiration
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}

	return
}
