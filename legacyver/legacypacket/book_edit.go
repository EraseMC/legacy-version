package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

const (
	BookActionReplacePage = iota
	BookActionAddPage
	BookActionDeletePage
	BookActionSwapPages
	BookActionSign
)

// BookEdit is sent by the client when it edits a book. It is sent each time a modification was made and the
// player stops its typing 'session', rather than simply after closing the book.
type BookEdit struct {
	// InventorySlot is the slot in which the book that was edited may be found. Typically, the server should
	// check if this slot matches the held item slot of the player.
	InventorySlot int32
	// ActionType is the type of the book edit action. The data obtained depends on what type this is. The
	// action type is one of the constants above.
	ActionType uint32
	// PageNumber is the number of the page that the book edit action concerns. It applies for all actions
	// but the BookActionSign. In BookActionSwapPages, it is one of the pages that was swapped.
	PageNumber int32
	// SecondaryPageNumber is the page number of the second page that the action concerned. It is only set for
	// the BookActionSwapPages action, in which case it is the other page that is swapped.
	SecondaryPageNumber int32
	// Text is the text that was written in a particular page of the book. It applies for the
	// BookActionAddPage and BookActionReplacePage only.
	Text string
	// PhotoName is the name of the photo on the page in the book. It applies for the BookActionAddPage and
	// BookActionReplacePage only.
	// Unfortunately, the functionality of this field was removed from the default Minecraft Bedrock Edition.
	// It is still available on Education Edition.
	PhotoName string
	// Title is the title that the player has given the book. It applies only for the BookActionSign action.
	Title string
	// Author is the author that the player has given the book. It applies only for the BookActionSign action.
	// Note that the author may be freely changed, so no assumptions can be made on if the author is actually
	// the name of a player.
	Author string
	// XUID is the XBOX Live User ID of the player that edited the book. The field is rather pointless, as the
	// server is already aware of the XUID of the player anyway.
	XUID string
}

// ID ...
func (*BookEdit) ID() uint32 {
	return packet.IDBookEdit
}

func (pk *BookEdit) Marshal(io protocol.IO) {
	if proto.IsProtoLT(io, proto.ID924) {
		x := byte(pk.ActionType)
		io.Uint8(&x)
		pk.ActionType = uint32(x)
	} else {
		io.Varint32(&pk.InventorySlot)
	}
	if proto.IsProtoLT(io, proto.ID924) {
		x := byte(pk.InventorySlot)
		io.Uint8(&x)
		pk.InventorySlot = int32(x)
	} else {
		io.Varuint32(&pk.ActionType)
	}
	switch pk.ActionType {
	case BookActionReplacePage, BookActionAddPage:
		if proto.IsProtoLT(io, proto.ID924) {
			x := byte(pk.PageNumber)
			io.Uint8(&x)
			pk.PageNumber = int32(x)
		} else {
			io.Varint32(&pk.PageNumber)
		}
		io.String(&pk.Text)
		io.String(&pk.PhotoName)
	case BookActionDeletePage:
		if proto.IsProtoLT(io, proto.ID924) {
			x := byte(pk.PageNumber)
			io.Uint8(&x)
			pk.PageNumber = int32(x)
		} else {
			io.Varint32(&pk.PageNumber)
		}
	case BookActionSwapPages:
		if proto.IsProtoLT(io, proto.ID924) {
			x := byte(pk.PageNumber)
			io.Uint8(&x)
			pk.PageNumber = int32(x)

			y := byte(pk.SecondaryPageNumber)
			io.Uint8(&y)
			pk.SecondaryPageNumber = int32(y)
		} else {
			io.Varint32(&pk.PageNumber)
			io.Varint32(&pk.SecondaryPageNumber)
		}
	case BookActionSign:
		io.String(&pk.Title)
		io.String(&pk.Author)
		io.String(&pk.XUID)
	default:
		io.UnknownEnumOption(pk.ActionType, "book edit action type")
	}
}
