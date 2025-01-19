package loader

import (
	"errors"
	"fmt"
	"iter"
	"path"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type ImageAssets interface {
	NotLoadedPaths() iter.Seq[string]
	Put(path string, val *ebiten.Image)
}

type Image struct {
	AssetDir string
	Assets   ImageAssets
}

func (i *Image) Run() error {
	var errs []error

	for assetPath := range i.Assets.NotLoadedPaths() {
		img, _, err := ebitenutil.NewImageFromFile(path.Join(i.AssetDir, assetPath))
		if err != nil {
			errs = append(errs, fmt.Errorf("load %q image: %w", assetPath, err))

			continue
		}

		i.Assets.Put(assetPath, img)
	}

	return errors.Join(errs...)
}
