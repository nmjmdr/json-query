package main

import "fmt"
import "filter"
import "os"
import "queryfunctions"

func main() {
	type Message struct {
		Drm          bool
		EpisodeCount int
	}
	var m Message
	var found []Message

	f := filter.NewItemsFilter(&m, func(v interface{}) {
		foundMessage := v.(*Message)
		found = append(found, *foundMessage)
	})

	reader, err := os.Open("./sample.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	fieldQuery1 := queryfunctions.IsTrue("drm")
	fieldQuery2 := queryfunctions.IsGreaterThanN("episodeCount", 3)
	err = f.Filter(reader, []filter.FieldQuery{fieldQuery1, fieldQuery2})
	fmt.Println(found)
}
