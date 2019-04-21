// Basic datatypes of wikibase
// They are: String, integer, nonNegativeInteger, decimal, IRI, GlobalSiteIdentifier, UserLanguageCode
// Described here: https://www.mediawiki.org/wiki/Wikibase/DataModel#Defining_data_structures_in_UML
// and there exported name are captialized, for example decimal is exported as Decimal

package gowd

import (
    "math/big"
)


// basic data type "string" is simply same as string in go

// Interger corresponds to the basic data type "integer"
type Integer struct {
    big.Int
}

func NewInteger(i int64) (x *Integer) {
    x = new(Integer)
    x.Int.SetInt64(i)
    return
}

func (x *Integer) SetString(s string) bool {
    i := big.NewInt(0)
    _, ok := i.SetString(s,10)
    return ok
}


// Unsigned corresponds to the basic data type "nonNegativeInteger"
type Unsigned Integer

func NewUnsigned(i uint64) (x *Unsigned) {
    x = new(Unsigned)
    x.SetUint64(i)
    return
}

func (x *Unsigned) SetUint64(i uint64) {
    x.Int.SetUint64(i)
}

func (x *Unsigned) Set(i *big.Int) bool {
    if i.Sign()<0 {
        return false
    }
    x.Int.Set(i)
    return true
}

func (x *Unsigned) SetString(s string) bool {
    i := big.NewInt(0)
    _, ok := i.SetString(s,10)
    if ok && i.Sign()>=0 {
        x.Int.Set(i)
        return true
    }else {
        return false
    }
}


// Decimal corresponds to the basic data type "decimal"
type Decimal struct {
    big.Float
}

func NewDecimal(f float64) (x *Decimal) {
    x = new(Decimal)
    x.Float.SetFloat64(f)
    return
}


// TODO: enable format check
type IRI string
type GlobalSiteIdentifier string
type UserLanguageCode string


