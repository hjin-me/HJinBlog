package user

const (
	PrivilegePostRead int = 1 << iota
	PrivilegePostWrite
	PrivilegePostDelete
	PrivilegeUserRead
	PrivilegeUserWrite
	PrivilegeUserDelete
	PrivilegeCategoryRead
	PrivilegeCategoryWrite
	PrivilegeCategoryDelete
)

type Privilege int

func (p Privilege) String() string {
	s := "forbidden"
	pInt := int(p)
	switch {
	case 0 < pInt&(PrivilegeUserWrite|PrivilegeUserRead|PrivilegeUserDelete):
		s = "admin"
	case 0 < pInt&(PrivilegeCategoryWrite|PrivilegePostWrite):
		s = "editor"
	case 0 < pInt&(PrivilegeCategoryRead|PrivilegePostRead):
		s = "viewer"
	}

	return s

}
