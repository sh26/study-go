package code

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// https://school.programmers.co.kr/learn/courses/30/lessons/120880
// 특이한 정렬
func q120880() {

	solution := func(numlist []int, n int) (nearlist []int) {

		sort.Ints(numlist)

		numsort := map[int]interface{}{}

		for _, num := range numlist {
			diff := int(math.Abs(float64(num - n)))

			if _, ok := numsort[diff]; !ok {
				numsort[diff] = []int{}
			}
			numsort[diff] = append(numsort[diff].([]int), num)
		}

		for idx := 0; ; idx++ {
			if _, ok := numsort[idx]; !ok {
				continue
			}

			sort.Sort(sort.Reverse(sort.IntSlice(numsort[idx].([]int))))

			nearlist = append(nearlist, numsort[idx].([]int)...)

			if len(nearlist) == len(numlist) {
				break
			}
		}

		return nearlist
	}

	fmt.Println(solution([]int{1, 2, 3, 4, 5, 6}, 4))                    // [4, 5, 3, 6, 2, 1]
	fmt.Println(solution([]int{10000, 20, 36, 47, 40, 6, 10, 7000}, 30)) // [36, 40, 20, 47, 10, 6, 7000, 10000]
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120882
// 등수 매기기
func q120882() {

	solution := func(score ...[]int) (ranks []int) {

		aves := []float64{}
		rank := make(map[float64]int)

		for _, student := range score {
			sum := float64(student[0] + student[1])
			aves = append(aves, sum)
			if sum != 0 {
				aves[len(aves)-1] /= 2
			}
		}

		sort.Sort(sort.Reverse(sort.Float64Slice(aves)))

		for idx := len(score) - 1; idx >= 0; idx-- {
			rank[aves[idx]] = idx + 1

			if idx > 0 {
				if aves[idx] == aves[idx-1] {
					rank[aves[idx-1]] = idx
					idx--
				}
			}
		}

		for _, student := range score {
			if sum := float64(student[0] + student[1]); sum > 0 {
				ranks = append(ranks, rank[sum/2])
			} else {
				ranks = append(ranks, rank[0])
			}
		}

		return
	}

	fmt.Println(solution([]int{80, 70}, []int{90, 50}, []int{40, 70}, []int{50, 80}))                                                  // [1, 2, 4, 3]
	fmt.Println(solution([]int{80, 70}, []int{70, 80}, []int{30, 50}, []int{90, 100}, []int{100, 90}, []int{100, 100}, []int{10, 30})) // [4, 4, 6, 2, 2, 1, 7]
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120956
// 옹알이 (1)
func q120956() {

	solution := func(babbling []string) (res int) {
		rep := strings.NewReplacer("aya", "", "ye", "", "woo", "", "ma", "")

		for _, baa := range babbling {
			if rep.Replace(baa) == "" {
				res++
			}
		}

		return
	}

	fmt.Println(solution([]string{"aya", "yee", "u", "maa", "wyeoo"}))         // 1
	fmt.Println(solution([]string{"ayaye", "uuuma", "ye", "yemawoo", "ayaa"})) // 3
}

// https://school.programmers.co.kr/learn/courses/30/lessons/120883
// 로그인 성공?
func q120883() {

	solution := func(id_pw []string, db ...[]string) string {

		for _, user := range db {
			if id_pw[0] == user[0] {
				if id_pw[1] == user[1] {
					return "login"
				}
				return "wrong pw"
			}
		}
		return "fail"
	}

	fmt.Println(solution([]string{"meosseugi", "1234"}, []string{"rardss", "123"}, []string{"yyoom", "1234"}, []string{"meosseugi", "1234"}))                       // "login"
	fmt.Println(solution([]string{"programmer01", "15789"}, []string{"programmer02", "111111"}, []string{"programmer00", "134"}, []string{"programmer01", "1145"})) // "wrong pw"
	fmt.Println(solution([]string{"rabbit04", "98761"}, []string{"jaja11", "98761"}, []string{"krong0313", "29440"}, []string{"rabbit00", "111333"}))               // "fail"
}
