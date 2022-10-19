package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"regexp"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/mukul1234567/Library-Management-System/api"
	"github.com/mukul1234567/Library-Management-System/db"
	"golang.org/x/crypto/bcrypt"
)

type JWTClaim struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}
type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var jwtKey = []byte("yu78jhe5$r")
var id_auth, email_auth, role_auth string
var values_role = map[string]int{"user": 2, "admin": 1, "superadmin": 0}

// values_role["user"] = 2
// values_role["admin"] = 1
// values_role["superadmin"] = 0
var v db.User
var count int = 0

func Login(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// fmt.Println("Hi")
		// var j JWTClaim
		var x Authentication
		// fmt.Println("Hiii")
		resp, err1 := service.List(req.Context())
		// fmt.Println("Hello")
		if err1 == errNoUsers {
			// api.Success(rw, http.StatusCreated, api.Response{Message: "Enter List"})
			// fmt.Println("Hello1")
			api.Error(rw, http.StatusNotFound, api.Response{Message: err1.Error()})
			return
		}
		if err1 != nil {
			// fmt.Println("Hello2")
			api.Success(rw, http.StatusCreated, api.Response{Message: "Enter List"})
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err1.Error()})
			return
		}

		// api.Success(rw, http.StatusOK, resp)

		err := json.NewDecoder(req.Body).Decode(&x)
		// fmt.Println("Hello00")
		if err != nil {
			// fmt.Println("Hello3")
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}
		var jwtString string
		var isValid int
		fmt.Println(x.Email, x.Password)
		for _, v = range resp.User {
			// fmt.Println("Hello4")
			fmt.Println(v.Email, v.Password)
			// fmt.Println(v)
			// fmt.Print(v.Role, v.ID)
			// fmt.Println(CheckPasswordHash(x.Password, v.Password))
			if v.Email == x.Email {
				if CheckPasswordHash(x.Password, v.Password) {
					isValid = 0
					fmt.Println("Found")
					jwtString, err = GenerateJWT(v.ID, v.Email, v.Role)
					break
				}
			}
		}
		if isValid != 0 {
			// fmt.Println("Hello5")
			return
		}

		if err != nil {
			// fmt.Println("Hello6")
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		api.Success(rw, http.StatusCreated, jwtString)
		fmt.Println(jwtString)

	})
}
func CheckPasswordHash(password, hash string) bool {
	fmt.Println("Hello00")
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	fmt.Println(err)
	return err == nil
}
func Authorize(handler http.HandlerFunc, role int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")
		if tokenString != "" {
			fmt.Println(tokenString)
			token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
			if err != nil {
				fmt.Printf("Error %s", err)
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				fmt.Println(claims)
				id_auth = fmt.Sprint(claims["id"])
				email_auth = fmt.Sprint(claims["email"])
				role_auth = fmt.Sprint(claims["role"])

				fmt.Println("Id = ", id_auth)
				fmt.Println("Email =", email_auth)
				fmt.Println("Role =", role_auth)

			} else {
				api.Success(w, http.StatusCreated, api.Response{Message: "Invalid Token String"})
			}
			fmt.Println(values_role[role_auth], role)
			if values_role[role_auth] > role {
				fmt.Println("You dont have access to perform this action")
				return
			}

			handler.ServeHTTP(w, r)
		} else {
			api.Success(w, http.StatusOK, api.Response{Message: "Please Enter token string"})
		}
	}
}

func GenerateJWT(ID string, Email string, Role string) (tokenString string, err error) {

	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		ID:    ID,
		Email: Email,
		Role:  Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func Create(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var c CreateRequest
		err := json.NewDecoder(req.Body).Decode(&c)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}
		err = c.Validate()

		if err != nil {
			// fmt.Println("error1")
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		err = service.Create(req.Context(), c)
		if isBadRequest(err) {
			// fmt.Println("error2")
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		if err != nil {
			// fmt.Println("error3")
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		api.Success(rw, http.StatusCreated, api.Response{Message: "Created Successfully"})
	})
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func List(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// api.Success(rw, http.StatusCreated, api.Response{Message: "Enter List"})
		resp, err := service.List(req.Context())
		if err == errNoUsers {
			// api.Success(rw, http.StatusCreated, api.Response{Message: "Enter List"})
			api.Error(rw, http.StatusNotFound, api.Response{Message: err.Error()})
			return
		}
		if err != nil {
			api.Success(rw, http.StatusCreated, api.Response{Message: "Enter List"})
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		api.Success(rw, http.StatusOK, resp)
	})
}

func Show(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// api.Success(rw, http.StatusCreated, api.Response{Message: "Enter List"})
		resp, err := service.Show(req.Context())
		if err == errNoUsers {
			// api.Success(rw, http.StatusCreated, api.Response{Message: "Enter List"})
			api.Error(rw, http.StatusNotFound, api.Response{Message: err.Error()})
			return
		}
		if err != nil {
			api.Success(rw, http.StatusCreated, api.Response{Message: "Enter List"})
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		api.Success(rw, http.StatusOK, resp)
	})
}

func FindByID(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		resp, err := service.FindByID(req.Context(), vars["id"])

		if err == errNoUserId {
			api.Error(rw, http.StatusNotFound, api.Response{Message: err.Error()})
			return
		}
		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		api.Success(rw, http.StatusOK, resp)
	})
}

func DeleteByID(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		err := service.DeleteByID(req.Context(), vars["id"])
		if err == errNoUserId {
			api.Error(rw, http.StatusNotFound, api.Response{Message: err.Error()})
		}
		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		api.Success(rw, http.StatusOK, api.Response{Message: "Deleted Successfully"})
	})
}

func Update(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var c UpdateRequest

		err := json.NewDecoder(req.Body).Decode(&c)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		err = service.Update(req.Context(), c)
		if isBadRequest(err) {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		api.Success(rw, http.StatusOK, api.Response{Message: "Updated Successfully"})
	})
}

func UpdatePassword(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var c UpdatePasswordStruct
		resp, err := service.List(req.Context())
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		err = json.NewDecoder(req.Body).Decode(&c)

		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}
		count = 0
		for _, v = range resp.User {
			if v.ID == c.ID {
				fmt.Println(c.Password, v.Password)
				if CheckPasswordHash(c.Password, v.Password) {
					count = 1
					err = service.UpdatePassword(req.Context(), c)
					if isBadRequest(err) {
						api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
						return
					}

					if err != nil {
						api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
						return
					}

					api.Success(rw, http.StatusOK, api.Response{Message: "Updated Successfully"})
				}
			}
		}
		if count != 1 {
			api.Success(rw, http.StatusOK, api.Response{Message: "Invalid Details"})
		}

	})
}

func isBadRequest(err error) bool {
	return err == errEmptyName || err == errEmptyID
}

/*
{
    "duedate":52,
    "book_id": "501cf224-5009-4d6f-b522-b7b84feeee09",
    "user_id": "d563e110-ac05-4904-be9c-1cbf42939833"
}*/
