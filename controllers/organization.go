package controllers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"atlas-api/config/schema"
	"atlas-api/db"
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
	var (
		organizationReq OrganizationReq
		organization    schema.Organization
	)

	body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))
	if err != nil {
		helper.OrganizationResponse(rw, req, 200, organization, nil)
		return
	}
	if err := req.Body.Close(); err != nil {
		helper.OrganizationResponse(rw, req, 200, organization, nil)
		return
	}

	if err := json.Unmarshal(body, &organizationReq); err != nil {
		helper.OrganizationResponse(rw, req, 200, organization, nil)
		return
	}

	organization.TeamName = organizationReq.TeamName
	organization.ContactName = organizationReq.ContactName
	organization.ContactEmail = organizationReq.ContactEmail
	organization.ContactPhone = organizationReq.ContactPhone

	database, err := db.Connection()
	if err != nil {
		log.Fatal(err)
	}

	_, err = database.Query("INSERT INTO organizations (team_name, contact_name, contact_email, contact_phone) VALUES ($1, $2, $3, $4)",
		organization.TeamName,
		organization.ContactName,
		organization.ContactEmail,
		organization.ContactPhone,
	)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = database.Close()
	if err != nil {
		log.Fatal(err)
	}

	helper.OrganizationResponse(rw, req, 200, organization, nil)
}
