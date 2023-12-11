package klaviyo

import (
	"context"
	"fmt"
	"strings"
)

type ProfileServiceOp struct {
	client *Client
}

type ProfileService interface {
	Read(context.Context, ProfileRequest) (*ProfileResponse, error)
	Browse(context.Context, ProfileRequest) (*ProfilesResponse, error)
	Edit(context.Context, EditProfileRequest) (*ProfilesResponse, error)
	Create(context.Context, CreateProfile) (*ProfilesResponse, error)
}

type ProfileRequest struct {
	ID     string
	Emails *[]string
}

type EditProfileRequest struct {
	ID             string
	Emails         *[]string
	EditAttributes *EditAttributeRequest
	EditLocation   *EditLocationRequest
}

type EditAttributeRequest struct {
	Email        *string
	PhoneNumber  *string
	FirstName    *string
	LastName     *string
	Organization *string
	Title        *string
	Image        *string
}

type EditLocationRequest struct {
	Address1  *string
	Address2  *string
	City      *string
	Country   *string
	Latitude  *string
	Longitude *string
	Region    *string
	Zip       *string
	Timezone  *string
	IP        *string
}

type CreateProfile struct {
	Email        *string
	PhoneNumber  *string
	FirstName    *string
	LastName     *string
	Organization *string
	Title        *string
	Image        *string
	Properties   *map[string]string
	Location     *struct {
		Address1  *string
		Address2  *string
		City      *string
		Country   *string
		Latitude  *string
		Longitude *string
		Region    *string
		Zip       *string
		Timezone  *string
		IP        *string
	}
}

type ProfileResponse struct {
	Data     ProfileResponseData      `json:"data,omitempty"`
	Included []map[string]interface{} `json:"included,omitempty"`
}

type ProfilesResponse struct {
	Links map[string]string     `json:"links,omitempty"`
	Data  []ProfileResponseData `json:"data,omitempty"`
}

type ProfileResponseData struct {
	Type       string                     `json:"type,omitempty"`
	ID         string                     `json:"id,omitempty"`
	Links      map[string]string          `json:"links,omitempty"`
	Attributes *ProfileResponseAttributes `json:"attributes,omitempty"`
}

type ProfileResponseAttributes struct {
	Email               string                              `json:"email,omitempty"`
	PhoneNumber         string                              `json:"phone_number,omitempty"`
	ExternalID          string                              `json:"external_id,omitempty"`
	FirstName           string                              `json:"first_name,omitempty"`
	LastName            string                              `json:"last_name,omitempty"`
	Organization        string                              `json:"organization,omitempty"`
	Title               string                              `json:"title,omitempty"`
	Image               string                              `json:"image,omitempty"`
	Created             string                              `json:"created,omitempty"`
	Updated             string                              `json:"updated,omitempty"`
	LastEventDate       string                              `json:"last_event_date,omitempty"`
	Properties          map[string]string                   `json:"properties,omitempty"`
	Location            *ProfileResponseLocation            `json:"location,omitempty"`
	Subscriptions       *ProfileResponseSubscriptions       `json:"subscriptions,omitempty"`
	PredictiveAnalytics *ProfileResponsePredictiveAnalytics `json:"predictive_analytics,omitempty"`
}

type ProfileResponseLocation struct {
	Address1  string `json:"address1,omitempty"`
	Address2  string `json:"address2,omitempty"`
	City      string `json:"city,omitempty"`
	Country   string `json:"country,omitempty"`
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
	Region    string `json:"region,omitempty"`
	Zip       string `json:"zip,omitempty"`
	Timezone  string `json:"timezone,omitempty"`
	IP        string `json:"ip,omitempty"`
}

type ProfileResponseSubscriptions struct {
	Email *ProfileResponseSubscriptionEmail `json:"email,omitempty"`
	SMS   *ProfileResponseSubscriptionSMS   `json:"sms,omitempty"`
}

type ProfileResponsePredictiveAnalytics struct {
	HistoricCLV              int    `json:"historic_clv,omitempty"`
	PredictedCLV             int    `json:"predicted_clv,omitempty"`
	TotalCLV                 int    `json:"total_clv,omitempty"`
	HistoricNumberOfOrders   int    `json:"historic_number_of_orders,omitempty"`
	PredictedNumberOfOrders  int    `json:"predicted_number_of_orders,omitempty"`
	AverageDaysBetweenOrders int    `json:"average_days_between_orders,omitempty"`
	AverageOrderValue        int    `json:"average_order_value,omitempty"`
	ChurnProbability         int    `json:"churn_probability,omitempty"`
	ExpectedDateOfNextOrder  string `json:"expected_date_of_next_order,omitempty"`
}

