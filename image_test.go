package rgba4444

import (
	"image"
	"image/color"
	"testing"
)

type img interface {
	image.Image
	Opaque() bool
	Set(int, int, color.Color)
	SubImage(image.Rectangle) image.Image
}

func cmp(t *testing.T, cm color.Model, c0, c1 color.Color) bool {
	r0, g0, b0, a0 := cm.Convert(c0).RGBA()
	r1, g1, b1, a1 := cm.Convert(c1).RGBA()
	return r0 == r1 && g0 == g1 && b0 == b1 && a0 == a1
}

func TestImage(t *testing.T) {
	var m img = New(image.Rect(0, 0, 10, 10))
	if !image.Rect(0, 0, 10, 10).Eq(m.Bounds()) {
		t.Errorf("%T: want bounds %v, got %v", m, image.Rect(0, 0, 10, 10), m.Bounds())
	}
	if !cmp(t, m.ColorModel(), image.Transparent, m.At(6, 3)) {
		t.Errorf("%T: at (6, 3), want a zero color, got %v", m, m.At(6, 3))
	}
	m.Set(6, 3, image.Opaque)
	if !cmp(t, m.ColorModel(), image.Opaque, m.At(6, 3)) {
		t.Errorf("%T: at (6, 3), want a non-zero color, got %v", m, m.At(6, 3))
	}
	if !m.SubImage(image.Rect(6, 3, 7, 4)).(img).Opaque() {
		t.Errorf("%T: at (6, 3) was not opaque", m)
	}

	m = m.SubImage(image.Rect(3, 2, 9, 8)).(img)
	if !image.Rect(3, 2, 9, 8).Eq(m.Bounds()) {
		t.Errorf("%T: sub-image want bounds %v, got %v", m, image.Rect(3, 2, 9, 8), m.Bounds())
	}
	if !cmp(t, m.ColorModel(), image.Opaque, m.At(6, 3)) {
		t.Errorf("%T: sub-image at (6, 3), want a non-zero color, got %v", m, m.At(6, 3))
	}
	if !cmp(t, m.ColorModel(), image.Transparent, m.At(3, 3)) {
		t.Errorf("%T: sub-image at (3, 3), want a zero color, got %v", m, m.At(3, 3))
	}
	m.Set(3, 3, image.Opaque)
	if !cmp(t, m.ColorModel(), image.Opaque, m.At(3, 3)) {
		t.Errorf("%T: sub-image at (3, 3), want a non-zero color, got %v", m, m.At(3, 3))
	}

	// Test that taking an empty sub-image starting at a corner does not panic.
	m.SubImage(image.Rect(0, 0, 0, 0))
	m.SubImage(image.Rect(10, 0, 10, 0))
	m.SubImage(image.Rect(0, 10, 0, 10))
	m.SubImage(image.Rect(10, 10, 10, 10))
}
