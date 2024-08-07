// Code generated by fastssz. DO NOT EDIT.
// Hash: 849887b610983adf4efa1abfd912b25aa4f0ffbde8211f1e3909136539b1c2d9
// Version: 0.1.4
package testcases

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the Case1A object
func (c *Case1A) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(c)
}

// MarshalSSZTo ssz marshals the Case1A object to a target array
func (c *Case1A) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(4)

	// Offset (0) 'Foo'
	dst = ssz.WriteOffset(dst, offset)

	// Field (0) 'Foo'
	if size := len(c.Foo); size > 2048 {
		err = ssz.ErrBytesLengthFn("Case1A.Foo", size, 2048)
		return
	}
	dst = append(dst, c.Foo...)

	return
}

// UnmarshalSSZ ssz unmarshals the Case1A object
func (c *Case1A) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 4 {
		return ssz.ErrSize
	}

	tail := buf
	var o0 uint64

	// Offset (0) 'Foo'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 != 4 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (0) 'Foo'
	{
		buf = tail[o0:]
		if len(buf) > 2048 {
			return ssz.ErrBytesLength
		}
		if cap(c.Foo) == 0 {
			c.Foo = make([]byte, 0, len(buf))
		}
		c.Foo = append(c.Foo, buf...)
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Case1A object
func (c *Case1A) SizeSSZ() (size int) {
	size = 4

	// Field (0) 'Foo'
	size += len(c.Foo)

	return
}

// HashTreeRoot ssz hashes the Case1A object
func (c *Case1A) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(c)
}

// HashTreeRootWith ssz hashes the Case1A object with a hasher
func (c *Case1A) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Foo'
	{
		elemIndx := hh.Index()
		byteLen := uint64(len(c.Foo))
		if byteLen > 2048 {
			err = ssz.ErrIncorrectListSize
			return
		}
		hh.Append(c.Foo)
		hh.MerkleizeWithMixin(elemIndx, byteLen, (2048+31)/32)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Case1A object
func (c *Case1A) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(c)
}

// MarshalSSZ ssz marshals the Case1B object
func (c *Case1B) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(c)
}

// MarshalSSZTo ssz marshals the Case1B object to a target array
func (c *Case1B) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(4)

	// Offset (0) 'Bar'
	dst = ssz.WriteOffset(dst, offset)

	// Field (0) 'Bar'
	if size := len(c.Bar); size > 32 {
		err = ssz.ErrBytesLengthFn("Case1B.Bar", size, 32)
		return
	}
	dst = append(dst, c.Bar...)

	return
}

// UnmarshalSSZ ssz unmarshals the Case1B object
func (c *Case1B) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 4 {
		return ssz.ErrSize
	}

	tail := buf
	var o0 uint64

	// Offset (0) 'Bar'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 != 4 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (0) 'Bar'
	{
		buf = tail[o0:]
		if len(buf) > 32 {
			return ssz.ErrBytesLength
		}
		if cap(c.Bar) == 0 {
			c.Bar = make([]byte, 0, len(buf))
		}
		c.Bar = append(c.Bar, buf...)
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Case1B object
func (c *Case1B) SizeSSZ() (size int) {
	size = 4

	// Field (0) 'Bar'
	size += len(c.Bar)

	return
}

// HashTreeRoot ssz hashes the Case1B object
func (c *Case1B) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(c)
}

// HashTreeRootWith ssz hashes the Case1B object with a hasher
func (c *Case1B) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Bar'
	{
		elemIndx := hh.Index()
		byteLen := uint64(len(c.Bar))
		if byteLen > 32 {
			err = ssz.ErrIncorrectListSize
			return
		}
		hh.Append(c.Bar)
		hh.MerkleizeWithMixin(elemIndx, byteLen, (32+31)/32)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the Case1B object
func (c *Case1B) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(c)
}
