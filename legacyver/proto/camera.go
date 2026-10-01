package proto

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// CameraPreset represents a basic preset that can be extended upon by more complex instructions.
type CameraPreset struct {
	// Name is the name of the preset. Each preset must have their own unique name.
	Name string
	// Parent is the name of the preset that this preset extends upon. This can be left empty.
	Parent string
	// PosX is the default X position of the camera.
	PosX protocol.Optional[float32]
	// PosY is the default Y position of the camera.
	PosY protocol.Optional[float32]
	// PosZ is the default Z position of the camera.
	PosZ protocol.Optional[float32]
	// RotX is the default pitch of the camera.
	RotX protocol.Optional[float32]
	// RotY is the default yaw of the camera.
	RotY protocol.Optional[float32]
	// RotationSpeed is the speed at which the camera should rotate.
	RotationSpeed protocol.Optional[float32]
	// SnapToTarget determines whether the camera should snap to the target entity or not.
	SnapToTarget protocol.Optional[bool]
	// HorizontalRotationLimit is the horizontal rotation limit of the camera.
	HorizontalRotationLimit protocol.Optional[mgl32.Vec2]
	// VerticalRotationLimit is the vertical rotation limit of the camera.
	VerticalRotationLimit protocol.Optional[mgl32.Vec2]
	// ContinueTargeting determines whether the camera should continue targeting when using aim assist.
	ContinueTargeting protocol.Optional[bool]
	// TrackingRadius is the radius around the camera that the aim assist should track targets.
	TrackingRadius protocol.Optional[float32]
	// ViewOffset is only used in a follow_orbit camera and controls an offset based on a pivot point to the
	// player, causing it to be shifted in a certain direction.
	ViewOffset protocol.Optional[mgl32.Vec2]
	// EntityOffset controls the offset from the entity that the camera should be rendered at.
	EntityOffset protocol.Optional[mgl32.Vec3]
	// Radius is only used in a follow_orbit camera and controls how far away from the player the camera should
	// be rendered.
	Radius protocol.Optional[float32]
	// MinYawLimit is the minimum yaw limit of the camera.
	MinYawLimit protocol.Optional[float32]
	// MaxYawLimit is the maximum yaw limit of the camera.
	MaxYawLimit protocol.Optional[float32]
	// AudioListener defines where the audio should be played from when using this preset. This is one of the
	// constants above.
	AudioListener protocol.Optional[byte]
	// PlayerEffects is currently unknown.
	PlayerEffects protocol.Optional[bool]
	// AlignTargetAndCameraForward determines whether the camera should align the target and the camera forward
	// or not.
	AlignTargetAndCameraForward protocol.Optional[bool]
	// AimAssist defines the aim assist to use when using this preset.
	AimAssist protocol.Optional[protocol.CameraPresetAimAssist]
	// ControlScheme is the control scheme that the client should use in this camera. It is one of the following:
	//  - ControlSchemeLockedPlayerRelativeStrafe is the default behaviour, this cannot be set when the client
	//    is in a custom camera.
	//  - ControlSchemeCameraRelative makes movement relative to the camera's transform, with the client's
	//    rotation being relative to the client's movement.
	//  - ControlSchemeCameraRelativeStrafe makes movement relative to the camera's transform, with the
	//    client's rotation being locked.
	//  - ControlSchemePlayerRelative makes movement relative to the player's transform, meaning holding
	//    left/right will make the player turn in a circle.
	//  - ControlSchemePlayerRelativeStrafe makes movement the same as the default behaviour, but can be
	//    used in a custom camera.
	ControlScheme protocol.Optional[byte]
}

