package payload

import (
	"github.com/goccy/go-json"
)

// Block defines a common structure of a singular block layout
// from PayloadCMS.
//
// See: https://payloadcms.com/docs/fields/blocks
type Block struct {
	// The ID of the block, this is generated by Payload and is used to
	// uniquely identify the block.
	Id string `json:"id,omitempty"`

	// The lockType is saved as the slug of the block that has been selected.
	BlockType any `json:"block_type"`

	// The Admin panel provides each block with a blockName field which optionally
	// allows editors to label their blocks for better readability.
	BlockName *string `json:"block_name,omitempty"`

	// Key-value pairs of the block's fields, these pairs are defined by the block's
	// schema and vary depending on the block type.
	Fields map[string]any

	// RawJSON is the raw byte slice of the block, which can be used to decode
	// the block into a specific type.
	RawJSON json.RawMessage
}

// Blocks is a collection of Block types.
type Blocks []Block

// Decode decodes the block's raw JSON into the provided interface.
// This is optional, to be more performant, you make just access
// the fields directly.
func (b *Block) Decode(v any) error {
	return json.Unmarshal(b.RawJSON, v)
}

// UnmarshalJSON unmarshals the JSON data into the Block struct.
// This method is used to extract known fields and assign the remaining
// fields to the fields map.
func (b *Block) UnmarshalJSON(data []byte) error {
	b.RawJSON = data

	// Create a map to hold the JSON object
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}

	// Extract known fields and remove them from the map
	if blockType, ok := m["blockType"].(string); ok {
		b.BlockType = blockType
	}

	if blockName, ok := m["blockName"].(string); ok {
		b.BlockName = &blockName
	}

	if id, ok := m["id"].(string); ok {
		b.Id = id
	}

	// Remove the known fields from the map
	delete(m, "blockType")
	delete(m, "blockName")
	delete(m, "id")

	// Assign the remaining fields to the Other map
	b.Fields = m

	return nil
}
