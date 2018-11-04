package coordinates

import (
	"math"
	"sync"
)

func ParallelSieve(pos int) (ulamPos int) {
	ulamPos = 0
	if pos > 1 {
		numsCnt := pos + 1
		nums := make([]int, numsCnt)
		nums[0] = 1
		nums[1] = 0
		nums[2] = 1
		if pos > 2 {
			nums[3] = 1
		}
		if pos > 4 {
			nums[5] = 1
		}
		stepsDifs := [8]int{4, 2, 4, 2, 4, 6, 2, 6}
		k := 7
		for i := 0; k <= pos; i++ {
			nums[k] = 1
			k += stepsDifs[i % 8]
		}
		primeProgressionSteps := [8]int{7, 11, 13, 17, 19, 23, 29, 31}
		notPrimes := make(chan int, 10)
		var wg sync.WaitGroup
		wg.Add(8)
		for i := range primeProgressionSteps {
			numsCp := make([]int, numsCnt)
			copy(numsCp, nums)
			step := primeProgressionSteps[i]
			go func() {
				defer wg.Done()
				j := 0
				g := 0
				for k := step; int(math.Pow(float64(k), 2)) <= pos; k = 30 * j + step {
					if nums[k] == 1 {
						for l := int(math.Pow(float64(k), 2)); l <= pos; l += k {
							notPrimes <- l
							g++
						}
					}
					j++
				}
				println(j + g)
			}()
		}
		go func() {
			for m := range notPrimes {
				nums[m] = 0
			}
		}()
		wg.Wait()
		if nums[pos] == 1 {
			ulamPos = -1
		} else {
			for i := range nums {
				if nums[i] == 0 {
					ulamPos += 1
				}
			}
		}
	} else {
		ulamPos = pos
	}
	return ulamPos
}

func Sieve(pos int) (ulamPos int) {
	ulamPos = 0
	if pos > 1 {
		var nums = make([]int, pos + 1)
		nums[0] = 1
		nums[1] = 0
		nums[2] = 1
		for i := 3; i <= pos; i++ {
			if i % 2 > 0 {
				nums[i] = 1
			} else {
				nums[i] = 0
			}
		}
		j := 0
		g := 0
		for k := 3; int(math.Pow(float64(k), 2)) <= pos; k += 2 {
			if nums[k] == 1 {
				for l := int(math.Pow(float64(k), 2)); l <= pos; l += 2 * k {
					nums[l] = 0
					g++
				}
			}
			j++
		}
		println(j + g)
		if nums[pos] == 1 {
			ulamPos = -1
		} else {
			for i := range nums {
				if nums[i] == 0 {
					ulamPos += 1
				}
			}
		}
	} else {
		ulamPos = pos
	}
	return ulamPos
}

func UlamToCartesian(pos int) (x, y int) {
	closestSqrt := int(math.Floor(math.Sqrt(float64(pos))))
	minOddSqrt := closestSqrt
	if closestSqrt != 0 && closestSqrt % 2 == 0 {
		minOddSqrt = closestSqrt - 1
	}
	if minOddSqrt == 0 {
		x = 0
		y = 0
	} else {
		contourNum := (minOddSqrt + 1) / 2
		minOdd := int(math.Pow(float64(minOddSqrt), float64(2)))
		contourPos := pos - minOdd + 1
		secondCoord := contourPos % contourNum
		if secondCoord == 0 && (pos / contourNum) % 2 == 0 {
			secondCoord = contourNum
		}
		if pos >= minOdd && pos < minOdd + contourNum {
			x = -(contourNum - secondCoord)
			if secondCoord == 0 {
				x = 0
			}
			y = contourNum
		} else if pos >= minOdd + contourNum && pos < minOdd + 2 * contourNum {
			x = secondCoord
			y = contourNum
		} else if pos >= minOdd + 2 * contourNum && pos < minOdd + 3 * contourNum {
			x = contourNum
			y = contourNum - secondCoord
			if secondCoord == 0 {
				y = 0
			}
		} else if pos >= minOdd + 3 * contourNum && pos < minOdd + 4 * contourNum {
			x = contourNum
			y = -secondCoord
		} else if pos >= minOdd + 4 * contourNum && pos < minOdd + 5 * contourNum {
			x = contourNum - secondCoord
			if secondCoord == 0 {
				x = 0
			}
			y = -contourNum
		} else if pos >= minOdd + 5 * contourNum && pos < minOdd + 6 * contourNum {
			x = -secondCoord
			y = -contourNum
		} else if pos >= minOdd + 6 * contourNum && pos < minOdd + 7 * contourNum {
			x = -contourNum
			y = -(contourNum - secondCoord)
			if secondCoord == 0 {
				y = 0
			}
		} else if pos >= minOdd + 7 * contourNum && pos < minOdd + 8 * contourNum {
			x = -contourNum
			y = secondCoord
		}
	}
	return x, y
}
