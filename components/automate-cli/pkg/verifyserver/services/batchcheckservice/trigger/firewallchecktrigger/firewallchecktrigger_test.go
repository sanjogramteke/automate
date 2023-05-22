package firewallchecktrigger

import (
	"encoding/json"

	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/chef/automate/components/automate-cli/pkg/verifyserver/constants"
	"github.com/chef/automate/components/automate-cli/pkg/verifyserver/models"
	"github.com/chef/automate/lib/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	hardware = models.Hardware{
		AutomateNodeIps:        []string{"10.0.0.1"},
		PostgresqlNodeIps:      []string{"10.0.0.3", "10.0.0.6"},
		OpenSearchNodeIps:      []string{"10.0.0.5", "10.0.0.10"},
		ChefInfraServerNodeIps: []string{"10.0.0.7"},
	}

	nodeCert = models.NodeCert{
		IP:  "10.0.0.1",
		Key: "test-key",
	}

	nodes = []models.NodeCert{nodeCert}
)

const (
	automatepgTcpPass = `{
		"status": "SUCCESS",
		"result": {
		  "passed": true,
		  "checks": [
			{
			  "title": "Check for reachability of service at destination port from the source node",
			  "passed": true,
			  "success_msg": "The <protocol> service running at <destination_node_ip>:<destination_node_port> is reachable fro <source_node_ip>",
			  "error_msg": "",
			  "resolution_msg": ""
			}
		  ]
		}
	  }`

	automatePgFailureResponse = `
	{
			"status": "SUCCESS",
			"result": {
			  "passed": false,
			  "checks": [
				{
				  "title": "Check for reachability of service at destination port",
				  "passed": false,
				  "success_msg": "",
				  "error_msg": "The <protocol> service running at <destination_node_ip>:<destination_node_port> is not reachable from <source_ip>",
				  "resolution_msg": "Check your firewall settings to provide access to <destination_node_port> port at <destination_node_ip> from <source_node_ip>"
				}
			  ]
			}
		  }
	`
	apiTriggerAutomatePgPass = `[
		{
			"status": "SUCCESS",
			"host": "10.0.0.1",
			"node_type": "automate",
			"result": {
				"passed": true,
				"checks": [
					{
						"title": "Check for reachability of service at destination port from the source node",
						"passed": true,
						"success_msg": "The <protocol> service running at <destination_node_ip>:<destination_node_port> is reachable fro <source_node_ip>",
						"error_msg": "",
						"resolution_msg": ""
					}
				]
			}
		},
		{
			"status": "SUCCESS",
			"host": "127.0.0.1",
			"node_type": "bastion",
			"result": {
				"passed": true,
				"checks": [
					{
						"title": "Check for reachability of service at destination port from the source node",
						"passed": true,
						"success_msg": "The <protocol> service running at <destination_node_ip>:<destination_node_port> is reachable fro <source_node_ip>",
						"error_msg": "",
						"resolution_msg": ""
					}
				]
			}
		},
		{
			"status": "SUCCESS",
			"host": "127.0.0.1",
			"node_type": "bastion",
			"result": {
				"passed": true,
				"checks": [
					{
						"title": "Check for reachability of service at destination port from the source node",
						"passed": true,
						"success_msg": "The <protocol> service running at <destination_node_ip>:<destination_node_port> is reachable fro <source_node_ip>",
						"error_msg": "",
						"resolution_msg": ""
					}
				]
			}
		},
		{
			"status": "SUCCESS",
			"host": "127.0.0.1",
			"node_type": "bastion",
			"result": {
				"passed": true,
				"checks": [
					{
						"title": "Check for reachability of service at destination port from the source node",
						"passed": true,
						"success_msg": "The <protocol> service running at <destination_node_ip>:<destination_node_port> is reachable fro <source_node_ip>",
						"error_msg": "",
						"resolution_msg": ""
					}
				]
			}
		},
		{
			"status": "SUCCESS",
			"host": "127.0.0.1",
			"node_type": "bastion",
			"result": {
				"passed": true,
				"checks": [
					{
						"title": "Check for reachability of service at destination port from the source node",
						"passed": true,
						"success_msg": "The <protocol> service running at <destination_node_ip>:<destination_node_port> is reachable fro <source_node_ip>",
						"error_msg": "",
						"resolution_msg": ""
					}
				]
			}
		},
		{
			"status": "SUCCESS",
			"host": "127.0.0.1",
			"node_type": "bastion",
			"result": {
				"passed": true,
				"checks": [
					{
						"title": "Check for reachability of service at destination port from the source node",
						"passed": true,
						"success_msg": "The <protocol> service running at <destination_node_ip>:<destination_node_port> is reachable fro <source_node_ip>",
						"error_msg": "",
						"resolution_msg": ""
					}
				]
			}
		},
		{
			"status": "SUCCESS",
			"host": "127.0.0.1",
			"node_type": "bastion",
			"result": {
				"passed": true,
				"checks": [
					{
						"title": "Check for reachability of service at destination port from the source node",
						"passed": true,
						"success_msg": "The <protocol> service running at <destination_node_ip>:<destination_node_port> is reachable fro <source_node_ip>",
						"error_msg": "",
						"resolution_msg": ""
					}
				]
			}
		},
		{
			"status": "SUCCESS",
			"host": "127.0.0.1",
			"node_type": "bastion",
			"result": {
				"passed": true,
				"checks": [
					{
						"title": "Check for reachability of service at destination port from the source node",
						"passed": true,
						"success_msg": "The <protocol> service running at <destination_node_ip>:<destination_node_port> is reachable fro <source_node_ip>",
						"error_msg": "",
						"resolution_msg": ""
					}
				]
			}
		}
	]
	 `

	apiTriggerFailureResponse = `[
		{
			"status": "SUCCESS",
			"host": "10.0.0.1",
			"node_type": "automate",
			"result": {
				"passed": false,
				"checks": [
					{
						"title": "Check for reachability of service at destination port",
						"passed": false,
						"success_msg": "",
						"error_msg": "The <protocol> service running at <destination_node_ip>:<destination_node_port> is not reachable from <source_ip>",
						"resolution_msg": "Check your firewall settings to provide access to <destination_node_port> port at <destination_node_ip> from <source_node_ip>"
					}
				]
			}
		},
		{
			"status": "SUCCESS",
			"host": "127.0.0.1",
			"node_type": "bastion",
			"result": {
				"passed": false,
				"checks": [
					{
						"title": "Check for reachability of service at destination port",
						"passed": false,
						"success_msg": "",
						"error_msg": "The <protocol> service running at <destination_node_ip>:<destination_node_port> is not reachable from <source_ip>",
						"resolution_msg": "Check your firewall settings to provide access to <destination_node_port> port at <destination_node_ip> from <source_node_ip>"
					}
				]
			}
		},
		{
			"status": "SUCCESS",
			"host": "127.0.0.1",
			"node_type": "bastion",
			"result": {
				"passed": false,
				"checks": [
					{
						"title": "Check for reachability of service at destination port",
						"passed": false,
						"success_msg": "",
						"error_msg": "The <protocol> service running at <destination_node_ip>:<destination_node_port> is not reachable from <source_ip>",
						"resolution_msg": "Check your firewall settings to provide access to <destination_node_port> port at <destination_node_ip> from <source_node_ip>"
					}
				]
			}
		},
		{
			"status": "SUCCESS",
			"host": "127.0.0.1",
			"node_type": "bastion",
			"result": {
				"passed": false,
				"checks": [
					{
						"title": "Check for reachability of service at destination port",
						"passed": false,
						"success_msg": "",
						"error_msg": "The <protocol> service running at <destination_node_ip>:<destination_node_port> is not reachable from <source_ip>",
						"resolution_msg": "Check your firewall settings to provide access to <destination_node_port> port at <destination_node_ip> from <source_node_ip>"
					}
				]
			}
		},
		{
			"status": "SUCCESS",
			"host": "127.0.0.1",
			"node_type": "bastion",
			"result": {
				"passed": false,
				"checks": [
					{
						"title": "Check for reachability of service at destination port",
						"passed": false,
						"success_msg": "",
						"error_msg": "The <protocol> service running at <destination_node_ip>:<destination_node_port> is not reachable from <source_ip>",
						"resolution_msg": "Check your firewall settings to provide access to <destination_node_port> port at <destination_node_ip> from <source_node_ip>"
					}
				]
			}
		},
		{
			"status": "SUCCESS",
			"host": "127.0.0.1",
			"node_type": "bastion",
			"result": {
				"passed": false,
				"checks": [
					{
						"title": "Check for reachability of service at destination port",
						"passed": false,
						"success_msg": "",
						"error_msg": "The <protocol> service running at <destination_node_ip>:<destination_node_port> is not reachable from <source_ip>",
						"resolution_msg": "Check your firewall settings to provide access to <destination_node_port> port at <destination_node_ip> from <source_node_ip>"
					}
				]
			}
		},
		{
			"status": "SUCCESS",
			"host": "127.0.0.1",
			"node_type": "bastion",
			"result": {
				"passed": false,
				"checks": [
					{
						"title": "Check for reachability of service at destination port",
						"passed": false,
						"success_msg": "",
						"error_msg": "The <protocol> service running at <destination_node_ip>:<destination_node_port> is not reachable from <source_ip>",
						"resolution_msg": "Check your firewall settings to provide access to <destination_node_port> port at <destination_node_ip> from <source_node_ip>"
					}
				]
			}
		},
		{
			"status": "SUCCESS",
			"host": "127.0.0.1",
			"node_type": "bastion",
			"result": {
				"passed": false,
				"checks": [
					{
						"title": "Check for reachability of service at destination port",
						"passed": false,
						"success_msg": "",
						"error_msg": "The <protocol> service running at <destination_node_ip>:<destination_node_port> is not reachable from <source_ip>",
						"resolution_msg": "Check your firewall settings to provide access to <destination_node_port> port at <destination_node_ip> from <source_node_ip>"
					}
				]
			}
		}
	]`

	apiTriggerInvalidParsedResponse = `{
		"status": "SUCCESS",
		"host": "10.0.0.1",
		"node_type": "automate",
		"result": {
			"passed": true,
			"checks": [
				{
					"title": "Check for reachability of service at destination port from the source node",
					"passed": true,
					"success_msg": "The <protocol> service running at <destination_node_ip>:<destination_node_port> is reachable fro <source_node_ip>",
					"error_msg": "",
		
	 
	 `
)

