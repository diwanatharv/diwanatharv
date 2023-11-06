package repository

import (
	"errors"
	"log"

	"github.com/authnull0/user-service/src/db"
	"github.com/authnull0/user-service/src/models"
	"golang.org/x/crypto/bcrypt"
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
	err := db.Model(&models.User{}).Where(field+" = ?", value).Count(&count).Error
	if err != nil {
		log.Print(err.Error())
		return false, err
	}

	// Return true if the count is greater than one, otherwise return false.
	return count > 1, nil
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func GetUserByEmail(db *gorm.DB, email string) (*models.User, error) {
	var user models.User
	err := db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("No such mail exists")
		}
		return nil, err
	}

	return &user, nil
}
func Checkpassword(reqpassword string, original string) error {
	return bcrypt.CompareHashAndPassword([]byte(reqpassword), []byte(original))
}
