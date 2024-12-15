package namesilo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/er1c-zh/cert-manager-webhook-namesilo/utils"
)

type Response struct {
	Reply struct {
		Code   CodeWrapper `json:"code"`
		Detail string      `json:"detail"`
	} `json:"reply"`
}

// https://www.namesilo.com/api-reference#dns/dns-list-records
type DnsRecordListResponse struct {
	Reply struct {
		Code           CodeWrapper `json:"code"`
		Detail         string      `json:"detail"`
		ResourceRecord []struct {
			ResourceID string `json:"record_id"`
			Type       string `json:"type"`
			Host       string `json:"host"`
			Value      string `json:"value"`
		} `json:"resource_record"`
	} `json:"reply"`
}

// FIXME: namesilo api's response code sometimes is string instead of int
type CodeWrapper string

func (c *CodeWrapper) UnmarshalJSON(data []byte) error {
	*c = CodeWrapper(string(data))
	*c = CodeWrapper(strings.Trim(string(data), "\""))
	return nil
}

func Call[Resp any](apiKey string, operation string, params map[string]string) (Resp, error) {
	var resp Resp
	req, err := http.NewRequest("GET", "https://www.namesilo.com/api/"+operation, nil)
	if err != nil {
		return resp, err
	}

	queryParams := req.URL.Query()
	// common params
	queryParams.Set("version", "1")
	queryParams.Set("type", "json")
	queryParams.Set("key", apiKey)
	// add record api params
	for k, v := range params {
		queryParams.Set(k, v)
	}
	req.URL.RawQuery = queryParams.Encode()
	httpResp, err := http.DefaultClient.Do(req)
	if err != nil {
		return resp, err
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode != 200 {
		return resp, fmt.Errorf("namesilo: %s", httpResp.Status)
	}
	b, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return resp, err
	}
	if err := json.Unmarshal(b, &resp); err != nil {
		utils.Log("namesilo unmarsha fail: %s, data: %s", err.Error(), string(b))
		return resp, fmt.Errorf("namesilo unmarshal json fail: %s", err.Error())
	}
	return resp, nil
}
