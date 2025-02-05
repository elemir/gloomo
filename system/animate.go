package system

import (
	"errors"
	"fmt"
	"iter"

	gid "github.com/elemir/gloomo/id"
	"github.com/elemir/gloomo/model"
	"github.com/elemir/gloomo/node"
)

var errUnknownCurrentAnimation = errors.New("unknown current animation")

const (
	Speed = 10
)

type SpriteRepo interface {
	Upsert(id gid.ID, sprite node.Sprite)
}

type AnimationRepo interface {
	Upsert(id gid.ID, anim model.AnimatedSprite)
	List() iter.Seq2[gid.ID, model.AnimatedSprite]
}

type Animate struct {
	SpriteRepo    SpriteRepo
	AnimationRepo AnimationRepo
}

func (a *Animate) Run() error {
	for id, anim := range a.AnimationRepo.List() {
		currentSteps := anim.Animation.Steps[anim.Current]
		if len(currentSteps) == 0 {
			return fmt.Errorf("animation %q: %w", anim.Current, errUnknownCurrentAnimation)
		}

		anim.Counter = (anim.Counter + 1) % (len(currentSteps) * Speed)
		a.AnimationRepo.Upsert(id, anim)

		a.SpriteRepo.Upsert(id, node.Sprite{
			Image:    currentSteps[anim.Counter/Speed],
			Position: anim.Position,
			ZIndex:   anim.ZIndex,
		})
	}

	return nil
}
