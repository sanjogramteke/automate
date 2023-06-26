package s3backupchecktrigger

import (
	"net/http"

	"github.com/chef/automate/components/automate-cli/pkg/verifyserver/constants"
	"github.com/chef/automate/components/automate-cli/pkg/verifyserver/models"
	"github.com/chef/automate/components/automate-cli/pkg/verifyserver/services/batchcheckservice/trigger"
	"github.com/chef/automate/components/automate-cli/pkg/verifyserver/utils/checkutils"
	"github.com/chef/automate/lib/logger"
)

type S3BackupConfigCheck struct {
	log  logger.Logger
	port string
}

func NewS3BackupConfigCheck(log logger.Logger, port string) *S3BackupConfigCheck {
	return &S3BackupConfigCheck{
		log:  log,
		port: port,
	}
}

func (svc *S3BackupConfigCheck) Run(config *models.Config) []models.CheckTriggerResponse {
	if config.Hardware == nil || config.Backup.ObjectStorage == nil {
		return []models.CheckTriggerResponse{
			trigger.SkippedTriggerCheckResp(constants.UNKNOWN_HOST, constants.S3_BACKUP_CONFIG, constants.AUTOMATE),
		}
	}

	if isObjectStorage(config.Backup) {
		return emptyResp(config, constants.S3_BACKUP_CONFIG)
	}

	req := getS3CheckRequest(config.Backup.ObjectStorage)
	return runCheckForS3Config(config.Hardware.AutomateNodeIps, svc.log, svc.port, http.MethodPost, req)
}

func (ss *S3BackupConfigCheck) GetPortsForMockServer() map[string]map[string][]int {
	nodeTypePortMap := make(map[string]map[string][]int)
	return nodeTypePortMap
}

// runCheckForS3Config triggers the API on gives node automate nodes only for validating s3 backup config
func runCheckForS3Config(nodeIps []string, log logger.Logger, port string, method string, reqBody models.S3ConfigRequest) []models.CheckTriggerResponse {
	log.Debugf("Triggering the api call for automate nodes only for s3 backup config")
	outputCh := make(chan models.CheckTriggerResponse)
	for _, ip := range nodeIps {
		log.Debugf("Triggering api %s for the node %s", constants.S3_BACKUP_CHECK_API_PATH, ip)
		endpoint := checkutils.PrepareEndPoint(ip, port, constants.S3_BACKUP_CHECK_API_PATH)
		go trigger.TriggerCheckAPI(endpoint, ip, constants.AUTOMATE, method, outputCh, reqBody)
	}

	response := getResultFromOutputChan(len(nodeIps), outputCh)
	close(outputCh)
	return response
}

// getResultFromOutputChan gets the result from output channel
func getResultFromOutputChan(reqList int, outputCh chan models.CheckTriggerResponse) []models.CheckTriggerResponse {
	var result []models.CheckTriggerResponse

	for i := 0; i < reqList; i++ {
		res := <-outputCh
		result = append(result, res)
	}

	return result
}

func getS3CheckRequest(object *models.ObjectStorage) models.S3ConfigRequest {
	//passing a default value if s3 bucket path not provided
	if object.BasePath == "" {
		object.BasePath = "automate"
	}
	return models.S3ConfigRequest{
		Endpoint:   object.Endpoint,
		BucketName: object.BucketName,
		BasePath:   object.BasePath,
		AccessKey:  object.AccessKey,
		SecretKey:  object.SecretKey,
		Region:     object.AWSRegion,
	}
}

func isObjectStorage(backup *models.Backup) bool {
	return backup.ObjectStorage.BucketName == "" ||
		backup.ObjectStorage.AccessKey == "" ||
		backup.ObjectStorage.SecretKey == "" ||
		backup.ObjectStorage.AWSRegion == ""
}

func emptyResp(config *models.Config, checktype string) []models.CheckTriggerResponse {
	resps := []models.CheckTriggerResponse{}
	for _, ip := range config.Hardware.AutomateNodeIps {
		resps = append(resps, trigger.ErrTriggerCheckResp(ip, checktype, constants.AUTOMATE, constants.S3_BACKUP_MISSING))
	}

	return resps
}
