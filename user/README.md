## User
Contains the user properties specified in the [Soundcloud API Docs](https://developers.soundcloud.com/docs/api/reference#me)
```
type User struct {
  ID                    int    
  Permalink             string
  Username              string
  URI                   string
  PermalinkURL          string
  AvatarURL             string
  Country               string
  FullName              string
  City                  string
  Description           string
  DiscogsName           string
  MySpaceName           string
  Website               string
  WebsiteTitle          string
  Online                bool
  TrackCount            int
  PlaylistCount         int
  FollowersCount        int
  FollowingsCount       int
  PublicFavoritesCount  int
  Plan                  string
  PrivateTracksCount    int
  PrivatePlaylistCount  int
  PrimaryEmailConfirmed int
}
```
