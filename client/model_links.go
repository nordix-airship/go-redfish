/*
 * Redfish OAPI specification
 *
 * Partial Redfish OAPI specification for a limited client
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Links struct {
	Chassis []IdRef `json:"Chassis,omitempty"`
	ManagedBy []IdRef `json:"ManagedBy,omitempty"`
}