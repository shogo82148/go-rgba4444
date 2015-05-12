package rgba4444

import "image/color"

// RGBA represents a traditional 16-bit alpha-premultiplied color,
// having 4 bits for each of red, green, blue and alpha.
type Color struct {
	R, G, B, A uint8
}

func (c Color) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R)
	r |= r << 4
	r |= r << 8
	g = uint32(c.G)
	g |= g << 4
	g |= g << 8
	b = uint32(c.B)
	b |= b << 4
	b |= b << 8
	a = uint32(c.A)
	a |= a << 4
	a |= a << 8
	return
}

var Model color.Model = color.ModelFunc(rgba4444Model)

func rgba4444Model(c color.Color) color.Color {
	if _, ok := c.(Color); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	return Color{uint8(r >> 12), uint8(g >> 12), uint8(b >> 12), uint8(a >> 12)}
}
