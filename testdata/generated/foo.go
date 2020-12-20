// Code generated by bebopc-go; DO NOT EDIT.

package generated

import (
	"bytes"
	"io"

	"github.com/200sc/bebop"
	"github.com/200sc/bebop/iohelp"
)

var _ bebop.Record = &Foo{}

type Foo struct {
	Bar Bar
}

func (bbp Foo) MarshalBebop() []byte {
	buf := make([]byte, bbp.bodyLen())
	bbp.MarshalBebopTo(buf)
	return buf
}

func (bbp Foo) MarshalBebopTo(buf []byte) {
	at := 0
	(bbp.Bar).MarshalBebopTo(buf[at:])
	at += (bbp.Bar).bodyLen()
}

func (bbp *Foo) UnmarshalBebop(buf []byte) (err error) {
	at := 0
	bbp.Bar, err = makeBarFromBytes(buf[at:])
	if err != nil {
		 return err
	}
	at += (bbp.Bar).bodyLen()
	return nil
}

func (bbp *Foo) MustUnmarshalBebop(buf []byte) {
	at := 0
	bbp.Bar = mustMakeBarFromBytes(buf[at:])
	at += (bbp.Bar).bodyLen()
}

func (bbp Foo) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	err = (bbp.Bar).EncodeBebop(w)
	if err != nil {
		return err
	}
	return w.Err
}

func (bbp *Foo) DecodeBebop(ior io.Reader) (err error) {
	r := iohelp.NewErrorReader(ior)
	(bbp.Bar), err = makeBar(r)
	if err != nil {
		return err
	}
	return r.Err
}

func (bbp Foo) bodyLen() int {
	bodyLen := 0
	bodyLen += (bbp.Bar).bodyLen()
	return bodyLen
}

func makeFoo(r iohelp.ErrorReader) (Foo, error) {
	v := Foo{}
	err := v.DecodeBebop(r)
	return v, err
}

func makeFooFromBytes(buf []byte) (Foo, error) {
	v := Foo{}
	err := v.UnmarshalBebop(buf)
	return v, err
}

func mustMakeFooFromBytes(buf []byte) Foo {
	v := Foo{}
	v.MustUnmarshalBebop(buf)
	return v
}

var _ bebop.Record = &Bar{}

type Bar struct {
	X *float64
	Y *float64
	Z *float64
}

func (bbp Bar) MarshalBebop() []byte {
	buf := make([]byte, bbp.bodyLen())
	bbp.MarshalBebopTo(buf)
	return buf
}

func (bbp Bar) MarshalBebopTo(buf []byte) {
	at := 0
	iohelp.WriteUint32Bytes(buf[at:], uint32(bbp.bodyLen()-4))
	at += 4
	if bbp.X != nil {
		buf[at] = 1
		at++
		iohelp.WriteFloat64Bytes(buf[at:], *bbp.X)
		at += 8
	}
	if bbp.Y != nil {
		buf[at] = 2
		at++
		iohelp.WriteFloat64Bytes(buf[at:], *bbp.Y)
		at += 8
	}
	if bbp.Z != nil {
		buf[at] = 3
		at++
		iohelp.WriteFloat64Bytes(buf[at:], *bbp.Z)
		at += 8
	}
}

func (bbp *Bar) UnmarshalBebop(buf []byte) (err error) {
	at := 0
	_ = iohelp.ReadUint32Bytes(buf[at:])
	buf = buf[4:]
	for {
		switch buf[at] {
		case 1:
			at += 1
			bbp.X = new(float64)
			if len(buf[at:]) < 8 {
				 return iohelp.ErrTooShort
			}
			(*bbp.X) = iohelp.ReadFloat64Bytes(buf[at:])
			at += 8
		case 2:
			at += 1
			bbp.Y = new(float64)
			if len(buf[at:]) < 8 {
				 return iohelp.ErrTooShort
			}
			(*bbp.Y) = iohelp.ReadFloat64Bytes(buf[at:])
			at += 8
		case 3:
			at += 1
			bbp.Z = new(float64)
			if len(buf[at:]) < 8 {
				 return iohelp.ErrTooShort
			}
			(*bbp.Z) = iohelp.ReadFloat64Bytes(buf[at:])
			at += 8
		default:
			return nil
		}
	}
}

func (bbp *Bar) MustUnmarshalBebop(buf []byte) {
	at := 0
	for {
		switch buf[at] {
		case 1:
			bbp.X = new(float64)
			(*bbp.X) = iohelp.ReadFloat64Bytes(buf[at:])
			at += 8
		case 2:
			bbp.Y = new(float64)
			(*bbp.Y) = iohelp.ReadFloat64Bytes(buf[at:])
			at += 8
		case 3:
			bbp.Z = new(float64)
			(*bbp.Z) = iohelp.ReadFloat64Bytes(buf[at:])
			at += 8
		default:
			return
		}
	}
}

func (bbp Bar) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	iohelp.WriteUint32(w, uint32(bbp.bodyLen()))
	if bbp.X != nil {
		w.Write([]byte{1})
		iohelp.WriteFloat64(w, *bbp.X)
	}
	if bbp.Y != nil {
		w.Write([]byte{2})
		iohelp.WriteFloat64(w, *bbp.Y)
	}
	if bbp.Z != nil {
		w.Write([]byte{3})
		iohelp.WriteFloat64(w, *bbp.Z)
	}
	w.Write([]byte{0})
	return w.Err
}

func (bbp *Bar) DecodeBebop(ior io.Reader) (err error) {
	er := iohelp.NewErrorReader(ior)
	bodyLen := iohelp.ReadUint32(er)
	body := make([]byte, bodyLen)
	er.Read(body)
	r := iohelp.NewErrorReader(bytes.NewReader(body))
	for {
		switch iohelp.ReadByte(r) {
		case 1:
			bbp.X = new(float64)
			*bbp.X = iohelp.ReadFloat64(r)
		case 2:
			bbp.Y = new(float64)
			*bbp.Y = iohelp.ReadFloat64(r)
		case 3:
			bbp.Z = new(float64)
			*bbp.Z = iohelp.ReadFloat64(r)
		default:
			return er.Err
		}
	}
}

func (bbp Bar) bodyLen() int {
	bodyLen := 5
	if bbp.X != nil {
		bodyLen += 1
		bodyLen += 8
	}
	if bbp.Y != nil {
		bodyLen += 1
		bodyLen += 8
	}
	if bbp.Z != nil {
		bodyLen += 1
		bodyLen += 8
	}
	return bodyLen
}

func makeBar(r iohelp.ErrorReader) (Bar, error) {
	v := Bar{}
	err := v.DecodeBebop(r)
	return v, err
}

func makeBarFromBytes(buf []byte) (Bar, error) {
	v := Bar{}
	err := v.UnmarshalBebop(buf)
	return v, err
}

func mustMakeBarFromBytes(buf []byte) Bar {
	v := Bar{}
	v.MustUnmarshalBebop(buf)
	return v
}

