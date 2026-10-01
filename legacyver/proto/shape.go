package proto

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"image/color"
)

// PrimitiveShape defines a single debug shape to be rendered on the client.
// Each shape has a unique NetworkID and a set of optional parameters depending on its type.
type PrimitiveShape struct {
	// NetworkID is the network ID of the shape.
	NetworkID uint64
	// DimensionID is the optional dimension ID where the shape is rendered.
	DimensionID protocol.Optional[int32]
	// AttachedToEntityID is the optional runtime ID of the entity the shape is attached to.
	// should be uint64 pre-975
	AttachedToEntityID protocol.Optional[int64]
	// Type is the type of the shape.
	// If not set, the set shape will be cleared.
	Type protocol.Optional[uint8]
	// Location is the location of the shape.
	Location protocol.Optional[mgl32.Vec3]
	// Scale is the scale of the shape.
	Scale protocol.Optional[float32]
	// Rotation is the rotation of the shape.
	Rotation protocol.Optional[mgl32.Vec3]
	// TotalTimeLeft is the total time left of the shape.
	TotalTimeLeft protocol.Optional[float32]
	// Colour is the ARGB colour of the shape.
	Colour protocol.Optional[color.RGBA]
	// ExtraShapeData holding data specific to the type of shape (such as text string for the text shape).
	ExtraShapeData protocol.ShapeData
}

// FromLatest ...
func (x *PrimitiveShape) FromLatest(y protocol.PrimitiveShape) PrimitiveShape {
	x.NetworkID = y.NetworkID
	x.DimensionID = y.DimensionID
	x.AttachedToEntityID = y.AttachedToEntityID
	x.Type = y.Type
	x.Location = y.Location
	x.Scale = y.Scale
	x.Rotation = y.Rotation
	x.TotalTimeLeft = y.TotalTimeLeft
	x.Colour = y.Colour
	extraShapeData := y.ExtraShapeData
	switch data := extraShapeData.(type) {
	case *protocol.TextShape:
		extraShapeData = (&TextShape{}).FromLatest(*data)
	}
	x.ExtraShapeData = extraShapeData
	return *x
}

// ToLatest ...
func (x *PrimitiveShape) ToLatest() protocol.PrimitiveShape {
	extraShapeData := x.ExtraShapeData
	switch data := extraShapeData.(type) {
	case *TextShape:
		extraShapeData = data.ToLatest()
	}
	return protocol.PrimitiveShape{
		NetworkID:          x.NetworkID,
		DimensionID:        x.DimensionID,
		AttachedToEntityID: x.AttachedToEntityID,
		Type:               x.Type,
		Location:           x.Location,
		Scale:              x.Scale,
		Rotation:           x.Rotation,
		TotalTimeLeft:      x.TotalTimeLeft,
		Colour:             x.Colour,
		ExtraShapeData:     extraShapeData,
	}
}

// Marshal ...
func (x *PrimitiveShape) Marshal(io protocol.IO) {
	io.Varuint64(&x.NetworkID)
	protocol.OptionalFunc(io, &x.Type, io.Uint8)
	protocol.OptionalFunc(io, &x.Location, io.Vec3)
	protocol.OptionalFunc(io, &x.Scale, io.Float32)
	protocol.OptionalFunc(io, &x.Rotation, io.Vec3)
	protocol.OptionalFunc(io, &x.TotalTimeLeft, io.Float32)
	protocol.OptionalFunc(io, &x.Colour, io.BEARGB)
	if IsProtoGTE(io, ID924) {
		protocol.OptionalFunc(io, &x.DimensionID, io.Varint32)
		protocol.OptionalFunc(io, &x.AttachedToEntityID, io.Varint64)
	} else {
		dimensionId, _ := x.DimensionID.Value()
		io.Varint32(&dimensionId)
		x.DimensionID = protocol.Option(dimensionId)
	}
	IOShapeData(io, &x.ExtraShapeData)
}

