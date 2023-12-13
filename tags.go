package klaviyo

import (
	"context"
	"fmt"
	"strings"
)

type TagServiceOp struct {
	client *Client
}

type TagService interface {
	Read(context.Context, ReadTagRequest) (*TagResponse, error)
	Browse(context.Context, BrowseTagRequest) (*TagResponses, error)
	Edit(context.Context, EditTagRequest) (*TagResponse, error)
	Create(context.Context, CreateTagRequest) (*TagResponse, error)
}

type BrowseTagRequest struct {
	Name *string
}

type ReadTagRequest struct {
	ID string
}

type EditTagRequest struct {
	ID            string
	AttributeName *string
}

type CreateTagRequest struct {
	Name       string
	TagGroupID *string
}

type TagResponse struct {
	Data     TagResponseData `json:"data,omitempty"`
	Included []TagIncluded   `json:"included,omitempty"`
}

type TagResponses struct {
	Links    map[string]string `json:"links,omitempty"`
	Data     []TagResponseData `json:"data,omitempty"`
	Included []TagIncluded     `json:"included,omitempty"`
}

type TagResponseData struct {
	ID            *string                `json:"id,omitempty"`
	Type          *string                `json:"type,omitempty"`
	Attributes    *TagResponseAttributes `json:"attributes,omitempty"`
	Links         *map[string]string     `json:"links,omitempty"`
	Relationships *interface{}           `json:"relationships,omitempty"`
}

type TagIncluded struct {
	ID        *string                `json:"id,omitempty"`
	Type      *string                `json:"type,omitempty"`
	Attibutes *TagResponseAttributes `json:"attributes,omitempty"`
	Links     map[string]string      `json:"links,omitempty"`
}

type TagResponseAttributes struct {
	Name      *string `json:"name,omitempty"`
	Exclusive *bool   `json:"exclusive,omitempty"`
	Default   *bool   `json:"default,omitempty"`
}

const tagURL = "https://a.klaviyo.com/api/tags/"

func (s *TagServiceOp) Read(ctx context.Context, params ReadTagRequest) (*TagResponse, error) {

	var resp TagResponse

	url := tagURL + params.ID

	errRequest := s.client.Request("GET", url, strings.Reader{}, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}

func (s *TagServiceOp) Browse(ctx context.Context, params BrowseTagRequest) (*TagResponses, error) {

	var resp TagResponses

	errRequest := s.client.Request("GET", tagURL, strings.Reader{}, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}

func (s *TagServiceOp) Edit(ctx context.Context, params EditTagRequest) (*TagResponse, error) {

	var resp TagResponse
	payloadBuild := []string{
		"\"type\":\"tag\"",
		"\"id\":\"" + params.ID + "\"",
	}
	url := tagURL + params.ID

	if params.AttributeName != nil {
		payloadBuild = append(payloadBuild,
			fmt.Sprintf("\"attributes\":{\"name\":\"%v\"}", *params.AttributeName))
	}

	payloadString := fmt.Sprintf("{\"data\":{%v}}", strings.Join(payloadBuild, ","))
	payload := strings.NewReader(payloadString)

	errRequest := s.client.Request("PATCH", url, *payload, nil)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}

func (s *TagServiceOp) Create(ctx context.Context, params CreateTagRequest) (*TagResponse, error) {

	var resp TagResponse
	payloadBuild := []string{
		"\"type\":\"tag\"",
		"\"attributes\":{\"name\":\"" + params.Name + "\"}",
	}

	if params.TagGroupID != nil {
		payloadBuild = append(payloadBuild,
			fmt.Sprintf("\"relationships\":{\"tag-group\":{\"data\":{\"type\":\"tag-group\",\"id\":\"%v\"}}}", *params.TagGroupID))
	}

	payloadString := fmt.Sprintf("{\"data\":{%v}}", strings.Join(payloadBuild, ","))
	payload := strings.NewReader(payloadString)

	errRequest := s.client.Request("POST", tagURL, *payload, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}
