/**
 * Copyright 2019 Whiteblock Inc. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
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
