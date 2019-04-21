package gowd

import (
    "encoding/json"
)


type entityType string
const (
    EntityType_Item entityType = "item"
    EntityType_Property entityType = "property"
    EntityType_Lexeme entityType = "lexeme"
)


type Entity struct{
    Id              string      `json:"id"`          // Q111 for example
    Type            entityType  `json:"type"`        // "item"|"property"|"lexeme"
    
    Labels          map[string]MonolingualText      `json:"labels"`
    Descriptions    map[string]MonolingualText      `json:"descriptions"`
    Aliases         map[string][]MonolingualText    `json:"aliases"`
    Claims          map[string][]Statement          `json:"claims"`
    Sitelinks       map[string]SiteLink             `json:"sitelinks"`
    
    // Revision infos
    Lastrevid       Integer `json:"lastrevid"`   // The document's version (MediaWiki revision ID)
    Modified        string  `json:"modified"`    // The document's Publication date (MediaWiki revision timestamp)

    // Extra informations
    Title 	        string  `json:"title"`      // The title of the page the entity is stored on
    PageId 	        Integer `json:"pageid"`     // The page id the entity is stored on
    Ns              Integer `json:"ns"`         // The namespace id the entity is stored on
}


type Statement struct {
    MainSnak    Snak                    `json:"mainsnak"`
    Qualifiers  map[string][]Qualifier  `json:"qualifiers"`
    Id          string                  `json:"id"`             // The id of this statement
    Rank        string                  `json:"rank"`
    References  []Reference             `json:"references"`
}

type Snak struct {
    SnakType    string      `json:"snaktype"`    // value|somevalue|novalue
    Property    string      `json:"property"`    // The ID of the property this Snak is about
    DataType    string      `json:"datatype"`    // Data types specified by Wikibase data model
    
    // The actual value associated with the Snak (if the snaktype is "value")
    DataValueInterpreter    `json:"datavalue"`
}

func (snak *Snak) UnmarshalJSON(b []byte) error {
    var tempSnak struct {
        SnakType string `json:"snaktype"`
        Property string `json:"property"`
        DataType string `json:"datatype"`
        DataValue json.RawMessage `json:datavalue`
    }

    err := json.Unmarshal(b, &tempSnak)
    if err != nil {
        return err
    }

    snak.SnakType = tempSnak.SnakType
    snak.Property = tempSnak.Property
    snak.DataType = tempSnak.DataType
    err = json.Unmarshal(tempSnak.DataValue, &snak.DataValueInterpreter)
    return err
}

type Qualifier struct{
    Hash    string      `json:"hash"`
    Snak
}

type Reference struct {
    Hash    string              `json:"hash"`
    Snaks   map[string][]Snak   `json:"snaks"`
}

type SiteLink struct {
    Site    string      `json:"site"`
    Title   string      `json:"title"`
    Badges  []string    `json:"badges"`
    Url     string      `json:"url"`
}