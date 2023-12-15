package tests

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"testing"

	"github.com/iamsumit/sample-go-app/pkg/api"
	"github.com/iamsumit/sample-go-app/pkg/db"
	dbtest "github.com/iamsumit/sample-go-app/pkg/db/test"
	"github.com/iamsumit/sample-go-app/pkg/logger"
	configtest "github.com/iamsumit/sample-go-app/sample/internal/config/test"
	userv1 "github.com/iamsumit/sample-go-app/sample/internal/handler/entitygrp/user/v1"
	"github.com/iamsumit/sample-go-app/sample/internal/handler/router"
	"github.com/mitchellh/mapstructure"
)

var (
	TestDBName = "sample_db_test"
)

func TestV1Routes(t *testing.T) {

	// Create the required tables in the test database.
	sqlDB, teardown := setupTestDB(t)
	defer teardown()

	//------------------------------------------------
	// Routes
	//------------------------------------------------
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	h := router.ConfigureRoutes(shutdown, nil, nil, router.Config{
		Log: logger.Default(),
		DB:  sqlDB,
	}, nil)

	t.Run("UserInvalidID", func(t *testing.T) {
		testUserInvalidID(t, h)
	})

	t.Run("UserNotFound", func(t *testing.T) {
		testUserNotFound(t, h)
	})

	t.Run("UserCreation", func(t *testing.T) {
		testUserCreation(t, h)
	})

	t.Run("UserDuplication", func(t *testing.T) {
		testUserDuplication(t, h)
	})

	t.Run("UserFound", func(t *testing.T) {
		testUserFound(t, h)
	})

	t.Run("UserList", func(t *testing.T) {
		testUserList(t, h)
	})

}

// setupTestDB creates the tables in the test database.
func setupTestDB(t *testing.T) (*sql.DB, func()) {
	//------------------------------------------------
	// Test config
	//------------------------------------------------
	cfg, err := configtest.New(t, "../../../../config")
	if err != nil {
		t.Error(err)
	}

	//------------------------------------------------
	// Database connection
	//------------------------------------------------
	sqlDB, teardown, err := dbtest.New(t, db.Config{
		Type:     db.MySQL,
		Host:     cfg.MySQL.Host,
		Port:     cfg.MySQL.Port,
		Name:     cfg.MySQL.Name,
		User:     cfg.MySQL.User,
		Password: cfg.MySQL.Password,
	}, TestDBName)
	if err != nil {
		t.Error(err)
	}

	//------------------------------------------------
	// Table creation and truncate before start.
	//------------------------------------------------

	// Create the required tables in the database.
	for _, tableName := range []string{"users", "user_settings"} {
		q := fmt.Sprintf("create table if not exists %s.%s like %s.%s", TestDBName, tableName, cfg.MySQL.Name, tableName)
		if _, err := sqlDB.Exec(q); err != nil {
			t.Error(fmt.Errorf("creating %s.%s test table: %v", TestDBName, tableName, err))
		}

		q = fmt.Sprintf("truncate %s.%s", TestDBName, tableName)
		if _, err := sqlDB.Exec(q); err != nil {
			t.Error(fmt.Errorf("truncating %s.%s test table: %v", TestDBName, tableName, err))
		}
	}

	return sqlDB, teardown
}

// testUserNotFound tests the user not found error.
func testUserNotFound(t *testing.T, h http.Handler) {
	// Create a GET request for a user that does not exist yet.
	req, err := http.NewRequest(http.MethodGet, "/v1/user/100000", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record the response
	recorder := httptest.NewRecorder()
	// Serve the request to the router
	h.ServeHTTP(recorder, req)

	// Check the status code
	if status := recorder.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}

	if body := string(recorder.Body.Bytes()); strings.Contains(body, userv1.ErrUserNotFound.Error()) != true {
		t.Errorf("body returned wrong response: expected response body %s to contain the error string %s", body, userv1.ErrUserNotFound)
	}
}

