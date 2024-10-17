package resourcename

import (
	"github.com/shopspring/decimal"
)


const(
	blankOrderNameFormat        = "carts/{cart}/orders/-"
	redisProductKey           = "arise-b2b-saleor-pid-%s-%s"
	variantResourceName = "products/{product}/variants/{variant}"
)

type Price struct {
	BasePrice          decimal.Decimal `json:"base_price,omitempty"`
	PayablePrice       decimal.Decimal `json:"total_price,omitempty"`
	Tax                decimal.Decimal `json:"tax,omitempty"`
	Discount           decimal.Decimal `json:"discount,omitempty"`
	PotentialSaving    decimal.Decimal `json:"potential_saving,omitempty"`
	CurrencyCode       string          `json:"currency_code,omitempty"`
	DiscountPercentage float64         `json:"discount_percentage,omitempty"`
	TaxPercentage      float64         `json:"tax_percentage,omitempty"`
	SavingsPercentage  float64         `json:"savings_percentage,omitempty"`
	ExtraCharges       decimal.Decimal `json:"extra_charges,omitempty"`
}

func Start() error {
	//fmt.Println(resourcename.Sprint(blankOrderNameFormat,"1"))
	// product:="products/UHJvZHVjdDoxNjU=/variants/UHJvZHVjdFZhcmlhbnQ6NDAz"


	// var pid string
	// var channelName string

	// err:= resourcename.Sscan(product,variantResourceName,&pid,&channelName)
	// if err!=nil{
	// 	fmt.Println("inside error")
	// 	fmt.Println(err)
	// 	//fmt.Println(n)
	// }
	// fmt.Println(pid,channelName)
	return nil
}