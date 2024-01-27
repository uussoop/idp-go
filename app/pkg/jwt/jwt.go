package jwt

import (
	"crypto/rsa"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

// var secret = []byte(base64.StdEncoding.EncodeToString(utils.RandomBytes(32)))

var PrivateKey *rsa.PrivateKey

func CreateToken(Username string, t time.Duration) (string, error) {

	// Verify username and password

	token := jwt.New(jwt.SigningMethodRS512)
	claims := token.Claims.(jwt.MapClaims)

	// Set claims
	claims["username"] = Username
	claims["exp"] = time.Now().Add(t).Unix()

	// Create token
	tokenString, err := token.SignedString(PrivateKey)

	return tokenString, err
}

func ValidateToken(tokenString string) (string, bool, error) {
	pub := PrivateKey.PublicKey
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return &pub, nil
	})
	if err != nil {
		return "", false, err
	}

	// Validate claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		// log.Println(token.SignedString(secret))
		return username, true, nil

	} else {
		return "", false, errors.New("invalid token")
	}

}
func UsernameFromToken(T *string) (string, error) {
	token, err := jwt.Parse(*T, func(token *jwt.Token) (interface{}, error) {
		return PrivateKey.Public, nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if username := claims["username"].(string); username != "" {

			return username, nil
		} else {
			return "", err
		}
	} else {
		return "", err
	}
}
func CreateRefreshToken(username string) (string, error) {
	// Lookup refresh token in DB and make sure valid

	// Create new JWT token
	refreshToken, err := CreateToken(username, time.Duration(168*time.Hour))
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}

// func JWTHeader() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		//check if bearer token available if not return error

// 		var token string
// 		token = c.GetHeader("Authorization")
// 		if token == "" {
// 			c.JSON(
// 				http.StatusOK,
// 				cerror.ErrorResponseMessage{
// 					Status: cerror.ERROR_AUTH_TOKEN,
// 					Error:  "token not found",
// 				},
// 			)
// 			c.Abort()
// 			return
// 		} else {

// 			token = strings.Replace(token, "Bearer ", "", 1)
// 			username, usernameerr := UsernameFromToken(&token)
// 			if usernameerr != nil {
// 				c.JSON(http.StatusOK, cerror.ErrorResponseMessage{Status: cerror.ERROR_AUTH_TOKEN, Error: "token is curropt"})
// 				c.Abort()
// 				return
// 			}
// 			if ok, _ := models.UserExists(&username); !ok {
// 				c.JSON(http.StatusOK, cerror.ErrorResponseMessage{Status: cerror.ERROR_EXIST_USER, Error: "user doesnt exist"})
// 				c.Abort()
// 				return
// 			}

// 			valid, err := ValidateToken(token)
// 			if !valid {
// 				c.JSON(http.StatusOK, cerror.ErrorResponseMessage{Status: cerror.ERROR_AUTH_TOKEN, Error: "token not valid"})
// 				c.Abort()
// 				return
// 			}
// 			if err != nil {
// 				c.JSON(http.StatusOK, cerror.ErrorResponseMessage{Status: cerror.ERROR_AUTH_TOKEN, Error: "token check ended with error"})
// 				c.Abort()
// 				return
// 			}
// 			fmt.Println(username)
// 			user, err := models.GetUserByUsername(username)
// 			if err != nil {

// 				fmt.Println(err)
// 				c.JSON(http.StatusOK, cerror.ErrorResponseMessage{Status: cerror.ERROR_AUTH_TOKEN, Error: "token check ended with error"})
// 				c.Abort()
// 				return
// 			}
// 			if !user.Verified {
// 				fmt.Println(err)
// 				c.JSON(http.StatusOK, cerror.ErrorResponseMessage{Status: cerror.ERROR_USER_NOT_VERIFIED, Error: "user not verified"})
// 				c.Abort()
// 				return
// 			}
// 			c.Set("user", user)

// 			c.Next()
// 		}

// 	}
// }
