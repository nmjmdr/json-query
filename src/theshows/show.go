package theshows

type Image struct {
  ShowImage string
}

type Season struct {
  Slug string
}

type Show struct {
  Country string
  Description string
  Drm bool
  EpisodeCount int
  Genre string
  Image Image
  Language string
  NextEpisode interface{}
  PrimaryColour string
  Seasons []Season
  Slug string
  Title string
  TvChannel string
}

type Request struct {
  Payload []Show
}
