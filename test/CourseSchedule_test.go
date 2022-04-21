package main

import (
	"fmt"
	"testing"
)

func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edge   = make([][]int, numCourses)
		visit  = make([]int, numCourses)
		result []int
		link   bool
		dfs    func(u int)
	)
	// 0表示未搜索，1表示搜索中，2表示搜索完成
	dfs = func(u int) {
		visit[u] = 1
		for _, v := range edge[u] {
			// 未搜索
			if visit[v] == 0 {
				dfs(v)
				if link == false {
					return
				}
				// 搜索中，遇到环，无法拓扑排序
			} else if visit[v] == 1 {
				link = false
				return
			}
		}
		visit[u] = 2
		result = append(result, u)
	}
	link = true
	// 建立无向图
	for i := 0; i < len(prerequisites); i++ {
		edge[prerequisites[i][1]] = append(edge[prerequisites[i][1]], prerequisites[i][0])
	}
	// 从入度为0的点开始递归
	for i := 0; i < numCourses && link; i++ {
		if visit[i] == 0 {
			dfs(i)
		}
	}
	return link
}

func canFinish1(numCourses int, prerequisites [][]int) bool {
	var (
		edge   = make([][]int, numCourses)
		indeg  = make([]int, numCourses)
		result []int
	)
	// 首先建立无向图
	for i := 0; i < len(prerequisites); i++ {
		edge[prerequisites[i][1]] = append(edge[prerequisites[i][1]], prerequisites[i][0])
		// 计算入度
		indeg[prerequisites[i][0]]++
	}

	// 创建队列
	var q []int

	// 寻找入度为0的点
	for i := 0; i < len(indeg); i++ {
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}

	// 将队列的结果放进结果集，并对相邻的点减少入度
	for len(q) > 0 {
		u := q[0]
		result = append(result, u)
		q = q[1:]
		for _, v := range edge[u] {
			indeg[v]--
			if indeg[v] == 0 {
				q = append(q, v)
			}
		}
	}
	return len(result) == numCourses
}

func TestCourseSchedule(t *testing.T) {
	fmt.Println(canFinish1(2, [][]int{{1, 0}, {0, 1}}))
}