func TestMakeRequests(t *testing.T) {
	// Create a sample configuration
	config := models.Config{
		Hardware: hardware,
		Certificate: models.Certificate{

			Nodes: []models.NodeCert{
				{IP: "10.0.0.1",
					Cert: "cert",
					Key:  "key",
				},
			},
			RootCert: "root_cert",
		},
	}

	// Define the expected result
	expected := []models.FirewallRequest{
		// Expected requests for Chef Automate to all the OS and PG nodes
		{
			SourceNodeIP:               "10.0.0.1",
			DestinationNodeIP:          "10.0.0.3",
			DestinationServicePort:     "7432",
			DestinationServiceProtocol: "tcp",
		},
		{
			SourceNodeIP:               "10.0.0.1",
			DestinationNodeIP:          "10.0.0.5",
			DestinationServicePort:     "7432",
			DestinationServiceProtocol: "tcp",
		},
		{
			SourceNodeIP:               "10.0.0.1",
			DestinationNodeIP:          "10.0.0.10",
			DestinationServicePort:     "9200",
			DestinationServiceProtocol: "tcp",
		},
		{
			SourceNodeIP:               "10.0.0.1",
			DestinationNodeIP:          "10.0.0.6",
			DestinationServicePort:     "9200",
			DestinationServiceProtocol: "tcp",
		},
	}

	mapRequests := make(map[string][]models.FirewallRequest)
	// Call the function
	makeRequests(config, mapRequests)
	requestsForautomate, ok := mapRequests[constants.AUTOMATE]
	assert.True(t, ok)
	assert.Equal(t, len(requestsForautomate), 4)

	requestsForchefServer, ok := mapRequests[constants.CHEF_INFRA_SERVER]
	assert.True(t, ok)
	assert.Equal(t, len(requestsForchefServer), 5)

	//As there are two postgress nodes which needs to interact with eachother
	requestsForPostgres, ok := mapRequests[constants.POSTGRESQL]
	assert.True(t, ok)
	assert.Equal(t, len(requestsForPostgres), 12)

	requestsForOpensearch, ok := mapRequests[constants.OPENSEARCH]
	assert.True(t, ok)
	assert.Equal(t, len(requestsForOpensearch), 10)

	require.Equal(t, expected[0].SourceNodeIP, requestsForautomate[0].SourceNodeIP)
	require.Equal(t, expected[0].DestinationNodeIP, requestsForautomate[0].DestinationNodeIP)
}

