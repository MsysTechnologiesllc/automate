// Code generated by protoc-gen-policy. DO NOT EDIT.
// source: api/external/infra_proxy/infra_proxy.proto

package infra_proxy

import (
	request "github.com/chef/automate/api/external/infra_proxy/request"
	policy "github.com/chef/automate/components/automate-gateway/api/authz/policy"
)

func init() {
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetVersion", "service_info:version", "read", "GET", "/infra_proxy/version", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetServers", "infra_proxy:servers", "read", "GET", "/infra_proxy/servers", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetServer", "infra_proxy:servers:{id}", "read", "GET", "/infra_proxy/servers/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.GetServer); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
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
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetServerByName", "infra_proxy:servers:{name}", "read", "GET", "/infra_proxy/servers/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.GetServerByName); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
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
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/CreateServer", "infra_proxy:servers", "create", "POST", "/infra_proxy/servers", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.CreateServer); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
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
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/UpdateServer", "infra_proxy:servers:{id}", "update", "PUT", "/infra_proxy/servers/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.UpdateServer); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
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
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/DeleteServer", "infra_proxy:servers:{id}", "delete", "DELETE", "/infra_proxy/servers/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.DeleteServer); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
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
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetOrgs", "infra_proxy:servers:{server_id}", "read", "GET", "/infra_proxy/servers/{server_id}/orgs", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.GetOrgs); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
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
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetOrg", "infra_proxy:servers:{server_id}:orgs:{id}", "read", "GET", "/infra_proxy/servers/{server_id}/orgs/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.GetOrg); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
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
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/GetOrgByName", "infra_proxy:servers:{server_id}:orgs:{name}", "read", "GET", "/infra_proxy/servers/{server_id}/orgs/{name}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.GetOrgByName); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
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
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/CreateOrg", "infra:servers:{server_id}:orgs", "create", "POST", "/infra_proxy/servers/{server_id}/orgs", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.CreateOrg); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
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
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/UpdateOrg", "infra_proxy:servers:{server_id}:orgs:{id}", "update", "PUT", "/infra_proxy/servers/{server_id}/orgs/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.UpdateOrg); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
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
	policy.MapMethodTo("/chef.automate.api.infra_proxy.InfraProxy/DeleteOrg", "infra_proxy:servers:{server_id}:orgs:{id}", "delete", "DELETE", "/infra_proxy/servers/{server_id}/orgs/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.DeleteOrg); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
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
}
