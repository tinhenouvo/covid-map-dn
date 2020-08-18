package utils

import (
	"encoding/base64"
	"encoding/json"
	"time"

	"crawlerdatacovid/internal/pkg/config"
)

// PageToken ...
type PageTokenRequest struct {
	PageToken string
	Page      int64
	Offset    int
	Date      time.Time
}

type PaginationToken struct {
	CurrentPageToken string `json:"currentPageToken"`
	NextPageToken    string `json:"nextPageToken"`
}

type PageTokenList struct {
	PaginationToken
	Items interface{} `json:"items"`
}

func getDefaultPageToken() (response PageTokenRequest) {
	var currentPage int64 = 0
	limit := config.Limit20
	response.Page = currentPage
	response.Offset = limit * int(currentPage)
	response.Date = time.Now()
	response.PageToken = PageTokenEncode(currentPage, time.Now())
	return response
}

// PageTokenEncode encode next page token for api response
func PageTokenEncode(page int64, date time.Time) string {
	tokenData := PageTokenRequest{
		Page: page,
		Date: date,
	}
	tokenString, _ := json.Marshal(tokenData)
	encodedString := base64.StdEncoding.EncodeToString(tokenString)
	return encodedString
}

// PageTokenUsingPage generate next page token using "page"
func PageTokenUsingPage(totalRecords int64, limit int64, nextPage int64) string {
	if totalRecords == limit {
		return PageTokenEncode(nextPage, time.Now())
	}
	return ""
}

// PageTokenUsingDate generate next page token using "date"
func PageTokenUsingDate(totalRecords int64, limit int64, date time.Time) string {
	if totalRecords == limit {
		return PageTokenEncode(0, date)
	}
	return ""
}

func NextPageToken(totalsItems int, currentPageToken string) string {
	decoded, err := base64.StdEncoding.DecodeString(currentPageToken)
	if err != nil {
		return ""
	}
	// Parse token
	var pageToken PageTokenRequest
	err = json.Unmarshal(decoded, &pageToken)
	if err != nil {
		print(err)
	}
	return PageTokenUsingPage(int64(totalsItems), config.Limit20, pageToken.Page+1)
}

func PaginationTokenResponse(currentPageToken string, totalItems int, list interface{}) PageTokenList {
	paginationTokenList := PageTokenList{}
	nextPage := NextPageToken(totalItems, currentPageToken)
	paginationTokenList.CurrentPageToken = currentPageToken
	paginationTokenList.NextPageToken = nextPage
	paginationTokenList.Items = list
	return paginationTokenList
}

func NewPaginationTokenRequest(currentToken string) PageTokenRequest {
	if currentToken != "" {
		return getDefaultPageToken()
	}
	// Decode string
	decoded, err := base64.StdEncoding.DecodeString(currentToken)
	if err != nil {
		return getDefaultPageToken()
	}
	// Parse token
	var pageToken PageTokenRequest
	pageToken.PageToken = string(decoded)
	err = json.Unmarshal(decoded, &pageToken)
	if err != nil {
		return getDefaultPageToken()
	}
	return pageToken
}
