package path_calculator

const (
	A = "A"
	B = "B"
	C = "C"
)

type Section struct {
	A int
	B int
	C int
}

type RoadSystem []Section

type Label string
type Movement struct {
	Label Label
	Cost int
}
type Path []Movement

func optimalPath(system RoadSystem) (path Path) {
	optimalA, optimalB := stepRecursive(Path{}, Path{}, system)

	if sum(optimalA) < sum(optimalB) {
		return optimalA
	}

	return optimalB
}

func stepRecursive(pathA Path, pathB Path, system RoadSystem) (newPathA, newPathB Path) {
	newPathA, newPathB = step(pathA, pathB, system[0])

	if len(system) == 1 {
		return
	}

	return stepRecursive(newPathA, newPathB, system[1:])
}

func sum(path Path) (sum int) {
	for _, movement := range path {
		sum += movement.Cost
	}

	return
}

func step(pathA Path, pathB Path, section Section) (newPathA Path, newPathB Path) {
	var priceA int
	for _, path := range pathA {
		priceA += path.Cost
	}
	var priceB int
	for _, path := range pathB {
		priceB += path.Cost
	}

	forwardToA := priceA + section.A
	forwardToB := priceB + section.B
	crossToA := priceB + section.B + section.C
	crossToB := priceA + section.A + section.C

	if forwardToA <= crossToA {
		newPathA = append(pathA, Movement{A, section.A})
	} else {
		newPathA = append(pathB, Movement{B, section.B}, Movement{C, section.C})
	}

	if forwardToB <= crossToB {
		newPathB = append(pathB, Movement{B, section.B})
	} else {
		newPathB = append(pathA, Movement{A, section.A}, Movement{C, section.C})
	}

	return
}


func (movement Movement) equals(other Movement) bool {
	return movement.Cost == other.Cost && movement.Label == other.Label
}

func (path Path) equals(other Path) bool {
	if len(path) != len(other) {
		return false
	}

	for i := range path {
		if !path[i].equals(other[i]) {
			return false
		}
	}

	return true
}

