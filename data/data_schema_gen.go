package data

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *AuthorAffiliationPair) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0001}
		return
	}
	z.AuthorIndex, err = dc.ReadInt()
	if err != nil {
		err = msgp.WrapError(err, "AuthorIndex")
		return
	}
	z.AffiliationIndex, err = dc.ReadInt()
	if err != nil {
		err = msgp.WrapError(err, "AffiliationIndex")
		return
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z AuthorAffiliationPair) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 2
	err = en.Append(0x92)
	if err != nil {
		return
	}
	err = en.WriteInt(z.AuthorIndex)
	if err != nil {
		err = msgp.WrapError(err, "AuthorIndex")
		return
	}
	err = en.WriteInt(z.AffiliationIndex)
	if err != nil {
		err = msgp.WrapError(err, "AffiliationIndex")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z AuthorAffiliationPair) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 2
	o = append(o, 0x92)
	o = msgp.AppendInt(o, z.AuthorIndex)
	o = msgp.AppendInt(o, z.AffiliationIndex)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *AuthorAffiliationPair) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 2 {
		err = msgp.ArrayError{Wanted: 2, Got: zb0001}
		return
	}
	z.AuthorIndex, bts, err = msgp.ReadIntBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "AuthorIndex")
		return
	}
	z.AffiliationIndex, bts, err = msgp.ReadIntBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "AffiliationIndex")
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z AuthorAffiliationPair) Msgsize() (s int) {
	s = 1 + msgp.IntSize + msgp.IntSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ID) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 uint32
		zb0001, err = dc.ReadUint32()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = ID(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z ID) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteUint32(uint32(z))
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z ID) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendUint32(o, uint32(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ID) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 uint32
		zb0001, bts, err = msgp.ReadUint32Bytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = ID(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z ID) Msgsize() (s int) {
	s = msgp.Uint32Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *IDList) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(IDList, zb0002)
	}
	for zb0001 := range *z {
		{
			var zb0003 uint32
			zb0003, err = dc.ReadUint32()
			if err != nil {
				err = msgp.WrapError(err, zb0001)
				return
			}
			(*z)[zb0001] = ID(zb0003)
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z IDList) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteArrayHeader(uint32(len(z)))
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0004 := range z {
		err = en.WriteUint32(uint32(z[zb0004]))
		if err != nil {
			err = msgp.WrapError(err, zb0004)
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z IDList) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, uint32(len(z)))
	for zb0004 := range z {
		o = msgp.AppendUint32(o, uint32(z[zb0004]))
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *IDList) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(IDList, zb0002)
	}
	for zb0001 := range *z {
		{
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, zb0001)
				return
			}
			(*z)[zb0001] = ID(zb0003)
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z IDList) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize + (len(z) * (msgp.Uint32Size))
	return
}

