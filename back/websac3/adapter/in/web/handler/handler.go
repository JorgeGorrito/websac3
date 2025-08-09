package handler

import "slices"

type Authenticable struct {
	PermissionsRequired []string
}

func (a *Authenticable) ValidatePermissions(permissionsGot []string) bool {
	var hasPermission bool = true
	for _, permissionRequired := range a.PermissionsRequired {
		hasPermission = hasPermission && slices.Contains(permissionsGot, permissionRequired)
	}
	return hasPermission
}
