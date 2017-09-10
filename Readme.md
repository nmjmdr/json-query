# JSON-Query

I have developed JSON-Query as a generic libraray to query JSON data. The http server specifies the particualr conditions of (drm == true and episodeCount > 0) using the developed json-query libraray.

  - JSON-query library can be used to query any JSON data
  - Currently only two query functions have been built (IsTrue and IsGreaterThanN). But the library can be extended to supply custom query functions

# Usage:

POST: http://ec2-52-63-121-208.ap-southeast-2.compute.amazonaws.com/
With appropriate JSON data set as body, returns the required results

![Usage Image]()

### Build & Tests
_Downlaod & Setup_:
1. git clone https://github.com/nmjmdr/json-query.git
2. set GOPATH to the cloned directory. For example on Mac export GOPATH=path-to-cloned-dir (I know that this not the standard, ideally I would have created the folder structure as reccomended by GO)

_Building:_
1. `cd` to cloned directory and run `go build`. This should create an executable in the current directory

_Running the unit tests:_
1. The tests are included in tests folder
1. `cd` to cloned directory and run `go test ./tests`. This should run the unit tests

### Design
The main aim was to keep the design simple and generic. Especially it is beneficial if the logic to filter items is kept generic and does not depend upon the structure of JSON. This library acheives this using `reflect` library.

Queries can be specified as an array of type `FieldQuery` which has the following structure:
```
type Compare func(reflect.Value) bool
type FieldQuery struct {
	Field   string
	Compare Compare
}
```
The Compare functions for FieldQuery can be set using the built-in `queryfunctions`. Currently I have defined only two queryfunctions:
```
func IsTrue(fieldName string) filter.FieldQuery {
	fieldName = strings.Title(fieldName)
	f := func(v reflect.Value) bool {
		val := v.Bool()
		return val == true
	}
	return filter.FieldQuery{Field: fieldName, Compare: f}
}
```
and
```
func IsGreaterThanN(fieldName string, number int) filter.FieldQuery {
	fieldName = strings.Title(fieldName)
	f := func(v reflect.Value) bool {
		val := int(v.Int())
		return val > number
	}
	return filter.FieldQuery{Field: fieldName, Compare: f}
}
```

The only place in current design where the query is created is in server.go:
```
itemsFilter := filter.NewItemsFilter(&show)
fieldQueryDrm := queryfunctions.IsTrue("drm")
fieldQueryEpisodeCount := queryfunctions.IsGreaterThanN("episodeCount", 0)
for _, item := range request.Payload {
    if itemsFilter.IsMatch(item, []filter.FieldQuery{fieldQueryDrm, fieldQueryEpisodeCount}){
      results = append(results, theshows.ToResult(item))
    }
}
```
