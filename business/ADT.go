package business

import "time"

type CommonCriteria struct {
	Skills     []string `json:"skills"`
	Experience string   `json:"experience"`
	Location   string   `json:"location"`
}

type SearchCriteria struct {
	CommonCriteria        // Composición de los campos comunes
	Category       string `json:"category"`
}

type CallCriteria struct {
	CommonCriteria                        // Composición de los campos comunes
	Deadline       time.Time              `json:"deadline"`
	AdditionalInfo map[string]interface{} `json:"additional_info"`
}

type UserType interface {
	isUserType()
}

// Buscador representa un usuario que busca talentos
type Sponsor struct{}

func (Sponsor) isUserType() {}

// Talento representa un usuario que desea ser patrocinado, contratado, etc.
type Talent struct{}

func (Talent) isUserType() {}

// Representante representa un usuario que organiza convocatorias o representa talentos
type Representative struct{}

func (Representative) isUserType() {}

type Objective interface {
	isObjective()
}

type Hiring struct{}

func (Hiring) isObjective() {}

type Collaborating struct{}

func (Collaborating) isObjective() {}

type Sponsoring struct{}

func (Sponsoring) isObjective() {}

type Training struct{}

func (Training) isObjective() {}

type Mediating struct{}

func (Mediating) isObjective() {}

type ProposalResponse interface {
	isProposalResponse()
}

type Accept struct{}

func (Accept) isProposalResponse() {}

type Reject struct{}

func (Reject) isProposalResponse() {}
