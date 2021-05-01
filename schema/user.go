package schema

const (
	ProjectAdmin = iota + 1
	Developer
	Guest
	Maintainer
)

const (
	LDAPGroup = iota + 1
	HTTPGroup
)

type ProjectMember struct {
	RoleID      int        `json:"role_id,omitempty"`
	MemberGroup UserGroup  `json:"member_group"`
	MemberUser  UserEntity `json:"member_user"`
}

type UserGroup struct {
	GroupName   string `json:"group_name,omitempty"`
	LdapGroupDN string `json:"ldap_group_dn,omitempty"`
	GroupType   int    `json:"group_type,omitempty"`
	ID          int64  `json:"id,omitempty"`
}

type UserEntity struct {
	Username string `json:"username,omitempty"`
	UserID   int64  `json:"user_id,omitempty"`
}

type QueryUserOptions struct {
	Name     string
	Page     string
	PageSize string
}