// Helper function to create a dummy server
func createDummyServer(t *testing.T, requiredStatusCode int, invalidParseResponse bool, isFailed bool) (*httptest.Server, string, string) {
	if requiredStatusCode == http.StatusOK {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == constants.FIREWALL_API_PATH {
				body, err := ioutil.ReadAll(r.Body)
				require.NoError(t, err)

				var req models.FirewallRequest

				err = json.Unmarshal(body, &req)
				require.NoError(t, err)
				require.NotZero(t, req)
				require.NotEmpty(t, req)
				if invalidParseResponse {
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(apiTriggerInvalidParsedResponse))

				} else if isFailed {
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(automatePgFailureResponse))
				} else {
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(automatepgTcpPass))
				}

			}
		}))

		// Extract IP and port from the server's URL
		address := server.URL[strings.Index(server.URL, "//")+2:]
		colonIndex := strings.Index(address, ":")
		ip := address[:colonIndex]
		port := address[colonIndex+1:]

		return server, ip, port
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(requiredStatusCode)
	}))

	// Extract IP and port from the server's URL
	address := server.URL[strings.Index(server.URL, "//")+2:]
	colonIndex := strings.Index(address, ":")
	ip := address[:colonIndex]
	port := address[colonIndex+1:]

	return server, ip, port
}

