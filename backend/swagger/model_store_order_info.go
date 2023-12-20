/*
 * PUN street Universal Access - OpenAPI 3.0
 *
 * pua
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type StoreOrderInfo struct {
	StoreId int64 `json:"store_id"`

	StoreName string `json:"store_name"`

	StoreShippingFee int64 `json:"store_shipping_fee"`

	ProductOrder []ProductOrderInfo `json:"product_order"`

	ShippingDiscount *ShippingDiscount `json:"shipping_discount"`

	ShippingDiscountBool bool `json:"shipping_discount_bool"`

	SeasoningDiscount *SeasoningDiscount `json:"seasoning_discount"`

	SeasoningDiscountBool bool `json:"seasoning_discount_bool"`

	TotalPrice int64 `json:"total_price"`
}