// testUserInvalidID tests the invalid user error.
func testUserInvalidID(t *testing.T, h http.Handler) {
	req, err := http.NewRequest(http.MethodGet, "/v1/user/ds", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record the response
	recorder := httptest.NewRecorder()
	// Serve the request to the router
	h.ServeHTTP(recorder, req)

	// Check the status code
	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	if body := string(recorder.Body.Bytes()); strings.Contains(body, userv1.ErrInvalidID.Error()) != true {
		t.Errorf("body returned wrong response: expected response body %s to contain the error string %s", body, userv1.ErrInvalidID)
	}
}

// testUserNotFound tests the user not found error.
func testUserCreation(t *testing.T, h http.Handler) {
	//------------------------------------------------
	// Request with invalid payload
	//------------------------------------------------

	// Create a test request with JSON payload
	invalidPayload := []byte(`{"email": "test", "date_of_birth": "1993-12"}`)
	req, err := http.NewRequest(http.MethodPost, "/v1/user", bytes.NewBuffer(invalidPayload))
	if err != nil {
		t.Fatal(err)
	}

	// Record the response
	recorder := httptest.NewRecorder()
	// Serve the request to the router
	h.ServeHTTP(recorder, req)

	// Check the status code
	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	//------------------------------------------------
	// Verify response body for validation errors
	//------------------------------------------------

	body := string(recorder.Body.Bytes())

	// A validation failure message must exists.
	errStr := "field validation failed"
	if strings.Contains(body, errStr) != true {
		t.Errorf("body returned wrong response: expected response body %s to contain the error string %s", body, errStr)
	}

	// The validation must return the following errors.
	errStr = "dateofbirth is not a valid format"
	if strings.Contains(body, errStr) != true {
		t.Errorf("body returned wrong response: expected response body %s to contain the error string %s", body, errStr)
	}

	// The validation must return the following errors.
	errStr = "email must be a valid email address"
	if strings.Contains(body, errStr) != true {
		t.Errorf("body returned wrong response: expected response body %s to contain the error string %s", body, errStr)
	}

	// The validation must return the following errors.
	errStr = "name is a required field"
	if strings.Contains(body, errStr) != true {
		t.Errorf("body returned wrong response: expected response body %s to contain the error string %s", body, errStr)
	}

	//------------------------------------------------
	// Request with valid payload
	//------------------------------------------------

	// Create a test request with JSON payload
	validPayload := []byte(
		`{
				"name": "Some Name",
				"email": "someemail@test.com", 
				"date_of_birth": "1993-12-05",
				"biography": "Some bio"
			}`,
	)

	req, err = http.NewRequest(http.MethodPost, "/v1/user", bytes.NewBuffer(validPayload))
	if err != nil {
		t.Fatal(err)
	}

	// Record the response
	recorder = httptest.NewRecorder()
	// Serve the request to the router
	h.ServeHTTP(recorder, req)

	// Check the status code
	if status := recorder.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	//------------------------------------------------
	// Verify response body for correct response
	//------------------------------------------------

	response := &api.Response{}
	if err := json.Unmarshal(recorder.Body.Bytes(), response); err != nil {
		t.Error(err)
	}

	// The data must be decoded into the user struct.
	users := responseToUsers(t, response)

	if len(users) < 1 {
		t.Errorf("body returned wrong response: expected to have atleast 1 user got %d", len(users))
		return
	}

	if users[0].ID != 1 {
		t.Errorf("body returned wrong response: expected to have the user with id 1 got %d", users[0].ID)
	}

	if users[0].Name != "Some Name" {
		t.Errorf("body returned wrong response: expected to have the user with name Some Name got %s", users[0].Name)
	}

	if *users[0].Settings.Biography != "Some bio" {
		t.Errorf("body returned wrong response: expected to have the user with bio Some bio got %s", *users[0].Settings.Biography)
	}

	if *users[0].Settings.Email != "someemail@test.com" {
		t.Errorf("body returned wrong response: expected to have the user with email someemail@test.com got %s", *users[0].Settings.Email)
	}
}

func testUserDuplication(t *testing.T, h http.Handler) {
	//------------------------------------------------
	// Request with valid payload
	//------------------------------------------------

	// Create a test request with JSON payload
	validPayload := []byte(
		`{
				"name": "Some Name",
				"email": "someemail@test.com", 
				"date_of_birth": "1993-12-05",
				"biography": "Some bio"
			}`,
	)

	req, err := http.NewRequest(http.MethodPost, "/v1/user", bytes.NewBuffer(validPayload))
	if err != nil {
		t.Fatal(err)
	}

	// Record the response
	recorder := httptest.NewRecorder()
	// Serve the request to the router
	h.ServeHTTP(recorder, req)

	// Check the status code
	if status := recorder.Code; status != http.StatusConflict {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusConflict)
	}

	if body := string(recorder.Body.Bytes()); strings.Contains(body, userv1.ErrDuplicateEmail.Error()) != true {
		t.Errorf("body returned wrong response: expected response body %s to contain the error string %s", body, userv1.ErrDuplicateEmail)
	}

}

