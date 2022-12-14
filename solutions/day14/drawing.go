package day14

func (s scan) drawHorizontal(from, to position, e entity) {
	start, end := to.x, from.x
	if start > end {
		end, start = start, end
	}
	for i := start; i <= end; i++ {
		s.entities[position{from.y, i}] = e
	}
}

func (s scan) drawVertical(from, to position, e entity) {
	start, end := to.y, from.y
	if start > end {
		end, start = start, end
	}
	for i := start; i <= end; i++ {
		s.entities[position{i, from.x}] = e
	}
}

func (s scan) drawLine(from, to position, e entity) {
	if from.x != to.x {
		s.drawHorizontal(from, to, e)
	} else if from.y != to.y {
		s.drawVertical(from, to, e)
	} else {
		s.entities[from] = e
	}
}
