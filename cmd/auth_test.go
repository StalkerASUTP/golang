package main

import (
	"bytes"
	"encoding/json"
	"go/adv-api/internal/auth"
	"go/adv-api/internal/user"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDb() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
func initData(db *gorm.DB) {
	db.Create(&user.User{
		Email:    "a2@a.ru",
		Password: "$2a$10$T2M4TH3A01pNs3XF0uAL5OsOPba/mMvZXTngQRRL0zj8u/p6AFUvK",
		Name:     "Вася",
	})
}
func removeData(db *gorm.DB) {
	db.Unscoped().
		Where("email = ?", "a2@a.ru").
		Delete(&user.User{})
}
func TestLoginSuccess(t *testing.T) {
	//prepare
	db := initDb()
	initData(db)
	defer removeData(db)
	ts := httptest.NewServer(App())
	defer ts.Close()
	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "a2@a.ru",
		Password: "1",
	})
	res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	var respData auth.LoginResponse
	err = json.Unmarshal(body, &respData)
	if err != nil {
		t.Fatal(err)
	}
	if respData.Token == "" {
		t.Fatal("Token empty")
	}
	if res.StatusCode != 200 {
		t.Fatalf("Expexted %d got %d", 200, res.StatusCode)
	}
	
}
func TestLoginFail(t *testing.T) {
	db := initDb()
	initData(db)
	defer removeData(db)
	ts := httptest.NewServer(App())
	defer ts.Close()
	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "a2@a.ru",
		Password: "2",
	})
	res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != 401 {
		t.Fatalf("Expexted %d got %d", 401, res.StatusCode)
	}
}