// TextShape represents a text debug shape.
type TextShape struct {
	// Text is the text of the debug text shape.
	Text string
	// UseRotation is if the text should use the provided rotation, meaning it will be static and does not follow the
	// camera. Use false for default behaviour.
	UseRotation bool
	// BackgroundColour is the RGBA colour to use for the text background. This is a translucent black colour by default.
	BackgroundColour protocol.Optional[color.RGBA]
	// DepthTest is whether the text should show through walls. Use true for default behaviour.
	DepthTest bool
	// ShowBackface is if the background should render on the back side of the shape. This only has a visible effect when
	// UseRotation is true since you cannot see the back side of the text otherwise. Use true for default behaviour.
	ShowBackface bool
	// ShowBackfaceText is if the text should render on the back side of the shape. This only has a visible effect when
	// UseRotation is true since you cannot see the back side of the text otherwise. Use true for default behaviour.
	ShowBackfaceText bool
}

// Marshal ...
func (shape *TextShape) Marshal(io protocol.IO) {
	io.String(&shape.Text)
	if IsProtoGTE(io, ID975) {
		io.Bool(&shape.UseRotation)
		protocol.OptionalFunc(io, &shape.BackgroundColour, io.BEARGB)
		io.Bool(&shape.DepthTest)
		io.Bool(&shape.ShowBackface)
		io.Bool(&shape.ShowBackfaceText)
	}
}

// ToLatest ...
func (shape *TextShape) ToLatest() *protocol.TextShape {
	return &protocol.TextShape{
		Text:             shape.Text,
		UseRotation:      shape.UseRotation,
		BackgroundColour: shape.BackgroundColour,
		DepthTest:        shape.DepthTest,
		ShowBackface:     shape.ShowBackface,
		ShowBackfaceText: shape.ShowBackfaceText,
	}
}

// FromLatest ...
func (shape *TextShape) FromLatest(latest protocol.TextShape) *TextShape {
	shape.Text = latest.Text
	shape.UseRotation = latest.UseRotation
	shape.BackgroundColour = latest.BackgroundColour
	shape.DepthTest = latest.DepthTest
	shape.ShowBackface = latest.ShowBackface
	shape.ShowBackfaceText = latest.ShowBackfaceText
	return shape
}

// lookupShapeData looks up an ShapeData matching the shape data type passed.
// False is returned if no such shape data exists.
func lookupShapeData(shapeDataType uint32, x *protocol.ShapeData) bool {
	switch shapeDataType {
	case protocol.ShapeDataLast:
		*x = &protocol.LastShape{}
	case protocol.ShapeDataArrow:
		*x = &protocol.ArrowShape{}
	case protocol.ShapeDataText:
		*x = &TextShape{}
	case protocol.ShapeDataBox:
		*x = &protocol.BoxShape{}
	case protocol.ShapeDataLine:
		*x = &protocol.LineShape{}
	case protocol.ShapeDataSphere:
		*x = &protocol.SphereShape{}
	default:
		return false
	}
	return true
}

// lookupShapeDataType looks up a debug shape type that matches the ShapeData passed.
func lookupShapeDataType(x protocol.ShapeData, shapeDataType *uint32) bool {
	switch x.(type) {
	case *protocol.LastShape:
		*shapeDataType = protocol.ShapeDataLast
	case *protocol.ArrowShape:
		*shapeDataType = protocol.ShapeDataArrow
	case *TextShape:
		*shapeDataType = protocol.ShapeDataText
	case *protocol.BoxShape:
		*shapeDataType = protocol.ShapeDataBox
	case *protocol.LineShape:
		*shapeDataType = protocol.ShapeDataLine
	case *protocol.SphereShape:
		*shapeDataType = protocol.ShapeDataSphere
	default:
		return false
	}
	return true
}

// IOShapeData ...
func IOShapeData(io protocol.IO, x *protocol.ShapeData) {
	if IsWriter(io) {
		var shapeDataType uint32
		if !lookupShapeDataType(*x, &shapeDataType) {
			io.UnknownEnumOption(fmt.Sprintf("%T", *x), "debug shape data type")
		}
		io.Varuint32(&shapeDataType)
		(*x).Marshal(io)
		return
	}
	var shapeDataType uint32
	io.Varuint32(&shapeDataType)
	if !lookupShapeData(shapeDataType, x) {
		io.UnknownEnumOption(shapeDataType, "debug shape data type")
		return
	}
	(*x).Marshal(io)
}
