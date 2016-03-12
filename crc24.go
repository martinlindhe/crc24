package crc24

const (
	crc24Init = 0xb704ce
	crc24Poly = 0x1864cfb
	crc24Mask = 0xffffff
)

// ChecksumOpenPGP calculates the crc-24 as used by OpenPGP
func ChecksumOpenPGP(p []byte) uint32 {

	return crc24(0, p)
}

// crc24 calculates the OpenPGP checksum as specified in RFC 4880, section 6.1
func crc24(crc uint32, d []byte) uint32 {
	for _, b := range d {
		crc ^= uint32(b) << 16
		for i := 0; i < 8; i++ {
			crc <<= 1
			if crc&0x1000000 != 0 {
				crc ^= crc24Poly
			}
		}
	}
	return crc
}
