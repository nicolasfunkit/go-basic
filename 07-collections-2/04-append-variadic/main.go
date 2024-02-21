package main

func Merge(a []string, b []string) []string {
	return append(a, b...)
}

func Remove(a []string, index int) []string {
	if index >= len(a) || index < 0 {
		return a
	}

	if index == 0 {
		return a[1:]
	}
	if index == len(a)-1 {
		return a[:(index - 1)]
	}

	return append(a[:(index)], a[(index+1):]...)
}

func RemoveLast(a []string) []string {
	count := len(a)
	if count == 0 {
		return a
	}

	return a[:(count - 1)]
}
