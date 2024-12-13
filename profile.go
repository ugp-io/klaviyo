package klaviyo

import (
	"context"
	"strings"
	"time"
)

type ProfileServiceOp struct {
	client *Client
}

type ProfileService interface {
	Read(context.Context, ProfileRequest) (*ProfileResponse, error)
	Browse(context.Context, ProfileRequest) (*ProfilesResponse, error)
	Edit(context.Context, EditProfile) (*ProfileResponse, error)
	Create(context.Context, CreateProfile) (*ProfilesResponse, error)
}

type ProfileRequest struct {
	ID     string
	Emails *[]string
}

type EditProfile struct {
	Data EditProfileData `json:"data,omitempty"`
}

type EditProfileData struct {
	Type       string                 `json:"type,omitempty"`
	ID         string                 `json:"id,omitempty"`
	Attributes *EditProfileAttributes `json:"attributes,omitempty"`
	Meta       *EditProfileMeta       `json:"meta,omitempty"`
}

type EditProfileAttributes struct {
	Email        *string              `json:"email,omitempty"`
	PhoneNumber  *string              `json:"phone_number,omitempty"`
	ExternalID   *string              `json:"external_id,omitempty"`
	AnonymousID  *string              `json:"anonymous_id,omitempty"`
	FirstName    *string              `json:"first_name,omitempty"`
	LastName     *string              `json:"last_name,omitempty"`
	Organization *string              `json:"organization,omitempty"`
	Title        *string              `json:"title,omitempty"`
	Image        *string              `json:"image,omitempty"`
	Location     *EditProfileLocation `json:"location,omitempty"`
	Properties   *ProfileProperties   `json:"properties,omitempty"`
}

type ProfileProperties struct {
	LeadGroupNumber               string      `json:"lead.group_number,omitempty"`
	LeadDesignCount               int         `json:"lead.design_count,omitempty"`
	LeadQuoteCount                int         `json:"lead.quote_count,omitempty"`
	LeadStatus                    *string     `json:"lead.status,omitempty"`
	LeadProgress                  *string     `json:"lead.progress,omitempty"`
	LeadMinVal                    *int        `json:"lead.min_val,omitempty"`
	LeadMaxVal                    *int        `json:"lead.max_val,omitempty"`
	LeadAvgVal                    *int        `json:"lead.avg_val,omitempty"`
	LeadMinCreated                *time.Time  `json:"lead.min_created,omitempty"`
	LeadMaxCreated                *time.Time  `json:"lead.max_created,omitempty"`
	LeadMinArrival                *time.Time  `json:"lead.min_arrival,omitempty"`
	LeadMaxArrival                *time.Time  `json:"lead.max_arrival,omitempty"`
	LeadAdminInitiated            *bool       `json:"lead.admin_initiated,omitempty"`
	LeadAdminInteracted           *bool       `json:"lead.admin_interacted,omitempty"`
	LeadCustomPricing             *bool       `json:"lead.custom_pricing,omitempty"`
	LeadConverted                 *bool       `json:"lead.converted,omitempty"`
	LeadPublicLink                *string     `json:"lead.public_link,omitempty"`
	LeadEmployeeFirstName         *string     `json:"lead.employee.first_name,omitempty"`
	LeadEmployeeLastName          *string     `json:"lead.employee.last_name,omitempty"`
	LeadEmployeeFullName          *string     `json:"lead.employee.full_name,omitempty"`
	LeadEmployeeEmail             *string     `json:"lead.employee.email,omitempty"`
	LeadStoreCode                 *string     `json:"lead.store.code,omitempty"`
	LeadStoreAddress              *string     `json:"lead.store.address,omitempty"`
	LeadStoreEmail                *string     `json:"lead.store.email,omitempty"`
	LeadStorePhone                *string     `json:"lead.store.phone,omitempty"`
	LeadStoreRMEmail              *string     `json:"lead.store.RMEmail,omitempty"`
	LeadShippingLocationFirstName *string     `json:"lead.shipping_location.first_name,omitempty"`
	LeadShippingLocationLastName  *string     `json:"lead.shipping_location.last_name,omitempty"`
	LeadShippingLocationFullName  *string     `json:"lead.shipping_location.full_name,omitempty"`
	LeadShippingLocationAddress1  *string     `json:"lead.shipping_location.address1,omitempty"`
	LeadShippingLocationAddress2  *string     `json:"lead.shipping_location.address2,omitempty"`
	LeadShippingLocationCity      *string     `json:"lead.shipping_location.city,omitempty"`
	LeadShippingLocationState     *string     `json:"lead.shipping_location.state,omitempty"`
	LeadShippingLocationZip       *string     `json:"lead.shipping_location.zip,omitempty"`
	LeadShippingLocationCountry   *string     `json:"lead.shipping_location.country,omitempty"`
	LeadDesigns                   interface{} `json:"lead.designs,omitempty"`
	StatsCustomerLifetimeVal      *int        `json:"stats.customer_lifetime_val,omitempty"`
	StatsFirstOrderID             *int        `json:"stats.first_order.id,omitempty"`
	StatsFirstOrderDate           *time.Time  `json:"stats.first_order.date,omitempty"`
	StatsFirstOrderValue          *int        `json:"stats.first_order.value,omitempty"`
	StatsMostRecentOrderID        *int        `json:"stats.most_recent_order.id,omitempty"`
	StatsMostRecentOrderDate      *time.Time  `json:"stats.most_recent_order.date,omitempty"`
	StatsMostRecentOrderValue     *int        `json:"stats.most_recent_order.value,omitempty"`
}

