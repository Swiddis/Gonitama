package onitama

type Move struct {
	Dx int `json:"dx"`
	Dy int `json:"dy"`
}

type BitMove struct {
	Card uint
	Move uint
	Mask uint
}

// A bitmask, centered at bit 12, depicting a move.
// Has a 1 at the origin and a 1 at the destination.
func getMoveMask(move Move) uint {
	var moveMask uint = 1 << 12
	if move.Dx >= 0 {
		moveMask >>= move.Dx
	} else {
		moveMask <<= -move.Dx
	}
	if move.Dy >= 0 {
		moveMask <<= 5 * move.Dy
	} else {
		moveMask >>= -5 * move.Dy
	}
	moveMask |= 1 << 12
	return moveMask
}

// A bitmask corresponding to which squares might be able to make a move.
// This is effectively where bounds checking happens.
func getMoveableMask(move Move) uint {
	var moveableMask uint = 0b1111111111111111111111111
	switch move.Dx {
	case 2:
		moveableMask &= ^uint(0b0001100011000110001100011)
	case 1:
		moveableMask &= ^uint(0b0000100001000010000100001)
	case -1:
		moveableMask &= ^uint(0b1000010000100001000010000)
	case -2:
		moveableMask &= ^uint(0b1100011000110001100011000)
	default:
		break
	}

	switch move.Dy {
	case 2:
		moveableMask &= ^uint(0b1111111111000000000000000)
	case 1:
		moveableMask &= ^uint(0b1111100000000000000000000)
	case -1:
		moveableMask &= ^uint(0b0000000000000000000011111)
	case -2:
		moveableMask &= ^uint(0b0000000000000001111111111)
	default:
		break
	}
	return moveableMask
}
