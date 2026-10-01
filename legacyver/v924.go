package legacyver

import (
	_ "embed"

	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/akmalfairuz/legacy-version/mapping"
)

const (
	// ItemVersion924 ...
	ItemVersion924 = 251
	// BlockVersion924 ...
	BlockVersion924 int32 = (1 << 24) | (26 << 16) | (0 << 8)
)

var (
	//go:embed data/required_item_list_924.json
	requiredItemList924 []byte
	//go:embed data/block_states_924.nbt
	blockStateData924 []byte
)

func New924(dragonflyMapping bool) *Protocol {
	itemMapping := mapping.NewItemMapping(requiredItemList924, ItemVersion924)
	blockTranslator := lookupOrCreateBlockTranslator(924, BlockVersion924, blockStateData924)
	return &Protocol{
		ver:             "1.26.0",
		id:              proto.ID924,
		blockTranslator: blockTranslator,
		itemTranslator:  NewItemTranslator(itemMapping, itemMappingLatest(dragonflyMapping), blockTranslator.BlockMapping(), blockMappingLatest),
	}
}