func (x *CameraPreset) FromLatest(cp protocol.CameraPreset) CameraPreset {
	x.Name = cp.Name
	x.Parent = cp.Parent
	x.PosX = cp.PosX
	x.PosY = cp.PosY
	x.PosZ = cp.PosZ
	x.RotX = cp.RotX
	x.RotY = cp.RotY
	x.RotationSpeed = cp.RotationSpeed
	x.SnapToTarget = cp.SnapToTarget
	x.HorizontalRotationLimit = cp.HorizontalRotationLimit
	x.VerticalRotationLimit = cp.VerticalRotationLimit
	x.ContinueTargeting = cp.ContinueTargeting
	x.TrackingRadius = cp.TrackingRadius
	x.ViewOffset = cp.ViewOffset
	x.EntityOffset = cp.EntityOffset
	x.Radius = cp.Radius
	x.AudioListener = cp.AudioListener
	x.PlayerEffects = cp.PlayerEffects
	x.AimAssist = cp.AimAssist
	x.ControlScheme = cp.ControlScheme
	return *x
}

func (x *CameraPreset) ToLatest() protocol.CameraPreset {
	return protocol.CameraPreset{
		Name:                    x.Name,
		Parent:                  x.Parent,
		PosX:                    x.PosX,
		PosY:                    x.PosY,
		PosZ:                    x.PosZ,
		RotX:                    x.RotX,
		RotY:                    x.RotY,
		RotationSpeed:           x.RotationSpeed,
		SnapToTarget:            x.SnapToTarget,
		HorizontalRotationLimit: x.HorizontalRotationLimit,
		VerticalRotationLimit:   x.VerticalRotationLimit,
		ContinueTargeting:       x.ContinueTargeting,
		TrackingRadius:          x.TrackingRadius,
		ViewOffset:              x.ViewOffset,
		EntityOffset:            x.EntityOffset,
		Radius:                  x.Radius,
		MinYawLimit:             x.MinYawLimit,
		MaxYawLimit:             x.MaxYawLimit,
		AudioListener:           x.AudioListener,
		PlayerEffects:           x.PlayerEffects,
		AimAssist:               x.AimAssist,
		ControlScheme:           x.ControlScheme,
	}
}

// Marshal encodes/decodes a CameraPreset.
func (x *CameraPreset) Marshal(r protocol.IO) {
	r.String(&x.Name)
	r.String(&x.Parent)
	protocol.OptionalFunc(r, &x.PosX, r.Float32)
	protocol.OptionalFunc(r, &x.PosY, r.Float32)
	protocol.OptionalFunc(r, &x.PosZ, r.Float32)
	protocol.OptionalFunc(r, &x.RotX, r.Float32)
	protocol.OptionalFunc(r, &x.RotY, r.Float32)
	if IsProtoGTE(r, ID729) {
		protocol.OptionalFunc(r, &x.RotationSpeed, r.Float32)
		protocol.OptionalFunc(r, &x.SnapToTarget, r.Bool)
	}
	if IsProtoGTE(r, ID748) {
		protocol.OptionalFunc(r, &x.HorizontalRotationLimit, r.Vec2)
		protocol.OptionalFunc(r, &x.VerticalRotationLimit, r.Vec2)
		protocol.OptionalFunc(r, &x.ContinueTargeting, r.Bool)
	}
	if IsProtoGTE(r, ID766) {
		protocol.OptionalFunc(r, &x.TrackingRadius, r.Float32)
	}
	if IsProtoGTE(r, ID776) {
		protocol.OptionalFunc(r, &x.MinYawLimit, r.Float32)
		protocol.OptionalFunc(r, &x.MaxYawLimit, r.Float32)
	}
	protocol.OptionalFunc(r, &x.ViewOffset, r.Vec2)
	if IsProtoGTE(r, ID729) {
		protocol.OptionalFunc(r, &x.EntityOffset, r.Vec3)
	}
	protocol.OptionalFunc(r, &x.Radius, r.Float32)
	protocol.OptionalFunc(r, &x.AudioListener, r.Uint8)
	protocol.OptionalFunc(r, &x.PlayerEffects, r.Bool)
	if IsProtoGTE(r, ID748) && IsProtoLT(r, ID818) {
		protocol.OptionalFunc(r, &x.AlignTargetAndCameraForward, r.Bool)
	}
	if IsProtoGTE(r, ID766) {
		protocol.OptionalMarshaler(r, &x.AimAssist)
	}
	if IsProtoGTE(r, ID800) {
		protocol.OptionalFunc(r, &x.ControlScheme, r.Uint8)
	}
}

