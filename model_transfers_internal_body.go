/*
 * Bloc API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package bloc

type TransfersInternalBody struct {
	Amount int32 `json:"amount"`
	FromAccountId string `json:"from_account_id"`
	ToAccountId string `json:"to_account_id"`
	Narration string `json:"narration"`
	Reference string `json:"reference"`
	MetaData *interface{} `json:"meta_data"`
}