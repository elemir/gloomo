package loader

import (
	"errors"
	"fmt"
	"image"
	"io"
	"iter"
	"os"
	"path"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"gopkg.in/yaml.v3"

	"github.com/elemir/gloomo/model"
)

type AnimAsset struct {
	SpriteSheet string   `yaml:"spritesheet"`
	Size        AnimSize `yaml:"size"`
	Animations  []Anim   `yaml:"steps"`
}

type AnimSize struct {
	X int `yaml:"x"`
	Y int `yaml:"y"`
}

type Anim struct {
	Name  string `yaml:"name"`
	Steps []int  `yaml:"steps"`
}

type AnimAssets interface {
	NotLoadedPaths() iter.Seq[string]
	Put(path string, typ string, val model.Animation)
}

type Animation struct {
	AssetDir string
	Assets   AnimAssets
}

func (a *Animation) Run() error {
	var errs []error

	for assetPath := range a.Assets.NotLoadedPaths() {
		animAsset, err := loadAnimationAsset(path.Join(a.AssetDir, assetPath))
		if err != nil {
			errs = append(errs, fmt.Errorf("load animation asset file %q: %w", assetPath, err))
			continue
		}

		// THINK(evgenii.omelchenko): should we use loader here?
		spriteSheet, _, err := ebitenutil.NewImageFromFile(path.Join(a.AssetDir, animAsset.SpriteSheet))
		if err != nil {
			errs = append(errs, fmt.Errorf("load image %q: %w", assetPath, err))
			continue
		}

		for _, anim := range animAsset.Animations {
			var animation model.Animation

			for _, step := range anim.Steps {
				frame := getSpecificFrame(spriteSheet, animAsset.Size, step)
				animation.Steps = append(animation.Steps, frame)
			}

			a.Assets.Put(assetPath, anim.Name, animation)
		}
	}

	return errors.Join(errs...)
}

func loadAnimationAsset(path string) (AnimAsset, error) {
	var animAsset AnimAsset

	assetFile, err := os.Open(path)
	if err != nil {
		return AnimAsset{}, fmt.Errorf("open file %q: %w", path, err)
	}

	defer assetFile.Close()

	assetData, err := io.ReadAll(assetFile)
	if err != nil {
		return AnimAsset{}, fmt.Errorf("read data from file %q: %w", path, err)
	}

	err = yaml.Unmarshal(assetData, &animAsset)
	if err != nil {
		return AnimAsset{}, fmt.Errorf("unmarshal data: %w", err)
	}

	return animAsset, nil
}

func getSpecificFrame(spriteSheet *ebiten.Image, size AnimSize, step int) *ebiten.Image {
	width := spriteSheet.Bounds().Dx()
	countInRow := width / size.X
	x, y := step%countInRow, step/countInRow

	rect := image.Rect(x*size.X, y*size.Y, (x+1)*size.X, (y+1)*size.Y)

	//nolint:forcetypeassert // Typing issue: really ebiten.Image.SubImage always returns ebiten.Image
	return spriteSheet.SubImage(rect).(*ebiten.Image)
}
