# Sendtext API SDK for Go

## Introduction

This Go SDK enables developers to easily integrate Sendtext API functionalities into their applications. The Sendtext API provides services for sending SMS, managing SMS campaigns, checking SMS balance, and viewing SMS sending history.

## Installation

To use this SDK, include it in your Go project:

```bash
go get -u github.com/abdotop/go-sendtext
```

## Configuration

Start by creating a client instance using your API keys:

```go
    import "github.com/abdotop/go-sendtext"

    client := sendtext.NewClient("your_api_key", "your_api_secret")
```

## Data Structures

### SMSRequest

Represents the request to send a single SMS.

```go
    type SMSRequest struct {
        SenderName   string `json:"sender_name"`
        SMSType      string `json:"sms_type"`
        Phone        string `json:"phone"`
        ScheduledAt  string `json:"scheduled_at,omitempty"`
        Text         string `json:"text"`
    }
```

### SMSResponse

Represents the response from the Sendtext API after sending an SMS.

```go
    type SMSResponse struct {
        Type             string `json:"type"`
        Text             string `json:"text"`
        Phone            string `json:"phone"`
        SenderName       string `json:"senderName"`
        MessageID        string `json:"messageId"`
        ScheduledAt      string `json:"scheduledAt"`
        SendtextSmsCount int    `json:"sendtextSmsCount"`
        StatusID         int    `json:"statusId"`
        StatusDescription string `json:"statusDescription"`
    }
```

### CampaignRequest

Represents the payload for sending a bulk SMS campaign.

```go
    type CampaignRequest struct {
        SenderName      string `json:"sender_name"`
        SMSType         string `json:"sms_type"`
        ScheduledAt     string `json:"scheduled_at,omitempty"`
        CampaignName    string `json:"campaign_name"`
        CampaignLines   []struct {
            Phone string `json:"phone"`
            Text  string `json:"text"`
        } `json:"campaign_lines"`
    }
```

### CampaignResponse

Represents the response from the Sendtext API after sending a bulk SMS campaign.

```go
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
```

### BalanceResponse

Represents the response from the Sendtext API for a balance check request.

```go
    type BalanceResponse struct {
	    Balance   int    `json:"balance"`
	    ExpiresAt string `json:"expires_at"`
	    UpdatedAt string `json:"updated_at"`
    }
```

### HistoryItem

Represents a single SMS transaction in the history.

```go
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
```

### HistoryResponse

Is a list of HistoryItem, representing the complete SMS sending history.

```go
    type HistoryResponse []HistoryItem
```

## SDK Methods

### Send an SMS

Send an SMS to a specific number.

```go
    smsReq := sendtext.SMSRequest{
        SenderName: "Bonlux SN",
        SMSType: "normal",
        Phone: "221763983535",
        Text: "Bonjour. Votre code d'activation est 3725.",
    }

    response, err := client.SendSMS(smsReq)
    if err != nil {
        log.Fatalf("Error sending SMS: %s", err)
    }
    fmt.Printf("Response: %+v\n", response)
```

### Send an SMS Campaign

Send an SMS campaign to multiple numbers.

```go
    campaignReq := sendtext.CampaignRequest{
        SenderName: "Ecole SN",
        SMSType: "flash",
        CampaignName: "Admission 2025",
        CampaignLines: []struct{
            Phone string `json:"phone"`
            Text  string `json:"text"`
        }{
            {Phone: "221701234567", Text: "Bonjour Monsieur Sow. Votre enfant a été admis à l'école SN."},
            {Phone: "221761234567", Text: "Bonjour Madame Diop. Votre enfant a été admis à l'école SN."},
        },
    }

    campaignResp, err := client.SendCampaign(campaignReq)
    if err != nil {
        log.Fatalf("Error sending campaign: %s", err)
    }
    fmt.Printf("Campaign Response: %+v\n", campaignResp)
```

### Check SMS Balance

Check the available SMS balance.

```go
    balance, err := client.CheckBalance()
    if err != nil {
        log.Fatalf("Error checking balance: %s", err)
    }
    fmt.Printf("Balance: %+v\n", balance)
```

### Retrieve SMS Sending History

Get the history of sent SMS messages.

```go
    history, err := client.GetSMSHistory()
    if err != nil {
        log.Fatalf("Error retrieving SMS history: %s", err)
    }
    fmt.Printf("SMS History: %+v\n", history)
```

## Conclusion

This SDK provides a straightforward interface to integrate Sendtext API functionalities into your Go applications. Use the provided methods to send SMS, manage campaigns, and check your balance and sending history.
