## Track
Contains the track properties specified in the [Soundcloud API Docs](https://developers.soundcloud.com/docs/api/reference#tracks)
```
//Track structure
type Track struct {
	ID                  int         
	CreatedAt           string      
	UserID              int         
	Duration            int         
	Commentable         bool        
	State               string      
	Sharing             string      
	TagList             string      
	Permalink           string      
	Description         string      
	Streamable          bool        
	Downloadable        bool        
	Genre               string      
	Release             int         
	PurchaseURL         string     
	LabelID             int        
	LabelName           string     
	Isrc                string     
	VideoURL            string     
	TrackType           string     
	KeySignature        string     
	BPM                 int        
	Title               string     
	ReleaseYear         int         
	ReleaseMonth        int         
	ReleaseDay          int         
	OriginalFormat      string      
	OriginalContentSize int         
	License             string      
	URI                 string      
	PermalinkURL        string      
	ArtworkURL          string      
	WaveformURL         string      
	User                ShortUser  
	StreamURL           string     
	DownloadURL         string     
	PlaybackCount       int        
	DownloadCount       int        
	FavoritingsCount    int        
	CommentCount        int        
	CreatedWith         CreatedWith
	AttachmentsURI      string      
}

//ShortUser fields
type ShortUser struct {
	ID           int    
	Permalink    string
	Username     string
	URI          string
	PermalinkURL string
	AvatarURL    string
}

//CreatedWith struct
type CreatedWith struct {
	ID           int   
	Name         string
	URI          string
	PermalinkURL string
}

```
