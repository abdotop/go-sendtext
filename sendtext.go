package sendtext

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const baseURL = "https://api.sendtext.sn/v1"

// Client holds the API credentials and the HTTP client used for making requests.
type Client struct {
	APIKey     string
	APISecret  string
	httpClient *http.Client
}

// SMSRequest represents the payload for sending a single SMS.
type SMSRequest struct {
	SenderName  string `json:"sender_name"`
	SMSType     string `json:"sms_type"`
	Phone       string `json:"phone"`
	ScheduledAt string `json:"scheduled_at,omitempty"`
	Text        string `json:"text"`
}

// SMSResponse represents the response from the Sendtext API after sending an SMS.
type SMSResponse struct {
	Type              string `json:"type"`
	Text              string `json:"text"`
	Phone             string `json:"phone"`
	SenderName        string `json:"senderName"`
	MessageID         string `json:"messageId"`
	ScheduledAt       string `json:"scheduledAt"`
	SendtextSmsCount  int    `json:"sendtextSmsCount"`
	StatusID          int    `json:"statusId"`
	StatusDescription string `json:"statusDescription"`
}

// CampaignRequest represents the payload for sending a bulk SMS campaign.
type CampaignRequest struct {
	SenderName    string `json:"sender_name"`
	SMSType       string `json:"sms_type"`
	ScheduledAt   string `json:"scheduled_at,omitempty"`
	CampaignName  string `json:"campaign_name"`
	CampaignLines []struct {
		Phone string `json:"phone"`
		Text  string `json:"text"`
	} `json:"campaign_lines"`
}

// CampaignResponse represents the response from the Sendtext API after sending a bulk SMS campaign.
type CampaignResponse struct {
	Type          string `json:"type"`
	SenderName    string `json:"senderName"`
	ScheduledAt   string `json:"scheduledAt"`
	CampaignId    int    `json:"campaignId"`
	CampaignLines []struct {
		Text              string `json:"text"`
		Phone             string `json:"phone"`
		MessageID         string `json:"messageId"`
		SendtextSmsCount  int    `json:"sendtextSmsCount"`
		StatusId          int    `json:"statusId"`
		StatusDescription string `json:"statusDescription"`
	} `json:"campaignLines"`
}

// BalanceResponse represents the response from the Sendtext API for a balance check request.
type BalanceResponse struct {
	Balance   int    `json:"balance"`
	ExpiresAt string `json:"expires_at"`
	UpdatedAt string `json:"updated_at"`
}

// HistoryItem represents a single SMS transaction in the history.
type HistoryItem struct {
	Type              string `json:"type"`
	SenderName        string `json:"sendername"`
	MessageID         string `json:"messageId"`
	SendtextSmsCount  int    `json:"sendtextSmsCount"`
	Phone             string `json:"phone"`
	CreatedAt         string `json:"createdAt"`
	ScheduledAt       string `json:"scheduledAt"`
	Text              string `json:"text"`
	StatusId          int    `json:"statusId"`
	StatusDescription string `json:"statusDescription"`
}

// HistoryResponse is a list of HistoryItem, representing the complete SMS sending history.
type HistoryResponse []HistoryItem

// NewClient creates a new client with the provided API key and secret.
func NewClient(apiKey, apiSecret string) *Client {
	return &Client{
		APIKey:     apiKey,
		APISecret:  apiSecret,
		httpClient: &http.Client{},
	}
}

// SendSMS sends a single SMS using the Sendtext API.
func (c *Client) SendSMS(req SMSRequest) (*SMSResponse, error) {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", baseURL+"/sms", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("SNT-API-KEY", c.APIKey)
	request.Header.Set("SNT-API-SECRET", c.APISecret)

	resp, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var smsResp SMSResponse
	if err := json.NewDecoder(resp.Body).Decode(&smsResp); err != nil {
		return nil, err
	}

	return &smsResp, nil
}

// SendCampaign sends a bulk SMS campaign using the Sendtext API.
func (c *Client) SendCampaign(req CampaignRequest) (*CampaignResponse, error) {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", baseURL+"/bulk_sms", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("SNT-API-KEY", c.APIKey)
	request.Header.Set("SNT-API-SECRET", c.APISecret)

	resp, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var campaignResp CampaignResponse
	if err := json.NewDecoder(resp.Body).Decode(&campaignResp); err != nil {
		return nil, err
	}

	return &campaignResp, nil
}

// CheckBalance retrieves the current SMS balance from the Sendtext API.
func (c *Client) CheckBalance() (*BalanceResponse, error) {
	request, err := http.NewRequest("GET", baseURL+"/balance", nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("SNT-API-KEY", c.APIKey)
	request.Header.Set("SNT-API-SECRET", c.APISecret)

	resp, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var balanceResp BalanceResponse
	if err := json.NewDecoder(resp.Body).Decode(&balanceResp); err != nil {
		return nil, err
	}

	return &balanceResp, nil
}

// GetSMSHistory retrieves the history of sent SMS messages using the Sendtext API.
func (c *Client) GetSMSHistory() (*HistoryResponse, error) {
	request, err := http.NewRequest("GET", baseURL+"/history", nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("SNT-API-KEY", c.APIKey)
	request.Header.Set("SNT-API-SECRET", c.APISecret)

	resp, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var historyResp HistoryResponse
	if err := json.NewDecoder(resp.Body).Decode(&historyResp); err != nil {
		return nil, err
	}

	return &historyResp, nil
}
