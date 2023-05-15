package model

import "time"

type Tenant struct {
	ID                 string                 `json:"id,omitempty"`
	RecordID           string                 `json:"record_id,omitempty"`
	DisplayName        string                 `json:"display_name,omitempty"`
	Domain             string                 `json:"domain,omitempty"`
	Title              string                 `json:"title,omitempty"`
	Name               string                 `json:"name,omitempty"`
	TenantLogo         string                 `json:"tenant_logo,omitempty"`
	ChannelName        string                 `json:"channel_name,omitempty"`
	DefaultLanguage    string                 `json:"default_language,omitempty"`
	PaymentURI         string                 `json:"payment_uri,omitempty"`
	ProviderID         int                    `json:"provider_id,omitempty"`
	SupportEmail       string                 `json:"support_email,omitempty"`
	RedirectURI        string                 `json:"redirect_uri,omitempty"`
	Country            Country                `json:"country,omitempty"`
	SdkAttributes      SDKAttributes          `json:"sdk_attributes,omitempty"`
	Categories         []Category             `json:"categories,omitempty"`
	CheckoutConfig     CheckoutConfig         `json:"checkout_config,omitempty"`
	ConfigAttr         ConfigAttributes       `json:"config_attr,omitempty"`
	NotificationAttr   NotificationAttributes `json:"notification_attr,omitempty"`
	SubscriptionConfig SubscriptionConfig     `json:"subscription_config,omitempty"`
	OtpConfig          OtpConfig              `json:"otp_config,omitempty"`
	SupportAttributes  SupportAttributes      `json:"support_attributes,omitempty"`
	UserProfile        UserProfile            `json:"user_profile,omitempty"`
	RenewalConfig      RenewalConfig          `json:"renewal_config,omitempty"`
	ExchangeRates      []ExchangeRate         `json:"exchange_rates,omitempty"`
}

type SupportAttributes struct {
	TicketTagID string `json:"ticket_tag_id,omitempty"`
}

type SubscriptionConfig struct {
	MandateEndTime    int              `json:"mandate_end_time,omitempty"`
	MandateAmountType string           `json:"mandate_amount_type,omitempty"`
	ConfigAttr        ConfigAttributes `json:"config_attr,omitempty"`
}
type Category struct {
	CategoryID                 string  `json:"category_id"`
	Name                       string  `json:"name"`
	UserTransactionCountLimit  int     `json:"user_transaction_count_limit"`
	UserTransactionAmountLimit float64 `json:"user_transaction_amount_limit"`
}

type CheckoutConfig struct {
	CartRefundToWallet   bool                 `json:"cart_refund_to_wallet,omitempty"`
	ConvenienceFeeConfig ConvenienceFeeConfig `json:"convenience_fee_config,omitempty"`
	PaymentModes         []PaymentMode        `json:"payment_modes,omitempty"`
	DefaultPaymentMode   PaymentMode          `json:"default_payment_mode,omitempty"`
	PaymentProvider      PaymentProvider      `json:"payment_provider,omitempty"`
}

type ConvenienceFeeConfig struct {
	ExecutionScript string `json:"execution_script"`
	Title           string `json:"title"`
	SubText         string `json:"sub_text"`
	PopupText       string `json:"popup_text"`
}

