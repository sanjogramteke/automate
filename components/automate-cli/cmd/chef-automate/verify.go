package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/chef/automate/components/automate-cli/pkg/docs"
	"github.com/chef/automate/components/automate-cli/pkg/verifyserver/constants"
	"github.com/chef/automate/components/automate-cli/pkg/verifyserver/models"
	"github.com/chef/automate/components/automate-cli/pkg/verifyserver/response"
	"github.com/chef/automate/components/automate-cli/pkg/verifyserver/server"
	"github.com/chef/automate/components/automate-cli/pkg/verifysystemdcreate"
	"github.com/chef/automate/components/automate-deployment/pkg/cli"
	"github.com/chef/automate/lib/config"
	"github.com/chef/automate/lib/httputils"
	"github.com/chef/automate/lib/logger"
	"github.com/chef/automate/lib/platform/command"
	"github.com/chef/automate/lib/sshutils"
	"github.com/chef/automate/lib/stringutils"
	"github.com/chef/automate/lib/version"
	"github.com/spf13/cobra"
)

const (
	VERIFY_SERVER_PORT        = "VERIFY_SERVER_PORT"
	BINARY_DESTINATION_FOLDER = "/etc/automate-verify"
	SYSTEMD_PATH              = "/etc/systemd/system"
	REMOTE_NODES              = "remote nodes"
	BASTION                   = "bastion"
	SLEEP_DURATION            = 2 * time.Second
)

var (
	statusAPIEndpoint     = fmt.Sprintf("http://localhost:%s/status", getPort())
	batchCheckAPIEndpoint = fmt.Sprintf("http://localhost:%s/api/v1/checks/batch-checks", getPort())
)

type verifyCmdFlags struct {
	config      string
	debug       bool
	prettyPrint bool
}

type verifyCmdFlow struct {
	Client               httputils.HTTPClient
	CreateSystemdService verifysystemdcreate.CreateSystemdService
	SystemdCreateUtils   verifysystemdcreate.SystemdCreateUtils
	Config               *config.HaDeployConfig
	SSHUtil              sshutils.SSHUtil
	Writer               *cli.Writer
	prettyPrint          bool
}

func NewVerifyCmdFlow(client httputils.HTTPClient, createSystemdService verifysystemdcreate.CreateSystemdService, systemdCreateUtils verifysystemdcreate.SystemdCreateUtils, config *config.HaDeployConfig, sshutil sshutils.SSHUtil, writer *cli.Writer) *verifyCmdFlow {
	return &verifyCmdFlow{
		Client:               client,
		CreateSystemdService: createSystemdService,
		SystemdCreateUtils:   systemdCreateUtils,
		Config:               config,
		SSHUtil:              sshutil,
		Writer:               writer,
	}
}

type verifyServeCmdFlow struct{}

type verifySystemdCreateFlow struct{}

func init() {
	flagsObj := verifyCmdFlags{}

	verifyCmd := &cobra.Command{
		Use:   "verify",
		Short: "Verify Chef Automate setup",
		Long:  "Verify Chef Automate config files and Infrastructure",
		Annotations: map[string]string{
			docs.Compatibility: docs.Compatibility,
		},
		Args:   cobra.RangeArgs(0, 1),
		RunE:   verifyCmdFunc(&flagsObj),
		Hidden: true,
	}

	verifyServeCmd := &cobra.Command{
		Use:   "serve",
		Short: "Start verify server",
		Long:  "Start verify server for running checks",
		Annotations: map[string]string{
			docs.Compatibility: docs.Compatibility,
		},
		Args: cobra.ExactArgs(0),
		RunE: verifyServeCmdFunc(&flagsObj),
	}

	verifySystemdServiceCmd := &cobra.Command{
		Use:   "systemd-service COMMAND",
		Short: "Systemd utilities for verify command",
		Annotations: map[string]string{
			docs.Compatibility: docs.Compatibility,
		},
	}

	verifySystemdServiceCreateCmd := &cobra.Command{
		Use:   "create",
		Short: "Start verify server as systemd service",
		Annotations: map[string]string{
			docs.Compatibility: docs.Compatibility,
		},
		Args: cobra.ExactArgs(0),
		RunE: verifySystemdCreateFunc(&flagsObj),
	}

	verifyCmd.PersistentFlags().StringVarP(
		&flagsObj.config,
		"config",
		"c",
		"",
		"Config file that needs to be verified")

	verifyCmd.Flags().BoolVarP(
		&flagsObj.debug,
		"debug",
		"d",
		false,
		"enable debugging for verify command")

	verifyCmd.Flags().BoolVarP(
		&flagsObj.prettyPrint,
		"pretty-print",
		"p",
		false,
		"pretty-print json response")

	verifyServeCmd.Flags().BoolVarP(
		&flagsObj.debug,
		"debug",
		"d",
		false,
		"enable debugging for verify serve command")

	verifySystemdServiceCreateCmd.Flags().BoolVarP(
		&flagsObj.debug,
		"debug",
		"d",
		false,
		"enable debugging for verify systemd-service create command")

	verifySystemdServiceCmd.AddCommand(verifySystemdServiceCreateCmd)
	verifyCmd.AddCommand(verifySystemdServiceCmd)
	verifyCmd.AddCommand(verifyServeCmd)
	RootCmd.AddCommand(verifyCmd)
}

