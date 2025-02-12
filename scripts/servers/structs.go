package main

import "time"

type Category struct {
	Name        string
	Description string
	Servers     []ServerResponse
}

type Server struct {
	URL              string
	ForceCategory    *Category
	MastodonCovenant *bool
}

type ServerResponse struct {
	// Custom config
	MastodonCovenant *bool

	// API response
	Domain      string `json:"domain"`
	Title       string `json:"title"`
	Version     string `json:"version"`
	SourceURL   string `json:"source_url"`
	Description string `json:"description"`
	Usage       struct {
		Users struct {
			ActiveMonth int `json:"active_month"`
		} `json:"users"`
	} `json:"usage"`
	Thumbnail struct {
		URL      string `json:"url"`
		Blurhash string `json:"blurhash"`
		Versions struct {
			One_X string `json:"@1x"`
			Two_X string `json:"@2x"`
		} `json:"versions"`
	} `json:"thumbnail"`
	Languages     []string `json:"languages"`
	Configuration struct {
		Urls struct {
			Streaming string `json:"streaming"`
		} `json:"urls"`
		Accounts struct {
			MaxFeaturedTags int `json:"max_featured_tags"`
		} `json:"accounts"`
		Statuses struct {
			MaxCharacters            int `json:"max_characters"`
			MaxMediaAttachments      int `json:"max_media_attachments"`
			CharactersReservedPerURL int `json:"characters_reserved_per_url"`
		} `json:"statuses"`
		MediaAttachments struct {
			SupportedMimeTypes  []string `json:"supported_mime_types"`
			ImageSizeLimit      int      `json:"image_size_limit"`
			ImageMatrixLimit    int      `json:"image_matrix_limit"`
			VideoSizeLimit      int      `json:"video_size_limit"`
			VideoFrameRateLimit int      `json:"video_frame_rate_limit"`
			VideoMatrixLimit    int      `json:"video_matrix_limit"`
		} `json:"media_attachments"`
		Polls struct {
			MaxOptions             int `json:"max_options"`
			MaxCharactersPerOption int `json:"max_characters_per_option"`
			MinExpiration          int `json:"min_expiration"`
			MaxExpiration          int `json:"max_expiration"`
		} `json:"polls"`
		Translation struct {
			Enabled bool `json:"enabled"`
		} `json:"translation"`
	} `json:"configuration"`
	Registrations struct {
		Enabled          bool        `json:"enabled"`
		ApprovalRequired bool        `json:"approval_required"`
		Message          interface{} `json:"message"`
	} `json:"registrations"`
	Contact struct {
		Email   string `json:"email"`
		Account struct {
			ID             string        `json:"id"`
			Username       string        `json:"username"`
			Acct           string        `json:"acct"`
			DisplayName    string        `json:"display_name"`
			Locked         bool          `json:"locked"`
			Bot            bool          `json:"bot"`
			Discoverable   bool          `json:"discoverable"`
			Group          bool          `json:"group"`
			CreatedAt      time.Time     `json:"created_at"`
			Note           string        `json:"note"`
			URL            string        `json:"url"`
			Avatar         string        `json:"avatar"`
			AvatarStatic   string        `json:"avatar_static"`
			Header         string        `json:"header"`
			HeaderStatic   string        `json:"header_static"`
			FollowersCount int           `json:"followers_count"`
			FollowingCount int           `json:"following_count"`
			StatusesCount  int           `json:"statuses_count"`
			LastStatusAt   string        `json:"last_status_at"`
			Noindex        bool          `json:"noindex"`
			Emojis         []interface{} `json:"emojis"`
			Fields         []struct {
				Name       string      `json:"name"`
				Value      string      `json:"value"`
				VerifiedAt interface{} `json:"verified_at"`
			} `json:"fields"`
		} `json:"account"`
	} `json:"contact"`
	Rules []struct {
		ID   string `json:"id"`
		Text string `json:"text"`
	} `json:"rules"`
}

func (s *ServerResponse) Categorize(server Server) *Category {
	if server.ForceCategory != nil {
		return server.ForceCategory
	}

	if s.Registrations.Enabled && !s.Registrations.ApprovalRequired {
		return OpenCategory
	}

	if s.Registrations.Enabled {
		return OpenCategoryWithApproval
	}

	return ClosedCategory
}
