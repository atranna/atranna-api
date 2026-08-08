package auth

type PermissionKey string

type Permission struct {
	Key PermissionKey
	Resource string
	Action   string
	Description string
}

var Catalog = []Permission{
	{Key: "devices:read", Resource: "devices", Action: "read", Description: "Read actions on devices (e.g. GET /devices, GET /devices/:id)"},
	{Key: "devices:write", Resource: "devices", Action: "write", Description: "Write actions on devices (e.g. POST /devices, DELETE /devices/:id)"},
	{Key: "interfaces:read", Resource: "interfaces", Action: "read", Description: "Read actions on interfaces (e.g. GET /interfaces, GET /interfaces/:id)"},
	{Key: "interfaces:write", Resource: "interfaces", Action: "write", Description: "Write actions on interfaces (e.g. POST /interfaces, DELETE /interfaces/:id)"},
	{Key: "networks:read", Resource: "networks", Action: "read", Description: "Read actions on networks (e.g. GET /networks, GET /networks/:id)"},
	{Key: "networks:write", Resource: "networks", Action: "write", Description: "Write actions on networks (e.g. POST /networks, DELETE /networks/:id)"},
	{Key: "organization:read", Resource: "organizations", Action: "read", Description: "Read actions on organizations (e.g. GET /organizations, GET /organizations/:id)"},
	{Key: "organization:write", Resource: "organizations", Action: "write", Description: "Write actions on organizations (e.g. DELETE /organizations/:id)"},
	{Key: "users:read", Resource: "users", Action: "read", Description: "Read actions on users (e.g. GET /users, GET /users/:id)"},
	{Key: "users:write", Resource: "users", Action: "write", Description: "Write actions on users (e.g. POST /users, DELETE /users/:id)"},
	{Key: "organization_members:read", Resource: "organization_members", Action: "read", Description: "Read actions on organization members (e.g. GET /organizations/:id/members, GET /organizations/:id/members/:user_id)"},
	{Key: "organization_members:write", Resource: "organization_members", Action: "write", Description: "Write actions on organization members (e.g. POST /organizations/:id/members, DELETE /organizations/:id/members/:user_id"},
	{Key: "roles:read", Resource: "roles", Action: "read", Description: "Read actions on roles (e.g. GET /roles, GET /roles/:id)"},
	{Key: "roles:write", Resource: "roles", Action: "write", Description: "Write actions on roles (e.g. POST /roles, DELETE /roles/:id)"},
}

const DevicesRead PermissionKey = "devices:read"
const DevicesWrite PermissionKey = "devices:write"
const InterfacesRead PermissionKey = "interfaces:read"
const InterfacesWrite PermissionKey = "interfaces:write"
const NetworksRead PermissionKey = "networks:read"
const NetworksWrite PermissionKey = "networks:write"
const OrganizationRead PermissionKey = "organization:read"
const OrganizationWrite PermissionKey = "organization:write"
const UsersRead PermissionKey = "users:read"
const UsersWrite PermissionKey = "users:write"
const OrganizationMembersRead PermissionKey = "organization_members:read"
const OrganizationMembersWrite PermissionKey = "organization_members:write"
const RolesRead PermissionKey = "roles:read"
const RolesWrite PermissionKey = "roles:write"