package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
	"image/color"
)

const (
	MapUpdateFlagTexture = 1 << (iota + 1)
	MapUpdateFlagDecoration
	MapUpdateFlagInitialisation
)

// ClientBoundMapItemData is sent by the server to the client to update the data of a map shown to the client.
// It is sent with a combination of flags that specify what data is updated.
// The ClientBoundMapItemData packet may be used to update specific parts of the map only. It is not required
// to send the entire map each time when updating one part.
type ClientBoundMapItemData struct {
	// MapID is the unique identifier that represents the map that is updated over network. It remains
	// consistent across sessions.
	MapID int64
	// UpdateFlags is a combination of flags found above that indicate what parts of the map should be updated
	// client-side.
	UpdateFlags uint32
	// Dimension is the dimension of the map that should be updated, for example the overworld (0), the nether
	// (1) or the end (2).
	Dimension byte
	// LockedMap specifies if the map that was updated was a locked map, which may be done using a cartography
	// table.
	LockedMap bool
	// Origin is the center position of the map being updated.
	Origin protocol.BlockPos
	// Scale is the scale of the map as it is shown in-game. It is written when any of the MapUpdateFlags are
	// set to the UpdateFlags field.
	Scale byte

	// The following fields apply only for the MapUpdateFlagInitialisation.

	// MapsIncludedIn holds an array of map IDs that the map updated is included in. This has to do with the
	// scale of the map: Each map holds its own map ID and all map IDs of maps that include this map and have
	// a bigger scale. This means that a scale 0 map will have 5 map IDs in this slice, whereas a scale 4 map
	// will have only 1 (its own).
	// The actual use of this field remains unknown.
	MapsIncludedIn []int64

	// The following fields apply only for the MapUpdateFlagDecoration.

	// TrackedObjects is a list of tracked objects on the map, which may either be entities or blocks. The
	// client makes sure these tracked objects are actually tracked. (position updated etc.)
	TrackedObjects []proto.MapTrackedObject
	// Decorations is a list of fixed decorations located on the map. The decorations will not change
	// client-side, unless the server updates them.
	Decorations []protocol.MapDecoration

	// The following fields apply only for the MapUpdateFlagTexture update flag.

	// Height is the height of the texture area that was updated. The height may be a subset of the total
	// height of the map.
	Height int32
	// Width is the width of the texture area that was updated. The width may be a subset of the total width
	// of the map.
	Width int32
	// XOffset is the X offset in pixels at which the updated texture area starts. From this X, the updated
	// texture will extend exactly Width pixels to the right.
	XOffset int32
	// YOffset is the Y offset in pixels at which the updated texture area starts. From this Y, the updated
	// texture will extend exactly Height pixels up.
	YOffset int32
	// Pixels is a list of pixel colours for the new texture of the map. It is indexed as Pixels[y*height + x].
	Pixels []color.RGBA
}

// ID ...
func (*ClientBoundMapItemData) ID() uint32 {
	return packet.IDClientBoundMapItemData
}

func (pk *ClientBoundMapItemData) Marshal(io protocol.IO) {
	io.Varint64(&pk.MapID)
	io.Varuint32(&pk.UpdateFlags)
	io.Uint8(&pk.Dimension)
	io.Bool(&pk.LockedMap)
	io.BlockPos(&pk.Origin)

	if pk.UpdateFlags&MapUpdateFlagInitialisation != 0 {
		protocol.FuncSlice(io, &pk.MapsIncludedIn, io.Varint64)
	}
	if pk.UpdateFlags&(MapUpdateFlagInitialisation|MapUpdateFlagDecoration|MapUpdateFlagTexture) != 0 {
		io.Uint8(&pk.Scale)
	}
	if pk.UpdateFlags&MapUpdateFlagDecoration != 0 {
		protocol.Slice(io, &pk.TrackedObjects)
		protocol.Slice(io, &pk.Decorations)
	}
	if pk.UpdateFlags&MapUpdateFlagTexture != 0 {
		io.Varint32(&pk.Width)
		io.Varint32(&pk.Height)
		io.Varint32(&pk.XOffset)
		io.Varint32(&pk.YOffset)
		protocol.FuncSlice(io, &pk.Pixels, io.VarRGBA)
	}
}

// ToLatest ...
func (pk *ClientBoundMapItemData) ToLatest() *packet.ClientBoundMapItemData {
	trackedObjects := make([]protocol.MapTrackedObject, len(pk.TrackedObjects))
	for i, trackedObject := range pk.TrackedObjects {
		trackedObjects[i] = trackedObject.ToLatest()
	}
	return &packet.ClientBoundMapItemData{
		MapID:          pk.MapID,
		UpdateFlags:    pk.UpdateFlags,
		Dimension:      pk.Dimension,
		LockedMap:      pk.LockedMap,
		Origin:         pk.Origin,
		Scale:          pk.Scale,
		MapsIncludedIn: pk.MapsIncludedIn,
		TrackedObjects: trackedObjects,
		Decorations:    pk.Decorations,
		Height:         pk.Height,
		Width:          pk.Width,
		XOffset:        pk.XOffset,
		YOffset:        pk.YOffset,
		Pixels:         pk.Pixels,
	}
}

// FromLatest ...
func (pk *ClientBoundMapItemData) FromLatest(latest *packet.ClientBoundMapItemData) *ClientBoundMapItemData {
	trackedObjects := make([]proto.MapTrackedObject, len(latest.TrackedObjects))
	for i, trackedObject := range latest.TrackedObjects {
		trackedObjects[i] = (&proto.MapTrackedObject{}).FromLatest(trackedObject)
	}
	pk.MapID = latest.MapID
	pk.UpdateFlags = latest.UpdateFlags
	pk.Dimension = latest.Dimension
	pk.LockedMap = latest.LockedMap
	pk.Origin = latest.Origin
	pk.Scale = latest.Scale
	pk.MapsIncludedIn = latest.MapsIncludedIn
	pk.TrackedObjects = trackedObjects
	pk.Decorations = latest.Decorations
	pk.Height = latest.Height
	pk.Width = latest.Width
	pk.XOffset = latest.XOffset
	pk.YOffset = latest.YOffset
	pk.Pixels = latest.Pixels
	return pk
}
