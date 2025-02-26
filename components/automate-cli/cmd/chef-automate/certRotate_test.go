package main

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/chef/automate/api/config/deployment"
	"github.com/chef/automate/api/config/shared"
	"github.com/chef/automate/lib/io/fileutils"
	"github.com/chef/automate/lib/sshutils"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const (
	RemoteFilePath = "198.51.100.0:/home/ec2-user/certs/public.pem"
	LocalFilePath  = "/home/ec2-user/certs/public.pem"
	ValidIP        = "198.51.100.0"
	ValidIP1       = "198.51.100.1"
	ValidIP2       = "198.51.100.2"
	ValidIP3       = "198.51.100.3"
	ValidIP4       = "198.51.100.4"
	ValidIP5       = "198.51.100.5"
	ValidIP6       = "198.51.100.6"
	ValidIP7       = "198.51.100.7"
	ValidIP8       = "198.51.100.8"
	ValidIP9       = "198.51.100.9"
	FileContent    = `-----BEGIN CERTIFICATE-----
MIIEDzCCAvegAwIBAgIBADANBgkqhkiG9w0BAQUFADBoMQswCQYDVQQGEwJVUzEl
MCMGA1UEChMcU3RhcmZpZWxkIFRlY2hub2xvZ2llcywgSW5jLjEyMDAGA1UECxMp
U3RhcmZpZWxkIENsYXNzIDIgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkwHhcNMDQw
NjI5MTczOTE2WhcNMzQwNjI5MTczOTE2WjBoMQswCQYDVQQGEwJVUzElMCMGA1UE
ChMcU3RhcmZpZWxkIFRlY2hub2xvZ2llcywgSW5jLjEyMDAGA1UECxMpU3RhcmZp
ZWxkIENsYXNzIDIgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkwggEgMA0GCSqGSIb3
DQEBAQUAA4IBDQAwggEIAoIBAQC3Msj+6XGmBIWtDBFk385N78gDGIc/oav7PKaf
8MOh2tTYbitTkPskpD6E8J7oX+zlJ0T1KKY/e97gKvDIr1MvnsoFAZMej2YcOadN
+lq2cwQlZut3f+dZxkqZJRRU6ybH838Z1TBwj6+wRir/resp7defqgSHo9T5iaU0
X9tDkYI22WY8sbi5gv2cOj4QyDvvBmVmepsZGD3/cVE8MC5fvj13c7JdBmzDI1aa
K4UmkhynArPkPw2vCHmCuDY96pzTNbO8acr1zJ3o/WSNF4Azbl5KXZnJHoe0nRrA
1W4TNSNe35tfPe/W93bC6j67eA0cQmdrBNj41tpvi/JEoAGrAgEDo4HFMIHCMB0G
A1UdDgQWBBS/X7fRzt0fhvRbVazc1xDCDqmI5zCBkgYDVR0jBIGKMIGHgBS/X7fR
zt0fhvRbVazc1xDCDqmI56FspGowaaBCDEfGA1UEBhMCVVMxJTAjBgNVBAoTHFN0
YXJmaWVsZCBUZWNobm9sb2dpZXMsIEluYy4xMjAwBgNVBAsTKVN0YXJmaWVsZCBD
bGFzcyAyIENlcnRpZmljYXRpb24gQXV0aG9yaXR5ggEAMAwGA1UdEwQFMAMBAf8w
DQYJKoZIhvcNAQEFBQADggEBAAWdP4id0ckaVaGsafPzWdqbAYcaT1epoXkJKtv3
L7IezMdeatiDh6GX70k1PncGQVhiv45YuApnP+yz3SFmH8lU+nLMPUxA2IGvd56D
eruix/U0F47ZEUD0/CwqTRV/p2JdLiXTAAsgGh1o+Re49L2L7ShZ3U0WixeDyLJl
xy16paq8U4Zt3VekyvggQQto8PT7dL5WXXp59fkdheMtlb71cZBDzI0fmgAKhynp
VSJYACPq4xJDKVtHCN2MQWplBqjlIapBtJUhlbl90TSrE9atvNziPTnNvT51cKEY
WQPJIrSPnNVeKtelttQKbfi3QBFGmh95DmK/D5fs4C8xY5Z=
-----END CERTIFICATE-----`
	ValidCertPath                = "./certRotate.go"
	TestOpensearchAdminAndRootCA = `
	[tls]
	rootCA = """-----BEGIN CERTIFICATE-----
	MIIEDzCCAvegAwIBAgIBADANBgkqhkiG9w0BAQUFADBoMQswCQYDVQQGEwJVUzEl
	MCMGA1UEChMcU3RhcmZpZWxkIFRlY2hub2xvZ2llcywgSW5jLjEyMDAGA1UECxMp
	U3RhcmZpZWxkIENsYXNzIDIgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkwHhcNMDQw
	NjI5MTczOTE2WhcNMzQwNjI5MTczOTE2WjBoMQswCQYDVQQGEwJVUzElMCMGA1UE
	ChMcU3RhcmZpZWxkIFRlY2hub2xvZ2llcywgSW5jLjEyMDAGA1UECxMpU3RhcmZp
	ZWxkIENsYXNzIDIgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkwggEgMA0GCSqGSIb3
	DQEBAQUAA4IBDQAwggEIAoIBAQC3Msj+6XGmBIWtDBFk385N78gDGIc/oav7PKaf
	8MOh2tTYbitTkPskpD6E8J7oX+zlJ0T1KKY/e97gKvDIr1MvnsoFAZMej2YcOadN
	+lq2cwQlZut3f+dZxkqZJRRU6ybH838Z1TBwj6+wRir/resp7defqgSHo9T5iaU0
	X9tDkYI22WY8sbi5gv2cOj4QyDvvBmVmepsZGD3/cVE8MC5fvj13c7JdBmzDI1aa
	K4UmkhynArPkPw2vCHmCuDY96pzTNbO8acr1zJ3o/WSNF4Azbl5KXZnJHoe0nRrA
	1W4TNSNe35tfPe/W93bC6j67eA0cQmdrBNj41tpvi/JEoAGrAgEDo4HFMIHCMB0G
	A1UdDgQWBBS/X7fRzt0fhvRbVazc1xDCDqmI5zCBkgYDVR0jBIGKMIGHgBS/X7fR
	zt0fhvRbVazc1xDCDqmI56FspGowaaBCDEfGA1UEBhMCVVMxJTAjBgNVBAoTHFN0
	YXJmaWVsZCBUZWNobm9sb2dpZXMsIEluYy4xMjAwBgNVBAsTKVN0YXJmaWVsZCBD
	bGFzcyAyIENlcnRpZmljYXRpb24gQXV0aG9yaXR5ggEAMAwGA1UdEwQFMAMBAf8w
	DQYJKoZIhvcNAQEFBQADggEBAAWdP4id0ckaVaGsafPzWdqbAYcaT1epoXkJKtv3
	L7IezMdeatiDh6GX70k1PncGQVhiv45YuApnP+yz3SFmH8lU+nLMPUxA2IGvd56D
	eruix/U0F47ZEUD0/CwqTRV/p2JdLiXTAAsgGh1o+Re49L2L7ShZ3U0WixeDyLJl
	xy16paq8U4Zt3VekyvggQQto8PT7dL5WXXp59fkdheMtlb71cZBDzI0fmgAKhynp
	VSJYACPq4xJDKVtHCN2MQWplBqjlIapBtJUhlbl90TSrE9atvNziPTnNvT51cKEY
	WQPJIrSPnNVeKtelttQKbfi3QBFGmh95DmK/D5fs4C8xY5Z=
	-----END CERTIFICATE-----"""
	admin_cert = """-----BEGIN CERTIFICATE-----
	MIIEDzCCAvegAwIBAgIBADANBgkqhkiG9w0BAQUFADBoMQswCQYDVQQGEwJVUzEl
	MCMGA1UEChMcU3RhcmZpZWxkIFRlY2hub2xvZ2llcywgSW5jLjEyMDAGA1UECxMp
	U3RhcmZpZWxkIENsYXNzIDIgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkwHhcNMDQw
	NjI5MTczOTE2WhcNMzQwNjI5MTczOTE2WjBoMQswCQYDVQQGEwJVUzElMCMGA1UE
	ChMcU3RhcmZpZWxkIFRlY2hub2xvZ2llcywgSW5jLjEyMDAGA1UECxMpU3RhcmZp
	ZWxkIENsYXNzIDIgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkwggEgMA0GCSqGSIb3
	DQEBAQUAA4IBDQAwggEIAoIBAQC3Msj+6XGmBIWtDBFk385N78gDGIc/oav7PKaf
	8MOh2tTYbitTkPskpD6E8J7oX+zlJ0T1KKY/e97gKvDIr1MvnsoFAZMej2YcOadN
	+lq2cwQlZut3f+dZxkqZJRRU6ybH838Z1TBwj6+wRir/resp7defqgSHo9T5iaU0
	X9tDkYI22WY8sbi5gv2cOj4QyDvvBmVmepsZGD3/cVE8MC5fvj13c7JdBmzDI1aa
	K4UmkhynArPkPw2vCHmCuDY96pzTNbO8acr1zJ3o/WSNF4Azbl5KXZnJHoe0nRrA
	1W4TNSNe35tfPe/W93bC6j67eA0cQmdrBNj41tpvi/JEoAGrAgEDo4HFMIHCMB0G
	A1UdDgQWBBS/X7fRzt0fhvRbVazc1xDCDqmI5zCBkgYDVR0jBIGKMIGHgBS/X7fR
	zt0fhvRbVazc1xDCDqmI56FspGowaaBCDEfGA1UEBhMCVVMxJTAjBgNVBAoTHFN0
	YXJmaWVsZCBUZWNobm9sb2dpZXMsIEluYy4xMjAwBgNVBAsTKVN0YXJmaWVsZCBD
	bGFzcyAyIENlcnRpZmljYXRpb24gQXV0aG9yaXR5ggEAMAwGA1UdEwQFMAMBAf8w
	DQYJKoZIhvcNAQEFBQADggEBAAWdP4id0ckaVaGsafPzWdqbAYcaT1epoXkJKtv3
	L7IezMdeatiDh6GX70k1PncGQVhiv45YuApnP+yz3SFmH8lU+nLMPUxA2IGvd56D
	eruix/U0F47ZEUD0/CwqTRV/p2JdLiXTAAsgGh1o+Re49L2L7ShZ3U0WixeDyLJl
	xy16paq8U4Zt3VekyvggQQto8PT7dL5WXXp59fkdheMtlb71cZBDzI0fmgAKhynp
	VSJYACPq4xJDKVtHCN2MQWplBqjlIapBtJUhlbl90TSrE9atvNziPTnNvT51cKEY
	WQPJIrSPnNVeKtelttQKbfi3QBFGmh95DmK/D5fs4C8xY5Z=
	-----END CERTIFICATE-----"""
	admin_key = """-----BEGIN CERTIFICATE-----
	MIIEDzCCAvegAwIBAgIBADANBgkqhkiG9w0BAQUFADBoMQswCQYDVQQGEwJVUzEl
	MCMGA1UEChMcU3RhcmZpZWxkIFRlY2hub2xvZ2llcywgSW5jLjEyMDAGA1UECxMp
	U3RhcmZpZWxkIENsYXNzIDIgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkwHhcNMDQw
	NjI5MTczOTE2WhcNMzQwNjI5MTczOTE2WjBoMQswCQYDVQQGEwJVUzElMCMGA1UE
	ChMcU3RhcmZpZWxkIFRlY2hub2xvZ2llcywgSW5jLjEyMDAGA1UECxMpU3RhcmZp
	ZWxkIENsYXNzIDIgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkwggEgMA0GCSqGSIb3
	DQEBAQUAA4IBDQAwggEIAoIBAQC3Msj+6XGmBIWtDBFk385N78gDGIc/oav7PKaf
	8MOh2tTYbitTkPskpD6E8J7oX+zlJ0T1KKY/e97gKvDIr1MvnsoFAZMej2YcOadN
	+lq2cwQlZut3f+dZxkqZJRRU6ybH838Z1TBwj6+wRir/resp7defqgSHo9T5iaU0
	X9tDkYI22WY8sbi5gv2cOj4QyDvvBmVmepsZGD3/cVE8MC5fvj13c7JdBmzDI1aa
	K4UmkhynArPkPw2vCHmCuDY96pzTNbO8acr1zJ3o/WSNF4Azbl5KXZnJHoe0nRrA
	1W4TNSNe35tfPe/W93bC6j67eA0cQmdrBNj41tpvi/JEoAGrAgEDo4HFMIHCMB0G
	A1UdDgQWBBS/X7fRzt0fhvRbVazc1xDCDqmI5zCBkgYDVR0jBIGKMIGHgBS/X7fR
	zt0fhvRbVazc1xDCDqmI56FspGowaaBCDEfGA1UEBhMCVVMxJTAjBgNVBAoTHFN0
	YXJmaWVsZCBUZWNobm9sb2dpZXMsIEluYy4xMjAwBgNVBAsTKVN0YXJmaWVsZCBD
	bGFzcyAyIENlcnRpZmljYXRpb24gQXV0aG9yaXR5ggEAMAwGA1UdEwQFMAMBAf8w
	DQYJKoZIhvcNAQEFBQADggEBAAWdP4id0ckaVaGsafPzWdqbAYcaT1epoXkJKtv3
	L7IezMdeatiDh6GX70k1PncGQVhiv45YuApnP+yz3SFmH8lU+nLMPUxA2IGvd56D
	eruix/U0F47ZEUD0/CwqTRV/p2JdLiXTAAsgGh1o+Re49L2L7ShZ3U0WixeDyLJl
	xy16paq8U4Zt3VekyvggQQto8PT7dL5WXXp59fkdheMtlb71cZBDzI0fmgAKhynp
	VSJYACPq4xJDKVtHCN2MQWplBqjlIapBtJUhlbl90TSrE9atvNziPTnNvT51cKEY
	WQPJIrSPnNVeKtelttQKbfi3QBFGmh95DmK/D5fs4C8xY5Z=
	-----END CERTIFICATE-----"""
	ssl_cert = """-----BEGIN CERTIFICATE-----
	MIIEDzCCAvegAwIBAgIBADANBgkqhkiG9w0BAQUFADBoMQswCQYDVQQGEwJVUzEl
	MCMGA1UEChMcU3RhcmZpZWxkIFRlY2hub2xvZ2llcywgSW5jLjEyMDAGA1UECxMp
	U3RhcmZpZWxkIENsYXNzIDIgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkwHhcNMDQw
	NjI5MTczOTE2WhcNMzQwNjI5MTczOTE2WjBoMQswCQYDVQQGEwJVUzElMCMGA1UE
	ChMcU3RhcmZpZWxkIFRlY2hub2xvZ2llcywgSW5jLjEyMDAGA1UECxMpU3RhcmZp
	ZWxkIENsYXNzIDIgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkwggEgMA0GCSqGSIb3
	DQEBAQUAA4IBDQAwggEIAoIBAQC3Msj+6XGmBIWtDBFk385N78gDGIc/oav7PKaf
	8MOh2tTYbitTkPskpD6E8J7oX+zlJ0T1KKY/e97gKvDIr1MvnsoFAZMej2YcOadN
	+lq2cwQlZut3f+dZxkqZJRRU6ybH838Z1TBwj6+wRir/resp7defqgSHo9T5iaU0
	X9tDkYI22WY8sbi5gv2cOj4QyDvvBmVmepsZGD3/cVE8MC5fvj13c7JdBmzDI1aa
	K4UmkhynArPkPw2vCHmCuDY96pzTNbO8acr1zJ3o/WSNF4Azbl5KXZnJHoe0nRrA
	1W4TNSNe35tfPe/W93bC6j67eA0cQmdrBNj41tpvi/JEoAGrAgEDo4HFMIHCMB0G
	A1UdDgQWBBS/X7fRzt0fhvRbVazc1xDCDqmI5zCBkgYDVR0jBIGKMIGHgBS/X7fR
	zt0fhvRbVazc1xDCDqmI56FspGowaaBCDEfGA1UEBhMCVVMxJTAjBgNVBAoTHFN0
	YXJmaWVsZCBUZWNobm9sb2dpZXMsIEluYy4xMjAwBgNVBAsTKVN0YXJmaWVsZCBD
	bGFzcyAyIENlcnRpZmljYXRpb24gQXV0aG9yaXR5ggEAMAwGA1UdEwQFMAMBAf8w
	DQYJKoZIhvcNAQEFBQADggEBAAWdP4id0ckaVaGsafPzWdqbAYcaT1epoXkJKtv3
	L7IezMdeatiDh6GX70k1PncGQVhiv45YuApnP+yz3SFmH8lU+nLMPUxA2IGvd56D
	eruix/U0F47ZEUD0/CwqTRV/p2JdLiXTAAsgGh1o+Re49L2L7ShZ3U0WixeDyLJl
	xy16paq8U4Zt3VekyvggQQto8PT7dL5WXXp59fkdheMtlb71cZBDzI0fmgAKhynp
	VSJYACPq4xJDKVtHCN2MQWplBqjlIapBtJUhlbl90TSrE9atvNziPTnNvT51cKEY
	WQPJIrSPnNVeKtelttQKbfi3QBFGmh95DmK/D5fs4C8xY5Z=
	-----END CERTIFICATE-----"""
	ssl_key = """-----BEGIN CERTIFICATE-----
	MIIEDzCCAvegAwIBAgIBADANBgkqhkiG9w0BAQUFADBoMQswCQYDVQQGEwJVUzEl
	MCMGA1UEChMcU3RhcmZpZWxkIFRlY2hub2xvZ2llcywgSW5jLjEyMDAGA1UECxMp
	U3RhcmZpZWxkIENsYXNzIDIgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkwHhcNMDQw
	NjI5MTczOTE2WhcNMzQwNjI5MTczOTE2WjBoMQswCQYDVQQGEwJVUzElMCMGA1UE
	ChMcU3RhcmZpZWxkIFRlY2hub2xvZ2llcywgSW5jLjEyMDAGA1UECxMpU3RhcmZp
	ZWxkIENsYXNzIDIgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkwggEgMA0GCSqGSIb3
	DQEBAQUAA4IBDQAwggEIAoIBAQC3Msj+6XGmBIWtDBFk385N78gDGIc/oav7PKaf
	8MOh2tTYbitTkPskpD6E8J7oX+zlJ0T1KKY/e97gKvDIr1MvnsoFAZMej2YcOadN
	+lq2cwQlZut3f+dZxkqZJRRU6ybH838Z1TBwj6+wRir/resp7defqgSHo9T5iaU0
	X9tDkYI22WY8sbi5gv2cOj4QyDvvBmVmepsZGD3/cVE8MC5fvj13c7JdBmzDI1aa
	K4UmkhynArPkPw2vCHmCuDY96pzTNbO8acr1zJ3o/WSNF4Azbl5KXZnJHoe0nRrA
	1W4TNSNe35tfPe/W93bC6j67eA0cQmdrBNj41tpvi/JEoAGrAgEDo4HFMIHCMB0G
	A1UdDgQWBBS/X7fRzt0fhvRbVazc1xDCDqmI5zCBkgYDVR0jBIGKMIGHgBS/X7fR
	zt0fhvRbVazc1xDCDqmI56FspGowaaBCDEfGA1UEBhMCVVMxJTAjBgNVBAoTHFN0
	YXJmaWVsZCBUZWNobm9sb2dpZXMsIEluYy4xMjAwBgNVBAsTKVN0YXJmaWVsZCBD
	bGFzcyAyIENlcnRpZmljYXRpb24gQXV0aG9yaXR5ggEAMAwGA1UdEwQFMAMBAf8w
	DQYJKoZIhvcNAQEFBQADggEBAAWdP4id0ckaVaGsafPzWdqbAYcaT1epoXkJKtv3
	L7IezMdeatiDh6GX70k1PncGQVhiv45YuApnP+yz3SFmH8lU+nLMPUxA2IGvd56D
	eruix/U0F47ZEUD0/CwqTRV/p2JdLiXTAAsgGh1o+Re49L2L7ShZ3U0WixeDyLJl
	xy16paq8U4Zt3VekyvggQQto8PT7dL5WXXp59fkdheMtlb71cZBDzI0fmgAKhynp
	VSJYACPq4xJDKVtHCN2MQWplBqjlIapBtJUhlbl90TSrE9atvNziPTnNvT51cKEY
	WQPJIrSPnNVeKtelttQKbfi3QBFGmh95DmK/D5fs4C8xY5Z=
	-----END CERTIFICATE-----"""
[plugins.security.authcz]
	admin_dn = '- CN=chefadmin,O=Chef Software Inc,L=Seattle,ST=Washington,C=US'
[plugins.security.ssl.transport]
	enforce_hostname_verification = false
	resolve_hostname = false
[plugins.security]
	nodes_dn = '- CN=chefnode1,O=Chef Software Inc,L=Seattle,ST=Washington,C=US'`
	TestFrontendConfig = `[[load_balancer.v1.sys.frontend_tls]]
	cert = """-----BEGIN CERTIFICATE-----
	MIIEDzCCAvegAwIBAgIBADANBgkqhkiG9w0BAQUFADBoMQswCQYDVQQGEwJVUzEl
	MCMGA1UEChMcU3RhcmZpZWxkIFRlY2hub2xvZ2llcywgSW5jLjEyMDAGA1UECxMp
	U3RhcmZpZWxkIENsYXNzIDIgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkwHhcNMDQw
	NjI5MTczOTE2WhcNMzQwNjI5MTczOTE2WjBoMQswCQYDVQQGEwJVUzElMCMGA1UE
	ChMcU3RhcmZpZWxkIFRlY2hub2xvZ2llcywgSW5jLjEyMDAGA1UECxMpU3RhcmZp
	ZWxkIENsYXNzIDIgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkwggEgMA0GCSqGSIb3
	DQEBAQUAA4IBDQAwggEIAoIBAQC3Msj+6XGmBIWtDBFk385N78gDGIc/oav7PKaf
	8MOh2tTYbitTkPskpD6E8J7oX+zlJ0T1KKY/e97gKvDIr1MvnsoFAZMej2YcOadN
	+lq2cwQlZut3f+dZxkqZJRRU6ybH838Z1TBwj6+wRir/resp7defqgSHo9T5iaU0
	X9tDkYI22WY8sbi5gv2cOj4QyDvvBmVmepsZGD3/cVE8MC5fvj13c7JdBmzDI1aa
	K4UmkhynArPkPw2vCHmCuDY96pzTNbO8acr1zJ3o/WSNF4Azbl5KXZnJHoe0nRrA
	1W4TNSNe35tfPe/W93bC6j67eA0cQmdrBNj41tpvi/JEoAGrAgEDo4HFMIHCMB0G
	A1UdDgQWBBS/X7fRzt0fhvRbVazc1xDCDqmI5zCBkgYDVR0jBIGKMIGHgBS/X7fR
	zt0fhvRbVazc1xDCDqmI56FspGowaaBCDEfGA1UEBhMCVVMxJTAjBgNVBAoTHFN0
	YXJmaWVsZCBUZWNobm9sb2dpZXMsIEluYy4xMjAwBgNVBAsTKVN0YXJmaWVsZCBD
	bGFzcyAyIENlcnRpZmljYXRpb24gQXV0aG9yaXR5ggEAMAwGA1UdEwQFMAMBAf8w
	DQYJKoZIhvcNAQEFBQADggEBAAWdP4id0ckaVaGsafPzWdqbAYcaT1epoXkJKtv3
	L7IezMdeatiDh6GX70k1PncGQVhiv45YuApnP+yz3SFmH8lU+nLMPUxA2IGvd56D
	eruix/U0F47ZEUD0/CwqTRV/p2JdLiXTAAsgGh1o+Re49L2L7ShZ3U0WixeDyLJl
	xy16paq8U4Zt3VekyvggQQto8PT7dL5WXXp59fkdheMtlb71cZBDzI0fmgAKhynp
	VSJYACPq4xJDKVtHCN2MQWplBqjlIapBtJUhlbl90TSrE9atvNziPTnNvT51cKEY
	WQPJIrSPnNVeKtelttQKbfi3QBFGmh95DmK/D5fs4C8xY5Z=
	-----END CERTIFICATE-----"""
	key = """-----BEGIN CERTIFICATE-----
	MIIEDzCCAvegAwIBAgIBADANBgkqhkiG9w0BAQUFADBoMQswCQYDVQQGEwJVUzEl
	MCMGA1UEChMcU3RhcmZpZWxkIFRlY2hub2xvZ2llcywgSW5jLjEyMDAGA1UECxMp
	U3RhcmZpZWxkIENsYXNzIDIgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkwHhcNMDQw
	NjI5MTczOTE2WhcNMzQwNjI5MTczOTE2WjBoMQswCQYDVQQGEwJVUzElMCMGA1UE
	ChMcU3RhcmZpZWxkIFRlY2hub2xvZ2llcywgSW5jLjEyMDAGA1UECxMpU3RhcmZp
	ZWxkIENsYXNzIDIgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkwggEgMA0GCSqGSIb3
	DQEBAQUAA4IBDQAwggEIAoIBAQC3Msj+6XGmBIWtDBFk385N78gDGIc/oav7PKaf
	8MOh2tTYbitTkPskpD6E8J7oX+zlJ0T1KKY/e97gKvDIr1MvnsoFAZMej2YcOadN
	+lq2cwQlZut3f+dZxkqZJRRU6ybH838Z1TBwj6+wRir/resp7defqgSHo9T5iaU0
	X9tDkYI22WY8sbi5gv2cOj4QyDvvBmVmepsZGD3/cVE8MC5fvj13c7JdBmzDI1aa
	K4UmkhynArPkPw2vCHmCuDY96pzTNbO8acr1zJ3o/WSNF4Azbl5KXZnJHoe0nRrA
	1W4TNSNe35tfPe/W93bC6j67eA0cQmdrBNj41tpvi/JEoAGrAgEDo4HFMIHCMB0G
	A1UdDgQWBBS/X7fRzt0fhvRbVazc1xDCDqmI5zCBkgYDVR0jBIGKMIGHgBS/X7fR
	zt0fhvRbVazc1xDCDqmI56FspGowaaBCDEfGA1UEBhMCVVMxJTAjBgNVBAoTHFN0
	YXJmaWVsZCBUZWNobm9sb2dpZXMsIEluYy4xMjAwBgNVBAsTKVN0YXJmaWVsZCBD
	bGFzcyAyIENlcnRpZmljYXRpb24gQXV0aG9yaXR5ggEAMAwGA1UdEwQFMAMBAf8w
	DQYJKoZIhvcNAQEFBQADggEBAAWdP4id0ckaVaGsafPzWdqbAYcaT1epoXkJKtv3
	L7IezMdeatiDh6GX70k1PncGQVhiv45YuApnP+yz3SFmH8lU+nLMPUxA2IGvd56D
	eruix/U0F47ZEUD0/CwqTRV/p2JdLiXTAAsgGh1o+Re49L2L7ShZ3U0WixeDyLJl
	xy16paq8U4Zt3VekyvggQQto8PT7dL5WXXp59fkdheMtlb71cZBDzI0fmgAKhynp
	VSJYACPq4xJDKVtHCN2MQWplBqjlIapBtJUhlbl90TSrE9atvNziPTnNvT51cKEY
	WQPJIrSPnNVeKtelttQKbfi3QBFGmh95DmK/D5fs4C8xY5Z=
	-----END CERTIFICATE-----"""
[[global.v1.frontend_tls]]
	cert = """-----BEGIN CERTIFICATE-----
	MIIEDzCCAvegAwIBAgIBADANBgkqhkiG9w0BAQUFADBoMQswCQYDVQQGEwJVUzEl
	MCMGA1UEChMcU3RhcmZpZWxkIFRlY2hub2xvZ2llcywgSW5jLjEyMDAGA1UECxMp
	U3RhcmZpZWxkIENsYXNzIDIgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkwHhcNMDQw
	NjI5MTczOTE2WhcNMzQwNjI5MTczOTE2WjBoMQswCQYDVQQGEwJVUzElMCMGA1UE
	ChMcU3RhcmZpZWxkIFRlY2hub2xvZ2llcywgSW5jLjEyMDAGA1UECxMpU3RhcmZp
	ZWxkIENsYXNzIDIgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkwggEgMA0GCSqGSIb3
	DQEBAQUAA4IBDQAwggEIAoIBAQC3Msj+6XGmBIWtDBFk385N78gDGIc/oav7PKaf
	8MOh2tTYbitTkPskpD6E8J7oX+zlJ0T1KKY/e97gKvDIr1MvnsoFAZMej2YcOadN
	+lq2cwQlZut3f+dZxkqZJRRU6ybH838Z1TBwj6+wRir/resp7defqgSHo9T5iaU0
	X9tDkYI22WY8sbi5gv2cOj4QyDvvBmVmepsZGD3/cVE8MC5fvj13c7JdBmzDI1aa
	K4UmkhynArPkPw2vCHmCuDY96pzTNbO8acr1zJ3o/WSNF4Azbl5KXZnJHoe0nRrA
	1W4TNSNe35tfPe/W93bC6j67eA0cQmdrBNj41tpvi/JEoAGrAgEDo4HFMIHCMB0G
	A1UdDgQWBBS/X7fRzt0fhvRbVazc1xDCDqmI5zCBkgYDVR0jBIGKMIGHgBS/X7fR
	zt0fhvRbVazc1xDCDqmI56FspGowaaBCDEfGA1UEBhMCVVMxJTAjBgNVBAoTHFN0
	YXJmaWVsZCBUZWNobm9sb2dpZXMsIEluYy4xMjAwBgNVBAsTKVN0YXJmaWVsZCBD
	bGFzcyAyIENlcnRpZmljYXRpb24gQXV0aG9yaXR5ggEAMAwGA1UdEwQFMAMBAf8w
	DQYJKoZIhvcNAQEFBQADggEBAAWdP4id0ckaVaGsafPzWdqbAYcaT1epoXkJKtv3
	L7IezMdeatiDh6GX70k1PncGQVhiv45YuApnP+yz3SFmH8lU+nLMPUxA2IGvd56D
	eruix/U0F47ZEUD0/CwqTRV/p2JdLiXTAAsgGh1o+Re49L2L7ShZ3U0WixeDyLJl
	xy16paq8U4Zt3VekyvggQQto8PT7dL5WXXp59fkdheMtlb71cZBDzI0fmgAKhynp
	VSJYACPq4xJDKVtHCN2MQWplBqjlIapBtJUhlbl90TSrE9atvNziPTnNvT51cKEY
	WQPJIrSPnNVeKtelttQKbfi3QBFGmh95DmK/D5fs4C8xY5Z=
	-----END CERTIFICATE-----"""
	key = """-----BEGIN CERTIFICATE-----
	MIIEDzCCAvegAwIBAgIBADANBgkqhkiG9w0BAQUFADBoMQswCQYDVQQGEwJVUzEl
	MCMGA1UEChMcU3RhcmZpZWxkIFRlY2hub2xvZ2llcywgSW5jLjEyMDAGA1UECxMp
	U3RhcmZpZWxkIENsYXNzIDIgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkwHhcNMDQw
	NjI5MTczOTE2WhcNMzQwNjI5MTczOTE2WjBoMQswCQYDVQQGEwJVUzElMCMGA1UE
	ChMcU3RhcmZpZWxkIFRlY2hub2xvZ2llcywgSW5jLjEyMDAGA1UECxMpU3RhcmZp
	ZWxkIENsYXNzIDIgQ2VydGlmaWNhdGlvbiBBdXRob3JpdHkwggEgMA0GCSqGSIb3
	DQEBAQUAA4IBDQAwggEIAoIBAQC3Msj+6XGmBIWtDBFk385N78gDGIc/oav7PKaf
	8MOh2tTYbitTkPskpD6E8J7oX+zlJ0T1KKY/e97gKvDIr1MvnsoFAZMej2YcOadN
	+lq2cwQlZut3f+dZxkqZJRRU6ybH838Z1TBwj6+wRir/resp7defqgSHo9T5iaU0
	X9tDkYI22WY8sbi5gv2cOj4QyDvvBmVmepsZGD3/cVE8MC5fvj13c7JdBmzDI1aa
	K4UmkhynArPkPw2vCHmCuDY96pzTNbO8acr1zJ3o/WSNF4Azbl5KXZnJHoe0nRrA
	1W4TNSNe35tfPe/W93bC6j67eA0cQmdrBNj41tpvi/JEoAGrAgEDo4HFMIHCMB0G
	A1UdDgQWBBS/X7fRzt0fhvRbVazc1xDCDqmI5zCBkgYDVR0jBIGKMIGHgBS/X7fR
	zt0fhvRbVazc1xDCDqmI56FspGowaaBCDEfGA1UEBhMCVVMxJTAjBgNVBAoTHFN0
	YXJmaWVsZCBUZWNobm9sb2dpZXMsIEluYy4xMjAwBgNVBAsTKVN0YXJmaWVsZCBD
	bGFzcyAyIENlcnRpZmljYXRpb24gQXV0aG9yaXR5ggEAMAwGA1UdEwQFMAMBAf8w
	DQYJKoZIhvcNAQEFBQADggEBAAWdP4id0ckaVaGsafPzWdqbAYcaT1epoXkJKtv3
	L7IezMdeatiDh6GX70k1PncGQVhiv45YuApnP+yz3SFmH8lU+nLMPUxA2IGvd56D
	eruix/U0F47ZEUD0/CwqTRV/p2JdLiXTAAsgGh1o+Re49L2L7ShZ3U0WixeDyLJl
	xy16paq8U4Zt3VekyvggQQto8PT7dL5WXXp59fkdheMtlb71cZBDzI0fmgAKhynp
	VSJYACPq4xJDKVtHCN2MQWplBqjlIapBtJUhlbl90TSrE9atvNziPTnNvT51cKEY
	WQPJIrSPnNVeKtelttQKbfi3QBFGmh95DmK/D5fs4C8xY5Z=
	-----END CERTIFICATE-----"""`
	testfile = `./testfiles/ssh`
)

