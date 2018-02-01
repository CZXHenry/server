package controllers

import (
	"html/template"

	"github.com/sunyatsuntobee/server/models"
)

type layout struct {
	Title   string
	Content template.HTML
}

type contact struct {
	Name  string
	Phone string
}

// PhotoLiveDetail is a complete information of an
// Entity PhotoLive
type PhotoLiveDetail struct {
	PhotoLive           *models.PhotoLive
	Organization        *models.Organization
	Activity            *models.Activity
	ActivityStage       *models.ActivityStage
	Manager             *models.User
	PhotographerManager *models.User
	Supervisors         []*models.User
}

// OrganizationDetail is a complete information of an
// Entity Organization
type OrganizationDetail struct {
	Organization *models.Organization
	Contactors   []*models.User
	Departments  []*models.OrganizationDepartment
	LoginLogs    []*models.OrganizationLoginLog
}
