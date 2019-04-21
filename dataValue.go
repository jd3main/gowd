// Described in <https://www.mediawiki.org/wiki/Wikibase/DataModel/JSON#Data_Values>

package gowd

import (
    "encoding/json"
    "strings"
    "errors"
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

func IsImplementedDataValueType(s string) bool {
    return DataValueTypeSupports[s]
}


type DataValue interface{
    _isDataValue()
    Type() string
}


func (*String)              _isDataValue() {}
func (*WikibaseEntityId)    _isDataValue() {}
func (*GlobeCoordinate)     _isDataValue() {}
func (*Quantity)            _isDataValue() {}
func (*Time)                _isDataValue() {}
func (*MonolingualText)     _isDataValue() {}

func (*String)              Type() string { return "string" }
func (*WikibaseEntityId)    Type() string { return "wikibase-entityid" }
func (*GlobeCoordinate)     Type() string { return "globecoordinate" }
func (*Quantity)            Type() string { return "quantity" }
func (*Time)                Type() string { return "time" }
func (*MonolingualText)     Type() string { return "monolingualtext" }


type String string

type WikibaseEntityId struct {
    Value       string  `json:"id"`
    EntityType  string  `json:"entity-type"`
    NumericId   Integer `json:"numeric-id"`
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
    Language    string  `json:"language"`
    Value       string  `json:"value"`
}



type DataValueInterpreter struct {
    DataValue   `json:"value"`
}

func (data *DataValueInterpreter) UnmarshalJSON(b []byte) error {
    var tempData struct {
        Type    string          `json:"type"`
        Value   json.RawMessage `json:"value"`
    }

	if err := json.Unmarshal(b, &tempData); err != nil {
		return err
    }

	switch strings.ToLower(tempData.Type) {
	case "string":
        data.DataValue = new(String)
	case "wikibase-entityid":
        data.DataValue = new(WikibaseEntityId)
	case "globecoordinate":
        data.DataValue = new(GlobeCoordinate)
	case "quantity":
        data.DataValue = new(Quantity)
	case "time":
        data.DataValue = new(Time)
	case "monolingualtext":
        data.DataValue = new(MonolingualText)
	default:
        return errors.New(`Unsupported value type: "`+tempData.Type+`"`)
    }

    err := json.Unmarshal(tempData.Value, data.DataValue)
    
	return err
}