// CameraInstructionSet represents a camera instruction that sets the camera to a specified preset and can be extended
// with easing functions and translations to the camera's position and rotation.
type CameraInstructionSet struct {
	// Preset is the index of the preset in the CameraPresets packet sent to the player.
	Preset uint32
	// Ease represents the easing function that is used by the instruction.
	Ease protocol.Optional[protocol.CameraEase]
	// Position represents the position of the camera.
	Position protocol.Optional[mgl32.Vec3]
	// Rotation represents the rotation of the camera.
	Rotation protocol.Optional[mgl32.Vec2]
	// Facing is a vector that the camera will always face towards during the duration of the instruction.
	Facing protocol.Optional[mgl32.Vec3]
	// ViewOffset is an offset based on a pivot point to the player, causing the camera to be shifted in a
	// certain direction.
	ViewOffset protocol.Optional[mgl32.Vec2]
	// EntityOffset is an offset from the entity that the camera should be rendered at.
	EntityOffset protocol.Optional[mgl32.Vec3]
	// Default determines whether the camera is a default camera or not.
	Default protocol.Optional[bool]
	// IgnoreStartingValuesComponent behavior is currently unknown.
	IgnoreStartingValuesComponent bool
}

func (x *CameraInstructionSet) FromLatest(cis protocol.CameraInstructionSet) CameraInstructionSet {
	x.Preset = cis.Preset
	x.Ease = cis.Ease
	x.Position = cis.Position
	x.Rotation = cis.Rotation
	x.Facing = cis.Facing
	x.ViewOffset = cis.ViewOffset
	x.EntityOffset = cis.EntityOffset
	x.Default = cis.Default
	x.IgnoreStartingValuesComponent = cis.IgnoreStartingValuesComponent
	return *x
}

func (x *CameraInstructionSet) ToLatest() protocol.CameraInstructionSet {
	return protocol.CameraInstructionSet{
		Preset:                        x.Preset,
		Ease:                          x.Ease,
		Position:                      x.Position,
		Rotation:                      x.Rotation,
		Facing:                        x.Facing,
		ViewOffset:                    x.ViewOffset,
		EntityOffset:                  x.EntityOffset,
		Default:                       x.Default,
		IgnoreStartingValuesComponent: x.IgnoreStartingValuesComponent,
	}
}

// Marshal encodes/decodes a CameraInstructionSet.
func (x *CameraInstructionSet) Marshal(r protocol.IO) {
	r.Uint32(&x.Preset)
	protocol.OptionalMarshaler(r, &x.Ease)
	protocol.OptionalFunc(r, &x.Position, r.Vec3)
	protocol.OptionalFunc(r, &x.Rotation, r.Vec2)
	protocol.OptionalFunc(r, &x.Facing, r.Vec3)
	protocol.OptionalFunc(r, &x.ViewOffset, r.Vec2)
	protocol.OptionalFunc(r, &x.EntityOffset, r.Vec3)
	protocol.OptionalFunc(r, &x.Default, r.Bool)
	if IsProtoGTE(r, ID818) {
		r.Bool(&x.IgnoreStartingValuesComponent)
	}
}

// CameraAimAssistPriorities represents the block and entity specific priorities for targetting. The aim
// assist will select the block or entity with the highest priority within the specified thresholds.
type CameraAimAssistPriorities struct {
	// Entities is a list of priorities for specific entity identifiers.
	Entities []protocol.CameraAimAssistPriority
	// Blocks is a list of priorities for specific block identifiers.
	Blocks []protocol.CameraAimAssistPriority
	// BlockTags is a list of priorities for specific block tags.
	BlockTags []protocol.CameraAimAssistPriority
	// EntityTypeFamilies is a list of priorities for specific entity type families.
	EntityTypeFamilies []protocol.CameraAimAssistPriority
	// EntityDefault is the default priority for entities.
	EntityDefault protocol.Optional[int32]
	// BlockDefault is the default priority for blocks.
	BlockDefault protocol.Optional[int32]
}

