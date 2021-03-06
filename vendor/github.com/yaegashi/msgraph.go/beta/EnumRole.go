// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// RoleAssignmentScopeType undocumented
type RoleAssignmentScopeType string

const (
	// RoleAssignmentScopeTypeVResourceScope undocumented
	RoleAssignmentScopeTypeVResourceScope RoleAssignmentScopeType = "resourceScope"
	// RoleAssignmentScopeTypeVAllDevices undocumented
	RoleAssignmentScopeTypeVAllDevices RoleAssignmentScopeType = "allDevices"
	// RoleAssignmentScopeTypeVAllLicensedUsers undocumented
	RoleAssignmentScopeTypeVAllLicensedUsers RoleAssignmentScopeType = "allLicensedUsers"
	// RoleAssignmentScopeTypeVAllDevicesAndLicensedUsers undocumented
	RoleAssignmentScopeTypeVAllDevicesAndLicensedUsers RoleAssignmentScopeType = "allDevicesAndLicensedUsers"
)

var (
	// RoleAssignmentScopeTypePResourceScope is a pointer to RoleAssignmentScopeTypeVResourceScope
	RoleAssignmentScopeTypePResourceScope = &_RoleAssignmentScopeTypePResourceScope
	// RoleAssignmentScopeTypePAllDevices is a pointer to RoleAssignmentScopeTypeVAllDevices
	RoleAssignmentScopeTypePAllDevices = &_RoleAssignmentScopeTypePAllDevices
	// RoleAssignmentScopeTypePAllLicensedUsers is a pointer to RoleAssignmentScopeTypeVAllLicensedUsers
	RoleAssignmentScopeTypePAllLicensedUsers = &_RoleAssignmentScopeTypePAllLicensedUsers
	// RoleAssignmentScopeTypePAllDevicesAndLicensedUsers is a pointer to RoleAssignmentScopeTypeVAllDevicesAndLicensedUsers
	RoleAssignmentScopeTypePAllDevicesAndLicensedUsers = &_RoleAssignmentScopeTypePAllDevicesAndLicensedUsers
)

var (
	_RoleAssignmentScopeTypePResourceScope              = RoleAssignmentScopeTypeVResourceScope
	_RoleAssignmentScopeTypePAllDevices                 = RoleAssignmentScopeTypeVAllDevices
	_RoleAssignmentScopeTypePAllLicensedUsers           = RoleAssignmentScopeTypeVAllLicensedUsers
	_RoleAssignmentScopeTypePAllDevicesAndLicensedUsers = RoleAssignmentScopeTypeVAllDevicesAndLicensedUsers
)

// RoleSummaryStatus undocumented
type RoleSummaryStatus string

const (
	// RoleSummaryStatusVOk undocumented
	RoleSummaryStatusVOk RoleSummaryStatus = "ok"
	// RoleSummaryStatusVBad undocumented
	RoleSummaryStatusVBad RoleSummaryStatus = "bad"
)

var (
	// RoleSummaryStatusPOk is a pointer to RoleSummaryStatusVOk
	RoleSummaryStatusPOk = &_RoleSummaryStatusPOk
	// RoleSummaryStatusPBad is a pointer to RoleSummaryStatusVBad
	RoleSummaryStatusPBad = &_RoleSummaryStatusPBad
)

var (
	_RoleSummaryStatusPOk  = RoleSummaryStatusVOk
	_RoleSummaryStatusPBad = RoleSummaryStatusVBad
)