func TestFirewallCheck_Run(t *testing.T) {

	type args struct {
		config models.Config
	}
	tests := []struct {
		name                 string
		args                 args
		response             string
		httpStatusCode       int
		isError              bool
		invalidParseResponse bool
		isFailed             bool
	}{
		{
			name: "Automate Pg Passed",
			args: args{
				config: models.Config{
					Hardware: models.Hardware{
						AutomateNodeIps:   []string{"10.0.0.1"},
						PostgresqlNodeIps: []string{"10.0.0.3"},
					},
					Certificate: models.Certificate{
						RootCert: "test-cert",
						Nodes:    nodes,
					},
				},
			},
			response:       apiTriggerAutomatePgPass,
			httpStatusCode: http.StatusOK,
		},
		{
			name: "Automate Pg failure",
			args: args{
				config: models.Config{
					Hardware: models.Hardware{
						AutomateNodeIps:   []string{"10.0.0.1"},
						PostgresqlNodeIps: []string{"10.0.0.3"},
					},
					Certificate: models.Certificate{
						RootCert: "test-cert",
						Nodes:    nodes,
					},
				},
			},
			isFailed:       true,
			response:       apiTriggerFailureResponse,
			httpStatusCode: http.StatusOK,
		},
		{
			name: "Internal Server Error For automate and other nodes",
			args: args{
				config: models.Config{
					Hardware: models.Hardware{
						AutomateNodeIps:   []string{"10.0.0.1"},
						PostgresqlNodeIps: []string{"10.0.0.3"},
					},
					Certificate: models.Certificate{
						RootCert: "test-cert",
						Nodes:    nodes,
					},
				},
			},
			response:             "error while connecting to the endpoint, received invalid status code",
			httpStatusCode:       http.StatusInternalServerError,
			isError:              true,
			invalidParseResponse: false,
		},
		{
			name: "Error in parsing the response",
			args: args{
				config: models.Config{
					Hardware: models.Hardware{
						AutomateNodeIps:   []string{"10.0.0.1"},
						PostgresqlNodeIps: []string{"10.0.0.3"},
					},
					Certificate: models.Certificate{
						RootCert: "test-cert",
						Nodes:    nodes,
					},
				},
			},
			response:             "error while parsing the response data",
			httpStatusCode:       http.StatusOK,
			isError:              true,
			invalidParseResponse: true,
		}, {
			name: "Status Bad Request",
			args: args{
				config: models.Config{
					Hardware: models.Hardware{
						AutomateNodeIps:   []string{"10.0.0.1"},
						PostgresqlNodeIps: []string{"10.0.0.3"},
					},
					Certificate: models.Certificate{
						RootCert: "test-cert",
						Nodes:    nodes,
					},
				},
			},
			response:             "error while connecting to the endpoint, received invalid status code",
			httpStatusCode:       http.StatusBadRequest,
			isError:              true,
			invalidParseResponse: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				// Create a dummy server
				server, host, port := createDummyServer(t, tt.httpStatusCode, tt.invalidParseResponse, tt.isFailed)
				defer server.Close()

				fwc := NewFirewallCheck(logger.NewLogrusStandardLogger(), port)
				fwc.host = host
				got := fwc.Run(tt.args.config)

				if tt.isError {
					assert.NotNil(t, got[0].Error)
					assert.Contains(t, got[0].Error.Error(), tt.response)
					if tt.httpStatusCode == http.StatusBadRequest {
						assert.Equal(t, got[0].Error.Code, tt.httpStatusCode)
					} else {
						assert.Equal(t, got[0].Error.Code, http.StatusInternalServerError)
					}
					for i := range got {
						assert.NotNil(t, got[i].NodeType)
						assert.NotNil(t, got[i].Host)
					}

				} else {
					var want []models.CheckTriggerResponse
					json.Unmarshal([]byte(tt.response), &want)
					assert.Equal(t, len(got), len(want))
					for i := range want {
						//Checking each want trigger response is present in what we go in case of non-error
						assert.Contains(t, got, want[i])
					}
				}

			})

		})

	}
}

