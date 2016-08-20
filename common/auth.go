package common

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

//Using asymmetric crypto/RSA keys
const (
	//openssl genrsa -out app.rsa 1024
	privKeyPath = "keys/app.rsa"
	//openssl rsa -in app.rsa -pubout > app.rsa.pub
	pubKeyPath = "keys/app.rsa.pub"
)

//Private key for signin and public key for verification
var (
	//verifyKey, signKey []byte
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

//UserInfo type for Claims
type UserInfo struct {
	Email string
	Role  string
}

//AdminClaims struct
type AdminClaims struct {
	UserInfo
	jwt.StandardClaims
}

//read the key files before starting http handlers
func initKeys() {
	signBytes, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
	}
}

//GenerateJWT generates a JWT Token with claims
func GenerateJWT(email, role string) (string, error) {
	//Create the claims
	claims := AdminClaims{
		UserInfo{
			Email: email,
			Role:  role,
		},
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 20).Unix(),
			Issuer:    "admin",
		},
	}
	//create signer for rsa 256 with claims
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

//Authorize func is a middleware for validating JWT tokens
func Authorize(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	//validate the tokens
	tokenString, err := request.AuthorizationHeaderExtractor.ExtractToken(r)
	if err != nil {
		log.Fatalf("[Authorize]: %s\n", err)
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return verifyKey, nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		next(w, r)
	} else {
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorExpired: //JWT expired
				DisplayAppError(
					w,
					err,
					"Access Token is expired, get a new Token",
					401,
				)
				return
			default:
				DisplayAppError(
					w,
					err,
					"Error while parsing the Access Token!",
					500,
				)
				return
			}
		default:
			DisplayAppError(
				w,
				err,
				"Error while parsing the Access Token!",
				500,
			)
			return

		}
	}
}
