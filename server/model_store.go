/*
 * Takí Project
 *
 * IF683 Takí project
 *
 * API version: 1.0.0
 * Contact: mvgmb@cin.ufpe.br
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type Store struct {

	// Unique identifier representing a specific Store
	Id int32 `json:"_id,omitempty"`

	// Name of store
	Name string `json:"name,omitempty"`

	Map_ *ModelMap `json:"map,omitempty"`
}
