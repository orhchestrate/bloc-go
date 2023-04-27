/*
 * Bloc API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package bloc

type AllOfTransactionbyreferenceData struct {
	Id string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Amount int32 `json:"amount"`
	Reference string `json:"reference"`
	Status string `json:"status"`
	Shared bool `json:"shared"`
	Currency string `json:"currency"`
	Environment string `json:"environment"`
	PaymentMethod string `json:"payment_method"`
	Provider string `json:"provider"`
	PaymentType string `json:"payment_type"`
	Source string `json:"source"`
	MetaData string `json:"meta_data"`
	OrganizationId string `json:"organization_id"`
	CustomerId string `json:"customer_id"`
	Fee int32 `json:"fee"`
	BillingId string `json:"billing_id"`
	CustomerDetail *Object `json:"customer_detail"`
	Reversal bool `json:"reversal"`
	PreviousAccountBalance int32 `json:"previous_account_balance"`
	CurrentAccountBalance int32 `json:"current_account_balance"`
	AccountId string `json:"account_id"`
	Drcr string `json:"drcr"`
}