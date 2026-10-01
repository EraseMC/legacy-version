package proto

import (
	"github.com/google/uuid"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// GatheringJoinInfo contains information about the gathering (experience) the player is joining.
type GatheringJoinInfo struct {
	// ExperienceID is the UUID of the experience.
	ExperienceID uuid.UUID
	// ExperienceName is the name of the experience.
	ExperienceName string
	// ExperienceWorldID is the UUID of the experience world.
	ExperienceWorldID uuid.UUID
	// ExperienceWorldName is the world name of the experience.
	ExperienceWorldName string
	// CreatorID is the ID of the creator.
	CreatorID string
	// TargetID is the session ID of the experience.
	TargetID uuid.UUID
	// ScenarioID is the scenario ID of experience.
	ScenarioID string
	// ServerID is the server identifier.
	ServerID string
}

// Marshal encodes/decodes a GatheringJoinInfo.
func (x *GatheringJoinInfo) Marshal(r protocol.IO) {
	if IsProtoGTE(r, ID944) {
		r.UUID(&x.ExperienceID)
	} else {
		IOStringUUID(r, &x.ExperienceID)
	}
	r.String(&x.ExperienceName)
	if IsProtoGTE(r, ID944) {
		r.UUID(&x.ExperienceWorldID)
	} else {
		IOStringUUID(r, &x.ExperienceWorldID)
	}
	r.String(&x.ExperienceWorldName)
	r.String(&x.CreatorID)
	if IsProtoGTE(r, ID944) {
		r.UUID(&x.TargetID)
		if IsProtoGTE(r, ID975) {
			r.String(&x.ScenarioID)
		} else {
			if IsWriter(r) {
				scenarioUUID, _ := uuid.Parse(x.ScenarioID)
				r.UUID(&scenarioUUID)
				x.ScenarioID = scenarioUUID.String()
			} else {
				var scenarioUUID uuid.UUID
				r.UUID(&scenarioUUID)
				x.ScenarioID = scenarioUUID.String()
			}
		}
		r.String(&x.ServerID)
	}
}

// ToLatest ...
func (x *GatheringJoinInfo) ToLatest() protocol.GatheringJoinInfo {
	return protocol.GatheringJoinInfo{
		ExperienceID:        x.ExperienceID,
		ExperienceName:      x.ExperienceName,
		ExperienceWorldID:   x.ExperienceWorldID,
		ExperienceWorldName: x.ExperienceWorldName,
		CreatorID:           x.CreatorID,
		TargetID:            x.TargetID,
		ScenarioID:          x.ScenarioID,
		ServerID:            x.ServerID,
	}
}

// FromLatest ...
func (x *GatheringJoinInfo) FromLatest(latest protocol.GatheringJoinInfo) GatheringJoinInfo {
	x.ExperienceID = latest.ExperienceID
	x.ExperienceName = latest.ExperienceName
	x.ExperienceWorldID = latest.ExperienceWorldID
	x.ExperienceWorldName = latest.ExperienceWorldName
	x.CreatorID = latest.CreatorID
	x.TargetID = latest.TargetID
	x.ScenarioID = latest.ScenarioID
	x.ServerID = latest.ServerID
	return *x
}

// ServerJoinInformation contains optional information about the server the player is joining.
type ServerJoinInformation struct {
	// GatheringJoinInfo is optional information about the gathering being joined.
	GatheringJoinInfo protocol.Optional[GatheringJoinInfo]
	// StoreEntryPointInfo is optional information about the store entry point.
	StoreEntryPointInfo protocol.Optional[protocol.StoreEntryPointInfo]
	// PresenceInfo is optional presence information.
	PresenceInfo protocol.Optional[protocol.PresenceInfo]
}

// Marshal encodes/decodes a ServerJoinInformation.
func (x *ServerJoinInformation) Marshal(r protocol.IO) {
	protocol.OptionalMarshaler(r, &x.GatheringJoinInfo)
	if IsProtoGTE(r, ID944) {
		protocol.OptionalMarshaler(r, &x.StoreEntryPointInfo)
		protocol.OptionalMarshaler(r, &x.PresenceInfo)
	}
}

// ToLatest ...
func (x *ServerJoinInformation) ToLatest() (ret protocol.ServerJoinInformation) {
	if v, ok := x.GatheringJoinInfo.Value(); ok {
		ret.GatheringJoinInfo = protocol.Option(v.ToLatest())
	}
	ret.StoreEntryPointInfo = x.StoreEntryPointInfo
	ret.PresenceInfo = x.PresenceInfo
	return ret
}

// FromLatest ...
func (x *ServerJoinInformation) FromLatest(latest protocol.ServerJoinInformation) ServerJoinInformation {
	if v, ok := latest.GatheringJoinInfo.Value(); ok {
		x.GatheringJoinInfo = protocol.Option((&GatheringJoinInfo{}).FromLatest(v))
	}
	x.StoreEntryPointInfo = latest.StoreEntryPointInfo
	x.PresenceInfo = latest.PresenceInfo
	return *x
}
