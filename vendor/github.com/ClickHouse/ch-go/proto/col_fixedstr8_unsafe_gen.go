//go:build (amd64 || arm64 || riscv64) && !purego

// Code generated by ./cmd/ch-gen-col, DO NOT EDIT.

package proto

import (
	"unsafe"

	"github.com/go-faster/errors"
)

// DecodeColumn decodes FixedStr8 rows from *Reader.
func (c *ColFixedStr8) DecodeColumn(r *Reader, rows int) error {
	if rows == 0 {
		return nil
	}
	*c = append(*c, make([][8]byte, rows)...)
	s := *(*slice)(unsafe.Pointer(c))
	const size = 8
	s.Len *= size
	s.Cap *= size
	dst := *(*[]byte)(unsafe.Pointer(&s))
	if err := r.ReadFull(dst); err != nil {
		return errors.Wrap(err, "read full")
	}
	return nil
}

// EncodeColumn encodes FixedStr8 rows to *Buffer.
func (c ColFixedStr8) EncodeColumn(b *Buffer) {
	v := c
	if len(v) == 0 {
		return
	}
	offset := len(b.Buf)
	const size = 8
	b.Buf = append(b.Buf, make([]byte, size*len(v))...)
	s := *(*slice)(unsafe.Pointer(&v))
	s.Len *= size
	s.Cap *= size
	src := *(*[]byte)(unsafe.Pointer(&s))
	dst := b.Buf[offset:]
	copy(dst, src)
}
