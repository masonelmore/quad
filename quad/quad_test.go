package quad

import (
	"image"
	"image/png"
	"os"
	"testing"
)

func TestAspectNoPanic(t *testing.T) {
	tests := []struct {
		filename string
		name     string
	}{
		{"beach-320x320.png", "square"},
		{"beach-480x320.png", "notsquare"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			img, err := openImage(test.filename)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			Quadify(img, 0)
		})
	}
}

func BenchmarkQuadify(b *testing.B) {
	// Testing multiple resoultions is probably overkill.  The performance
	// between resolutions is very close.
	benchmarks := []struct {
		filename  string
		tolerance float64
		name      string
	}{
		{"lowdetail-480.png", 0.1, "low-480-0.1"},
		{"lowdetail-960.png", 0.1, "low-960-0.1"},
		{"lowdetail-1920.png", 0.1, "low-1920-0.1"},
		{"highdetail-480.png", 0.1, "high-480-0.1"},
		{"highdetail-960.png", 0.1, "high-960-0.1"},
		{"highdetail-1920.png", 0.1, "high-1920-0.1"},

		{"lowdetail-480.png", 0.05, "low-480-0.05"},
		{"lowdetail-960.png", 0.05, "low-960-0.05"},
		{"lowdetail-1920.png", 0.05, "low-1920-0.05"},
		{"highdetail-480.png", 0.05, "high-480-0.05"},
		{"highdetail-960.png", 0.05, "high-960-0.05"},
		{"highdetail-1920.png", 0.05, "high-1920-0.05"},

		{"lowdetail-480.png", 0.0, "low-480-0.0"},
		{"lowdetail-960.png", 0.0, "low-960-0.0"},
		{"lowdetail-1920.png", 0.0, "low-1920-0.0"},
		{"highdetail-480.png", 0.0, "high-480-0.0"},
		{"highdetail-960.png", 0.0, "high-960-0.0"},
		{"highdetail-1920.png", 0.0, "high-1920-0.0"},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				img, err := openImage(benchmark.filename)
				if err != nil {
					b.Error(err)
					b.FailNow()
				}
				Quadify(img, 0)
			}
		})
	}
}

func openImage(filename string) (image.Image, error) {
	f, err := os.Open("testdata/" + filename)
	if err != nil {
		return nil, err
	}
	i, err := png.Decode(f)
	if err != nil {
		return nil, err
	}
	return i, nil
}
