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
	case pInt > PrivilegeUserWrite|PrivilegeUserRead|PrivilegeUserDelete:
		s = "admin"
	case pInt > PrivilegeCategoryWrite|PrivilegePostWrite:
		s = "editor"
	case pInt > PrivilegeCategoryRead|PrivilegePostRead:
		s = "viewer"
	}

	return s

}
