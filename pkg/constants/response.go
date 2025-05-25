package constants

const (
	deliverCampaignsSuccessCode string = "2001"
	deliverCampaignsErrorCode   string = "4001"

	deliverCampaignsSuccessMessage string = "Campaigns fetched successfully"
	malformedRequestMessage        string = "Malformed request"
	defaultErrorMessage            string = "Something went wrong!"
)

var (
	DeliverCampaignsSuccess          = response{code: deliverCampaignsSuccessCode, message: deliverCampaignsSuccessMessage}
	DeliverCampaignsMalformedRequest = response{code: deliverCampaignsErrorCode, message: malformedRequestMessage}
	DeliverCampaignsError            = response{code: deliverCampaignsErrorCode, message: defaultErrorMessage}
)

type response struct {
	code    string
	message string
}

func (response response) GetSuccessCode() string {
	return "S01" + response.code
}

func (response response) GetErrorCode() string {
	return "E01" + response.code
}

func (response response) GetMessage() string {
	return response.message
}
