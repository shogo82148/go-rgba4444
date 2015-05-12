package rgba4444

import (
	"image"
	"image/color"
)

// Image is an in-memory image whose At method returns rgba4444.Color values.
type Image struct {
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*2].
	Pix []uint8
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect image.Rectangle
}

func New(r image.Rectangle) *Image {
	w, h := r.Dx(), r.Dy()
	pix := make([]uint8, 2*w*h)
	return &Image{pix, 2 * w, r}
}

func (p *Image) ColorModel() color.Model { return Model }

func (p *Image) Bounds() image.Rectangle { return p.Rect }

func (p *Image) At(x, y int) color.Color {
	return p.RGBA4444At(x, y)
}

func (p *Image) RGBA4444At(x, y int) Color {
	if !(image.Point{x, y}.In(p.Rect)) {
		return Color{}
	}
	i := p.PixOffset(x, y)
	return Color{
		uint8(p.Pix[i+0] >> 4),
		uint8(p.Pix[i+0] & 0x0F),
		uint8(p.Pix[i+1] >> 4),
		uint8(p.Pix[i+1] & 0x0F),
	}
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *Image) PixOffset(x, y int) int {
	return (y-p.Rect.Min.Y)*p.Stride + (x-p.Rect.Min.X)*2
}

func (p *Image) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := Model.Convert(c).(Color)
	p.Pix[i+0] = (c1.R << 4) | (c1.G & 0x0F)
	p.Pix[i+1] = (c1.B << 4) | (c1.A & 0x0F)
}

func (p *Image) SetRGBA4444(x, y int, c Color) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	p.Pix[i+0] = (c.R << 4) | (c.G & 0x0F)
	p.Pix[i+1] = (c.B << 4) | (c.A & 0x0F)
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *Image) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.Rect)
	if r.Empty() {
		return &Image{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return &Image{
		Pix:    p.Pix[i:],
		Stride: p.Stride,
		Rect:   r,
	}
}

// Opaque scans the entire image and reports whether it is fully opaque.
func (p *Image) Opaque() bool {
	if p.Rect.Empty() {
		return true
	}
	i0, i1 := 1, p.Rect.Dx()*2
	for y := p.Rect.Min.Y; y < p.Rect.Max.Y; y++ {
		for i := i0; i < i1; i += 2 {
			if (p.Pix[i] & 0x0F) != 0x0F {
				return false
			}
		}
		i0 += p.Stride
		i1 += p.Stride
	}
	return true
}
