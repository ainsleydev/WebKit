package payloadcms

import "time"

// Meta defines the data generated by the SEO plugin from Payload
// along with additional fields such as Private & Canonical.
//
// The SEO plugin appears in the majority of collections and in both
// the Global Settings and Page level fields.
type Meta struct {
	Title          *string `json:"title,omitempty"`
	Description    *string `json:"description,omitempty"`
	Image          *int    `json:"image,omitempty"`
	Private        *bool   `json:"private,omitempty"`
	CanonicalURL   *string `json:"canonicalURL,omitempty"`
	StructuredData any     `json:"structuredData,omitempty"`
}

// Settings defines the common global collection type within Payload
// that allows users to change site settings.
type Settings struct {
	Id            float64        `json:"id"`
	SiteName      *string        `json:"siteName,omitempty"`
	TagLine       *string        `json:"tagLine,omitempty"`
	Locale        string         `json:"locale,omitempty"` // In en_GB format
	Logo          *int           `json:"logo,omitempty"`
	Meta          Meta           `json:"meta"`
	Robots        *string        `json:"robots,omitempty"`
	CodeInjection *CodeInjection `json:"codeInjection,omitempty"`
	Maintenance   *Maintenance   `json:"maintenance,omitempty"`
	Contact       *Contact       `json:"contact,omitempty"`
	UpdatedAt     *time.Time     `json:"updatedAt,omitempty"`
	CreatedAt     *time.Time     `json:"createdAt,omitempty"`
}

// CodeInjection defines the fields for injecting code into the head
// or foot of the frontend.
type CodeInjection struct {
	Footer *string `json:"footer,omitempty"`
	Head   *string `json:"head,omitempty"`
}

// Maintenance defines the fields for displaying an offline page to
// the front-end when it's been enabled within PayloadCMS.
type Maintenance struct {
	Content *string `json:"content,omitempty"`
	Enabled *bool   `json:"enabled,omitempty"`
	Title   *string `json:"title,omitempty"`
}

// Contact defines the fields for Contact Details within PayloadCMS.
type Contact struct {
	Address   *string `json:"address,omitempty"`
	Email     *string `json:"email,omitempty"`
	Facebook  *string `json:"facebook,omitempty"`
	Instagram *string `json:"instagram,omitempty"`
	LinkedIn  *string `json:"linkedIn,omitempty"`
	Telephone *string `json:"telephone,omitempty"`
	Tiktok    *string `json:"tiktok,omitempty"`
	X         *string `json:"x,omitempty"`
	Youtube   *string `json:"youtube,omitempty"`
}
