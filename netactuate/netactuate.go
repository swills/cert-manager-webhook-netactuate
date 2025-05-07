package netactuate

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func GetDomainFromZone(fqdn string) string {
	return strings.TrimSuffix(fqdn, ".")
}

// GetZoneID returns the zone ID for a domain name, if it exists
func GetZoneID(domainName string, apiKey string) (int, error) {
	zoneList, err := DNSZoneGet(apiKey)

	if err != nil {
		return 0, fmt.Errorf("error getting zone: %w", err)
	}

	for _, zone := range zoneList.Data {
		if strings.EqualFold(zone.Name, strings.TrimRight(domainName, ".")) {
			return zone.ID, nil
		}
	}

	return 0, ErrDomainNotFound
}

// see https://docs.netactuate.com/reference/dns

// DNSZoneGet returns a list of all DNS Zones for an account
func DNSZoneGet(apiKey string) (*ZoneList, error) {
	var err error

	backgroundContext := context.Background()

	url := "https://vapi2.netactuate.com/api/dns/zones?type=NATIVE&key=" + apiKey

	var req *http.Request

	req, err = http.NewRequestWithContext(backgroundContext, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating http request: %w", err)
	}

	req.Header.Add("accept", "application/json")

	var res *http.Response

	res, err = http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making http request: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error response from netactuate api: %s, %w", res.Status, ErrHTTPNotOK)
	}

	defer func() {
		_ = res.Body.Close()
	}()

	var body []byte

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var zoneList ZoneList

	err = json.Unmarshal(body, &zoneList)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response body: %w", err)
	}

	return &zoneList, nil
}

// DNSRecordPost Adds a new DNS record to a Zone
func DNSRecordPost(apiKey string, domainName string, recordType string, recordName string, recordContent string) error {
	var err error

	var zoneID int

	zoneID, err = GetZoneID(domainName, apiKey)
	if err != nil {
		return fmt.Errorf("error getting zone ID: %w", err)
	}

	url := "https://vapi2.netactuate.com/api/dns/record?domain_id=" + strconv.FormatInt(int64(zoneID), 10) +
		"&name=" + strings.TrimRight(recordName, ".") + "&type=" + recordType + "&record_content=" + recordContent + "&key=" + apiKey

	backgroundContext := context.Background()

	var req *http.Request

	req, err = http.NewRequestWithContext(backgroundContext, http.MethodPost, url, nil)
	if err != nil {
		return fmt.Errorf("error creating http request: %w", err)
	}

	req.Header.Add("accept", "application/json")

	var res *http.Response

	res, err = http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error making http request: %w", err)
	}

	defer func() {
		_ = res.Body.Close()
	}()

	var body []byte

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	var dnsRecordPostResponse DNSRecordPostResponse

	err = json.Unmarshal(body, &dnsRecordPostResponse)
	if err != nil {
		return fmt.Errorf("error unmarshaling response body: %w", err)
	}

	if res.StatusCode == http.StatusOK && dnsRecordPostResponse.Code == http.StatusOK {
		return nil
	}

	return ErrUnknown
}

// DNSRecordsGet gets a list of DNS records for the given domain
func DNSRecordsGet(apiKey string, domainName string) ([]DNSRecord, error) {
	var err error

	var zoneID int

	zoneID, err = GetZoneID(domainName, apiKey)
	if err != nil {
		return nil, fmt.Errorf("error getting zone ID: %w", err)
	}

	url := "https://vapi2.netactuate.com/api/dns/records/" + strconv.FormatInt(int64(zoneID), 10) + "?key=" + apiKey

	backgroundContext := context.Background()

	var req *http.Request

	req, err = http.NewRequestWithContext(backgroundContext, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating http request: %w", err)
	}

	req.Header.Add("accept", "application/json")

	var res *http.Response

	res, err = http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making http request: %w", err)
	}

	defer func() {
		_ = res.Body.Close()
	}()

	var body []byte

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, ErrHTTPNotOK
	}

	var dnsRecordListResponse DNSRecordListResponse

	err = json.Unmarshal(body, &dnsRecordListResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response body: %w", err)
	}

	return dnsRecordListResponse.Data, nil
}

// DNSRecordDelete deletes a DNS record
func DNSRecordDelete(apiKey string, recordID int) error {
	var err error

	url := "https://vapi2.netactuate.com/api/dns/record/" + strconv.FormatInt(int64(recordID), 10) +
		"?key=" + apiKey

	backgroundContext := context.Background()

	var req *http.Request

	req, err = http.NewRequestWithContext(backgroundContext, http.MethodDelete, url, nil)
	if err != nil {
		return fmt.Errorf("error creating http request: %w", err)
	}

	req.Header.Add("accept", "application/json")

	var res *http.Response

	res, err = http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error making http request: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("error response from netactuate api: %s, %w", res.Status, ErrHTTPNotOK)
	}

	defer func() {
		_ = res.Body.Close()
	}()

	var body []byte

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	var zoneList ZoneList

	err = json.Unmarshal(body, &zoneList)
	if err != nil {
		return fmt.Errorf("error unmarshaling response body: %w", err)
	}

	return nil
}
