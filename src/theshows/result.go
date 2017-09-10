package theshows

type Result struct {
  Image string
  Slug string
  Title string
}

func ToResult(show Show) Result {
  return Result { Image: show.Image.ShowImage, Slug: show.Slug, Title: show.Title }
}
