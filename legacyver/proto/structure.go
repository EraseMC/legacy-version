package proto

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// StructureSettings is a struct holding settings of a structure block. Its fields may be changed using the
// in-game UI on the client-side.
type StructureSettings struct {
	// PaletteName is the name of the palette used in the structure. Currently, it seems that this field is
	// always 'default'.
	PaletteName string
	// IgnoreEntities specifies if the structure should ignore entities or include them. If set to false,
	// entities will also show up in the exported structure.
	IgnoreEntities bool
	// IgnoreBlocks specifies if the structure should ignore blocks or include them. If set to false, blocks
	// will show up in the exported structure.
	IgnoreBlocks bool
	// AllowNonTickingChunks specifies if the structure should allow non-ticking chunks. If set to false, the structure
	// will export non-ticking chunks.
	AllowNonTickingChunks bool
	// Size is the size of the area that is about to be exported. The area exported will start at the
	// Position + Offset, and will extend as far as Size specifies.
	Size protocol.BlockPos
	// Offset is the offset position that was set in the structure block. The area exported is offset by this
	// position.
	Offset protocol.BlockPos
	// LastEditingPlayerUniqueID is the unique ID of the player that last edited the structure block that
	// these settings concern.
	LastEditingPlayerUniqueID int64
	// Rotation is the rotation that the structure block should obtain. See the constants above for available
	// options.
	Rotation byte
	// Mirror specifies the way the structure should be mirrored. It is either no mirror at all, mirror on the
	// x/z axis or both.
	Mirror byte
	// AnimationMode ...
	AnimationMode byte
	// AnimationDuration ...
	AnimationDuration float32
	// Integrity is usually 1, but may be set to a number between 0 and 1 to omit blocks randomly, using
	// the Seed that follows.
	Integrity float32
	// Seed is the seed used to omit blocks if Integrity is not equal to one. If the Seed is 0, a random
	// seed is selected to omit blocks.
	Seed uint32
	// Pivot is the pivot around which the structure may be rotated.
	Pivot mgl32.Vec3
}

// Marshal reads/writes StructureSettings x using IO r.
func (x *StructureSettings) Marshal(r protocol.IO) {
	r.String(&x.PaletteName)
	r.Bool(&x.IgnoreEntities)
	r.Bool(&x.IgnoreBlocks)
	r.Bool(&x.AllowNonTickingChunks)
	if IsProtoGTE(r, ID944) {
		r.BlockPos(&x.Size)
		r.BlockPos(&x.Offset)
	} else {
		UBlockPos(r, &x.Size)
		UBlockPos(r, &x.Offset)
	}
	r.Varint64(&x.LastEditingPlayerUniqueID)
	r.Uint8(&x.Rotation)
	r.Uint8(&x.Mirror)
	r.Uint8(&x.AnimationMode)
	r.Float32(&x.AnimationDuration)
	r.Float32(&x.Integrity)
	r.Uint32(&x.Seed)
	r.Vec3(&x.Pivot)
}

// ToLatest ...
func (x *StructureSettings) ToLatest() protocol.StructureSettings {
	return protocol.StructureSettings{
		PaletteName:               x.PaletteName,
		IgnoreEntities:            x.IgnoreEntities,
		IgnoreBlocks:              x.IgnoreBlocks,
		AllowNonTickingChunks:     x.AllowNonTickingChunks,
		Size:                      x.Size,
		Offset:                    x.Offset,
		LastEditingPlayerUniqueID: x.LastEditingPlayerUniqueID,
		Rotation:                  x.Rotation,
		Mirror:                    x.Mirror,
		AnimationMode:             x.AnimationMode,
		AnimationDuration:         x.AnimationDuration,
		Integrity:                 x.Integrity,
		Seed:                      x.Seed,
		Pivot:                     x.Pivot,
	}
}

// FromLatest ...
func (x *StructureSettings) FromLatest(latest protocol.StructureSettings) StructureSettings {
	x.PaletteName = latest.PaletteName
	x.IgnoreEntities = latest.IgnoreEntities
	x.IgnoreBlocks = latest.IgnoreBlocks
	x.AllowNonTickingChunks = latest.AllowNonTickingChunks
	x.Size = latest.Size
	x.Offset = latest.Offset
	x.LastEditingPlayerUniqueID = latest.LastEditingPlayerUniqueID
	x.Rotation = latest.Rotation
	x.Mirror = latest.Mirror
	x.AnimationMode = latest.AnimationMode
	x.AnimationDuration = latest.AnimationDuration
	x.Integrity = latest.Integrity
	x.Seed = latest.Seed
	x.Pivot = latest.Pivot
	return *x
}
