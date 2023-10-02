package acceptance

import (
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

var dev = `{"email":"dev@test.com","languages":["go"],"expertise":2}`

type DevCrudSuite struct {
	suite.Suite
}

func (d *DevCrudSuite) SetupTest() {
	url := "http://devs-crud:8000/create"
	makeRequest(d.T(), http.MethodPost, url, strings.NewReader(dev))
}

func (d *DevCrudSuite) TearDownTest() {
	url := "http://devs-crud:8000/delete?email=dev@test.com"
	makeRequest(d.T(), http.MethodDelete, url, nil)
}

func (d *DevCrudSuite) Test_GetDev() {
	url := "http://devs-crud:8000/get?email=dev@test.com"
	statusCode, body := makeRequest(d.T(), http.MethodGet, url, nil)
	d.Require().Equal(http.StatusOK, statusCode)

	trimmedBody := strings.TrimSuffix(body, "\n")
	d.Require().Equal(dev, trimmedBody)
}

func (d *DevCrudSuite) Test_GetEmptyDev() {
	url := "http://devs-crud:8000/get?email=no-dev@test.com"
	statusCode, body := makeRequest(d.T(), http.MethodGet, url, nil)
	d.Require().Equal(http.StatusNotFound, statusCode)

	trimmedBody := strings.TrimSuffix(body, "\n")
	d.Require().Equal("Resource not found", trimmedBody)
}

func (d *DevCrudSuite) Test_CreateDev() {
	var newDev = `{"email":"anotherdev@test.com","languages":["go"],"expertise":2}`
	url := "http://devs-crud:8000/create"
	statusCode, _ := makeRequest(d.T(), http.MethodPost, url, strings.NewReader(newDev))
	d.Require().Equal(http.StatusOK, statusCode)

	url = "http://devs-crud:8000/get?email=anotherdev@test.com"
	statusCode, body := makeRequest(d.T(), http.MethodGet, url, nil)
	d.Require().Equal(http.StatusOK, statusCode)

	trimmedBody := strings.TrimSuffix(body, "\n")
	d.Require().Equal(newDev, trimmedBody)

	url = "http://devs-crud:8000/delete?email=anotherdev@test.com"
	statusCode, _ = makeRequest(d.T(), http.MethodDelete, url, nil)
	d.Require().Equal(http.StatusOK, statusCode)
}

func (d *DevCrudSuite) Test_UpdateDev() {
	var updatedDev = `{"email":"dev@test.com","languages":["go, ruby"],"expertise":4}`
	url := "http://devs-crud:8000/update?email=dev@test.com"
	statusCode, _ := makeRequest(d.T(), http.MethodPatch, url, strings.NewReader(updatedDev))
	d.Require().Equal(http.StatusOK, statusCode)

	url = "http://devs-crud:8000/get?email=dev@test.com"
	statusCode, body := makeRequest(d.T(), http.MethodGet, url, nil)
	d.Require().Equal(http.StatusOK, statusCode)

	trimmedBody := strings.TrimSuffix(body, "\n")
	d.Require().Equal(updatedDev, trimmedBody)
}

func (d *DevCrudSuite) Test_DeleteDev() {
	url := "http://devs-crud:8000/delete?email=dev@test.com"
	statusCode, _ := makeRequest(d.T(), http.MethodDelete, url, nil)
	d.Require().Equal(http.StatusOK, statusCode)

	url = "http://devs-crud:8000/get?email=dev@test.com"
	statusCode, body := makeRequest(d.T(), http.MethodGet, url, nil)
	d.Require().Equal(http.StatusNotFound, statusCode)

	trimmedBody := strings.TrimSuffix(body, "\n")
	d.Require().Equal("Resource not found", trimmedBody)
}

func (d *DevCrudSuite) TestMalformedBody() {
	var malformedBody = `this is not a JSON`
	url := "http://devs-crud:8000/create"
	statusCode, body := makeRequest(d.T(), http.MethodPost, url, strings.NewReader(malformedBody))
	d.Require().Equal(http.StatusBadRequest, statusCode)

	trimmedBody := strings.TrimSuffix(body, "\n")
	d.Require().Contains(trimmedBody, "error parsing JSON input")

}

func (d *DevCrudSuite) TestinvalidArguments() {
	var newDev = `{"email":"anotherdev@test.com","languages":["go"],"expertise":-2}`
	url := "http://devs-crud:8000/create"
	statusCode, body := makeRequest(d.T(), http.MethodPost, url, strings.NewReader(newDev))
	d.Require().Equal(http.StatusBadRequest, statusCode)

	trimmedBody := strings.TrimSuffix(body, "\n")
	d.Require().Contains(trimmedBody, "error validating struct: Key: 'Dev.Expertise'")
}

func (d *DevCrudSuite) TestinvalidEmail() {
	var newDev = `{"email":"not+valid+email.com","languages":["go"],"expertise":2}`
	url := "http://devs-crud:8000/create"
	statusCode, body := makeRequest(d.T(), http.MethodPost, url, strings.NewReader(newDev))
	d.Require().Equal(http.StatusBadRequest, statusCode)

	trimmedBody := strings.TrimSuffix(body, "\n")
	d.Require().Contains(trimmedBody, "error validating struct: Key: 'Dev.Email'")
}

// TODO: think more test cases
// func anotherTestcase() {}

func TestDevCrudSuite(t *testing.T) {
	suite.Run(t, new(DevCrudSuite))
}
