package payload

import (
	"github.com/perimeterx/marshmallow"
	"net/url"
)

// Navigation defines a common structure of a navigation layout
type Navigation struct {
	// Header navigational links that are typically required
	// within Payload settings.
	Header NavigationItem `json:"header,omitempty"`

	// Footer navigational links that are optionally required.
	Footer NavigationItems `json:"footer,omitempty"`

	// Arbitrary key-value pairs of any other fields that appear within
	// the schema but are not defined in the struct.
	Fields map[string]any `json:"-"`
}

// UnmarshalJSON unmarshalls the JSON data into the Navigation type.
// This method is used to extract known fields and assign the remaining
// fields to the Fields map.
func (n *Navigation) UnmarshalJSON(data []byte) error {
	var temp Navigation
	result, err := marshmallow.Unmarshal(data,
		&temp,
		marshmallow.WithExcludeKnownFieldsFromMap(true),
	)
	if err != nil {
		return err
	}

	*n = temp
	n.Fields = result

	return nil
}

// NavigationItems is a collection of NavigationItem types.
type NavigationItems []NavigationItem

// NavigationItem defines a common structure of a singular navigation item
// from PayloadCMS.
type NavigationItem struct {
	// The ID of the item, this is generated by Payload and is used to
	// uniquely identify the li k.
	ID string `json:"id,omitempty"`

	// Title or label of the navigation item.
	Title string `json:"title"`

	// The URL that the navigation item should link to.
	URL url.URL `json:"url"`

	// AN optional image media object associated with the link.
	Image Media `json:"image,omitempty"`

	// Optional children items of the navigation item. Maximum depth is
	// defined within the Payload settings.
	Children NavigationItems `json:"children"`

	// Arbitrary key-value pairs of any other fields that appear within
	// the schema but are not defined in the struct.
	Fields map[string]any `json:"-"`
}

// UnmarshalJSON unmarshalls the JSON data into the NavigationItem type.
// This method is used to extract known fields and assign the remaining
// fields to the Fields map.
func (i *NavigationItem) UnmarshalJSON(data []byte) error {
	var temp NavigationItem
	result, err := marshmallow.Unmarshal(data,
		&temp,
		marshmallow.WithExcludeKnownFieldsFromMap(true),
	)
	if err != nil {
		return err
	}

	*i = temp
	i.Fields = result

	return nil
}