func verifyServeCmdFunc(flagsObj *verifyCmdFlags) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		c := verifyServeCmdFlow{}
		return c.runVerifyServeCmd(cmd, args, flagsObj.debug)
	}
}

func (v *verifyServeCmdFlow) runVerifyServeCmd(cmd *cobra.Command, args []string, debug bool) error {
	port := getPort()
	writer.Println("Using port " + port)
	vs, err := server.NewVerifyServer(port, debug)
	if err != nil {
		return err
	}
	return vs.Start()
}

func verifySystemdCreateFunc(flagsObj *verifyCmdFlags) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		c := verifySystemdCreateFlow{}
		return c.runVerifySystemdCreateCmd(cmd, args, flagsObj.debug)
	}
}

func (v *verifySystemdCreateFlow) runVerifySystemdCreateCmd(cmd *cobra.Command, args []string, debug bool) error {
	createSystemdServiceWithBinary, err := verifysystemdcreate.NewCreateSystemdService(
		verifysystemdcreate.NewSystemdCreateUtilsImpl(),
		BINARY_DESTINATION_FOLDER,
		SYSTEMD_PATH,
		debug,
		writer,
	)
	if err != nil {
		return err
	}
	return createSystemdServiceWithBinary.Create()
}

func verifyCmdFunc(flagsObj *verifyCmdFlags) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		defaultLevel := "info"
		if flagsObj.debug {
			defaultLevel = "debug"
		}

		log, err := logger.NewLogger("text", defaultLevel)
		if err != nil {
			return err
		}

		createSystemdServiceWithBinary, err := verifysystemdcreate.NewCreateSystemdService(
			verifysystemdcreate.NewSystemdCreateUtilsImpl(),
			BINARY_DESTINATION_FOLDER,
			SYSTEMD_PATH,
			flagsObj.debug,
			writer,
		)
		if err != nil {
			return err
		}

		c := NewVerifyCmdFlow(httputils.NewClient(log), createSystemdServiceWithBinary, verifysystemdcreate.NewSystemdCreateUtilsImpl(), config.NewHaDeployConfig(), sshutils.NewSSHUtilWithCommandExecutor(sshutils.NewSshClient(), log, command.NewExecExecutor()), writer)
		return c.runVerifyCmd(cmd, args, flagsObj)
	}
}

func (v *verifyCmdFlow) runVerifyCmd(cmd *cobra.Command, args []string, flagsObj *verifyCmdFlags) error {
	v.prettyPrint = flagsObj.prettyPrint
	return v.RunVerify(flagsObj.config)
}

func (v *verifyCmdFlow) RunVerify(config string) error {
	var configPath string
	// TODO : config flag is optional for now. Need to handle the default config path
	if len(strings.TrimSpace(config)) > 0 {
		configPath = config
	}

	err := v.Config.ParseAndVerify(configPath)
	if err != nil {
		return err
	}

	// Get config required for batch-check API call
	batchCheckConfig := &models.Config{
		Hardware: &models.Hardware{},
		SSHUser:  &models.SSHUser{},
		Backup: &models.Backup{
			FileSystem:    &models.FileSystem{},
			ObjectStorage: &models.ObjectStorage{},
		},
	}
	err = batchCheckConfig.PopulateWith(v.Config)
	if err != nil {
		return err
	}

	// TODO: For now just print the result as json. Need to merge the result with the response from batch-check API call for remote.

	resBastion, err := v.runVerifyServiceForBastion(*batchCheckConfig)
	if err != nil {
		return err
	}
	v.Writer.Printf("Response for batch-check API on bastion: \n%s\n", string(resBastion))

	resRemote, err := v.runVerifyServiceForRemote(*batchCheckConfig)
	if err != nil {
		return err
	}
	v.Writer.Printf("Response for batch-check API on remote nodes: \n%s\n", string(resRemote))

	return nil
}

