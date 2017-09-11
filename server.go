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
    w.WriteHeader(http.StatusBadRequest)
    json.NewEncoder(w).Encode(ErrorResponse{ Error: "Could not decode request: JSON parsing failed"})
    return
  }

  if request.Payload == nil || len(request.Payload) == 0 {
     w.WriteHeader(http.StatusBadRequest)
    json.NewEncoder(w).Encode(ErrorResponse{ Error: "No items supplied in payload"})
    return
  }

  var results []theshows.Result
  var show theshows.Show

	itemsFilter := filter.NewItemsFilter(&show)

  fieldQueryDrm := queryfunctions.IsTrue("drm")
  fieldQueryEpisodeCount := queryfunctions.IsGreaterThanN("episodeCount", 0)

  for _, item := range request.Payload {
    if itemsFilter.IsMatch(item, []filter.FieldQuery{fieldQueryDrm, fieldQueryEpisodeCount}) {
      results = append(results, theshows.ToResult(item))
    }
  }
  type responseObject struct {
    Response []theshows.Result
  }
  responseToSend :=  responseObject{ Response: results }
  fmt.Println(responseToSend)
  json.NewEncoder(w).Encode(responseToSend)
}

func main() {
  router := mux.NewRouter()
  router.HandleFunc("/", RootHandler).Methods("POST")
  log.Fatal(http.ListenAndServe(":9090", router))
}
