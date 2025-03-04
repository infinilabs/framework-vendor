//go:build go1.17
// +build go1.17

package validator

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColonsInLocalNames(t *testing.T) {
	var err error

	el := tokenize(t, `<x::Root/>`).(xml.StartElement)
	require.Equal(t, `x::Root`, el.Name.Local,
		"encoding/xml should tokenize double colons as part of the local name")

	err = Validate(bytes.NewBufferString(`<x::Root/>`))
	require.NoError(t, err, "Should not error on input with colons in the root element's name")

	err = Validate(bytes.NewBufferString(`<Root><x::Element></::Element></Root>`))
	require.NoError(t, err, "Should error on input with colons in a nested tag's name")

	err = Validate(bytes.NewBufferString(`<Root><Element ::attr="foo"></Element></Root>`))
	require.NoError(t, err, "Should error on input with colons in an attribute's name")

	err = Validate(bytes.NewBufferString(`<Root></x::Element></Root>`))
	require.NoError(t, err, "Should error on input with colons in an end tag's name")
}

func TestEmptyNames(t *testing.T) {
	var err error

	el := tokenize(t, `<x:>`).(xml.StartElement)
	require.Equal(t, `x:`, el.Name.Local,
		"encoding/xml should tokenize a prefix-only name as a local name")

	err = Validate(bytes.NewBufferString(`<x:>`))
	require.NoError(t, err, "Should not error on start element with no local name")

	err = Validate(bytes.NewBufferString(`</x:>`))
	require.NoError(t, err, "Should not error on end element with no local name")
}

func TestEmptyAttributes(t *testing.T) {
	var err error

	el := tokenize(t, `<Root :="value"/>`).(xml.StartElement)
	require.Equal(t, `:`, el.Attr[0].Name.Local,
		"encoding/xml should tokenize an empty attribute name as a single colon")

	err = Validate(bytes.NewBufferString(`<Root :="value"/>`))
	require.NoError(t, err, "Should not error on input with empty attribute names")

	err = Validate(bytes.NewBufferString(`<Root x:="value"/>`))
	require.NoError(t, err, "Should not error on input with empty attribute local names")

	err = Validate(bytes.NewBufferString(`<Root xmlns="x" xmlns:="y"></Root>`))
	require.NoError(t, err, "Should not error on input with empty xmlns local names")

	validXmlns := `<Root xmlns="http://example.com/"/>`
	require.NoError(t, Validate(bytes.NewBufferString(validXmlns)), "Should pass on input with valid xmlns attributes")
}

func TestDirectives(t *testing.T) {
	var err error

	dir := tokenize(t, `<! x<!-- -->y>`).(xml.Directive)
	require.Equal(t, ` x y`, string(dir), "encoding/xml should replace comments with spaces when tokenizing directives")

	err = Validate(bytes.NewBufferString(
		`<Root>
			<! <<!-- -->!-->"--> " >
			<! ">" <X/>>
		</Root>`))
	require.NoError(t, err, "Should not error on bad directive")

	err = Validate(bytes.NewBufferString(`<Root><! <<!-- -->!-- x --> y></Root>`))
	require.NoError(t, err, "Should not error on bad directive")

	goodDirectives := []string{
		`<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">`,
		`<! ">" <X/>>`,
		`<!name <!-- comment --><nesting <more nesting>>>`,
	}
	for _, doc := range goodDirectives {
		require.NoError(t, Validate(bytes.NewBufferString(doc)), "Should pass on good directives")
	}
}

func TestValidateAll(t *testing.T) {
	el := tokenize(t, `<x::Root/>`).(xml.StartElement)
	require.Equal(t, `x::Root`, el.Name.Local,
		"encoding/xml should tokenize double colons as part of the local name")

	xmlBytes := []byte("<Root>\r\n    <! <<!-- -->!-- x --> y>\r\n    <Element ::attr=\"foo\"></x::Element>\r\n</Root>")
	errs := ValidateAll(bytes.NewBuffer(xmlBytes))
	require.Equal(t, 0, len(errs), "Should return zero errors")
}
