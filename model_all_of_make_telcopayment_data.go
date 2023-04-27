/*
 * Bloc API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package bloc

type AllOfMakeTelcopaymentData struct {
	CreatedAt string `json:"created_at"`
	Status string `json:"status"`
	Amount int32 `json:"amount"`
	Reference string `json:"reference"`
	CustomerName string `json:"customer_name"`
	OperatorId string `json:"operator_id"`
	ProductId string `json:"product_id"`
	BillType string `json:"bill_type"`
	MetaData *Object `json:"meta_data"`
}