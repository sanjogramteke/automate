package trigger

import "github.com/chef/automate/components/automate-cli/pkg/verifyserver/models"

type SshUserAccessCheck struct {
}

func NewSshUserAccessCheck() *SshUserAccessCheck {
	return &SshUserAccessCheck{}
}

func (suc *SshUserAccessCheck) Run(config models.Config) []models.CheckTriggerResponse {
	m := []models.CheckTriggerResponse{
		models.CheckTriggerResponse{},
	}
	return m
}