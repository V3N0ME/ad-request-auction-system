package usecase

import (
	template "janusapi/pkg/template"
)

//Usecase is the instance of template
type Usecase struct {
	repo template.Repository
}

//Create creates a new template
func (uc *Usecase) Create() {

}

//Get returns a template
func (uc *Usecase) Get() {

}

//New returns a new instance of usecase
func New(repo template.Repository) *Usecase {
	return &Usecase{repo}
}
