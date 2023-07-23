package stack

import (
	"sort"
	"strings"
)

func braceExpansionII(expression string) []string {
	op := []byte{}
	st := stack{}
	for i := 0; i < len(expression); i++ {
		c := expression[i]
		if i == 0 {
			if c >= 'a' && c <= 'z' {
				st.push(string(c))
			} else {
				op = append(op, c)
			}
			continue
		}
		if c == '}' {
			for op[len(op)-1] != '{' {
				s1 := st.pop()
				s2 := st.pop()
				s := ""
				if op[len(op)-1] == '*' {
					s = multi(s1, s2)
				} else {
					s = add(s1, s2)
				}
				st.push(s)
				op = op[:len(op)-1]
			}
			op = op[:len(op)-1]
			continue
		}
		if c == '{' {
			if expression[i-1] == '}' {
				op = append(op, '*')
			}
			if expression[i-1] <= 'z' && expression[i-1] >= 'a' {
				op = append(op, '*')
			}
			op = append(op, c)
			continue
		}
		if c >= 'a' && c <= 'z' {
			if expression[i-1] == '}' {
				op = append(op, '*')
			}
			if expression[i-1] >= 'a' && expression[i-1] <= 'z' {
				op = append(op, '*')
			}
			st = append(st, string(c))
			continue
		}
		for op[len(op)-1] == '*' {
			s1 := st.pop()
			s2 := st.pop()
			s := multi(s1, s2)
			st.push(s)
			op = op[:len(op)-1]
		}
		op = append(op, '+')
	}
	for len(op) > 0 {
		c := op[len(op)-1]
		s1 := st.pop()
		s2 := st.pop()
		s := ""
		if c == '*' {
			s = multi(s1, s2)
		} else {
			s = add(s1, s2)
		}
		st.push(s)
		op = op[:len(op)-1]
	}
	arr := strings.Split(st[0], ",")
	m := map[string]bool{}
	for _, v := range arr {
		m[v] = true
	}
	arr = []string{}
	for k := range m {
		arr = append(arr, k)
	}
	sort.Strings(arr)
	return arr
}

type stack []string

func (st *stack) push(s string) {
	*st = append(*st, s)
}
func (st *stack) pop() string {
	old := *st
	l := len(old)
	if l == 0 {
		return ""
	}
	s := old[l-1]
	*st = old[:l-1]
	return s
}

func multi(s1, s2 string) string {
	arr1 := strings.Split(s2, ",")
	arr2 := strings.Split(s1, ",")
	arr := []string{}
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			arr = append(arr, arr1[i]+arr2[j])
		}
	}
	return strings.Join(arr, ",")
}

func add(s1, s2 string) string {
	return s2 + "," + s1
}
