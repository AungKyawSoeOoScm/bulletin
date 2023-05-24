package request

type UpdatePostsRequest struct {
	Id          int    `validate:"required"`
	Title       string `validate:"required,min=1,max=200" json:"name"`
	Description string `validate:"required" json:"description"`
	Status      *int   `json:"status"`
}
