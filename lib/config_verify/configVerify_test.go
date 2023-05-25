package config_verify

import (
	"testing"

	"github.com/chef/automate/lib/config_parser"
)

func TestConfigValidateOnPrem(t *testing.T) {
	t.Run("Test Success config Validation", func(t *testing.T) {
		config := GetTestOnPremConfigData()

		cv := &ConfigVerifyImpl{}
		err := cv.ConfigValidateOnPrem(config)
		if err != nil {
			t.Errorf("ConfigInitials validation failed: %s", err)
		}
	})

	t.Run("Test Failed Config Validation", func(t *testing.T) {
		config := GetTestOnPremFailedConfigData()

		cv := &ConfigVerifyImpl{}
		err := cv.ConfigValidateOnPrem(config)
		if err != nil {
			return
		}
		// If there is no error, fail the test case
		t.Fail()
	})

	t.Run("Test Failed On prem AWS managed Config Validation", func(t *testing.T) {
		config := GetTestOnPremFailedConfigData()
		config.ExternalDB.Database.Type = "aws"
		cv := &ConfigVerifyImpl{}
		err := cv.ConfigValidateOnPrem(config)
		if err != nil {
			return
		}
		// If there is no error, fail the test case
		t.Fail()
	})

	t.Run("Test Failed Self-managed Config Validation", func(t *testing.T) {
		config := GetTestOnPremFailedConfigData()
		config.ExternalDB.Database.Type = "self-managed"
		cv := &ConfigVerifyImpl{}
		err := cv.ConfigValidateOnPrem(config)
		if err != nil {
			return
		}
		// If there is no error, fail the test case
		t.Fail()
	})

	t.Run("Test Success AWS Config Validation", func(t *testing.T) {
		config := GetTestAWSConfigData()

		cv := &ConfigVerifyImpl{}
		err := cv.ConfigValidateAWS(config)
		if err != nil {
			t.Errorf("ConfigInitials validation failed: %s", err)
		}
	})

	t.Run("Test Failed AWS Config Validation", func(t *testing.T) {
		config := GetFailedTestAWSConfigData()

		cv := &ConfigVerifyImpl{}
		err := cv.ConfigValidateAWS(config)
		if err != nil {
			return
		}
		// If there is no error, fail the test case
		t.Fail()
	})

	t.Run("Test Success Automate SC Config Validation", func(t *testing.T) {
		config := GetAutomateScConfigTestData()
		cv := &ConfigVerifyImpl{}
		err := cv.ConfigValidateStandalone(config)
		if err != nil {
			t.Errorf("ConfigInitials validation failed: %s", err)
		}
	})

	t.Run("Test Failed Automate SC Config Validation", func(t *testing.T) {
		config := GetFailedAutomateScConfigTestData()
		cv := &ConfigVerifyImpl{}
		err := cv.ConfigValidateStandalone(config)
		if err != nil {
			return
		}
		// If there is no error, fail the test case
		t.Fail()
	})

	t.Run("Test Failed Scenario where HA Onprem Automate CertByIp is not provided", func(t *testing.T) {
		config := GetTestOnPremFailedConfigData()
		config.Automate.Config.EnableCustomCerts = true
		config.Automate.Config.PrivateKey = ""
		config.Automate.Config.PublicKey = "/path/key"
		config.Automate.Config.CertsByIP = []config_parser.CertByIP{}
		cv := &ConfigVerifyImpl{}
		err := cv.ConfigValidateOnPrem(config)
		if err != nil {
			return
		}
		// If there is no error, fail the test case
		t.Fail()
	})

	t.Run("Test Failed Scenario where HA Onprem ChefServer CertByIp is not provided", func(t *testing.T) {
		config := GetTestOnPremFailedConfigData()
		config.ChefServer.Config.EnableCustomCerts = true
		config.ChefServer.Config.PrivateKey = ""
		config.ChefServer.Config.PrivateKey = ""
		config.ChefServer.Config.CertsByIP = []config_parser.CertByIP{}
		cv := &ConfigVerifyImpl{}
		err := cv.ConfigValidateOnPrem(config)
		if err != nil {
			return
		}
		// If there is no error, fail the test case
		t.Fail()
	})

	t.Run("Test Failed Scenario where HA Onprem Postgres CertByIp is not provided", func(t *testing.T) {
		config := GetTestOnPremFailedConfigData()
		config.Postgresql.Config.EnableCustomCerts = true
		config.Postgresql.Config.PrivateKey = ""
		config.Postgresql.Config.PrivateKey = ""
		config.Postgresql.Config.CertsByIP = []config_parser.CertByIP{}
		cv := &ConfigVerifyImpl{}
		err := cv.ConfigValidateOnPrem(config)
		if err != nil {
			return
		}
		// If there is no error, fail the test case
		t.Fail()
	})

	t.Run("Test Failed Scenario where HA Onprem OpenSearch CertByIp is not provided", func(t *testing.T) {
		config := GetTestOnPremFailedConfigData()
		config.Opensearch.Config.EnableCustomCerts = true
		config.Opensearch.Config.PrivateKey = ""
		config.Opensearch.Config.PrivateKey = ""
		config.Opensearch.Config.CertsByIP = []config_parser.CertByIP{}
		cv := &ConfigVerifyImpl{}
		err := cv.ConfigValidateOnPrem(config)
		if err != nil {
			return
		}
		// If there is no error, fail the test case
		t.Fail()
	})

}
