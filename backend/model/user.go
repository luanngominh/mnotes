package model

var (
	//UserInactive user have not verify yet
	UserInactive = uint8(0)
	//UserActive user have already verify
	UserActive = uint8(1)
	//UserDeactive account blocked
	UserDeactive = uint8(2)
)

//User store user infomation
type User struct {
	//Line below which is called embed struct
	Model

	Name       string `json:"name"`
	Email      string `json:"email"`
	HashedPass string `json:"hashed_pass"`
	Status     uint8  `json:"status"`
	Verify     string `json:"verify"`
	AccessCode string `json:"access_code"`
}
