package quad

import (
	"image"
	"image/color"
	"math"
)

// Quadify transforms an image into a bunch of squares.  A higher tolerance
// lowers the amount of detail in the final image.
func Quadify(img image.Image, tolerance float64) *image.RGBA {
	qt := newQuadTree(img, tolerance)
	out := image.NewRGBA(img.Bounds())

	nodes := []*node{qt.root}
	for len(nodes) > 0 {
		n := nodes[0]
		nodes = nodes[1:]

		n.fillRegion(out)
		n.drawBorder(out)

		if len(n.children) != 0 {
			nodes = append(nodes, n.children...)
		}
	}

	return out
}

type quadtree struct {
	root *node
}

func newQuadTree(img image.Image, tolerance float64) *quadtree {
	q := new(quadtree)
	q.root = newNode(img, img.Bounds(), tolerance)
	return q
}

type node struct {
	img      image.Image
	region   image.Rectangle
	color    color.Color
	diff     float64
	children []*node
}

func newNode(img image.Image, region image.Rectangle, tolerance float64) *node {
	n := new(node)
	n.img = img
	n.region = region
	n.color = averageColor(img, region)
	n.diff = difference(img, region, n.color)

	if n.diff > tolerance {
		a, b, c, d := quarter(region)
		n.children = []*node{
			newNode(img, a, tolerance),
			newNode(img, b, tolerance),
			newNode(img, c, tolerance),
			newNode(img, d, tolerance),
		}
	}

	return n
}

func (n *node) fillRegion(img *image.RGBA) {
	for y := n.region.Min.Y; y < n.region.Max.Y; y++ {
		for x := n.region.Min.X; x < n.region.Max.X; x++ {
			img.Set(x, y, n.color)
		}
	}
}

func (n *node) drawBorder(img *image.RGBA) {
	// Vertical borders
	for y := n.region.Min.Y; y < n.region.Max.Y; y++ {
		img.Set(n.region.Min.X, y, color.Black)
		img.Set(n.region.Max.X, y, color.Black)
	}

	// Horizontal borders
	for x := n.region.Min.X; x < n.region.Max.X; x++ {
		img.Set(x, n.region.Min.Y, color.Black)
		img.Set(x, n.region.Max.Y, color.Black)
	}
}

func averageColor(imag image.Image, rect image.Rectangle) color.Color {
	var r, g, b uint64
	for py := rect.Min.Y; py < rect.Max.Y; py++ {
		for px := rect.Min.X; px < rect.Max.X; px++ {
			pr, pg, pb, _ := imag.At(px, py).RGBA()
			r += uint64(pr * pr)
			g += uint64(pg * pg)
			b += uint64(pb * pb)
		}
	}

	area := uint64(rect.Dx() * rect.Dy())
	r = uint64(math.Sqrt(float64(r / area)))
	g = uint64(math.Sqrt(float64(g / area)))
	b = uint64(math.Sqrt(float64(b / area)))

	return color.RGBA{uint8(r / 0x101), uint8(g / 0x101), uint8(b / 0x101), 255}
}

func difference(img image.Image, region image.Rectangle, avg color.Color) float64 {
	const maxDifference = 65535
	r, g, b, _ := avg.RGBA()
	var sum float64
	for y := region.Min.Y; y < region.Max.Y; y++ {
		for x := region.Min.X; x < region.Max.X; x++ {
			c := img.At(x, y)
			rr, gg, bb, _ := c.RGBA()
			sum += math.Abs(float64(r) - float64(rr))
			sum += math.Abs(float64(g) - float64(gg))
			sum += math.Abs(float64(b) - float64(bb))
		}
	}

	area := region.Dx() * region.Dy()
	return (sum / float64(3*area)) / maxDifference
}

func quarter(r image.Rectangle) (a, b, c, d image.Rectangle) {
	halfX := (r.Dx() / 2) + r.Min.X
	halfY := (r.Dy() / 2) + r.Min.Y

	a = image.Rect(r.Min.X, r.Min.Y, halfX, halfY)
	b = image.Rect(halfX, r.Min.Y, r.Max.X, halfY)
	c = image.Rect(r.Min.X, halfY, halfX, r.Max.Y)
	d = image.Rect(halfX, halfY, r.Max.X, r.Max.Y)

	return a, b, c, d
}
