package main

import "github.com/gorilla/mux"
import "net/http"
import "log"
import "fmt"
import "filter"
import "queryfunctions"
import "theshows"
import "encoding/json"

var itemsFilter filter.ItemsFilter

type ErrorResponse struct {
  Error string
}

func RootHandler(w http.ResponseWriter, r *http.Request) {

  w.Header().Set("content-type", "application/json")

  var request theshows.Request
  err := json.NewDecoder(r.Body).Decode(&request)
  if err != nil {
    json.NewEncoder(w).Encode(ErrorResponse{ Error: "Could not decode request: JSON parsing failed"})
    return
  }
  fmt.Println(request.Payload)
  var results []theshows.Result
  var show theshows.Show

	itemsFilter := filter.NewItemsFilter(&show, func(value interface{}) {
		found := value.(*theshows.Show)
		result := theshows.ToResult(*found)
		results = append(results, result)
	})

  fieldQuery1 := queryfunctions.IsTrue("drm")
  fieldQuery2 := queryfunctions.IsGreaterThanN("episodeCount", 3)

  err = itemsFilter.Filter(r.Body, []filter.FieldQuery{fieldQuery1, fieldQuery2})
  fmt.Println(results, err)
  json.NewEncoder(w).Encode(results)
}

func main() {
  router := mux.NewRouter()
  router.HandleFunc("/", RootHandler).Methods("POST")
  log.Fatal(http.ListenAndServe(":9090", router))
}
