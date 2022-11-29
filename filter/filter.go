package filter

import "github.com/disintegration/imaging"

type Filter interface {
	Process(srcPath, dstPath string) error
}
type Grayscale struct{}

func (g *Grayscale) Process(srcPath, dstPath string) error {
	src, err := imaging.Open(srcPath)
	return nil
}
