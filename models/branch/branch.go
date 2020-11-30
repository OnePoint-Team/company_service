package branch

import (
	"github.com/OnePoint-Team/company_service/initDB"
	"log"
	"regexp"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

var db *gorm.DB = initDB.InitDB()

type Base struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Branch struct
type Branch struct {
	Base      Base      `gorm:"embedded"`
	Name      string    `gorm:"column:name;size:128;not null;"`
	CompanyID uuid.UUID `json:"-"`
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

// Insert function is used to insert data into database
// SECURITY ISSUES: NOT CHEKCED BEFORE INSERTION
func (branch *Branch) Insert() {
	db.Create(branch)
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
	db.First(&branch, uid)
}

// Update function is used to update data in the database
// SECURITY ISSUES: NOT CHEKCED BEFORE UPDATE
func (branch *Branch) Update() {
	db.Save(&branch)
	log.Println("Updated -> ", branch)
}

// Delete function is used to delete data into database
func (branch *Branch) Delete() {
	db.Delete(&branch)
	log.Println("Deleted -> ", branch)
}

// ############################## //

func sanitarize(id string) {
	matched, _ := regexp.MatchString("^[a-zA-Z0-9-]*$", id)
	if !matched {
		log.Fatalln("Injection occuried")
	}
}
