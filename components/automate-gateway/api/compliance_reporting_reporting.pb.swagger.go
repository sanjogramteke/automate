package api

func init() {
	Swagger.Add("compliance_reporting_reporting", `{
  "swagger": "2.0",
  "info": {
    "title": "components/automate-gateway/api/compliance/reporting/reporting.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/compliance/reporting/controls": {
      "post": {
        "summary": "List controls",
        "description": "List all controls optionally using filters. Supports filtering, but not pagination or sorting.",
        "operationId": "ListControlItems",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.ControlItems"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.ControlItemRequest"
            }
          }
        ],
        "tags": [
          "Compliance Reporting Controls"
        ]
      }
    },
    "/compliance/reporting/nodes/id/{id}": {
      "get": {
        "summary": "Fetch a node",
        "description": "Fetch a specific node by id.\nDoes not support filtering, pagination or sorting.",
        "operationId": "ReadNode",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.Node"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "description": "The id of the node to fetch",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "Compliance Reporting Nodes"
        ]
      }
    },
    "/compliance/reporting/nodes/search": {
      "post": {
        "summary": "List nodes",
        "description": "List all nodes optionally using filters. Supports pagination, filtering, and sorting.\n\n| Sort paramter | Sort value |\n| --- | --- |\n| environment | environment.lower |\n| latest_report.controls.failed.critical | controls_sums.failed.critical |\n| latest_report.controls.failed.total | controls_sums.failed.total |\n| latest_report.end_time (default) | end_time |\n| latest_report.status | status |\n| name | node_name.lower |\n| platform | platform.full |\n| status | status |",
        "operationId": "ListNodes",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.Nodes"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.Query"
            }
          }
        ],
        "tags": [
          "Compliance Reporting Nodes"
        ]
      }
    },
    "/compliance/reporting/profiles": {
      "post": {
        "summary": "List profiles",
        "description": "List all profiles optionally using filters. Supports pagination, filtering, and sorting.\n\n| Sort paramter | Sort value |\n| --- | --- |\n| name | name.lower |\n| title (default) | title.lower |",
        "operationId": "ListProfiles",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.ProfileMins"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.Query"
            }
          }
        ],
        "tags": [
          "Compliance Reporting Profiles"
        ]
      }
    },
    "/compliance/reporting/report-ids": {
      "post": {
        "summary": "List report IDs",
        "description": "List all report IDs optionally using filters. Supports filtering, but not pagination or sorting.\nIncluding more than one value for ` + "`" + `control` + "`" + `, ` + "`" + `profile_id` + "`" + `, or ` + "`" + `profile_name` + "`" + ` is not allowed.\nIncluding values for both ` + "`" + `profile_id` + "`" + ` and ` + "`" + `profile_name` + "`" + ` in one request is not allowed.",
        "operationId": "ListReportIds",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.ReportIds"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.Query"
            }
          }
        ],
        "tags": [
          "Compliance Reporting Report IDs"
        ]
      }
    },
    "/compliance/reporting/reports": {
      "post": {
        "summary": "List reports",
        "description": "List all reports optionally using filters. Supports pagination, filtering, and sorting.\n\n| Sort paramter | Sort value |\n| --- | --- |\n| latest_report.controls.failed.critical | controls_sums.failed.critical |\n| latest_report.controls.failed.total | controls_sums.failed.total |\n| latest_report.end_time (default) | end_time |\n| latest_report.status | status |\n| node_name | node_name.lower |",
        "operationId": "ListReports",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.Reports"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.Query"
            }
          }
        ],
        "tags": [
          "Compliance Reporting Reports"
        ]
      }
    },
    "/compliance/reporting/reports/id/{id}": {
      "post": {
        "summary": "Fetch a report",
        "description": "Fetch a specific report by id. Supports filtering, but not pagination or sorting.\nIncluding more than one value for ` + "`" + `profile_id` + "`" + ` is not allowed.",
        "operationId": "ReadReport",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.Report"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "description": "Used by the ReadReport endpoint to specify which report to return",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.Query"
            }
          }
        ],
        "tags": [
          "Compliance Reporting Reports"
        ]
      }
    },
    "/compliance/reporting/suggestions": {
      "post": {
        "summary": "List suggestions",
        "description": "Get suggestions for compliance reporting resources based on matching text substrings.\nSupports filtering, but not pagination or sorting.\n` + "`" + `type` + "`" + ` parameter is required. It must be one of the parameters from the following table.\n\n| Suggestion type parameter | Suggestion type value |\n| --- | --- |\n| chef_server | source_fqdn |\n| chef_tags | chef_tags |\n| control | profiles.controls.title |\n| control_tag_key | profiles.controls.string_tags.key |\n| control_tag_value | profiles.controls.string_tags.values |\n| environment | environment |\n| inspec_version | version |\n| node | node_name |\n| organization | organization_name |\n| platform | platform.name |\n| platform_with_version | platform.full |\n| policy_group | policy_group |\n| policy_name | policy_name |\n| profile | profiles.title |\n| profile_with_version | profiles.full |\n| recipe | recipes |\n| role | roles |",
        "operationId": "ListSuggestions",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.Suggestions"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.SuggestionRequest"
            }
          }
        ],
        "tags": [
          "Compliance Reporting Suggestions"
        ]
      }
    },
    "/compliance/reporting/version": {
      "get": {
        "operationId": "GetVersion",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.common.version.VersionInfo"
            }
          }
        },
        "tags": [
          "hidden"
        ]
      }
    }
  },
  "definitions": {
    "chef.automate.api.common.version.VersionInfo": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "version": {
          "type": "string"
        },
        "sha": {
          "type": "string"
        },
        "built": {
          "type": "string"
        }
      }
    },
    "chef.automate.api.compliance.reporting.v1.Attribute": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "title": "The name of the attribute"
        },
        "options": {
          "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.Option",
          "title": "The options defined for the attribute"
        }
      }
    },
    "chef.automate.api.compliance.reporting.v1.Control": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "title": "The unique id of this control"
        },
        "code": {
          "type": "string",
          "title": "The full ruby code of the control defined in the profile"
        },
        "desc": {
          "type": "string",
          "title": "The full description of the control"
        },
        "impact": {
          "type": "number",
          "format": "float",
          "title": "The severity of the control"
        },
        "title": {
          "type": "string",
          "title": "The compact description of the control"
        },
        "source_location": {
          "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.SourceLocation",
          "title": "Intentionally blank"
        },
        "results": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.Result"
          },
          "title": "The results of running all tests defined in the control against the node"
        },
        "refs": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.Ref"
          },
          "title": "External supporting documents for the control"
        },
        "tags": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          },
          "title": "Metadata defined on the control in key-value format"
        }
      }
    },
    "chef.automate.api.compliance.reporting.v1.ControlItem": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "title": "The unique id of this control"
        },
        "title": {
          "type": "string",
          "title": "The compact description of the control"
        },
        "profile": {
          "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.ProfileMin",
          "title": "Intentionally blank"
        },
        "impact": {
          "type": "number",
          "format": "float",
          "title": "The severity of the control"
        },
        "end_time": {
          "type": "string",
          "format": "date-time",
          "title": "The time the report using the control was submitted at"
        },
        "control_summary": {
          "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.ControlSummary",
          "title": "Intentionally blank"
        }
      }
    },
    "chef.automate.api.compliance.reporting.v1.ControlItemRequest": {
      "type": "object",
      "properties": {
        "text": {
          "type": "string",
          "title": "The term to use to match resources on"
        },
        "size": {
          "type": "integer",
          "format": "int32",
          "title": "The maximum number of controls to return (Default 100)"
        },
        "filters": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.ListFilter"
          },
          "title": "The criteria used to filter the controls returned"
        }
      }
    },
    "chef.automate.api.compliance.reporting.v1.ControlItems": {
      "type": "object",
      "properties": {
        "control_items": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.ControlItem"
          },
          "title": "The paginated results of controls matching the filters"
        }
      }
    },
    "chef.automate.api.compliance.reporting.v1.ControlSummary": {
      "type": "object",
      "properties": {
        "total": {
          "type": "integer",
          "format": "int32",
          "title": "The total number of controls"
        },
        "passed": {
          "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.Total",
          "title": "Intentionally blank"
        },
        "skipped": {
          "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.Total",
          "title": "Intentionally blank"
        },
        "failed": {
          "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.Failed",
          "title": "Intentionally blank"
        }
      },
      "title": "A minimal represenation of the statuses of the controls"
    },
    "chef.automate.api.compliance.reporting.v1.Dependency": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "title": "The name of the profile"
        },
        "url": {
          "type": "string",
          "title": "The URL of the profile accessible over HTTP or HTTPS"
        },
        "path": {
          "type": "string",
          "title": "The path to the profile on disk"
        },
        "git": {
          "type": "string",
          "title": "The git URL of the profile"
        },
        "branch": {
          "type": "string",
          "title": "The specific git branch of the dependency"
        },
        "tag": {
          "type": "string",
          "title": "The specific git tag of the dependency"
        },
        "commit": {
          "type": "string",
          "title": "The specific git commit of the dependency"
        },
        "version": {
          "type": "string",
          "title": "The specific git version of the dependency"
        },
        "supermarket": {
          "type": "string",
          "title": "The name of the dependency stored in Chef Supermarket"
        },
        "github": {
          "type": "string",
          "title": "The short name of the dependency stored on Github"
        },
        "compliance": {
          "type": "string",
          "title": "The short name of the dependency stored on the Chef Automate or Chef Compliance server"
        },
        "status": {
          "type": "string",
          "title": "The status of the dependency in the report"
        },
        "skip_message": {
          "type": "string",
          "title": "The reason this profile was skipped in the generated report, if any"
        }
      }
    },
    "chef.automate.api.compliance.reporting.v1.ExportData": {
      "type": "object",
      "properties": {
        "content": {
          "type": "string",
          "format": "byte",
          "title": "The exported reports in the requested format"
        }
      }
    },
    "chef.automate.api.compliance.reporting.v1.Failed": {
      "type": "object",
      "properties": {
        "total": {
          "type": "integer",
          "format": "int32",
          "title": "The total number of failed controls"
        },
        "minor": {
          "type": "integer",
          "format": "int32",
          "title": "The number of failed controls with minor severity"
        },
        "major": {
          "type": "integer",
          "format": "int32",
          "title": "The number of failed controls with major severity"
        },
        "critical": {
          "type": "integer",
          "format": "int32",
          "title": "The number of failed controls with critical severity"
        }
      },
      "title": "Stats of failed controls"
    },
    "chef.automate.api.compliance.reporting.v1.Group": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "title": "The name of the file the controls are defined in"
        },
        "title": {
          "type": "string",
          "title": "The title of control group"
        },
        "controls": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "The ids of the controls defined in this file"
        }
      }
    },
    "chef.automate.api.compliance.reporting.v1.Kv": {
      "type": "object",
      "properties": {
        "key": {
          "type": "string",
          "title": "The key of the tag"
        },
        "value": {
          "type": "string",
          "title": "The value of the tag"
        }
      }
    },
    "chef.automate.api.compliance.reporting.v1.LatestReportSummary": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "title": "The id of the latest report"
        },
        "end_time": {
          "type": "string",
          "format": "date-time",
          "title": "The time the report was submitted at"
        },
        "status": {
          "type": "string",
          "title": "The status of the run the report was made from"
        },
        "controls": {
          "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.ControlSummary",
          "title": "Intentionally blank"
        }
      },
      "title": "A summary of the information contained in the latest report for this node"
    },
    "chef.automate.api.compliance.reporting.v1.ListFilter": {
      "type": "object",
      "properties": {
        "values": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "The values to filter for"
        },
        "type": {
          "type": "string",
          "title": "The field to filter on"
        }
      }
    },
    "chef.automate.api.compliance.reporting.v1.Node": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "title": "The id of this node"
        },
        "name": {
          "type": "string",
          "title": "The name assigned to the node"
        },
        "platform": {
          "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.Platform",
          "title": "Intentionally blank"
        },
        "environment": {
          "type": "string",
          "title": "The environment assigned to the node"
        },
        "latest_report": {
          "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.LatestReportSummary",
          "title": "A summary of the information contained in the latest report for this node"
        },
        "tags": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.Kv"
          },
          "title": "The tags assigned to this node"
        },
        "profiles": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.ProfileMeta"
          },
          "title": "A minimal represenation of the compliance profiles run against the node"
        }
      }
    },
    "chef.automate.api.compliance.reporting.v1.Nodes": {
      "type": "object",
      "properties": {
        "nodes": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.Node"
          },
          "title": "The nodes matching the request filters"
        },
        "total": {
          "type": "integer",
          "format": "int32",
          "title": "The total number of nodes matching the filters"
        },
        "total_passed": {
          "type": "integer",
          "format": "int32",
          "title": "The total number of passing nodes matching the filters"
        },
        "total_failed": {
          "type": "integer",
          "format": "int32",
          "title": "The total number of failed nodes matching the filters"
        },
        "total_skipped": {
          "type": "integer",
          "format": "int32",
          "title": "The total number of skipped nodes matching the filters"
        }
      }
    },
    "chef.automate.api.compliance.reporting.v1.Option": {
      "type": "object",
      "properties": {
        "description": {
          "type": "string",
          "title": "The description of the attribute"
        },
        "default": {
          "type": "string",
          "title": "The default value of the attribute"
        }
      }
    },
    "chef.automate.api.compliance.reporting.v1.Platform": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "title": "The name of the node's operating system"
        },
        "release": {
          "type": "string",
          "title": "The version of the node's operating system"
        },
        "full": {
          "type": "string",
          "title": "The combined name and version of the node's operating system"
        }
      },
      "title": "The name and version of the node's operating system"
    },
    "chef.automate.api.compliance.reporting.v1.Profile": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "title": "The name of the profile. Must be unique"
        },
        "title": {
          "type": "string",
          "title": "The human-readable name of the profile"
        },
        "maintainer": {
          "type": "string",
          "title": "The maintainer listed in the profile metadata"
        },
        "copyright": {
          "type": "string",
          "title": "The name of the copyright holder"
        },
        "copyright_email": {
          "type": "string",
          "title": "The contact information for the copyright holder"
        },
        "license": {
          "type": "string",
          "title": "The license the profile is released under"
        },
        "summary": {
          "type": "string",
          "title": "A short description of the profile"
        },
        "version": {
          "type": "string",
          "title": "The version of the profile"
        },
        "owner": {
          "type": "string",
          "title": "The name of the account that uploaded the profile to Automate"
        },
        "full": {
          "type": "string",
          "title": "The combined name and version of the profile"
        },
        "supports": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.Support"
          },
          "title": "The supported platform targets"
        },
        "depends": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.Dependency"
          },
          "title": "Other profiles that this profile depends on"
        },
        "sha256": {
          "type": "string",
          "title": "A unique value generated from the profile used to identify it"
        },
        "groups": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.Group"
          },
          "title": "The groups of controls defined in the profile"
        },
        "controls": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.Control"
          },
          "title": "The controls defined on the profile"
        },
        "attributes": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.Attribute"
          },
          "title": "The attributes defined on the profile"
        },
        "latest_version": {
          "type": "string",
          "title": "The highest version number of the profile stored in Automate"
        },
        "status": {
          "type": "string",
          "title": "The status of the profile in the generated report"
        },
        "skip_message": {
          "type": "string",
          "title": "The reason this profile was skipped in the generated report, if any"
        }
      }
    },
    "chef.automate.api.compliance.reporting.v1.ProfileCounts": {
      "type": "object",
      "properties": {
        "total": {
          "type": "integer",
          "format": "int32",
          "title": "The total number of nodes matching the filters"
        },
        "failed": {
          "type": "integer",
          "format": "int32",
          "title": "The total number of failed nodes matching the filters"
        },
        "skipped": {
          "type": "integer",
          "format": "int32",
          "title": "The total number of skipped nodes matching the filters"
        },
        "passed": {
          "type": "integer",
          "format": "int32",
          "title": "The total number of passing nodes matching the filters"
        }
      },
      "title": "Stats on the statuses of nodes matching the filters"
    },
    "chef.automate.api.compliance.reporting.v1.ProfileMeta": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "title": "The name of the profile"
        },
        "version": {
          "type": "string",
          "title": "The version of the profile"
        },
        "id": {
          "type": "string",
          "title": "The unique id of the profile"
        },
        "status": {
          "type": "string",
          "title": "The status of the profile run against the node"
        },
        "full": {
          "type": "string",
          "title": "The combined name and version of the profile"
        }
      }
    },
    "chef.automate.api.compliance.reporting.v1.ProfileMin": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "title": "The name of the profile"
        },
        "title": {
          "type": "string",
          "title": "The title of the profile"
        },
        "id": {
          "type": "string",
          "title": "The id of the profile"
        },
        "version": {
          "type": "string",
          "title": "The version of the profile"
        },
        "status": {
          "type": "string",
          "title": "The aggregated status of the profile across the nodes it has been run on"
        }
      },
      "title": "Minimal represenation of a profile"
    },
    "chef.automate.api.compliance.reporting.v1.ProfileMins": {
      "type": "object",
      "properties": {
        "profiles": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.ProfileMin"
          },
          "title": "Minimal represenations of the profiles matching the filters"
        },
        "counts": {
          "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.ProfileCounts",
          "title": "Intentionally blank"
        }
      }
    },
    "chef.automate.api.compliance.reporting.v1.Query": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "title": "Used by the ReadReport endpoint to specify which report to return"
        },
        "type": {
          "type": "string",
          "title": "Used by the ListSuggestions endpoint to control the type of suggestions requested, used by the Export endpoint to control the file format of the returned documents"
        },
        "filters": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.ListFilter"
          },
          "title": "The list of filters used to narrow down the list"
        },
        "order": {
          "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.Query.OrderType",
          "title": "Whether to sort in ascending or descending order"
        },
        "sort": {
          "type": "string",
          "title": "The field to sort the list of results by"
        },
        "page": {
          "type": "integer",
          "format": "int32",
          "title": "The offset to use when paginating requests"
        },
        "per_page": {
          "type": "integer",
          "format": "int32",
          "title": "The number of results to return with each paginated request"
        }
      }
    },
    "chef.automate.api.compliance.reporting.v1.Query.OrderType": {
      "type": "string",
      "enum": [
        "ASC",
        "DESC"
      ],
      "default": "ASC",
      "title": "The two allowed values for ordering results"
    },
    "chef.automate.api.compliance.reporting.v1.Ref": {
      "type": "object",
      "properties": {
        "url": {
          "type": "string",
          "title": "The URL of the external document"
        },
        "ref": {
          "type": "string",
          "title": "The description of the external document"
        }
      }
    },
    "chef.automate.api.compliance.reporting.v1.Report": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "title": "The id of the report"
        },
        "node_id": {
          "type": "string",
          "title": "The id of the node generating the report"
        },
        "node_name": {
          "type": "string",
          "title": "The name of the node generating the report"
        },
        "end_time": {
          "type": "string",
          "format": "date-time",
          "title": "The time the report was submitted at"
        },
        "status": {
          "type": "string",
          "title": "The status of the run the report was made from"
        },
        "controls": {
          "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.ControlSummary",
          "title": "Intentionally blank"
        },
        "environment": {
          "type": "string",
          "title": "The environment of the node generating the report"
        },
        "version": {
          "type": "string",
          "title": "The version of the report"
        },
        "platform": {
          "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.Platform",
          "title": "Intentionally blank"
        },
        "statistics": {
          "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.Statistics",
          "title": "Intentionally blank"
        },
        "profiles": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.Profile"
          },
          "title": "The profiles run as part of this report"
        },
        "job_id": {
          "type": "string",
          "title": "The id of the job associated with the report"
        },
        "ipaddress": {
          "type": "string",
          "title": "The IP address of the node generating the report"
        },
        "fqdn": {
          "type": "string",
          "title": "The FQDN (fully qualified domain name) of the node generating the report"
        }
      }
    },
    "chef.automate.api.compliance.reporting.v1.ReportIds": {
      "type": "object",
      "properties": {
        "ids": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "title": "The list of report ids found matching the query"
        }
      }
    },
    "chef.automate.api.compliance.reporting.v1.Reports": {
      "type": "object",
      "properties": {
        "reports": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.Report"
          },
          "title": "The paginated results of reports matching the filters"
        },
        "total": {
          "type": "integer",
          "format": "int32",
          "title": "The total number of reports matching the filters"
        }
      }
    },
    "chef.automate.api.compliance.reporting.v1.Result": {
      "type": "object",
      "properties": {
        "status": {
          "type": "string",
          "title": "The status of the test"
        },
        "code_desc": {
          "type": "string",
          "title": "The description of the test"
        },
        "run_time": {
          "type": "number",
          "format": "float",
          "title": "The time taken to run the test"
        },
        "start_time": {
          "type": "string",
          "title": "The timestamp of when this individual test was run"
        },
        "message": {
          "type": "string",
          "title": "The reason the test failed, if any"
        },
        "skip_message": {
          "type": "string",
          "title": "The reason the test was skipped, if any"
        }
      }
    },
    "chef.automate.api.compliance.reporting.v1.SourceLocation": {
      "type": "object",
      "properties": {
        "ref": {
          "type": "string",
          "title": "The source code file the control is defined in"
        },
        "line": {
          "type": "integer",
          "format": "int32",
          "title": "The line number the control is defined on"
        }
      }
    },
    "chef.automate.api.compliance.reporting.v1.Statistics": {
      "type": "object",
      "properties": {
        "duration": {
          "type": "number",
          "format": "float",
          "title": "The duration of the report's generation time"
        }
      },
      "title": "Statistics of the report's run"
    },
    "chef.automate.api.compliance.reporting.v1.Suggestion": {
      "type": "object",
      "properties": {
        "text": {
          "type": "string",
          "title": "The content that matched the search term"
        },
        "id": {
          "type": "string",
          "title": "The id of the resource that was suggested"
        },
        "score": {
          "type": "number",
          "format": "float",
          "title": "The confidence in the match quality"
        },
        "version": {
          "type": "string",
          "title": "The version of the suggestion"
        }
      }
    },
    "chef.automate.api.compliance.reporting.v1.SuggestionRequest": {
      "type": "object",
      "properties": {
        "type": {
          "type": "string",
          "title": "The type of resource to get suggestions for"
        },
        "text": {
          "type": "string",
          "title": "The term to use to match resources on"
        },
        "size": {
          "type": "integer",
          "format": "int32",
          "title": "The maximum number of suggestions to return (Default 100)"
        },
        "filters": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.ListFilter"
          },
          "title": "The criteria used to filter the suggestions returned"
        }
      }
    },
    "chef.automate.api.compliance.reporting.v1.Suggestions": {
      "type": "object",
      "properties": {
        "suggestions": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.Suggestion"
          },
          "title": "The list of returned suggestions"
        }
      }
    },
    "chef.automate.api.compliance.reporting.v1.Support": {
      "type": "object",
      "properties": {
        "os_name": {
          "type": "string",
          "title": "The name of the supported operating system"
        },
        "os_family": {
          "type": "string",
          "title": "The name of the broader category of the supported platform (eg, linux, windows)"
        },
        "release": {
          "type": "string",
          "title": "The specific release of the operating system this profile supports"
        },
        "inspec_version": {
          "type": "string",
          "title": "The supported inspec version this profile was made to run on"
        },
        "platform": {
          "type": "string",
          "title": "The platform name and version combined"
        }
      }
    },
    "chef.automate.api.compliance.reporting.v1.Total": {
      "type": "object",
      "properties": {
        "total": {
          "type": "integer",
          "format": "int32",
          "title": "The number of controls"
        }
      },
      "title": "A subtotal of controls"
    },
    "google.protobuf.Any": {
      "type": "object",
      "properties": {
        "type_url": {
          "type": "string",
          "description": "A URL/resource name that uniquely identifies the type of the serialized\nprotocol buffer message. This string must contain at least\none \"/\" character. The last segment of the URL's path must represent\nthe fully qualified name of the type (as in\n` + "`" + `path/google.protobuf.Duration` + "`" + `). The name should be in a canonical form\n(e.g., leading \".\" is not accepted).\n\nIn practice, teams usually precompile into the binary all types that they\nexpect it to use in the context of Any. However, for URLs which use the\nscheme ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + `, or no scheme, one can optionally set up a type\nserver that maps type URLs to message definitions as follows:\n\n* If no scheme is provided, ` + "`" + `https` + "`" + ` is assumed.\n* An HTTP GET on the URL must yield a [google.protobuf.Type][]\n  value in binary format, or produce an error.\n* Applications are allowed to cache lookup results based on the\n  URL, or have them precompiled into a binary to avoid any\n  lookup. Therefore, binary compatibility needs to be preserved\n  on changes to types. (Use versioned type names to manage\n  breaking changes.)\n\nNote: this functionality is not currently available in the official\nprotobuf release, and it is not used for type URLs beginning with\ntype.googleapis.com.\n\nSchemes other than ` + "`" + `http` + "`" + `, ` + "`" + `https` + "`" + ` (or the empty scheme) might be\nused with implementation specific semantics."
        },
        "value": {
          "type": "string",
          "format": "byte",
          "description": "Must be a valid serialized protocol buffer of the above specified type."
        }
      },
      "description": "` + "`" + `Any` + "`" + ` contains an arbitrary serialized protocol buffer message along with a\nURL that describes the type of the serialized message.\n\nProtobuf library provides support to pack/unpack Any values in the form\nof utility functions or additional generated methods of the Any type.\n\nExample 1: Pack and unpack a message in C++.\n\n    Foo foo = ...;\n    Any any;\n    any.PackFrom(foo);\n    ...\n    if (any.UnpackTo(\u0026foo)) {\n      ...\n    }\n\nExample 2: Pack and unpack a message in Java.\n\n    Foo foo = ...;\n    Any any = Any.pack(foo);\n    ...\n    if (any.is(Foo.class)) {\n      foo = any.unpack(Foo.class);\n    }\n\n Example 3: Pack and unpack a message in Python.\n\n    foo = Foo(...)\n    any = Any()\n    any.Pack(foo)\n    ...\n    if any.Is(Foo.DESCRIPTOR):\n      any.Unpack(foo)\n      ...\n\n Example 4: Pack and unpack a message in Go\n\n     foo := \u0026pb.Foo{...}\n     any, err := ptypes.MarshalAny(foo)\n     ...\n     foo := \u0026pb.Foo{}\n     if err := ptypes.UnmarshalAny(any, foo); err != nil {\n       ...\n     }\n\nThe pack methods provided by protobuf library will by default use\n'type.googleapis.com/full.type.name' as the type URL and the unpack\nmethods only use the fully qualified type name after the last '/'\nin the type URL, for example \"foo.bar.com/x/y.z\" will yield type\nname \"y.z\".\n\n\nJSON\n====\nThe JSON representation of an ` + "`" + `Any` + "`" + ` value uses the regular\nrepresentation of the deserialized, embedded message, with an\nadditional field ` + "`" + `@type` + "`" + ` which contains the type URL. Example:\n\n    package google.profile;\n    message Person {\n      string first_name = 1;\n      string last_name = 2;\n    }\n\n    {\n      \"@type\": \"type.googleapis.com/google.profile.Person\",\n      \"firstName\": \u003cstring\u003e,\n      \"lastName\": \u003cstring\u003e\n    }\n\nIf the embedded message type is well-known and has a custom JSON\nrepresentation, that representation will be embedded adding a field\n` + "`" + `value` + "`" + ` which holds the custom JSON in addition to the ` + "`" + `@type` + "`" + `\nfield. Example (for message [google.protobuf.Duration][]):\n\n    {\n      \"@type\": \"type.googleapis.com/google.protobuf.Duration\",\n      \"value\": \"1.212s\"\n    }"
    },
    "grpc.gateway.runtime.StreamError": {
      "type": "object",
      "properties": {
        "grpc_code": {
          "type": "integer",
          "format": "int32"
        },
        "http_code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        },
        "http_status": {
          "type": "string"
        },
        "details": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/google.protobuf.Any"
          }
        }
      }
    }
  },
  "x-stream-definitions": {
    "chef.automate.api.compliance.reporting.v1.ExportData": {
      "type": "object",
      "properties": {
        "result": {
          "$ref": "#/definitions/chef.automate.api.compliance.reporting.v1.ExportData"
        },
        "error": {
          "$ref": "#/definitions/grpc.gateway.runtime.StreamError"
        }
      },
      "title": "Stream result of chef.automate.api.compliance.reporting.v1.ExportData"
    }
  }
}
`)
}