// Marshal encodes/decodes a CameraAimAssistPriorities.
func (x *CameraAimAssistPriorities) Marshal(r protocol.IO) {
	protocol.Slice(r, &x.Entities)
	protocol.Slice(r, &x.Blocks)
	protocol.Slice(r, &x.BlockTags)
	if IsProtoGTE(r, ID924) {
		protocol.Slice(r, &x.EntityTypeFamilies)
	}
	protocol.OptionalFunc(r, &x.EntityDefault, r.Int32)
	protocol.OptionalFunc(r, &x.BlockDefault, r.Int32)
}

// ToLatest ...
func (x *CameraAimAssistPriorities) ToLatest() protocol.CameraAimAssistPriorities {
	return protocol.CameraAimAssistPriorities{
		Entities:           x.Entities,
		Blocks:             x.Blocks,
		BlockTags:          x.BlockTags,
		EntityTypeFamilies: x.EntityTypeFamilies,
		EntityDefault:      x.EntityDefault,
		BlockDefault:       x.BlockDefault,
	}
}

// FromLatest ...
func (x *CameraAimAssistPriorities) FromLatest(latest protocol.CameraAimAssistPriorities) CameraAimAssistPriorities {
	x.Entities = latest.Entities
	x.Blocks = latest.Blocks
	x.BlockTags = latest.BlockTags
	x.EntityTypeFamilies = latest.EntityTypeFamilies
	x.EntityDefault = latest.EntityDefault
	x.BlockDefault = latest.BlockDefault
	return *x
}

// CameraAimAssistPreset defines a base preset that can be extended upon when sending an aim assist.
type CameraAimAssistPreset struct {
	// Identifier represents the identifier of this preset.
	Identifier string
	// BlockExclusions is a list of block identifiers that should be ignored by the aim assist.
	BlockExclusions []string
	// EntityExclusions is a list of entity identifiers that should be ignored by the aim assist.
	EntityExclusions []string
	// BlockTagExclusions is a list of block tags that should be ignored by the aim assist.
	BlockTagExclusions []string
	// EntityTypeFamilyExclusions is a list of entity type families that should be ignored by the aim assist.
	EntityTypeFamilyExclusions []string
	// LiquidTargets is a list of entity identifiers that should be targetted when inside of a liquid.
	LiquidTargets []string
	// ItemSettings is a list of settings for specific item identifiers. If an item is not listed here, it
	// will fallback to DefaultItemSettings or HandSettings if no item is held.
	ItemSettings []protocol.CameraAimAssistItemSettings
	// DefaultItemSettings is the identifier of a category to use when the player is not holding an item
	// listed in ItemSettings. This must be the identifier of a category within the Categories slice.
	DefaultItemSettings protocol.Optional[string]
	// HandSettings is the identifier of a category to use when the player is not holding an item. This must
	// be the identifier of a category within Categories slice.
	HandSettings protocol.Optional[string]
}

// Marshal encodes/decodes a CameraAimAssistPreset.
func (x *CameraAimAssistPreset) Marshal(r protocol.IO) {
	r.String(&x.Identifier)
	protocol.FuncSlice(r, &x.BlockExclusions, r.String)
	protocol.FuncSlice(r, &x.EntityExclusions, r.String)
	protocol.FuncSlice(r, &x.BlockTagExclusions, r.String)
	if IsProtoGTE(r, ID924) {
		protocol.FuncSlice(r, &x.EntityTypeFamilyExclusions, r.String)
	}
	protocol.FuncSlice(r, &x.LiquidTargets, r.String)
	protocol.Slice(r, &x.ItemSettings)
	protocol.OptionalFunc(r, &x.DefaultItemSettings, r.String)
	protocol.OptionalFunc(r, &x.HandSettings, r.String)
}

