/*
 * Bloc API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package bloc

type CardsDisputeBody struct {
	TransactionId string `json:"transaction_id"`
	Reason string `json:"reason"`
	Explanation string `json:"explanation"`
	MetaData *Object `json:"meta_data"`
}