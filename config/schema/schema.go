package schema

import ()

// User schema
type User struct {
	FirstName    string
	LastName     string
	Email        string
	PasswordHash string
	PasswordSalt string
	Disabled     bool
}

// Organization schema
type Organization struct {
	TeamName     string
	ContactName  string
	ContactEmail string
	ContactPhone string
	Projects     []Project
}

// Project schema
type Project struct {
	Name           string
	Client         string
	SlackChannel   string
	StartDate      string
	OrganizationID int
	Organization   Organization
	Platforms      []Platform
	Pages          []Page
	Tasks          []Task
	Roles          []Role
	Groups         []Group
}

// Platform schema
type Platform struct {
	Name string
}

// Page schema
type Page struct {
	Name string
}

// Task schema
type Task struct {
	Name string
}

// Role schema
type Role struct {
	Name string
}

// Group schema
type Group struct {
	Name string
}
