package controllers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"atlas-api/config/schema"
	"atlas-api/db"
	"atlas-api/helpers"

	"github.com/gorilla/mux"
)

// OrganizationReq ...
type OrganizationReq struct {
	ID           int
	TeamName     string           `json:"teamName"`
	ContactName  string           `json:"contactName"`
	ContactEmail string           `json:"contactEmail"`
	ContactPhone string           `json:"contactPhone"`
	Projects     []schema.Project `json:"projects"`
}

// GetOrganization will get the current organization
func GetOrganization(rw http.ResponseWriter, req *http.Request) {
	var (
		organization schema.Organization
		name         string
		client       string
		startDate    string
		junk         string
	)
	vars := mux.Vars(req)
	organizationID := vars["id"]
	oid, _ := strconv.Atoi(organizationID)

	database, err := db.Connection()
	if err != nil {
		helper.CreateResponse(rw, req, 500, nil, err)
	}

	rows, err := database.Query("SELECT * FROM organizations INNER JOIN projects on organizations.id=projects.organization_id WHERE organizations.id=$1", organizationID)
	if err != nil {
		helper.CreateResponse(rw, req, 500, nil, err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&junk, &organization.TeamName, &organization.ContactName, &organization.ContactEmail, &organization.ContactPhone, &junk, &name, &client, &startDate, &junk)
		if err != nil {
			log.Fatal(err)
		}
		project := schema.Project{name, client, "", startDate, oid, organization, nil, nil, nil, nil, nil}
		organization.Projects = append(organization.Projects, project)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	defer database.Close()

	helper.CreateResponse(rw, req, 200, organization, nil)
}

// CreateOrganization will create a new project
func CreateOrganization(rw http.ResponseWriter, req *http.Request) {
	var (
		organizationReq OrganizationReq
		organization    schema.Organization
	)

	body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))
	if err != nil {
		helper.CreateResponse(rw, req, 500, nil, err)
		return
	}
	if err := req.Body.Close(); err != nil {
		helper.CreateResponse(rw, req, 500, nil, err)
		return
	}

	if err := json.Unmarshal(body, &organizationReq); err != nil {
		helper.CreateResponse(rw, req, 500, nil, err)
		return
	}

	organization.TeamName = organizationReq.TeamName
	organization.ContactName = organizationReq.ContactName
	organization.ContactEmail = organizationReq.ContactEmail
	organization.ContactPhone = organizationReq.ContactPhone

	database, err := db.Connection()
	if err != nil {
		helper.CreateResponse(rw, req, 500, nil, err)
	}

	_, err = database.Query("INSERT INTO organizations (team_name, contact_name, contact_email, contact_phone) VALUES ($1, $2, $3, $4)",
		organization.TeamName,
		organization.ContactName,
		organization.ContactEmail,
		organization.ContactPhone,
	)
	if err != nil {
		helper.CreateResponse(rw, req, 500, nil, err)
		return
	}

	err = database.Close()
	if err != nil {
		helper.CreateResponse(rw, req, 500, nil, err)
	}

	helper.CreateResponse(rw, req, 200, organization, nil)
}
