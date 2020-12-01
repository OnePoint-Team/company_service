package branch

import (
	"log"
	"regexp"

	"github.com/OnePoint-Team/company_service/initdb"
	base "github.com/OnePoint-Team/company_service/models"
	"github.com/OnePoint-Team/company_service/models/agent"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// Branch struct
type Branch struct {
	Base      base.Base     `gorm:"embedded"`
	Name      string        `gorm:"column:name;size:128;not null;"`
	CompanyID uuid.UUID     `json:"-"`
	Agents    []agent.Agent `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// BeforeCreate method run before every create call via the ORM.
func (branch *Branch) BeforeCreate(db *gorm.DB) (err error) {
	uuid := uuid.NewV4()
	if err != nil {
		return err
	}
	log.Println("UUID IS GENERATED")
	branch.Base.ID = uuid
	return
}

//Tabler for gorm get table name
type Tabler interface {
	TableName() string
}

// TableName for change table name
func (Branch) TableName() string {
	return "branch"
}

// Insert function is used to insert data into database
// SECURITY ISSUES: NOT CHEKCED BEFORE INSERTION
func (branch *Branch) Insert() {
	initdb.DbInstance.Create(branch)
	log.Println("Created -> ", branch)
}

// Select by id ############################## //
func (branch *Branch) Select(id string) {

	// Chekc if all is digit or letter
	sanitarize(id)

	uid, err := uuid.FromString(id)
	if err != nil {
		log.Fatalln("Error occuried ->", err)
	}

	// SELECT * FROM users WHERE id = id;
	initdb.DbInstance.First(&branch, uid)
}

// Update function is used to update data in the database
// SECURITY ISSUES: NOT CHEKCED BEFORE UPDATE
func (branch *Branch) Update() {
	initdb.DbInstance.Save(&branch)
	log.Println("Updated -> ", branch)
}

// Delete function is used to delete data into database
func (branch *Branch) Delete() {
	initdb.DbInstance.Delete(&branch)
	log.Println("Deleted -> ", branch)
}

// ############################## //

func sanitarize(id string) {
	matched, _ := regexp.MatchString("^[a-zA-Z0-9-]*$", id)
	if !matched {
		log.Fatalln("Injection occuried")
	}
}
