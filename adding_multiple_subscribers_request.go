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

type AddingMultipleSubscribersRequest struct {

	// A list of subscribers to add to the mailing list. You may specify the email address, the name and the custom fields for each subscriber. The following example demonstrates the format each subscriber should be written in your request, assuming that you also have two custom fields named Age and Country.
	Subscribers []Subscribers `json:"Subscribers"`
}