type ProfileResponseSubscriptionEmail struct {
	Marketing struct {
		CanReceiveEmailMarketing string                                              `json:"can_receive_email_marketing,omitempty"`
		Consent                  string                                              `json:"consent,omitempty"`
		ConsentTimestamp         string                                              `json:"consent_timestamp,omitempty"`
		LastUpdated              string                                              `json:"last_updated,omitempty"`
		Method                   string                                              `json:"method,omitempty"`
		MethodDetail             string                                              `json:"method_detail,omitempty"`
		CustomMethodDetail       string                                              `json:"custom_method_detail,omitempty"`
		DoubleOptin              string                                              `json:"double_optin,omitempty"`
		Suppressions             *[]ProfileResponseSubscriptionEmailSuppressions     `json:"suppression,omitempty"`
		ListSuppressions         *[]ProfileResponseSubscriptionEmailListSuppressions `json:"list_suppressions,omitempty"`
	} `json:"marketing,omitempty"`
}

type ProfileResponseSubscriptionEmailSuppressions struct {
	Reason    string `json:"reason,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
}

type ProfileResponseSubscriptionEmailListSuppressions struct {
	ListID    string `json:"list_id,omitempty"`
	Reason    string `json:"reason,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
}

type ProfileResponseSubscriptionSMS struct {
	Marketing struct {
		CanReceiveSMSMarketing string `json:"can_receive_sms_marketing,omitempty"`
		Consent                string `json:"consent,omitempty"`
		ConsentTimestamp       string `json:"consent_timestamp,omitempty"`
		Method                 string `json:"method,omitempty"`
		MethodDetail           string `json:"method_detail,omitempty"`
		LastUpdated            string `json:"last_updated,omitempty"`
	} `json:"marketing,omitempty"`
}

const profileURL = "https://a.klaviyo.com/api/profiles/"

