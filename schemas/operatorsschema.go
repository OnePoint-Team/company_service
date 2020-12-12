package schemas



type OperatorCreate struct {
	UserID string `json:"user_id" binding:"required,uuid4"`
}


type OperatorPathVar struct{
	LID string `uri:"lid" binding:"required,uuid4"`
}


type OperatorPath struct{
	UserID string `uri:"oid" binding:"required,uuid4"`
}



