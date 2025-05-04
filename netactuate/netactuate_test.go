package netactuate

import (
	"os"
	"strconv"
	"testing"
	"time"
)

var netactuateAPIKey = "your-api-key"
var testDomain = "example.com"

func TestGetZoneID(t *testing.T) {
	t.Parallel()

	netactuateAPIKey = os.Getenv("NETACTUATE_API_KEY")
	if netactuateAPIKey == "" || netactuateAPIKey == "your-api-key" {
		t.Fatal("error getting api key")
	}

	testDomain = os.Getenv("TEST_DOMAIN")
	if testDomain == "" || testDomain == "example.com" {
		t.Fatal("error getting domain")
	}

	type args struct {
		domainName string
		apiKey     string
	}

	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    testDomain,
			args:    args{domainName: testDomain, apiKey: netactuateAPIKey},
			want:    296650,
			wantErr: false,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			got, err := GetZoneID(testCase.args.domainName, testCase.args.apiKey)
			if (err != nil) != testCase.wantErr {
				t.Errorf("GetZoneID() error = %v, wantErr %v", err, testCase.wantErr)

				return
			}

			if got != testCase.want {
				t.Errorf("GetZoneID() = %v, want %v", got, testCase.want)
			}
		})
	}
}

func TestDNSRecordPost(t *testing.T) {
	t.Parallel()

	netactuateAPIKey = os.Getenv("NETACTUATE_API_KEY")
	if netactuateAPIKey == "" || netactuateAPIKey == "your-api-key" {
		t.Fatal("error getting api key")
	}

	testDomain = os.Getenv("TEST_DOMAIN")
	if testDomain == "" || testDomain == "example.com" {
		t.Fatal("error getting domain")
	}

	type args struct {
		apiKey        string
		domainName    string
		recordType    string
		recordName    string
		recordContent string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				apiKey:        netactuateAPIKey,
				domainName:    testDomain,
				recordType:    "A",
				recordName:    "test" + strconv.FormatInt(time.Now().Unix(), 10),
				recordContent: "1.1.1.1",
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			err := DNSRecordPost(
				testCase.args.apiKey,
				testCase.args.domainName,
				testCase.args.recordType,
				testCase.args.recordName,
				testCase.args.recordContent,
			)

			if (err != nil) != testCase.wantErr {
				t.Errorf("DNSRecordPost() error = %v, wantErr %v", err, testCase.wantErr)
			}
		})
	}
}

func TestDNSRecordsGet(t *testing.T) {
	t.Parallel()

	netactuateAPIKey = os.Getenv("NETACTUATE_API_KEY")
	if netactuateAPIKey == "" || netactuateAPIKey == "your-api-key" {
		t.Fatal("error getting api key")
	}

	testDomain = os.Getenv("TEST_DOMAIN")
	if testDomain == "" || testDomain == "example.com" {
		t.Fatal("error getting domain")
	}

	type args struct {
		apiKey     string
		domainName string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				apiKey:     netactuateAPIKey,
				domainName: testDomain,
			},
			wantErr: false,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			_, err := DNSRecordsGet(testCase.args.apiKey, testCase.args.domainName)
			if (err != nil) != testCase.wantErr {
				t.Errorf("DNSRecordsGet() error = %v, wantErr %v", err, testCase.wantErr)

				return
			}
		})
	}
}

func TestDNSRecordDelete(t *testing.T) {
	t.Parallel()

	netactuateAPIKey = os.Getenv("NETACTUATE_API_KEY")
	if netactuateAPIKey == "" || netactuateAPIKey == "your-api-key" {
		t.Fatal("error getting api key")
	}

	testDomain = os.Getenv("TEST_DOMAIN")
	if testDomain == "" || testDomain == "example.com." ||
		testDomain == "example.com" || testDomain == "example.coM." {
		t.Fatal("error getting domain")
	}

	testRecordID := os.Getenv("TEST_RECORD_ID")
	if testRecordID == "" || testRecordID == "123456" {
		t.Fatal("error getting test record ID")
	}

	testRecordIDInt64, err := strconv.ParseInt(testRecordID, 10, 64)
	if err != nil {
		t.Fatal("error getting test record ID")
	}

	type args struct {
		apiKey   string
		recordID int
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				apiKey:   netactuateAPIKey,
				recordID: int(testRecordIDInt64),
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			err := DNSRecordDelete(testCase.args.apiKey, testCase.args.recordID)
			if (err != nil) != testCase.wantErr {
				t.Errorf("DNSRecordDelete() error = %v, wantErr %v", err, testCase.wantErr)
			}
		})
	}
}
