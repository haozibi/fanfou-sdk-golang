package fanfou

// Status statuses data
type Status struct {
	CreatedAt           string  `json:"created_at"`
	ID                  string  `json:"id"`
	RawID               int64   `json:"rawid"`
	Text                string  `json:"text"`
	Source              string  `json:"source"`
	Location            string  `json:"location"`
	Truncated           bool    `json:"truncated"`
	InReplyToStatusID   string  `json:"in_reply_to_status_id"`
	InReplyToUserID     string  `json:"in_reply_to_user_id"`
	InReplyToScreenName string  `json:"in_reply_to_screen_name"`
	RepostStatusID      string  `json:"repost_status_id"`
	RepostStatus        *Status `json:"repost_status"`
	RepostUserID        string  `json:"repost_user_id"`
	RepostScreenName    string  `json:"repost_screen_name"`
	Favorited           bool    `json:"favorited"`
	User                *User   `json:"user"`
	Photo               *Photo  `json:"photo"`
}

// User User
type User struct {
	ID                        string  `json:"id"`
	Name                      string  `json:"name"`
	ScreenName                string  `json:"screen_name"`
	Location                  string  `json:"location"`
	Gender                    string  `json:"gender"`
	Birthday                  string  `json:"birthday"`
	Description               string  `json:"description"`
	ProfileImageURL           string  `json:"profile_image_url"`
	ProfileImageURLLarge      string  `json:"profile_image_url_large"`
	URL                       string  `json:"url"`
	Protected                 bool    `json:"protected"`
	FollowersCount            int     `json:"followers_count"`
	FriendsCount              int     `json:"friends_count"`
	FavouritesCount           int     `json:"favourites_count"`
	StatusesCount             int64   `json:"statuses_count"`
	Following                 bool    `json:"following"`
	Notifications             bool    `json:"notifications"`
	CreatedAt                 string  `json:"created_at"`
	UtcOffset                 int64   `json:"utc_offset"`
	ProfileBackgroundColor    string  `json:"profile_background_color"`
	ProfileTextColor          string  `json:"profile_text_color"`
	ProfileLinkColor          string  `json:"profile_link_color"`
	ProfileSidebarFillColor   string  `json:"profile_sidebar_fill_color"`
	ProfileSidebarBorderColor string  `json:"profile_sidebar_border_color"`
	ProfileBackgroundImageURL string  `json:"profile_background_image_url"`
	ProfileBackgroundTile     bool    `json:"profile_background_tile"`
	Status                    *Status `json:"status"`
}

// Photo Photo
type Photo struct {
	Imageurl string `json:"imageurl"`
	Thumburl string `json:"thumburl"`
	Largeurl string `json:"largeurl"`
}

// Trends Trends
type Trends struct {
	Name  string `json:"name"`
	Query string `json:"query"`
	URL   string `json:"url"`
}