type Country struct {
	ID                    string `json:"id,omitempty"`
	Name                  string `json:"name,omitempty"`
	FullName              string `json:"full_name,omitempty"`
	CurrencySymbol        string `json:"currency_symbol,omitempty"`
	CountryCode           string `json:"country_code,omitempty"`
	CurrencyCode          string `json:"currency_code,omitempty"`
	CurrencyName          string `json:"currency_name,omitempty"`
	IsdCode               string `json:"isd_code,omitempty"`
	CurrencyDecimals      int32  `json:"currency_decimals,omitempty"`
	FlagLogo              string `json:"flag_logo,omitempty" copier:"FlagLogoURL"`
	MobileValidationRegex string `json:"mobile_validation_regex,omitempty"`
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

type ConfigAttributes struct {
	TenantID                string         `json:"tenant_id,omitempty"`
	AnalyticsProjectID      string         `json:"analytics_project_id,omitempty"`
	ImageBaseURI            string         `json:"image_base_uri,omitempty"`
	PartnerTitle            string         `json:"partner_title,omitempty"`
	Locale                  string         `json:"locale,omitempty"`
	SupportEmail            string         `json:"support_email,omitempty"`
	BrowserTabIconURI       string         `json:"browser_tab_icon_uri,omitempty"`
	BrowserTabTitle         string         `json:"browser_tab_title,omitempty"`
	ThemeData               []MapItem      `json:"theme_data,omitempty"`
	HomeScreenBanner        Banner         `json:"home_screen_banner,omitempty"`
	OnboardingBanners       []Banner       `json:"onboarding_banners,omitempty"`
	EnableCountryCode       []string       `json:"enable_country_code,omitempty"`
	MobileOnboardingBanners []Banner       `json:"mobile_onboarding_banners,omitempty"`
	RecommendedSectionTitle string         `json:"recommended_section_title,omitempty"`
	RecommendedSectionTag   string         `json:"recommended_section_tag,omitempty"`
	PlatformFeatures        []Banner       `json:"platform_features,omitempty"`
	ProblemSection          ProblemSection `json:"problem_section,omitempty"`
	DisplayCategoryIDs      []string       `json:"display_category_ids,omitempty"`
}

type Banner struct {
	Title       string `json:"title,omitempty"`
	Subtitle    string `json:"subtitle,omitempty"`
	Description string `json:"description,omitempty"`
	ImageURI    string `json:"image_uri,omitempty"`
}

type PaymentProvider struct {
	ProviderName       string `json:"provider_name,omitempty"`
	ProviderID         int    `json:"provider_id,omitempty"`
	PaymentGatewayID   string `json:"payement_gateway_id,omitempty"`
	PaymentGatewayType string `json:"payment_gateway_type,omitempty"`
	AdditionalInfo     string `json:"additional_info,omitempty"`
	PaymentModel       string `json:"payment_model,omitempty"`
	SetMandate         bool   `json:"set_mandate,omitempty"`
}

type NotificationAttributes struct {
	RefundComment     string `json:"refund_comment,omitempty"`
	BgImage           string `json:"bg_image,omitempty"`
	ContactImage      string `json:"contact_image,omitempty"`
	ThemePrimaryColor string `json:"theme_primary_color,omitempty"`
	Timezone          string `json:"timezone,omitempty"`
	SenderName        string `json:"sender_name,omitempty"`
	SenderEmail       string `json:"sender_email,omitempty"`
}

type OtpConfig struct {
	MaxGenerateOtpAllowed int `json:"max_generate_otp_allowed,omitempty"`
	RetryWait             int `json:"retry_wait,omitempty"`
	WindowDuration        int `json:"window_duration,omitempty"`
	MaxVerifyAttempt      int `json:"max_verify_attempt,omitempty"`
	ValidityDurationInMin int `json:"validity_duration_in_min,omitempty"`
}

type State struct {
	ID        int     `json:"record_id"`
	StateName *string `json:"state_name"`
}

type UserProfile struct {
	ProfileSchemaID string `json:"profile_schema_id,omitempty"`
}

type RenewalConfig struct {
	AutoPayBannerMsg   string `json:"auto_pay_banner_msg,omitempty"`
	ManualPayBannerMsg string `json:"manual_pay_banner_msg,omitempty"`
	MandateAmountLimit int    `json:"mandate_amount_limit,omitempty"`
	SendInvoice        bool   `json:"send_invoice,omitempty"`
}

type ProblemSection struct {
	Title      string    `json:"title,omitempty"`
	Subtitle   string    `json:"subtitle,omitempty"`
	DataPoints []MapItem `json:"data_points,omitempty"`
}

type ExchangeRate struct {
	ID           string    `json:"id,omitempty"`
	FromCurrency string    `json:"from_currency,omitempty"`
	ToCurrency   string    `json:"to_currency,omitempty"`
	StartAt      time.Time `json:"start_at,omitempty"`
	ExpireAt     time.Time `json:"expire_at,omitempty"`
	ExchangeRate float64   `json:"exchange_rate,omitempty"`
}
