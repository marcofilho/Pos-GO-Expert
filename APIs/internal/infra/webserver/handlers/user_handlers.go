package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/marcofilho/Pos-GO-Expert/APIs/internal/dto"
	"github.com/marcofilho/Pos-GO-Expert/APIs/internal/entity"
	"github.com/marcofilho/Pos-GO-Expert/APIs/internal/infra/database"
)

type UserHandler struct {
	UserDB       database.UserInterface
	Jwt          *jwtauth.JWTAuth
	JwtExpiresIn int
}

func NewUserHandler(userDB database.UserInterface, jwt *jwtauth.JWTAuth, jwtExpriesIn int) *UserHandler {
	return &UserHandler{
		UserDB:       userDB,
		Jwt:          jwt,
		JwtExpiresIn: jwtExpriesIn,
	}
}

func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	var userLogin dto.GetJWTInput

	error := json.NewDecoder(r.Body).Decode(&userLogin)
	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.UserDB.FindByEmail(userLogin.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !user.ValidatePassword(userLogin.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, tokenString, _ := h.Jwt.Encode(map[string]interface{}{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(h.JwtExpiresIn)).Unix(),
	})

	accessToken := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: tokenString,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)

}

func (u *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userInput dto.CreateUserInput

	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := entity.NewUser(userInput.Name, userInput.Email, userInput.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = u.UserDB.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
