// Code generated by fastssz. DO NOT EDIT.
// Hash: b39bb565b34add934304fdd716d14cb913acca74f226a8fd1f0c0c390e3bc247
package sync_pb

import (
	types_pb "github.com/astranetworld/ast/api/protocol/types_pb"
	ssz "github.com/prysmaticlabs/fastssz"
)

// MarshalSSZ ssz marshals the HeadersByRangeRequest object
func (h *HeadersByRangeRequest) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(h)
}

// MarshalSSZTo ssz marshals the HeadersByRangeRequest object to a target array
func (h *HeadersByRangeRequest) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(20)

	// Offset (0) 'StartBlockNumber'
	dst = ssz.WriteOffset(dst, offset)
	if h.StartBlockNumber == nil {
		h.StartBlockNumber = new(types_pb.H256)
	}
	offset += h.StartBlockNumber.SizeSSZ()

	// Field (1) 'Count'
	dst = ssz.MarshalUint64(dst, h.Count)

	// Field (2) 'Step'
	dst = ssz.MarshalUint64(dst, h.Step)

	// Field (0) 'StartBlockNumber'
	if dst, err = h.StartBlockNumber.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the HeadersByRangeRequest object
func (h *HeadersByRangeRequest) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 20 {
		return ssz.ErrSize
	}

	tail := buf
	var o0 uint64

	// Offset (0) 'StartBlockNumber'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 20 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (1) 'Count'
	h.Count = ssz.UnmarshallUint64(buf[4:12])

	// Field (2) 'Step'
	h.Step = ssz.UnmarshallUint64(buf[12:20])

	// Field (0) 'StartBlockNumber'
	{
		buf = tail[o0:]
		if h.StartBlockNumber == nil {
			h.StartBlockNumber = new(types_pb.H256)
		}
		if err = h.StartBlockNumber.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the HeadersByRangeRequest object
func (h *HeadersByRangeRequest) SizeSSZ() (size int) {
	size = 20

	// Field (0) 'StartBlockNumber'
	if h.StartBlockNumber == nil {
		h.StartBlockNumber = new(types_pb.H256)
	}
	size += h.StartBlockNumber.SizeSSZ()

	return
}

// HashTreeRoot ssz hashes the HeadersByRangeRequest object
func (h *HeadersByRangeRequest) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(h)
}

// HashTreeRootWith ssz hashes the HeadersByRangeRequest object with a hasher
func (h *HeadersByRangeRequest) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'StartBlockNumber'
	if err = h.StartBlockNumber.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'Count'
	hh.PutUint64(h.Count)

	// Field (2) 'Step'
	hh.PutUint64(h.Step)

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}

// MarshalSSZ ssz marshals the Ping object
func (p *Ping) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(p)
}

// MarshalSSZTo ssz marshals the Ping object to a target array
func (p *Ping) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'SeqNumber'
	dst = ssz.MarshalUint64(dst, p.SeqNumber)

	return
}

// UnmarshalSSZ ssz unmarshals the Ping object
func (p *Ping) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 8 {
		return ssz.ErrSize
	}

	// Field (0) 'SeqNumber'
	p.SeqNumber = ssz.UnmarshallUint64(buf[0:8])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Ping object
func (p *Ping) SizeSSZ() (size int) {
	size = 8
	return
}

// HashTreeRoot ssz hashes the Ping object
func (p *Ping) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(p)
}

// HashTreeRootWith ssz hashes the Ping object with a hasher
func (p *Ping) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'SeqNumber'
	hh.PutUint64(p.SeqNumber)

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}

// MarshalSSZ ssz marshals the Status object
func (s *Status) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(s)
}

// MarshalSSZTo ssz marshals the Status object to a target array
func (s *Status) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(8)

	// Offset (0) 'GenesisHash'
	dst = ssz.WriteOffset(dst, offset)
	if s.GenesisHash == nil {
		s.GenesisHash = new(types_pb.H256)
	}
	offset += s.GenesisHash.SizeSSZ()

	// Offset (1) 'CurrentHeight'
	dst = ssz.WriteOffset(dst, offset)
	if s.CurrentHeight == nil {
		s.CurrentHeight = new(types_pb.H256)
	}
	offset += s.CurrentHeight.SizeSSZ()

	// Field (0) 'GenesisHash'
	if dst, err = s.GenesisHash.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (1) 'CurrentHeight'
	if dst, err = s.CurrentHeight.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the Status object
func (s *Status) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 8 {
		return ssz.ErrSize
	}

	tail := buf
	var o0, o1 uint64

	// Offset (0) 'GenesisHash'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 8 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (1) 'CurrentHeight'
	if o1 = ssz.ReadOffset(buf[4:8]); o1 > size || o0 > o1 {
		return ssz.ErrOffset
	}

	// Field (0) 'GenesisHash'
	{
		buf = tail[o0:o1]
		if s.GenesisHash == nil {
			s.GenesisHash = new(types_pb.H256)
		}
		if err = s.GenesisHash.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (1) 'CurrentHeight'
	{
		buf = tail[o1:]
		if s.CurrentHeight == nil {
			s.CurrentHeight = new(types_pb.H256)
		}
		if err = s.CurrentHeight.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the Status object
func (s *Status) SizeSSZ() (size int) {
	size = 8

	// Field (0) 'GenesisHash'
	if s.GenesisHash == nil {
		s.GenesisHash = new(types_pb.H256)
	}
	size += s.GenesisHash.SizeSSZ()

	// Field (1) 'CurrentHeight'
	if s.CurrentHeight == nil {
		s.CurrentHeight = new(types_pb.H256)
	}
	size += s.CurrentHeight.SizeSSZ()

	return
}

// HashTreeRoot ssz hashes the Status object
func (s *Status) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(s)
}

// HashTreeRootWith ssz hashes the Status object with a hasher
func (s *Status) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'GenesisHash'
	if err = s.GenesisHash.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'CurrentHeight'
	if err = s.CurrentHeight.HashTreeRootWith(hh); err != nil {
		return
	}

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}

