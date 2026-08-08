package auth

var SystemRoles = map[string][]PermissionKey{
	"owner":  { "devices:read", "devices:write", "interfaces:read", "interfaces:write", "networks:read", "networks:write", "organization:read", "organization:write", "users:read", "users:write", "organization_members:read", "organization_members:write" },
	"admin":  { "devices:read", "devices:write", "interfaces:read", "interfaces:write", "networks:read", "networks:write", "organization:read", "users:read", "users:write", "organization_members:read", "organization_members:write" },
	"operator":   { "devices:read", "devices:write", "interfaces:read", "interfaces:write", "networks:read", "networks:write", "organization:read", "users:read" },
	"viewer": { "devices:read", "interfaces:read", "networks:read", "organization:read", "users:read" },
}