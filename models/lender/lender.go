package lender

import (
	"fmt"
	"log"

	"github.com/OnePoint-Team/company_service/initdb"
	base "github.com/OnePoint-Team/company_service/models"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// Lender struct for loan companies
type Lender struct {
	base.Base
	Name string `gorm:"column:name" json:"name"`
}

// Tabler for gorm get table name
type Tabler interface {
	TableName() string
}

// TableName for change table name
func (Lender) TableName() string {
	return "lenders"
}

// BeforeCreate method run before every create call via the ORM.
func (l *Lender) BeforeCreate(DbInstance *gorm.DB) (err error) {
	uuid := uuid.NewV4()
	fmt.Println("UUID is generated")
	l.Base.ID = uuid
	return
}

func checkExistLender(name string) error {
	l := Lender{}
	r := initdb.DbInstance.Table("lenders").Where(&Lender{Name: name}).First(&l)
	log.Println("existing data -> ", l)
	return r.Error
}

// Select quering for Lender
func (l *Lender) Select(id string) error {
	uid, err := uuid.FromString(id)
	if err != nil {
		log.Println("error when convert uuid")
		return err
	}
	result := initdb.DbInstance.First(&l, uid)
	return result.Error
}

// Insert insert to db
func (l *Lender) Insert() error {
	if err := checkExistLender(l.Name); err == nil {
		return err
	}
	result := initdb.DbInstance.Create(&l)
	return result.Error
}

// All select all lender records
func (l *Lender) All(lenders *[]Lender) error {
	result := initdb.DbInstance.Find(lenders)
	return result.Error
}

func (l *Lender) Delete() error {
	r := initdb.DbInstance.Delete(&l)
	return r.Error
}
