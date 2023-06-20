package sshuseraccesschecktrigger

import (
	"fmt"
	"net/http"

	"github.com/chef/automate/components/automate-cli/pkg/verifyserver/constants"
	"github.com/chef/automate/components/automate-cli/pkg/verifyserver/models"
	"github.com/chef/automate/components/automate-cli/pkg/verifyserver/services/batchcheckservice/trigger"
	"github.com/chef/automate/components/automate-cli/pkg/verifyserver/utils/configutils"
	"github.com/chef/automate/lib/logger"
)

type SshUserAccessCheck struct {
	host string
	port string
	log  logger.Logger
}

func NewSshUserAccessCheck(log logger.Logger, port string) *SshUserAccessCheck {
	return &SshUserAccessCheck{
		log:  log,
		host: constants.LOCAL_HOST_URL,
		port: port,
	}
}

func (ss *SshUserAccessCheck) Run(config *models.Config) []models.CheckTriggerResponse {
	ss.log.Info("Performing SSH user access check from batch check ")

	// Check if certificate is empty or nil
	if config.Hardware == nil {
		return trigger.HardwareNil(constants.SSH_USER, true, true, false)
	}
	if config.SSHUser == nil {
		return trigger.ConstructNilResp(config, constants.SSH_USER)
	}
	if IsSSHUserEmpty(config.SSHUser) {
		return trigger.ConstructEmptyResp(config, constants.SSH_USER, "SSH credentials is missing")
	}

	count := config.Hardware.AutomateNodeCount + config.Hardware.ChefInfraServerNodeCount +
		config.Hardware.PostgresqlNodeCount + config.Hardware.OpenSearchNodeCount

	outputCh := make(chan models.CheckTriggerResponse, count)
	url := fmt.Sprintf("%s:%s%s", ss.host, ss.port, constants.SSH_USER_CHECK_API_PATH)

	var finalResult []models.CheckTriggerResponse
	hostMap := configutils.GetNodeTypeMap(config.Hardware)
	for ip, types := range hostMap {
		for i := 0; i < len(types); i++ {
			requestBody := getSShUserAPIRquest(ip, config.SSHUser)
			go trigger.TriggerCheckAPI(url, ip, types[i], http.MethodPost, outputCh, requestBody)
		}
	}
	for i := 0; i < count; i++ {
		resp := <-outputCh
		finalResult = append(finalResult, resp)
	}
	close(outputCh)

	return finalResult
}

func getSShUserAPIRquest(ip string, sshUser *models.SSHUser) models.SShUserRequest {

	return models.SShUserRequest{
		IP:           ip,
		Username:     sshUser.Username,
		Port:         sshUser.Port,
		SudoPassword: sshUser.SudoPassword,
		PrivateKey:   sshUser.PrivateKey,
	}
}

func (ss *SshUserAccessCheck) GetPortsForMockServer() map[string]map[string][]int {
	nodeTypePortMap := make(map[string]map[string][]int)
	return nodeTypePortMap
}

func IsSSHUserEmpty(sshUser *models.SSHUser) bool {
	return (sshUser.Username == "" || sshUser.Port == "" || sshUser.PrivateKey == "")
}
