package helperfunc

type TimeInterval struct {
	From string
	To   string
}

func MaxOftoInt(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func MinOftoInt(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
