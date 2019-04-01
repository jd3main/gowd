// Full list of data type: <https://www.wikidata.org/wiki/Special:ListDatatypes>

package gowd

import (
)

// DataType := commonsMedia | globe-coordinate |
//              wikibase-item | wikibase-property | 
//              string | monolingualtext |
//              external-id | quantity | time |
//              url | math | geo-shape | musical-notation | tabular-data |
//              wikibase-lexeme | wikibase-form | wikibase-sense

var DataTypeSupports map[string]bool

func init() {
    DataTypeSupports = map[string]bool {
        "commonsMedia":         true,
        "globe-coordinate":     true,
        "wikibase-item":        true,
        "wikibase-property":    true,
        "string":               true,
        "monolingualtext":      true,
        "external-id":          true,
        "quantity":             true,
        "time":                 true,
        "url":                  true,
        "math":                 true,
        "geo-shape":            true,
        "musical-notation":     true,
        "tabular-data":         true,
        "wikibase-lexeme":      true,
        "wikibase-form":        true,
        "wikibase-sense":       true,
    }
}

func IsDataType(s string) bool {
    _, exist :=  DataTypeSupports[s]
    return exist
}
