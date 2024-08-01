package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m, t int

	fmt.Fscanf(in, "%d %d %d\n", &n, &m, &t)

	mp := make([][]int, n+10)
	bp := make([][]int, n+10)

	for i := 0; i < n; i++ {
		mp[i] = make([]int, m+10)
		bp[i] = make([]int, m+10)

		up := m - 1
		for j := 0; j < up; j++ {
			fmt.Fscanf(in, "%d", &mp[i][j])
		}
		fmt.Fscanf(in, "%d\n", &mp[i][up])
	}

	bp[n] = make([]int, m+10)

	getBP(mp, bp, n, m)

	var x1, y1, x2, y2, c int
	for i := 0; i < t; i++ {
		fmt.Fscanf(in, "%d %d %d %d %d\n", &x1, &y1, &x2, &y2, &c)
		insert(bp, x1-1, y1-1, x2-1, y2-1, c)
	}

	printMP(bp, n, m)
}

func getBP(mp, bp [][]int, n, m int) {
	for i := 0; i < n; i++ {
		if i == 0 {
			bp[0][0] = mp[0][0]
			for j := 1; j < m; j++ {
				bp[0][j] = mp[0][j] - mp[0][j-1]
			}
		} else {
			bp[i][0] = mp[i][0] - mp[i-1][0]
			for j := 1; j < m; j++ {
				bp[i][j] = mp[i][j] - mp[i-1][j] - mp[i][j-1] + mp[i-1][j-1]
			}
		}
	}
}

func insert(bp [][]int, x1, y1, x2, y2, c int) {
	bp[x1][y1] += c
	bp[x1][y2+1] -= c
	bp[x2+1][y1] -= c
	bp[x2+1][y2+1] += c
}

func printMP(bp [][]int, n, m int) {

	for i := 0; i < n; i++ {
		if i == 0 {
			fmt.Printf("%d", bp[0][0])
			for j := 1; j < m; j++ {
				bp[0][j] += bp[0][j-1]
				fmt.Printf(" %d", bp[0][j])
			}
			fmt.Println()
		} else {
			bp[i][0] += bp[i-1][0]
			fmt.Printf("%d", bp[i][0])
			for j := 1; j < m; j++ {
				bp[i][j] += bp[i-1][j] + bp[i][j-1] - bp[i-1][j-1]
				fmt.Printf(" %d", bp[i][j])
			}
			fmt.Println()
		}
	}
}