// ToLatest ...
func (x *CameraAimAssistPreset) ToLatest() protocol.CameraAimAssistPreset {
	return protocol.CameraAimAssistPreset{
		Identifier:                 x.Identifier,
		BlockExclusions:            x.BlockExclusions,
		EntityExclusions:           x.EntityExclusions,
		BlockTagExclusions:         x.BlockTagExclusions,
		EntityTypeFamilyExclusions: x.EntityTypeFamilyExclusions,
		LiquidTargets:              x.LiquidTargets,
		ItemSettings:               x.ItemSettings,
		DefaultItemSettings:        x.DefaultItemSettings,
		HandSettings:               x.HandSettings,
	}
}

// FromLatest ...
func (x *CameraAimAssistPreset) FromLatest(latest protocol.CameraAimAssistPreset) CameraAimAssistPreset {
	x.Identifier = latest.Identifier
	x.BlockExclusions = latest.BlockExclusions
	x.EntityExclusions = latest.EntityExclusions
	x.BlockTagExclusions = latest.BlockTagExclusions
	x.EntityTypeFamilyExclusions = latest.EntityTypeFamilyExclusions
	x.LiquidTargets = latest.LiquidTargets
	x.ItemSettings = latest.ItemSettings
	x.DefaultItemSettings = latest.DefaultItemSettings
	x.HandSettings = latest.HandSettings
	return *x
}

// CameraRotationOption represents a rotation option for camera spline instructions.
type CameraRotationOption struct {
	// Value is the rotation value.
	Value mgl32.Vec3
	// Time is the time for this rotation option.
	Time float32
	// EaseType is the optional easing function used to interpolate towards this rotation key frame.
	// This is one of the EasingType constants.
	EaseType int32
}

// Marshal encodes/decodes a CameraRotationOption.
func (x *CameraRotationOption) Marshal(r protocol.IO) {
	easingType := easingTypeToString(x.EaseType)
	r.Vec3(&x.Value)
	r.Float32(&x.Time)
	if IsProtoGTE(r, ID924) {
		r.String(&easingType)
		easingTypeFromString(r, &x.EaseType, easingType)
	}
}

// ToLatest ...
func (x *CameraRotationOption) ToLatest() protocol.CameraRotationOption {
	return protocol.CameraRotationOption{
		Value:    x.Value,
		Time:     x.Time,
		EaseType: x.EaseType,
	}
}

// FromLatest ...
func (x *CameraRotationOption) FromLatest(latest protocol.CameraRotationOption) CameraRotationOption {
	x.Value = latest.Value
	x.Time = latest.Time
	x.EaseType = latest.EaseType
	return *x
}

// CameraSplineInstruction represents a camera instruction that creates a spline path for the camera to follow.
type CameraSplineInstruction struct {
	// TotalTime is the total time for the spline animation.
	TotalTime float32
	// SplineType is the optional spline interpolation type. This is one of the SplineEaseType constants.
	SplineType protocol.Optional[uint8]
	// Curve is a list of points that define the spline curve.
	Curve []mgl32.Vec3
	// ProgressKeyFrames is a list of progress key frames for the spline.
	ProgressKeyFrames []protocol.CameraProgressOption
	// RotationOptions is a list of rotation options for the spline.
	RotationOptions []CameraRotationOption
	// SplineIdentifier is an optional identifier for referencing the spline by name.
	SplineIdentifier protocol.Optional[string]
	// LoadFromJson optionally determines whether the spline should be loaded from a JSON definition.
	LoadFromJson protocol.Optional[bool]
}

