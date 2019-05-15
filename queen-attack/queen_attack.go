package queenattack


import "errors"


type Queen struct {
	x, y int
}

func newQueen(pos string) (*Queen, error) {
	if len(pos) != 2 {
		return nil, errors.New("Position invalid")
	}
	x, y := pos[0] - 'a' + 1, pos[1] - '0'
	if x <= 0 || x > 8 {
		return nil, errors.New("Off board")
	}
	if y <= 0 || y > 8 {
		return nil, errors.New("Off board")
	}
	return &Queen{int(x), int(y)}, nil
}

func canAttack(q1, q2 *Queen) bool {
	return ((q1.x == q2.x) ||
		    (q1.y == q2.y) ||
		    (q1.x + q1.y == q2.x + q2.y) ||
		    (q1.x - q1.y == q2.x - q2.y))
}

func CanQueenAttack(w, b string) (bool, error) {
	var err error
	var queen1, queen2 *Queen
	queen1, err = newQueen(w)
	if err != nil {
		return false, err
	}
	queen2, err = newQueen(b)
	if err != nil {
		return false, err
	}
	if *queen1 == *queen2 {
		return false, errors.New("Same square")
	}
	return canAttack(queen1, queen2), nil
}