func (s *ProfileServiceOp) Read(ctx context.Context, params ProfileRequest) (*ProfileResponse, error) {

	var resp ProfileResponse

	url := profileURL + params.ID + "/"

	errRequest := s.client.Request("GET", url, strings.Reader{}, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}

func (s *ProfileServiceOp) Browse(ctx context.Context, params ProfileRequest) (*ProfilesResponse, error) {

	var resp ProfilesResponse

	errRequest := s.client.Request("GET", profileURL, strings.Reader{}, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}

func (s *ProfileServiceOp) Edit(ctx context.Context, params EditProfileRequest) (*ProfilesResponse, error) {

	var resp ProfilesResponse

	url := fmt.Sprintf("%v%v/", profileURL, params.ID)

	payloadBuild := []string{
		"\"type\":\"profile\"",
		"\"id\":\"" + params.ID + "\"",
	}
	if params.EditAttributes != nil {

		editAttributes := params.EditAttributes
		var attributeBody []string

		if editAttributes.Email != nil {
			attributeBody = append(attributeBody, fmt.Sprintf("\"email\":\"%v\"", *editAttributes.Email))
		}

		if editAttributes.PhoneNumber != nil {
			attributeBody = append(attributeBody, fmt.Sprintf("\"phone_number\":\"%v\"", *editAttributes.PhoneNumber))
		}

		if editAttributes.FirstName != nil {
			attributeBody = append(attributeBody, fmt.Sprintf("\"first_name\":\"%v\"", *editAttributes.FirstName))
		}

		if editAttributes.LastName != nil {
			attributeBody = append(attributeBody, fmt.Sprintf("\"last_name\":\"%v\"", *editAttributes.LastName))
		}

		if editAttributes.Organization != nil {
			attributeBody = append(attributeBody, fmt.Sprintf("\"organization\":\"%v\"", *editAttributes.Organization))
		}

		if editAttributes.Title != nil {
			attributeBody = append(attributeBody, fmt.Sprintf("\"title\":\"%v\"", *editAttributes.Title))
		}

		if editAttributes.Image != nil {
			attributeBody = append(attributeBody, fmt.Sprintf("\"image\":\"%v\"", *editAttributes.Image))
		}

		if params.EditLocation != nil {

			editLocation := params.EditLocation
			var locationBody []string

			if editLocation.Address1 != nil {
				locationBody = append(locationBody, fmt.Sprintf("\"address1\":\"%v\"", *editLocation.Address1))
			}

			if editLocation.Address2 != nil {
				locationBody = append(locationBody, fmt.Sprintf("\"address2\":\"%v\"", *editLocation.Address2))
			}

			if editLocation.City != nil {
				locationBody = append(locationBody, fmt.Sprintf("\"city\":\"%v\"", *editLocation.City))
			}

			if editLocation.Country != nil {
				locationBody = append(locationBody, fmt.Sprintf("\"country\":\"%v\"", *editLocation.Country))
			}

			if editLocation.Region != nil {
				locationBody = append(locationBody, fmt.Sprintf("\"region\":\"%v\"", *editLocation.Region))
			}

			if editLocation.Zip != nil {
				locationBody = append(locationBody, fmt.Sprintf("\"zip\":\"%v\"", *editLocation.Zip))
			}

			attributeBody = append(attributeBody, fmt.Sprintf("\"location\":{%v}", strings.Join(locationBody, ",")))
		}

		payloadBuild = append(payloadBuild, fmt.Sprintf("\"attributes\":{%v}", strings.Join(attributeBody, ",")))
	}
	//
	payloadString := fmt.Sprintf("{\"data\":{%v}}", strings.Join(payloadBuild, ","))
	payload := strings.NewReader(payloadString)

	errRequest := s.client.Request("PATCH", url, *payload, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return nil, nil
}

func (s *ProfileServiceOp) Create(ctx context.Context, params CreateProfile) (*ProfilesResponse, error) {

	var resp ProfilesResponse

	payloadBuild := []string{
		"\"type\":\"profile\"",
	}

	var attributeBody []string

	if params.Email != nil {
		attributeBody = append(attributeBody, fmt.Sprintf("\"email\":\"%v\"", *params.Email))
	}

	if params.PhoneNumber != nil {
		attributeBody = append(attributeBody, fmt.Sprintf("\"phone_number\":\"%v\"", *params.PhoneNumber))
	}

	if params.FirstName != nil {
		attributeBody = append(attributeBody, fmt.Sprintf("\"first_name\":\"%v\"", *params.FirstName))
	}

	if params.LastName != nil {
		attributeBody = append(attributeBody, fmt.Sprintf("\"last_name\":\"%v\"", *params.LastName))
	}

	if params.Organization != nil {
		attributeBody = append(attributeBody, fmt.Sprintf("\"organization\":\"%v\"", *params.Organization))
	}

	if params.Title != nil {
		attributeBody = append(attributeBody, fmt.Sprintf("\"title\":\"%v\"", *params.Title))
	}

	if params.Image != nil {
		attributeBody = append(attributeBody, fmt.Sprintf("\"image\":\"%v\"", *params.Image))
	}

	if params.Location != nil {

		editLocation := params.Location
		var locationBody []string

		if editLocation.Address1 != nil {
			locationBody = append(locationBody, fmt.Sprintf("\"address1\":\"%v\"", *editLocation.Address1))
		}

		if editLocation.Address2 != nil {
			locationBody = append(locationBody, fmt.Sprintf("\"address2\":\"%v\"", *editLocation.Address2))
		}

		if editLocation.City != nil {
			locationBody = append(locationBody, fmt.Sprintf("\"city\":\"%v\"", *editLocation.City))
		}

		if editLocation.Country != nil {
			locationBody = append(locationBody, fmt.Sprintf("\"country\":\"%v\"", *editLocation.Country))
		}

		if editLocation.Region != nil {
			locationBody = append(locationBody, fmt.Sprintf("\"region\":\"%v\"", *editLocation.Region))
		}

		if editLocation.Zip != nil {
			locationBody = append(locationBody, fmt.Sprintf("\"zip\":\"%v\"", *editLocation.Zip))
		}

		attributeBody = append(attributeBody, fmt.Sprintf("\"location\":{%v}", strings.Join(locationBody, ",")))
	}

	payloadBuild = append(payloadBuild, fmt.Sprintf("\"attributes\":{%v}", strings.Join(attributeBody, ",")))

	payloadString := fmt.Sprintf("{\"data\":{%v}}", strings.Join(payloadBuild, ","))
	payload := strings.NewReader(payloadString)

	errRequest := s.client.Request("POST", profileURL, *payload, nil)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}
