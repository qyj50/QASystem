package qanda

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

//
// Auxiliary functions for Auxiliary functions
//

func GetEchoTestEnv(filename string) *echo.Echo {
	return New("/var/empty", "sqlite3", "file:"+filename+"?mode=memory&cache=shared&_fk=1", "super-secret-key")
}

func GetTokenFromBody(rec *httptest.ResponseRecorder, t *testing.T) string {
	resp := new(struct {
		Token string `json:"token"`
	})
	err := json.NewDecoder(rec.Body).Decode(resp)
	if t != nil && err != nil {
		t.Fatal(err)
	}
	return resp.Token
}

//
// Auxiliary functions
//

func AuxUserRegister(e *echo.Echo, t *testing.T, name string, password string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodPost, "/v1/user/register", bytes.NewBufferString(`
{
	"username": "`+name+`",
	"password": "`+password+`"
}
    `))
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if t != nil && rec.Result().StatusCode != http.StatusOK {
		t.Fatal("register failed")
	}

	return rec
}

func AuxUserLogin(e *echo.Echo, t *testing.T, name string, password string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodPost, "/v1/user/login", bytes.NewBufferString(`
{
	"username": "`+name+`",
	"password": "`+password+`"
}
    `))
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if t != nil && rec.Result().StatusCode != http.StatusOK {
		t.Fatal("login failed")
	}

	return rec
}

func AuxUserInfo(e *echo.Echo, t *testing.T, token string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodGet, "/v1/user/info", nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if t != nil && rec.Result().StatusCode != http.StatusOK {
		t.Fatal("Get user info failed")
	}

	return rec
}

func AuxUserEdit(e *echo.Echo, t *testing.T, token string, body string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodPost, "/v1/user/edit", bytes.NewBufferString(body))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if t != nil && rec.Result().StatusCode != http.StatusOK {
		t.Fatal("edit user failed")
	}

	return rec
}

func AuxUserFilter(e *echo.Echo, t *testing.T, query string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodGet, "/v1/user/filter"+query, nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if t != nil && rec.Result().StatusCode != http.StatusOK {
		t.Fatal("filter user failed")
	}

	return rec
}

func AuxQuestionSubmit(e *echo.Echo, t *testing.T, body string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodPost, "/v1/question/submit", bytes.NewBufferString(body))
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if t != nil && rec.Result().StatusCode != http.StatusOK {
		t.Fatal("submit question failed")
	}

	return rec
}

func AuxQuestionPay(e *echo.Echo, t *testing.T, questionid int, token string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodPost, "/v1/question/pay", bytes.NewBufferString(`
{
	"questionid": `+strconv.Itoa(questionid)+`
}
    `))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if t != nil && rec.Result().StatusCode != http.StatusOK {
		t.Fatal("question pay failed")
	}

	return rec
}

func AuxQuestionQuery(e *echo.Echo, t *testing.T, questionid int) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodGet, "/v1/question/"+strconv.Itoa(questionid), nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if t != nil && rec.Result().StatusCode != http.StatusOK {
		t.Fatal("question query failed")
	}

	return rec
}

func AuxQuestionList(e *echo.Echo, t *testing.T) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodGet, "/v1/question/list", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if t != nil && rec.Result().StatusCode != http.StatusOK {
		t.Fatal("question list failed")
	}

	return rec
}

func AuxQuestionMine(e *echo.Echo, t *testing.T, token string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodGet, "/v1/question/mine", nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if t != nil && rec.Result().StatusCode != http.StatusOK {
		t.Fatal("question mine failed")
	}

	return rec
}

func AuxQuestionAccept(e *echo.Echo, t *testing.T, questionid int, choice bool, token string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodPost, "/v1/question/accept", bytes.NewBufferString(`
{
	"questionid": `+strconv.Itoa(questionid)+`,
	"choice": `+fmt.Sprintf(`%t`, choice)+`
}
    `))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if t != nil && rec.Result().StatusCode != http.StatusOK {
		t.Fatal("question accept failed")
	}

	return rec
}

func AuxQuestionClose(e *echo.Echo, t *testing.T, questionid int, token string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodPost, "/v1/question/close", bytes.NewBufferString(`
{
	"questionid": `+strconv.Itoa(questionid)+`
}
    `))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if t != nil && rec.Result().StatusCode != http.StatusOK {
		t.Fatal("question close failed")
	}

	return rec
}

//
// test functions
//

