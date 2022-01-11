package justwatch

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

type Contents struct {
	Page         int     `json:"page"`
	PageSize     int     `json:"page_size"`
	TotalPages   int     `json:"total_pages"`
	TotalResults int     `json:"total_results"`
	Items        []Items `json:"items"`
}

type Items struct {
	JwEntityID     string    `json:"jw_entity_id"`
	ID             int       `json:"id"`
	Title          string    `json:"title"`
	FullPath       string    `json:"full_path"`
	TmdbPopularity float64   `json:"tmdb_popularity,omitempty"`
	ObjectType     string    `json:"object_type"`
	OriginalTitle  string    `json:"original_title,omitempty"`
	Offers         []Offers  `json:"offers"`
	Scoring        []Scoring `json:"scoring,omitempty"`
	ShowID         int       `json:"show_id,omitempty"`
	ShowTitle      string    `json:"show_title,omitempty"`
	SeasonNumber   int       `json:"season_number,omitempty"`
}

type Offers struct {
	JwEntityID       string `json:"jw_entity_id"`
	MonetizationType string `json:"monetization_type"`
	ProviderID       int    `json:"provider_id"`
	PackageShortName string `json:"package_short_name"`
	Currency         string `json:"currency"`
	Urls             Urls   `json:"urls"`
	PresentationType string `json:"presentation_type"`
	Country          string `json:"country"`
}

type Urls struct {
	StandardWeb    string `json:"standard_web"`
	DeeplinkWeb    string `json:"deeplink_web"`
	DeeplinkRokuos string `json:"deeplink_rokuos"`
}

type Scoring struct {
	ProviderType string  `json:"provider_type"`
	Value        float64 `json:"value"`
}

type ContentBodyOptions struct {
	AgeCertifications          []string          `json:"age_certifications"`
	ContentTypes               []string          `json:"content_types"`
	ExcludeProviders           []string          `json:"exclude_providers"`
	Genres                     []string          `json:"genres"`
	ExcludeGenres              []string          `json:"exclude_genres"`
	Languages                  string            `json:"languages"`
	MinPrice                   int               `json:"min_price"`
	MinRuntime                 int               `json:"min_runtime"`
	MatchingOffersOnly         bool              `json:"matching_offers_only"`
	MaxPrice                   int               `json:"max_price"`
	MaxRuntime                 int               `json:"max_runtime"`
	MonetizationTypes          []string          `json:"monetization_types"`
	PresentationTypes          []string          `json:"presentation_types"`
	Providers                  []string          `json:"providers"`
	ReleaseYearFrom            int               `json:"release_year_from"`
	ReleaseYearUntil           int               `json:"release_year_until"`
	ScoringFilterTypes         map[string]string `json:"scoring_filter_types"`
	TimelineType               interface{}       `json:"timeline_type"`
	SortBy                     string            `json:"sort_by"`
	SortAsc                    bool              `json:"sort_asc"`
	EnableProviderFilter       bool              `json:"enable_provider_filter"`
	ProductionCountries        []string          `json:"production_countries"`
	ExcludeProductionCountries []string          `json:"exclude_production_countries"`
	IsUpcoming                 bool              `json:"is_upcoming"`
	Provider                   string            `json:"provider"`
	Date                       string            `json:"date"`
	Page                       int               `json:"page"`
	PageSize                   int               `json:"page_size"`
}

type ContentOptions struct {
	Body               string `url:"body"`
	FilterPriceChanges bool   `url:"filter_price_changes"`
}

func (c *Client) ListNewContens(ctx context.Context, opts *ContentBodyOptions) (*Contents, error) {
	u := fmt.Sprintf("content/titles/%s/new/single_provider", c.locale)
	return c.listContents(ctx, u, opts)
}

func (c *Client) listContents(ctx context.Context, endpoint string, opts *ContentBodyOptions) (*Contents, error) {
	contentOpts, err := buildQuery(opts)
	if err != nil {
		return nil, err
	}

	u, err := addOptions(endpoint, contentOpts)
	if err != nil {
		return nil, err
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	contents := new(Contents)
	resp, err := c.Do(req, contents)
	if err != nil {
		return nil, err
	}
	log.Printf("request url: %s\n", resp.Request.URL.String())

	return contents, nil
}

func buildQuery(opts *ContentBodyOptions) (*ContentOptions, error) {
	body, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}

	contentOpts := &ContentOptions{
		Body:               string(body),
		FilterPriceChanges: false,
	}
	return contentOpts, nil
}
