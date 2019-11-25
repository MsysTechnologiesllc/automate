// Code generated by protoc-gen-policy. DO NOT EDIT.
// source: api/external/infra_proxy/infra_proxy.proto

package infra_proxy

import (
	request "github.com/chef/automate/api/external/infra_proxy/request"
	policyv2 "github.com/chef/automate/components/automate-gateway/authz/policy_v2"
)

func init() {
	policyv2.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetVersion", "system:service:version", "system:serviceVersion:get", "GET", "/infra_proxy/version", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policyv2.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetServers", "infra:servers", "infra:servers:list", "GET", "/infra_proxy/servers", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policyv2.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetServer", "infra:servers:{id}", "infra:servers:get", "GET", "/infra_proxy/servers/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.GetServer); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetServerByName", "infra:servers:{name}", "infra:servers:get", "GET", "/infra_proxy/servers/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.GetServerByName); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/CreateServer", "infra:servers", "infra:servers:create", "POST", "/infra_proxy/servers", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.CreateServer); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "name":
					return m.Name
				case "description":
					return m.Description
				case "fqdn":
					return m.Fqdn
				case "ip_address":
					return m.IpAddress
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/UpdateServer", "infra:servers:{id}", "infra:servers:update", "PUT", "/infra_proxy/servers/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.UpdateServer); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				case "description":
					return m.Description
				case "fqdn":
					return m.Fqdn
				case "ip_address":
					return m.IpAddress
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/DeleteServer", "infra:servers:{id}", "infra:servers:delete", "DELETE", "/infra_proxy/servers/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.DeleteServer); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetOrgs", "infra:servers:{server_id}:orgs", "infra:serverOrgs:get", "GET", "/infra_proxy/servers/{server_id}/orgs", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.GetOrgs); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "server_id":
					return m.ServerId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetOrg", "infra:servers:{server_id}:orgs:{id}", "infra:serverOrgs:get", "GET", "/infra_proxy/servers/{server_id}/orgs/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.GetOrg); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "server_id":
					return m.ServerId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetOrgByName", "infra:servers:{server_id}:orgs:{name}", "infra:serverOrgs:get", "GET", "/infra_proxy/servers/{server_id}/orgs/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.GetOrgByName); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "name":
					return m.Name
				case "server_id":
					return m.ServerId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/CreateOrg", "infra:servers:{server_id}:orgs", "infra:serverOrgs:create", "POST", "/infra_proxy/servers/{server_id}/orgs", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.CreateOrg); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "name":
					return m.Name
				case "admin_user":
					return m.AdminUser
				case "admin_key":
					return m.AdminKey
				case "server_id":
					return m.ServerId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/UpdateOrg", "infra:servers:{server_id}:orgs:{id}", "infra:serverOrgs:update", "PUT", "/infra_proxy/servers/{server_id}/orgs/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.UpdateOrg); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				case "admin_user":
					return m.AdminUser
				case "admin_key":
					return m.AdminKey
				case "server_id":
					return m.ServerId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/DeleteOrg", "infra:servers:{server_id}:orgs:{id}", "infra:serverOrgs:delete", "DELETE", "/infra_proxy/servers/{server_id}/orgs/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.DeleteOrg); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "server_id":
					return m.ServerId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetCookbooks", "infra:servers:{server_id}:orgs:{org_id}:cookbooks", "infra:serverOrgCookbooks:get", "GET", "/infra_proxy/servers/{server_id}/orgs/{org_id}/cookbooks", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Cookbooks); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetCookbooksAvailableVersions", "infra:servers:{server_id}:orgs:{org_id}:cookbooks", "infra:serverOrgCookbooks:get", "GET", "/infra_proxy/servers/{server_id}/orgs/{org_id}/cookbooks/num_versions", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.CookbooksAvailableVersions); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "org_id":
					return m.OrgId
				case "server_id":
					return m.ServerId
				case "num_versions":
					return m.NumVersions
				default:
					return ""
				}
			})
		}
		return ""
	})
}
