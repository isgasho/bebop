// Code generated by bebopc-go; DO NOT EDIT.

package generated

import (
	"bytes"
	"io"

	"github.com/200sc/bebop"
	"github.com/200sc/bebop/iohelp"
)

type DepE uint32

const (
	// Deprecated: X in DepE
	DepE_X DepE = 1
)

// Documented enum 
type DocE uint32

const (
	// Documented constant 
	DocE_X DocE = 1
	// Deprecated: Y in DocE
	DocE_Y DocE = 2
	// Deprecated, documented constant 
	// Deprecated: Z in DocE
	DocE_Z DocE = 3
)

var _ bebop.Record = &DocS{}

// Documented struct 
type DocS struct {
	// Documented field 
	X int32
}

func (bbp DocS) MarshalBebop() []byte {
	buf := make([]byte, bbp.bodyLen())
	bbp.MarshalBebopTo(buf)
	return buf
}

func (bbp DocS) MarshalBebopTo(buf []byte) {
	at := 0
	iohelp.WriteInt32Bytes(buf[at:], bbp.X)
	at += 4
}

func (bbp *DocS) UnmarshalBebop(buf []byte) (err error) {
	at := 0
	if len(buf[at:]) < 4 {
		 return iohelp.ErrTooShort
	}
	bbp.X = iohelp.ReadInt32Bytes(buf[at:])
	at += 4
	return nil
}

func (bbp *DocS) MustUnmarshalBebop(buf []byte) {
	at := 0
	bbp.X = iohelp.ReadInt32Bytes(buf[at:])
	at += 4
}

func (bbp DocS) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	iohelp.WriteInt32(w, bbp.X)
	return w.Err
}

func (bbp *DocS) DecodeBebop(ior io.Reader) (err error) {
	r := iohelp.NewErrorReader(ior)
	bbp.X = iohelp.ReadInt32(r)
	return r.Err
}

func (bbp DocS) bodyLen() int {
	bodyLen := 0
	bodyLen += 4
	return bodyLen
}

func makeDocS(r iohelp.ErrorReader) (DocS, error) {
	v := DocS{}
	err := v.DecodeBebop(r)
	return v, err
}

func makeDocSFromBytes(buf []byte) (DocS, error) {
	v := DocS{}
	err := v.UnmarshalBebop(buf)
	return v, err
}

func mustMakeDocSFromBytes(buf []byte) DocS {
	v := DocS{}
	v.MustUnmarshalBebop(buf)
	return v
}

var _ bebop.Record = &DepM{}

type DepM struct {
	// Deprecated: x in DepM
	X *int32
}

func (bbp DepM) MarshalBebop() []byte {
	buf := make([]byte, bbp.bodyLen())
	bbp.MarshalBebopTo(buf)
	return buf
}

func (bbp DepM) MarshalBebopTo(buf []byte) {
	at := 0
	iohelp.WriteUint32Bytes(buf[at:], uint32(bbp.bodyLen()-4))
	at += 4
	if bbp.X != nil {
		buf[at] = 1
		at++
		iohelp.WriteInt32Bytes(buf[at:], *bbp.X)
		at += 4
	}
}

func (bbp *DepM) UnmarshalBebop(buf []byte) (err error) {
	at := 0
	_ = iohelp.ReadUint32Bytes(buf[at:])
	buf = buf[4:]
	for {
		switch buf[at] {
		case 1:
			at += 1
			bbp.X = new(int32)
			if len(buf[at:]) < 4 {
				 return iohelp.ErrTooShort
			}
			(*bbp.X) = iohelp.ReadInt32Bytes(buf[at:])
			at += 4
		default:
			return nil
		}
	}
}

func (bbp *DepM) MustUnmarshalBebop(buf []byte) {
	at := 0
	for {
		switch buf[at] {
		case 1:
			bbp.X = new(int32)
			(*bbp.X) = iohelp.ReadInt32Bytes(buf[at:])
			at += 4
		default:
			return
		}
	}
}

func (bbp DepM) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	iohelp.WriteUint32(w, uint32(bbp.bodyLen()))
	if bbp.X != nil {
		w.Write([]byte{1})
		iohelp.WriteInt32(w, *bbp.X)
	}
	w.Write([]byte{0})
	return w.Err
}

func (bbp *DepM) DecodeBebop(ior io.Reader) (err error) {
	er := iohelp.NewErrorReader(ior)
	bodyLen := iohelp.ReadUint32(er)
	body := make([]byte, bodyLen)
	er.Read(body)
	r := iohelp.NewErrorReader(bytes.NewReader(body))
	for {
		switch iohelp.ReadByte(r) {
		case 1:
			bbp.X = new(int32)
			*bbp.X = iohelp.ReadInt32(r)
		default:
			return er.Err
		}
	}
}

func (bbp DepM) bodyLen() int {
	bodyLen := 5
	if bbp.X != nil {
		bodyLen += 1
		bodyLen += 4
	}
	return bodyLen
}

func makeDepM(r iohelp.ErrorReader) (DepM, error) {
	v := DepM{}
	err := v.DecodeBebop(r)
	return v, err
}

