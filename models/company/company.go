package company

import (
	"log"
	"regexp"

	"github.com/OnePoint-Team/company_service/initdb"
	base "github.com/OnePoint-Team/company_service/models"
	"github.com/OnePoint-Team/company_service/models/agent"
	"github.com/OnePoint-Team/company_service/models/branch"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//Company struct
type Company struct {
	Base     base.Base       `gorm:"embedded"`
	Name     string          `gorm:"column:name;size:128;not null;unique;"`
	Branches []branch.Branch `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Agents   []agent.Agent   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

//Tabler for gorm get table name
type Tabler interface {
	TableName() string
}

// TableName for change table name
func (Company) TableName() string {
	return "company"
}

// BeforeCreate method run before every create call via the ORM.
func (company *Company) BeforeCreate(db *gorm.DB) (err error) {
	uuid := uuid.NewV4()

	log.Println("UUID IS GENERATED")
	company.Base.ID = uuid
	return
}

// Select by id ############################## //
func (company *Company) Select(id string) error {

	// Chekc if all is digit or letter
	sanitarize(id)

	uid, err := uuid.FromString(id)
	if err != nil {
		log.Println("uuid.FromString Error ->", err)
		return err
	}

	// SELECT * FROM users WHERE id = id;
	result := initdb.DbInstance.First(&company, uid)
	if result.Error == nil {
		initdb.DbInstance.Preload(clause.Associations).Find(company)
	}

	return result.Error
}

// All all ############################## //
func (company *Company) All(companies *[]Company) error {
	// SELECT * FROM users WHERE id = id;
	result := initdb.DbInstance.Find(&companies)

	if result.Error != nil {
		log.Println("Data not found->", result.Error)
		return result.Error
	}
	// Preload Branches of Company.
	initdb.DbInstance.Preload(clause.Associations).Find(&companies)

	return result.Error

}

// Insert function is used to insert data into database
// SECURITY ISSUES: NOT CHEKCED BEFORE INSERTION
func (company *Company) Insert() error {
	result := initdb.DbInstance.Create(company)
	log.Println("Created -> ", result)
	return result.Error
}

// ############################## //

// Update function is used to update data in the database
// SECURITY ISSUES: NOT CHEKCED BEFORE UPDATE
func (company *Company) Update() {
	initdb.DbInstance.Save(&company)
	log.Println("Updated -> ", company)
}

// Delete function is used to delete data into database
func (company *Company) Delete() {
	initdb.DbInstance.Delete(&company)
	log.Println("Deleted -> ", company)
}

func sanitarize(id string) {
	matched, _ := regexp.MatchString("^[a-zA-Z0-9-]*$", id)
	if !matched {
		log.Fatalln("Injection occuried")
	}
}
