package auth

var SystemRoles = map[string][]PermissionKey{
	"owner":  {}, // Has permission to do anything, so no need to list permissions here
	"admin":  { DevicesRead, DevicesWrite, InterfacesRead, InterfacesWrite, NetworksRead, NetworksWrite, OrganizationRead, OrganizationWrite, OrganizationMembersRead, OrganizationMembersWrite, RolesRead, RolesWrite },
	"operator":   { OrganizationMembersRead, DevicesRead, DevicesWrite, InterfacesRead, InterfacesWrite, NetworksRead, NetworksWrite, OrganizationRead, UsersRead },
	"viewer": { DevicesRead, InterfacesRead, NetworksRead, OrganizationRead, UsersRead },
}