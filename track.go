package soundcloud

//Track structure
type Track struct {
	ID                  int         `json:"id"`
	CreatedAt           string      `json:"created_at"`
	UserID              int         `json:"user_id"`
	Duration            int         `json:"duration"`
	Commentable         bool        `json:"commentable"`
	State               string      `json:"state"`
	Sharing             string      `json:"sharing"`
	TagList             string      `json:"tag_list"`
	Permalink           string      `json:"permalink"`
	Description         string      `json:"decription"`
	Streamable          bool        `json:"streamable"`
	Downloadable        bool        `json:"downloadable"`
	Genre               string      `json:"genre"`
	Release             int         `json:"release"`
	PurchaseURL         string      `json:"purchase_url"`
	LabelID             int         `json:"label_id"`
	LabelName           string      `json:"label_name"`
	Isrc                string      `json:"isrc"`
	VideoURL            string      `json:"video_url"`
	TrackType           string      `json:"recording"`
	KeySignature        string      `json:"key_signature"`
	BPM                 int         `json:"bpm"`
	Title               string      `json:"title"`
	ReleaseYear         int         `json:"release_year"`
	ReleaseMonth        int         `json:"release_month"`
	ReleaseDay          int         `json:"release_day"`
	OriginalFormat      string      `json:"original_format"`
	OriginalContentSize int         `json:"original_content_size"`
	License             string      `json:"license"`
	URI                 string      `json:"uri"`
	PermalinkURL        string      `json:"permalink_url"`
	ArtworkURL          string      `json:"artwork_url"`
	WaveformURL         string      `json:"waveform_url"`
	User                ShortUser   `json:"user"`
	StreamURL           string      `json:"stream_url"`
	DownloadURL         string      `json:"download_url"`
	PlaybackCount       int         `json:"playback_count"`
	DownloadCount       int         `json:"download_count"`
	FavoritingsCount    int         `json:"favoritings_count"`
	CommentCount        int         `json:"comment_count"`
	CreatedWith         CreatedWith `json:"created_with"`
	AttachmentsURI      string      `json:"attachments_uri"`
}

//ShortUser fields
type ShortUser struct {
	ID           int    `json:"id"`
	Permalink    string `json:"permalink"`
	Username     string `json:"username"`
	URI          string `json:"uri"`
	PermalinkURL string `json:"permalink_url"`
	AvatarURL    string `json:"avatar_url"`
}

//CreatedWith struct
type CreatedWith struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	URI          string `json:"uri"`
	PermalinkURL string `json:"permalink_url"`
}
