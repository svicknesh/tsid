package tsid

// build the alphabets list and their values for easier acquisition
var uppercase = []rune("0123456789ABCDEFGHJKMNPQRSTVWXYZ")
var alphabets []uint64

func init() {

	alphabets = make([]uint64, 128)

	// this definition lets us to a mapping of the alphabets to numbers easily

	// define the numbers
	alphabets['0'] = 0x00
	alphabets['1'] = 0x01
	alphabets['2'] = 0x02
	alphabets['3'] = 0x03
	alphabets['4'] = 0x04
	alphabets['5'] = 0x05
	alphabets['6'] = 0x06
	alphabets['7'] = 0x07
	alphabets['8'] = 0x08
	alphabets['9'] = 0x09

	// define the lowercase, characters 'i' and 'l' are 1 and 'o' is 0, 'u' odes not exist
	alphabets['a'] = 0x0a
	alphabets['b'] = 0x0b
	alphabets['c'] = 0x0c
	alphabets['d'] = 0x0d
	alphabets['e'] = 0x0e
	alphabets['f'] = 0x0f
	alphabets['g'] = 0x10
	alphabets['h'] = 0x11
	alphabets['i'] = 0x01
	alphabets['j'] = 0x12
	alphabets['k'] = 0x13
	alphabets['l'] = 0x01
	alphabets['m'] = 0x14
	alphabets['n'] = 0x15
	alphabets['o'] = 0x00
	alphabets['p'] = 0x16
	alphabets['q'] = 0x17
	alphabets['r'] = 0x18
	alphabets['s'] = 0x19
	alphabets['t'] = 0x1a
	alphabets['v'] = 0x1b
	alphabets['w'] = 0x1c
	alphabets['x'] = 0x1d
	alphabets['y'] = 0x1e
	alphabets['z'] = 0x1f

	// define the uppercase, characters 'I' and 'L' are 1 and 'O' is 0, 'U' odes not exist
	alphabets['A'] = 0x0a
	alphabets['B'] = 0x0b
	alphabets['C'] = 0x0c
	alphabets['D'] = 0x0d
	alphabets['E'] = 0x0e
	alphabets['F'] = 0x0f
	alphabets['G'] = 0x10
	alphabets['H'] = 0x11
	alphabets['I'] = 0x01
	alphabets['J'] = 0x12
	alphabets['K'] = 0x13
	alphabets['L'] = 0x01
	alphabets['M'] = 0x14
	alphabets['N'] = 0x15
	alphabets['O'] = 0x00
	alphabets['P'] = 0x16
	alphabets['Q'] = 0x17
	alphabets['R'] = 0x18
	alphabets['S'] = 0x19
	alphabets['T'] = 0x1a
	alphabets['V'] = 0x1b
	alphabets['W'] = 0x1c
	alphabets['X'] = 0x1d
	alphabets['Y'] = 0x1e
	alphabets['Z'] = 0x1f

}
