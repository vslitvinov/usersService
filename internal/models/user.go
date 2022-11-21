package models


type User struct {
	UUID        uint64 `json:"uuid"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Password    string `json:"password"`
}
