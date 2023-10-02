package transport

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ZeusPerez/go-crud-skeleton/internal/errors"
	"github.com/ZeusPerez/go-crud-skeleton/internal/models"
	"github.com/ZeusPerez/go-crud-skeleton/internal/services"
	log "github.com/sirupsen/logrus"
)

type HttpConfig struct {
	Timeout time.Duration `default:"1s"`
}

type HttpDevAdapter interface {
	AddHandlers(mux *http.ServeMux)
}

type httpDevAdapter struct {
	cfg      HttpConfig
	provider services.Devs
}

func NewHttpAdapter(cfg HttpConfig, devsProvider services.Devs) httpDevAdapter {
	return httpDevAdapter{
		cfg:      cfg,
		provider: devsProvider}
}

func (a httpDevAdapter) AddHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/status", a.status)
	mux.HandleFunc("/get", a.get)
	mux.HandleFunc("/create", a.create)
	mux.HandleFunc("/update", a.update)
	mux.HandleFunc("/delete", a.delete)
}

func (a httpDevAdapter) status(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	statusMsg := `{"status":"ok"}`
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(statusMsg))

}

func (a httpDevAdapter) get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), a.cfg.Timeout)
	defer cancel()

	email := r.URL.Query().Get("email")
	dev, err := a.provider.Get(ctx, email)
	if err != nil {
		parseServiceErr(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dev)
}

func (a httpDevAdapter) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), a.cfg.Timeout)
	defer cancel()

	dev, err := parseBody(w, r.Body)
	if err != nil {
		return
	}

	err = a.provider.Create(ctx, dev)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

}

func (a httpDevAdapter) update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), a.cfg.Timeout)
	defer cancel()

	dev, err := parseBody(w, r.Body)
	if err != nil {
		return
	}

	dev, err = a.provider.Update(ctx, dev)
	if err != nil {
		parseServiceErr(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dev)
}

func (a httpDevAdapter) delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), a.cfg.Timeout)
	defer cancel()

	email := r.URL.Query().Get("email")
	err := a.provider.Delete(ctx, email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

}

func parseBody(w http.ResponseWriter, body io.ReadCloser) (models.Dev, error) {
	bytesBody, err := ioutil.ReadAll(body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return models.Dev{}, err
	}

	dev, err := models.JsonToDev(bytesBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	return dev, err
}

func parseServiceErr(w http.ResponseWriter, err error) {
	log.Error(err)
	switch err.(type) {
	case errors.Internal:
		w.WriteHeader(http.StatusInternalServerError)
	case errors.NotFound:
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusBadRequest)

	}

	w.Write([]byte(err.Error()))
}
