/*
 * Bloc API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package bloc

type ResolvedAccount struct {
	Success bool `json:"success"`
	Data *AllOfResolvedAccountData `json:"data"`
	Message string `json:"message"`
}