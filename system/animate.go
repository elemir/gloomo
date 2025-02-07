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
	for id, sprite := range a.AnimationRepo.List() {
		animation := sprite.AnimationSheet.Animations[sprite.Current]
		if len(animation.Steps) == 0 {
			return fmt.Errorf("animation %q: %w", sprite.Current, errUnknownCurrentAnimation)
		}

		sprite.Counter = (sprite.Counter + 1) % (len(animation.Steps) * Speed)
		a.AnimationRepo.Upsert(id, sprite)

		a.SpriteRepo.Upsert(id, node.Sprite{
			Image:    animation.Steps[sprite.Counter/Speed],
			Position: sprite.Position.Round(),
			ZIndex:   sprite.ZIndex,
			Mirror:   animation.Mirror,
		})
	}

	return nil
}
