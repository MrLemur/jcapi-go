/*
 * JumpCloud APIs
 *
 * V1 & V2 versions of JumpCloud's API. The previous version of JumpCloud's API. This set of endpoints allows JumpCloud customers to manage commands, systems, & system users.
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package v1

type ApplicationConfigConstantAttributesValue struct {

	Name string `json:"name,omitempty"`

	Value string `json:"value,omitempty"`

	Required bool `json:"required,omitempty"`

	Visible bool `json:"visible,omitempty"`

	ReadOnly bool `json:"readOnly,omitempty"`
}
