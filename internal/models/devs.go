package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
)

var validate *validator.Validate

type Languages []string

func (s Languages) Value() (driver.Value, error) {
	if len(s) == 0 {
		return "[]", nil
	}
	return fmt.Sprintf(`["%s"]`, strings.Join(s, `","`)), nil
}

func (s *Languages) Scan(src interface{}) (err error) {
	var languages []string
	switch src := src.(type) {
	case string:
		err = json.Unmarshal([]byte(src), &languages)
	case []byte:
		err = json.Unmarshal(src, &languages)
	default:
		return errors.New("incompatible type for Languages")
	}
	if err != nil {
		return
	}
	*s = languages
	return nil
}

type Dev struct {
	Email     string    `validate:"required,email" json:"email" sql:"email"`
	Languages Languages `json:"languages" sql:"languages"`
	Expertise int       `validate:"required,min=0,max=5" json:"expertise" sql:"expertise"`
}

func JsonToDev(input []byte) (Dev, error) {
	var dev Dev
	err := json.Unmarshal(input, &dev)
	if err != nil {
		log.Errorf("error parsing JSON input: %v", err.Error())
		return dev, fmt.Errorf("error parsing JSON input: %v", err.Error())
	}
	validate = validator.New()
	err = validate.Struct(dev)
	if err != nil {
		log.Errorf("error validating struct: %v", err.Error())
		return dev, fmt.Errorf("error validating struct: %v", err.Error())
	}

	return dev, nil
}
