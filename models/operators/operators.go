package operators

import (
	"fmt"
	"github.com/satori/go.uuid"
	"log"
	"github.com/OnePoint-Team/company_service/initdb"
	base "github.com/OnePoint-Team/company_service/models"

)


type  Operators struct {
	LenderID uuid.UUID `json:lender_id`
}
 


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