func TestTriggerFirewallAPI(t *testing.T) {
	request := models.FirewallRequest{
		SourceNodeIP:      "test",
		DestinationNodeIP: "testdest",
	}
	t.Run("Endpoint not reachable", func(t *testing.T) {
		endPoint := "http://nonexistent-api.com"
		host := "example.com"
		output := make(chan models.CheckTriggerResponse)

		// Call the function under test
		go triggerFirewallAPI(endPoint, host, constants.AUTOMATE, http.MethodPost, output, request)

		// Wait for the response
		response := <-output

		// Assert the expected error response
		require.NotNil(t, response.Error)
		assert.Equal(t, http.StatusInternalServerError, response.Error.Code)
		assert.Contains(t, response.Error.Message, `"http://nonexistent-api.com": dial tcp: lookup nonexistent-api.com`)
		require.Equal(t, "automate", response.NodeType)

	})
	t.Run("Non-OK status code", func(t *testing.T) {
		// Create a test server to mock the API endpoint
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Return a non-OK status code
			w.WriteHeader(http.StatusBadRequest)
		}))
		defer server.Close()

		output := make(chan models.CheckTriggerResponse)

		// Call the function under test
		go triggerFirewallAPI(server.URL+constants.FIREWALL_API_PATH, server.URL, constants.AUTOMATE, http.MethodPost, output, request)

		// Wait for the response
		response := <-output

		// Assert the expected error response
		require.NotNil(t, response.Error)
		assert.Equal(t, http.StatusBadRequest, response.Error.Code)
		assert.Equal(t, "error while connecting to the endpoint, received invalid status code", response.Error.Message)
	})

	t.Run("Request creation error", func(t *testing.T) {
		endPoint := "http://example.com/api/v1/checks/firewall"
		host := "example.com"
		output := make(chan models.CheckTriggerResponse)

		// Call the function under test

		go triggerFirewallAPI(endPoint, host, constants.AUTOMATE, http.MethodPost, output, request)
		// Wait for the response
		response := <-output
		// Assert the expected error response
		require.NotNil(t, response.Error)
		require.Equal(t, http.StatusNotFound, response.Error.Code)
		assert.Equal(t, "error while connecting to the endpoint, received invalid status code", response.Error.Message)
	})
}