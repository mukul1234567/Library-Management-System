package user

import (
	"fmt"
	"strings"

	"github.com/mukul1234567/Library-Management-System/db"
)

type CreateRequest struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	Age       int    `json:"age"`
	Address   string `json:"address"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	MobileNum string `json:"mob_no"`
	Role      string `json:"role"`
}

type UpdateRequest struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	Age       int    `json:"age"`
	Address   string `json:"address"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	MobileNum string `json:"mob_no"`
	Role      string `json:"role"`
}

type UpdatePasswordStruct struct {
	ID          string `json:"id"`
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}

type FindByIDResponse struct {
	User db.User `json:"user"`
}

type ListResponse struct {
	User []db.User `json:"users"`
}

type ListResponser struct {
	Users []db.Userlist `json:"users"`
}

func (cr CreateRequest) Validate() (err error) {
	if cr.FirstName == "" || cr.LastName == "" || cr.Gender == "" || cr.Address == "" || cr.Email == "" || cr.Password == "" || cr.MobileNum == "" || cr.Role == "" {
		return errEmptyName
	} else if (strings.ToLower(cr.Role) != "user") && (strings.ToLower(cr.Role) != "admin") && (strings.ToLower(cr.Role) != "superadmin") {
		fmt.Println(strings.ToLower(cr.Role))
		return errInvalidRole
	} else if !isEmailValid(cr.Email) {
		return errInvalidEmail
	} else if len(cr.MobileNum) != 10 {
		return errInvalidMobNo
	} else if cr.Age < 0 {
		return errInvalidAge
	}
	return
}

func (cr UpdateRequest) Validate() (err error) {
	if cr.FirstName == "" || cr.LastName == "" || cr.Gender == "" || cr.Address == "" || cr.Email == "" || cr.MobileNum == "" {
		return errEmptyName
	}
	return
}
