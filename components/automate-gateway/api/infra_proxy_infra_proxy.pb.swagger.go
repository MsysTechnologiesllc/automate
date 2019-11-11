package api

func init() {
	Swagger.Add("infra_proxy_infra_proxy", `{
  "swagger": "2.0",
  "info": {
    "title": "components/automate-gateway/api/infra_proxy/infra_proxy.proto",
    "version": "version not set"
  },
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/infra_proxy/servers": {
      "get": {
        "operationId": "GetServers",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.GetServers"
            }
          }
        },
        "tags": [
          "InfraProxy"
        ]
      }
    },
    "/infra_proxy/servers/orgs": {
      "get": {
        "operationId": "GetOrgs",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.infra_proxy.response.GetOrgs"
            }
          }
        },
        "tags": [
          "InfraProxy"
        ]
      }
    }
  },
  "definitions": {
    "chef.automate.api.infra_proxy.response.GetOrgs": {
      "type": "object",
      "properties": {
        "orgs": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Org"
          }
        }
      }
    },
    "chef.automate.api.infra_proxy.response.GetServers": {
      "type": "object",
      "properties": {
        "servers": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.infra_proxy.response.Server"
          }
        }
      }
    },
    "chef.automate.api.infra_proxy.response.Org": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "admin_user": {
          "type": "string"
        },
        "admin_key": {
          "type": "string"
        },
        "server_id": {
          "type": "string"
        }
      }
    },
    "chef.automate.api.infra_proxy.response.Server": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "description": {
          "type": "string"
        },
        "fqdn": {
          "type": "string"
        },
        "ip_address": {
          "type": "string"
        },
        "orgs_count": {
          "type": "integer",
          "format": "int32"
        }
      }
    }
  }
}
`)
}
