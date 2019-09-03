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


//HasPermission checks whether the given role has the given permission
func HasPermission(role Role, perm Permission) bool {

	switch role {
	case Owner:
		if isOneOfThese(perm, MakeOrgOwner, DeleteOrg, ModifyOrgAdmin, ModifyOrgOwner) {
			return true
		}
		fallthrough
	case Admin:
		if isOneOfThese(perm, MakeOrgAdmin, MakeOrgUser, ModifyOrgUser, UpdateOrg, DeleteOrgRequest) {
			return true
		}
		fallthrough
	case User:
		if isOneOfThese(perm, ModifyOrgGuest, MakeOrgGuest, ViewOrgRequest) {
			return true
		}
		fallthrough
	case Guest:
		if isOneOfThese(perm, ViewOrgMembers) {
			return true
		}
	}
	return false
}
