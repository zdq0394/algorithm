package main

func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		result[i] = make([]int, i+1)
		result[i][0] = 1
		result[i][i] = 1
		for j := 1; j <= i-1; j++ {
			result[i][j] = result[i-1][j-1] + result[i-1][j]
		}
	}
	return result
}

func getRow(rowIndex int) []int {
	var pre, next []int
	pre = []int{1}
	for i := 1; i <= rowIndex; i++ {
		next = make([]int, i+1)
		next[0] = 1
		next[i] = 1
		for j := 1; j <= i-1; j++ {
			next[j] = pre[j-1] + pre[j]
		}
		pre = next
	}
	return pre
}

func main() {

}
