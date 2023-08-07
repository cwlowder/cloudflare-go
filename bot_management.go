package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// AuthIdCharacteristics a single option from
// configuration?properties=auth_id_characteristics.
type BotManagement struct {
	EnableJS                     *bool   `json:"enable_js,omitempty"`
	FightMode                    *bool   `json:"fight_mode,omitempty"`
	SBFMDefinitelyAutomated      *string `json:"sbfm_definitely_automated,omitempty"`
	SBFMLikelyAutomated          *string `json:"sbfm_likely_automated,omitempty"`
	SBFMVerifiedBots             *string `json:"sbfm_verified_bots,omitempty"`
	SBFMStaticResourceProtection *bool   `json:"sbfm_static_resource_protection,omitempty"`
	OptimizeWordpress            *bool   `json:"optimize_wordpress,omitempty"`
	SuppressSessionScore         *bool   `json:"suppress_session_score,omitempty"`
	AutoUpdateModel              *bool   `json:"auto_update_model,omitempty"`
	UsingLatestModel             *bool   `json:"using_latest_model,omitempty"`
}

// BotManagementResponse represents the response from the bot_management endpoint.
type BotManagementResponse struct {
	Result BotManagement `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

type UpdateBotManagementParams struct {
	EnableJS                     *bool   `json:"enable_js,omitempty"`
	FightMode                    *bool   `json:"fight_mode,omitempty"`
	SBFMDefinitelyAutomated      *string `json:"sbfm_definitely_automated,omitempty"`
	SBFMLikelyAutomated          *string `json:"sbfm_likely_automated,omitempty"`
	SBFMVerifiedBots             *string `json:"sbfm_verified_bots,omitempty"`
	SBFMStaticResourceProtection *bool   `json:"sbfm_static_resource_protection,omitempty"`
	OptimizeWordpress            *bool   `json:"optimize_wordpress,omitempty"`
	SuppressSessionScore         *bool   `json:"suppress_session_score,omitempty"`
	AutoUpdateModel              *bool   `json:"auto_update_model,omitempty"`
}

// GetAPIShieldConfiguration gets a zone API shield configuration.
//
// API documentation: https://api.cloudflare.com/#api-shield-settings-retrieve-information-about-specific-configuration-properties
func (api *API) GetBotManagement(ctx context.Context, rc *ResourceContainer) (BotManagement, ResultInfo, error) {
	uri := fmt.Sprintf("/zones/%s/bot_management", rc.Identifier)

	res, err := api.makeRequestContextWithHeaders(ctx, http.MethodGet, uri, nil, botV2Header())
	if err != nil {
		return BotManagement{}, ResultInfo{}, err
	}
	var asResponse BotManagementResponse
	err = json.Unmarshal(res, &asResponse)
	if err != nil {
		return BotManagement{}, ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return asResponse.Result, asResponse.ResultInfo, nil
}

// UpdateAPIShieldConfiguration sets a zone API shield configuration.
//
// API documentation: https://api.cloudflare.com/#api-shield-settings-set-configuration-properties
func (api *API) UpdateBotManagement(ctx context.Context, rc *ResourceContainer, params UpdateBotManagementParams) (BotManagement, error) {
	uri := fmt.Sprintf("/zones/%s/bot_management", rc.Identifier)

	res, err := api.makeRequestContextWithHeaders(ctx, http.MethodPut, uri, params, botV2Header())
	if err != nil {
		return BotManagement{}, err
	}

	var bmResponse BotManagementResponse
	err = json.Unmarshal(res, &bmResponse)
	if err != nil {
		return BotManagement{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return bmResponse.Result, nil
}

// We are currently undergoing the process of updating the bot management API.
// The older 1.0.0 version of the is still the default version, so we will need
// to explicitly set this special header on all requests. We will eventually
// make 2.0.0 the default version, and later we will remove the 1.0.0 entirely.
func botV2Header() http.Header {
	header := make(http.Header)
	header.Set("Cloudflare-Version", "2.0.0")

	return header
}
