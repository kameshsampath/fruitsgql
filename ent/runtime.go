// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/kameshsampath/fruitsgql/ent/fruit"
	"github.com/kameshsampath/fruitsgql/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	fruitFields := schema.Fruit{}.Fields()
	_ = fruitFields
	// fruitDescName is the schema descriptor for name field.
	fruitDescName := fruitFields[0].Descriptor()
	// fruit.NameValidator is a validator for the "name" field. It is called by the builders before save.
	fruit.NameValidator = fruitDescName.Validators[0].(func(string) error)
	// fruitDescSeason is the schema descriptor for season field.
	fruitDescSeason := fruitFields[1].Descriptor()
	// fruit.SeasonValidator is a validator for the "season" field. It is called by the builders before save.
	fruit.SeasonValidator = fruitDescSeason.Validators[0].(func(string) error)
}