package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/shivamk2406/Practice/graphql/graph/model"
)

var (
	DisplayName     = "DisplayName"
	Title           = "title"
	Name            = "Name"
	ChannelName     = "Channel Name"
	DefaultLanguage = "English"
	FullName        = "Full Name"
)

func GenerateJSON() {
	TenantData := model.Tenant{
		ID:              "2",
		RecordID:        "1",
		DisplayName:     "IndusInd",
		Title:           "Indusland-India",
		Name:            "Indusland",
		ChannelName:     "Nuclei-Staging",
		DefaultLanguage: "English",
		PaymentURI:      "https://arise-dev.gonuclei.com/mock-payment",
		Country: model.Country{
			ID:             "356",
			Name:           "India",
			Fullname:       "Republic of India",
			CurrencySymbol: "₹",
			CountryCode:    "IN",
			CurrencyCode:   "INR",
		},

		SdkAttributes: model.SDKAttributes{
			WalletEnabled: true,
		},
		Categories: []model.Category{
			{
				CategoryID:                 "1",
				Name:                       "Category-1",
				UserTransactionCountLimit:  100,
				UserTransactionAmountLimit: 2000,
			}, {
				CategoryID:                 "2",
				Name:                       "Category-2",
				UserTransactionCountLimit:  200,
				UserTransactionAmountLimit: 2000,
			},
		},
		CheckoutConfig: model.CheckoutConfig{
			CartRefundToWallet: false,
			ConvenienceFeeConfig: model.ConvenienceFeeConfig{
				ExecutionScript: "",
				Title:           "Convenience Fee",
				SubText:         "",
				PopupText:       "",
			},
			PaymentModes: []model.PaymentMode{
				{
					PaymentModeID: "1",
					Name:          "UPI",
					LanguageDisplayText: []model.MapItem{
						{
							Key:   "en",
							Value: "English",
						},
					},
					ImageURL: "Image Url",
				},
				{
					PaymentModeID: "2",
					Name:          "CREDIT_CARD",
					LanguageDisplayText: []model.MapItem{
						{
							Key:   "HN",
							Value: "HINDI",
						},
					},
					ImageURL: "Image URL",
				},
			},
			DefaultPaymentMode: model.PaymentMode{
				PaymentModeID: "1",
				Name:          "UPI",
				LanguageDisplayText: []model.MapItem{
					{
						Key:   "en",
						Value: "English",
					},
				},
				ImageURL: "Image Url",
			},
		},
	}

	fmt.Println("Tenant Proto Data")
	fmt.Println(TenantData)
	fmt.Println("---------------------------------------------------------------")
	byte, err := json.Marshal(TenantData)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(byte))
	fmt.Println("After Unmarshalling")
	var res model.Tenant
	err = json.Unmarshal(byte, &res)
	fmt.Println(res)
}

func main() {
	GenerateJSON()
}
