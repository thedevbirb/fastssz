// Code generated by fastssz. DO NOT EDIT.
// Hash: 7d191912ad4d9d8f069f503d82d98d5011cfa1e35dfe0a2e7378082bb6ae8b1d
// Version: 0.1.4
package testcases

import (
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the BytesWrapper object
func (b *BytesWrapper) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BytesWrapper object to a target array
func (b *BytesWrapper) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'Bytes'
	if size := len(b.Bytes); size != 48 {
		err = ssz.ErrBytesLengthFn("BytesWrapper.Bytes", size, 48)
		return
	}
	dst = append(dst, b.Bytes...)

	return
}

// UnmarshalSSZ ssz unmarshals the BytesWrapper object
func (b *BytesWrapper) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 48 {
		return ssz.ErrSize
	}

	// Field (0) 'Bytes'
	if cap(b.Bytes) == 0 {
		b.Bytes = make([]byte, 0, len(buf[0:48]))
	}
	b.Bytes = append(b.Bytes, buf[0:48]...)

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BytesWrapper object
func (b *BytesWrapper) SizeSSZ() (size int) {
	size = 48
	return
}

// HashTreeRoot ssz hashes the BytesWrapper object
func (b *BytesWrapper) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BytesWrapper object with a hasher
func (b *BytesWrapper) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Bytes'
	if size := len(b.Bytes); size != 48 {
		err = ssz.ErrBytesLengthFn("BytesWrapper.Bytes", size, 48)
		return
	}
	hh.PutBytes(b.Bytes)

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the BytesWrapper object
func (b *BytesWrapper) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(b)
}

// MarshalSSZ ssz marshals the ListC object
func (l *ListC) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(l)
}

// MarshalSSZTo ssz marshals the ListC object to a target array
func (l *ListC) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(4)

	// Offset (0) 'Elems'
	dst = ssz.WriteOffset(dst, offset)

	// Field (0) 'Elems'
	if size := len(l.Elems); size > 32 {
		err = ssz.ErrListTooBigFn("ListC.Elems", size, 32)
		return
	}
	for ii := 0; ii < len(l.Elems); ii++ {
		if dst, err = l.Elems[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	return
}

// UnmarshalSSZ ssz unmarshals the ListC object
func (l *ListC) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 4 {
		return ssz.ErrSize
	}

	tail := buf
	var o0 uint64

	// Offset (0) 'Elems'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 != 4 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (0) 'Elems'
	{
		buf = tail[o0:]
		num, err := ssz.DivideInt2(len(buf), 48, 32)
		if err != nil {
			return err
		}
		l.Elems = make([]BytesWrapper, num)
		for ii := 0; ii < num; ii++ {
			if err = l.Elems[ii].UnmarshalSSZ(buf[ii*48 : (ii+1)*48]); err != nil {
				return err
			}
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the ListC object
func (l *ListC) SizeSSZ() (size int) {
	size = 4

	// Field (0) 'Elems'
	size += len(l.Elems) * 48

	return
}

// HashTreeRoot ssz hashes the ListC object
func (l *ListC) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(l)
}

// HashTreeRootWith ssz hashes the ListC object with a hasher
func (l *ListC) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Elems'
	{
		subIndx := hh.Index()
		num := uint64(len(l.Elems))
		if num > 32 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range l.Elems {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 32)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the ListC object
func (l *ListC) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(l)
}

// MarshalSSZ ssz marshals the ListP object
func (l *ListP) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(l)
}

// MarshalSSZTo ssz marshals the ListP object to a target array
func (l *ListP) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(4)

	// Offset (0) 'Elems'
	dst = ssz.WriteOffset(dst, offset)

	// Field (0) 'Elems'
	if size := len(l.Elems); size > 32 {
		err = ssz.ErrListTooBigFn("ListP.Elems", size, 32)
		return
	}
	for ii := 0; ii < len(l.Elems); ii++ {
		if dst, err = l.Elems[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	return
}

// UnmarshalSSZ ssz unmarshals the ListP object
func (l *ListP) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 4 {
		return ssz.ErrSize
	}

	tail := buf
	var o0 uint64

	// Offset (0) 'Elems'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 != 4 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (0) 'Elems'
	{
		buf = tail[o0:]
		num, err := ssz.DivideInt2(len(buf), 48, 32)
		if err != nil {
			return err
		}
		l.Elems = make([]*BytesWrapper, num)
		for ii := 0; ii < num; ii++ {
			if l.Elems[ii] == nil {
				l.Elems[ii] = new(BytesWrapper)
			}
			if err = l.Elems[ii].UnmarshalSSZ(buf[ii*48 : (ii+1)*48]); err != nil {
				return err
			}
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the ListP object
func (l *ListP) SizeSSZ() (size int) {
	size = 4

	// Field (0) 'Elems'
	size += len(l.Elems) * 48

	return
}

// HashTreeRoot ssz hashes the ListP object
func (l *ListP) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(l)
}

// HashTreeRootWith ssz hashes the ListP object with a hasher
func (l *ListP) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Elems'
	{
		subIndx := hh.Index()
		num := uint64(len(l.Elems))
		if num > 32 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range l.Elems {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 32)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the ListP object
func (l *ListP) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(l)
}
