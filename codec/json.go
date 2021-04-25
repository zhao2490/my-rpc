package codec

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
)

type JsonCodec struct {
	conn io.ReadWriteCloser
	buf  *bufio.Writer
	enc  *json.Encoder
	dec  *json.Decoder
}

var _ Codec = (*JsonCodec)(nil)

func NewJsonCodec(conn io.ReadWriteCloser) Codec {
	buf := bufio.NewWriter(conn)
	return &JsonCodec{
		conn: conn,
		buf:  buf,
		enc:  json.NewEncoder(buf),
		dec:  json.NewDecoder(conn),
	}
}

func (c *JsonCodec) ReadHeader(header *Header) error {
	return c.enc.Encode(header)
}

func (c *JsonCodec) ReadBody(body interface{}) error {
	return c.enc.Encode(body)
}

func (c *JsonCodec) Write(header *Header, body interface{}) (err error) {
	defer func() {
		_ = c.buf.Flush()
		if err != nil {
			_ = c.Close()
		}
	}()
	if err = c.enc.Encode(header); err != nil {
		log.Println("rpc codec: json error encoding header: ", err)
		return err
	}
	if err = c.enc.Encode(body); err != nil {
		log.Println("rpc codec: json error encoding body: ", err)
		return err
	}
	return nil
}

func (c *JsonCodec) Close() error {
	return c.conn.Close()
}
