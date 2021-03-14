package heap

type minInt int

func (s minInt) Less(v minInt) bool {
	return s < v
}

type maxInt int

func (s maxInt) Less(v maxInt) bool {
	return s > v
}
