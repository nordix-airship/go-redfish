/*
 * Redfish OAPI specification
 *
 * Partial Redfish OAPI specification for a limited client
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// RedfishError Contains an error payload from a Redfish Service.
type RedfishError struct {
	Error RedfishErrorError `json:"error"`
}
