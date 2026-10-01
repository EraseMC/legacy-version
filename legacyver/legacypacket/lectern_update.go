package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

// LecternUpdate is sent by the client to update the server on which page was opened in a book on a lectern,
// or if the book should be removed from it.
type LecternUpdate struct {
	// Page is the page number in the book that was opened by the player on the lectern.
	Page byte
	// PageCount is the number of pages that the book opened in the lectern has.
	PageCount byte
	// Position is the position of the lectern that was updated. If no lectern is at the block position,
	// the packet should be ignored.
	Position protocol.BlockPos
}

// ID ...
func (*LecternUpdate) ID() uint32 {
	return packet.IDLecternUpdate
}

func (pk *LecternUpdate) Marshal(io protocol.IO) {
	io.Uint8(&pk.Page)
	io.Uint8(&pk.PageCount)
	if proto.IsProtoGTE(io, proto.ID944) {
		io.BlockPos(&pk.Position)
	} else {
		proto.UBlockPos(io, &pk.Position)
	}
}

// ToLatest ...
func (pk *LecternUpdate) ToLatest() *packet.LecternUpdate {
	return &packet.LecternUpdate{
		Page:      pk.Page,
		PageCount: pk.PageCount,
		Position:  pk.Position,
	}
}

// FromLatest ...
func (pk *LecternUpdate) FromLatest(latest *packet.LecternUpdate) *LecternUpdate {
	pk.Page = latest.Page
	pk.PageCount = latest.PageCount
	pk.Position = latest.Position
	return pk
}
