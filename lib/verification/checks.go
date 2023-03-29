package verification

import (
	sc "github.com/chef/automate/api/config/deployment"
	"github.com/chef/automate/lib/reporting"
)

type Result struct {
	Valid  bool
	Report reporting.Info
}

func validateAWSProvisionConfig(config *HAAwsConfigToml) {

	/* Write the checks to validate if all the required fields for AWS Infra Provision without Managed Services
	are filled in the parsed config */

	if config.Aws.Config.SetupManagedServices {

	}

}

func validateAWSManagedProvisionConfig(config *HAAwsConfigToml) {

	/* Write the checks to validate if all the required fields for AWS Infra Provision with Managed Services
	are filled in the parsed config */

}

func validateAWSDeploymentConfig(config *HAAwsConfigToml) {

	/* Write the checks to validate if all the required fields for AWS Deployment without Managed Services
	are filled in the parsed config */

}

func validateAWSManagedDeploymentConfig(config *HAAwsConfigToml) {

	/* Write the checks to validate if all the required fields for AWS Deployment with Managed Services
	are filled in the parsed config */

}

func validateOnPremConfig(config *HAOnPremConfigToml) {

	/* Write the checks to validate if all the required fields for On-Prem Deployment with Chef Managed Services
	are filled in the parsed config */

}

func validateOnPremAWSConfig(config *HAOnPremConfigToml) {

	/* Write the checks to validate if all the required fields for On-Prem Deployment with AWS Managed Services
	are filled in the parsed config */

}

func validateOnPremCustomerConfig(config *HAOnPremConfigToml) {

	/* Write the checks to validate if all the required fields for On-Prem Deployment with Customer Managed Services
	are filled in the parsed config */

}

func validateStandaloneDeploymentConfig(config *sc.AutomateConfig) {

	/* Write the checks to validate if all the required fields for On-Prem Deployment with Customer Managed Services
	are filled in the parsed config */

}

func validateNodeReachability(ipaddress string, channel chan reporting.VerfictionReport) {

	/* Write the checks to validate if the nodes are reachable */

	result := reporting.Info{
		Hostip:    ipaddress,
		Parameter: "Node Reachabilility",
		Status:    "Success",
		StatusMessage: &reporting.StatusMessage{
			MainMessage: "Node A is reachable",
		},
	}

	chanWriter(channel, "Automate", result)

}

func validateCertificateExpiry(certContents string) Result {

	return Result{
		Valid: true,
		Report: reporting.Info{
			Hostip:    "172.02.0.01",
			Parameter: "Certificates",
			Status:    "Success",
			StatusMessage: &reporting.StatusMessage{
				MainMessage: "Certificate A validated successfully",
			},
		},
	}

	// return Result{
	// 	Valid: false,
	// 	Report: reporting.Info{
	// 		Hostip:    "172.02.0.01",
	// 		Parameter: "Certificates",
	// 		Status:    "Failed",
	// 		StatusMessage: &reporting.StatusMessage{
	// 			MainMessage: "Certificate B validation failed",
	// 			SubMessage:  []string{"Certificate is not valid"},
	// 			ToResolve:   []string{"1. Renew the expired certificates"},
	// 		},
	// 	},
	// }

	/* Write the checks to validate the expiry date of the certificate */

}

func validateCertificateFormat(certContents string) Result {

	return Result{
		Valid: true,
		Report: reporting.Info{
			Hostip:    "172.02.20.01",
			Parameter: "Certificates",
			Status:    "Failed",
			StatusMessage: &reporting.StatusMessage{
				MainMessage: "Certificate B validation failed",
				SubMessage:  []string{"Certificate is not formatted properly"},
				ToResolve:   []string{"1. Add the start and end certificate markers"},
			},
		},
	}

	/* Write the checks to validate the expiry date of the certificate */

}
