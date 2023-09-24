package user

type DetailRequest struct {
	ID int `uri:"id" binding:"required"`
}
