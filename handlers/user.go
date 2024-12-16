package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/BacataCode/SmartCardConnectApiRest/middleware"
	"github.com/BacataCode/SmartCardConnectApiRest/models"
	"github.com/BacataCode/SmartCardConnectApiRest/repository"
	"github.com/BacataCode/SmartCardConnectApiRest/responses"
	"github.com/BacataCode/SmartCardConnectApiRest/server"
	"github.com/BacataCode/SmartCardConnectApiRest/structures"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		neededRoles := []string{middleware.Admin}

		//Token validation
		user, err := middleware.ValidateToken(s, w, r)

		// Roles validation
		if err != nil || !middleware.ValidateRoles(w, neededRoles, user.Roles) {
			return
		}

		// Handle request
		w.Header().Set("Content-Type", "application/json")

		var req = structures.CreateRequest{}
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responses.BadRequest(w, "Invalid request body")
			return
		}

		_, err = repository.GetUserByEmail(r.Context(), req.Email)
		if err == nil {
			responses.BadRequest(w, "User already exists")
			return
		}

		// Hash password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			responses.NoAuthResponse(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		createUser := models.InsertUser{
			Email:    req.Email,
			Password: string(hashedPassword),
			Name:     req.Name,
			Roles:    req.Roles,
		}
		profile, err := repository.InsertUser(r.Context(), &createUser)
		if err != nil {
			responses.BadRequest(w, "Error creating user")
			return
		}

		//websocked
		neededRolesWs := []string{"admin"}
		neededModulesWs := []string{"1"}
		var planMessage = models.WebsocketMessage{
			// codes are used to identify to where (modules) and what to does the message (create, update, delete, etc.)
			Code:    "0000",
			Payload: profile,
			User:    user.Name,
		}
		s.Hub().Broadcast(planMessage, neededRolesWs, neededModulesWs)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(profile)
	}
}

func LoginHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req = structures.LoginRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responses.BadRequest(w, "Invalid request body")
			return
		}

		user, _ := repository.GetUserByEmail(r.Context(), req.Email)
		if user == nil {
			responses.NoAuthResponse(w, http.StatusUnauthorized, "Invalid credentials")
			return
		}

		// Compare passwords
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			responses.NoAuthResponse(w, http.StatusUnauthorized, "Invalid credentials")
			return
		}

		// Generate token
		claim := models.AppClaims{
			UserId: user.Id,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
		tokenString, err := token.SignedString([]byte(s.Config().JWTSecret))
		if err != nil {
			responses.NoAuthResponse(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responses.LoginResponse{
			Message: "Welcome, you are logged in!",
			Token:   tokenString,
		})
	}
}

func ProfileHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//Token validation
		profile, err := middleware.ValidateToken(s, w, r)
		if err != nil {
			return
		}

		// request
		w.Header().Set("Content-Type", "application/json")

		// Handle request
		json.NewEncoder(w).Encode(profile)
	}
}

func ListUsersHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		neededRoles := []string{middleware.Admin}

		//Token validation
		user, err := middleware.ValidateToken(s, w, r)

		// Roles validation
		if err != nil || !middleware.ValidateRoles(w, neededRoles, user.Roles) {
			return
		}

		// Handle request
		w.Header().Set("Content-Type", "application/json")
		profiles, err := repository.ListUsers(r.Context())
		if err != nil {
			responses.NotFound(w, "Error getting users")
			return
		}

		json.NewEncoder(w).Encode(profiles)
	}
}

func UpdateAnyUserHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		neededRoles := []string{middleware.Admin}

		//Token validation
		user, err := middleware.ValidateToken(s, w, r)

		// Roles validation
		if err != nil || !middleware.ValidateRoles(w, neededRoles, user.Roles) {
			return
		}

		// Handle request
		w.Header().Set("Content-Type", "application/json")
		var req = structures.UpdateUserRequest{}
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			responses.BadRequest(w, "Invalid request body")
			return
		}

		params := mux.Vars(r)
		data := models.UpdateUser{
			Id:    params["id"],
			Name:  req.Name,
			Roles: req.Roles,
		}
		updatedUser, err := repository.UpdateUser(r.Context(), data)
		if err != nil {
			responses.BadRequest(w, "Error updating user")
			return
		}

		//websocked
		neededRolesWs := []string{"admin"}
		neededModulesWs := []string{"1"}
		var planMessage = models.WebsocketMessage{
			// codes are used to identify to where (modules) and what to does the message (create, update, delete, etc.)
			Code:    "0000",
			Payload: updatedUser,
			User:    user.Name,
		}
		s.Hub().Broadcast(planMessage, neededRolesWs, neededModulesWs)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedUser)
	}
}

func DeleteUserHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		neededRoles := []string{middleware.Admin}

		//Token validation
		user, err := middleware.ValidateToken(s, w, r)

		// Roles validation
		if err != nil || !middleware.ValidateRoles(w, neededRoles, user.Roles) {
			return
		}

		// Handle request
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		err = repository.DeleteUser(r.Context(), params["id"])
		if err != nil {
			responses.BadRequest(w, "Error deleting user")
			return
		}

		//websocked
		neededRolesWs := []string{"admin"}
		neededModulesWs := []string{"1"}
		var planMessage = models.WebsocketMessage{
			// codes are used to identify to where (modules) and what to does the message (create, update, delete, etc.)
			Code:    "0000",
			Payload: params["id"],
			User:    user.Name,
		}
		s.Hub().Broadcast(planMessage, neededRolesWs, neededModulesWs)

		w.WriteHeader(http.StatusOK)
	}
}
