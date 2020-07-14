package collection

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

var (
	NonMapTypeError = errors.New("the given parameter v is a non-map type, " +
		"parameter v should be a map")
)

// The New() function creates a new Collection.
func New() *Collection {
	return &Collection{ make([]*Pair, 0) }
}

// The NewFromJsonString() function parses a given JSON string, and returns
// the result as a new Collection. It uses "encoding/json" package, so it will
// returns an InvalidUnmarshalError if v is nil or not a pointer. Because
// the "encoding/json" package uses Go native map internally which does not
// support element ordering, the returned Collection will not have the same
// order as defined in the JSON string. Instead, the key ordering will be
// sorted alphabetically.
func NewFromJsonString(v string) (*Collection, error) {
	object := make(map[string]interface{})
	decoder := json.NewDecoder(strings.NewReader(v))
	if err := decoder.Decode(&object); err != nil {
		return nil, err
	}
	return NewFromMap(object)
}

// The NewFromMap() function returns a new Collection from a given map. Because
// the Go native map does not support element ordering, the returned
// Collection will not have the same order as defined in the map. Instead, the
// key ordering will be sorted alphabetically. If the given parameter v is not
// a map, it will returns NonMapTypeError.
//
// Internally, this function will replaces any occurrence of child element with
// map type and replaces it to *Collection. So if the map v contains an element
// --with key named k--- that points to a map, the returned Collection will have
// element with key named k to be pointing to a new Collection, instead of the
// map, thus the impact is that any changes made to the map element will not be
// reflected to the Collection.
//
// For any map element that is not map or map pointer, it will be saved as it is.
func NewFromMap(v interface{}) (*Collection, error) {
	valueOfV := reflect.ValueOf(v)
	if valueOfV.Kind() != reflect.Map {
		return nil, NonMapTypeError
	}
	keys := make([]string, valueOfV.Len())
	for index, key := range valueOfV.MapKeys() {
		keys[index] = key.String()
	}
	sort.Strings(keys)
	collection := New()
	for _, key := range keys {
		pair := Pair{ key: key }
		value := reflect.ValueOf(valueOfV.MapIndex(reflect.ValueOf(key)).Interface())
		switch value.Kind() {
		case reflect.Map:
			convertedValue, err := NewFromMap(value.Interface())
			if err != nil {
				return nil, err
			}
			pair.value = convertedValue
			break
		case reflect.Ptr:
			convertedValue, err := NewFromMap(value.Elem().Interface())
			if err != nil {
				return nil, err
			}
			pair.value = convertedValue
			break
		default:
			pair.value = value.Interface()
		}
		collection.pairs = append(collection.pairs, &pair)
	}
	return collection, nil
}

// Collection defines a Collection Type. See "collection" package documentation
// for more information.
type Collection struct {
	pairs	[]*Pair
}

// The Add() function adds or updates an element with a specified key and value
// to the Collection. Since the Add() function returns back the same Collection,
// you can chain the function call.
//
// If an element with the specified key exists, it will replace the
// key-value pair by deleting current key-value pair and then insert a new
// key-value pair, thus changing the order of insertion. See Set() function
// to update the value and thus will NOT change the order of insertion.
func (collection *Collection) Add(key string, value interface{}) *Collection {
	collection.Delete(key)
	collection.pairs = append(collection.pairs, &Pair{ key, value })
	return collection
}

// The Clear() function removes all elements from the Collection.
func (collection *Collection) Clear() *Collection {
	collection.pairs = make([]*Pair, 0)
	return collection
}

// The Delete() function removes the specified element from the Collection by key.
func (collection *Collection) Delete(key string) *Collection {
	index := collection.IndexOf(key)
	if index != -1 {
		collection.pairs = append(collection.pairs[:index], collection.pairs[index+1:]...)
	}
	return collection
}

// The ForEach() function executes a provided function once for each Collection element.
func (collection *Collection) ForEach(function ForEachFunc) {
	for _, pair := range collection.pairs {
		function(pair.key, pair.value)
	}
}

// The Get() function returns a specified element from the Collection.
// If the element with the given key does not exist, it will returns nil.
func (collection *Collection) Get(key string) interface{} {
	pair := collection.PairOf(key)
	if pair == nil {
		return nil
	}
	return pair.value
}

// The Has() function returns a boolean indicating whether an element with
// the specified key exists or not.
func (collection *Collection) Has(key string) bool {
	return collection.PairOf(key) != nil
}

