package model

//User store user infomation
type User struct {
	//Line below which is called embed struct
	Model

	Name       string `json:"name"`
	Email      string `json:"email"`
	HashedPass string `json:"hashed_pass"`
	Status     string `json:"status"`
	AccessCode string `json:"access_code"`
}
