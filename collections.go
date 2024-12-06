package kettle

func SliceToSet[T comparable](ts []T) map[T]struct{} {
	result := map[T]struct{}{}
	for _, t := range ts {
		result[t] = struct{}{}
	}
	return result
}

func SetToSlice[T comparable](m map[T]struct{}) []T {
	result := make([]T, 0, len(m))
	for k := range m {
		result = append(result, k)
	}
	return result
}

// OverwriteMerge merge string maps
// if m1 and m2 has same key, m2's value will overwrite m1's value in the result map
func OverwriteMerge(m1 map[string]string, m2 map[string]string) map[string]string {
	result := map[string]string{}
	for k, v := range m1 {
		result[k] = v
	}
	for k, v := range m2 {
		result[k] = v
	}
	return result
}

func CovertMapT[T1 any, T2 any](m1 map[string]T1, converter func(T1) T2) map[string]T2 {
	result := map[string]T2{}
	for k, v := range m1 {
		result[k] = converter(v)
	}
	return result
}

func CovertSliceT[T1 any, T2 any](s1 []T1, converter func(T1) T2) []T2 {
	var result []T2
	for _, v := range s1 {
		result = append(result, converter(v))
	}
	return result
}

func Contains[T comparable](s []T, v T) bool {
	for _, _v := range s {
		if _v == v {
			return true
		}
	}
	return false
}

func Removes[T comparable](s []T, v T) []T {
	var result []T
	for _, _v := range s {
		if _v != v {
			result = append(result, _v)
		}
	}
	return result
}
