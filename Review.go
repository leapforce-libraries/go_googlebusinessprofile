package googlebusinessprofile

import (
	"fmt"
	types "github.com/leapforce-libraries/go_googlebusinessprofile/types"
	"net/http"
	"net/url"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type ReviewsResponse struct {
	Reviews          []Review `json:"reviews"`
	AverageRating    float64  `json:"averageRating"`
	TotalReviewCount float64  `json:"totalReviewCount"`
	NextPageToken    string   `json:"nextPageToken"`
}

type Review struct {
	Name        string               `json:"name"`
	ReviewId    string               `json:"reviewId"`
	Reviewer    Reviewer             `json:"reviewer"`
	StarRating  string               `json:"starRating"`
	Comment     string               `json:"comment"`
	CreateTime  types.DateTimeString `json:"createTime"`
	UpdateTime  types.DateTimeString `json:"updateTime"`
	ReviewReply ReviewReply          `json:"reviewReply"`
}

type Reviewer struct {
	ProfilePhotoUrl string `json:"profilePhotoUrl"`
	DisplayName     string `json:"displayName"`
	IsAnonymous     bool   `json:"isAnonymous"`
}

type ReviewReply struct {
	Comment    string               `json:"comment"`
	UpdateTime types.DateTimeString `json:"updateTime"`
}

type ReviewsConfig struct {
	AccountName  string
	LocationName string
	PageSize     *int
	OrderBy      *string
}

func (service *Service) Reviews(config *ReviewsConfig) (*[]Review, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("ReviewsConfig must not be nil")
	}
	values := url.Values{}

	if config.PageSize != nil {
		values.Set("pageSize", fmt.Sprintf("%v", *config.PageSize))
	}
	if config.OrderBy != nil {
		values.Set("orderBy", *config.OrderBy)
	}

	var reviews []Review

	for {
		url := fmt.Sprintf("https://mybusiness.googleapis.com/v4/%s/%s/reviews?%s", config.AccountName, config.LocationName, values.Encode())

		reviewsReponse := ReviewsResponse{}

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodGet,
			Url:           url,
			ResponseModel: &reviewsReponse,
		}

		_, _, e := service.googleService.HttpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		reviews = append(reviews, reviewsReponse.Reviews...)

		if reviewsReponse.NextPageToken == "" {
			break
		}

		values.Set("pageToken", reviewsReponse.NextPageToken)
	}

	return &reviews, nil
}
