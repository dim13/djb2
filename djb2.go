// Package djb2 implements Bernstein hash algorithm
package djb2

type digest struct {
	hash uint32
}

func (d *digest) BlockSize() int {
	return 1
}

func (d *digest) Reset() {
	d.hash = 5381
}

func (d *digest) Size() int {
	return 8
}

func (d *digest) Sum(b []byte) []byte {
	s := d.Sum32()
	for i := d.Size() - 1; i >= 0; i-- {
		b = append(b, byte(s>>uint(8*i)))
	}
	return b
}

func (d *digest) Sum32() uint32 {
	return d.hash
}

func (d *digest) Write(p []byte) (n int, err error) {
	for _, v := range p {
		d.hash = ((d.hash << 5) + d.hash) + uint32(v)
	}
	return len(p), nil
}

func Sum(b []byte) uint32 {
	d := new(digest)
	d.Reset()
	d.Write(b)
	return d.Sum32()
}

func SumString(s string) uint32 {
	return Sum([]byte(s))
}
