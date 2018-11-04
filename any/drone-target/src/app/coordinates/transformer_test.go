package coordinates

import (
	"testing"
)

func TestUlamToCartesian(t *testing.T) {
	var x, y int

	// corners

	x, y = UlamToCartesian(36)
	if x != 3 || y != -3 {
		t.Error("Expected 3 -3, got ", x, y)
	}

	x, y = UlamToCartesian(42)
	if x != -3 || y != -3 {
		t.Error("Expected -3 -3, got ", x, y)
	}

	x, y = UlamToCartesian(56)
	if x != 4 || y != 4 {
		t.Error("Expected 4 4, got ", x, y)
	}

	x, y = UlamToCartesian(80)
	if x != -4 || y != 4 {
		t.Error("Expected -4 4, got ", x, y)
	}

	// zero coordinate

	x, y = UlamToCartesian(33)
	if x != 3 || y != 0 {
		t.Error("Expected 3 0, got ", x, y)
	}

	x, y = UlamToCartesian(39)
	if x != 0 || y != -3 {
		t.Error("Expected 0 -3, got ", x, y)
	}

	x, y = UlamToCartesian(45)
	if x != -3 || y != 0 {
		t.Error("Expected -3 0, got ", x, y)
	}

	x, y = UlamToCartesian(52)
	if x != 0 || y != 4 {
		t.Error("Expected 0 4, got ", x, y)
	}

	// other

	x, y = UlamToCartesian(35)
	if x != 3 || y != -2 {
		t.Error("Expected 3 -2, got ", x, y)
	}

	x, y = UlamToCartesian(40)
	if x != -1 || y != -3 {
		t.Error("Expected -1 -3, got ", x, y)
	}

	x, y = UlamToCartesian(127)
	if x != 1 || y != 6 {
		t.Error("Expected 1 6, got ", x, y)
	}

	x, y = UlamToCartesian(77)
	if x != -4 || y != 1 {
		t.Error("Expected 0 4, got ", x, y)
	}
}

func TestSieve(t *testing.T) {
	var ulamCoord int

	ulamCoord = Sieve(0)
	if ulamCoord != 0 {
		t.Error("Expected 0, got ", ulamCoord)
	}

	ulamCoord = Sieve(1)
	if ulamCoord != 1 {
		t.Error("Expected 1, got ", ulamCoord)
	}

	ulamCoord = Sieve(33)
	if ulamCoord != 22 {
		t.Error("Expected 22, got ", ulamCoord)
	}

	ulamCoord = Sieve(55555)
	if ulamCoord != 49918 {
		t.Error("Expected 49918, got ", ulamCoord)
	}

	// Prime numbers

	ulamCoord = Sieve(2)
	if ulamCoord != -1 {
		t.Error("Expected -1, got ", ulamCoord)
	}

	ulamCoord = Sieve(7)
	if ulamCoord != -1 {
		t.Error("Expected -1, got ", ulamCoord)
	}
}
