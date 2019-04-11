package lib

// PrimeSieve - Implementation of Sieve of Eratosthenes
// Initiate with NewPrimeSieve
type PrimeSieve struct {
	primeBitSet []uint64 // If bit is 0, means isPrime. (Saves time setting up the array)
}

func (p *PrimeSieve) getLocation(location int) (int, uint) {
	return location >> 6, uint(location & 0x3f) // divide by 64, modulo 64
}

func (p *PrimeSieve) setBit(location int) {
	x, y := p.getLocation(location)

	p.primeBitSet[x] |= 1 << y
}

func (p *PrimeSieve) clearBit(location int) {
	x, y := p.getLocation(location)

	p.primeBitSet[x] &= ^(1 << y)
}

func (p *PrimeSieve) getBit(location int) bool {
	x, y := p.getLocation(location)

	return (p.primeBitSet[x] & (1 << y)) != 0
}

// NewPrimeSieve -
func NewPrimeSieve(maxSearch int) *PrimeSieve {
	bitset := make([]uint64, (maxSearch/64)+1)
	p := &PrimeSieve{
		primeBitSet: bitset,
	}

	// 0 and 1 are not primes. Remember that we are doing the reverse, 1 means not prime.
	p.setBit(0)
	p.setBit(1)

	// Execute sieve
	for i := 2; i <= maxSearch; i++ {
		if p.getBit(i) == false { // Doing the reverse, 0 means prime
			for j := 2; j*i <= maxSearch; j++ {
				p.setBit(j * i) // Doing reverse, 1 means not prime
			}
		}
	}

	return p
}

// IsPrime -
func (p *PrimeSieve) IsPrime(num int) bool {
	return !p.getBit(num) // Doing reverse, 0 means prime, 1 means not prime
}