// MarshalSSZ ssz marshals the ForkData object
func (f *ForkData) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(f)
}

// MarshalSSZTo ssz marshals the ForkData object to a target array
func (f *ForkData) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(8)

	// Offset (0) 'CurrentVersion'
	dst = ssz.WriteOffset(dst, offset)
	if f.CurrentVersion == nil {
		f.CurrentVersion = new(types_pb.H256)
	}
	offset += f.CurrentVersion.SizeSSZ()

	// Offset (1) 'GenesisValidatorsRoot'
	dst = ssz.WriteOffset(dst, offset)
	if f.GenesisValidatorsRoot == nil {
		f.GenesisValidatorsRoot = new(types_pb.H256)
	}
	offset += f.GenesisValidatorsRoot.SizeSSZ()

	// Field (0) 'CurrentVersion'
	if dst, err = f.CurrentVersion.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (1) 'GenesisValidatorsRoot'
	if dst, err = f.GenesisValidatorsRoot.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the ForkData object
func (f *ForkData) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 8 {
		return ssz.ErrSize
	}

	tail := buf
	var o0, o1 uint64

	// Offset (0) 'CurrentVersion'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 8 {
		return ssz.ErrInvalidVariableOffset
	}

	// Offset (1) 'GenesisValidatorsRoot'
	if o1 = ssz.ReadOffset(buf[4:8]); o1 > size || o0 > o1 {
		return ssz.ErrOffset
	}

	// Field (0) 'CurrentVersion'
	{
		buf = tail[o0:o1]
		if f.CurrentVersion == nil {
			f.CurrentVersion = new(types_pb.H256)
		}
		if err = f.CurrentVersion.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}

	// Field (1) 'GenesisValidatorsRoot'
	{
		buf = tail[o1:]
		if f.GenesisValidatorsRoot == nil {
			f.GenesisValidatorsRoot = new(types_pb.H256)
		}
		if err = f.GenesisValidatorsRoot.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the ForkData object
func (f *ForkData) SizeSSZ() (size int) {
	size = 8

	// Field (0) 'CurrentVersion'
	if f.CurrentVersion == nil {
		f.CurrentVersion = new(types_pb.H256)
	}
	size += f.CurrentVersion.SizeSSZ()

	// Field (1) 'GenesisValidatorsRoot'
	if f.GenesisValidatorsRoot == nil {
		f.GenesisValidatorsRoot = new(types_pb.H256)
	}
	size += f.GenesisValidatorsRoot.SizeSSZ()

	return
}

// HashTreeRoot ssz hashes the ForkData object
func (f *ForkData) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(f)
}

// HashTreeRootWith ssz hashes the ForkData object with a hasher
func (f *ForkData) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'CurrentVersion'
	if err = f.CurrentVersion.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'GenesisValidatorsRoot'
	if err = f.GenesisValidatorsRoot.HashTreeRootWith(hh); err != nil {
		return
	}

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}

// MarshalSSZ ssz marshals the BodiesByRangeRequest object
func (b *BodiesByRangeRequest) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BodiesByRangeRequest object to a target array
func (b *BodiesByRangeRequest) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(20)

	// Offset (0) 'StartBlockNumber'
	dst = ssz.WriteOffset(dst, offset)
	if b.StartBlockNumber == nil {
		b.StartBlockNumber = new(types_pb.H256)
	}
	offset += b.StartBlockNumber.SizeSSZ()

	// Field (1) 'Count'
	dst = ssz.MarshalUint64(dst, b.Count)

	// Field (2) 'Step'
	dst = ssz.MarshalUint64(dst, b.Step)

	// Field (0) 'StartBlockNumber'
	if dst, err = b.StartBlockNumber.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the BodiesByRangeRequest object
func (b *BodiesByRangeRequest) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 20 {
		return ssz.ErrSize
	}

	tail := buf
	var o0 uint64

	// Offset (0) 'StartBlockNumber'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 20 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (1) 'Count'
	b.Count = ssz.UnmarshallUint64(buf[4:12])

	// Field (2) 'Step'
	b.Step = ssz.UnmarshallUint64(buf[12:20])

	// Field (0) 'StartBlockNumber'
	{
		buf = tail[o0:]
		if b.StartBlockNumber == nil {
			b.StartBlockNumber = new(types_pb.H256)
		}
		if err = b.StartBlockNumber.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BodiesByRangeRequest object
func (b *BodiesByRangeRequest) SizeSSZ() (size int) {
	size = 20

	// Field (0) 'StartBlockNumber'
	if b.StartBlockNumber == nil {
		b.StartBlockNumber = new(types_pb.H256)
	}
	size += b.StartBlockNumber.SizeSSZ()

	return
}

// HashTreeRoot ssz hashes the BodiesByRangeRequest object
func (b *BodiesByRangeRequest) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BodiesByRangeRequest object with a hasher
func (b *BodiesByRangeRequest) HashTreeRootWith(hh *ssz.Hasher) (err error) {
	indx := hh.Index()

	// Field (0) 'StartBlockNumber'
	if err = b.StartBlockNumber.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (1) 'Count'
	hh.PutUint64(b.Count)

	// Field (2) 'Step'
	hh.PutUint64(b.Step)

	if ssz.EnableVectorizedHTR {
		hh.MerkleizeVectorizedHTR(indx)
	} else {
		hh.Merkleize(indx)
	}
	return
}
