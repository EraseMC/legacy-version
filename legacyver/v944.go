package legacyver

import (
	_ "embed"

	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/akmalfairuz/legacy-version/mapping"
)

const (
	// ItemVersion944 ...
	ItemVersion944 = 251
	// BlockVersion944 ...
	BlockVersion944 int32 = (1 << 24) | (26 << 16) | (0 << 8)
)

var (
	//go:embed data/required_item_list_944.json
	requiredItemList944 []byte
	//go:embed data/block_states_944.nbt
	blockStateData944 []byte
)

func New944(dragonflyMapping bool) *Protocol {
	itemMapping := mapping.NewItemMapping(requiredItemList944, ItemVersion944)
	blockTranslator := lookupOrCreateBlockTranslator(944, BlockVersion944, blockStateData944)
	return &Protocol{
		ver:             "1.26.10",
		id:              proto.ID944,
		blockTranslator: blockTranslator,
		itemTranslator:  NewItemTranslator(itemMapping, itemMappingLatest(dragonflyMapping), blockTranslator.BlockMapping(), blockMappingLatest),
	}
}
