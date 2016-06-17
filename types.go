package mashable

import "time"

type Error struct {
	Error   int    `json:"error"`
	Message string `json:"message"`
}

type ListOptions struct {
	Page    int `json:"page" url:"page"`
	PerPage int `json:"per_page" url:"per_page"`
}

type Collection struct {
	Total    int           `json:"total"`
	Page     int           `json:"page"`
	PerPage  int           `json:"per_page"`
	Pages    int           `json:"pages"`
	Flags    []interface{} `json:"flags"`
	Metadata struct {
		Targeting map[string]interface{} `json:"targeting"`
	} `json:"metadata"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
}

type Post struct {
	*Error      `json:"inline,omitempty"`
	ID          string            `json:"id"`
	SortKey     string            `json:"sort_key"`
	Title       string            `json:"title"`
	TitleTag    interface{}       `json:"title_tag"`
	Images      map[string]string `json:"images"`
	Topics      []string          `json:"topics"`
	Channel     string            `json:"channel"`
	ChannelName string            `json:"channel_name"`
	SubChannels []string          `json:"subchannels"`
	Author      string            `json:"author"`
	AuthorID    string            `json:"author_id"`
	PostDate    time.Time         `json:"post_date"`
	PostDateRfc string            `json:"post_date_rfc"`
	Link        string            `json:"link"`
	Content     struct {
		Excerpt string      `json:"excerpt"`
		Full    string      `json:"full"`
		Intro   interface{} `json:"intro"`
		Plain   string      `json:"plain"`
	} `json:"content"`
	Shares struct {
		Facebook   int `json:"facebook"`
		Twitter    int `json:"twitter"`
		GooglePlus int `json:"google_plus"`
		Pinterest  int `json:"pinterest"`
		LinkedIn   int `json:"linked_in"`
		Total      int `json:"total"`
	} `json:"shares"`
	CommentsCount int                    `json:"comments_count"`
	URL           string                 `json:"url"`
	Velocity      []int                  `json:"velocity"`
	ShortURL      string                 `json:"short_url"`
	Targeting     map[string]interface{} `json:"targeting"`
	ContentSource string                 `json:"content_source"`
	LeadType      string                 `json:"lead_type"`
	CommentsURL   string                 `json:"comments_url"`
	ShortcodeData struct {
		Gallery []interface{} `json:"gallery"`
	} `json:"shortcode_data"`
	Webview        bool   `json:"webview"`
	SponsoredBy    string `json:"sponsored_by"`
	SponsoredByURL string `json:"sponsored_by_url"`
	SeriesType     string `json:"series_type"`
	SeriesSlug     string `json:"series_slug"`
}

type PostList struct {
	*Error     `json:"inline,omitempty"`
	Collection *Collection `json:"collection"`
	Posts      []*Post     `json:"posts"`
}

type Topic struct {
	*Error        `json:"inline,omitempty"`
	Name          string `json:"name"`
	Followers     int    `json:"followers"`
	Posts         int    `json:"posts"`
	Article       string `json:"article"`
	ArticleSource struct {
		Text string `json:"text"`
		URL  string `json:"url"`
	} `json:"article_source"`
	Taxonomies []string `json:"taxonomies"`
	Thumb      string   `json:"thumb"`
	Slug       string   `json:"slug"`
	URL        string   `json:"url"`
}

type TopicList struct {
	*Error     `json:"inline,omitempty"`
	Collection *Collection `json:"collection"`
	Topics     []*Topic    `json:"topics"`
}

type User struct {
	*Error      `json:"inline,omitempty"`
	ID          string    `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	About       string    `json:"about"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"created_at"`
	Following   int       `json:"following"`
	Location    string    `json:"location"`
	Badges      int       `json:"badges"`
	Comments    int       `json:"comments"`
	Gender      string    `json:"gender"`
	Connections []string  `json:"connections"`
	ShareTo     []string  `json:"share_to"`
	URL         string    `json:"url"`
}

type SearchResult struct {
	Search struct {
		Query       string      `json:"query"`
		Status      string      `json:"status"`
		Suggestions interface{} `json:"suggestions"`
		Hits        interface{} `json:"hits"`
		Results     struct {
			Posts  []*Post  `json:"posts"`
			Users  []*User  `json:"users"`
			Topics []*Topic `json:"topics"`
		} `json:"results"`
	} `json:"search"`
}
