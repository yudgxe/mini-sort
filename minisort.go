package minisort

import (
	"bufio"
	"io"
	"sort"
	"strconv"
	"strings"

	"github.com/yudgxe/mini-sort/config"
	"github.com/yudgxe/mini-sort/cstrings"
	"github.com/yudgxe/mini-sort/months"
)

type Table struct {
	Matrix [][]string

	c *config.FlagConfig
}

func New(c *config.FlagConfig) *Table {
	return &Table{
		c: c,
	}
}

func (t *Table) len(col int) int {
	var l int = 0

	for i := 0; i < len(t.Matrix); i++ {
		if len(t.Matrix[i]) > col {
			l++
		}
	}

	return l
}

func (t *Table) Len() int {
	return t.len(0)
}

func (t *Table) Less(i, j int) bool {
	v1 := t.GetValue(i, t.c.K)
	v2 := t.GetValue(j, t.c.K)

	condition := v1 < v2

	f := func(f func(string) (int, error)) {
		v1, err1 := f(v1)
		v2, err2 := f(v2)

		if err1 == nil && err2 == nil {
			condition = v1 < v2
		}
	}

	if t.c.M {
		f(months.Parse)
	}

	if t.c.N {
		f(func(s string) (int, error) {
			v, err := strconv.Atoi(strings.TrimLeft(s, " "))
			if err != nil {
				return 0, err
			}

			return v, nil
		})
	}

	if t.c.R {
		return !condition
	}

	return condition
}

func (t *Table) GetValue(i, j int) string {
	if j > len(t.Matrix[i])-1 {
		return ""
	}

	if t.c.B {
		return strings.TrimLeft(t.Matrix[i][j], " ")
	}

	return t.Matrix[i][j]
}

func (t *Table) Swap(i, j int) {
	t.Matrix[i], t.Matrix[j] = t.Matrix[j], t.Matrix[i]
}

func (t *Table) Read(r io.Reader, sep string) {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		strs := cstrings.Split(scanner.Text(), sep)
		t.Matrix = append(t.Matrix, strs)
	}
}

func (t *Table) Write(w io.Writer) {
	var lastStr string

	for i, row := range t.Matrix {
		v := t.GetValue(i, t.c.K)
		if v == lastStr && t.c.U && i != 0 {
			continue
		}
		lastStr = v

		for _, v := range row {
			w.Write([]byte(v + " "))
		}
		w.Write([]byte("\n"))
	}
}

func (t *Table) Sort() {
	sort.Sort(t)
}

func (t *Table) Check() int {
	for i := 0; i < t.Len()-1; i++ {
		if t.Less(i+1, i) {
			return i + 1
		}
	}

	return -1
}

func (t *Table) GetRow(i int) []string {
	if i < t.Len() {
		return t.Matrix[i]
	}

	return nil
}
