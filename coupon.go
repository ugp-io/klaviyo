package klaviyo

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

type CouponServiceOp struct {
	client *Client
}

type CouponCodeServiceOp struct {
	client *Client
}

type CouponService interface {
	Read(context.Context, ReadCouponRequest) (*CouponResponse, error)
	Browse(context.Context, BrowseCouponRequest) (*CouponResponses, error)
	Edit(context.Context, EditCouponRequest) (*CouponResponse, error)
	Create(context.Context, CreateCouponRequest) (*CouponResponse, error)
	ReadCode(context.Context, ReadCouponCodeRequest) (*CouponCodeResponse, error)
	BrowseCodes(context.Context, BrowseCouponCodeRequest) (*CouponCodeResponses, error)
	EditCode(context.Context, EditCouponCodeRequest) (*CouponCodeResponse, error)
	CreateCode(context.Context, CreateCouponCodeRequest) (*CouponCodeResponse, error)
}

type BrowseCouponRequest struct {
}

type ReadCouponRequest struct {
	ID string
}

type EditCouponRequest struct {
	ID          string
	Description *string
}

type CreateCouponRequest struct {
	ExternalID  string
	Description *string
}

type CouponResponse struct {
	Data CouponResponseData `json:"data,omitempty"`
}

type CouponResponses struct {
	Data  []CouponResponseData `json:"data,omitempty"`
	Links map[string]string    `json:"links,omitempty"`
}

type CouponResponseData struct {
	ID         *string                   `json:"id,omitempty"`
	Type       *string                   `json:"type,omitempty"`
	Attributes *CouponResponseAttributes `json:"attributes,omitempty"`
	Links      *map[string]string        `json:"links,omitempty"`
}

type CouponResponseAttributes struct {
	ExternalID  *string `json:"external_id,omitempty"`
	Description *string `json:"description,omitempty"`
}

type BrowseCouponCodeRequest struct {
	CouponID  *[]string
	ProfileID *[]string
}

type ReadCouponCodeRequest struct {
	ID string
}

type EditCouponCodeRequest struct {
	ID        string
	Status    *string
	ExpiresAt *string
}

type CreateCouponCodeRequest struct {
	Code     string
	CouponID string
	ExiredAt *string
}

type CouponCodeResponse struct {
	Data CouponCodeResponseData
}

type CouponCodeResponses struct {
	Data     []CouponCodeResponseData `json:"data,omitempty"`
	Links    map[string]string        `json:"links,omitempty"`
	Included []CouponResponseData     `json:"included,omitempty"`
}

type CouponCodeResponseData struct {
	ID            *string                       `json:"id,omitempty"`
	Type          *string                       `json:"type,omitempty"`
	Attributes    *CouponCodeResponseAttributes `json:"attributes,omitempty"`
	Links         *map[string]string            `json:"links,omitempty"`
	Relationships *interface{}                  `json:"relationships,omitempty"`
}

type CouponCodeResponseAttributes struct {
	UniqueCode *string `json:"unique_code,omitempty"`
	ExpiresAt  *string `json:"expires_at,omitempty"`
	Status     *string `json:"status,omitempty"`
}

const couponURL = "https://a.klaviyo.com/api/coupons/"
const couponCodeURL = "https://a.klaviyo.com/api/coupon-codes/"

