package legacyver

import (
	_ "embed"

	"github.com/akmalfairuz/legacy-version/mapping"
)

const (
	// ItemVersion975 ...
	ItemVersion975 = 261
	// BlockVersion975 ...
	BlockVersion975 int32 = (1 << 24) | (26 << 16) | (10 << 8)
)

var (
	//go:embed data/dragonfly_items.json
	dragonflyLatestItemList []byte
	//go:embed data/required_item_list_975.json
	requiredItemList975 []byte
	//go:embed data/block_states_975.nbt
	blockStateData975 []byte

	itemMappingLatestPocketMine = mapping.NewItemMapping(requiredItemList975, ItemVersion975)
	itemMappingLatestDragonfly  = mapping.NewItemMapping(dragonflyLatestItemList, ItemVersion975)
	blockMappingLatest          = mapping.NewBlockMapping(blockStateData975)
)

func itemMappingLatest(dragonflyMapping bool) mapping.Item {
	if dragonflyMapping {
		return itemMappingLatestDragonfly
	}
	return itemMappingLatestPocketMine
}
