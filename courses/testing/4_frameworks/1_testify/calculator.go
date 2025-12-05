package frameworks_testify

type Calculator struct {
}

func (c *Calculator) Sum(values ...int64) int64 {
	return 0
}

func (c *Calculator) Multiply(val, by int64) int64 {
	return 0
}

func (c *Calculator) SpecialSub(val, sub int64) int64 {
	if val > 100 {
		if val == 110 {
			return 42
		}
		return val - sub*2
	}
	return val - sub
}
