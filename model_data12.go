/*
 * Bloc API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package bloc

type Data12 struct {
	Category string `json:"category"`
	Desc string `json:"desc"`
	FeeType string `json:"fee_type"`
	Id string `json:"id"`
	Meta *AllOfData12Meta `json:"meta"`
	Name string `json:"name"`
	Operator string `json:"operator"`
}