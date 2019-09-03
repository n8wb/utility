package access

type Role string

const (
	Owner Role = "owner"
	Admin Role = "admin"
	User Role = "user"
	Guest Role = "guest"
)

type Permission int64

const (
	MakeOrgGuest Permission = iota + 1
	MakeOrgUser
	MakeOrgAdmin
	MakeOrgOwner

	ModifyOrgGuest
	ModifyOrgUser
	ModifyOrgAdmin
	ModifyOrgOwner

	UpdateOrg
	DeleteOrg

	DeleteOrgRequest
	ViewOrgRequest

	ViewOrgMembers

	CreateBiome
	DestroyBiome
)

//IsOneOfThese checks to see if any of the other given permissions match the first
//permission given
func IsOneOfThese(perm Permission, perms ...Permission) bool {
	for _, p := range perms {
		if p == perm {
			return true
		}
	}
	return false
}

//HasPermission checks whether the given role has the given permission
func HasPermission(role Role, perm Permission) bool {
	switch role {
	case Owner:
		if IsOneOfThese(perm, MakeOrgOwner, DeleteOrg, ModifyOrgAdmin, ModifyOrgOwner) {
			return true
		}
		fallthrough
	case Admin:
		if IsOneOfThese(perm, MakeOrgAdmin, MakeOrgUser, ModifyOrgUser, UpdateOrg, DeleteOrgRequest) {
			return true
		}
		fallthrough
	case User:
		if IsOneOfThese(perm, ModifyOrgGuest, MakeOrgGuest, ViewOrgRequest) {
			return true
		}
		fallthrough
	case Guest:
		if IsOneOfThese(perm, ViewOrgMembers) {
			return true
		}
	}
	return false
}
