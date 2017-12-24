package refract

// A Track represents a video or a song;
type Track struct {
	ID       uint64 `json:"id"`
	Title    string `json:"title"`
	Duration uint64 `json:"duration"`
}

// A YoutubeTrack represents a youtube video
type YoutubeTrack struct {
	Track
	URL string `json:"url"`
}

// A Team represents a slack team
// there will usually only be one slack team.
type Team struct {
	Name   string `name:"string"`
	Domain string `domain:"string"`
}
