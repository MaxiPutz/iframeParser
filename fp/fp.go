package fp

func Map(arr []string, f func(ele string) string) []string {
	for i, ele := range arr {
		arr[i] = f(ele)
	}
	return arr
}
