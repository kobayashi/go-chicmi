package chicmi

import (
	"net/http"
)

// Results is used as returned parameters for Search
type Results struct {
	Values struct {
		Timezone   string        `json:"timezone"`
		CityID     int           `json:"city_id"`
		Events     []Event       `json:"events"`
		Headliners []interface{} `json:"headliners"`
		VipHidden  int           `json:"vip_hidden"`
	} `json:"values"`
	Meta Meta `json:"meta"`
}

// Event returns calender event
type Event struct {
	AddressBusinessName    string      `json:"address_business_name"`
	AddressStreet1         string      `json:"address_street_1"`
	AddressStreet2         string      `json:"address_street_2"`
	AddressCity            string      `json:"address_city"`
	AddressZip             string      `json:"address_zip"`
	CountryID              string      `json:"country_id"`
	EventID                string      `json:"event_id"`
	EventNameEn            string      `json:"event_name_en"`
	EventNameZh            string      `json:"event_name_zh"`
	EventIsFeatured        string      `json:"event_is_featured"`
	IsSponsored            string      `json:"is_sponsored"`
	EventEligibilityTypeID string      `json:"event_eligibility_type_id"`
	EventLogo              string      `json:"event_logo"`
	EventHero              string      `json:"event_hero"`
	StartDate              string      `json:"start_date"`
	UtcDiff                string      `json:"utc_diff"`
	EndDate                string      `json:"end_date"`
	EventTypeID            string      `json:"event_type_id"`
	Latitude               string      `json:"latitude"`
	Longitude              string      `json:"longitude"`
	IsHot                  string      `json:"is_hot"`
	Created                string      `json:"created"`
	PageViewsTotal         string      `json:"page_views_total"`
	HotnessValue           string      `json:"hotness_value"`
	CommentCount           string      `json:"comment_count"`
	IsCancelled            string      `json:"is_cancelled"`
	IsSoldOut              string      `json:"is_sold_out"`
	Slug                   string      `json:"slug"`
	Icons                  string      `json:"icons"`
	VipLevel               string      `json:"vip_level"`
	TicketMode             string      `json:"ticket_mode"`
	IsImagesAlignBottom    string      `json:"is_images_align_bottom"`
	IsImagesAlignCenter    string      `json:"is_images_align_center"`
	IsHeadliner            string      `json:"is_headliner"`
	HeadlinerCopyEn        interface{} `json:"headliner_copy_en"`
	HeadlinerCopyZh        interface{} `json:"headliner_copy_zh"`
	ShareCount             string      `json:"share_count"`
	Days                   string      `json:"days"`
	SummaryEn              string      `json:"summary_en"`
	SummaryZh              string      `json:"summary_zh"`
	LocationEn             string      `json:"location_en"`
	LocationZh             string      `json:"location_zh"`
	Price                  string      `json:"price"`
	ActionType             string      `json:"action_type"`
	ActionURL              string      `json:"action_url"`
	EventName              string      `json:"event_name"`
	Summary                string      `json:"summary"`
	Location               string      `json:"location"`
	HeadlinerCopy          interface{} `json:"headliner_copy"`
	FormattedAddress       string      `json:"formatted_address"`
	DetailURL              string      `json:"detail_url"`
	DateLabel              string      `json:"date_label"`
	DateGroup              string      `json:"date_group"`
	DateLocal              string      `json:"date_local"`
	EventType              string      `json:"event_type"`
	EventTypeNoun          string      `json:"event_type_noun"`
	EventTagColour         string      `json:"event_tag_colour"`
	EventEligibilityType   string      `json:"event_eligibility_type"`
	EventHeroURL           string      `json:"event_hero_url"`
	EventPreviewURL        string      `json:"event_preview_url"`
	EventThumbURL          string      `json:"event_thumb_url"`
	EventCardURL           string      `json:"event_card_url"`
	ImageAlign             string      `json:"image_align"`
	EventLogoPin           string      `json:"event_logo_pin"`
	EventLogoURL           string      `json:"event_logo_url"`
}

// EventType is type of event
type EventType struct {
	ID              string `json:"id"`
	TypeNameEn      string `json:"type_name_en"`
	TypeNameZh      string `json:"type_name_zh"`
	TypeNameRu      string `json:"type_name_ru"`
	TypeNameAr      string `json:"type_name_ar"`
	TypeNounEn      string `json:"type_noun_en"`
	Slug            string `json:"slug"`
	Emoji           string `json:"emoji"`
	Tags            string `json:"tags"`
	IsSale          string `json:"is_sale"`
	IsPro           string `json:"is_pro"`
	HasSectors      string `json:"has_sectors"`
	IsDefaultFilter string `json:"is_default_filter"`
	SortOrder       string `json:"sort_order"`
	DefaultImageID  string `json:"default_image_id"`
	TypeTitleEn     string `json:"type_title_en"`
	TypeTitleZh     string `json:"type_title_zh"`
	TagColour       string `json:"tag_colour"`
	TypeImage       string `json:"type_image"`
	TypeLinkEn      string `json:"type_link_en"`
	TypeLinkZh      string `json:"type_link_zh"`
	TypeName        string `json:"type_name"`
	TypeNoun        string `json:"type_noun"`
	TypeLink        string `json:"type_link"`
}

// SectorType is type of sector
type SectorType struct {
	ID                 string      `json:"id"`
	TypeNameEn         string      `json:"type_name_en"`
	TypeNameZh         string      `json:"type_name_zh"`
	TypeNameRu         interface{} `json:"type_name_ru"`
	TypeNameAr         interface{} `json:"type_name_ar"`
	TypeNounEn         string      `json:"type_noun_en"`
	TypeNounZh         string      `json:"type_noun_zh"`
	TypeNounRu         interface{} `json:"type_noun_ru"`
	TypeNounAr         interface{} `json:"type_noun_ar"`
	IsDefaultFilter    string      `json:"is_default_filter"`
	IsEventSector      string      `json:"is_event_sector"`
	IsOnlineSaleSector string      `json:"is_online_sale_sector"`
	SortOrder          string      `json:"sort_order"`
	Slug               string      `json:"slug"`
	TypeName           string      `json:"type_name"`
	TypeNoun           string      `json:"type_noun"`
}

// EventEligibilityType is type of event eligibility
type EventEligibilityType struct {
	ID              string `json:"id"`
	TypeNameEn      string `json:"type_name_en"`
	TypeNameZh      string `json:"type_name_zh"`
	TypeNameRu      string `json:"type_name_ru"`
	TypeNameAr      string `json:"type_name_ar"`
	IsDefaultFilter string `json:"is_default_filter"`
	SortOrder       string `json:"sort_order"`
	Slug            string `json:"slug"`
	TypeName        string `json:"type_name"`
}

// Meta is meta for Results
type Meta struct {
	Version    string `json:"version"`
	ServerTime int    `json:"server-time"`
	Step       int    `json:"step"`
	Geoip      string `json:"geoip"`
	Language   string `json:"language"`
	Remote     bool   `json:"remote"`
}

// Search returns events from chicmi based requested query.
func (s *Service) Search(params *SearchParams) (Results, *http.Response, error) {
	events := new(Results)
	apiError := new(APIError)
	resp, err := s.sling.New().Get("calendar_in_city/").QueryStruct(params).Receive(events, apiError)
	if err == nil {
		err = apiError
	}
	return *events, resp, err
}
