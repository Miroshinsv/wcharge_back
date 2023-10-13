package entity

type Role struct {
	ID             int    `json:"id"`
	RoleName       string `json:"role_name"`
	RolePrivileges int    `json:"role_privileges"`
}
