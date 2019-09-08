/*
 * Redfish OAPI specification
 *
 * Partial Redfish OAPI specification for a limited client
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
type ResetType string

// List of ResetType
const (
	ON ResetType = "On"
	FORCE_OFF ResetType = "ForceOff"
	GRACEFUL_SHUTDOWN ResetType = "GracefulShutdown"
	GRACEFUL_RESTART ResetType = "GracefulRestart"
	FORCE_RESTART ResetType = "ForceRestart"
	NMI ResetType = "Nmi"
	FORCE_ON ResetType = "ForceOn"
	PUSH_POWER_BUTTON ResetType = "PushPowerButton"
	POWER_CYCLE ResetType = "PowerCycle"
)