// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ConditionalAccessConditionSet undocumented
type ConditionalAccessConditionSet struct {
	// Object is the base model of ConditionalAccessConditionSet
	Object
	// Applications undocumented
	Applications *ConditionalAccessApplications `json:"applications,omitempty"`
	// Users undocumented
	Users *ConditionalAccessUsers `json:"users,omitempty"`
	// SignInRiskLevels undocumented
	SignInRiskLevels []RiskLevel `json:"signInRiskLevels,omitempty"`
	// Platforms undocumented
	Platforms *ConditionalAccessPlatforms `json:"platforms,omitempty"`
	// Locations undocumented
	Locations *ConditionalAccessLocations `json:"locations,omitempty"`
	// ClientAppTypes undocumented
	ClientAppTypes []ConditionalAccessClientApp `json:"clientAppTypes,omitempty"`
	// DeviceStates undocumented
	DeviceStates *ConditionalAccessDeviceStates `json:"deviceStates,omitempty"`
}