// Marshal encodes/decodes a CameraSplineInstruction.
func (x *CameraSplineInstruction) Marshal(r protocol.IO) {
	r.Float32(&x.TotalTime)
	if IsProtoGTE(r, ID924) {
		protocol.OptionalFunc(r, &x.SplineType, r.Uint8)
	} else {
		easeType, _ := x.SplineType.Value()
		r.Uint8(&easeType)
		x.SplineType = protocol.Option(easeType)
	}
	protocol.FuncSlice(r, &x.Curve, r.Vec3)
	if IsProtoGTE(r, ID924) {
		protocol.Slice(r, &x.ProgressKeyFrames)
	} else {
		progressKeyFrames := make([]mgl32.Vec2, len(x.ProgressKeyFrames))
		for i := 0; i < len(x.ProgressKeyFrames); i++ {
			progressKeyFrames[i] = mgl32.Vec2{x.ProgressKeyFrames[i].Value, x.ProgressKeyFrames[i].Time}
		}
		protocol.FuncSlice(r, &progressKeyFrames, r.Vec2)
		x.ProgressKeyFrames = make([]protocol.CameraProgressOption, len(progressKeyFrames))
		for i := 0; i < len(progressKeyFrames); i++ {
			x.ProgressKeyFrames[i].Value = progressKeyFrames[i][0]
			x.ProgressKeyFrames[i].Time = progressKeyFrames[i][1]
		}
	}
	protocol.Slice(r, &x.RotationOptions)
	if IsProtoGTE(r, ID944) {
		protocol.OptionalFunc(r, &x.SplineIdentifier, r.String)
		protocol.OptionalFunc(r, &x.LoadFromJson, r.Bool)
	}
}

// ToLatest ...
func (x *CameraSplineInstruction) ToLatest() protocol.CameraSplineInstruction {
	rotationOptions := make([]protocol.CameraRotationOption, len(x.RotationOptions))
	for i, option := range x.RotationOptions {
		rotationOptions[i] = option.ToLatest()
	}
	return protocol.CameraSplineInstruction{
		TotalTime:         x.TotalTime,
		SplineType:        x.SplineType,
		Curve:             x.Curve,
		ProgressKeyFrames: x.ProgressKeyFrames,
		RotationOptions:   rotationOptions,
		SplineIdentifier:  x.SplineIdentifier,
		LoadFromJson:      x.LoadFromJson,
	}
}

// FromLatest ...
func (x *CameraSplineInstruction) FromLatest(latest protocol.CameraSplineInstruction) CameraSplineInstruction {
	x.TotalTime = latest.TotalTime
	x.SplineType = latest.SplineType
	x.Curve = latest.Curve
	x.ProgressKeyFrames = latest.ProgressKeyFrames
	x.RotationOptions = make([]CameraRotationOption, len(latest.RotationOptions))
	for i, option := range latest.RotationOptions {
		var legacyOption CameraRotationOption
		x.RotationOptions[i] = legacyOption.FromLatest(option)
	}
	x.SplineIdentifier = latest.SplineIdentifier
	x.LoadFromJson = latest.LoadFromJson
	return *x
}

// CameraSplineDefinition represents a named camera spline definition.
type CameraSplineDefinition struct {
	// Name is the name of the spline definition.
	Name string
	// TotalTime is the total time for the spline animation.
	TotalTime float32
	// SplineType is the optional spline interpolation type.
	SplineType protocol.Optional[string]
	// ControlPoints is a list of points that define the spline curve.
	ControlPoints []mgl32.Vec3
	// ProgressKeyFrames is a list of progress key frames for the spline.
	ProgressKeyFrames []protocol.CameraProgressOption
	// RotationKeyFrames is a list of rotation key frames for the spline.
	RotationKeyFrames []CameraRotationOption
}

// Marshal encodes/decodes a CameraSplineDefinition.
func (x *CameraSplineDefinition) Marshal(r protocol.IO) {
	r.String(&x.Name)
	r.Float32(&x.TotalTime)
	protocol.OptionalFunc(r, &x.SplineType, r.String)
	protocol.FuncSlice(r, &x.ControlPoints, r.Vec3)
	protocol.Slice(r, &x.ProgressKeyFrames)
	protocol.Slice(r, &x.RotationKeyFrames)
}