// DecodeMsg implements msgp.Decodable
func (z *PaperInfo) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 5 {
		err = msgp.ArrayError{Wanted: 5, Got: zb0001}
		return
	}
	z.Year, err = dc.ReadInt()
	if err != nil {
		err = msgp.WrapError(err, "Year")
		return
	}
	z.VenueIndex, err = dc.ReadInt()
	if err != nil {
		err = msgp.WrapError(err, "VenueIndex")
		return
	}
	z.Citation, err = dc.ReadInt()
	if err != nil {
		err = msgp.WrapError(err, "Citation")
		return
	}
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err, "RefList")
		return
	}
	if cap(z.RefList) >= int(zb0002) {
		z.RefList = (z.RefList)[:zb0002]
	} else {
		z.RefList = make([]int, zb0002)
	}
	for za0001 := range z.RefList {
		z.RefList[za0001], err = dc.ReadInt()
		if err != nil {
			err = msgp.WrapError(err, "RefList", za0001)
			return
		}
	}
	var zb0003 uint32
	zb0003, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err, "AuthorAffList")
		return
	}
	if cap(z.AuthorAffList) >= int(zb0003) {
		z.AuthorAffList = (z.AuthorAffList)[:zb0003]
	} else {
		z.AuthorAffList = make([]AuthorAffiliationPair, zb0003)
	}
	for za0002 := range z.AuthorAffList {
		var zb0004 uint32
		zb0004, err = dc.ReadArrayHeader()
		if err != nil {
			err = msgp.WrapError(err, "AuthorAffList", za0002)
			return
		}
		if zb0004 != 2 {
			err = msgp.ArrayError{Wanted: 2, Got: zb0004}
			return
		}
		z.AuthorAffList[za0002].AuthorIndex, err = dc.ReadInt()
		if err != nil {
			err = msgp.WrapError(err, "AuthorAffList", za0002, "AuthorIndex")
			return
		}
		z.AuthorAffList[za0002].AffiliationIndex, err = dc.ReadInt()
		if err != nil {
			err = msgp.WrapError(err, "AuthorAffList", za0002, "AffiliationIndex")
			return
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *PaperInfo) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 5
	err = en.Append(0x95)
	if err != nil {
		return
	}
	err = en.WriteInt(z.Year)
	if err != nil {
		err = msgp.WrapError(err, "Year")
		return
	}
	err = en.WriteInt(z.VenueIndex)
	if err != nil {
		err = msgp.WrapError(err, "VenueIndex")
		return
	}
	err = en.WriteInt(z.Citation)
	if err != nil {
		err = msgp.WrapError(err, "Citation")
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.RefList)))
	if err != nil {
		err = msgp.WrapError(err, "RefList")
		return
	}
	for za0001 := range z.RefList {
		err = en.WriteInt(z.RefList[za0001])
		if err != nil {
			err = msgp.WrapError(err, "RefList", za0001)
			return
		}
	}
	err = en.WriteArrayHeader(uint32(len(z.AuthorAffList)))
	if err != nil {
		err = msgp.WrapError(err, "AuthorAffList")
		return
	}
	for za0002 := range z.AuthorAffList {
		// array header, size 2
		err = en.Append(0x92)
		if err != nil {
			return
		}
		err = en.WriteInt(z.AuthorAffList[za0002].AuthorIndex)
		if err != nil {
			err = msgp.WrapError(err, "AuthorAffList", za0002, "AuthorIndex")
			return
		}
		err = en.WriteInt(z.AuthorAffList[za0002].AffiliationIndex)
		if err != nil {
			err = msgp.WrapError(err, "AuthorAffList", za0002, "AffiliationIndex")
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *PaperInfo) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 5
	o = append(o, 0x95)
	o = msgp.AppendInt(o, z.Year)
	o = msgp.AppendInt(o, z.VenueIndex)
	o = msgp.AppendInt(o, z.Citation)
	o = msgp.AppendArrayHeader(o, uint32(len(z.RefList)))
	for za0001 := range z.RefList {
		o = msgp.AppendInt(o, z.RefList[za0001])
	}
	o = msgp.AppendArrayHeader(o, uint32(len(z.AuthorAffList)))
	for za0002 := range z.AuthorAffList {
		// array header, size 2
		o = append(o, 0x92)
		o = msgp.AppendInt(o, z.AuthorAffList[za0002].AuthorIndex)
		o = msgp.AppendInt(o, z.AuthorAffList[za0002].AffiliationIndex)
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *PaperInfo) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 5 {
		err = msgp.ArrayError{Wanted: 5, Got: zb0001}
		return
	}
	z.Year, bts, err = msgp.ReadIntBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Year")
		return
	}
	z.VenueIndex, bts, err = msgp.ReadIntBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "VenueIndex")
		return
	}
	z.Citation, bts, err = msgp.ReadIntBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "Citation")
		return
	}
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "RefList")
		return
	}
	if cap(z.RefList) >= int(zb0002) {
		z.RefList = (z.RefList)[:zb0002]
	} else {
		z.RefList = make([]int, zb0002)
	}
	for za0001 := range z.RefList {
		z.RefList[za0001], bts, err = msgp.ReadIntBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, "RefList", za0001)
			return
		}
	}
	var zb0003 uint32
	zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "AuthorAffList")
		return
	}
	if cap(z.AuthorAffList) >= int(zb0003) {
		z.AuthorAffList = (z.AuthorAffList)[:zb0003]
	} else {
		z.AuthorAffList = make([]AuthorAffiliationPair, zb0003)
	}
	for za0002 := range z.AuthorAffList {
		var zb0004 uint32
		zb0004, bts, err = msgp.ReadArrayHeaderBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, "AuthorAffList", za0002)
			return
		}
		if zb0004 != 2 {
			err = msgp.ArrayError{Wanted: 2, Got: zb0004}
			return
		}
		z.AuthorAffList[za0002].AuthorIndex, bts, err = msgp.ReadIntBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, "AuthorAffList", za0002, "AuthorIndex")
			return
		}
		z.AuthorAffList[za0002].AffiliationIndex, bts, err = msgp.ReadIntBytes(bts)
		if err != nil {
			err = msgp.WrapError(err, "AuthorAffList", za0002, "AffiliationIndex")
			return
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *PaperInfo) Msgsize() (s int) {
	s = 1 + msgp.IntSize + msgp.IntSize + msgp.IntSize + msgp.ArrayHeaderSize + (len(z.RefList) * (msgp.IntSize)) + msgp.ArrayHeaderSize + (len(z.AuthorAffList) * (30 + msgp.IntSize + msgp.IntSize))
	return
}

