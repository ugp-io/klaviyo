package klaviyo

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Client struct {
	APIKey    string
	APISecret string
	Profile   ProfileService
	Event     EventService
	Metric    MetricService
	Tag       TagService
	Coupon    CouponService
}

func NewClient(apiKey string) *Client {

	c := &Client{
		APIKey: apiKey,
	}

	c.Profile = &ProfileServiceOp{client: c}
	c.Event = &EventServiceOp{client: c}
	c.Metric = &MetricServiceOp{client: c}
	c.Tag = &TagServiceOp{client: c}
	c.Coupon = &CouponServiceOp{client: c}

	return c

}

func (c *Client) Request(method string, url string, body interface{}, v interface{}) error {

	var bodyReader *strings.Reader
	if body != nil {
		requestJson, errMarshal := json.Marshal(body)
		if errMarshal != nil {
			return errMarshal
		}
		bodyReader = strings.NewReader(string(requestJson))
	}

	httpReq, errNewRequest := http.NewRequest(method, url, bodyReader)
	if errNewRequest != nil {
		return errNewRequest
	}

	// Content Type
	httpReq.Header.Add("accept", "application/json")
	httpReq.Header.Add("revision", "2023-10-15")
	httpReq.Header.Add("content-type", "application/json")
	httpReq.Header.Add("Authorization", "Klaviyo-API-Key "+c.APIKey)

	//Client
	client := &http.Client{}
	resp, errDo := client.Do(httpReq)

	// fmt.Println(resp)
	// fmt.Println()
	// fmt.Println(errDo)

	if resp != nil {
		defer resp.Body.Close()
	}
	if errDo != nil {
		return errDo
	}

	if v != nil {
		decoder := json.NewDecoder(resp.Body)
		errDecode := decoder.Decode(&v)
		if errDecode != nil {
			return errDecode
		}
	}
	return nil
}