type EditProfileLocation struct {
	Address1  *string `json:"address1,omitempty"`
	Address2  *string `json:"address2,omitempty"`
	City      *string `json:"city,omitempty"`
	Country   *string `json:"country,omitempty"`
	Latitude  *string `json:"latitude,omitempty"`
	Longitude *string `json:"longitude,omitempty"`
	Region    *string `json:"region,omitempty"`
	Zip       *string `json:"zip,omitempty"`
	Timezone  *string `json:"timezone,omitempty"`
	IP        *string `json:"ip,omitempty"`
}

type EditProfileMeta struct {
	PatchProperties *EditProfileMetaPatchProperties `json:"patch_properties,omitempty"`
}

type EditProfileMetaPatchProperties struct {
	Unset    *string     `json:"unset,omitempty"`
	Append   interface{} `json:"append,omitempty"`
	Unappend interface{} `json:"unappend,omitempty"`
}

type CreateProfile struct {
	Data *CreateProfileData `json:"data,omitempty"`
}

type CreateProfileData struct {
	Type       string                   `json:"type,omitempty"`
	ID         *string                  `json:"id,omitempty"`
	Attributes *CreateProfileAttributes `json:"attributes,omitempty"`
}

type CreateProfileAttributes struct {
	Email        *string                `json:"email,omitempty"`
	PhoneNumber  *string                `json:"phone_number,omitempty"`
	FirstName    *string                `json:"first_name,omitempty"`
	LastName     *string                `json:"last_name,omitempty"`
	Organization *string                `json:"organization,omitempty"`
	Title        *string                `json:"title,omitempty"`
	Image        *string                `json:"image,omitempty"`
	Properties   *map[string]string     `json:"properties,omitempty"`
	Location     *CreateProfileLocation `json:"location,omitempty"`
}

type CreateProfileLocation struct {
	Address1  *string `json:"address1,omitempty"`
	Address2  *string `json:"address2,omitempty"`
	City      *string `json:"city,omitempty"`
	Country   *string `json:"country,omitempty"`
	Latitude  *string `json:"latitude,omitempty"`
	Longitude *string `json:"longitude,omitempty"`
	Region    *string `json:"region,omitempty"`
	Zip       *string `json:"zip,omitempty"`
	Timezone  *string `json:"timezone,omitempty"`
	IP        *string `json:"ip,omitempty"`
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
	Properties          interface{}                         `json:"properties,omitempty"`
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

func (s *ProfileServiceOp) Edit(ctx context.Context, params EditProfile) (*ProfileResponse, error) {

	var resp ProfileResponse
	url := profileURL + params.Data.ID

	errRequest := s.client.Request("PATCH", url, params, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}

func (s *ProfileServiceOp) Create(ctx context.Context, params CreateProfile) (*ProfilesResponse, error) {

	var resp ProfilesResponse

	errRequest := s.client.Request("POST", profileURL, params, nil)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}
