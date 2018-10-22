package zubrin

import (
	"fmt"
	"reflect"

	"github.com/fatih/structtag"
)

// EnvTagPrefixer is a tag modifier which adds a prefix to the values of
// `env` tags. This can be used to register one config multiple times and
// have their initialization be read from different environment variables.
type EnvTagPrefixer struct {
	prefix string
}

// NewEnvTagPrefixer creates a new EnvTagPrefixer.
func NewEnvTagPrefixer(prefix string) TagModifier {
	return &EnvTagPrefixer{
		prefix: prefix,
	}
}

// AlterFieldTag adds the env prefixer's prefix to the `env` tag value, if one is set.
func (p *EnvTagPrefixer) AlterFieldTag(fieldType reflect.StructField, tags *structtag.Tags) error {
	tag, err := tags.Get("env")
	if err != nil {
		return nil
	}

	return tags.Set(&structtag.Tag{
		Key:  "env",
		Name: fmt.Sprintf("%s_%s", p.prefix, tag.Name),
	})
}