// ToLatest ...
func (x *CameraSplineDefinition) ToLatest() protocol.CameraSplineDefinition {
	rotationKeyFrames := make([]protocol.CameraRotationOption, len(x.RotationKeyFrames))
	for i, option := range x.RotationKeyFrames {
		rotationKeyFrames[i] = option.ToLatest()
	}
	return protocol.CameraSplineDefinition{
		Name:              x.Name,
		TotalTime:         x.TotalTime,
		SplineType:        x.SplineType,
		ControlPoints:     x.ControlPoints,
		ProgressKeyFrames: x.ProgressKeyFrames,
		RotationKeyFrames: rotationKeyFrames,
	}
}

// FromLatest ...
func (x *CameraSplineDefinition) FromLatest(latest protocol.CameraSplineDefinition) CameraSplineDefinition {
	x.Name = latest.Name
	x.TotalTime = latest.TotalTime
	x.SplineType = latest.SplineType
	x.ControlPoints = latest.ControlPoints
	x.ProgressKeyFrames = latest.ProgressKeyFrames
	x.RotationKeyFrames = make([]CameraRotationOption, len(latest.RotationKeyFrames))
	for i, option := range latest.RotationKeyFrames {
		var legacyOption CameraRotationOption
		x.RotationKeyFrames[i] = legacyOption.FromLatest(option)
	}
	return *x
}

// CameraAimAssistCategory is an aim assist category that defines priorities for specific blocks and entities.
type CameraAimAssistCategory struct {
	// Name is the name of the category which can be used by a CameraAimAssistPreset.
	Name string
	// Priorities represents the block and entity specific priorities as well as the default priorities for
	// this category.
	Priorities CameraAimAssistPriorities
}

// Marshal encodes/decodes a CameraAimAssistCategory.
func (x *CameraAimAssistCategory) Marshal(r protocol.IO) {
	r.String(&x.Name)
	protocol.Single(r, &x.Priorities)
}

// ToLatest ...
func (x *CameraAimAssistCategory) ToLatest() protocol.CameraAimAssistCategory {
	return protocol.CameraAimAssistCategory{
		Name:       x.Name,
		Priorities: x.Priorities.ToLatest(),
	}
}

// FromLatest ...
func (x *CameraAimAssistCategory) FromLatest(latest protocol.CameraAimAssistCategory) CameraAimAssistCategory {
	x.Name = latest.Name
	x.Priorities = x.Priorities.FromLatest(latest.Priorities)
	return *x
}

const (
	EasingTypeLinear = iota
	EasingTypeSpring
	EasingTypeInQuad
	EasingTypeOutQuad
	EasingTypeInOutQuad
	EasingTypeInCubic
	EasingTypeOutCubic
	EasingTypeInOutCubic
	EasingTypeInQuart
	EasingTypeOutQuart
	EasingTypeInOutQuart
	EasingTypeInQuint
	EasingTypeOutQuint
	EasingTypeInOutQuint
	EasingTypeInSine
	EasingTypeOutSine
	EasingTypeInOutSine
	EasingTypeInExpo
	EasingTypeOutExpo
	EasingTypeInOutExpo
	EasingTypeInCirc
	EasingTypeOutCirc
	EasingTypeInOutCirc
	EasingTypeInBounce
	EasingTypeOutBounce
	EasingTypeInOutBounce
	EasingTypeInBack
	EasingTypeOutBack
	EasingTypeInOutBack
	EasingTypeInElastic
	EasingTypeOutElastic
	EasingTypeInOutElastic
	EasingTypeInverseLerp
)

