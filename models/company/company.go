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
	Base     base.Base       `gorm:"embedded" serialize:"id:Base.ID,created:Base.created"`
	Name     string          `gorm:"column:name;size:128;not null;unique;" serialize:"name"`
	Branches []branch.Branch `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Agents   []agent.Agent   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

//Tabler for gorm get table name
// type Tabler interface {
// 	TableName() string
// }

// TableName for change table name
// func (Company) TableName() string {
// 	return "company"
// }

// Schema for change table name

// BeforeCreate method run before every create call via the ORM.
func (c *Company) BeforeCreate(db *gorm.DB) (err error) {
	uuid := uuid.NewV4()

	log.Println("UUID IS GENERATED")
	c.Base.ID = uuid
	return
}

// Select by id ############################## //
func (c *Company) Select(id string) error {

	// Chekc if all is digit or letter
	sanitarize(id)

	uid, err := uuid.FromString(id)
	if err != nil {
		log.Println("uuid.FromString Error ->", err)
		return err
	}

	// SELECT * FROM users WHERE id = id;
	result := initdb.DbInstance.First(&c, uid)
	if result.Error == nil {
		initdb.DbInstance.Preload(clause.Associations).Find(c)
	}

	return result.Error
}

// All all ############################## //
func (c *Company) All(companies *[]Company) error {
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
func (c *Company) Insert() error {
	result := initdb.DbInstance.Create(c)
	log.Println("Created -> ", result)
	return result.Error
}

// ############################## //

// Update function is used to update data in the database
// SECURITY ISSUES: NOT CHEKCED BEFORE UPDATE
func (c *Company) Update() {
	initdb.DbInstance.Save(&c)
	log.Println("Updated -> ", c)
}

// Delete function is used to delete data into database
func (c *Company) Delete() {
	initdb.DbInstance.Delete(&c)
	log.Println("Deleted -> ", c)
}

func sanitarize(id string) {
	matched, _ := regexp.MatchString("^[a-zA-Z0-9-]*$", id)
	if !matched {
		log.Fatalln("Injection occuried")
	}
}