// DecodeMsg implements msgp.Decodable
func (z *RawData) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0001 uint32
	zb0001, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 6 {
		err = msgp.ArrayError{Wanted: 6, Got: zb0001}
		return
	}
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err, "VenueList")
		return
	}
	if cap(z.VenueList) >= int(zb0002) {
		z.VenueList = (z.VenueList)[:zb0002]
	} else {
		z.VenueList = make([]ID, zb0002)
	}
	for za0001 := range z.VenueList {
		{
			var zb0003 uint32
			zb0003, err = dc.ReadUint32()
			if err != nil {
				err = msgp.WrapError(err, "VenueList", za0001)
				return
			}
			z.VenueList[za0001] = ID(zb0003)
		}
	}
	var zb0004 uint32
	zb0004, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err, "AffiliationList")
		return
	}
	if cap(z.AffiliationList) >= int(zb0004) {
		z.AffiliationList = (z.AffiliationList)[:zb0004]
	} else {
		z.AffiliationList = make([]ID, zb0004)
	}
	for za0002 := range z.AffiliationList {
		{
			var zb0005 uint32
			zb0005, err = dc.ReadUint32()
			if err != nil {
				err = msgp.WrapError(err, "AffiliationList", za0002)
				return
			}
			z.AffiliationList[za0002] = ID(zb0005)
		}
	}
	var zb0006 uint32
	zb0006, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err, "CountryList")
		return
	}
	if cap(z.CountryList) >= int(zb0006) {
		z.CountryList = (z.CountryList)[:zb0006]
	} else {
		z.CountryList = make([]ID, zb0006)
	}
	for za0003 := range z.CountryList {
		{
			var zb0007 uint32
			zb0007, err = dc.ReadUint32()
			if err != nil {
				err = msgp.WrapError(err, "CountryList", za0003)
				return
			}
			z.CountryList[za0003] = ID(zb0007)
		}
	}
	var zb0008 uint32
	zb0008, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err, "AuthorList")
		return
	}
	if cap(z.AuthorList) >= int(zb0008) {
		z.AuthorList = (z.AuthorList)[:zb0008]
	} else {
		z.AuthorList = make([]ID, zb0008)
	}
	for za0004 := range z.AuthorList {
		{
			var zb0009 uint32
			zb0009, err = dc.ReadUint32()
			if err != nil {
				err = msgp.WrapError(err, "AuthorList", za0004)
				return
			}
			z.AuthorList[za0004] = ID(zb0009)
		}
	}
	var zb0010 uint32
	zb0010, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err, "PaperList")
		return
	}
	if cap(z.PaperList) >= int(zb0010) {
		z.PaperList = (z.PaperList)[:zb0010]
	} else {
		z.PaperList = make([]ID, zb0010)
	}
	for za0005 := range z.PaperList {
		{
			var zb0011 uint32
			zb0011, err = dc.ReadUint32()
			if err != nil {
				err = msgp.WrapError(err, "PaperList", za0005)
				return
			}
			z.PaperList[za0005] = ID(zb0011)
		}
	}
	var zb0012 uint32
	zb0012, err = dc.ReadArrayHeader()
	if err != nil {
		err = msgp.WrapError(err, "PaperInfoList")
		return
	}
	if cap(z.PaperInfoList) >= int(zb0012) {
		z.PaperInfoList = (z.PaperInfoList)[:zb0012]
	} else {
		z.PaperInfoList = make([]PaperInfo, zb0012)
	}
	for za0006 := range z.PaperInfoList {
		err = z.PaperInfoList[za0006].DecodeMsg(dc)
		if err != nil {
			err = msgp.WrapError(err, "PaperInfoList", za0006)
			return
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *RawData) EncodeMsg(en *msgp.Writer) (err error) {
	// array header, size 6
	err = en.Append(0x96)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.VenueList)))
	if err != nil {
		err = msgp.WrapError(err, "VenueList")
		return
	}
	for za0001 := range z.VenueList {
		err = en.WriteUint32(uint32(z.VenueList[za0001]))
		if err != nil {
			err = msgp.WrapError(err, "VenueList", za0001)
			return
		}
	}
	err = en.WriteArrayHeader(uint32(len(z.AffiliationList)))
	if err != nil {
		err = msgp.WrapError(err, "AffiliationList")
		return
	}
	for za0002 := range z.AffiliationList {
		err = en.WriteUint32(uint32(z.AffiliationList[za0002]))
		if err != nil {
			err = msgp.WrapError(err, "AffiliationList", za0002)
			return
		}
	}
	err = en.WriteArrayHeader(uint32(len(z.CountryList)))
	if err != nil {
		err = msgp.WrapError(err, "CountryList")
		return
	}
	for za0003 := range z.CountryList {
		err = en.WriteUint32(uint32(z.CountryList[za0003]))
		if err != nil {
			err = msgp.WrapError(err, "CountryList", za0003)
			return
		}
	}
	err = en.WriteArrayHeader(uint32(len(z.AuthorList)))
	if err != nil {
		err = msgp.WrapError(err, "AuthorList")
		return
	}
	for za0004 := range z.AuthorList {
		err = en.WriteUint32(uint32(z.AuthorList[za0004]))
		if err != nil {
			err = msgp.WrapError(err, "AuthorList", za0004)
			return
		}
	}
	err = en.WriteArrayHeader(uint32(len(z.PaperList)))
	if err != nil {
		err = msgp.WrapError(err, "PaperList")
		return
	}
	for za0005 := range z.PaperList {
		err = en.WriteUint32(uint32(z.PaperList[za0005]))
		if err != nil {
			err = msgp.WrapError(err, "PaperList", za0005)
			return
		}
	}
	err = en.WriteArrayHeader(uint32(len(z.PaperInfoList)))
	if err != nil {
		err = msgp.WrapError(err, "PaperInfoList")
		return
	}
	for za0006 := range z.PaperInfoList {
		err = z.PaperInfoList[za0006].EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "PaperInfoList", za0006)
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *RawData) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// array header, size 6
	o = append(o, 0x96)
	o = msgp.AppendArrayHeader(o, uint32(len(z.VenueList)))
	for za0001 := range z.VenueList {
		o = msgp.AppendUint32(o, uint32(z.VenueList[za0001]))
	}
	o = msgp.AppendArrayHeader(o, uint32(len(z.AffiliationList)))
	for za0002 := range z.AffiliationList {
		o = msgp.AppendUint32(o, uint32(z.AffiliationList[za0002]))
	}
	o = msgp.AppendArrayHeader(o, uint32(len(z.CountryList)))
	for za0003 := range z.CountryList {
		o = msgp.AppendUint32(o, uint32(z.CountryList[za0003]))
	}
	o = msgp.AppendArrayHeader(o, uint32(len(z.AuthorList)))
	for za0004 := range z.AuthorList {
		o = msgp.AppendUint32(o, uint32(z.AuthorList[za0004]))
	}
	o = msgp.AppendArrayHeader(o, uint32(len(z.PaperList)))
	for za0005 := range z.PaperList {
		o = msgp.AppendUint32(o, uint32(z.PaperList[za0005]))
	}
	o = msgp.AppendArrayHeader(o, uint32(len(z.PaperInfoList)))
	for za0006 := range z.PaperInfoList {
		o, err = z.PaperInfoList[za0006].MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "PaperInfoList", za0006)
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *RawData) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	if zb0001 != 6 {
		err = msgp.ArrayError{Wanted: 6, Got: zb0001}
		return
	}
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "VenueList")
		return
	}
	if cap(z.VenueList) >= int(zb0002) {
		z.VenueList = (z.VenueList)[:zb0002]
	} else {
		z.VenueList = make([]ID, zb0002)
	}
	for za0001 := range z.VenueList {
		{
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "VenueList", za0001)
				return
			}
			z.VenueList[za0001] = ID(zb0003)
		}
	}
	var zb0004 uint32
	zb0004, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "AffiliationList")
		return
	}
	if cap(z.AffiliationList) >= int(zb0004) {
		z.AffiliationList = (z.AffiliationList)[:zb0004]
	} else {
		z.AffiliationList = make([]ID, zb0004)
	}
	for za0002 := range z.AffiliationList {
		{
			var zb0005 uint32
			zb0005, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "AffiliationList", za0002)
				return
			}
			z.AffiliationList[za0002] = ID(zb0005)
		}
	}
	var zb0006 uint32
	zb0006, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "CountryList")
		return
	}
	if cap(z.CountryList) >= int(zb0006) {
		z.CountryList = (z.CountryList)[:zb0006]
	} else {
		z.CountryList = make([]ID, zb0006)
	}
	for za0003 := range z.CountryList {
		{
			var zb0007 uint32
			zb0007, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "CountryList", za0003)
				return
			}
			z.CountryList[za0003] = ID(zb0007)
		}
	}
	var zb0008 uint32
	zb0008, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "AuthorList")
		return
	}
	if cap(z.AuthorList) >= int(zb0008) {
		z.AuthorList = (z.AuthorList)[:zb0008]
	} else {
		z.AuthorList = make([]ID, zb0008)
	}
	for za0004 := range z.AuthorList {
		{
			var zb0009 uint32
			zb0009, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "AuthorList", za0004)
				return
			}
			z.AuthorList[za0004] = ID(zb0009)
		}
	}
	var zb0010 uint32
	zb0010, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "PaperList")
		return
	}
	if cap(z.PaperList) >= int(zb0010) {
		z.PaperList = (z.PaperList)[:zb0010]
	} else {
		z.PaperList = make([]ID, zb0010)
	}
	for za0005 := range z.PaperList {
		{
			var zb0011 uint32
			zb0011, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "PaperList", za0005)
				return
			}
			z.PaperList[za0005] = ID(zb0011)
		}
	}
	var zb0012 uint32
	zb0012, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err, "PaperInfoList")
		return
	}
	if cap(z.PaperInfoList) >= int(zb0012) {
		z.PaperInfoList = (z.PaperInfoList)[:zb0012]
	} else {
		z.PaperInfoList = make([]PaperInfo, zb0012)
	}
	for za0006 := range z.PaperInfoList {
		bts, err = z.PaperInfoList[za0006].UnmarshalMsg(bts)
		if err != nil {
			err = msgp.WrapError(err, "PaperInfoList", za0006)
			return
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *RawData) Msgsize() (s int) {
	s = 1 + msgp.ArrayHeaderSize + (len(z.VenueList) * (msgp.Uint32Size)) + msgp.ArrayHeaderSize + (len(z.AffiliationList) * (msgp.Uint32Size)) + msgp.ArrayHeaderSize + (len(z.CountryList) * (msgp.Uint32Size)) + msgp.ArrayHeaderSize + (len(z.AuthorList) * (msgp.Uint32Size)) + msgp.ArrayHeaderSize + (len(z.PaperList) * (msgp.Uint32Size)) + msgp.ArrayHeaderSize
	for za0006 := range z.PaperInfoList {
		s += z.PaperInfoList[za0006].Msgsize()
	}
	return
}