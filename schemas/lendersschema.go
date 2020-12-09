package schemas

type LenderCreate struct {
	Name string `json:"name" binding:"required"`
}
