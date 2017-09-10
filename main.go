package main

import "fmt"
import "filter"
import "os"
import "queryfunctions"

func main() {
	type Message struct {
		Drm  bool
		EpisodeCount int
	}
	var m Message
	f := filter.NewItemsFilter(&m)

	reader, err := os.Open("./sample.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	var values []interface{}


	fieldQuery :=  queryfunctions.IsTrue("Drm")
	values,err = f.Filter(reader, []filter.FieldQuery{fieldQuery})

	for _, value := range values {
		m, ok := value.(*Message)
		fmt.Println(m, ok)
	}

}
