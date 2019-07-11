package arrays

func AddToArrayForm(A []int, K int) []int {
	A = reverseArray(A)
	k := int2ReverseArray(K)
	p := len(A)
	q := len(k)
	r := []int{}
	addon := 0
	i := 0
	j := 0
	for i < p && j < q {
		s := A[i] + k[j] + addon
		r = append(r, s%10)
		addon = s / 10
		i++
		j++
	}
	for i < p {
		s := A[i] + addon
		r = append(r, s%10)
		addon = s / 10
		i++
	}
	for j < q {
		s := k[j] + addon
		r = append(r, s%10)
		addon = s / 10
		j++
	}
	if addon > 0 {
		r = append(r, addon)
	}
	res := reverseArray(r)
	return res
}

func int2ReverseArray(n int) []int {
	queue := []int{}
	for {
		if n < 10 {
			queue = append(queue, n)
			break
		} else {
			queue = append(queue, n%10)
			n = n / 10
		}
	}
	return queue
}

func reverseArray(queue []int) []int {
	l := len(queue)
	for i := 0; i < l/2; i++ {
		queue[i], queue[l-1-i] = queue[l-1-i], queue[i]
	}
	return queue
}
