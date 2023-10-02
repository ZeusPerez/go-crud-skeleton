package transport

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ZeusPerez/go-crud-skeleton/internal/models"
	"github.com/ZeusPerez/go-crud-skeleton/internal/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	testDev = models.Dev{
		Email:     "dev@test.com",
		Languages: []string{"go"},
		Expertise: 2,
	}
	jsonDev = `{"email": "dev@test.com", "languages":["go"], "expertise": 2}`
)

func TestGetOK(t *testing.T) {
	mockDevs := &services.MockDevs{}
	devsHttpAdapter := NewHttpAdapter(HttpConfig{}, mockDevs)
	req := httptest.NewRequest(http.MethodGet, "/get?email=dev@test.com", nil)
	w := httptest.NewRecorder()

	mockDevs.On("Get", mock.Anything, "dev@test.com").
		Once().
		Return(testDev, nil)

	devsHttpAdapter.get(w, req)

	res := w.Result()
	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	var expectedDev models.Dev
	json.Unmarshal(responseBody, &expectedDev)
	assert.EqualValues(t, expectedDev, testDev)

	mockDevs.AssertExpectations(t)
}

func TestGetError(t *testing.T) {
	mockDevs := &services.MockDevs{}
	devsHttpAdapter := NewHttpAdapter(HttpConfig{}, mockDevs)
	req := httptest.NewRequest(http.MethodGet, "/get?email=dev@test.com", nil)
	w := httptest.NewRecorder()

	mockDevs.On("Get", mock.Anything, "dev@test.com").
		Once().
		Return(models.Dev{}, errors.New("test-error"))

	devsHttpAdapter.get(w, req)

	res := w.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)

	mockDevs.AssertExpectations(t)
}

func TestCreateOK(t *testing.T) {
	mockDevs := &services.MockDevs{}
	devsHttpAdapter := NewHttpAdapter(HttpConfig{}, mockDevs)
	req := httptest.NewRequest(http.MethodPost, "/create", bytes.NewBuffer([]byte(jsonDev)))
	w := httptest.NewRecorder()

	mockDevs.On("Create", mock.Anything, mock.AnythingOfType("models.Dev")).
		Once().
		Return(nil)

	devsHttpAdapter.create(w, req)

	res := w.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)

	mockDevs.AssertExpectations(t)
}

func TestCreateError(t *testing.T) {
	mockDevs := &services.MockDevs{}
	devsHttpAdapter := NewHttpAdapter(HttpConfig{}, mockDevs)
	req := httptest.NewRequest(http.MethodPost, "/create", bytes.NewBuffer([]byte(jsonDev)))
	w := httptest.NewRecorder()

	mockDevs.On("Create", mock.Anything, mock.AnythingOfType("models.Dev")).
		Once().
		Return(errors.New("test-error"))

	devsHttpAdapter.create(w, req)

	res := w.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)

	mockDevs.AssertExpectations(t)
}

func TestUpdateOK(t *testing.T) {
	mockDevs := &services.MockDevs{}
	devsHttpAdapter := NewHttpAdapter(HttpConfig{}, mockDevs)
	req := httptest.NewRequest(http.MethodPatch, "/update", bytes.NewBuffer([]byte(jsonDev)))
	w := httptest.NewRecorder()

	mockDevs.On("Update", mock.Anything, mock.AnythingOfType("models.Dev")).
		Once().
		Return(testDev, nil)

	devsHttpAdapter.update(w, req)

	res := w.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)

	mockDevs.AssertExpectations(t)
}

func TestUpdateError(t *testing.T) {
	mockDevs := &services.MockDevs{}
	devsHttpAdapter := NewHttpAdapter(HttpConfig{}, mockDevs)
	req := httptest.NewRequest(http.MethodPatch, "/update", bytes.NewBuffer([]byte(jsonDev)))
	w := httptest.NewRecorder()

	mockDevs.On("Update", mock.Anything, mock.AnythingOfType("models.Dev")).
		Once().
		Return(models.Dev{}, errors.New("test-error"))

	devsHttpAdapter.update(w, req)

	res := w.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)

	mockDevs.AssertExpectations(t)
}

func TestDeleteOK(t *testing.T) {
	mockDevs := &services.MockDevs{}
	devsHttpAdapter := NewHttpAdapter(HttpConfig{}, mockDevs)
	req := httptest.NewRequest(http.MethodDelete, "/delete?email=dev@test.com", nil)
	w := httptest.NewRecorder()

	mockDevs.On("Delete", mock.Anything, "dev@test.com").
		Once().
		Return(nil)

	devsHttpAdapter.delete(w, req)

	res := w.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)

	mockDevs.AssertExpectations(t)
}

func TestDeleteError(t *testing.T) {
	mockDevs := &services.MockDevs{}
	devsHttpAdapter := NewHttpAdapter(HttpConfig{}, mockDevs)
	req := httptest.NewRequest(http.MethodDelete, "/delete?email=dev@test.com", nil)
	w := httptest.NewRecorder()

	mockDevs.On("Delete", mock.Anything, "dev@test.com").
		Once().
		Return(errors.New("test-error"))

	devsHttpAdapter.delete(w, req)

	res := w.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)

	mockDevs.AssertExpectations(t)
}

func TestWrongMethod(t *testing.T) {
	mockDevs := &services.MockDevs{}
	devsHttpAdapter := NewHttpAdapter(HttpConfig{}, mockDevs)
	req := httptest.NewRequest(http.MethodDelete, "/get?email=dev@test.com", nil)
	w := httptest.NewRecorder()

	devsHttpAdapter.get(w, req)

	res := w.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusMethodNotAllowed, res.StatusCode)

	mockDevs.AssertExpectations(t)
}
