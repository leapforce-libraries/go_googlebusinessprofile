package googlebusinessprofile

import (
	"fmt"
	"net/http"
	"net/url"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type LocationsResponse struct {
	Locations     []Location `json:"locations"`
	NextPageToken string     `json:"nextPageToken"`
}

type Location struct {
	Name         string `json:"name"`
	LanguageCode string `json:"languageCode"`
	StoreCode    string `json:"storeCode"`
	Title        string `json:"title"`
	//PhoneNumbers              string   `json:"phoneNumbers"`
	//Categories                string   `json:"categories"`
	//StorefrontAddress         string   `json:"storefrontAddress"`
	WebsiteUri string `json:"websiteUri"`
	//RegularHours              string   `json:"regularHours"`
	//SpecialHours              string   `json:"specialHours"`
	//ServiceArea               string   `json:"serviceArea"`
	//Labels                    []string `json:"labels"`
	//AdWordsLocationExtensions string   `json:"adWordsLocationExtensions"`
	//Latlng                    string   `json:"latlng"`
	//OpenInfo                  string   `json:"openInfo"`
	//Metadata                  string   `json:"metadata"`
	//Profile                   string   `json:"profile"`
	//RelationshipData          string   `json:"relationshipData"`
	//MoreHours                 string   `json:"moreHours"`
	//ServiceItems              string   `json:"serviceItems"`
}

type LocationsConfig struct {
	Account  string
	PageSize *int
	Filter   *string
	OrderBy  *string
	ReadMask *string
}

func (service *Service) Locations(config *LocationsConfig) (*[]Location, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("LocationsConfig must not be nil")
	}
	var values url.Values

	if config.PageSize != nil {
		values.Set("pageSize", fmt.Sprintf("%v", *config.PageSize))
	}
	if config.Filter != nil {
		values.Set("filter", *config.Filter)
	}
	if config.OrderBy != nil {
		values.Set("orderBy", *config.OrderBy)
	}
	if config.ReadMask != nil {
		values.Set("readMask", *config.ReadMask)
	}

	var locations []Location

	for {
		url := fmt.Sprintf("https://mybusinessbusinessinformation.googleapis.com/v1/accounts/%s/locations?%s", config.Account, values.Encode())

		locationsReponse := LocationsResponse{}

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           url,
			ResponseModel: &locationsReponse,
		}

		_, _, e := service.googleService.HttpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		locations = append(locations, locationsReponse.Locations...)

		if locationsReponse.NextPageToken == "" {
			break
		}

		values.Set("pageToken", locationsReponse.NextPageToken)
	}

	return &locations, nil
}
