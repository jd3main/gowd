package gowd

import (
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
    Lastrevid       string  `json:"lastrevid"`   // The document's version (MediaWiki revision ID)
    Modified        string  `json:"modified"`    // The document's Publication date (MediaWiki revision timestamp)

    // Extra informations
    Title 	        string  `json:"title"`      // The title of the page the entity is stored on
    Pageid 	        string  `json:"pageid"`     // The page id the entity is stored on
    Ns              string  `json:"ns"`         // The namespace id the entity is stored on
}


type Statement struct {
    Id          string  `json:"id"`             // The id of this statement.
    MainSnak    Snak    `json:"mainsnak"`
    Rank        string  `json:"rank"`
    Qualifiers  map[string][]Qualifier  `json:"qualifiers"`
    References  []Reference `json:"references"`
}

type Snak struct {
    SnakType    string      `json:"snaktype"`    // value|somevalue|novalue
    Property    string      `json:"property"`    // The ID of the property this Snak is about
    DataType    string      `json:"datatype"`    // Datatype indicates how the value of the Snak can be interpreted
    DataValue   dataValue   `json:"datavalue"`   // the actual value the Snak associates with (if the snaktype is "value")
}

type Qualifier struct{
    Hash    string  `json:"hash"`
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