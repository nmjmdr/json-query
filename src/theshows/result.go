package theshows

type Result struct {
  Image string
  Slug string
  Title string
}

func toResult(show Show) {
  return Result { Image: show.Image.ShowImage, Slug: show.Slug, Title: show.Title }
}
