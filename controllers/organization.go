package controllers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	// "atlas-api/config/db"
	"atlas-api/config/schema"
	"atlas-api/helpers"
)

// OrganizationReq ...
type OrganizationReq struct {
	TeamName     string `json:"teamName"`
	ContactName  string `json:"contactName"`
	ContactEmail string `json:"contactEmail"`
	ContactPhone string `json:"contactPhone"`
}

// CreateOrganization will create a new project
func CreateOrganization(rw http.ResponseWriter, req *http.Request) {
	var organizationReq OrganizationReq
	var organization schema.Organization

	body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))
	if err != nil {
		log.Fatal(err)
	}
	if err := req.Body.Close(); err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(body, &organizationReq); err != nil {
		helper.JSONHandler(rw, req)

		rw.WriteHeader(422)
		err = json.NewEncoder(rw).Encode(err)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	organization.TeamName = organizationReq.TeamName
	organization.ContactName = organizationReq.ContactName
	organization.ContactEmail = organizationReq.ContactEmail
	organization.ContactPhone = organizationReq.ContactPhone

	// if err := db.DB.Create(&organization).Error; err != nil {
	//
	// 	err = helper.HandleError(rw, req, 400, err)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	//
	// 	return
	// }
	//
	helper.HandleError(rw, req, 200, nil)
}
