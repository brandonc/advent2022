package day11

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type position byte

type parser struct {
	pos    position
	input  string
	monkey *monkey
}

var (
	header        position = 0
	startingItems position = 1
	operation     position = 2
	test          position = 3
	testPass      position = 4
	testFail      position = 5
)

func parse(reader io.Reader) []*monkey {
	p := parser{
		monkey: &monkey{},
	}
	scanner := bufio.NewScanner(reader)
	monkeys := make([]*monkey, 0, 8)

	for scanner.Scan() {
		p.input = scanner.Text()
		if len(p.input) == 0 {
			continue
		}

		switch p.pos {
		case header:
			p.pos = startingItems
		case startingItems:
			itemsRaw := strings.Split(p.input[len("  Starting items: "):], ", ")
			p.monkey.items = make([]int, len(itemsRaw))
			for i, r := range itemsRaw {
				item, _ := strconv.Atoi(r)
				p.monkey.items[i] = item
			}
			p.pos = operation
		case operation:
			op := p.input[len("  Operation: new = old ")]

			operandRaw := p.input[strings.LastIndexByte(p.input, op)+2:]
			if operandRaw != "old" {
				x, _ := strconv.Atoi(operandRaw)
				switch op {
				case '*':
					p.monkey.operation = mult(x)
				case '+':
					p.monkey.operation = add(x)
				default:
					panic(fmt.Sprintf("invalid op %c", op))
				}
			} else {
				p.monkey.operation = square
			}
			p.pos = test
		case test:
			raw := p.input[len("  Test: divisible by "):]
			test, _ := strconv.Atoi(raw)
			p.monkey.test = test
			p.pos = testPass
		case testPass:
			to, _ := strconv.Atoi(p.input[len("    If true: throw to monkey "):])
			p.monkey.ifTrueMonkey = to
			p.pos = testFail
		case testFail:
			to, _ := strconv.Atoi(p.input[len("    If false: throw to monkey "):])
			p.monkey.ifFalseMonkey = to
			p.pos = header
			monkeys = append(monkeys, p.monkey)
			p.monkey = &monkey{}
		}
	}
	return monkeys
}
