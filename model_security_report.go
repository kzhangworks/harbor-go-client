/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// The security information of the chart
type SecurityReport struct {
	Signature *DigitalSignature `json:"signature,omitempty"`
}