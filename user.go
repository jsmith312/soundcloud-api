package user

//User SoundCloud user field when calling /user
type User struct {
	ID                    int    `json:"id"`
	Permalink             string `json:"permalink"`
	Username              string `json:"username"`
	URI                   string `json:"uri"`
	PermalinkURL          string `json:"permalink_url"`
	AvatarURL             string `json:"avatar_url"`
	Country               string `json:"country"`
	FullName              string `json:"full_name"`
	City                  string `json:"city"`
	Description           string `json:"description"`
	DiscogsName           string `json:"discogs_name"`
	MySpaceName           string `json:"myspace_name"`
	Website               string `json:"website"`
	WebsiteTitle          string `json:"website_title"`
	Online                bool   `json:"online"`
	TrackCount            int    `json:"track_count"`
	PlaylistCount         int    `json:"playlist_count"`
	FollowersCount        int    `json:"followers_count"`
	FollowingsCount       int    `json:"followings_count"`
	PublicFavoritesCount  int    `json:"public_favorites_count"`
	Plan                  string `json:"plan"`
	PrivateTracksCount    int    `json:"private_tracks_count"`
	PrivatePlaylistCount  int    `json:"private_playlists_count"`
	PrimaryEmailConfirmed int    `json:"primary_email_confirmed"`
}

//GetNumTracks
func (u User) GetNumTracks() int {
	return u.TrackCount
}
