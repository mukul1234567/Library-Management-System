package transaction

import (
	"encoding/json"
	"net/http"

	"github.com/mukul1234567/Library-Management-System/api"
	
)

func Create(service Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// api.Success(rw, http.StatusOK, api.Response{Message: "hi"})
		var c createRequest
		err := json.NewDecoder(req.Body).Decode(&c)
		if err != nil {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		err = service.create(req.Context(), c)
		if isBadRequest(err) {
			api.Error(rw, http.StatusBadRequest, api.Response{Message: err.Error()})
			return
		}

		if err != nil {
			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
			return
		}

		api.Success(rw, http.StatusCreated, api.Response{Message: "Created Successfully"})
	})
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

// func FindByID(service Service) http.HandlerFunc {
// 	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
// 		vars := mux.Vars(req)

// 		resp, err := service.findByID(req.Context(), vars["id"])

// 		if err == errNoUserId {
// 			api.Error(rw, http.StatusNotFound, api.Response{Message: err.Error()})
// 			return
// 		}
// 		if err != nil {
// 			api.Error(rw, http.StatusInternalServerError, api.Response{Message: err.Error()})
// 			return
// 		}

// 		api.Success(rw, http.StatusOK, resp)
// 	})
// }


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
	return err == errEmptyBookID || err == errEmptyUserID
}
