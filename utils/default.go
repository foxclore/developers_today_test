package utils

func UnpackDefault[K any](arr []K, defVal K) K {
	if len(arr) > 0 {
		return arr[0]
	}
	return defVal
}
