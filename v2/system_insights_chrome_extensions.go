/*
 * JumpCloud APIs
 *
 *  JumpCloud's V2 API. This set of endpoints allows JumpCloud customers to manage objects, groupings and mappings and interact with the JumpCloud Graph.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package v2

type SystemInsightsChromeExtensions struct {

	Uid string `json:"uid,omitempty"`

	Name string `json:"name,omitempty"`

	Identifier string `json:"identifier,omitempty"`

	Version string `json:"version,omitempty"`

	Description string `json:"description,omitempty"`

	Locale string `json:"locale,omitempty"`

	UpdateUrl string `json:"update_url,omitempty"`

	Author string `json:"author,omitempty"`

	Persistent int32 `json:"persistent,omitempty"`

	Path string `json:"path,omitempty"`

	Permissions string `json:"permissions,omitempty"`

	JcCollectionTime string `json:"jc_collection_time,omitempty"`

	JcSystemId string `json:"jc_system_id,omitempty"`
}
