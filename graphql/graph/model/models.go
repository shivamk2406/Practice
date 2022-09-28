package model

//go:generate go get github.com/99designs/gqlgen generate

type Tenant struct {
	ID              string         `json:"id,omitempty"`
	RecordID        string         `json:"record_id,omitempty"`
	DisplayName     string         `json:"display_name,omitempty"`
	Title           string         `json:"title,omitempty"`
	Name            string         `json:"name,omitempty"`
	ChannelName     string         `json:"channel_name,omitempty"`
	DefaultLanguage string         `json:"default_language,omitempty"`
	PaymentURI      string         `json:"payment_uri,omitempty"`
	Country         Country        `json:"country,omitempty"`
	SdkAttributes   SDKAttributes  `json:"sdk_attributes,omitempty"`
	Categories      []Category     `json:"categories,omitempty"`
	CheckoutConfig  CheckoutConfig `json:"checkout_config,omitempty"`
}
type Category struct {
	CategoryID                 string  `json:"category_id"`
	Name                       string  `json:"name"`
	UserTransactionCountLimit  int     `json:"user_transaction_count_limit"`
	UserTransactionAmountLimit float64 `json:"user_transaction_amount_limit"`
}

type CheckoutConfig struct {
	CartRefundToWallet   bool                 `json:"cart_refund_to_wallet"`
	ConvenienceFeeConfig ConvenienceFeeConfig `json:"convenience_fee_config"`
	PaymentModes         []PaymentMode        `json:"payment_modes"`
	DefaultPaymentMode   PaymentMode          `json:"default_payment_mode"`
}

type ConvenienceFeeConfig struct {
	ExecutionScript string `json:"execution_script"`
	Title           string `json:"title"`
	SubText         string `json:"sub_text"`
	PopupText       string `json:"popup_text"`
}

type Country struct {
	ID             string `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	Fullname       string `json:"fullname,omitempty"`
	CurrencySymbol string `json:"currency_symbol,omitempty"`
	CountryCode    string `json:"country_code,omitempty"`
	CurrencyCode   string `json:"currency_code,omitempty"`
}

type MapItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type PaymentMode struct {
	PaymentModeID       string    `json:"payment_mode_id"`
	Name                string    `json:"name"`
	LanguageDisplayText []MapItem `json:"language_display_text"`
	ImageURL            string    `json:"image_url"`
}

type SDKAttributes struct {
	WalletEnabled bool `json:"wallet_enabled,omitempty"`
}
