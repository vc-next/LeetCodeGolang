package simple

func IsValid(s string) bool {
	max := len(s)
	if max%2 != 0 {
		return false
	}

	type Pointer struct {
		flags []uint8
		size  int
	}
	pointer := Pointer{[]uint8{}, 0}

	for i := 0; i < max; i++ {
		flag := s[i]
		if flag == '{' || flag == '(' || flag == '[' {
			pointer.flags = append(pointer.flags, flag)
			pointer.size += 1
		} else {
			if pointer.size < 1 {
				return false
			}
			switch flag {
			case '}':
				if pointer.flags[pointer.size-1] != '{' {
					return false
				}
				break
			case ']':
				if pointer.flags[pointer.size-1] != '[' {
					return false
				}
				break
			case ')':
				if pointer.flags[pointer.size-1] != '(' {
					return false
				}
				break
			default:
				return false
			}
			pointer.flags = pointer.flags[:pointer.size-1]
			pointer.size -= 1
		}
	}
	return pointer.size == 0
}
