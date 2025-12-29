package auth_e2e

import (
	"auth/internal/app"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	testApp    http.Handler
	testServer *httptest.Server
)

func TestMain(m *testing.M) {
	setup()
	// Run tests
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	fmt.Println("Setup")
	app, _ := app.GetApp("../../../../../config.json")
	testApp = app
	ts := httptest.NewServer(testApp)
	testServer = ts
}

func teardown() {
	fmt.Println("Teardown")
	defer testServer.Close()
}
