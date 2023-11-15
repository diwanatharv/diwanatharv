package repository

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/authnull0/user-service/src/db"
	"github.com/authnull0/user-service/src/models/dto"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"golang.org/x/crypto/argon2"

	"github.com/authnull0/user-service/src/models"
	"gorm.io/gorm"
)

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	db := db.GetConnectiontoDatabaseDynamically(viper.GetString(viper.GetString("env") + ".db.dbname"))
	err := db.Where("email_address = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("No such mail exists")
		}
		return nil, err
	}

	return &user, nil
}

var (
	ErrInvalidHash         = errors.New("the encoded hash is not in the correct format")
	ErrIncompatibleVersion = errors.New("incompatible version of argon2")
)

var p = &dto.Params{
	Memory:      64 * 1024,
	Iterations:  3,
	Parallelism: 2,
	SaltLength:  16,
	KeyLength:   32,
}

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
func GenerateFromPassword(password string) (encodedHash string, err error) {
	salt, err := generateRandomBytes(p.SaltLength)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, p.Iterations, p.Memory, p.Parallelism, p.KeyLength)

	// Base64 encode the salt and hashed password.
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// Return a string using the standard encoded hash representation.
	encodedHash = fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, p.Memory, p.Iterations, p.Parallelism, b64Salt, b64Hash)

	return encodedHash, nil
}

func ComparePasswordAndHash(password, encodedHash string) (match bool, err error) {
	// Extract the parameters, salt and derived key from the encoded password
	// hash.
	p, salt, hash, err := decodeHash(encodedHash)
	if err != nil {
		return false, err
	}

	// Derive the key from the other password using the same parameters.
	otherHash := argon2.IDKey([]byte(password), salt, p.Iterations, p.Memory, p.Parallelism, p.KeyLength)

	// Check that the contents of the hashed passwords are identical. Note
	// that we are using the subtle.ConstantTimeCompare() function for this
	// to help prevent timing attacks.
	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}
	return false, nil
}
func decodeHash(encodedHash string) (p *dto.Params, salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, ErrInvalidHash
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, ErrIncompatibleVersion
	}

	p = &dto.Params{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.Memory, &p.Iterations, &p.Parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	p.SaltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.KeyLength = uint32(len(hash))

	return p, salt, hash, nil
}

func CreateToken(email string) (string, error) {
	// Generate a new jwt token object, specifying signing method and the claims
	// Your secret key (keep this secret in a production environment)
	secretKey := []byte("Authnull")

	// Create a claim with the email ID and expiration time
	claims := jwt.MapClaims{
		"email": email,                            // Replace with the email you want to include
		"exp":   time.Now().Add(time.Hour).Unix(), // Set the expiration time (1 hour in this example)
	}

	// Create a token with the claims and the signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with your secret key to generate the final JWT token
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println("Error creating JWT token:", err)
		return "", err
	}

	// Print the JWT token
	fmt.Println("JWT Token:", tokenString)

	return tokenString, nil
}

func VerifyToken(tokenString string) (string, error) {
	// Your secret key (keep this secret in a production environment)
	secretKey := []byte("Authnull")

	// Parse the JWT string and store the result in `token`
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// Return the secret key used to sign the token
		return secretKey, nil
	})
	if err != nil {
		return "", err
	}

	//GET THE EMAIL FROM THE TOKEN
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", nil
	}
	email := claims["email"].(string)

	return email, nil
}

func GetOrganization(email string) (*models.Organization, error) {

	var organization models.Organization
	db := db.GetConnectiontoDatabaseDynamically(viper.GetString(viper.GetString("env") + ".db.dbname"))

	err := db.Where("admin_email = ?", email).First(&organization).Error
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	return &organization, nil
}
