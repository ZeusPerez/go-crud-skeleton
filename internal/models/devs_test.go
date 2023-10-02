package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOK(t *testing.T) {
	jsonDev := `{"email": "dev@test.com", "languages":["go"], "expertise": 2}`

	_, err := JsonToDev([]byte(jsonDev))
	assert.NoError(t, err)
}

func TestEmailRequired(t *testing.T) {
	errStr := "error validating struct: Key: 'Dev.Email' Error:Field validation for 'Email' failed on the 'required' tag"

	jsonDev := `{"languages":["go"], "expertise": 2}`

	_, err := JsonToDev([]byte(jsonDev))
	assert.Error(t, err)
	assert.Equal(t, err.Error(), errStr)
}

func TestEmailFormat(t *testing.T) {
	errStr := "error validating struct: Key: 'Dev.Email' Error:Field validation for 'Email' failed on the 'email' tag"

	jsonDev := `{"email": "no-email-format", "languages":["go"], "expertise": 2}`

	_, err := JsonToDev([]byte(jsonDev))
	assert.Error(t, err)
	assert.Equal(t, err.Error(), errStr)
}

func TestExpertiseRequired(t *testing.T) {
	errStr := "error validating struct: Key: 'Dev.Expertise' Error:Field validation for 'Expertise' failed on the 'required' tag"

	jsonDev := `{"email": "dev@test.com", "languages":["go"]}`

	_, err := JsonToDev([]byte(jsonDev))
	assert.Error(t, err)
	assert.Equal(t, err.Error(), errStr)
}

func TestExpertiseMaxValue(t *testing.T) {
	errStr := "error validating struct: Key: 'Dev.Expertise' Error:Field validation for 'Expertise' failed on the 'max' tag"

	jsonDev := `{"email": "dev@test.com", "languages":["go"], "expertise": 20}`

	_, err := JsonToDev([]byte(jsonDev))
	assert.Error(t, err)
	assert.Equal(t, err.Error(), errStr)
}

func TestExpertiseMinValue(t *testing.T) {
	errStr := "error validating struct: Key: 'Dev.Expertise' Error:Field validation for 'Expertise' failed on the 'min' tag"

	jsonDev := `{"email": "dev@test.com", "languages":["go"], "expertise": -2}`

	_, err := JsonToDev([]byte(jsonDev))
	assert.Error(t, err)
	assert.Equal(t, err.Error(), errStr)
}
