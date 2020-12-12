package operators

import (
	"errors"
	"fmt"
	"log"

	"github.com/OnePoint-Team/company_service/initdb"
	base "github.com/OnePoint-Team/company_service/models"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// class OperationTypes(Model):
//     __tablename__ = "operation_types"

//     name = Column(String(),nullable=False)
//     created = Column(Datetime(timezone=True), default=func.now())
//     updated = Column(Datetime(timezone=True), onupdate=func.now(),nullable=True)

// class Operators(Model):
//     __tablename__ = "operators"

//     lender_id = Column(UUIDType(),ForeignKey("lenders.id",use_alter=True,ondelete="SET NULL"))
//     user_id = Column(UUIDType())
//     created = Column(Datetime(timezone=True), default=func.now())
//     updated = Column(Datetime(timezone=True), onupdate=func.now(),nullable=True)

// lender_id = Column(UUIDType(),ForeignKey("lenders.id",use_alter=True,ondelete="SET NULL"))
// user_id = Column(UUIDType())

// operation type name created updated ,

type Operators struct {
	base.Base
	LenderID uuid.UUID `gorm:"column:lender_id" json:"lender_id"`
	UserID   uuid.UUID `gorm:"column:user_id" json:"user_id"`
}

func (o *Operators) BeforeCreate(DbInstance *gorm.DB) (err error) {
	uuid := uuid.NewV4()

	o.Base.ID = uuid
	return
}

func CheckExistence(id uuid.UUID) error {
	o := Operators{}
	// db.First(&user, "10")

	r := initdb.DbInstance.Table("operators").Find(&o, Operators{LenderID: id})

	// r := initdb.DbInstance.Table("operators").First(&o, id)

	// r := initdb.DbInstance.Table("operators").Where(&Operators{LenderID: id}).First(&o)
	log.Println("data found ->", o)
	

	if r.Error != nil{
		return r.Error
	}
	if r.RowsAffected >= 1 {
		return errors.New("More than one record")
	}
	return nil


}

func (o *Operators) Select(id string) error {
	fmt.Println(id)
	uid, err := uuid.FromString(id)
	if err != nil {
		log.Println("error occured")
	}
	result := initdb.DbInstance.First(&o, uid)

	return result.Error
}

func (o *Operators) Insert(lid, userid string) error {

	lenderID, _ := uuid.FromString(lid)
	userID, _ := uuid.FromString(userid)

	err := CheckExistence(userID)

	if err != nil {
		log.Println("operator exist", err)
		return err
	}

	o.UserID = userID
	o.LenderID = lenderID
	fmt.Println(o.LenderID, o.UserID)
	result := initdb.DbInstance.Create(&o)
	return result.Error
}

func (o *Operators) All(operators *[]Operators) error {
	result := initdb.DbInstance.Find(operators)
	return result.Error
}

func (o *Operators) Delete(userid string) error {
	r := initdb.DbInstance.Table("operators").Where("id = ? ", userid).Delete(o)
	if r.Error != nil{
		return r.Error
	}
	return nil
}

func (o *Operators) Update(userid string) error{
	
	r := initdb.DbInstance.Table("operators").Where("id = ? ", userid).Save(&o)

	log.Println("THEE",r)
	return r.Error
}

