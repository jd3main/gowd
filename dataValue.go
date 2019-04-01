// Described in <https://www.mediawiki.org/wiki/Wikibase/DataModel/JSON#Data_Values>

package gowd

import (
    "encoding/json"
)


var DataValueTypeSupports map[string]bool

func init() {
    DataValueTypeSupports = map[string]bool{
        "string":               true,
        "wikibase-entityid":    true,
        "globecoordinate":      false,
        "quantity":        		false,
        "time":                 false,
        "monolingualtext":    	false,
    }
}

func IsDataValueType(s string) bool {
    _, exist :=  DataValueTypeSupports[s]
    return exist
}

func IsSupportedDataValueType(s string) bool {
    return DataValueTypeSupports[s]
}


type dataValue interface{}



type WikibaseEntityId struct {
    Value       string `json:"id"`
    EntityType  string `json:"entity-type"`
    NumericId   string `json:"numeric-id"`
}


type GlobeCoordinate struct {
    json.RawMessage
}


type Quantity struct {
    json.RawMessage
}


type Time struct {
    json.RawMessage
}


type MonolingualText struct {
    json.RawMessage
}
