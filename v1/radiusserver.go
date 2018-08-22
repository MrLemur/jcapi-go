/*
 * JumpCloud APIs
 *
 * V1 & V2 versions of JumpCloud's API. The previous version of JumpCloud's API. This set of endpoints allows JumpCloud customers to manage commands, systems, & system users.
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package v1

type Radiusserver struct {

	Id string `json:"_id,omitempty"`

	Organization string `json:"organization,omitempty"`

	NetworkSourceIp string `json:"networkSourceIp,omitempty"`

	SharedSecret string `json:"sharedSecret,omitempty"`

	Name string `json:"name,omitempty"`

	Tags []string `json:"tags,omitempty"`

	TagNames []string `json:"tagNames,omitempty"`
}
