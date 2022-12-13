package day13

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"reflect"
	"sort"
	"strconv"

	"github.com/brandonc/advent2022/internal/ui"
	"github.com/brandonc/advent2022/solutions/solution"
)

type day13 struct{}

// Factory must exist for codegen
func Factory() solution.Solver {
	return day13{}
}

// packet is a list that contains either more packets or integers
type packet []any

func parsePacket(raw string) (packet, int) {
	ui.Assert(len(raw) > 0 && raw[0] == '[', fmt.Sprintf("Expected list, got %s", raw))
	buf := bytes.NewBuffer(make([]byte, 0, 2))
	result := make(packet, 0)
	i := 1

	appendIntFromBuffer := func() {
		num, err := strconv.Atoi(buf.String())
		ui.Assert(err == nil, fmt.Sprintf("Error parsing buffer '%s' as integer", buf.String()))
		result = append(result, num)
		buf.Reset()
	}

	for {
		switch raw[i] {
		case '[':
			// Recursive case: packet contains a packet
			list, bytesRead := parsePacket(raw[i:])
			result = append(result, list)
			i += bytesRead
		case ',':
			// buffer can contain an int
			if buf.Len() > 0 {
				appendIntFromBuffer()
			}
			i += 1
		case ']':
			// buffer can contain an int
			if buf.Len() > 0 {
				appendIntFromBuffer()
			}
			// Exit case: end of packet
			return result, i + 1
		default:
			buf.WriteByte(raw[i])
			i += 1
		}
	}
}

func result(r bool) *bool {
	return &r
}

func ensurePacket(item any, t reflect.Kind) packet {
	if t != reflect.Slice {
		return packet{item}
	}
	return item.(packet)
}

func (p packet) compare(other packet) bool {
	result := p.comparePacketNext(other)
	ui.Assert(result != nil, "packets should be comparable")
	return *result
}

func (p packet) comparePacketNext(other packet) *bool {
	if len(p) == 0 && len(other) == 0 {
		// No more items to compare, signal that more items need to be compared
		return nil
	}

	if len(p) == 0 {
		// Left side ran out of items
		return result(true)
	}

	if len(other) == 0 {
		// Right side ran out of items
		return result(false)
	}

	leftItem, leftItemType := p[0], reflect.TypeOf(p[0]).Kind()
	rightItem, rightItemType := other[0], reflect.TypeOf(other[0]).Kind()

	// Slice comparison
	if leftItemType == reflect.Slice || rightItemType == reflect.Slice {
		if result := ensurePacket(leftItem, leftItemType).comparePacketNext(ensurePacket(rightItem, rightItemType)); result != nil {
			return result
		}
	}

	// Integer comparison
	if leftItemType == reflect.Int && rightItemType == reflect.Int {
		if leftItem.(int) > rightItem.(int) {
			return result(false)
		}
		if leftItem.(int) < rightItem.(int) {
			return result(true)
		}
	}

	// Compare next item
	return p[1:].comparePacketNext(other[1:])
}

func (d day13) Solve(reader io.Reader) (any, any, error) {
	scanner := bufio.NewScanner(reader)

	dividerPacket1 := &packet{
		packet{packet{2}},
	}

	dividerPacket2 := &packet{
		packet{packet{6}},
	}

	part1 := 0
	index := 0
	allPackets := []*packet{
		dividerPacket1,
		dividerPacket2,
	}

	for {
		ui.Assert(scanner.Scan(), "Expected more input")
		packet1, _ := parsePacket(scanner.Text())
		ui.Assert(scanner.Scan(), "Expected more input")
		packet2, _ := parsePacket(scanner.Text())

		if packet1.compare(packet2) {
			part1 += index + 1
		}

		allPackets = append(allPackets, &packet1, &packet2)

		if !scanner.Scan() {
			break
		}
		index += 1
	}

	sort.Slice(allPackets, func(i, j int) bool {
		return (*allPackets[i]).compare(*allPackets[j])
	})

	part2 := 0
	divider1 := -1
	for i := 0; i < len(allPackets); i++ {
		if allPackets[i] == dividerPacket1 {
			divider1 = i + 1
			continue
		}
		if allPackets[i] == dividerPacket2 {
			ui.Assert(divider1 > 1, "Expected divider packet 1 to be present")
			part2 = divider1 * (i + 1)
			break
		}
	}

	return part1, part2, nil
}
