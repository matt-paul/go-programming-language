package allinclusive

func appendIf(a []string, s string) []string {
	for _, el := range a {
		if el == s {
			return a
		}
	}
	return append(a, s)
}

func rotations(uu string) []string {
	res := []string{uu}
	rot := uu
	for i := 1; i < len(uu); i++ {
		rot := rot[i:] + rot[:i]
		res = appendIf(res, rot)
	}
	return res
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// ContainAllRots check that each rotation of the string is included
func ContainAllRots(strng string, arr []string) bool {
	if len(strng) == 0 {
		return true
	}
	rots := rotations(strng)
	for _, el := range rots {
		if contains(arr, el) == false {
			return false
		}
	}
	return true
}
