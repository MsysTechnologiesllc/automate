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
        "description": "Lists controls from the last run, with the option of applying filters. Supports filtering, but not pagination or sorting.",
        "operationId": "ListControlItems",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/v1ControlItems"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/v1ControlItemRequest"
            }
          }
        ],
        "tags": [
          "Controls"
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
              "$ref": "#/definitions/v1Node"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "description": "Unique identifier.",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "Nodes"
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
              "$ref": "#/definitions/v1Nodes"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/v1Query"
            }
          }
        ],
        "tags": [
          "Nodes"
        ]
      }
    },
    "/compliance/reporting/profiles": {
      "post": {
        "summary": "List profiles",
        "description": "List all profiles in use, with the option of applying filters. Adding a profile filter returns a list of all the nodes using that profile. Supports pagination, filtering, and sorting.\n\n| Sort paramter | Sort value |\n| --- | --- |\n| name | name.lower |\n| title (default) | title.lower |",
        "operationId": "ListProfiles",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/v1ProfileMins"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/v1Query"
            }
          }
        ],
        "tags": [
          "Profiles"
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
              "$ref": "#/definitions/v1ReportIds"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/v1Query"
            }
          }
        ],
        "tags": [
          "Reports"
        ]
      }
    },
    "/compliance/reporting/reports": {
      "post": {
        "summary": "List reports",
        "description": "Lists all reports. Apply an optional filter to return a select group of reports. Supports pagination, filtering, and sorting.\n\n| Sort paramter | Sort value |\n| --- | --- |\n| latest_report.controls.failed.critical | controls_sums.failed.critical |\n| latest_report.controls.failed.total | controls_sums.failed.total |\n| latest_report.end_time (default) | end_time |\n| latest_report.status | status |\n| node_name | node_name.lower |",
        "operationId": "ListReports",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/v1Reports"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/v1Query"
            }
          }
        ],
        "tags": [
          "Reports"
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
              "$ref": "#/definitions/v1Report"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "description": "Unique idenfifier.",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/v1Query"
            }
          }
        ],
        "tags": [
          "Reports"
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
              "$ref": "#/definitions/v1Suggestions"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/v1SuggestionRequest"
            }
          }
        ],
        "tags": [
          "Reports"
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
              "$ref": "#/definitions/versionVersionInfo"
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
    "QueryOrderType": {
      "type": "string",
      "enum": [
        "ASC",
        "DESC"
      ],
      "default": "ASC",
      "description": "Sort the results in ascending or descending order."
    },
    "protobufAny": {
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
    "runtimeStreamError": {
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
            "$ref": "#/definitions/protobufAny"
          }
        }
      }
    },
    "v1Attribute": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "The name of the attribute."
        },
        "options": {
          "$ref": "#/definitions/v1Option",
          "description": "The options defined for the attribute."
        }
      }
    },
    "v1Control": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "The unique id of this control."
        },
        "code": {
          "type": "string",
          "description": "The full ruby code of the control defined in the profile."
        },
        "desc": {
          "type": "string",
          "description": "The full description of the control."
        },
        "impact": {
          "type": "number",
          "format": "float",
          "description": "The severity of the control."
        },
        "title": {
          "type": "string",
          "description": "The compact description of the control."
        },
        "source_location": {
          "$ref": "#/definitions/v1SourceLocation",
          "description": "Intentionally blank."
        },
        "results": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/v1Result"
          },
          "description": "The results of running all tests defined in the control against the node."
        },
        "refs": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/v1Ref"
          },
          "description": "External supporting documents for the control."
        },
        "tags": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          },
          "description": "Metadata defined on the control in key-value format."
        }
      }
    },
    "v1ControlItem": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "The unique id of this control."
        },
        "title": {
          "type": "string",
          "description": "The compact description of the control."
        },
        "profile": {
          "$ref": "#/definitions/v1ProfileMin",
          "description": "Intentionally blank."
        },
        "impact": {
          "type": "number",
          "format": "float",
          "description": "The severity of the control."
        },
        "end_time": {
          "type": "string",
          "format": "date-time",
          "description": "The time the report using the control was submitted at."
        },
        "control_summary": {
          "$ref": "#/definitions/v1ControlSummary",
          "description": "Intentionally blank."
        }
      }
    },
    "v1ControlItemRequest": {
      "type": "object",
      "properties": {
        "text": {
          "type": "string",
          "description": "The term to use to match resources on."
        },
        "size": {
          "type": "integer",
          "format": "int32",
          "description": "The maximum number of controls to return (Default 100)."
        },
        "filters": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/v1ListFilter"
          },
          "description": "The criteria used to filter the controls returned."
        }
      }
    },
    "v1ControlItems": {
      "type": "object",
      "properties": {
        "control_items": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/v1ControlItem"
          },
          "description": "The paginated results of controls matching the filters."
        }
      }
    },
    "v1ControlSummary": {
      "type": "object",
      "properties": {
        "total": {
          "type": "integer",
          "format": "int32",
          "description": "The total number of controls in the report."
        },
        "passed": {
          "$ref": "#/definitions/v1Total",
          "description": "Intentionally blank."
        },
        "skipped": {
          "$ref": "#/definitions/v1Total",
          "description": "Intentionally blank."
        },
        "failed": {
          "$ref": "#/definitions/v1Failed",
          "description": "Intentionally blank."
        }
      },
      "description": "A minimal represenation of the statuses of the controls in the report."
    },
    "v1Dependency": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "The name of the profile."
        },
        "url": {
          "type": "string",
          "description": "The URL of the profile accessible over HTTP or HTTPS."
        },
        "path": {
          "type": "string",
          "description": "The path to the profile on disk."
        },
        "git": {
          "type": "string",
          "description": "The git URL of the profile."
        },
        "branch": {
          "type": "string",
          "description": "The specific git branch of the dependency."
        },
        "tag": {
          "type": "string",
          "description": "The specific git tag of the dependency."
        },
        "commit": {
          "type": "string",
          "description": "The specific git commit of the dependency."
        },
        "version": {
          "type": "string",
          "description": "The specific git version of the dependency."
        },
        "supermarket": {
          "type": "string",
          "description": "The name of the dependency stored in Chef Supermarket."
        },
        "github": {
          "type": "string",
          "description": "The short name of the dependency stored on Github."
        },
        "compliance": {
          "type": "string",
          "description": "The short name of the dependency stored on the Chef Automate or Chef Compliance server."
        },
        "status": {
          "type": "string",
          "description": "The status of the dependency in the report."
        },
        "skip_message": {
          "type": "string",
          "description": "The reason this profile was skipped in the generated report, if any."
        }
      }
    },
    "v1ExportData": {
      "type": "object",
      "properties": {
        "content": {
          "type": "string",
          "format": "byte",
          "description": "Exported reports in JSON or CSV."
        }
      }
    },
    "v1Failed": {
      "type": "object",
      "properties": {
        "total": {
          "type": "integer",
          "format": "int32",
          "description": "The total number of failed controls."
        },
        "minor": {
          "type": "integer",
          "format": "int32",
          "description": "The number of failed controls with minor severity."
        },
        "major": {
          "type": "integer",
          "format": "int32",
          "description": "The number of failed controls with major severity."
        },
        "critical": {
          "type": "integer",
          "format": "int32",
          "description": "The number of failed controls with critical severity."
        }
      },
      "description": "Stats of failed controls."
    },
    "v1Group": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "The name of the file the controls are defined in."
        },
        "title": {
          "type": "string",
          "description": "The title of control group."
        },
        "controls": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "The ids of the controls defined in this file."
        }
      }
    },
    "v1Kv": {
      "type": "object",
      "properties": {
        "key": {
          "type": "string",
          "description": "The key of the tag."
        },
        "value": {
          "type": "string",
          "description": "The value of the tag."
        }
      }
    },
    "v1LatestReportSummary": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "The id of the latest report."
        },
        "end_time": {
          "type": "string",
          "format": "date-time",
          "description": "The time the report was submitted at."
        },
        "status": {
          "type": "string",
          "description": "The status of the run the report was made from."
        },
        "controls": {
          "$ref": "#/definitions/v1ControlSummary",
          "description": "Intentionally blank."
        }
      },
      "description": "A summary of the information contained in the latest report for this node."
    },
    "v1ListFilter": {
      "type": "object",
      "properties": {
        "values": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "Filters applied to the list."
        },
        "type": {
          "type": "string",
          "description": "The field to filter on."
        }
      }
    },
    "v1Node": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "The id of this node."
        },
        "name": {
          "type": "string",
          "description": "The name assigned to the node."
        },
        "platform": {
          "$ref": "#/definitions/v1Platform",
          "description": "Intentionally blank."
        },
        "environment": {
          "type": "string",
          "description": "The environment assigned to the node."
        },
        "latest_report": {
          "$ref": "#/definitions/v1LatestReportSummary",
          "description": "A summary of the information contained in the latest report for this node."
        },
        "tags": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/v1Kv"
          },
          "description": "The tags assigned to this node."
        },
        "profiles": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/v1ProfileMeta"
          },
          "description": "A minimal represenation of the compliance profiles run against the node."
        }
      }
    },
    "v1Nodes": {
      "type": "object",
      "properties": {
        "nodes": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/v1Node"
          },
          "description": "The nodes matching the request filters."
        },
        "total": {
          "type": "integer",
          "format": "int32",
          "description": "The total number of nodes matching the filters."
        },
        "total_passed": {
          "type": "integer",
          "format": "int32",
          "description": "The total number of passing nodes matching the filters."
        },
        "total_failed": {
          "type": "integer",
          "format": "int32",
          "description": "The total number of failed nodes matching the filters."
        },
        "total_skipped": {
          "type": "integer",
          "format": "int32",
          "description": "The total number of skipped nodes matching the filters."
        }
      }
    },
    "v1Option": {
      "type": "object",
      "properties": {
        "description": {
          "type": "string",
          "description": "The description of the attribute."
        },
        "default": {
          "type": "string",
          "description": "The default value of the attribute."
        }
      }
    },
    "v1Platform": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "The name of the node's operating system."
        },
        "release": {
          "type": "string",
          "description": "The version of the node's operating system."
        },
        "full": {
          "type": "string",
          "description": "The combined name and version of the node's operating system."
        }
      },
      "description": "The name and version of the node's operating system."
    },
    "v1Profile": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "The name of the profile. Must be unique."
        },
        "title": {
          "type": "string",
          "description": "The profile title."
        },
        "maintainer": {
          "type": "string",
          "description": "The maintainer listed in the profile metadata."
        },
        "copyright": {
          "type": "string",
          "description": "The name of the copyright holder."
        },
        "copyright_email": {
          "type": "string",
          "description": "The contact information for the copyright holder."
        },
        "license": {
          "type": "string",
          "description": "The license the profile is released under."
        },
        "summary": {
          "type": "string",
          "description": "A short description of the profile."
        },
        "version": {
          "type": "string",
          "description": "The version of the profile."
        },
        "owner": {
          "type": "string",
          "description": "The name of the account that uploaded the profile to Automate."
        },
        "full": {
          "type": "string",
          "description": "The combined name and version of the profile."
        },
        "supports": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/v1Support"
          },
          "description": "The supported platform targets."
        },
        "depends": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/v1Dependency"
          },
          "description": "Other profiles that this profile depends on."
        },
        "sha256": {
          "type": "string",
          "description": "A unique value generated from the profile used to identify it."
        },
        "groups": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/v1Group"
          },
          "description": "The groups of controls defined in the profile."
        },
        "controls": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/v1Control"
          },
          "description": "The controls defined on the profile."
        },
        "attributes": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/v1Attribute"
          },
          "description": "The attributes defined on the profile."
        },
        "latest_version": {
          "type": "string",
          "description": "The highest version number of the profile stored in Automate."
        },
        "status": {
          "type": "string",
          "description": "The status of the profile in the generated report."
        },
        "skip_message": {
          "type": "string",
          "description": "The reason this profile was skipped in the generated report, if any."
        }
      }
    },
    "v1ProfileCounts": {
      "type": "object",
      "properties": {
        "total": {
          "type": "integer",
          "format": "int32",
          "description": "The total number of nodes matching the filters."
        },
        "failed": {
          "type": "integer",
          "format": "int32",
          "description": "The total number of failed nodes matching the filters."
        },
        "skipped": {
          "type": "integer",
          "format": "int32",
          "description": "The total number of skipped nodes matching the filters."
        },
        "passed": {
          "type": "integer",
          "format": "int32",
          "description": "The total number of passing nodes matching the filters."
        }
      },
      "description": "Stats on the statuses of nodes matching the filters."
    },
    "v1ProfileMeta": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "The name of the profile."
        },
        "version": {
          "type": "string",
          "description": "The version of the profile."
        },
        "id": {
          "type": "string",
          "description": "The unique id of the profile."
        },
        "status": {
          "type": "string",
          "description": "The status of the profile run against the node."
        },
        "full": {
          "type": "string",
          "description": "The combined name and version of the profile."
        }
      }
    },
    "v1ProfileMin": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string",
          "description": "The name of the profile."
        },
        "title": {
          "type": "string",
          "description": "The title of the profile."
        },
        "id": {
          "type": "string",
          "description": "The id of the profile."
        },
        "version": {
          "type": "string",
          "description": "The version of the profile."
        },
        "status": {
          "type": "string",
          "description": "The aggregated status of the profile across the nodes it has been run on."
        }
      },
      "description": "Minimal represenation of a profile."
    },
    "v1ProfileMins": {
      "type": "object",
      "properties": {
        "profiles": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/v1ProfileMin"
          },
          "description": "Minimal represenations of the profiles matching the filters."
        },
        "counts": {
          "$ref": "#/definitions/v1ProfileCounts",
          "description": "Intentionally blank."
        }
      }
    },
    "v1Query": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "Unique idenfifier."
        },
        "type": {
          "type": "string",
          "description": "File type, either JSON or CSV."
        },
        "filters": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/v1ListFilter"
          },
          "description": "Filters applied to the report results."
        },
        "order": {
          "$ref": "#/definitions/QueryOrderType"
        },
        "sort": {
          "type": "string",
          "description": "Sort the list of results by a field."
        },
        "page": {
          "type": "integer",
          "format": "int32",
          "description": "The offset for paginating requests. An offset defines a place in the results in order to fetch the next page of the results."
        },
        "per_page": {
          "type": "integer",
          "format": "int32",
          "description": "The number of results on each paginated request page."
        }
      }
    },
    "v1Ref": {
      "type": "object",
      "properties": {
        "url": {
          "type": "string",
          "description": "The external document URL."
        },
        "ref": {
          "type": "string",
          "description": "A description of the external document."
        }
      }
    },
    "v1Report": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "A unique report identifier."
        },
        "node_id": {
          "type": "string",
          "description": "The id of the node making the report."
        },
        "node_name": {
          "type": "string",
          "description": "The name of the node making the report."
        },
        "end_time": {
          "type": "string",
          "format": "date-time",
          "description": "The time that the report was completed."
        },
        "status": {
          "type": "string",
          "description": "The status of the run the report was made from."
        },
        "controls": {
          "$ref": "#/definitions/v1ControlSummary",
          "description": "Intentionally blank."
        },
        "environment": {
          "type": "string",
          "description": "The environment of the node making the report."
        },
        "version": {
          "type": "string",
          "description": "The version of the report."
        },
        "platform": {
          "$ref": "#/definitions/v1Platform",
          "description": "Intentionally blank."
        },
        "statistics": {
          "$ref": "#/definitions/v1Statistics",
          "description": "Intentionally blank."
        },
        "profiles": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/v1Profile"
          },
          "description": "The profiles run as part of this report."
        },
        "job_id": {
          "type": "string",
          "description": "The id of the compliance scan job associated with the report."
        },
        "ipaddress": {
          "type": "string",
          "description": "The IP address of the node making the report."
        },
        "fqdn": {
          "type": "string",
          "description": "The FQDN (fully qualified domain name) of the node making the report."
        }
      }
    },
    "v1ReportIds": {
      "type": "object",
      "properties": {
        "ids": {
          "type": "array",
          "items": {
            "type": "string"
          },
          "description": "The list of unique report identifiers found matching the query."
        }
      }
    },
    "v1Reports": {
      "type": "object",
      "properties": {
        "reports": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/v1Report"
          },
          "description": "Paginated results of reports matching the filters."
        },
        "total": {
          "type": "integer",
          "format": "int32",
          "description": "Total number of reports matching the filters."
        }
      }
    },
    "v1Result": {
      "type": "object",
      "properties": {
        "status": {
          "type": "string",
          "description": "The status of the test."
        },
        "code_desc": {
          "type": "string",
          "description": "The description of the test."
        },
        "run_time": {
          "type": "number",
          "format": "float",
          "description": "The time taken to run the test."
        },
        "start_time": {
          "type": "string",
          "description": "The timestamp of when this individual test was run."
        },
        "message": {
          "type": "string",
          "description": "The reason the test failed, if any."
        },
        "skip_message": {
          "type": "string",
          "description": "The reason the test was skipped, if any."
        }
      }
    },
    "v1SourceLocation": {
      "type": "object",
      "properties": {
        "ref": {
          "type": "string",
          "description": "The source code file the control is defined in."
        },
        "line": {
          "type": "integer",
          "format": "int32",
          "description": "The line number the control is defined on."
        }
      }
    },
    "v1Statistics": {
      "type": "object",
      "properties": {
        "duration": {
          "type": "number",
          "format": "float",
          "description": "The duration of the report's generation time."
        }
      },
      "description": "Statistics of the report's run."
    },
    "v1Suggestion": {
      "type": "object",
      "properties": {
        "text": {
          "type": "string",
          "description": "The content that matched the search term."
        },
        "id": {
          "type": "string",
          "description": "The id of the resource that was suggested."
        },
        "score": {
          "type": "number",
          "format": "float",
          "description": "The confidence in the match quality."
        },
        "version": {
          "type": "string",
          "description": "The version of the suggestion."
        }
      }
    },
    "v1SuggestionRequest": {
      "type": "object",
      "properties": {
        "type": {
          "type": "string",
          "description": "The type of resource to get suggestions for."
        },
        "text": {
          "type": "string",
          "description": "The term to use to match resources on."
        },
        "size": {
          "type": "integer",
          "format": "int32",
          "description": "The maximum number of suggestions to return."
        },
        "filters": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/v1ListFilter"
          },
          "description": "The criteria used to filter the suggestions returned."
        }
      }
    },
    "v1Suggestions": {
      "type": "object",
      "properties": {
        "suggestions": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/v1Suggestion"
          },
          "description": "The list of returned suggestions."
        }
      }
    },
    "v1Support": {
      "type": "object",
      "properties": {
        "os_name": {
          "type": "string",
          "description": "The name of the supported operating system."
        },
        "os_family": {
          "type": "string",
          "description": "The wider category of supported platform (e.g., linux, windows)."
        },
        "release": {
          "type": "string",
          "description": "The specific operating system release number this profile supports."
        },
        "inspec_version": {
          "type": "string",
          "description": "The supported inspec version for this profile."
        },
        "platform": {
          "type": "string",
          "description": "The platform name and version combined."
        }
      }
    },
    "v1Total": {
      "type": "object",
      "properties": {
        "total": {
          "type": "integer",
          "format": "int32",
          "description": "The total number of controls."
        }
      },
      "description": "A subtotal of controls."
    },
    "versionVersionInfo": {
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
    }
  },
  "x-stream-definitions": {
    "v1ExportData": {
      "type": "object",
      "properties": {
        "result": {
          "$ref": "#/definitions/v1ExportData"
        },
        "error": {
          "$ref": "#/definitions/runtimeStreamError"
        }
      },
      "title": "Stream result of v1ExportData"
    }
  }
}
`)
}
