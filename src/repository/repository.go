package repository

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/authnull0/user-service/src/models/dto"
	"golang.org/x/crypto/argon2"

	"github.com/authnull0/user-service/src/db"
	"github.com/authnull0/user-service/src/models"
	"gorm.io/gorm"
)

type Postgress struct {
	Db *gorm.DB
}

func Postgressmanager() *Postgress {
	return &Postgress{Db: db.Makegormserver()}
}

type postgressmethods interface {
	Insert(value interface{}) (tx *gorm.DB)
	Find(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Update(values interface{}) (tx *gorm.DB)
	Delete(value interface{}, conds ...interface{}) (tx *gorm.DB)
	UniqueId()
}

func (p *Postgress) Insert(value interface{}) (tx *gorm.DB) {
	return p.Db.Create(value)
}
func (p *Postgress) Find(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	return p.Db.Find(dest, conds...)
}

func (p *Postgress) Update(values interface{}) (tx *gorm.DB) {
	return p.Db.Model(values).Updates(values)
}
func (p *Postgress) Delete(value interface{}, conds ...interface{}) (tx *gorm.DB) {
	return p.Db.Delete(value, conds)
}
func (p *Postgress) Unique() (tx *gorm.DB) {
	return p.Db.Exec("ALTER TABLE User ALTER COLUMN id SET DEFAULT nextval('User');")
}
func IsFieldNotUnique(db *gorm.DB, field string, value string) (bool, error) {
	// Get the count of records that match the field value.
	var count int64
	err := db.Model(&models.Organization{}).Where(field+" = ?", value).Count(&count).Error
	if err != nil {
		log.Print(err.Error())
		return false, err
	}

	// Return true if the count is greater than one, otherwise return false.
	return count > 1, nil
}
func GetUserByEmailForOrganization(db *gorm.DB, email string) (*models.Organization, error) {
	var user models.Organization
	err := db.Where("email = ?", email).First(&user).Error
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
