/*
 * Redfish OAPI specification
 *
 * Partial Redfish OAPI specification for a limited client
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type ProcessorSummary struct {
	Count int32 `json:"Count,omitempty"`
	Status Status `json:"Status,omitempty"`
}
