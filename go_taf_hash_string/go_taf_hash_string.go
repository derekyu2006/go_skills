package go_taf_hash_string

func calcMagicHash(s string) uint32 {
	var value uint32

	for _, c := range s {
		value += uint32(c)
		value += value << 10
		value ^= value >> 6
	}

	value += value << 3
	value ^= value >> 11
	value += value << 15

	if value == 0 {
		value = 1
	}

	return value
}
