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

func Login(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// var j JWTClaim
		var x Authentication
		resp, err1 := service.list(req.Context())

		if err1 == errNoUsers {
			// api.Success(rw, http.StatusCreated, api.Response{Message: "Enter List"})
			api.Error(rw, http.StatusNotFound, api.Response{Message: err1.Error()})
			return
		}
		if err1 != nil {
			api.Success(rw, http.StatusCreated, api.Response{Message: "Enter List"})
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err1.Error()})
			return
		}

		// api.Success(rw, http.StatusOK, resp)

		err := json.NewDecoder(req.Body).Decode(&x)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}
		var jwtString string
		var isValid int
		for _, v = range resp.User {
			fmt.Println(x.Email, x.Password)
			// fmt.Print(v.Role, v.ID)
			if v.Email == x.Email && v.Password == x.Password {
				isValid = 0
				fmt.Println("Found")
				jwtString, err = GenerateJWT(v.ID, v.Email, v.Role)
				break
			}

		}
		if isValid != 0 {
			return
		}

		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		api.Success(rw, http.StatusCreated, jwtString)
		fmt.Println(jwtString)

	})
}

func Authorize(handler http.HandlerFunc, role int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")
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

		}
		fmt.Println(values_role[role_auth], role)
		if values_role[role_auth] > role {
			fmt.Println("You dont have access to perform this action")
			return
		}

		handler.ServeHTTP(w, r)
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
		var c createRequest
		err := json.NewDecoder(req.Body).Decode(&c)
		if err != nil {
			// fmt.Println("error1")
			api.Error(rw, http.StatusBadRequest, api.Response{Message: "Invalid Datatype Entered"})
			return
		}

		err = service.create(req.Context(), c)
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
		resp, err := service.list(req.Context())
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

		resp, err := service.findByID(req.Context(), vars["id"])

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

		err := service.deleteByID(req.Context(), vars["id"])
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
		var c updateRequest

		err := json.NewDecoder(req.Body).Decode(&c)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		err = service.update(req.Context(), c)
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

func isBadRequest(err error) bool {
	return err == errEmptyName || err == errEmptyID
}