// Runs automate-verify service on bastion
func (v *verifyCmdFlow) runVerifyServiceForBastion(batchCheckConfig models.Config) ([]byte, error) {
	v.Writer.Println("Checking automate-verify service on bastion")

	//Call status API to check if automate-verify service is running on bastion
	_, responseBody, err := v.Client.MakeRequest(http.MethodGet, statusAPIEndpoint, nil)
	if err != nil {
		v.Writer.Printf("error while checking automate-verify service on bastion: %v\n", err)
		if responseBody != nil {
			v.Writer.Printf("Error response body: %s\n", string(responseBody))
		}
		v.Writer.Println("Adding automate-verify service to systemd and starting the service")

		// create systemd service with current CLI binary
		err = v.createSystemdOnBastion()
		if err != nil {
			return nil, err
		}

		v.Writer.Println("Added automate-verify service to systemd and started the the service")
	} else {
		resultBytes, err := v.getResultFromResponseBody(responseBody)
		if err != nil {
			return nil, err
		}

		var result models.StatusDetails

		if err := json.Unmarshal(resultBytes, &result); err != nil {
			return nil, fmt.Errorf("failed to unmarshal Result field: %v", err)
		}

		// Upgrade if current CLI is latest version than the existing CLI on bastion
		latestCLIVersion := version.BuildTime
		if latestCLIVersion > result.CliVersion {
			v.Writer.Printf("Upgrading from CLI %s to latest %s on bastion\n", result.CliVersion, latestCLIVersion)
			// create systemd service with current CLI binary
			err = v.createSystemdOnBastion()
			if err != nil {
				return nil, err
			}
			v.Writer.Println("Upgrading from CLI completed on bastion")
		}
	}
	v.Writer.Println("Checking automate-verify service on bastion completed")

	// Doing batch-check API call for bastion
	batchCheckBastionReq := models.BatchCheckRequest{
		Checks: constants.GetBastionChecks(),
		Config: &batchCheckConfig,
	}

	return v.makeBatchCheckAPICall(batchCheckBastionReq, BASTION)
}

// Runs automate-verify service on remote nodes
func (v *verifyCmdFlow) runVerifyServiceForRemote(batchCheckConfig models.Config) ([]byte, error) {

	// TODO: Need to check if automate-verify service is already running on remote nodes and upgrade if needed.
	var port, keyFile, userName string
	var hostIPs []string

	if v.Config.IsExistingInfra() {
		port, keyFile, userName = v.getSSHConfig(v.Config.Architecture.ExistingInfra)
		hostIPs = v.getHostIPs(
			v.Config.ExistingInfra.Config.AutomatePrivateIps,
			v.Config.ExistingInfra.Config.ChefServerPrivateIps,
			v.Config.ExistingInfra.Config.PostgresqlPrivateIps,
			v.Config.ExistingInfra.Config.OpensearchPrivateIps,
		)

		sshConfig := sshutils.NewSshConfig("", port, keyFile, userName)

		destFileName := "chef-automate"

		// Copying CLI binary to remote nodes
		err := v.copyCLIOnRemoteNodes(destFileName, sshConfig, hostIPs)
		if err != nil {
			return nil, err
		}

		// Starting automate-verify service on remote nodes
		err = v.startServiceOnRemoteNodes(destFileName, sshConfig, hostIPs)
		if err != nil {
			return nil, err
		}
	}

	//TODO: Need to handle the case for AWS

	// Doing batch-check API call for remote nodes
	batchCheckRemoteReq := models.BatchCheckRequest{
		Checks: constants.GetRemoteChecks(),
		Config: &batchCheckConfig,
	}
	return v.makeBatchCheckAPICall(batchCheckRemoteReq, REMOTE_NODES)
}

// Creates systemd service with current CLI binary
func (v *verifyCmdFlow) createSystemdOnBastion() error {
	err := v.CreateSystemdService.Create()
	if err != nil {
		return err
	}
	// TODO: Added 2 seconds sleep to make sure service is started. Need to handle this in a better way.
	time.Sleep(SLEEP_DURATION)
	return nil
}

