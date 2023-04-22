package web

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Polidoro-root/go-inventory-management/configs"
	"github.com/Polidoro-root/go-inventory-management/internal/entity"
	"github.com/Polidoro-root/go-inventory-management/internal/usecase"
)

type SignInResponse struct {
	AccessToken string `json:"access_token"`
}

type WebUserHandler struct {
	UserRepository entity.UserRepositoryInterface
}

func NewWebUserHandler(userRepository entity.UserRepositoryInterface) *WebUserHandler {
	return &WebUserHandler{
		UserRepository: userRepository,
	}
}

func (h *WebUserHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	defer r.Body.Close()

	var input usecase.UserSignInInputDTO

	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userSignIn := usecase.NewUserSignInUseCase(h.UserRepository)

	output, err := userSignIn.Execute(input)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	configs := configs.LoadConfig()

	jwt := NewJWT(configs.JWTSecret)

	token, err := jwt.GenerateToken(
		output.UserID,
		time.Now().Add(time.Second*time.Duration(configs.JWTExpiresIn)).Unix(),
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := SignInResponse{
		AccessToken: token,
	}

	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
