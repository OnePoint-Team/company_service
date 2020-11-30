package company

import (
	"log"
	"regexp"
	"time"

	"github.com/OnePoint-Team/company_service/initDB"
	"github.com/OnePoint-Team/company_service/models/agent"
	"github.com/OnePoint-Team/company_service/models/branch"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var db *gorm.DB = initDB.InitDB()

type Base struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;"`
	CreatedAt time.Time `json:"created"`
	UpdatedAt time.Time `json:"updated"`
}

type Company struct {
	Base     Base            `gorm:"embedded"`
	Name     string          `gorm:"column:name;size:128;not null;unique;"`
	Branches []branch.Branch `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Agents   []agent.Agent   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// BeforeCreate method run before every create call via the ORM.
func (company *Company) BeforeCreate(db *gorm.DB) (err error) {
	uuid := uuid.NewV4()

	log.Println("UUID IS GENERATED")
	company.Base.ID = uuid
	return
}

// Select by id ############################## //
func (company *Company) Select(id string) *gorm.DB {

	// Chekc if all is digit or letter
	sanitarize(id)

	uid, err := uuid.FromString(id)
	if err != nil {
		log.Println("uuid.FromString Error ->", err)
	}

	// SELECT * FROM users WHERE id = id;
	result := db.First(&company, uid)
	if result.Error == nil {
		db.Preload(clause.Associations).Find(company)
	}

	return result
}

// SelectAll all ############################## //
func (company *Company) SelectAll(companies *[]Company) *gorm.DB {
	// SELECT * FROM users WHERE id = id;
	result := db.Find(&companies)

	if result.Error != nil {
		log.Println("Data not found->", result.Error)
	} else {
		// Preload Branches of Company.
		db.Preload(clause.Associations).Find(&companies)
	}

	return result

}

// Insert function is used to insert data into database
// SECURITY ISSUES: NOT CHEKCED BEFORE INSERTION
func (company *Company) Insert() *gorm.DB {
	result := db.Create(company)
	log.Println("Created -> ", result)
	return result
}

// ############################## //

// Update function is used to update data in the database
// SECURITY ISSUES: NOT CHEKCED BEFORE UPDATE
func (company *Company) Update() {
	db.Save(&company)
	log.Println("Updated -> ", company)
}

// Delete function is used to delete data into database
func (company *Company) Delete() {
	db.Delete(&company)
	log.Println("Deleted -> ", company)
}

func sanitarize(id string) {
	matched, _ := regexp.MatchString("^[a-zA-Z0-9-]*$", id)
	if !matched {
		log.Fatalln("Injection occuried")
	}
}