func (s *CouponServiceOp) Read(ctx context.Context, params ReadCouponRequest) (*CouponResponse, error) {

	var resp CouponResponse

	url := fmt.Sprintf("%v%v", couponURL, params.ID)

	errRequest := s.client.Request("GET", url, strings.Reader{}, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}

func (s *CouponServiceOp) Browse(ctx context.Context, params BrowseCouponRequest) (*CouponResponses, error) {

	var resp CouponResponses

	errRequest := s.client.Request("GET", couponURL, strings.Reader{}, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}

func (s *CouponServiceOp) Edit(ctx context.Context, params EditCouponRequest) (*CouponResponse, error) {

	var resp CouponResponse
	url := fmt.Sprintf("%v%v", couponURL, params.ID)
	payloadBuild := []string{
		"\"type\":\"coupon\"",
		"\"id\":\"" + params.ID + "\"",
	}

	if params.Description != nil {
		payloadBuild = append(payloadBuild,
			fmt.Sprintf("\"attributes\":{\"description\":\"%v\"}", *params.Description))
	}

	payloadString := fmt.Sprintf("{\"data\":{%v}}", strings.Join(payloadBuild, ","))
	payload := strings.NewReader(payloadString)

	errRequest := s.client.Request("PATCH", url, *payload, nil)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}

func (s *CouponServiceOp) Create(ctx context.Context, params CreateCouponRequest) (*CouponResponse, error) {

	var resp CouponResponse
	payloadBuild := []string{
		"\"type\":\"coupon\"",
	}
	attributeBuild := []string{
		"\"external_id\":\"" + params.ExternalID + "\"",
	}

	if params.Description != nil {
		attributeBuild = append(attributeBuild,
			fmt.Sprintf("\"description\":\"%v\"", *params.Description))
	}

	payloadBuild = append(payloadBuild,
		fmt.Sprintf("\"attributes\":{%v}", strings.Join(attributeBuild, ",")))

	payloadString := fmt.Sprintf("{\"data\":{%v}}", strings.Join(payloadBuild, ","))
	payload := strings.NewReader(payloadString)

	errRequest := s.client.Request("POST", couponURL, *payload, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}

func (s *CouponServiceOp) ReadCode(ctx context.Context, params ReadCouponCodeRequest) (*CouponCodeResponse, error) {

	var resp CouponCodeResponse

	url := fmt.Sprintf("%v%v", couponCodeURL, params.ID)

	errRequest := s.client.Request("GET", url, strings.Reader{}, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}

func (s *CouponServiceOp) BrowseCodes(ctx context.Context, params BrowseCouponCodeRequest) (*CouponCodeResponses, error) {

	var urlString string
	var resp CouponCodeResponses
	uri := url.Values{}

	if params.CouponID != nil {

		uri.Add("filter", "any(coupon.id,[\""+strings.Join(*params.CouponID, "\",\"")+"\"])")
		urlString = fmt.Sprintf("%v?%v", couponCodeURL, uri.Encode())
	}

	if params.ProfileID != nil {
		uri.Add("filter", "any(profile.id,[\""+strings.Join(*params.ProfileID, "\",\"")+"\"])")
		urlString = fmt.Sprintf("%v?%v", couponCodeURL, uri.Encode())
	}

	errRequest := s.client.Request("GET", urlString, strings.Reader{}, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}

func (s *CouponServiceOp) EditCode(ctx context.Context, params EditCouponCodeRequest) (*CouponCodeResponse, error) {

	var resp CouponCodeResponse
	url := fmt.Sprintf("%v%v", couponCodeURL, params.ID)
	payloadBuild := []string{
		"\"type\":\"coupon-code\"",
		"\"id\":\"" + params.ID + "\"",
	}

	var attributeBuild []string
	if params.ExpiresAt != nil {
		attributeBuild = append(attributeBuild,
			fmt.Sprintf("\"expires_at\":\"%v\"", *params.ExpiresAt))
	}

	if params.Status != nil {
		attributeBuild = append(attributeBuild,
			fmt.Sprintf("\"status\":\"%v\"", *params.Status))
	}

	payloadBuild = append(payloadBuild,
		fmt.Sprintf("\"attributes\":{%v}", strings.Join(attributeBuild, ",")))

	payloadString := fmt.Sprintf("{\"data\":{%v}}", strings.Join(payloadBuild, ","))
	payload := strings.NewReader(payloadString)

	errRequest := s.client.Request("PATCH", url, *payload, nil)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}

func (s *CouponServiceOp) CreateCode(ctx context.Context, params CreateCouponCodeRequest) (*CouponCodeResponse, error) {

	var resp CouponCodeResponse
	payloadBuild := []string{
		"\"type\":\"coupon-code\"",
		"\"relationships\":{\"coupon\":{\"data\":{\"type\":\"coupon\",\"id\":\"" + params.CouponID + "\"}}}",
	}

	attributesBuild := []string{
		"\"unique_code\":\"" + params.Code + "\"",
	}
	if params.ExiredAt != nil {
		attributesBuild = append(attributesBuild,
			fmt.Sprintf("\"expires_at\":\"%v\"", *params.ExiredAt))
	}
	payloadBuild = append(payloadBuild,
		fmt.Sprintf("\"attributes\":{%v}", strings.Join(attributesBuild, ",")))

	payloadString := fmt.Sprintf("{\"data\":{%v}}", strings.Join(payloadBuild, ","))
	payload := strings.NewReader(payloadString)

	errRequest := s.client.Request("POST", couponCodeURL, *payload, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}
