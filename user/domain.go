package user

import (
	"fmt"
	"strings"

	"github.com/mukul1234567/Library-Management-System/db"
)

type createRequest struct {
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

type updateRequest struct {
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

// OldPassword  string `json:"old_password"`
// NewPassword  string `json:"new_password"`
type findByIDResponse struct {
	User db.User `json:"user"`
}

type listResponse struct {
	User []db.User `json:"users"`
}

func (cr createRequest) Validate() (err error) {
	if cr.FirstName == "" || cr.LastName == "" || cr.Gender == "" || cr.Address == "" || cr.Email == "" || cr.Password == "" || cr.MobileNum == "" || cr.Role == "" {
		return errEmptyName
	}
	if (strings.ToLower(cr.Role) != "user") && (strings.ToLower(cr.Role) != "admin") && (strings.ToLower(cr.Role) != "superadmin") {
		fmt.Println(strings.ToLower(cr.Role))
		return errInvalidRole
	}
	if !isEmailValid(cr.Email) {
		return errInvalidEmail
	}
	if len(cr.MobileNum) != 10 {
		return errInvalidMobNo
	}
	if cr.Age < 0 {
		return errInvalidAge
	}
	return
}

func (cr updateRequest) Validate() (err error) {
	// if cr.FirstName == "" || cr.LastName == "" || cr.Gender == "" || cr.Address == "" || cr.Email == "" || cr.OldPassword == "" ||cr.NewPassword == "" || cr.MobileNum == "" || cr.Role == "" {
	if cr.FirstName == "" || cr.LastName == "" || cr.Gender == "" || cr.Address == "" || cr.Email == "" || cr.Password == "" || cr.MobileNum == "" || cr.Role == "" {
		return errEmptyName
	}
	return
}
