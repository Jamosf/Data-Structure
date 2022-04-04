package days

// tag-[��ѧ]
// ÿ��һ��
// leetcode2013: ���������
type DetectSquares struct {
	p map[int]map[int]int
}

func ConstructorDetectSquares() DetectSquares {
	return DetectSquares{make(map[int]map[int]int)}
}

func (d *DetectSquares) Add(point []int) {
	x, y := point[0], point[1]
	if _, ok := d.p[x]; !ok {
		d.p[x] = make(map[int]int)
	}
	d.p[x][y]++
}

func (d *DetectSquares) Count(point []int) int {
	x, y := point[0], point[1]
	if _, ok := d.p[x]; !ok {
		return 0
	}
	ans := 0
	for y1, c := range d.p[x] {
		if y != y1 {
			ans += c * d.p[x+abs(y-y1)][y] * d.p[x+abs(y-y1)][y1]
			ans += c * d.p[x-abs(y-y1)][y] * d.p[x-abs(y-y1)][y1]
		}
	}
	return ans
}
