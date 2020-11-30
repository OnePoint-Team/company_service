package agent

import (
	"log"
	"regexp"
	"time"

	"github.com/OnePoint-Team/company_service/initDB"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

var db *gorm.DB = initDB.InitDB()

type Base struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Agent struct
type Agent struct {
	Base      Base      `gorm:"embedded"`
	CompanyID uuid.UUID `json:"-"`
	BranchID  uuid.UUID `json:"-"`
	UserID    uuid.UUID `json:"-"`
}

// BeforeCreate method run before every create call via the ORM.
func (a *Agent) BeforeCreate(db *gorm.DB) (err error) {
	uuid := uuid.NewV4()
	if err != nil {
		return err
	}
	log.Println("UUID IS GENERATED")
	a.Base.ID = uuid
	return
}

// Insert function is used to insert data into database
// SECURITY ISSUES: NOT CHEKCED BEFORE INSERTION
func (a *Agent) Insert() {
	db.Create(a)
	log.Println("Created -> ", a)
}

// Select by id ############################## //
func (a *Agent) Select(id string) {

	// Chekc if all is digit or letter
	sanitarize(id)

	uid, err := uuid.FromString(id)
	if err != nil {
		log.Fatalln("Error occuried ->", err)
	}

	// SELECT * FROM users WHERE id = id;
	db.First(&a, uid)
}

// Update function is used to update data in the database
// SECURITY ISSUES: NOT CHEKCED BEFORE UPDATE
func (a *Agent) Update() {
	db.Save(&a)
	log.Println("Updated -> ", a)
}

// Delete function is used to delete data into database
func (a *Agent) Delete() {
	db.Delete(&a)
	log.Println("Deleted -> ", a)
}

// ############################## //

func sanitarize(id string) {
	matched, _ := regexp.MatchString("^[a-zA-Z0-9-]*$", id)
	if !matched {
		log.Fatalln("Injection occuried")
	}
}
