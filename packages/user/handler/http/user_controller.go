package delivery

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"restful_api/entities"
	"restful_api/packages/user"
	"restful_api/utils"
	"strconv"
)

// MakeUserHandler - define router
func MakeUserHandler(router *http.ServeMux, service *user.Service) {
	router.Handle("/api/users", FetchAll(service))      //GET /api/users
	router.Handle("/api/user", GetByID(service))        //GET /api/user?id=2
	router.Handle("/api/users/add", Store(service))     //POST /api/users/add
	router.Handle("/api/users/update", Update(service)) //PUT /api/users/update?id=6
	router.Handle("/api/users/delete", Delete(service)) //DELETE /api/users/delete?id=9
}

// FetchAll - Get All Users
func FetchAll(s *user.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !utils.CheckHTTPMethod("GET", r) {
			utils.Respond(w, utils.Message(http.StatusMethodNotAllowed, "Method Not Allowed"))
			return
		}

		users, err := s.FetchAll()
		if err != nil {
			utils.Respond(w, utils.Message(http.StatusInternalServerError, err.Error()))
			return
		}

		for _, user := range users {
			user.Password = ""
		}
		resp := utils.Message(http.StatusOK, "Get Users info successfully")
		resp["users"] = users
		utils.Respond(w, resp)
	})
}

// GetByID - Get User by ID
func GetByID(s *user.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !utils.CheckHTTPMethod("GET", r) {
			utils.Respond(w, utils.Message(http.StatusMethodNotAllowed, "Method Not Allowed"))
			return
		}

		userID, err := strconv.ParseInt(r.URL.Query()["id"][0], 10, 64)
		if err != nil {
			utils.Respond(w, utils.Message(false, "ID is not valid"))
		}
		u, err := s.FindByID(uint(userID))
		if err != nil {
			utils.Respond(w, utils.Message(false, err.Error()))
			return
		}
		if u.ID == 0 {
			utils.Respond(w, utils.Message(http.StatusNotFound, "User Not Found"))
			return
		}
		u.Password = ""
		resp := utils.Message(http.StatusOK, "Get User info Successfully!")
		resp["user"] = u
		utils.Respond(w, resp)
	})
}

// Store - Create new User
func Store(s *user.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !utils.CheckHTTPMethod("POST", r) {
			utils.Respond(w, utils.Message(http.StatusMethodNotAllowed, "Method Not Allowed"))
			return
		}

		user := entities.User{}
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			log.Println(err)
			utils.Respond(w, utils.Message(http.StatusBadRequest, errors.New("Bad Request").Error()))
			return
		}

		if ok, err := s.Store(&user); err != nil {
			utils.Respond(w, utils.Message(ok, err.Error()))
			return
		}

		utils.Respond(w, utils.Message(http.StatusCreated, string("Saved")))
		return
	})
}

// Update - Update user by userID
func Update(s *user.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !utils.CheckHTTPMethod("PUT", r) {
			utils.Respond(w, utils.Message(http.StatusMethodNotAllowed, "Method Not Allowed"))
			return
		}

		userID, err := strconv.ParseInt(r.URL.Query()["id"][0], 10, 64)
		if err != nil {
			utils.Respond(w, utils.Message(false, "ID is not valid"))
		}

		userParam := entities.User{}
		if err := json.NewDecoder(r.Body).Decode(&userParam); err != nil {
			log.Println(err)
			utils.Respond(w, utils.Message(http.StatusBadRequest, errors.New("Bad Request").Error()))
			return
		}
		user, err := s.FindByID(uint(userID))
		if err != nil {
			utils.Respond(w, utils.Message(http.StatusBadRequest, "User not found!"))
			return
		}

		if userParam.Username != "" {
			user.Username = userParam.Username
		}
		if len(userParam.Password) > 0 {
			user.Password, err = s.HashPassword(userParam.Password)
			if err != nil {
				utils.Respond(w, utils.Message(http.StatusBadRequest, err.Error()))
				return
			}
		}
		user.IsAdmin = userParam.IsAdmin

		if ok, err := s.Update(user); err != nil {
			utils.Respond(w, utils.Message(ok, err.Error()))
			return
		}

		utils.Respond(w, utils.Message(http.StatusOK, string("Saved")))
		return
	})
}

// Delete book by bookID
func Delete(s *user.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !utils.CheckHTTPMethod("DELETE", r) {
			utils.Respond(w, utils.Message(http.StatusMethodNotAllowed, "Method Not Allowed"))
			return
		}

		userID, err := strconv.ParseInt(r.URL.Query()["id"][0], 10, 64)
		if err != nil {
			utils.Respond(w, utils.Message(false, "ID is not valid"))
		}

		if ok, err := s.Delete(uint(userID)); err != nil {
			utils.Respond(w, utils.Message(ok, err.Error()))
			return
		}

		utils.Respond(w, utils.Message(http.StatusOK, "Deleted"))
		return
	})
}
