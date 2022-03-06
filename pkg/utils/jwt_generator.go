package utils

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"myponyasia.com/hub-api/app/models"
)

type JWT struct {
	privateKey []byte
	publicKey  []byte
}

func NewJWT(privateKey []byte, publicKey []byte) JWT {
	return JWT{
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}

// GenerateNewAccessToken func for generate a new Access token.
func GenerateNewAccessToken(user models.User) (string, string, error) {

	prvKey, err := ioutil.ReadFile(os.Getenv("JWT_FILE_SECRET"))
	if err != nil {
		log.Fatalln(err)
	}
	pubKey, err := ioutil.ReadFile(os.Getenv("JWT_FILE_PUBLIC"))
	if err != nil {
		log.Fatalln(err)
	}

	jwtToken := NewJWT(prvKey, pubKey)

	key, err := jwt.ParseRSAPrivateKeyFromPEM(jwtToken.privateKey)
	if err != nil {
		return "", "", err
	}

	// Set secret key from .env file.
	// secret := os.Getenv("JWT_SECRET_KEY")

	// Set expires minutes count for secret key from .env file.
	atExpired, _ := strconv.Atoi(os.Getenv("JWT_ACCESS_TOKEN_EXPIRED"))
	rtExpired, _ := strconv.Atoi(os.Getenv("JWT_REFRESH_TOKEN_EXPIRED"))

	// Create a new access token claims.
	atClaims := jwt.MapClaims{}
	atClaims["typ"] = "session"
	atClaims["iss"] = "myponyasia.com"
	atClaims["aud"] = "myponyasia.com"
	atClaims["iat"] = time.Now().Unix()
	atClaims["nbf"] = time.Now().Unix()
	atClaims["exp"] = time.Now().Add(time.Minute * time.Duration(atExpired)).Unix()
	atClaims["uuid"] = user.ID.String()
	atClaims["data"] = fiber.Map{
		"roles": []string{
			user.Role,
		},
		"permissions": []string{},
	}

	// Create a new JWT access token with claims.
	atToken := jwt.NewWithClaims(jwt.SigningMethodRS256, atClaims)
	at, err := atToken.SignedString(key)
	if err != nil {
		return "", "", err
	}

	// Create a new refresh token claims.
	rtClaims := jwt.MapClaims{}
	rtClaims["typ"] = "refresh"
	rtClaims["iss"] = "myponyasia.com"
	rtClaims["aud"] = "myponyasia.com"
	rtClaims["iat"] = time.Now().Unix()
	rtClaims["nbf"] = time.Now().Unix()
	rtClaims["exp"] = time.Now().Add((time.Hour * 24) * time.Duration(rtExpired)).Unix()
	rtClaims["uuid"] = user.ID.String()

	// Create a new JWT refresh token with claims.
	rtToken := jwt.NewWithClaims(jwt.SigningMethodRS256, rtClaims)
	rt, err := rtToken.SignedString(key)
	if err != nil {
		return "", "", err
	}

	return at, rt, nil
}