func TestUser(t *testing.T) {
	e := GetEchoTestEnv("ent")
	AuxUserRegister(e, t, "user1", "testpassword")
	rec := AuxUserLogin(e, t, "user1", "testpassword")
	token := GetTokenFromBody(rec, t)
	AuxUserInfo(e, t, token)
	AuxUserEdit(e, t, token, `
{
	"email":"hello@example",
	"phone":"12345678",
	"answerer":true,
	"price":-100,
	"profession":"Geschichte"
}
	`)
	AuxUserFilter(e, t, "?username=user1&email=&phone=&answerer=&price=-100&profession=")
}

// Register: bad json
func TestUserRegisterX1(t *testing.T) {
	e := GetEchoTestEnv("ent")
	if rec := AuxUserRegister(e, nil, "userX\"!!", "testpassword"); rec.Result().StatusCode != http.StatusBadRequest {
		t.Fatal("user register allows bad json")
	}
}

// Register: no password
func TestUserRegisterX2(t *testing.T) {
	e := GetEchoTestEnv("ent")
	if rec := AuxUserRegister(e, nil, "userX\"}", ""); rec.Result().StatusCode != http.StatusBadRequest {
		t.Fatal("user register allows no password")
	}
}

// Register: repeated register
func TestUserRegisterX3(t *testing.T) {
	e := GetEchoTestEnv("ent")
	if rec := AuxUserRegister(e, nil, "user1", "testpassword"); rec.Result().StatusCode != http.StatusBadRequest {
		t.Fatal("user register allows repeated register")
	}
}

// Login: bad json
func TestUserLoginX1(t *testing.T) {
	e := GetEchoTestEnv("ent")
	if rec := AuxUserLogin(e, nil, "user1\"!!", "testpassword"); rec.Result().StatusCode != http.StatusBadRequest {
		t.Fatal("user login allows bad json")
	}
}

// Login: no password
func TestUserLoginX2(t *testing.T) {
	e := GetEchoTestEnv("ent")
	if rec := AuxUserLogin(e, nil, "user1\"}", "testpassword"); rec.Result().StatusCode != http.StatusBadRequest {
		t.Fatal("user login allows no password")
	}
}

// Login: incorrect password
func TestUserLoginX3(t *testing.T) {
	e := GetEchoTestEnv("ent")
	if rec := AuxUserLogin(e, nil, "user1", "hello"); rec.Result().StatusCode != http.StatusBadRequest {
		t.Fatal("user login allows incorrect password")
	}
}

// Info: token verification
func TestUserInfoX1(t *testing.T) {
	e := GetEchoTestEnv("ent")
	token := GetTokenFromBody(AuxUserLogin(e, t, "user1", "testpassword"), t)
	if rec := AuxUserInfo(e, nil, token+"qwerty"); rec.Result().StatusCode != http.StatusForbidden {
		t.Fatal("user info allows incorrect token")
	}
}

// Info: inexistent user
func TestUserInfoX2(t *testing.T) {
	e := GetEchoTestEnv("ent")
	e1 := GetEchoTestEnv("entTestUserInfoX2")
	token := GetTokenFromBody(AuxUserRegister(e1, t, "user1X", "testpassword"), t)
	if rec := AuxUserInfo(e, nil, token); rec.Result().StatusCode != http.StatusBadRequest {
		t.Fatal("user info allows inexistent user")
	}
}

// Edit: bad json
func TestUserEditX1(t *testing.T) {
	e := GetEchoTestEnv("ent")
	token := GetTokenFromBody(AuxUserLogin(e, t, "user1", "testpassword"), t)
	if rec := AuxUserEdit(e, nil, token, "{\"}"); rec.Result().StatusCode != http.StatusBadRequest {
		t.Fatal("user edit allows bad json")
	}
}

// Filter: bad query
func TestUserFilterX1(t *testing.T) {
	e := GetEchoTestEnv("ent")
	if rec := AuxUserFilter(e, nil, "?answerer=trualse"); rec.Result().StatusCode != http.StatusBadRequest {
		t.Fatal("user filter allows bad query")
	}
}

/*
func TestQuestion(t *testing.T) {
	e := GetEchoTestEnv("ent1")

	rec := AuxUserRegister(e, t, "user1", "pass")
	token1 := GetTokenFromBody(rec, t)
	rec = AuxUserRegister(e, t, "user2", "pass")
	token2 := GetTokenFromBody(rec, t)

	AuxQuestionSubmit(e, t, `
{
	"title": "test title",
	"content":"test content",
	"questionerid":1,
	"answererid":2
}
	`)

	AuxQuestionPay(e, t, 1, token1)
	AuxQuestionQuery(e, t, 1)
	AuxQuestionList(e, t)
	AuxQuestionMine(e, t, token1)

	AuxQuestionAccept(e, t, 1, true, token2)
	AuxQuestionClose(e, t, 1, token1)
}
*/