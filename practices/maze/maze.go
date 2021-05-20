package main

import (
	"fmt"
	"os"
)

// 从左上角进，从右下角出迷宫
// 广度优先搜索，走到一个格子，看一下这个格子的上左下右
// 结束条件：
// 1. 走到终点即结束
// 2. 全是死路亦结束

// 把数据读到迷宫 slice 里
func readMaze(filepath string) [][]int {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)

		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

// 走的格子（到达的点）
type point struct {
	i, j int
}

// 走的方向
var directions = [4]point{
	{-1, 0}, // 上
	{0, -1}, // 左
	{1, 0},  // 下
	{0, 1},  // 右
}

// 点加法
func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

// 点所以的位置的值
func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}

	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}

// 走迷宫
func walk(maze [][]int, start point, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	queue := []point{start}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// 达到终点
		if current == end {
			break
		}

		for _, dir := range directions {
			next := current.add(dir)

			// maze at next is 0
			// and steps at next is 0
			// next != start

			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}

			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}

			if next == start {
				continue
			}

			currSteps, _ := current.at(steps)
			steps[next.i][next.j] = currSteps + 1

			queue = append(queue, next)
		}
	}

	return steps
}

func main() {
	pwd, _ := os.Getwd()
	fp := pwd + string(os.PathSeparator) + "maze.dat"
	fmt.Printf("data file: %s\n", fp)
	maze := readMaze(fp)

	fmt.Println("original maze:")
	for _, row := range maze {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}

	fmt.Println()

	start := point{0, 0}
	end := point{len(maze) - 1, len(maze[0]) - 1}
	steps := walk(maze, start, end)

	fmt.Println("steps:")
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}
