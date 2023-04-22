package web_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Polidoro-root/go-inventory-management/configs"
	"github.com/Polidoro-root/go-inventory-management/internal/infra/database"
	"github.com/Polidoro-root/go-inventory-management/internal/infra/web"
	"github.com/Polidoro-root/go-inventory-management/internal/testutils"
)

func setupTest(t *testing.T) *web.WebUserHandler {
	db := testutils.SetupDatabase(t)

	repository := database.NewUserRepository(db)

	return web.NewWebUserHandler(repository)
}

func TestSignIn(t *testing.T) {
	handler := setupTest(t)

	configs := configs.LoadTestEnv(t)

	reqPayload := strings.NewReader(`{"email": "jv.polidoro@outlook.com", "password": "password"}`)

	r := httptest.NewRequest(http.MethodPost, "/users/signin", reqPayload)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()

	handler.SignIn(w, r)

	res := w.Result()

	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		t.Fatalf("Expected response status %d but got %d", http.StatusCreated, res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)

	if err != nil {
		t.Fatal(err)
	}

	var response *web.SignInResponse

	err = json.Unmarshal(data, &response)

	if err != nil {
		t.Fatal(err)
	}

	jwt := web.NewJWT(configs.JWTSecret)

	payload, err := jwt.VerifyToken(response.AccessToken)

	if err != nil {
		t.Fatal(err)
	}

	if payload == nil {
		t.Fatal("payload should not be nil")
	}

	if payload.Subject == "" {
		t.Fatal("payload.Subject should not be empty")
	}
}
