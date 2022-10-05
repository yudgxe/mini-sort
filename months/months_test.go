package months_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/yudgxe/mini-sort/months"
)

var data = `March
February
Apr
April
August
July
June
November
September
1
4
03
06
01/03/19
01/08/19
02/09/18`

func TestParse(t *testing.T) {
	expected := []int{3, 2, 4, 4, 8, 7, 6, 11, 9, 1, 4, 3, 6, -1, -1, -1}

	r := strings.NewReader(data)
	sc := bufio.NewScanner(r)

	i := 0
	for sc.Scan() {
		m, _ := months.Parse(sc.Text())
		if m != expected[i] {
			t.Errorf("Incorrect result, Expect %d, got %d", expected[i], m)
		}
		i++
	}
}
