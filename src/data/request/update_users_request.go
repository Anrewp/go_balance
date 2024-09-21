package request

type UpdateUsersRequest struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}
