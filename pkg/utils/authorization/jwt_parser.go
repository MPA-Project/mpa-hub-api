package authorization

// import (
// 	"errors"
// 	"io/ioutil"
// 	"os"
// 	"strings"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/golang-jwt/jwt"
// )

// // TokenMetadata struct to describe metadata in JWT.
// type TokenMetadata struct {
// 	Expires int64
// }

// // ExtractTokenMetadata func to extract metadata from JWT.
// func ExtractTokenMetadata(c *fiber.Ctx) (*TokenMetadata, error) {
// 	token, err := verifyToken(c)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Setting and checking token and credentials.
// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if ok && token.Valid {
// 		// Expires time.
// 		expires := int64(claims["exp"].(float64))

// 		return &TokenMetadata{
// 			Expires: expires,
// 		}, nil
// 	}

// 	return nil, err
// }

// func extractToken(c *fiber.Ctx) string {
// 	bearToken := c.Get("Authorization")

// 	// Normally Authorization HTTP header.
// 	onlyToken := strings.Split(bearToken, " ")
// 	if len(onlyToken) == 2 {
// 		return onlyToken[1]
// 	}

// 	return ""
// }

// func verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
// 	tokenString := extractToken(c)

// 	prvKey, err := ioutil.ReadFile(os.Getenv("JWT_FILE_SECRET"))
// 	if err != nil {
// 		return nil, err
// 	}
// 	pubKey, err := ioutil.ReadFile(os.Getenv("JWT_FILE_PUBLIC"))
// 	if err != nil {
// 		return nil, err
// 	}

// 	jwtToken := NewJWT(prvKey, pubKey)

// 	key, err := jwt.ParseRSAPublicKeyFromPEM(jwtToken.publicKey)
// 	if err != nil {
// 		return nil, err
// 	}

// 	type MPAClaims struct {
// 		Uuid string `json:"uuid"`
// 		jwt.StandardClaims
// 	}

// 	token, err := jwt.ParseWithClaims(tokenString, &MPAClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
// 			return nil, errors.New("signin methode failed")
// 		}
// 		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
// 		return key, nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return token, nil
// }
