package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

// InventorySlot is sent by the server to update a single slot in one of the inventory windows that the client
// currently has opened. Usually this is the main inventory, but it may also be the off hand or, for example,
// a chest inventory.
type InventorySlot struct {
	// WindowID is the ID of the window that the packet modifies. It must point to one of the windows that the
	// client currently has opened.
	WindowID uint32
	// Slot is the index of the slot that the packet modifies. The new item will be set to the slot at this
	// index.
	Slot uint32
	// Container is the protocol.FullContainerName that describes the container that the content is for.
	Container protocol.Optional[proto.FullContainerName]
	// DynamicContainerSize ...
	DynamicContainerSize uint32
	// StorageItem is the item that is acting as the storage container for the inventory. If the inventory is
	// not a dynamic container then this field should be left empty. When set, only the item type is used by
	// the client and none of the other stack info.
	StorageItem protocol.Optional[protocol.ItemInstance]
	// NewItem is the item to be put in the slot at Slot. It will overwrite any item that may currently
	// be present in that slot.
	NewItem protocol.ItemInstance
}

// ID ...
func (*InventorySlot) ID() uint32 {
	return packet.IDInventorySlot
}

func (pk *InventorySlot) Marshal(io protocol.IO) {
	io.Varuint32(&pk.WindowID)
	io.Varuint32(&pk.Slot)
	if proto.IsProtoGTE(io, proto.ID729) {
		if proto.IsProtoGTE(io, proto.ID975) {
			protocol.OptionalMarshaler(io, &pk.Container)
		} else {
			var container proto.FullContainerName
			if v, ok := pk.Container.Value(); ok {
				container = v
			}
			protocol.Single(io, &container)
			pk.Container = protocol.Option(container)
		}
	}
	if proto.IsProtoGTE(io, proto.ID748) {
		if proto.IsProtoGTE(io, proto.ID975) {
			protocol.OptionalFunc(io, &pk.StorageItem, io.ItemInstanceNew)
		} else {
			var storageItem protocol.ItemInstance
			if v, ok := pk.StorageItem.Value(); ok {
				storageItem = v
			}
			io.ItemInstance(&storageItem)
			pk.StorageItem = protocol.Option(storageItem)
		}
	} else {
		if proto.IsProtoGTE(io, proto.ID712) {
			io.Varuint32(&pk.DynamicContainerSize)
		}
	}
	if proto.IsProtoGTE(io, proto.ID975) {
		io.ItemInstanceNew(&pk.NewItem)
	} else {
		io.ItemInstance(&pk.NewItem)
	}
}

// ToLatest ...
func (pk *InventorySlot) ToLatest() *packet.InventorySlot {
	var container protocol.Optional[protocol.FullContainerName]
	if v, ok := pk.Container.Value(); ok {
		container = protocol.Option(v.ToLatest())
	}
	return &packet.InventorySlot{
		WindowID:    pk.WindowID,
		Slot:        pk.Slot,
		Container:   container,
		StorageItem: pk.StorageItem,
		NewItem:     pk.NewItem,
	}
}

// FromLatest ...
func (pk *InventorySlot) FromLatest(latest *packet.InventorySlot) *InventorySlot {
	var container protocol.Optional[proto.FullContainerName]
	if v, ok := latest.Container.Value(); ok {
		container = protocol.Option((&proto.FullContainerName{}).FromLatest(v))
	}
	pk.WindowID = latest.WindowID
	pk.Slot = latest.Slot
	pk.Container = container
	pk.StorageItem = latest.StorageItem
	pk.NewItem = latest.NewItem
	return pk
}