// The HasAll() function is the same as Has() function, but accepts a slice
// of keys instead of a single key string. If the object has all the element
// for each given key, it will returns true. Otherwise, it will returns
// false.
func (collection *Collection) HasAll(keys ...string) bool {
	for _, key := range keys {
		if !collection.Has(key) {
			return false
		}
	}
	return true
}

// The HasSome() function is the same as Has() function, but accepts a slice
// of keys instead of a single key string. If the object has one or more
// element for each given key, it will returns true. Otherwise, it will
// returns false.
func (collection *Collection) HasSome(keys ...string) bool {
	for _, key := range keys {
		if collection.Has(key) {
			return true
		}
	}
	return false
}

// The IndexOf() function returns the index of Pair{} that represent the
// given key. This Pair{} is registered in the internal slice of Pair{}
// information inside the Collection. If there are no Pair{} registered with
// the given key, it will returns -1.
func (collection *Collection) IndexOf(key string) int {
	for index, pair := range collection.pairs {
		if pair.key == key {
			return index
		}
	}
	return -1
}

// The Keys() function returns a slice of string that contains the keys
// for each element in the Collection, based on the insertion order.
func (collection *Collection) Keys() []string {
	keys := make([]string, len(collection.pairs))
	for index, pair := range collection.pairs {
		keys[index] = pair.key
	}
	return keys
}

// The Length() function returns the number of elements contained inside the
// Collection.
func (collection *Collection) Length() int {
	return len(collection.pairs)
}

// The PairOf() function returns a pointer to the Pair{} that represent the
// given key. This Pair{} is registered in the internal slice of Pair{}
// information inside the Collection. If there are no Pair{} registered with
// the given key, it will returns nil.
func (collection *Collection) PairOf(key string) *Pair {
	for _, pair := range collection.pairs {
		if pair.key == key {
			return pair
		}
	}
	return nil
}

// The Present() function returns an Collection Presenter, which capable to
// returns the the collection to other predefined data type.
func (collection *Collection) Present() Presenter {
	return Presenter{ collection }
}

// The Reflect() function returns the Collection value of the given key as a 
// reflect.Value data. If there is no element saved with the given key, it will
// returns reflect.Value of nil.
func (collection *Collection) Reflect(key string) reflect.Value {
	if !collection.Has(key) {
		return reflect.ValueOf(nil)
	} 
	return reflect.ValueOf(collection.Get(key))
}

// The Reflects() function returns the map[string]reflect.Value representation
// of the Collection.
func (collection *Collection) Reflects() map[string]reflect.Value {
	reflection := make(map[string]reflect.Value)
	for _, pair := range collection.pairs {
		reflection[pair.key] = collection.Reflect(pair.key)
	}
	return reflection
}

// The Set() function adds or updates an element with a specified key and value
// to the Collection. Since the Set() function returns back the same Collection,
// you can chain the function call.
//
// If an element with the specified key exists, it will update the value
// and thus will NOT change the order of insertion. See Add() function
// to replace the key-value pair by deleting current key-value pair and then
// insert a new key-value pair, thus changing the order of insertion.
func (collection *Collection) Set(key string, value interface{}) *Collection {
	pair := collection.PairOf(key)
	if pair != nil {
		pair.key = key
		pair.value = value
	} else {
		collection.pairs = append(collection.pairs, &Pair{ key, value })
	}
	return collection
}

// The String() function returns a string representing the specified Collection
// and its elements.
func (collection *Collection) String() string {
	str := fmt.Sprint()
	for index, pair := range collection.pairs {
		if index != 0 {
			str += fmt.Sprint(",")
		}
		str += fmt.Sprintf("\"%s\":", pair.key)
		switch reflect.ValueOf(pair.value).Kind() {
		case reflect.String:
			str += fmt.Sprintf("\"%s\"", pair.value.(string))
			break
		default:
			str += fmt.Sprintf("%v", pair.value)
		}
	}
	return fmt.Sprintf("{%s}", str)
}

// The Values() function returns a Values that contains the values for each element
// in the Collection in insertion order.
func (collection *Collection) Values() []interface{} {
	values := make([]interface{}, len(collection.pairs))
	for index, pair := range collection.pairs {
		values[index] = pair.value
	}
	return values
}

// Pair defines key-value pair of an element in Collection.
type Pair struct {
	key 	string
	value 	interface{}
}

type ForEachFunc func (key string, value interface{})