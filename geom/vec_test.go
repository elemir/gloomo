package geom_test

import (
	"testing"

	"github.com/elemir/gloomo/geom"
	"github.com/stretchr/testify/require"
)

func TestAngles(t *testing.T) {
	var vecs []geom.Vec2

	for i := range 3 {
		for j := range 3 {
			vecs = append(vecs, geom.Vec2{
				float64(i - 1), float64(j - 1),
			})
		}
	}

	for _, v1 := range vecs {
		for _, v2 := range vecs {
			var rotatedV geom.Vec2

			angle := v1.AngleBetween(v2)
			if v1.Length() == 0 {
				rotatedV = v2.Normalize()
			} else if v2.Length() != 0 {
				rotatedV = v1.Rotate(angle).Normalize()
			}

			projV := rotatedV.Mul(rotatedV.Dot(v2))

			require.InDeltaf(t, 0.0, v2.Distance(projV), 0.01, "rotated vector should be the same: v1=%v, v2=%v, angle=%v, rotatedV=%v", v1, v2, angle, rotatedV)
		}
	}
}

func TestRotate(t *testing.T) {
	var vecs []geom.Vec2

	for i := range 3 {
		for j := range 3 {
			vecs = append(vecs, geom.Vec2{
				float64(i - 1), float64(j - 1),
			}.Normalize())
		}
	}

	for _, v1 := range vecs {
		for _, v2 := range vecs {
			var rotatedV geom.Vec2

			angle := v1.AngleBetween(v2)
			if v1.Length() == 0 {
				rotatedV = v2
			} else if v2.Length() != 0 {
				rotatedV = v1.Rotate(angle)
			}

			require.InDeltaf(t, 0.0, v2.Distance(rotatedV), 0.01, "rotated vector should be the same: v1=%v, v2=%v, angle=%v, rotatedV=%v", v1, v2, angle, rotatedV)
		}
	}
}
