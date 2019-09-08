/*
 * Redfish OAPI specification
 *
 * Partial Redfish OAPI specification for a limited client
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
type PowerState string

// List of PowerState
const (
	TRUE PowerState = "true"
	FALSE PowerState = "false"
	POWERING_ON PowerState = "PoweringOn"
	POWERING_OFF PowerState = "PoweringOff"
)