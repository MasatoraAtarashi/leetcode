package main

type TimeMap struct {
	m map[string][]TsValue
}

type TsValue struct {
	timestamp int
	value     string
}

func _Constructor() TimeMap {
	return TimeMap{
		m: map[string][]TsValue{},
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.m[key] = append(this.m[key], TsValue{
		timestamp: timestamp,
		value:     value,
	})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	vs := this.m[key]
	i := 0
	j := len(vs) - 1
	val := ""
	for i <= j {
		mid := (i + j) / 2
		if vs[mid].timestamp == timestamp {
			val = vs[mid].value
			break
		} else if vs[mid].timestamp < timestamp {
			val = vs[mid].value
			i = mid + 1
		} else {
			j = mid - 1
		}
	}

	return val
}
