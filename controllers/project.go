package controllers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"atlas-api/config/schema"
	"atlas-api/db"
	"atlas-api/helpers"
)

// ProjectReq ...
type ProjectReq struct {
	Name           string `json:"name"`
	Client         string `json:"client"`
	SlackChannel   string `json:"slackChannel"`
	StartDate      string `json:"startDate"`
	OrganizationID int    `json:"organizationId"`
}

// CreateProject will create a new project
func CreateProject(rw http.ResponseWriter, req *http.Request) {
	var projectReq ProjectReq
	var project schema.Project

	body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))
	if err != nil {
		helper.CreateResponse(rw, req, 500, nil, err)
		return
	}
	if err := req.Body.Close(); err != nil {
		helper.CreateResponse(rw, req, 500, nil, err)
		return
	}

	if err := json.Unmarshal(body, &projectReq); err != nil {
		helper.CreateResponse(rw, req, 500, nil, err)
		return
	}

	project.Name = projectReq.Name
	project.Client = projectReq.Client
	project.StartDate = projectReq.StartDate
	project.OrganizationID = projectReq.OrganizationID

	database, err := db.Connection()
	if err != nil {
		helper.CreateResponse(rw, req, 500, nil, err)
		return
	}

	// TODO: slack channel in database
	_, err = database.Query("INSERT INTO projects (name, client, start_date, organization_id) VALUES ($1, $2, $3, $4)",
		project.Name,
		project.Client,
		project.StartDate,
		project.OrganizationID,
	)
	if err != nil {
		helper.CreateResponse(rw, req, 500, nil, err)
		return
	}

	err = database.Close()
	if err != nil {
		helper.CreateResponse(rw, req, 500, nil, err)
		return
	}

	helper.CreateResponse(rw, req, 200, project, err)
}
