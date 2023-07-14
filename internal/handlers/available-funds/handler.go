package handlerAvailableFunds

import (
	"encoding/json"
	"net/http"

	"github.com/codebyjdd/tech-test-cushon/internal/db"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

type Handler interface {
	Handle(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
}

type handler struct {
	db  db.Db
	log *logrus.Logger
}

func (h handler) Handle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	funds, err := h.db.GetAvailablePublicFunds()
	h.log.Infof("%s : %d", InfoGotPubFunds, len(funds))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err = w.Write([]byte("could not get available funds")); err != nil {
			h.log.Errorf("%s : %s", ErrResp, err.Error())
		}
		return
	}

	b, err := json.Marshal(funds)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err = w.Write([]byte("could not marshal funds")); err != nil {
			h.log.Errorf("%s : %s", ErrResp, err.Error())
		}
	}
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(b); err != nil {
		h.log.Errorf("%s : %s", ErrResp, err.Error())
	}
}
