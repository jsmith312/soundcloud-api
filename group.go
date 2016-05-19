package soundcloud

//Group SoundCloud Group field
type Group struct {
	ID               int    `json:"id"`
	URI              string `json:"uri"`
	CreatedAt        string `json:"created_at"`
	Permalink        string `json:"permalink"`
	PermalinkURL     string `json:"permalink_url"`
	ArtworkURL       string `json:"artwork_url"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	ShortDescription string `json:"short_description"`
	Creator          string `json:"creator"`
}
