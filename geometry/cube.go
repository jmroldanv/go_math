package geometry

import "errors"

// CubeVolume calculates volumen of an int
func CubeVolume(n int) (int, error) {
	if n != 0 {
		return n * n * n, nil
	}

	return 0, errors.New("Zero length edge is not allowed")
}
