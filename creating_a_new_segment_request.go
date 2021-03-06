/* 
 * Moosend API
 *
 * TODO: Add a description
 *
 * OpenAPI spec version: 1.0
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package Moosend.Wrappers.GOWrapper

type CreatingANewSegmentRequest struct {

	// The name of the segment.
	Name string `json:"Name"`

	// Specifies how the segment's criteria will match together. This must be one of the following values. If not specified, `All` will be assumed.
	MatchType string `json:"MatchType,omitempty"`
}
