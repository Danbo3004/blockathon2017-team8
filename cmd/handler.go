package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"../model"
	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

var (
	companies      = make(map[string]model.Company)
	users          = make(map[string]model.User)
	usersCompanies map[string]model.UserCompany
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	defer closeRequest(r)

	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	user.ID = uuid.NewV4().String()
	users[user.ID] = user

	responseSuccess(w, http.StatusOK, map[string]string{"id": user.ID})
}

func BuyToken(w http.ResponseWriter, r *http.Request) {
	defer closeRequest(r)

	var userCompany model.UserCompany
	err := json.NewDecoder(r.Body).Decode(&userCompany)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	userCompany.ID = uuid.NewV4().String()
	usersCompanies[userCompany.ID] = userCompany

	// execute smart contract

	responseSuccess(w, http.StatusOK, map[string]bool{"success": true})
}

func RegisterCompany(w http.ResponseWriter, r *http.Request) {
	defer closeRequest(r)
	var company model.Company

	err := json.NewDecoder(r.Body).Decode(&company)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	company.ID = uuid.NewV4().String()
	companies[company.ID] = company

	// deploy a smart contract

	responseSuccess(w, http.StatusOK, map[string]string{"id": company.ID})
}

func GetCompany(w http.ResponseWriter, r *http.Request) {
	defer closeRequest(r)

	companyID := pathParam(r, "id")

	company, ok := companies[companyID]
	if !ok {
		responseError(w, http.StatusBadRequest, fmt.Sprintf("there is no company with id %s", companyID))
		return
	}

	responseSuccess(w, http.StatusOK, company)
}

func responseError(w http.ResponseWriter, statusCode int, msg string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(statusCode)
	if _, err := io.WriteString(w, msg); err != nil {
		logrus.Error("Write response message", err)
	}
}

func responseSuccess(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Context-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			logrus.Error("Marshal body object", err)
			responseError(w, http.StatusInternalServerError, err.Error())
			return
		}
		if _, err = w.Write(data); err != nil {
			logrus.Error("Write response body content", err)
		}
	}
}

// pathParam returns path parameter by name.
func pathParam(r *http.Request, name string) string {
	vars := mux.Vars(r)
	return vars[name]
}

// closeRequest closes request body.
func closeRequest(r *http.Request) {
	err := r.Body.Close()
	if err != nil {
		logrus.Error(err)
	}
}
