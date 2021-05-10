package test

import (
	"encoding/json"

	// "fmt"
	// "fmt"
	"net/http"

	// "net/http/httptest"
	"Friend_management/services"
	"testing"

	"Friend_management/models"
	// "Friend_management/repository"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	// "github.com/go-chi/chi"
	// "Friend_management/db"

	"Friend_management/db"
	"bytes"
	"context"
	"log"
	"net"
	"os"
	"time"
	// r_Response "Friend_management/models/response"
	// "github.com/go-chi/chi"
)

func WrongMin(x, y float64) float64 {
	if x > y {
		return x
	} else {
		return y
	}
}

// func TestMathBasics(t *testing.T) {
// 	v := WrongMin(10, 0)
// 	if v != 0 {
// 		t.Errorf("Failed the test!")
// 	}
// }

// func TestL(t *testing.T){
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/",services.GetAllUsers)
// 	writer := httptest.NewRecorder()
// 	request, _ := http.NewRequest("GET","/",nil)
// 	mux.ServeHTTP(writer, request)
// 	if writer.Code !=201{
// 		t.Errorf("Response code is %v", writer.Code)
// 	}
// }

func Greet(name string) string {
	return "Hello, " + name
}
func TestGreet(t *testing.T) {
	expected := "Hello, World!"
	received := Greet("World!")

	assert.Equal(t, expected, received)
}

type MockRepository struct {
	mock.Mock
}

func CreateConnection() {
	addr := ":8080"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
	}
	database, err := db.Initialize2()
	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}
	// defer database.Conn.Close()
	httpHandler := services.NewHandler(database)
	server := &http.Server{
		Handler: httpHandler,
	}
	go func() {
		server.Serve(listener)
	}()
	defer Stop(server)
}
func TestGetAllUsers(t *testing.T) {
	CreateConnection()
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(services.GetAllUsers)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	var X models.UserList
	err = json.Unmarshal([]byte(rr.Body.String()), &X)
	if err != nil {
		t.Errorf("something wrong")
	}
	if X.Users[0].Email != "hochilen" {
		t.Errorf("errors: %v ", X.Users[0].Email)
	}
	// Check the response body is what we expect.
	// expected := `[{"id":1,"first_name":"Krish","last_name":"Bhanushali","email_address":"krishsb@g.com","phone_number":"0987654321"},{"id":2,"first_name":"xyz","last_name":"pqr","email_address":"xyz@pqr.com","phone_number":"1234567890"},{"id":6,"first_name":"FirstNameSample","last_name":"LastNameSample","email_address":"lr@gmail.com","phone_number":"1111111111"}]`
	// expected :=`[{"users":[{"email":"hochilen"},{"email":"hochilen2"},{"email":"hochilen@gmail.com"}]}]`
	// expected := `{"users":[{"email":"hochilen"},{"email":"hochilen2"},{"email":"hochilen@gmail.com"}]}`
	// if rr.Body.String()!=expected{
	// 	t.Errorf("errors roi: %v",rr.Body.String())
	// }
	// assert.Equal(t, expected, rr.Body.String())
}
func TestGetUserByName(t *testing.T) {
	CreateConnection()
	req, err := http.NewRequest("GET", "/users/find", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("id", "hochilen@gmail.com")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(services.GetUser)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var X models.User
	err = json.Unmarshal([]byte(rr.Body.String()), &X)
	if err != nil {
		t.Errorf("something wrong")
	}
	if X.Email != "hochilen@gmail.com" {
		t.Errorf("errors: %v ", X.Email)
	}
}
func TestCreateUser(t *testing.T) {
	CreateConnection()
	var jsonStr = []byte(`{"email": "hochilen@gmail.com"}`)
	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(services.CreateUser)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	var X models.User
	err = json.Unmarshal([]byte(rr.Body.String()), &X)
	if err != nil {
		t.Errorf("something wrong")
	}
	if X.Email != "hochilen@gmail.com" {
		t.Errorf("errors: %v ", X.Email)
	}
}
func TestCreateUserWrongEmail(t *testing.T) {
	CreateConnection()
	var jsonStr = []byte(`{"email": "hochilen"}`)
	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(services.CreateUser)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v||%v",
			status, http.StatusOK,)
	}
}
func TestDeleteUserByName(t *testing.T) {
	CreateConnection()
	req, err := http.NewRequest("DELETE", "/users/delete", nil)
	if err != nil {
		t.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("id", "hochilen@gmail.com")
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(services.DeleteUser)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
func Stop(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Could not shut down server correctly: %v\n", err)
		os.Exit(1)
	}
}
