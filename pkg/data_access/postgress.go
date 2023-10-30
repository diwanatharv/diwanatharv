package data_access

import (
	"awesomeProject12/pkg/config/postgress"
	"awesomeProject12/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type Postgress struct {
	Db *gorm.DB
}

func Postgressmanager() *Postgress {
	return &Postgress{Db: postgress.Makegormserver()}
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
