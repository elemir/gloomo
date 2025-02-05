package repo

import (
	"image"
	"iter"

	gid "github.com/elemir/gloomo/id"
	"github.com/elemir/gloomo/model"
)

type AnimatedSprite struct {
	Animations        Collection[*model.Animation]
	Positions         Collection[image.Point]
	ZIndices          Collection[int]
	StepCounters      Collection[int]
	CurrentAnimations Collection[string]
}

func (a *AnimatedSprite) List() iter.Seq2[gid.ID, model.AnimatedSprite] {
	return func(yield func(gid.ID, model.AnimatedSprite) bool) {
		for id, anim := range a.Animations.Items() {
			counter, ok := a.StepCounters.Get(id)
			if !ok {
				continue
			}

			current, ok := a.CurrentAnimations.Get(id)
			if !ok {
				continue
			}

			pos, ok := a.Positions.Get(id)
			if !ok {
				continue
			}

			zIndex, ok := a.ZIndices.Get(id)
			if !ok {
				continue
			}

			sprite := model.AnimatedSprite{
				Animation: anim,
				Position:  pos,
				ZIndex:    zIndex,
				Counter:   counter,
				Current:   current,
			}

			if !yield(id, sprite) {
				return
			}
		}
	}
}

func (a *AnimatedSprite) Upsert(id gid.ID, sprite model.AnimatedSprite) {
	a.Animations.Set(id, sprite.Animation)
	a.StepCounters.Set(id, sprite.Counter)
	a.CurrentAnimations.Set(id, sprite.Current)
	a.Positions.Set(id, sprite.Position)
	a.ZIndices.Set(id, sprite.ZIndex)
}

func (a *AnimatedSprite) Get(id gid.ID) (model.AnimatedSprite, bool) {
	anim, ok := a.Animations.Get(id)
	if !ok {
		return model.AnimatedSprite{}, false
	}

	counter, ok := a.StepCounters.Get(id)
	if !ok {
		return model.AnimatedSprite{}, false
	}

	current, ok := a.CurrentAnimations.Get(id)
	if !ok {
		return model.AnimatedSprite{}, false
	}

	pos, ok := a.Positions.Get(id)
	if !ok {
		return model.AnimatedSprite{}, false
	}

	zIndex, ok := a.ZIndices.Get(id)
	if !ok {
		return model.AnimatedSprite{}, false
	}

	return model.AnimatedSprite{
		Animation: anim,
		Position:  pos,
		ZIndex:    zIndex,
		Counter:   counter,
		Current:   current,
	}, true
}
