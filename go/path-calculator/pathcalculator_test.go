package path_calculator

import "testing"

func TestWhenCalulatingHeathrowToLondon(t *testing.T) {
	heathrowToLondon := RoadSystem{
		Section{50, 10, 30},
		Section{5, 90, 20},
		Section{40, 2, 25},
		Section{10, 8, 0},
	}
	expected := Path{
		{B, 10},
		{C, 30},
		{A, 5},
		{C, 20},
		{B, 2},
		{B, 8},
	}
	actual := optimalPath(heathrowToLondon)
	t.Run("should return correct path", func(t *testing.T) {
		if !actual.equals(expected) {
			t.Fatalf("expected: %v got: %v", expected, actual)
		}
	})
}

func TestWhenCalculatingASingleStepAndNoPreExistingPaths(t *testing.T) {
	expectedA := Path{
		Movement{B,10},
		Movement{C,30},
	}
	expectedB := Path{
		Movement{B,10},
	}
	section := Section{50, 10, 30}
	actualA, actualB := step(Path{}, Path{}, section)

	t.Run("should return correct path a", func(t *testing.T){
		if !actualA.equals(expectedA) {
			t.Fail()
		}
	})
	t.Run("should return correct path b", func(t *testing.T){
		if !actualB.equals(expectedB) {
			t.Fail()
		}
	})
}

func TestWhenCalculatingASingleStepAndAHasLongerPreExistingPath(t *testing.T) {
	preA := Path{Movement{A,100}}
	preB := Path{Movement{B,50}}
	expectedA := Path{
		Movement{B,50},
		Movement{B,10},
		Movement{C,30},
	}
	expectedB := Path{
		Movement{B,50},
		Movement{B,10},
	}
	section := Section{50, 10, 30}
	actualA, actualB := step(preA, preB, section)

	t.Run("should return correct path a", func(t *testing.T){
		if !actualA.equals(expectedA) {
			t.Fail()
		}
	})
	t.Run("should return correct path b", func(t *testing.T){
		if !actualB.equals(expectedB) {
			t.Fail()
		}
	})
}

func TestWhenCalculatingASingleStepAndAHasLongerPreExistingPathAndCIsLong(t *testing.T) {
	preA := Path{Movement{A,100}}
	preB := Path{Movement{B,50}}
	expectedA := Path{
		Movement{A,100},
		Movement{A,50},
	}
	expectedB := Path{
		Movement{B,50},
		Movement{B,100},
	}
	section := Section{50, 100, 30}
	actualA, actualB := step(preA, preB, section)

	t.Run("should return correct path a", func(t *testing.T){
		if !actualA.equals(expectedA) {
			t.Fail()
		}
	})
	t.Run("should return correct path b", func(t *testing.T){
		if !actualB.equals(expectedB) {
			t.Fail()
		}
	})
}

func TestWhenCalculatingASingleStepAndBHasLongerPreExistingPath(t *testing.T) {
	preA := Path{Movement{A,60}}
	preB := Path{Movement{B,100}}
	expectedA := Path{
		Movement{A,60},
		Movement{A,50},
	}
	expectedB := Path{
		Movement{B,100},
		Movement{B,10},
	}
	section := Section{50, 10, 30}
	actualA, actualB := step(preA, preB, section)

	t.Run("should return correct path a", func(t *testing.T){
		if !actualA.equals(expectedA) {
			t.Fail()
		}
	})
	t.Run("should return correct path b", func(t *testing.T){
		if !actualB.equals(expectedB) {
			t.Fail()
		}
	})
}

func TestWhenCalculatingASingleStepAndBHasLongerPreExistingPathAndCIsLong(t *testing.T) {
	preA := Path{Movement{A,60}}
	preB := Path{Movement{A,50}}
	expectedA := Path{
		Movement{A,60},
		Movement{A,50},
	}
	expectedB := Path{
		Movement{A,60},
		Movement{A,50},
		Movement{C,30},
	}
	section := Section{50, 100, 30}
	actualA, actualB := step(preA, preB, section)

	t.Run("should return correct path a", func(t *testing.T){
		if !actualA.equals(expectedA) {
			t.Fail()
		}
	})
	t.Run("should return correct path b", func(t *testing.T){
		if !actualB.equals(expectedB) {
			t.Fail()
		}
	})
}
