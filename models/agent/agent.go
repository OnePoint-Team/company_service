package agent

import (
	"errors"
	"log"
	"regexp"

	"github.com/OnePoint-Team/company_service/initdb"
	base "github.com/OnePoint-Team/company_service/models"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// Agent struct
type Agent struct {
	base.Base
	CompanyID uuid.UUID `gorm:"column:company_id" json:"company_id"`
	BranchID  uuid.UUID `gorm:"column:branch_id" json:"branch_id"`
	UserID    uuid.UUID `gorm:"column:user_id" json:"user_id"`
}

//Tabler for gorm get table name
type Tabler interface {
	TableName() string
}

// TableName for change table name
func (Agent) TableName() string {
	return "agents"
}

// BeforeCreate method run before every create call via the ORM.
func (a *Agent) BeforeCreate(DbInstance *gorm.DB) (err error) {
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
func (a *Agent) Insert(cid, bid, uid string) error {
	if err := checkExistAgent(cid, bid, uid); err == nil {
		log.Println("agent exist")
		return errors.New("agent already exist")
	}
	companyID, _ := uuid.FromString(cid)
	branchID, _ := uuid.FromString(bid)
	userID, _ := uuid.FromString(uid)

	a.CompanyID = companyID
	a.BranchID = branchID
	a.UserID = userID
	r := initdb.DbInstance.Create(a)
	log.Println("Created -> ", a)
	if r.Error != nil {
		log.Println(r.Error)
	}
	return r.Error
}

// All fetch all agents from db
func (a *Agent) All(agents *[]Agent, cid, bid string) {
	query := initdb.DbInstance.Table("agents").Where("company_id = ? AND branch_id = ?", cid, bid).Find(&agents)
	if query.Error != nil {
		log.Println(query.Error)
	}
}

func checkExistAgent(cid, bid, uid string) error {
	a := Agent{}
	query := initdb.DbInstance.Table("agents").Where("company_id = ? AND branch_id = ? AND user_id= ?", cid, bid, uid).First(&a)
	log.Println("Exist data ->> ", a)
	return query.Error
}

// Select by id ############################## //
func (a *Agent) Select(aid, cid, bid string) error {

	q := initdb.DbInstance.First(&a, "id = ? and company_id = ? and branch_id = ?", aid, cid, bid)
	return q.Error
}

// Update function is used to update data in the database
// SECURITY ISSUES: NOT CHEKCED BEFORE UPDATE
func (a *Agent) Update() error {
	r := initdb.DbInstance.Save(&a)
	log.Println("Updated -> ", a)
	return r.Error
}

// Delete function is used to delete data into database
func (a *Agent) Delete() error {
	r := initdb.DbInstance.Delete(&a)
	log.Println("Deleted -> ", a)
	return r.Error
}

// ############################## //

func sanitarize(id string) {
	matched, _ := regexp.MatchString("^[a-zA-Z0-9-]*$", id)
	if !matched {
		log.Fatalln("Injection occuried")
	}
}