// Makes batch-check API call
func (v *verifyCmdFlow) makeBatchCheckAPICall(requestBody models.BatchCheckRequest, nodeType string) ([]byte, error) {
	v.Writer.Printf("Doing batch-check API call for %s\n", nodeType)
	_, responseBody, err := v.Client.MakeRequest(http.MethodPost, batchCheckAPIEndpoint, requestBody)
	if err != nil {
		if responseBody != nil {
			v.Writer.Printf("Error response body: %s\n", string(responseBody))
		}
		return nil, fmt.Errorf("error while doing batch-check API call for %s: %v", nodeType, err)
	}
	resultBytes, err := v.getResultFromResponseBody(responseBody)
	if err != nil {
		return nil, err
	}
	v.Writer.Printf("Batch-check API call for %s completed\n", nodeType)
	return resultBytes, nil
}

// Copies CLI binary to remote nodes
func (v *verifyCmdFlow) copyCLIOnRemoteNodes(destFileName string, sshConfig sshutils.SSHConfig, hostIPs []string) error {
	v.Writer.Println("Copying automate-verify CLI to remote nodes")
	currentBinaryPath, err := v.SystemdCreateUtils.GetBinaryPath()
	if err != nil {
		return err
	}
	copyResults := v.SSHUtil.CopyFileToRemoteConcurrently(sshConfig, currentBinaryPath, destFileName, false, hostIPs)
	isError := false
	for _, result := range copyResults {
		if result.Error != nil {
			v.Writer.Errorf("Remote copying automate-verify CLI failed on node : %s with error: %v\n", result.HostIP, result.Error)
			isError = true
		}
	}
	if isError {
		return fmt.Errorf("remote copying failed")
	}
	v.Writer.Println("Copying automate-verify CLI to remote nodes completed")
	return nil
}

// Starts automate-verify service on remote nodes
func (v *verifyCmdFlow) startServiceOnRemoteNodes(destFileName string, sshConfig sshutils.SSHConfig, hostIPs []string) error {
	v.Writer.Println("Creating automate-verify service on remote nodes")
	cmd := fmt.Sprintf("sudo /tmp/%s verify systemd-service create", destFileName)
	excuteResults := v.SSHUtil.ExecuteConcurrently(sshConfig, cmd, hostIPs)
	isError := false
	for _, result := range excuteResults {
		if result.Error != nil {
			v.Writer.Errorf("Remote execution of cmd '%s' automate-verify CLI failed on node : %s with error: %v\n", cmd, result.HostIP, result.Error)
			isError = true
		}
	}
	if isError {
		return fmt.Errorf("remote execution failed")
	}
	// TODO: Added 2 seconds sleep to make sure service is started. Need to handle this in a better way.
	time.Sleep(SLEEP_DURATION)
	v.Writer.Println("Created automate-verify service on remote nodes")
	return nil
}

// Returns result field from response body
func (v *verifyCmdFlow) getResultFromResponseBody(responseBody []byte) ([]byte, error) {
	var responseBodyStruct response.ResponseBody
	if err := json.Unmarshal(responseBody, &responseBodyStruct); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body: %v", err)
	}

	var resultBytes []byte
	var err error
	if v.prettyPrint {
		resultBytes, err = json.MarshalIndent(responseBodyStruct.Result, "", "    ")
	} else {
		resultBytes, err = json.Marshal(responseBodyStruct.Result)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to marshal result field: %v", err)
	}

	return resultBytes, nil
}

// Returns port, keyFile, userName
func (v *verifyCmdFlow) getSSHConfig(config *config.ConfigInitials) (string, string, string) {
	return config.SSHPort, config.SSHKeyFile, config.SSHUser
}

// Returns concatenated list of automatePrivateIps, chefServerPrivateIps, postgresqlPrivateIps, opensearchPrivateIps
func (v *verifyCmdFlow) getHostIPs(automatePrivateIps, chefServerPrivateIps, postgresqlPrivateIps, opensearchPrivateIps []string) []string {
	var hostIPs []string
	hostIPs = stringutils.ConcatSlice(hostIPs, automatePrivateIps)
	hostIPs = stringutils.ConcatSlice(hostIPs, chefServerPrivateIps)
	if !v.Config.IsExternalDb() {
		hostIPs = stringutils.ConcatSlice(hostIPs, postgresqlPrivateIps)
		hostIPs = stringutils.ConcatSlice(hostIPs, opensearchPrivateIps)
	}
	return hostIPs
}

// Returns port
func getPort() string {
	port := server.DEFAULT_PORT
	envPort := os.Getenv(VERIFY_SERVER_PORT)
	if envPort != "" {
		if _, err := strconv.Atoi(envPort); err == nil {
			port = envPort
		}
	}
	return port
}
