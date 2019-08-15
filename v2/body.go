/*
 * JumpCloud APIs
 *
 *  JumpCloud's V2 API. This set of endpoints allows JumpCloud customers to manage objects, groupings and mappings and interact with the JumpCloud Graph.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package v2

type Body struct {

	Id string `json:"id,omitempty"`

	UserLockoutAction *LdapServerAction `json:"userLockoutAction,omitempty"`

	UserPasswordExpirationAction *LdapServerAction `json:"userPasswordExpirationAction,omitempty"`
}
