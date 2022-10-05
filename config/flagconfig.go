package config

import "flag"

type FlagConfig struct {
	K int

	N bool
	R bool
	U bool
	B bool
	C bool
	M bool
}

func Export(flag *flag.FlagSet) func() *FlagConfig {
	var c FlagConfig

	flag.IntVar(&c.K, "k", 1, "номер столбца для сортировки, не может быть меньше 1")

	flag.BoolVar(&c.N, "n", false, "cортировать по числовому значению")
	flag.BoolVar(&c.R, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&c.U, "u", false, "не выводить повторяющиеся строки")
	flag.BoolVar(&c.B, "b", false, "игнорировать начальные пробелы")
	flag.BoolVar(&c.C, "c", false, "проверить отсортированы ли данные")
	flag.BoolVar(&c.M, "M", false, "отсортировать по месяцам")

	return func() *FlagConfig {
		c.K -= 1

		if c.K < 0 {
			c.K = 0
		}

		if c.M && c.N {
			c.M = false
		}

		return &c
	}
}
