/*
 * Redfish OAPI specification
 *
 * Partial Redfish OAPI specification for a limited client
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// InsertMediaRequestBody struct for InsertMediaRequestBody
type InsertMediaRequestBody struct {
	Image string `json:"Image"`
	Inserted bool `json:"Inserted,omitempty"`
	Password string `json:"Password,omitempty"`
	TransferMethod TransferMethod `json:"TransferMethod,omitempty"`
	TransferProtocolType TransferProtocolType `json:"TransferProtocolType,omitempty"`
	UserName string `json:"UserName,omitempty"`
	WriteProtected bool `json:"WriteProtected,omitempty"`
}
