/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Insufficient storage
type InsufficientStorageChartApiError struct {
	// The error message returned by the chart API
	Error_ string `json:"error"`
}