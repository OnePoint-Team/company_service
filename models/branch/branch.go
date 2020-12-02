package branch

import (
	"errors"
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
func (branch *Branch) Insert(id string) error {

	uid, err := uuid.FromString(id)
	if err != nil {
		log.Fatalln("Error occuried ->", err)
		return err
	}
	log.Println("before select -> ", branch)
	if err := checkExistBranch(branch.Name, uid); err != nil {
		log.Println("Selected -> ", branch)
		branch.CompanyID = uid
		initdb.DbInstance.Create(branch)
		log.Println("Err -> ", err)
		return nil
	}
	return errors.New("Branch already exist")

}

// Select by id ############################## //
func (branch *Branch) Select(bid, cid string) error {

	// Chekc if all is digit or letter
	sanitarize(bid)

	branchID, err := uuid.FromString(bid)
	if err != nil {
		log.Fatalln("Error occuried ->", err)
		return err
	}
	companyID, err := uuid.FromString(cid)
	if err != nil {
		log.Fatalln("Error occuried ->", err)
		return err
	}

	r := initdb.DbInstance.Table("branch").Where("id = ? AND company_id = ?", branchID, companyID).First(&branch)
	if errors.Is(r.Error, gorm.ErrRecordNotFound) {
		log.Println(r.Error)
		return r.Error
	}
	return nil
}

// All fetch all branch with foreignkey
func (branch *Branch) All(branches *[]Branch, id string) {
	uid, err := uuid.FromString(id)
	if err != nil {
		log.Fatalln("Error occuried ->", err)
		return
	}
	r := initdb.DbInstance.Where(&Branch{CompanyID: uid}).Find(&branches)
	if r.Error != nil {
		log.Println(r.Error)
	}

}

func checkExistBranch(name string, id uuid.UUID) error {
	b := Branch{}
	r := initdb.DbInstance.Table("branch").Where(&Branch{Name: name, CompanyID: id}).First(&b)
	log.Println("Exist data ->> ", b)
	return r.Error
}

// Update function is used to update data in the database
// SECURITY ISSUES: NOT CHEKCED BEFORE UPDATE
func (branch *Branch) Update() {
	initdb.DbInstance.Save(&branch)
	log.Println("Updated -> ", branch)
}

// Delete function is used to delete data into database
func (branch *Branch) Delete(bid, cid string) {
	initdb.DbInstance.Table("branch").Where("id = ? AND company_id = ?", bid, cid).Delete(branch)
	// initdb.DbInstance.Delete(&branch)
	log.Println("Deleted -> ", branch)
}

// ############################## //

func sanitarize(id string) {
	matched, _ := regexp.MatchString("^[a-zA-Z0-9-]*$", id)
	if !matched {
		log.Fatalln("Injection occuried")
	}
}
