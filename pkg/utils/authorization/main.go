package authorization

import (
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"myponyasia.com/hub-api/app/services/roles"
	"myponyasia.com/hub-api/pkg/entities"
)

type RefreshTokenClaims struct {
	UserId string `json:"uuid"`
	jwt.StandardClaims
}

type AccessTokenClaims struct {
	UserId string `json:"uuid"`
	jwt.StandardClaims
}

var (
	PrivateKey []byte
	PublicKey  []byte
)

func init() {
	var err error
	PrivateKey, err = ioutil.ReadFile(os.Getenv("JWT_FILE_SECRET"))
	if err != nil {
	}
	PublicKey, err = ioutil.ReadFile(os.Getenv("JWT_FILE_PUBLIC"))
	if err != nil {
	}
}

func GenerateNewAccessToken(user entities.User) (string, string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(PrivateKey)
	if err != nil {
		return "", "", err
	}

	// Set expires minutes count for secret key from .env file.
	atExpired, _ := strconv.Atoi(os.Getenv("JWT_ACCESS_TOKEN_EXPIRED"))
	rtExpired, _ := strconv.Atoi(os.Getenv("JWT_REFRESH_TOKEN_EXPIRED"))

	userRoles, err := roles.FindRolesByUserId(user)
	if err != nil {
		return "", "", err
	}

	// Create a roles name
	var roles []string
	for _, v := range userRoles {
		roles = append(roles, v.Role.Name)
	}

	// Create a new access token claims.
	atClaims := jwt.MapClaims{
		"typ":  "session",
		"iss":  "myponyasia.com",
		"aud":  "myponyasia.com",
		"iat":  time.Now().Unix(),
		"nbf":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Minute * time.Duration(atExpired)).Unix(),
		"uuid": user.ID.String(),
		"data": fiber.Map{
			"roles":       roles,
			"permissions": []string{},
		},
	}

	// Create a new JWT access token with claims.
	atToken := jwt.NewWithClaims(jwt.SigningMethodRS256, atClaims)
	at, err := atToken.SignedString(key)
	if err != nil {
		return "", "", err
	}

	// Create a new refresh token claims.
	rtClaims := jwt.MapClaims{
		"typ":  "refresh",
		"iss":  "myponyasia.com",
		"aud":  "myponyasia.com",
		"iat":  time.Now().Unix(),
		"nbf":  time.Now().Unix(),
		"exp":  time.Now().Add((time.Hour * 24) * time.Duration(rtExpired)).Unix(),
		"uuid": user.ID.String(),
	}

	// Create a new JWT refresh token with claims.
	rtToken := jwt.NewWithClaims(jwt.SigningMethodRS256, rtClaims)
	rt, err := rtToken.SignedString(key)
	if err != nil {
		return "", "", err
	}

	return at, rt, nil
}

// func jwtTokenRead(inToken interface{}) (interface{}, error) {
// 	publicRSA, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
// 	if err != nil {
// 		return nil, err
// 	}
// 	token, err := jwt.Parse(inToken.(string), func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
// 			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return publicRSA, err
// 	})

// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		return claims, nil
// 	} else {
// 		return nil, err
// 	}
// }

// func getTokenRemainingValidity(timestamp interface{}) int {
// 	expireOffset := 0
// 	if validity, ok := timestamp.(float64); ok {
// 		tm := time.Unix(int64(validity), 0)
// 		remainder := tm.Sub(time.Now())
// 		if remainder > 0 {
// 			fmt.Println(remainder)
// 			return int(remainder.Seconds()) + expireOffset
// 		}
// 	}
// 	return expireOffset
// }
