package pb

type Entity interface {
	MarkDirty(idx int, dt byte)
	MarkAllDirty(dt byte)
	BuildEntityInfo()
	SetParent(entity Entity, idx int, dt byte)
}

func RemoveIndexOf[T any](slice []T, idx int) []T {
	return append(slice[:idx], slice[idx+1:]...)
}
func RemoveValueOf[T comparable](slice []T, v T) []T {
	ret := make([]T, 0)
	for _, vv := range slice {
		if vv != v {
			ret = append(ret, vv)
		}
	}
	return ret
}

func MergeNotExistInto[T comparable](dest []T, src []T) []T {
	ret := make([]T, len(dest))
	copy(ret, dest)
	for _, s := range src {
		var exist = false
		for _, d := range dest {
			if s == d {
				exist = true
			}
		}
		if !exist {
			ret = append(ret, s)
		}
	}
	return ret
}
func AppendIfNotExist[T comparable](set []T, v T) []T {
	var exist = false
	for _, s := range set {
		if s == v {
			exist = true
		}
	}
	if !exist {
		return append(set, v)
	}
	return set
}
