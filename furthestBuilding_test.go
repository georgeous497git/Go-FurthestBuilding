package furthest

import "testing"

func TestFurthestBuilding(t *testing.T) {
	heights := []int{4, 12, 2, 7, 3, 18, 20, 3, 19}
	bricks := 10
	ladders := 2
	want := 7
	max := FurthestBuilding(heights, bricks, ladders)

	if want != max {
		t.Errorf("Getting %d but the furthest must be equals to: %d", max, want)
	}

	heights2 := []int{14, 3, 19, 3}
	bricks2 := 17
	ladders2 := 0
	want2 := 3
	max2 := FurthestBuilding(heights2, bricks2, ladders2)

	if want2 != max2 {
		t.Errorf("Getting %d but the furthest must be equals to: %d", max2, want2)
	}

	heights3 := []int{14, 3, 19, 3}
	bricks3 := 17
	ladders3 := 0
	want3 := 3
	max3 := FurthestBuilding(heights3, bricks3, ladders3)

	if want3 != max3 {
		t.Errorf("Getting %d but the furthest must be equals to: %d", max2, want2)
	}

	heights4 := []int{1, 5, 1, 2, 3, 4, 10000}
	bricks4 := 4
	ladders4 := 1
	want4 := 5
	max4 := FurthestBuilding(heights4, bricks4, ladders4)

	if want4 != max4 {
		t.Errorf("Getting %d but the furthest must be equals to: %d", max4, want4)
	}
}
