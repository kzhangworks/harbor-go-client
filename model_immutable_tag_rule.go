/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ImmutableTagRule struct {
	Id int64 `json:"id,omitempty"`
	ProjectId int64 `json:"project_id,omitempty"`
	TagFilter string `json:"tag_filter,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
}