var sshConfig = sshutils.SSHConfig{
	SshUser:    "ubuntu",
	SshPort:    "22",
	SshKeyFile: testfile,
	HostIP:     ValidIP,
	Timeout:    150,
}

func TestIsRemotePath(t *testing.T) {
	c := certRotateFlow{fileUtils: mockFS()}

	type testCaseInfo struct {
		testCaseDescription string
		input               string
		expected            bool
	}

	testCases := []testCaseInfo{
		{
			testCaseDescription: "Valid remote path",
			input:               RemoteFilePath,
			expected:            true,
		},
		{
			testCaseDescription: "Local path instead of remote path",
			input:               LocalFilePath,
			expected:            false,
		},
		{
			testCaseDescription: "Invalid remote path 1",
			input:               "/home/ec2-user/certs/public.pem198.51.100.0",
			expected:            false,
		},
		{
			testCaseDescription: "Invalid remote path 2",
			input:               "198.51.100.0/home/ec2-user/certs/public.pem",
			expected:            false,
		},
		{
			testCaseDescription: "Invalid remote path 3",
			input:               "\n   198.51.100.0:/home/ec2-user/certs/public.pem",
			expected:            false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testCaseDescription, func(t *testing.T) {
			actual := c.IsRemotePath(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestGetIPV4(t *testing.T) {
	c := certRotateFlow{fileUtils: mockFS()}

	type testCaseInfo struct {
		testCaseDescription string
		input               string
		expected            string
	}

	testCases := []testCaseInfo{
		{
			testCaseDescription: "Valid IP V4",
			input:               RemoteFilePath,
			expected:            ValidIP,
		},
		{
			testCaseDescription: "Valid IP V4 but invalid remote path 1",
			input:               "/home/ec2-user/certs/public.pem:127.0.0.1",
			expected:            "127.0.0.1",
		},
		{
			testCaseDescription: "Valid IP V4 but invalid remote path 2",
			input:               "/home/ec2-user/:0.0.0.0:/certs/public.pem",
			expected:            "0.0.0.0",
		},
		{
			testCaseDescription: "Invalid IP v4 and valid path",
			input:               "256.256.256.256:/home/ec2-user/certs/public.pem",
			expected:            "",
		},
		{
			testCaseDescription: "Invalid IP v4 and invalid path",
			input:               "/home/ec2-user/certs/public.pem:1.2.3",
			expected:            "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testCaseDescription, func(t *testing.T) {
			actual := c.GetIPV4(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestGetRemoteFileDetails(t *testing.T) {
	c := certRotateFlow{fileUtils: mockFS()}

	type testCaseInfo struct {
		testCaseDescription    string
		input                  string
		isError                bool
		expectedErrorMessage   string
		expectedRemoteFilePath string
		expectedFileName       string
		expectedHostIP         string
		clusterIPs             []string
	}

	testCases := []testCaseInfo{
		{
			testCaseDescription:    "Valid remote path",
			input:                  RemoteFilePath,
			isError:                false,
			expectedErrorMessage:   "",
			expectedRemoteFilePath: LocalFilePath,
			expectedFileName:       "public.pem",
			expectedHostIP:         ValidIP,
			clusterIPs:             []string{ValidIP},
		},
		{
			testCaseDescription:    "Invalid Remote Path - Local path",
			input:                  LocalFilePath,
			isError:                true,
			expectedErrorMessage:   " is not a valid IPv4 address",
			expectedRemoteFilePath: "",
			expectedFileName:       "",
			expectedHostIP:         "",
			clusterIPs:             []string{ValidIP},
		},
		{
			testCaseDescription:    "Invalid Remote Path - Colon missing",
			input:                  "198.51.100.0/home/ec2-user/certs/public.pem",
			isError:                true,
			expectedErrorMessage:   "Invalid remote path: 198.51.100.0/home/ec2-user/certs/public.pem",
			expectedRemoteFilePath: "",
			expectedFileName:       "",
			expectedHostIP:         "",
			clusterIPs:             []string{ValidIP},
		},
		{
			testCaseDescription:    "Invalid Remote Path - No filename",
			input:                  "198.51.100.0:/home/ec2-user/certs/public/",
			isError:                false,
			expectedErrorMessage:   "",
			expectedRemoteFilePath: "/home/ec2-user/certs/public",
			expectedFileName:       "public",
			expectedHostIP:         ValidIP,
			clusterIPs:             []string{ValidIP},
		},
		{
			testCaseDescription:    "Invalid Remote Path - Reverse",
			input:                  "/home/ec2-user/certs/public/:198.51.100.0",
			isError:                false,
			expectedErrorMessage:   "",
			expectedRemoteFilePath: ValidIP,
			expectedFileName:       ValidIP,
			expectedHostIP:         ValidIP,
			clusterIPs:             []string{ValidIP},
		},
		{
			testCaseDescription:    "Invalid Remote Path - Empty Path",
			input:                  ValidIP + ":",
			isError:                true,
			expectedErrorMessage:   "Invalid remote path: " + ValidIP + ":",
			expectedRemoteFilePath: "",
			expectedFileName:       "",
			expectedHostIP:         "",
			clusterIPs:             []string{ValidIP},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testCaseDescription, func(t *testing.T) {
			infra := &AutomateHAInfraDetails{}
			infra.Outputs.AutomatePrivateIps.Value = tc.clusterIPs
			remoteFilePathRes, fileNameRes, hostIPRes, err := c.GetRemoteFileDetails(tc.input, infra)
			if tc.isError {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedErrorMessage, err.Error())
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.expectedRemoteFilePath, remoteFilePathRes)
			assert.Equal(t, tc.expectedFileName, fileNameRes)
			assert.Equal(t, tc.expectedHostIP, hostIPRes)
		})
	}
}

func TestGetCerts(t *testing.T) {
	c := certRotateFlow{fileUtils: mockFS()}

	infra := &AutomateHAInfraDetails{}

	type testCaseInfo struct {
		testCaseDescription string
		flagsObj            certRotateFlags
		rootCaWant          string
		publicCertWant      string
		privateCertWant     string
		adminCertWant       string
		adminKeyWant        string
		isError             bool
	}

	testCases := []testCaseInfo{
		{
			testCaseDescription: "All paths given and flag is automate service",
			flagsObj: certRotateFlags{
				automate:        true,
				privateCertPath: ValidCertPath,
				publicCertPath:  ValidCertPath,
			},
			rootCaWant:      "",
			publicCertWant:  FileContent,
			privateCertWant: FileContent,
			adminCertWant:   "",
			adminKeyWant:    "",
			isError:         false,
		},
		{
			testCaseDescription: "All paths given except root-ca flag is automate service and node flag given",
			flagsObj: certRotateFlags{
				automate:        true,
				privateCertPath: ValidCertPath,
				publicCertPath:  ValidCertPath,
				node:            "ip-given",
			},
			rootCaWant:      "",
			publicCertWant:  FileContent,
			privateCertWant: FileContent,
			adminCertWant:   "",
			adminKeyWant:    "",
			isError:         false,
		},
		{
			testCaseDescription: "All paths given and flag is opensearch service",
			flagsObj: certRotateFlags{
				opensearch:      true,
				privateCertPath: ValidCertPath,
				publicCertPath:  ValidCertPath,
				rootCAPath:      ValidCertPath,
				adminCertPath:   ValidCertPath,
				adminKeyPath:    ValidCertPath,
			},
			rootCaWant:      FileContent,
			publicCertWant:  FileContent,
			privateCertWant: FileContent,
			adminCertWant:   FileContent,
			adminKeyWant:    FileContent,
			isError:         false,
		},
		{
			testCaseDescription: "All paths empty and flag for automate service",
			flagsObj:            certRotateFlags{},
			rootCaWant:          "",
			publicCertWant:      "",
			privateCertWant:     "",
			adminCertWant:       "",
			adminKeyWant:        "",
			isError:             true,
		},
		{
			testCaseDescription: "some invalid paths given and flag is automate service",
			flagsObj: certRotateFlags{
				automate:        true,
				privateCertPath: ValidCertPath,
				publicCertPath:  "./xyx-cert.go",
				rootCAPath:      ValidCertPath,
			},
			rootCaWant:      "",
			publicCertWant:  "",
			privateCertWant: "",
			adminCertWant:   "",
			adminKeyWant:    "",
			isError:         true,
		},
		{
			testCaseDescription: "All paths given but invalid (file not exist in (f.s)and flag is automate service",
			flagsObj: certRotateFlags{
				automate:        true,
				privateCertPath: "./xyz.go",
				publicCertPath:  "./xyz.go",
				rootCAPath:      "./xyx.go",
			},
			rootCaWant:      "",
			publicCertWant:  "",
			privateCertWant: "",
			adminCertWant:   "",
			adminKeyWant:    "",
			isError:         true,
		},
		{
			testCaseDescription: "All paths given except root-ca and flag is automate service",
			flagsObj: certRotateFlags{
				automate:        true,
				privateCertPath: ValidCertPath,
				publicCertPath:  ValidCertPath,
			},
			rootCaWant:      "",
			publicCertWant:  FileContent,
			privateCertWant: FileContent,
			adminCertWant:   "",
			adminKeyWant:    "",
			isError:         false,
		},
		{
			testCaseDescription: "Some mandatory path not given and flag is opensearch service",
			flagsObj: certRotateFlags{
				opensearch:      true,
				privateCertPath: ValidCertPath,
				publicCertPath:  ValidCertPath,
				rootCAPath:      ValidCertPath,
				adminCertPath:   ValidCertPath,
				adminKeyPath:    "",
			},
			rootCaWant:      "",
			publicCertWant:  "",
			privateCertWant: "",
			adminCertWant:   "",
			adminKeyWant:    "",
			isError:         true,
		},
		{
			testCaseDescription: "Invalid adminCert path and flag is opensearch service",
			flagsObj: certRotateFlags{
				opensearch:      true,
				privateCertPath: ValidCertPath,
				publicCertPath:  ValidCertPath,
				rootCAPath:      ValidCertPath,
				adminCertPath:   "./xyz-cert.go",
				adminKeyPath:    ValidCertPath,
			},
			rootCaWant:      "",
			publicCertWant:  "",
			privateCertWant: "",
			adminCertWant:   "",
			adminKeyWant:    "",
			isError:         true,
		},
		{
			testCaseDescription: "Invalid adminKey path and flag is opensearch service",
			flagsObj: certRotateFlags{
				opensearch:      true,
				privateCertPath: ValidCertPath,
				publicCertPath:  ValidCertPath,
				rootCAPath:      ValidCertPath,
				adminCertPath:   ValidCertPath,
				adminKeyPath:    "./xyz-cert.go",
			},
			rootCaWant:      "",
			publicCertWant:  "",
			privateCertWant: "",
			adminCertWant:   "",
			adminKeyWant:    "",
			isError:         true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testCaseDescription, func(t *testing.T) {
			certsGot, err := c.getCerts(infra, &tc.flagsObj)
			if tc.isError {
				assert.Error(t, err)
				assert.Nil(t, certsGot)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, certsGot)
			}
			if certsGot != nil {
				assert.Equal(t, tc.rootCaWant, certsGot.rootCA)
				assert.Equal(t, tc.publicCertWant, certsGot.publicCert)
				assert.Equal(t, tc.privateCertWant, certsGot.privateCert)
				assert.Equal(t, tc.adminCertWant, certsGot.adminCert)
				assert.Equal(t, tc.adminKeyWant, certsGot.adminKey)
			}
		})
	}
}

func TestCompareCurrentCertsWithNewCerts(t *testing.T) {
	c, infra := getMockCertRotateFlowAndInfra()
	type testCaseInfo struct {
		testCaseDescription string
		remoteService       string
		newCerts            *certificates
		flagsObj            *certRotateFlags
		currentCertsInfo    *certShowCertificates
		skipIpsList         []string
	}

	testCases := []testCaseInfo{
		{
			testCaseDescription: "Automate new certs are equal | No node flag",
			remoteService:       AUTOMATE,
			newCerts: &certificates{
				publicCert:  FileContent,
				privateCert: FileContent,
				rootCA:      FileContent,
			},
			flagsObj: &certRotateFlags{},
			currentCertsInfo: &certShowCertificates{
				AutomateRootCert: FileContent,
				AutomateCertsByIP: []CertByIP{
					{
						IP:         ValidIP,
						PublicKey:  FileContent,
						PrivateKey: FileContent,
					},
					{
						IP:         ValidIP1,
						PublicKey:  FileContent,
						PrivateKey: FileContent,
					},
				},
			},
			skipIpsList: infra.Outputs.AutomatePrivateIps.Value,
		},
		{
			testCaseDescription: "Automate new certs are different | No node flag",
			remoteService:       AUTOMATE,
			newCerts: &certificates{
				publicCert:  FileContent + "a",
				privateCert: FileContent,
				rootCA:      FileContent,
			},
			flagsObj: &certRotateFlags{},
			currentCertsInfo: &certShowCertificates{
				AutomateRootCert: FileContent,
				AutomateCertsByIP: []CertByIP{
					{
						IP:         ValidIP,
						PublicKey:  FileContent,
						PrivateKey: FileContent,
					},
					{
						IP:         ValidIP1,
						PublicKey:  FileContent,
						PrivateKey: FileContent,
					},
				},
			},
			skipIpsList: []string{},
		},
		{
			testCaseDescription: "Automate new certs are eqaul | Node flag",
			remoteService:       AUTOMATE,
			newCerts: &certificates{
				publicCert:  FileContent,
				privateCert: FileContent,
			},
			flagsObj: &certRotateFlags{
				node: ValidIP1,
			},
			currentCertsInfo: &certShowCertificates{
				AutomateCertsByIP: []CertByIP{
					{
						IP:         ValidIP,
						PublicKey:  FileContent,
						PrivateKey: FileContent,
					},
					{
						IP:         ValidIP1,
						PublicKey:  FileContent,
						PrivateKey: FileContent,
					},
				},
			},
			skipIpsList: []string{ValidIP1},
		},
		{
			testCaseDescription: "Chef-server new certs are equal | No node flag",
			remoteService:       CHEF_SERVER,
			newCerts: &certificates{
				publicCert:  FileContent,
				privateCert: FileContent,
			},
			flagsObj: &certRotateFlags{},
			currentCertsInfo: &certShowCertificates{
				ChefServerCertsByIP: []CertByIP{
					{
						IP:         ValidIP2,
						PublicKey:  FileContent,
						PrivateKey: FileContent,
					},
					{
						IP:         ValidIP3,
						PublicKey:  FileContent,
						PrivateKey: FileContent,
					},
				},
			},
			skipIpsList: infra.Outputs.ChefServerPrivateIps.Value,
		},
		{
			testCaseDescription: "Chef-server new certs are different | No node flag",
			remoteService:       CHEF_SERVER,
			newCerts: &certificates{
				publicCert:  FileContent + "a",
				privateCert: FileContent,
			},
			flagsObj: &certRotateFlags{},
			currentCertsInfo: &certShowCertificates{
				ChefServerCertsByIP: []CertByIP{
					{
						IP:         ValidIP2,
						PublicKey:  FileContent,
						PrivateKey: FileContent,
					},
					{
						IP:         ValidIP3,
						PublicKey:  FileContent,
						PrivateKey: FileContent,
					},
				},
			},
			skipIpsList: []string{},
		},
		{
			testCaseDescription: "Opensearch new certs are equals | No node flag",
			remoteService:       OPENSEARCH,
			newCerts: &certificates{
				publicCert:  FileContent,
				privateCert: FileContent,
				rootCA:      FileContent,
				adminCert:   FileContent,
				adminKey:    FileContent,
			},
			flagsObj: &certRotateFlags{},
			currentCertsInfo: &certShowCertificates{
				OpensearchAdminCert: FileContent,
				OpensearchAdminKey:  FileContent,
				OpensearchRootCert:  FileContent,
				OpensearchCertsByIP: []CertByIP{
					{
						IP:         ValidIP4,
						PublicKey:  FileContent,
						PrivateKey: FileContent,
					},
					{
						IP:         ValidIP5,
						PublicKey:  FileContent,
						PrivateKey: FileContent,
					},
					{
						IP:         ValidIP6,
						PublicKey:  FileContent,
						PrivateKey: FileContent,
					},
				},
			},
			skipIpsList: infra.Outputs.OpensearchPrivateIps.Value,
		},
		{
			testCaseDescription: "Opensearch new certs are different | No node flag",
			remoteService:       OPENSEARCH,
			newCerts: &certificates{
				publicCert:  FileContent + "a",
				privateCert: FileContent,
				rootCA:      FileContent,
				adminCert:   FileContent,
				adminKey:    FileContent,
			},
			flagsObj: &certRotateFlags{},
			currentCertsInfo: &certShowCertificates{
				OpensearchAdminCert: FileContent,
				OpensearchAdminKey:  FileContent,
				OpensearchRootCert:  FileContent,
				OpensearchCertsByIP: []CertByIP{
					{
						IP:         ValidIP4,
						PublicKey:  FileContent,
						PrivateKey: FileContent,
					},
					{
						IP:         ValidIP5,
						PublicKey:  FileContent,
						PrivateKey: FileContent,
					},
					{
						IP:         ValidIP6,
						PublicKey:  FileContent,
						PrivateKey: FileContent,
					},
				},
			},
			skipIpsList: []string{},
		},
		{
			testCaseDescription: "Postgresql new certs are eqaul | No node flag",
			remoteService:       POSTGRESQL,
			newCerts: &certificates{
				publicCert:  FileContent,
				privateCert: FileContent,
				rootCA:      FileContent,
			},
			flagsObj: &certRotateFlags{},
			currentCertsInfo: &certShowCertificates{
				PostgresqlRootCert: FileContent,
				PostgresqlCertsByIP: []CertByIP{
					{
						IP:         ValidIP7,
						PublicKey:  FileContent,
						PrivateKey: FileContent,
					},
					{
						IP:         ValidIP8,
						PublicKey:  FileContent,
						PrivateKey: FileContent,
					},
					{
						IP:         ValidIP9,
						PublicKey:  FileContent,
						PrivateKey: FileContent,
					},
				},
			},
			skipIpsList: infra.Outputs.PostgresqlPrivateIps.Value,
		},
		{
			testCaseDescription: "Postgresql new certs are different | No node flag",
			remoteService:       POSTGRESQL,
			newCerts: &certificates{
				publicCert:  FileContent,
				privateCert: FileContent + "a",
				rootCA:      FileContent,
			},
			flagsObj: &certRotateFlags{},
			currentCertsInfo: &certShowCertificates{
				PostgresqlRootCert: FileContent,
				PostgresqlCertsByIP: []CertByIP{
					{
						IP:         ValidIP7,
						PublicKey:  FileContent,
						PrivateKey: FileContent,
					},
					{
						IP:         ValidIP8,
						PublicKey:  FileContent,
						PrivateKey: FileContent,
					},
					{
						IP:         ValidIP9,
						PublicKey:  FileContent,
						PrivateKey: FileContent,
					},
				},
			},
			skipIpsList: []string{},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testCaseDescription, func(t *testing.T) {
			skipIpsListGot := c.compareCurrentCertsWithNewCerts(testCase.remoteService, testCase.newCerts, testCase.flagsObj, testCase.currentCertsInfo)
			assert.Equal(t, testCase.skipIpsList, skipIpsListGot)
		})
	}
}

func TestGetFrontEndIpsForSkippingCnAndRootCaPatching(t *testing.T) {
	c, infra := getMockCertRotateFlowAndInfra()
	nodesDn, _ := getDistinguishedNameFromKey(FileContent)
	newCn := nodesDn.CommonName

	type testCaseInfo struct {
		testCaseDescription string
		newRootCA           string
		newCn               string
		oldCn               string
		oldRootCA           string
		node                string
		infra               *AutomateHAInfraDetails
		isIpSkip            bool
	}

	testCases := []testCaseInfo{
		{
			testCaseDescription: "Comparing old cn nd old rootCA | same cn and root-ca",
			newRootCA:           FileContent,
			oldRootCA:           FileContent,
			newCn:               newCn,
			oldCn:               newCn,
			infra:               infra,
			isIpSkip:            true,
		},
		{
			testCaseDescription: "Comparing old cn nd old rootCA | same root-ca different cn",
			newRootCA:           FileContent,
			oldRootCA:           FileContent,
			newCn:               newCn + "n",
			oldCn:               "",
			infra:               infra,
			isIpSkip:            false,
		},
		{
			testCaseDescription: "Comparing old cn nd old rootCA | different root-ca",
			newRootCA:           FileContent + "a",
			oldRootCA:           FileContent,
			newCn:               newCn,
			oldCn:               newCn,
			infra:               infra,
			isIpSkip:            false,
		},
		{
			testCaseDescription: "Comparing old cn nd old rootCA | different root-ca same cn",
			newRootCA:           FileContent + "a",
			oldRootCA:           FileContent,
			newCn:               newCn,
			oldCn:               newCn,
			infra:               infra,
			isIpSkip:            false,
		},
		{
			testCaseDescription: "Comparing old cn | different cn with node flag",
			newCn:               newCn + "a",
			oldCn:               newCn,
			infra:               infra,
			node:                ValidIP,
			isIpSkip:            false,
		},
		{
			testCaseDescription: "Comparing old cn | same cn with node flag",
			newCn:               newCn,
			oldCn:               newCn,
			infra:               infra,
			node:                ValidIP,
			isIpSkip:            true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testCaseDescription, func(t *testing.T) {
			isIpSkipped := c.getFrontEndIpsForSkippingCnAndRootCaPatching(testCase.newRootCA, testCase.newCn, testCase.oldCn, testCase.oldRootCA, testCase.node)
			assert.Equal(t, testCase.isIpSkip, isIpSkipped)
		})
	}
}

func TestGetFilteredIps(t *testing.T) {
	c, infra := getMockCertRotateFlowAndInfra()

	type testCaseInfo struct {
		serviceIps          []string
		skipIpsList         []string
		filteredIpsExpected []string
	}

	testCases := []testCaseInfo{
		{
			serviceIps:          infra.Outputs.OpensearchPrivateIps.Value,
			skipIpsList:         []string{ValidIP5},
			filteredIpsExpected: []string{ValidIP4, ValidIP6},
		},
		{
			serviceIps:          infra.Outputs.OpensearchPrivateIps.Value,
			skipIpsList:         []string{},
			filteredIpsExpected: infra.Outputs.OpensearchPrivateIps.Value,
		},
		{
			serviceIps:          infra.Outputs.OpensearchPrivateIps.Value,
			skipIpsList:         infra.Outputs.OpensearchPrivateIps.Value,
			filteredIpsExpected: []string{},
		},
	}

	for _, testCase := range testCases {
		t.Run("TestingGetFilteredIps", func(t *testing.T) {
			filteredIpsGot := c.getFilteredIps(testCase.serviceIps, testCase.skipIpsList)
			assert.Equal(t, testCase.filteredIpsExpected, filteredIpsGot)
		})
	}
}

func TestGetFrontIpsToSkipRootCAandCNPatchingForOs(t *testing.T) {
	c, infra := getMockCertRotateFlowAndInfra()
	type testCaseInfo struct {
		testCaseDescription string
		automatesConfig     *deployment.AutomateConfig
		newRootCA           string
		oldRootCA           string
		newCn               string
		oldCn               string
		node                string
		infra               *AutomateHAInfraDetails
		skipIpsList         []string
	}

	testCases := []testCaseInfo{
		{
			testCaseDescription: "Opensearch root-ca and cn patching | same rootca and CN ",
			automatesConfig: &deployment.AutomateConfig{
				Global: &shared.GlobalConfig{
					V1: &shared.V1{
						External: &shared.External{
							Opensearch: &shared.External_Opensearch{
								Ssl: &shared.External_Opensearch_SSL{
									ServerName: &wrapperspb.StringValue{
										Value: "chefnode",
									},
									RootCert: &wrapperspb.StringValue{
										Value: FileContent,
									},
								},
							},
						},
					},
				},
			},
			oldRootCA:   FileContent,
			newRootCA:   FileContent,
			newCn:       "chefnode",
			oldCn:       "chefnode",
			infra:       infra,
			skipIpsList: []string{ValidIP},
		},
		{
			testCaseDescription: "Opensearch root-ca and cn patching | diff rootca and CN",
			automatesConfig: &deployment.AutomateConfig{
				Global: &shared.GlobalConfig{
					V1: &shared.V1{
						External: &shared.External{Opensearch: &shared.External_Opensearch{
							Ssl: &shared.External_Opensearch_SSL{
								ServerName: &wrapperspb.StringValue{
									Value: "chefnode",
								},
								RootCert: &wrapperspb.StringValue{
									Value: FileContent,
								},
							},
						},
						},
					},
				},
			},
			oldRootCA:   FileContent,
			newRootCA:   FileContent + "a",
			newCn:       "chefnode1",
			oldCn:       "chefnode",
			infra:       infra,
			skipIpsList: []string{},
		},
		{
			testCaseDescription: "Opensearch root-ca and cn patching | same rootca and diff CN",
			automatesConfig: &deployment.AutomateConfig{
				Global: &shared.GlobalConfig{
					V1: &shared.V1{
						External: &shared.External{Opensearch: &shared.External_Opensearch{
							Ssl: &shared.External_Opensearch_SSL{
								ServerName: &wrapperspb.StringValue{
									Value: "chefnode",
								},
								RootCert: &wrapperspb.StringValue{
									Value: FileContent,
								},
							},
						},
						},
					},
				},
			},
			oldRootCA:   FileContent,
			newRootCA:   FileContent,
			newCn:       "chefnode1",
			oldCn:       "chefnode",
			infra:       infra,
			skipIpsList: []string{},
		},
		{
			testCaseDescription: "Opensearch root-ca and cn patching | diff rootca and same CN",
			automatesConfig: &deployment.AutomateConfig{
				Global: &shared.GlobalConfig{
					V1: &shared.V1{
						External: &shared.External{Opensearch: &shared.External_Opensearch{
							Ssl: &shared.External_Opensearch_SSL{
								ServerName: &wrapperspb.StringValue{
									Value: "chefnode",
								},
								RootCert: &wrapperspb.StringValue{
									Value: FileContent,
								},
							},
						},
						},
					},
				},
			},
			oldRootCA:   FileContent,
			newRootCA:   FileContent + "a",
			newCn:       "chefnode",
			oldCn:       "chefnode",
			infra:       infra,
			skipIpsList: []string{},
		},
		{
			testCaseDescription: "Opensearch root-ca and cn patching | diff CN with node flag",
			automatesConfig: &deployment.AutomateConfig{
				Global: &shared.GlobalConfig{
					V1: &shared.V1{
						External: &shared.External{Opensearch: &shared.External_Opensearch{
							Ssl: &shared.External_Opensearch_SSL{
								ServerName: &wrapperspb.StringValue{
									Value: "chefnode",
								},
							},
						},
						},
					},
				},
			},
			newCn:       "chefnode1",
			oldCn:       "chefnode",
			node:        ValidIP,
			infra:       infra,
			skipIpsList: []string{},
		},
		{
			testCaseDescription: "Opensearch root-ca and cn patching | same CN with node flag",
			automatesConfig: &deployment.AutomateConfig{
				Global: &shared.GlobalConfig{
					V1: &shared.V1{
						External: &shared.External{Opensearch: &shared.External_Opensearch{
							Ssl: &shared.External_Opensearch_SSL{
								ServerName: &wrapperspb.StringValue{
									Value: "chefnode",
								},
							},
						},
						},
					},
				},
			},
			newCn:       "chefnode",
			oldCn:       "chefnode",
			node:        ValidIP,
			infra:       infra,
			skipIpsList: []string{ValidIP},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testCaseDescription, func(t *testing.T) {
			configMap := map[string]*deployment.AutomateConfig{
				ValidIP: testCase.automatesConfig,
			}
			skipIpsListGot := c.getFrontIpsToSkipRootCAandCNPatchingForOs(configMap, testCase.newRootCA, testCase.newCn, testCase.node, infra)
			assert.Equal(t, testCase.skipIpsList, skipIpsListGot)

		})
	}
}

func TestGetFrontendIPsToSkipRootCAPatchingForPg(t *testing.T) {
	c, infra := getMockCertRotateFlowAndInfra()
	type testCaseInfo struct {
		testCaseDescription string
		automatesConfig     *deployment.AutomateConfig
		newRootCA           string
		oldRootCA           string
		infra               *AutomateHAInfraDetails
		skipIpsList         []string
	}

	testCases := []testCaseInfo{
		{
			testCaseDescription: "Postgresql root-ca comparing | same rootca ",
			automatesConfig: &deployment.AutomateConfig{
				Global: &shared.GlobalConfig{
					V1: &shared.V1{
						External: &shared.External{
							Postgresql: &shared.External_Postgresql{
								Ssl: &shared.External_Postgresql_SSL{
									RootCert: &wrapperspb.StringValue{
										Value: FileContent,
									},
								},
							},
						},
					},
				},
			},
			oldRootCA:   FileContent,
			newRootCA:   FileContent,
			infra:       infra,
			skipIpsList: []string{ValidIP},
		},
		{
			testCaseDescription: "Postgresql root-ca comparing | diff rootca ",
			automatesConfig: &deployment.AutomateConfig{
				Global: &shared.GlobalConfig{
					V1: &shared.V1{
						External: &shared.External{
							Postgresql: &shared.External_Postgresql{
								Ssl: &shared.External_Postgresql_SSL{
									RootCert: &wrapperspb.StringValue{
										Value: FileContent,
									},
								},
							},
						},
					},
				},
			},
			oldRootCA:   FileContent,
			newRootCA:   FileContent + "a",
			infra:       infra,
			skipIpsList: []string{},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testCaseDescription, func(t *testing.T) {
			configMap := map[string]*deployment.AutomateConfig{
				ValidIP: testCase.automatesConfig,
			}
			skipIpsListGot := c.getFrontendIPsToSkipRootCAPatchingForPg(configMap, testCase.newRootCA, infra)
			assert.Equal(t, testCase.skipIpsList, skipIpsListGot)

		})
	}
}

func TestGetSkipIpsListForPgRootCAPatching(t *testing.T) {
	infra := NewMockInfra()
	c := NewCertRotate()

	type testCaseInfo struct {
		description     string
		sshUtil         SSHUtil
		certs           *certificates
		infra           *AutomateHAInfraDetails
		MockPullConfigs PullConfigs
		skipIpsList     []string
		ExpectedError   string
		wantError       bool
	}
	testCases := []testCaseInfo{
		{
			description: "same root ca",
			sshUtil:     GetMockSSHUtil(&SSHConfig{}, nil, completedMessage, nil, "", nil),
			certs: &certificates{
				publicCert:  FileContent,
				privateCert: FileContent,
				rootCA:      FileContent,
			},
			infra: infra,
			MockPullConfigs: &MockPullConfigs{
				pullAutomateConfigsFunc: func() (map[string]*deployment.AutomateConfig, error) {
					return map[string]*deployment.AutomateConfig{
						ValidIP7: {
							Global: &shared.GlobalConfig{
								V1: &shared.V1{
									External: &shared.External{
										Postgresql: &shared.External_Postgresql{
											Ssl: &shared.External_Postgresql_SSL{
												RootCert: &wrapperspb.StringValue{
													Value: FileContent,
												},
											},
										},
									},
								},
							},
						},
					}, nil
				},
				pullChefServerConfigsFunc: func() (map[string]*deployment.AutomateConfig, error) {
					return map[string]*deployment.AutomateConfig{
						ValidIP7: {
							Global: &shared.GlobalConfig{
								V1: &shared.V1{
									External: &shared.External{
										Postgresql: &shared.External_Postgresql{
											Ssl: &shared.External_Postgresql_SSL{
												RootCert: &wrapperspb.StringValue{
													Value: FileContent,
												},
											},
										},
									},
								},
							},
						},
					}, nil
				},
			},

			skipIpsList:   []string{ValidIP7, ValidIP7},
			ExpectedError: "",
			wantError:     false,
		},
		{
			description: "Different rootca",
			sshUtil:     GetMockSSHUtil(&SSHConfig{}, nil, completedMessage, nil, "", nil),
			certs: &certificates{
				rootCA: FileContent + "a",
			},
			infra: infra,
			MockPullConfigs: &MockPullConfigs{
				pullAutomateConfigsFunc: func() (map[string]*deployment.AutomateConfig, error) {
					return map[string]*deployment.AutomateConfig{
						ValidIP7: {
							Global: &shared.GlobalConfig{
								V1: &shared.V1{
									External: &shared.External{
										Postgresql: &shared.External_Postgresql{
											Ssl: &shared.External_Postgresql_SSL{
												RootCert: &wrapperspb.StringValue{
													Value: FileContent,
												},
											},
										},
									},
								},
							},
						},
					}, nil
				},
				pullChefServerConfigsFunc: func() (map[string]*deployment.AutomateConfig, error) {
					return map[string]*deployment.AutomateConfig{
						ValidIP7: {
							Global: &shared.GlobalConfig{
								V1: &shared.V1{
									External: &shared.External{
										Postgresql: &shared.External_Postgresql{
											Ssl: &shared.External_Postgresql_SSL{
												RootCert: &wrapperspb.StringValue{
													Value: FileContent,
												},
											},
										},
									},
								},
							},
						},
					}, nil
				},
			},

			skipIpsList:   []string{},
			ExpectedError: "",
			wantError:     false,
		},
		{
			description: "Error while pulling chefserver config",
			sshUtil:     GetMockSSHUtil(&SSHConfig{}, nil, completedMessage, nil, "", nil),
			certs: &certificates{
				publicCert:  FileContent,
				privateCert: FileContent,
				rootCA:      FileContent,
			},
			infra: infra,
			MockPullConfigs: &MockPullConfigs{
				pullAutomateConfigsFunc: func() (map[string]*deployment.AutomateConfig, error) {
					return map[string]*deployment.AutomateConfig{
						ValidIP7: {
							Global: &shared.GlobalConfig{
								V1: &shared.V1{
									External: &shared.External{
										Postgresql: &shared.External_Postgresql{
											Ssl: &shared.External_Postgresql_SSL{
												RootCert: &wrapperspb.StringValue{
													Value: FileContent,
												},
											},
										},
									},
								},
							},
						},
					}, nil
				},
				pullChefServerConfigsFunc: func() (map[string]*deployment.AutomateConfig, error) {
					return nil, errors.New("ERROR")
				},
			},

			skipIpsList:   []string{},
			ExpectedError: "ERROR",
			wantError:     true,
		},
		{
			description: "Error while pulling Automate config",
			sshUtil:     GetMockSSHUtil(&SSHConfig{}, nil, completedMessage, nil, "", nil),
			certs: &certificates{
				publicCert:  FileContent,
				privateCert: FileContent,
				rootCA:      FileContent,
			},
			infra: infra,
			MockPullConfigs: &MockPullConfigs{
				pullAutomateConfigsFunc: func() (map[string]*deployment.AutomateConfig, error) {
					return nil, errors.New("ERROR")
				},
				pullChefServerConfigsFunc: func() (map[string]*deployment.AutomateConfig, error) {
					return nil, errors.New("ERROR")
				},
			},

			skipIpsList:   []string{},
			ExpectedError: "ERROR",
			wantError:     true,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			c.pullConfigs = testCase.MockPullConfigs
			skipIpsListGot, err := c.getSkipIpsListForPgRootCAPatching(testCase.infra, testCase.sshUtil, testCase.certs)
			if !testCase.wantError {
				assert.Equal(t, testCase.skipIpsList, skipIpsListGot)
			} else {
				assert.Error(t, err)
			}

		})
	}
}

func TestGetSkipIpsListForOsRootCACNPatching(t *testing.T) {
	infra := NewMockInfra()
	c := NewCertRotate()

	type testCaseInfo struct {
		description     string
		sshUtil         SSHUtil
		certs           *certificates
		infra           *AutomateHAInfraDetails
		MockPullConfigs PullConfigs
		nodesCn         string
		flagsObj        *certRotateFlags
		skipIpsList     []string
		ExpectedError   string
		wantError       bool
	}
	testCases := []testCaseInfo{
		{
			description: "same root ca and cn in all ips",
			sshUtil:     GetMockSSHUtil(&SSHConfig{}, nil, completedMessage, nil, "", nil),
			certs: &certificates{
				rootCA: FileContent,
			},
			infra:    infra,
			nodesCn:  "chefnode",
			flagsObj: &certRotateFlags{},
			MockPullConfigs: &MockPullConfigs{
				pullAutomateConfigsFunc: func() (map[string]*deployment.AutomateConfig, error) {
					return map[string]*deployment.AutomateConfig{
						ValidIP: {
							Global: &shared.GlobalConfig{
								V1: &shared.V1{
									External: &shared.External{
										Opensearch: &shared.External_Opensearch{
											Ssl: &shared.External_Opensearch_SSL{
												ServerName: &wrapperspb.StringValue{
													Value: "chefnode",
												},
												RootCert: &wrapperspb.StringValue{
													Value: FileContent,
												},
											},
										},
									},
								},
							},
						},
					}, nil
				},
				pullChefServerConfigsFunc: func() (map[string]*deployment.AutomateConfig, error) {
					return map[string]*deployment.AutomateConfig{
						ValidIP1: {
							Global: &shared.GlobalConfig{
								V1: &shared.V1{
									External: &shared.External{
										Opensearch: &shared.External_Opensearch{
											Ssl: &shared.External_Opensearch_SSL{
												ServerName: &wrapperspb.StringValue{
													Value: "chefnode",
												},
												RootCert: &wrapperspb.StringValue{
													Value: FileContent,
												},
											},
										},
									},
								},
							},
						},
					}, nil
				},
			},

			skipIpsList:   []string{ValidIP, ValidIP1},
			ExpectedError: "",
			wantError:     false,
		},
		{
			description: "same root ca in all ips but different cn in one ip",
			sshUtil:     GetMockSSHUtil(&SSHConfig{}, nil, completedMessage, nil, "", nil),
			certs: &certificates{
				rootCA: FileContent,
			},
			infra:    infra,
			nodesCn:  "chefnode",
			flagsObj: &certRotateFlags{},
			MockPullConfigs: &MockPullConfigs{
				pullAutomateConfigsFunc: func() (map[string]*deployment.AutomateConfig, error) {
					return map[string]*deployment.AutomateConfig{
						ValidIP: {
							Global: &shared.GlobalConfig{
								V1: &shared.V1{
									External: &shared.External{
										Opensearch: &shared.External_Opensearch{
											Ssl: &shared.External_Opensearch_SSL{
												ServerName: &wrapperspb.StringValue{
													Value: "chefnode",
												},
												RootCert: &wrapperspb.StringValue{
													Value: FileContent,
												},
											},
										},
									},
								},
							},
						},
					}, nil
				},
				pullChefServerConfigsFunc: func() (map[string]*deployment.AutomateConfig, error) {
					return map[string]*deployment.AutomateConfig{
						ValidIP1: {
							Global: &shared.GlobalConfig{
								V1: &shared.V1{
									External: &shared.External{
										Opensearch: &shared.External_Opensearch{
											Ssl: &shared.External_Opensearch_SSL{
												ServerName: &wrapperspb.StringValue{
													Value: "chefnodee",
												},
												RootCert: &wrapperspb.StringValue{
													Value: FileContent,
												},
											},
										},
									},
								},
							},
						},
					}, nil
				},
			},

			skipIpsList:   []string{ValidIP},
			ExpectedError: "",
			wantError:     false,
		},
		{
			description: "diff root ca and cn in all ips",
			sshUtil:     GetMockSSHUtil(&SSHConfig{}, nil, completedMessage, nil, "", nil),
			certs: &certificates{
				rootCA: FileContent + "a",
			},
			infra:    infra,
			nodesCn:  "chefnodee",
			flagsObj: &certRotateFlags{},
			MockPullConfigs: &MockPullConfigs{
				pullAutomateConfigsFunc: func() (map[string]*deployment.AutomateConfig, error) {
					return map[string]*deployment.AutomateConfig{
						ValidIP: {
							Global: &shared.GlobalConfig{
								V1: &shared.V1{
									External: &shared.External{
										Opensearch: &shared.External_Opensearch{
											Ssl: &shared.External_Opensearch_SSL{
												ServerName: &wrapperspb.StringValue{
													Value: "chefnode",
												},
												RootCert: &wrapperspb.StringValue{
													Value: FileContent,
												},
											},
										},
									},
								},
							},
						},
					}, nil
				},
				pullChefServerConfigsFunc: func() (map[string]*deployment.AutomateConfig, error) {
					return map[string]*deployment.AutomateConfig{
						ValidIP1: {
							Global: &shared.GlobalConfig{
								V1: &shared.V1{
									External: &shared.External{
										Opensearch: &shared.External_Opensearch{
											Ssl: &shared.External_Opensearch_SSL{
												ServerName: &wrapperspb.StringValue{
													Value: "chefnode",
												},
												RootCert: &wrapperspb.StringValue{
													Value: FileContent,
												},
											},
										},
									},
								},
							},
						},
					}, nil
				},
			},

			skipIpsList:   []string{},
			ExpectedError: "",
			wantError:     false,
		},
		{
			description: "same CN with node flag",
			sshUtil:     GetMockSSHUtil(&SSHConfig{}, nil, completedMessage, nil, "", nil),
			certs: &certificates{
				rootCA: FileContent,
			},
			infra: infra,
			flagsObj: &certRotateFlags{
				node: ValidIP4,
			},
			nodesCn: "chefnode",
			MockPullConfigs: &MockPullConfigs{
				pullAutomateConfigsFunc: func() (map[string]*deployment.AutomateConfig, error) {
					return map[string]*deployment.AutomateConfig{
						ValidIP: {
							Global: &shared.GlobalConfig{
								V1: &shared.V1{
									External: &shared.External{
										Opensearch: &shared.External_Opensearch{
											Ssl: &shared.External_Opensearch_SSL{
												ServerName: &wrapperspb.StringValue{
													Value: "chefnode",
												},
											},
										},
									},
								},
							},
						},
					}, nil
				},
				pullChefServerConfigsFunc: func() (map[string]*deployment.AutomateConfig, error) {
					return map[string]*deployment.AutomateConfig{
						ValidIP1: {
							Global: &shared.GlobalConfig{
								V1: &shared.V1{
									External: &shared.External{
										Opensearch: &shared.External_Opensearch{
											Ssl: &shared.External_Opensearch_SSL{
												ServerName: &wrapperspb.StringValue{
													Value: "chefnode",
												},
											},
										},
									},
								},
							},
						},
					}, nil
				},
			},

			skipIpsList:   []string{ValidIP, ValidIP1},
			ExpectedError: "",
			wantError:     false,
		},
		{
			description: "diff CN with node flag",
			sshUtil:     GetMockSSHUtil(&SSHConfig{}, nil, completedMessage, nil, "", nil),
			certs: &certificates{
				rootCA: FileContent + "a",
			},
			infra: infra,
			flagsObj: &certRotateFlags{
				node: ValidIP4,
			},
			nodesCn: "chefnode",
			MockPullConfigs: &MockPullConfigs{
				pullAutomateConfigsFunc: func() (map[string]*deployment.AutomateConfig, error) {
					return map[string]*deployment.AutomateConfig{
						ValidIP: {
							Global: &shared.GlobalConfig{
								V1: &shared.V1{
									External: &shared.External{
										Opensearch: &shared.External_Opensearch{
											Ssl: &shared.External_Opensearch_SSL{
												ServerName: &wrapperspb.StringValue{
													Value: "chefnodee",
												},
											},
										},
									},
								},
							},
						},
					}, nil
				},
				pullChefServerConfigsFunc: func() (map[string]*deployment.AutomateConfig, error) {
					return map[string]*deployment.AutomateConfig{
						ValidIP1: {
							Global: &shared.GlobalConfig{
								V1: &shared.V1{
									External: &shared.External{
										Opensearch: &shared.External_Opensearch{
											Ssl: &shared.External_Opensearch_SSL{
												ServerName: &wrapperspb.StringValue{
													Value: "chefnodee",
												},
											},
										},
									},
								},
							},
						},
					}, nil
				},
			},

			skipIpsList:   []string{},
			ExpectedError: "",
			wantError:     false,
		},
		{
			description: "Error while fetching chefserver config",
			sshUtil:     GetMockSSHUtil(&SSHConfig{}, nil, completedMessage, nil, "", nil),
			certs: &certificates{
				publicCert:  FileContent,
				privateCert: FileContent,
				rootCA:      FileContent,
			},
			infra: infra,
			flagsObj: &certRotateFlags{
				node: ValidIP4,
			},
			MockPullConfigs: &MockPullConfigs{
				pullAutomateConfigsFunc: func() (map[string]*deployment.AutomateConfig, error) {
					return map[string]*deployment.AutomateConfig{
						ValidIP7: {
							Global: &shared.GlobalConfig{
								V1: &shared.V1{
									External: &shared.External{
										Postgresql: &shared.External_Postgresql{
											Ssl: &shared.External_Postgresql_SSL{
												RootCert: &wrapperspb.StringValue{
													Value: FileContent,
												},
											},
										},
									},
								},
							},
						},
					}, nil
				},
				pullChefServerConfigsFunc: func() (map[string]*deployment.AutomateConfig, error) {
					return nil, errors.New("ERROR")
				},
			},

			skipIpsList:   []string{},
			ExpectedError: "ERROR",
			wantError:     true,
		},
		{
			description: "Error while fetching automate config",
			sshUtil:     GetMockSSHUtil(&SSHConfig{}, nil, completedMessage, nil, "", nil),
			certs: &certificates{
				publicCert:  FileContent,
				privateCert: FileContent,
				rootCA:      FileContent,
			},
			infra: infra,
			MockPullConfigs: &MockPullConfigs{
				pullAutomateConfigsFunc: func() (map[string]*deployment.AutomateConfig, error) {
					return nil, errors.New("ERROR")
				},
				pullChefServerConfigsFunc: func() (map[string]*deployment.AutomateConfig, error) {
					return map[string]*deployment.AutomateConfig{
						ValidIP7: {
							Global: &shared.GlobalConfig{
								V1: &shared.V1{
									External: &shared.External{
										Postgresql: &shared.External_Postgresql{
											Ssl: &shared.External_Postgresql_SSL{
												RootCert: &wrapperspb.StringValue{
													Value: FileContent,
												},
											},
										},
									},
								},
							},
						},
					}, nil
				},
			},

			skipIpsList:   []string{},
			ExpectedError: "ERROR",
			wantError:     true,
		},
		{
			description: "same root ca and cn in one of the automate and chefserver ips ",
			sshUtil:     GetMockSSHUtil(&SSHConfig{}, nil, completedMessage, nil, "", nil),
			certs: &certificates{
				rootCA: FileContent,
			},
			infra:    infra,
			nodesCn:  "chefnode",
			flagsObj: &certRotateFlags{},
			MockPullConfigs: &MockPullConfigs{
				pullAutomateConfigsFunc: func() (map[string]*deployment.AutomateConfig, error) {
					return map[string]*deployment.AutomateConfig{
						ValidIP: {
							Global: &shared.GlobalConfig{
								V1: &shared.V1{
									External: &shared.External{
										Opensearch: &shared.External_Opensearch{
											Ssl: &shared.External_Opensearch_SSL{
												ServerName: &wrapperspb.StringValue{
													Value: "chefnodee",
												},
												RootCert: &wrapperspb.StringValue{
													Value: FileContent + "a",
												},
											},
										},
									},
								},
							},
						},
						ValidIP1: {
							Global: &shared.GlobalConfig{
								V1: &shared.V1{
									External: &shared.External{
										Opensearch: &shared.External_Opensearch{
											Ssl: &shared.External_Opensearch_SSL{
												ServerName: &wrapperspb.StringValue{
													Value: "chefnode",
												},
												RootCert: &wrapperspb.StringValue{
													Value: FileContent,
												},
											},
										},
									},
								},
							},
						},
					}, nil
				},
				pullChefServerConfigsFunc: func() (map[string]*deployment.AutomateConfig, error) {
					return map[string]*deployment.AutomateConfig{
						ValidIP2: {
							Global: &shared.GlobalConfig{
								V1: &shared.V1{
									External: &shared.External{
										Opensearch: &shared.External_Opensearch{
											Ssl: &shared.External_Opensearch_SSL{
												ServerName: &wrapperspb.StringValue{
													Value: "chefnode",
												},
												RootCert: &wrapperspb.StringValue{
													Value: FileContent,
												},
											},
										},
									},
								},
							},
						},
						ValidIP3: {
							Global: &shared.GlobalConfig{
								V1: &shared.V1{
									External: &shared.External{
										Opensearch: &shared.External_Opensearch{
											Ssl: &shared.External_Opensearch_SSL{
												ServerName: &wrapperspb.StringValue{
													Value: "chefnodee",
												},
												RootCert: &wrapperspb.StringValue{
													Value: FileContent + "a",
												},
											},
										},
									},
								},
							},
						},
					}, nil
				},
			},

			skipIpsList:   []string{ValidIP1, ValidIP2},
			ExpectedError: "",
			wantError:     false,
		},
		{
			description: "same rootca and cn in chefserver",
			sshUtil:     GetMockSSHUtil(&SSHConfig{}, nil, completedMessage, nil, "", nil),
			certs: &certificates{
				rootCA: FileContent,
			},
			infra:    infra,
			nodesCn:  "chefnode",
			flagsObj: &certRotateFlags{},
			MockPullConfigs: &MockPullConfigs{
				pullAutomateConfigsFunc: func() (map[string]*deployment.AutomateConfig, error) {
					return map[string]*deployment.AutomateConfig{
						ValidIP: {
							Global: &shared.GlobalConfig{
								V1: &shared.V1{
									External: &shared.External{
										Opensearch: &shared.External_Opensearch{
											Ssl: &shared.External_Opensearch_SSL{
												ServerName: &wrapperspb.StringValue{
													Value: "chefnodee",
												},
												RootCert: &wrapperspb.StringValue{
													Value: FileContent + "a",
												},
											},
										},
									},
								},
							},
						},
					}, nil
				},
				pullChefServerConfigsFunc: func() (map[string]*deployment.AutomateConfig, error) {
					return map[string]*deployment.AutomateConfig{
						ValidIP2: {
							Global: &shared.GlobalConfig{
								V1: &shared.V1{
									External: &shared.External{
										Opensearch: &shared.External_Opensearch{
											Ssl: &shared.External_Opensearch_SSL{
												ServerName: &wrapperspb.StringValue{
													Value: "chefnode",
												},
												RootCert: &wrapperspb.StringValue{
													Value: FileContent,
												},
											},
										},
									},
								},
							},
						},
						ValidIP3: {
							Global: &shared.GlobalConfig{
								V1: &shared.V1{
									External: &shared.External{
										Opensearch: &shared.External_Opensearch{
											Ssl: &shared.External_Opensearch_SSL{
												ServerName: &wrapperspb.StringValue{
													Value: "chefnode",
												},
												RootCert: &wrapperspb.StringValue{
													Value: FileContent,
												},
											},
										},
									},
								},
							},
						},
					}, nil
				},
			},

			skipIpsList:   []string{ValidIP2, ValidIP3},
			ExpectedError: "",
			wantError:     false,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			c.pullConfigs = testCase.MockPullConfigs
			skipIpsListGot, err := c.getSkipIpsListForOsRootCACNPatching(testCase.infra, testCase.sshUtil, testCase.certs, testCase.nodesCn, testCase.flagsObj)
			if !testCase.wantError {
				assert.Equal(t, testCase.skipIpsList, skipIpsListGot)
			} else {
				assert.Error(t, err)
			}

		})
	}
}

func TestPatchConfig(t *testing.T) {
	_, infra := getMockCertRotateFlowAndInfra()
	type testCaseInfo struct {
		description   string
		param         *patchFnParameters
		MockSSHUtil   sshutils.SSHUtil
		isError       bool
		ExpectedError string
	}
	testCases := []testCaseInfo{
		{
			description: "Test Case to rotate on backend",
			param: &patchFnParameters{
				sshUtil:       getMockSSHUtil(&SSHConfig{}, nil, "", nil),
				config:        TestOpensearchAdminAndRootCA,
				fileName:      "cert-rotate-os.toml",
				timestamp:     time.Now().Format("20060102150405"),
				remoteService: OPENSEARCH,
				concurrent:    false,
				infra:         infra,
				flagsObj: &certRotateFlags{
					opensearch: true,
				},
				skipIpsList: []string{},
			},
			MockSSHUtil: &sshutils.MockSSHUtilsImpl{
				CopyFileToRemoteConcurrentlyFunc: func(sshConfig sshutils.SSHConfig, srcFilePath string, destFileName string, removeFile bool, hostIPs []string) []sshutils.Result {
					return []sshutils.Result{
						{
							HostIP: "",
							Error:  nil,
							Output: "",
						},
					}
				},
			},
			isError:       false,
			ExpectedError: "",
		},
		{
			description: "Test Case to rotate on frontend",
			param: &patchFnParameters{
				sshUtil:       getMockSSHUtil(&SSHConfig{}, nil, "", nil),
				config:        TestFrontendConfig,
				fileName:      "cert-rotate-fe.toml",
				timestamp:     time.Now().Format("20060102150405"),
				remoteService: AUTOMATE,
				concurrent:    true,
				infra:         infra,
				flagsObj: &certRotateFlags{
					automate: true,
				},
				skipIpsList: []string{},
			},
			MockSSHUtil: &sshutils.MockSSHUtilsImpl{
				CopyFileToRemoteConcurrentlyFunc: func(sshConfig sshutils.SSHConfig, srcFilePath string, destFileName string, removeFile bool, hostIPs []string) []sshutils.Result {
					return []sshutils.Result{
						{
							HostIP: "",
							Error:  nil,
							Output: "",
						},
					}
				},
				ExecuteConcurrentlyFunc: func(sshConfig sshutils.SSHConfig, cmd string, hostIPs []string) []sshutils.Result {
					return []sshutils.Result{
						{
							HostIP: "",
							Error:  nil,
							Output: "",
						},
					}
				},
			},
			isError:       false,
			ExpectedError: "",
		},
		{
			description: "remote copying failed on on frontend",
			param: &patchFnParameters{
				config:        TestFrontendConfig,
				fileName:      "cert-rotate-fe.toml",
				timestamp:     time.Now().Format("20060102150405"),
				remoteService: AUTOMATE,
				concurrent:    true,
				infra:         infra,
				flagsObj: &certRotateFlags{
					automate: true,
				},
				skipIpsList: []string{},
			},
			MockSSHUtil: &sshutils.MockSSHUtilsImpl{
				CopyFileToRemoteConcurrentlyFunc: func(sshConfig sshutils.SSHConfig, srcFilePath string, destFileName string, removeFile bool, hostIPs []string) []sshutils.Result {
					return []sshutils.Result{
						{
							HostIP: "",
							Error:  errors.New("remote copying failed on node"),
							Output: "",
						},
					}
				},
				ExecuteConcurrentlyFunc: func(sshConfig sshutils.SSHConfig, cmd string, hostIPs []string) []sshutils.Result {
					return []sshutils.Result{
						{
							HostIP: "",
							Error:  nil,
							Output: "",
						},
					}
				},
			},
			isError:       true,
			ExpectedError: "remote copying failed on node",
		},
		{
			description: "remote execution failed on on frontend",
			param: &patchFnParameters{
				config:        TestFrontendConfig,
				fileName:      "cert-rotate-fe.toml",
				timestamp:     time.Now().Format("20060102150405"),
				remoteService: AUTOMATE,
				concurrent:    true,
				infra:         infra,
				flagsObj: &certRotateFlags{
					automate: true,
				},
				skipIpsList: []string{},
			},
			MockSSHUtil: &sshutils.MockSSHUtilsImpl{
				CopyFileToRemoteConcurrentlyFunc: func(sshConfig sshutils.SSHConfig, srcFilePath string, destFileName string, removeFile bool, hostIPs []string) []sshutils.Result {
					return []sshutils.Result{
						{
							HostIP: "",
							Error:  nil,
							Output: "",
						},
					}
				},
				ExecuteConcurrentlyFunc: func(sshConfig sshutils.SSHConfig, cmd string, hostIPs []string) []sshutils.Result {
					return []sshutils.Result{
						{
							HostIP: "",
							Error:  errors.New("remote execution failed"),
							Output: "",
						},
					}
				},
			},
			isError:       false,
			ExpectedError: "",
		},
		{
			description: "Deployment service error on frontend",
			param: &patchFnParameters{
				config:        TestFrontendConfig,
				fileName:      "cert-rotate-fe.toml",
				timestamp:     time.Now().Format("20060102150405"),
				remoteService: AUTOMATE,
				concurrent:    true,
				infra:         infra,
				flagsObj: &certRotateFlags{
					automate: true,
				},
				skipIpsList: []string{},
			},
			MockSSHUtil: &sshutils.MockSSHUtilsImpl{
				CopyFileToRemoteConcurrentlyFunc: func(sshConfig sshutils.SSHConfig, srcFilePath string, destFileName string, removeFile bool, hostIPs []string) []sshutils.Result {
					return []sshutils.Result{
						{
							HostIP: "",
							Error:  nil,
							Output: "",
						},
					}
				},
				ExecuteConcurrentlyFunc: func(sshConfig sshutils.SSHConfig, cmd string, hostIPs []string) []sshutils.Result {
					return []sshutils.Result{
						{
							HostIP: "",
							Error:  nil,
							Output: "DeploymentServiceCallError",
						},
					}
				},
			},
			isError:       false,
			ExpectedError: "",
		},
		{
			description: "Error occured while reading infra details on backend",
			param: &patchFnParameters{
				sshUtil:       getMockSSHUtil(&SSHConfig{}, nil, "", errors.Errorf("Error occured while reading infra details")),
				config:        TestOpensearchAdminAndRootCA,
				fileName:      "cert-rotate-os.toml",
				timestamp:     time.Now().Format("20060102150405"),
				remoteService: OPENSEARCH,
				concurrent:    false,
				infra:         infra,
				flagsObj: &certRotateFlags{
					opensearch: true,
				},
				skipIpsList: []string{},
			},
			isError:       true,
			ExpectedError: "Error occured while reading infra details",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			c := certRotateFlow{fileUtils: mockFS(),
				sshUtil: testCase.MockSSHUtil,
				writer:  getMockWriterImpl()}
			testDir := t.TempDir()
			fPath := filepath.Join(testDir, testCase.param.fileName)
			testCase.param.fileName = fPath
			output := c.patchConfig(testCase.param)
			if testCase.isError {
				assert.EqualError(t, output, testCase.ExpectedError)
			} else {
				assert.NoError(t, output)
			}
		})
	}
}

func getMockCertRotateFlowAndInfra() (*certRotateFlow, *AutomateHAInfraDetails) {
	return NewCertRotate(), NewMockInfra()
}

func NewMockInfra() *AutomateHAInfraDetails {
	infra := &AutomateHAInfraDetails{}
	infra.Outputs.AutomatePrivateIps.Value = []string{ValidIP, ValidIP1}
	infra.Outputs.ChefServerPrivateIps.Value = []string{ValidIP2, ValidIP3}
	infra.Outputs.OpensearchPrivateIps.Value = []string{ValidIP4, ValidIP5, ValidIP6}
	infra.Outputs.PostgresqlPrivateIps.Value = []string{ValidIP7, ValidIP8, ValidIP9}
	infra.Outputs.SSHUser.Value = "ubuntu"
	infra.Outputs.SSHKeyFile.Value = "new.pem"
	infra.Outputs.SSHPort.Value = "22"
	return infra
}

func NewCertRotate() *certRotateFlow {
	c := NewCertRotateFlow(mockFS(), &sshutils.MockSSHUtilsImpl{}, writer, &MockPullConfigs{})
	return c
}

func mockFS() *fileutils.MockFileSystemUtils {
	return &fileutils.MockFileSystemUtils{
		ReadFileFunc: func(filename string) ([]byte, error) {
			if _, err := os.Stat(filename); err == nil {
				// path/to/whatever exists
				return []byte(FileContent), nil
			} else if errors.Is(err, os.ErrNotExist) {
				// path/to/whatever does *not* exist
				return []byte{}, err
			} else {
				return []byte{}, err
			}
		},
	}
}
