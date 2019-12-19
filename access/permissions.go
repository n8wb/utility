/*
	Copyright 2019 whiteblock Inc.
	This file is a part of the utility.

	Utility is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	Utility is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package access

type Role string

const (
	Owner Role = "owner"
	Admin Role = "admin"
	User  Role = "user"
	Guest Role = "guest"
)

const (
	Owners string = "owners"
	Admins string = "admins"
	Users  string = "users"
	Guests string = "guests"
)

func GetGroupForRole(role Role) string {
	groups := make(map[Role]string)
	groups[Guest] = Guests
	groups[User] = Users
	groups[Admin] = Admins
	groups[Owner] = Owners
	return groups[role]
}

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

	DeleteOrgInvite
	ViewOrgInvite

	ViewOrgMembers

	CreateBiome
	DestroyBiome

	RunTestnet
	StopTestnet
	DeleteTestnet

	ViewPreviousTestnets
	CreateOrgInvite
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
		if IsOneOfThese(perm, MakeOrgAdmin, MakeOrgUser, ModifyOrgUser, UpdateOrg, DeleteOrgInvite, CreateOrgInvite) {
			return true
		}
		fallthrough
	case User:
		if IsOneOfThese(perm, ModifyOrgGuest, MakeOrgGuest, ViewOrgInvite) {
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
