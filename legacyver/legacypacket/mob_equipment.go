package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

// MobEquipment is sent by the client to the server and the server to the client to make the other side
// aware of the new item that an entity is holding. It is used to show the item in the hand of entities such
// as zombies too.
type MobEquipment struct {
	// EntityRuntimeID is the runtime ID of the entity. The runtime ID is unique for each world session, and
	// entities are generally identified in packets using this runtime ID.
	EntityRuntimeID uint64
	// NewItem is the new item held after sending the MobEquipment packet. The entity will be shown holding
	// that item to the player it was sent to.
	NewItem protocol.ItemInstance
	// InventorySlot is the slot in the inventory that was held. This is the same as HotBarSlot, and only
	// remains for backwards compatibility.
	InventorySlot byte
	// HotBarSlot is the slot in the hot bar that was held. It is the same as InventorySlot, which is only
	// there for backwards compatibility purposes.
	HotBarSlot byte
	// WindowID is the window ID of the window that had its equipped item changed. This is usually the window
	// ID of the normal inventory, but may also be something else, for example with the off hand.
	WindowID byte
}

// ID ...
func (*MobEquipment) ID() uint32 {
	return packet.IDMobEquipment
}

func (pk *MobEquipment) Marshal(io protocol.IO) {
	io.Varuint64(&pk.EntityRuntimeID)
	// The held item became a NetworkItemStackDescriptor (ItemInstanceNew) in
	// 1.26.20 / proto 975 — see the PMMP commit that renamed the call to
	// getNetworkItemStackDescriptor. gophertunnel v1.54.0 (proto 924 = 1.26.0)
	// still uses the legacy ItemInstance layout, so anything below 975 must
	// continue to use it. The tedac fork's "1.16.220+" comment is misleading
	// — only the latest two protocol bumps actually expect the descriptor.
	if proto.IsProtoGTE(io, proto.ID975) {
		io.ItemInstanceNew(&pk.NewItem)
	} else {
		io.ItemInstance(&pk.NewItem)
	}
	io.Uint8(&pk.InventorySlot)
	io.Uint8(&pk.HotBarSlot)
	io.Uint8(&pk.WindowID)
}

// ToLatest ...
func (pk *MobEquipment) ToLatest() *packet.MobEquipment {
	return &packet.MobEquipment{
		EntityRuntimeID: pk.EntityRuntimeID,
		NewItem:         pk.NewItem,
		InventorySlot:   pk.InventorySlot,
		HotBarSlot:      pk.HotBarSlot,
		WindowID:        pk.WindowID,
	}
}

// FromLatest ...
func (pk *MobEquipment) FromLatest(latest *packet.MobEquipment) *MobEquipment {
	pk.EntityRuntimeID = latest.EntityRuntimeID
	pk.NewItem = latest.NewItem
	pk.InventorySlot = latest.InventorySlot
	pk.HotBarSlot = latest.HotBarSlot
	pk.WindowID = latest.WindowID
	return pk
}
