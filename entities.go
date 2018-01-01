package refract

type TrackType int

const (
	_ TrackType = iota
	TrackTypeYoutube
)

// A Track represents a video or a song;
type Track struct {
	ID       uint64    `json:"id"`
	Title    string    `json:"title"`
	Duration uint64    `json:"duration"`
	URL      string    `json:"url"`
	Type     TrackType `json:"trackType"`
}

// A Team represents a slack team
// there will usually only be one slack team.
type Team struct {
	ID     uint64 `json:"id"`
	Name   string `json:"name"`
	Domain string `json:"domain"`
}
