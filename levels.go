package revclose

import "math"

func GetLevelsCrossedInRange(reversal, close float64, levels []LevelInterface) []LevelInterface {
	var crossedLevels []LevelInterface

	if reversal != math.SmallestNonzeroFloat64 {
		for _, level := range levels {
			if level.HasCrossed(reversal, close) {
				crossedLevels = append(crossedLevels, level)
			}
		}
	}

	return crossedLevels
}
