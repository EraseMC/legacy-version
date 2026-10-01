package proto

import "github.com/sandertv/gophertunnel/minecraft/protocol"

func UBlockPos(io protocol.IO, x *protocol.BlockPos) {
	if IsWriter(io) {
		io.Varint32(&x[0])
		y := uint32(x[1])
		io.Varuint32(&y)
		io.Varint32(&x[2])
		return
	}
	io.Varint32(&x[0])
	var y uint32
	io.Varuint32(&y)
	x[1] = int32(y)
	io.Varint32(&x[2])
}
