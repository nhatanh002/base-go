package utils

func Map[A any, B any](f func(a A) B, amap []A) []B {
	bmap := make([]B, len(amap))
	for k, v := range amap {
		bmap[k] = f(v)
	}
	return bmap
}

func Reduce[A any, B any](f func(b B, a A) B, seed B, amap []A) B {
	b := seed
	for _, v := range amap {
		b = f(b, v)
	}
	return b
}

func Filter[A any](f func(a A) bool, amap []A) []A {
	var res []A
	for _, v := range amap {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}

func TakeWhile[A any](f func(a A) bool, amap []A) []A {
	var res []A
	for _, v := range amap {
		if f(v) {
			res = append(res, v)
		} else {
			break
		}
	}
	return res
}

func DropWhile[A any](f func(a A) bool, amap []A) []A {
	var res []A
	for i, v := range amap {
		if !f(v) {
			res = append(res, amap[i:]...)
			break
		}
	}
	return res
}

func Member[A comparable](a A, amap []A) bool {
	for _, v := range amap {
		if v == a {
			return true
		}
	}
	return false
}

func Deduplicate[A comparable](alist []A) []A {
	allKeys := make(map[A]bool)
	list := []A{}
	for _, item := range alist {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func Compose[A, B, C any](fbc func(B) C, fab func(A) B) func(A) C {
	return func(a A) C {
		b := fab(a)
		return fbc(b)
	}
}

func Flip[A, B, C any](f func(A, B) C) func(B, A) C {
	return func(b B, a A) C {
		return f(a, b)
	}
}

func Pipe[A, B, C any](fab func(A) B, fbc func(B) C) func(A) C {
	return (Flip(Compose[A, B, C]))(fab, fbc)
}

func GroupBy[A any, K comparable](partition func(a A) K, alist []A) map[K]([]A) {
	m := make(map[K]([]A))
	for _, a := range alist {
		k := partition(a)
		m[k] = append(m[k], a)
	}
	return m
}