// testUserFound tests the user detail.
func testUserFound(t *testing.T, h http.Handler) {
	//------------------------------------------------
	// Verify response body for correct response
	//------------------------------------------------

	// Create a GET request for a user detail.
	req, err := http.NewRequest(http.MethodGet, "/v1/user/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record the response
	recorder := httptest.NewRecorder()
	// Serve the request to the router
	h.ServeHTTP(recorder, req)

	// Check the status code
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	response := &api.Response{}
	if err := json.Unmarshal(recorder.Body.Bytes(), response); err != nil {
		t.Error(err)
	}

	// The data must be decoded into the user struct.
	users := responseToUsers(t, response)

	if len(users) < 1 {
		t.Errorf("body returned wrong response: expected to have atleast 1 user got %d", len(users))
		return
	}

	if users[0].ID != 1 {
		t.Errorf("body returned wrong response: expected to have the user with id 1 got %d", users[0].ID)
	}

	if users[0].Name != "Some Name" {
		t.Errorf("body returned wrong response: expected to have the user with name Some Name got %s", users[0].Name)
	}

	if *users[0].Settings.Biography != "Some bio" {
		t.Errorf("body returned wrong response: expected to have the user with bio Some bio got %s", *users[0].Settings.Biography)
	}

	if *users[0].Settings.Email != "someemail@test.com" {
		t.Errorf("body returned wrong response: expected to have the user with email someemail@test.com got %s", *users[0].Settings.Email)
	}
}

// testUserList tests the user listing.
func testUserList(t *testing.T, h http.Handler) {
	//------------------------------------------------
	// Verify that no user is returned on second page
	//------------------------------------------------

	// Create a GET request for a user listing page returning no result on given page.
	req, err := http.NewRequest(http.MethodGet, "/v1/users?page=2", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record the response
	recorder := httptest.NewRecorder()
	// Serve the request to the router
	h.ServeHTTP(recorder, req)

	// Check the status code
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	response := &api.Response{}
	if err := json.Unmarshal(recorder.Body.Bytes(), response); err != nil {
		t.Error(err)
	}

	// The data must be decoded into the user struct.
	users := responseToUsers(t, response)

	if len(users) >= 1 {
		t.Errorf("body returned wrong response: expected to have atleast 0 user got %d", len(users))
	}

	//------------------------------------------------
	// Verify that user is returned on default page
	//------------------------------------------------

	// Create a GET request for a user that does not exist yet.
	req, err = http.NewRequest(http.MethodGet, "/v1/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record the response
	recorder = httptest.NewRecorder()
	// Serve the request to the router
	h.ServeHTTP(recorder, req)

	// Check the status code
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	response = &api.Response{}
	if err := json.Unmarshal(recorder.Body.Bytes(), response); err != nil {
		t.Error(err)
	}

	// The data must be decoded into the user struct.
	users = responseToUsers(t, response)

	if users[0].ID != 1 {
		t.Errorf("body returned wrong response: expected to have the user with id 1 got %d", users[0].ID)
	}

	if users[0].Name != "Some Name" {
		t.Errorf("body returned wrong response: expected to have the user with name Some Name got %s", users[0].Name)
	}

	if *users[0].Settings.Biography != "Some bio" {
		t.Errorf("body returned wrong response: expected to have the user with bio Some bio got %s", *users[0].Settings.Biography)
	}

	if *users[0].Settings.Email != "someemail@test.com" {
		t.Errorf("body returned wrong response: expected to have the user with email someemail@test.com got %s", *users[0].Settings.Email)
	}
}

// responseToUsers converts the response data to user type.
func responseToUsers(t *testing.T, response *api.Response) []*userv1.User {
	users := make([]*userv1.User, 0)

	// The data must be decoded into the user struct.
	err := mapstructure.Decode(response.Data, &users)
	if err != nil {
		t.Error(err)
	}

	return users
}
