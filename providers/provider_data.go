package providers

import (
	"net/url"
)

// TODO - add extra fields to capture google stuff
type ProviderData struct {
	ProviderName             string
	ClientID                 string
	ClientSecret             string
	LoginURL                 *url.URL
	RedeemURL                *url.URL
	ProfileURL               *url.URL
	ProtectedResource        *url.URL
	ValidateURL              *url.URL
	Scope                    string
	ApprovalPrompt           string
	GoogleAdminEmail         string
	GoogleServiceAccountJSON string
	GoogleSkipGroupAuth      bool
}

func (p *ProviderData) Data() *ProviderData { return p }