// SavedSearches SavedSearches
type SavedSearches struct {
	ID        int    `json:"id"`
	Query     string `json:"query"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

// DirectMessages DirectMessages
type DirectMessages struct {
	ID                  string          `json:"id"`
	Text                string          `json:"text"`
	SenderID            string          `json:"sender_id"`
	RecipientID         string          `json:"recipient_id"`
	CreatedAt           string          `json:"created_at"`
	SenderScreenName    string          `json:"sender_screen_name"`
	RecipientScreenName string          `json:"recipient_screen_name"`
	Sender              *User           `json:"sender"`
	Recipient           *User           `json:"recipient"`
	InReplyTo           *DirectMessages `json:"in_reply_to"`
}

var (
	allowParams = map[string][]string{
		"search/public_timeline.json": {
			"q",
			"since_id",
			"max_id",
			"count",
			"mode",
			"format",
			"callback",
		},
		"search/users.json": {
			"q",
			"count",
			"page",
			"mode",
			"format",
			"callback",
		},
		"search/user_timeline.json": {
			"q",
			"id",
			"since_id",
			"max_id",
			"count",
			"mode",
			"format",
			"callback",
		},
		"blocks/blocking.json": {
			"mode",
			"page",
			"count",
		},
		"blocks/create.json": {
			"id",
			"mode",
			"format",
			"callback",
		},
		"blocks/exists.json": {
			"id",
			"mode",
		},
		"blocks/destroy.json": {
			"id",
			"mode",
		},
		"users/tagged.json": {
			"tag",
			"count",
			"page",
			"mode",
			"format",
			"callback",
		},
		"users/show.json": {
			"id",
			"mode",
			"format",
			"callback",
		},
		"users/tag_list.json": {
			"id",
			"callback",
		},
		"users/followers.json": {
			"id",
			"count",
			"page",
			"mode",
			"format",
			"callback",
		},
		"2/users/recommendation.json": {
			"count",
			"page",
			"mode",
			"format",
			"callback",
		},
		"/2/users/cancel_recommendation.json": {
			"id",
			"mode",
			"format",
			"callback",
		},
		"users/friends.json": {
			"id",
			"count",
			"page",
			"mode",
			"format",
			"callback",
		},
		"account/verify_credentials.json": {
			"mode",
			"format",
			"callback",
		},
		"account/update_profile_image.json": {
			"mode",
			"format",
			"callback",
		},
		"account/rate_limit_status.json": {
			"callback",
		},
		"account/update_profile.json": {
			"url",
			"mode",
			"callback",
			"location",
			"description",
			"name",
			"email",
		},
		"account/update_notify_num.json": {
			"notify_num",
		},
		"saved_searches/create.json": {
			"query",
			"callback",
		},
		"saved_searches/destroy.json": {
			"id",
			"callback",
		},
		"saved_searches/show.json": {
			"id",
			"callback",
		},
		"saved_searches/list.json": {
			"callback",
		},
		"photos/user_timeline.json": {
			"id",
			"since_id",
			"max_id",
			"count",
			"page",
			"mode",
			"format",
			"callback",
		},
		"photos/upload.json": {
			"status",
			"source",
			"location",
			"mode",
			"format",
			"callback",
		},
		"trends/list.json": {
			"callback",
		},
		"followers/ids.json": {
			"id",
			"page",
			"count",
			"callback",
		},
		"favorites/destroy/": {
			"mode",
			"format",
			"callback",
		},
		"favorites/": {
			"id",
			"page",
			"count",
			"mode",
			"format",
			"callback",
		},
		"favorites/create/": {
			"id",
			"mode",
			"format",
			"callback",
		},
		"friendships/create.json": {
			"id",
			"mode",
		},
		"friendships/destroy.json": {
			"id",
			"mode",
			"format",
			"callback",
		},
		"friendships/requests.json": {
			"page",
			"count",
			"mode",
			"format",
			"callback",
		},
		"friendships/deny.json": {
			"id",
			"mode",
			"format",
			"callback",
		},
		"friendships/exists.json": {
			"user_a",
			"user_b",
		},
		"friendships/accept.json": {
			"id",
			"mode",
			"format",
			"callback",
		},
		"friendships/show.json": {
			"source_login_name",
			"source_id",
			"target_login_name",
			"target_id",
		},
		"friends/ids.json": {
			"id",
			"page",
			"count",
			"callback",
		},
		"statuses/destroy.json": {
			"id",
			"mode",
			"format",
			"callback",
		},
		"statuses/home_timeline.json": {
			"id",
			"since_id",
			"max_id",
			"count",
			"page",
			"mode",
			"format",
			"callback",
		},
		"statuses/public_timeline.json": {
			"count",
			"since_id",
			"max_id",
			"mode",
			"format",
			"callback",
		},
		"statuses/replies.json": {
			"since_id",
			"max_id",
			"count",
			"page",
			"mode",
			"format",
			"callback",
		},
		"statuses/followers.json": {
			"id",
			"count",
			"page",
			"mode",
			"format",
			"callback",
		},
		"statuses/update.json": {
			"status",
			"in_reply_to_status_id",
			"in_reply_to_user_id",
			"repost_status_id",
			"source",
			"mode",
			"format",
			"location",
			"callback",
		},
		"statuses/user_timeline.json": {
			"id",
			"since_id",
			"max_id",
			"count",
			"page",
			"mode",
			"format",
			"callback",
		},
		"statuses/friends.json": {
			"id",
			"count",
			"page",
			"mode",
			"callback",
		},
		"statuses/context_timeline.json": {
			"id",
			"mode",
			"format",
			"callback",
		},
		"statuses/mentions.json": {
			"since_id",
			"max_id",
			"count",
			"page",
			"mode",
			"format",
			"callback",
		},
		"statuses/show.json": {
			"id",
			"mode",
			"format",
			"callback",
		},
		"direct_messages/destroy.json": {
			"id",
			"callback",
		},
		"direct_messages/conversation.json": {
			"id",
			"count",
			"page",
			"since_id",
			"max_id",
			"mode",
			"callback",
		},
		"direct_messages/new.json": {
			"user",
			"text",
			"in_reply_to_id",
			"mode",
			"callback",
		},
		"direct_messages/conversation_list.json": {
			"count",
			"page",
			"mode",
			"callback",
		},
		"direct_messages/inbox.json": {
			"count",
			"page",
			"since_id",
			"max_id",
			"mode",
			"callback",
		},
		"direct_messages/sent.json": {
			"count",
			"page",
			"since_id",
			"max_id",
			"mode",
			"callback",
		},
	}
)