func makeDepMFromBytes(buf []byte) (DepM, error) {
	v := DepM{}
	err := v.UnmarshalBebop(buf)
	return v, err
}

func mustMakeDepMFromBytes(buf []byte) DepM {
	v := DepM{}
	v.MustUnmarshalBebop(buf)
	return v
}

var _ bebop.Record = &DocM{}

// Documented message 
type DocM struct {
	// Documented field 
	X *int32
	// Deprecated: y in DocM
	Y *int32
	// Deprecated, documented field 
	// Deprecated: z in DocM
	Z *int32
}

func (bbp DocM) MarshalBebop() []byte {
	buf := make([]byte, bbp.bodyLen())
	bbp.MarshalBebopTo(buf)
	return buf
}

func (bbp DocM) MarshalBebopTo(buf []byte) {
	at := 0
	iohelp.WriteUint32Bytes(buf[at:], uint32(bbp.bodyLen()-4))
	at += 4
	if bbp.X != nil {
		buf[at] = 1
		at++
		iohelp.WriteInt32Bytes(buf[at:], *bbp.X)
		at += 4
	}
	if bbp.Y != nil {
		buf[at] = 2
		at++
		iohelp.WriteInt32Bytes(buf[at:], *bbp.Y)
		at += 4
	}
	if bbp.Z != nil {
		buf[at] = 3
		at++
		iohelp.WriteInt32Bytes(buf[at:], *bbp.Z)
		at += 4
	}
}

func (bbp *DocM) UnmarshalBebop(buf []byte) (err error) {
	at := 0
	_ = iohelp.ReadUint32Bytes(buf[at:])
	buf = buf[4:]
	for {
		switch buf[at] {
		case 1:
			at += 1
			bbp.X = new(int32)
			if len(buf[at:]) < 4 {
				 return iohelp.ErrTooShort
			}
			(*bbp.X) = iohelp.ReadInt32Bytes(buf[at:])
			at += 4
		case 2:
			at += 1
			bbp.Y = new(int32)
			if len(buf[at:]) < 4 {
				 return iohelp.ErrTooShort
			}
			(*bbp.Y) = iohelp.ReadInt32Bytes(buf[at:])
			at += 4
		case 3:
			at += 1
			bbp.Z = new(int32)
			if len(buf[at:]) < 4 {
				 return iohelp.ErrTooShort
			}
			(*bbp.Z) = iohelp.ReadInt32Bytes(buf[at:])
			at += 4
		default:
			return nil
		}
	}
}

func (bbp *DocM) MustUnmarshalBebop(buf []byte) {
	at := 0
	for {
		switch buf[at] {
		case 1:
			bbp.X = new(int32)
			(*bbp.X) = iohelp.ReadInt32Bytes(buf[at:])
			at += 4
		case 2:
			bbp.Y = new(int32)
			(*bbp.Y) = iohelp.ReadInt32Bytes(buf[at:])
			at += 4
		case 3:
			bbp.Z = new(int32)
			(*bbp.Z) = iohelp.ReadInt32Bytes(buf[at:])
			at += 4
		default:
			return
		}
	}
}

func (bbp DocM) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	iohelp.WriteUint32(w, uint32(bbp.bodyLen()))
	if bbp.X != nil {
		w.Write([]byte{1})
		iohelp.WriteInt32(w, *bbp.X)
	}
	if bbp.Y != nil {
		w.Write([]byte{2})
		iohelp.WriteInt32(w, *bbp.Y)
	}
	if bbp.Z != nil {
		w.Write([]byte{3})
		iohelp.WriteInt32(w, *bbp.Z)
	}
	w.Write([]byte{0})
	return w.Err
}

func (bbp *DocM) DecodeBebop(ior io.Reader) (err error) {
	er := iohelp.NewErrorReader(ior)
	bodyLen := iohelp.ReadUint32(er)
	body := make([]byte, bodyLen)
	er.Read(body)
	r := iohelp.NewErrorReader(bytes.NewReader(body))
	for {
		switch iohelp.ReadByte(r) {
		case 1:
			bbp.X = new(int32)
			*bbp.X = iohelp.ReadInt32(r)
		case 2:
			bbp.Y = new(int32)
			*bbp.Y = iohelp.ReadInt32(r)
		case 3:
			bbp.Z = new(int32)
			*bbp.Z = iohelp.ReadInt32(r)
		default:
			return er.Err
		}
	}
}

func (bbp DocM) bodyLen() int {
	bodyLen := 5
	if bbp.X != nil {
		bodyLen += 1
		bodyLen += 4
	}
	if bbp.Y != nil {
		bodyLen += 1
		bodyLen += 4
	}
	if bbp.Z != nil {
		bodyLen += 1
		bodyLen += 4
	}
	return bodyLen
}

func makeDocM(r iohelp.ErrorReader) (DocM, error) {
	v := DocM{}
	err := v.DecodeBebop(r)
	return v, err
}

func makeDocMFromBytes(buf []byte) (DocM, error) {
	v := DocM{}
	err := v.UnmarshalBebop(buf)
	return v, err
}

func mustMakeDocMFromBytes(buf []byte) DocM {
	v := DocM{}
	v.MustUnmarshalBebop(buf)
	return v
}

