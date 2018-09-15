package chicmi

import "net/http"

// EventDetail is event detail with meta
type EventDetail struct {
	Values Values `json:"values"`
	Meta   Meta   `json:"meta"`
}

// Values is event detail information
type Values struct {
	ID                     string        `json:"id"`
	Created                string        `json:"created"`
	Updated                string        `json:"updated"`
	ReviewPrompted         interface{}   `json:"review_prompted"`
	EventDbID              string        `json:"event_db_id"`
	SectorTypeID           string        `json:"sector_type_id"`
	EventTypeID            string        `json:"event_type_id"`
	SecondaryEventTypeID   interface{}   `json:"secondary_event_type_id"`
	EventEligibilityTypeID string        `json:"event_eligibility_type_id"`
	RegistrationURL        string        `json:"registration_url"`
	EventNameEn            string        `json:"event_name_en"`
	EventNameZh            string        `json:"event_name_zh"`
	EventNameRu            interface{}   `json:"event_name_ru"`
	EventNameAr            interface{}   `json:"event_name_ar"`
	Slug                   string        `json:"slug"`
	ShortSlug              string        `json:"short_slug"`
	VanityURL              interface{}   `json:"vanity_url"`
	Icons                  string        `json:"icons"`
	StartDate              string        `json:"start_date"`
	StartMonth             string        `json:"start_month"`
	StartYear              string        `json:"start_year"`
	EndDate                string        `json:"end_date"`
	EndMonth               string        `json:"end_month"`
	EndYear                string        `json:"end_year"`
	VenueID                string        `json:"venue_id"`
	VenueNotified          interface{}   `json:"venue_notified"`
	SlackNotified          string        `json:"slack_notified"`
	AddressBusinessName    string        `json:"address_business_name"`
	AddressStreet1         string        `json:"address_street_1"`
	AddressStreet2         string        `json:"address_street_2"`
	AddressCity            string        `json:"address_city"`
	AddressZip             string        `json:"address_zip"`
	CountryID              string        `json:"country_id"`
	CityID                 string        `json:"city_id"`
	Latitude               string        `json:"latitude"`
	Longitude              string        `json:"longitude"`
	Gridref                string        `json:"gridref"`
	AvgRating              interface{}   `json:"avg_rating"`
	AvgRatingCount         interface{}   `json:"avg_rating_count"`
	IsLowInterest          string        `json:"is_low_interest"`
	IsFeatured             string        `json:"is_featured"`
	ProfileData            interface{}   `json:"profile_data"`
	EventLogo              string        `json:"event_logo"`
	EventHero              string        `json:"event_hero"`
	IsImagesContain        string        `json:"is_images_contain"`
	IsImagesMobileFull     string        `json:"is_images_mobile_full"`
	IsImagesWhite          string        `json:"is_images_white"`
	IsImagesAlignBottom    string        `json:"is_images_align_bottom"`
	IsImagesAlignCenter    string        `json:"is_images_align_center"`
	IsInstagramContain     string        `json:"is_instagram_contain"`
	CardsGenerated         string        `json:"cards_generated"`
	HasTwitterCard         string        `json:"has_twitter_card"`
	HasFacebookCard        string        `json:"has_facebook_card"`
	IsOnlineOnly           string        `json:"is_online_only"`
	IsDeleted              string        `json:"is_deleted"`
	Deleted                interface{}   `json:"deleted"`
	Embargo                interface{}   `json:"embargo"`
	EventCityUpdated       string        `json:"event_city_updated"`
	LinkChecked            string        `json:"link_checked"`
	EventSummaryUpdated    string        `json:"event_summary_updated"`
	UtcDiff                string        `json:"utc_diff"`
	CommentCount           string        `json:"comment_count"`
	TweetCount             string        `json:"tweet_count"`
	TweetUserCount         string        `json:"tweet_user_count"`
	HotnessValue           string        `json:"hotness_value"`
	IsHot                  string        `json:"is_hot"`
	RecurMode              interface{}   `json:"recur_mode"`
	RecurEnds              interface{}   `json:"recur_ends"`
	RecurUpdated           interface{}   `json:"recur_updated"`
	RecurLastAdded         interface{}   `json:"recur_last_added"`
	RecurSourceID          interface{}   `json:"recur_source_id"`
	FacebookPostID         interface{}   `json:"facebook_post_id"`
	PageViewsTotal         string        `json:"page_views_total"`
	PageViews7Days         string        `json:"page_views_7days"`
	PageViewsMax           string        `json:"page_views_max"`
	APIViewsTotal          string        `json:"api_views_total"`
	APIViews7Days          string        `json:"api_views_7days"`
	FacebookViewsTotal     string        `json:"facebook_views_total"`
	FacebookViews7Days     string        `json:"facebook_views_7days"`
	TotalViewsTotal        string        `json:"total_views_total"`
	TotalViews7Days        string        `json:"total_views_7days"`
	TotalViewsMax          string        `json:"total_views_max"`
	PaidViewsTotal         string        `json:"paid_views_total"`
	PaidViews7Days         string        `json:"paid_views_7days"`
	PageViewsUpdated       string        `json:"page_views_updated"`
	IsUnconfirmed          string        `json:"is_unconfirmed"`
	IsCancelled            string        `json:"is_cancelled"`
	IsSoldOut              string        `json:"is_sold_out"`
	CancelledNotified      interface{}   `json:"cancelled_notified"`
	IsUnlisted             string        `json:"is_unlisted"`
	IsLimitedRelease       string        `json:"is_limited_release"`
	IsNoSocialSharing      string        `json:"is_no_social_sharing"`
	AutolistAfterEventID   interface{}   `json:"autolist_after_event_id"`
	AdSpend                string        `json:"ad_spend"`
	AdSpendCurrency        string        `json:"ad_spend_currency"`
	AdSpendSpentUsd        interface{}   `json:"ad_spend_spent_usd"`
	AdSpendSpentLocal      interface{}   `json:"ad_spend_spent_local"`
	NameUpdated            string        `json:"name_updated"`
	FollowerCount          string        `json:"follower_count"`
	FollowerCount7Days     string        `json:"follower_count_7days"`
	TicketMode             string        `json:"ticket_mode"`
	VipLevel               string        `json:"vip_level"`
	VipRedirectID          interface{}   `json:"vip_redirect_id"`
	HasVipExclusive        string        `json:"has_vip_exclusive"`
	IsExcludePlatinum      string        `json:"is_exclude_platinum"`
	IsHeadliner            string        `json:"is_headliner"`
	HeadlinerCopyEn        interface{}   `json:"headliner_copy_en"`
	HeadlinerCopyZh        interface{}   `json:"headliner_copy_zh"`
	ShareCount             string        `json:"share_count"`
	Cca2                   string        `json:"cca2"`
	SummaryEn              string        `json:"summary_en"`
	SummaryZh              string        `json:"summary_zh"`
	IsPast                 string        `json:"is_past"`
	IsOnNow                string        `json:"is_on_now"`
	FormattedAddress       string        `json:"formatted_address"`
	DetailURL              string        `json:"detail_url"`
	EventLogoPin           string        `json:"event_logo_pin"`
	EventHeroURL           string        `json:"event_hero_url"`
	EventThumbnailURL      string        `json:"event_thumbnail_url"`
	ImageAlign             string        `json:"image_align"`
	IsSponsored            int           `json:"is_sponsored"`
	Summary                string        `json:"summary"`
	VenueSlug              string        `json:"venue_slug"`
	IsSpecialVenue         int           `json:"is_special_venue"`
	Dates                  []interface{} `json:"dates"`
	OpenNow                bool          `json:"open_now"`
	HasTimes               bool          `json:"has_times"`
	Timezone               string        `json:"timezone"`
	DatepickerFormat       string        `json:"datepicker_format"`
	Cities                 []string      `json:"cities"`
	Links                  []Links       `json:"links"`
	Designers              []interface{} `json:"designers"`
	OnlineSales            []interface{} `json:"online_sales"`
	IsSale                 string        `json:"is_sale"`
	IsPro                  string        `json:"is_pro"`
	TypeNameEn             string        `json:"type_name_en"`
	TypeNameZh             string        `json:"type_name_zh"`
	TypeName               string        `json:"type_name"`
	TypeSlug               string        `json:"type_slug"`
	DefaultImageID         string        `json:"default_image_id"`
	CommentUsers           string        `json:"comment_users"`
	Comments               Comments      `json:"comments"`
	SocialImages           SocialImages  `json:"social_images"`
	EventDescriptionHTML   string        `json:"event_description_html"`
	HasVipContent          bool          `json:"has_vip_content"`
	News                   []interface{} `json:"news"`
}

// Links returns website url
type Links struct {
	Website string `json:"website"`
}

// Comments is comments for a event
type Comments struct {
	EventID  string        `json:"event_id"`
	HasMore  bool          `json:"has_more"`
	Comments []interface{} `json:"comments"`
}

// SocialImages is images for a event
type SocialImages struct {
	EventID string        `json:"event_id"`
	Media   []interface{} `json:"media"`
}

// GetEvent returns Event detail based requested query.
func (s *Service) GetEvent(params *EventSearchParams) (EventDetail, *http.Response, error) {
	event := new(EventDetail)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("events_get/").QueryStruct(params).Receive(event, apiError)
	if err == nil {
		err = apiError
	}
	return *event, resp, err
}
