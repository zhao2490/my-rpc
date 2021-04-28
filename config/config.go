package config

import (
	"github.com/zhao2490/my-rpc/codec"
	"time"
)

const MagicNumber = 0x3bef5c

type Option struct {
	MagicNumber    int        // MagicNumber marks this's a my-rpc request
	CodecType      codec.Type // client may choose different Codec to encode body
	ConnectTimeout time.Duration
	HandleTimeout  time.Duration
}

var DefaultOption = &Option{
	MagicNumber:    MagicNumber,
	CodecType:      codec.GobType,
	ConnectTimeout: time.Second * 10,
}