// easingTypeFromString looks up an easing type from a string and writes the result to x.
func easingTypeFromString(io protocol.IO, x *int32, s string) {
	switch s {
	case "linear":
		*x = EasingTypeLinear
	case "spring":
		*x = EasingTypeSpring
	case "in_quad":
		*x = EasingTypeInQuad
	case "out_quad":
		*x = EasingTypeOutQuad
	case "in_out_quad":
		*x = EasingTypeInOutQuad
	case "in_cubic":
		*x = EasingTypeInCubic
	case "out_cubic":
		*x = EasingTypeOutCubic
	case "in_out_cubic":
		*x = EasingTypeInOutCubic
	case "in_quart":
		*x = EasingTypeInQuart
	case "out_quart":
		*x = EasingTypeOutQuart
	case "in_out_quart":
		*x = EasingTypeInOutQuart
	case "in_quint":
		*x = EasingTypeInQuint
	case "out_quint":
		*x = EasingTypeOutQuint
	case "in_out_quint":
		*x = EasingTypeInOutQuint
	case "in_sine":
		*x = EasingTypeInSine
	case "out_sine":
		*x = EasingTypeOutSine
	case "in_out_sine":
		*x = EasingTypeInOutSine
	case "in_expo":
		*x = EasingTypeInExpo
	case "out_expo":
		*x = EasingTypeOutExpo
	case "in_out_expo":
		*x = EasingTypeInOutExpo
	case "in_circ":
		*x = EasingTypeInCirc
	case "out_circ":
		*x = EasingTypeOutCirc
	case "in_out_circ":
		*x = EasingTypeInOutCirc
	case "in_back":
		*x = EasingTypeInBack
	case "out_back":
		*x = EasingTypeOutBack
	case "in_out_back":
		*x = EasingTypeInOutBack
	case "in_elastic":
		*x = EasingTypeInElastic
	case "out_elastic":
		*x = EasingTypeOutElastic
	case "in_out_elastic":
		*x = EasingTypeInOutElastic
	case "in_bounce":
		*x = EasingTypeInBounce
	case "out_bounce":
		*x = EasingTypeOutBounce
	case "in_out_bounce":
		*x = EasingTypeInOutBounce
	case "inverse_lerp":
		*x = EasingTypeInverseLerp
	default:
		io.InvalidValue(s, "easingType", "unknown easing type")
	}
}

// easingTypeToString looks up an easing type constant and returns the string representation.
func easingTypeToString(x int32) string {
	switch x {
	case EasingTypeLinear:
		return "linear"
	case EasingTypeSpring:
		return "spring"
	case EasingTypeInQuad:
		return "in_quad"
	case EasingTypeOutQuad:
		return "out_quad"
	case EasingTypeInOutQuad:
		return "in_out_quad"
	case EasingTypeInCubic:
		return "in_cubic"
	case EasingTypeOutCubic:
		return "out_cubic"
	case EasingTypeInOutCubic:
		return "in_out_cubic"
	case EasingTypeInQuart:
		return "in_quart"
	case EasingTypeOutQuart:
		return "out_quart"
	case EasingTypeInOutQuart:
		return "in_out_quart"
	case EasingTypeInQuint:
		return "in_quint"
	case EasingTypeOutQuint:
		return "out_quint"
	case EasingTypeInOutQuint:
		return "in_out_quint"
	case EasingTypeInSine:
		return "in_sine"
	case EasingTypeOutSine:
		return "out_sine"
	case EasingTypeInOutSine:
		return "in_out_sine"
	case EasingTypeInExpo:
		return "in_expo"
	case EasingTypeOutExpo:
		return "out_expo"
	case EasingTypeInOutExpo:
		return "in_out_expo"
	case EasingTypeInCirc:
		return "in_circ"
	case EasingTypeOutCirc:
		return "out_circ"
	case EasingTypeInOutCirc:
		return "in_out_circ"
	case EasingTypeInBack:
		return "in_back"
	case EasingTypeOutBack:
		return "out_back"
	case EasingTypeInOutBack:
		return "in_out_back"
	case EasingTypeInElastic:
		return "in_elastic"
	case EasingTypeOutElastic:
		return "out_elastic"
	case EasingTypeInOutElastic:
		return "in_out_elastic"
	case EasingTypeInBounce:
		return "in_bounce"
	case EasingTypeOutBounce:
		return "out_bounce"
	case EasingTypeInOutBounce:
		return "in_out_bounce"
	case EasingTypeInverseLerp:
		return "inverse_lerp"
	default:
		return "unknown"
	}
}
