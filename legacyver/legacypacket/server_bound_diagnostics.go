package legacypacket

import (
	"github.com/akmalfairuz/legacy-version/legacyver/proto"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

// ServerBoundDiagnostics is sent by the client to tell the server about the performance diagnostics
// of the client. It is sent by the client roughly every 500ms or 10 in-game ticks when the
// "Creator > Enable Client Diagnostics" setting is enabled.
type ServerBoundDiagnostics struct {
	// AverageFramesPerSecond is the average amount of frames per second that the client has been
	// running at.
	AverageFramesPerSecond float32
	// AverageServerSimTickTime is the average time that the server spends simulating a single tick
	// in milliseconds.
	AverageServerSimTickTime float32
	// AverageClientSimTickTime is the average time that the client spends simulating a single tick
	// in milliseconds.
	AverageClientSimTickTime float32
	// AverageBeginFrameTime is the average time that the client spends beginning a frame in
	// milliseconds.
	AverageBeginFrameTime float32
	// AverageInputTime is the average time that the client spends processing input in milliseconds.
	AverageInputTime float32
	// AverageRenderTime is the average time that the client spends rendering in milliseconds.
	AverageRenderTime float32
	// AverageEndFrameTime is the average time that the client spends ending a frame in milliseconds.
	AverageEndFrameTime float32
	// AverageRemainderTimePercent is the average percentage of time that the client spends on
	// tasks that are not accounted for.
	AverageRemainderTimePercent float32
	// AverageUnaccountedTimePercent is the average percentage of time that the client spends on
	// unaccounted tasks.
	AverageUnaccountedTimePercent float32
	// MemoryCategoryValues is a list of memory category counters sent by the client.
	MemoryCategoryValues []protocol.MemoryCategoryCounter
	// EntityDiagnostics is a list of entity timing entries sent by the client.
	EntityDiagnostics []protocol.EntityDiagnosticTimingInfo
	// SystemDiagnostics is a list of system timing entries sent by the client.
	SystemDiagnostics []protocol.SystemDiagnosticTimingInfo
}

// ID ...
func (*ServerBoundDiagnostics) ID() uint32 {
	return packet.IDServerBoundDiagnostics
}

func (pk *ServerBoundDiagnostics) Marshal(io protocol.IO) {
	io.Float32(&pk.AverageFramesPerSecond)
	io.Float32(&pk.AverageServerSimTickTime)
	io.Float32(&pk.AverageClientSimTickTime)
	io.Float32(&pk.AverageBeginFrameTime)
	io.Float32(&pk.AverageInputTime)
	io.Float32(&pk.AverageRenderTime)
	io.Float32(&pk.AverageEndFrameTime)
	io.Float32(&pk.AverageRemainderTimePercent)
	io.Float32(&pk.AverageUnaccountedTimePercent)
	if proto.IsProtoGTE(io, proto.ID924) {
		protocol.Slice(io, &pk.MemoryCategoryValues)
	}
	if proto.IsProtoGTE(io, proto.ID975) {
		protocol.Slice(io, &pk.EntityDiagnostics)
		protocol.Slice(io, &pk.SystemDiagnostics)
	}
}

// ToLatest ...
func (pk *ServerBoundDiagnostics) ToLatest() *packet.ServerBoundDiagnostics {
	return &packet.ServerBoundDiagnostics{
		AverageFramesPerSecond:        pk.AverageFramesPerSecond,
		AverageServerSimTickTime:      pk.AverageServerSimTickTime,
		AverageClientSimTickTime:      pk.AverageClientSimTickTime,
		AverageBeginFrameTime:         pk.AverageBeginFrameTime,
		AverageInputTime:              pk.AverageInputTime,
		AverageRenderTime:             pk.AverageRenderTime,
		AverageEndFrameTime:           pk.AverageEndFrameTime,
		AverageRemainderTimePercent:   pk.AverageRemainderTimePercent,
		AverageUnaccountedTimePercent: pk.AverageUnaccountedTimePercent,
		MemoryCategoryValues:          pk.MemoryCategoryValues,
		EntityDiagnostics:             pk.EntityDiagnostics,
		SystemDiagnostics:             pk.SystemDiagnostics,
	}
}

// FromLatest ...
func (pk *ServerBoundDiagnostics) FromLatest(latest *packet.ServerBoundDiagnostics) *ServerBoundDiagnostics {
	pk.AverageFramesPerSecond = latest.AverageFramesPerSecond
	pk.AverageServerSimTickTime = latest.AverageServerSimTickTime
	pk.AverageClientSimTickTime = latest.AverageClientSimTickTime
	pk.AverageBeginFrameTime = latest.AverageBeginFrameTime
	pk.AverageInputTime = latest.AverageInputTime
	pk.AverageRenderTime = latest.AverageRenderTime
	pk.AverageEndFrameTime = latest.AverageEndFrameTime
	pk.AverageRemainderTimePercent = latest.AverageRemainderTimePercent
	pk.AverageUnaccountedTimePercent = latest.AverageUnaccountedTimePercent
	pk.MemoryCategoryValues = latest.MemoryCategoryValues
	pk.EntityDiagnostics = latest.EntityDiagnostics
	pk.SystemDiagnostics = latest.SystemDiagnostics
	return pk
}
