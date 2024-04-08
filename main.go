package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
)

type Lotto struct {
	num1 int
	num2 int
	num3 int
	num4 int
	num5 int
	num6 int
}

func main() {
	// Seed를 설정하여 매번 다른 난수가 생성되도록 함
	var ticket int
	fmt.Println("Lotto How many ticket do you want to buy?")
	fmt.Scan(&ticket)

	userTickets := PickLottoNumber(ticket)
	var totalPrice int

	winningTicket := PickLottoNumber(1)[0]
	fmt.Println("Winning Ticket:", winningTicket)

	// 사용자 티켓과 당첨 티켓 비교
	for i, userTicket := range userTickets {
		matches := CompareLottoNumbers(userTicket, winningTicket)
		i++
		if matches == 4 {
			fmt.Printf("Ticket %d Matches: %d\n", i, matches)
			totalPrice += 50000
		}
		if matches == 5 {
			fmt.Printf("Ticket %d Matches: %d\n", i, matches)
			totalPrice += 10000000
		}
		if matches == 6 {
			fmt.Printf("Ticket %d Matches: %d\n", i, matches)
			totalPrice += 2000000000
		}
	}
	var totalSpend = ticket * 1000
	fmt.Println("Total Spend: ", totalSpend)
	fmt.Println("Your total Price is : ", totalPrice)

	total := totalPrice - totalSpend
	formattedNum := formatNumber(total)

	fmt.Println("your Profit is: ", formattedNum)

}
func formatNumber(num int) string {
	str := strconv.FormatInt(int64(num), 10) // 숫자를 문자열로 변환
	length := len(str)

	// 음수일 경우 맨 앞의 "-"는 유지하고 숫자 부분에만 세 자리마다 쉼표(,) 추가
	start := 0
	if num < 0 {
		start = 1
	}

	if length-start > 3 {
		for i := length - 3; i > start; i -= 3 {
			str = str[:i] + "," + str[i:]
		}
	}

	return str
}

func condition(num int, winningTicket Lotto) bool {
	return num == winningTicket.num1 ||
		num == winningTicket.num2 ||
		num == winningTicket.num3 ||
		num == winningTicket.num4 ||
		num == winningTicket.num5 ||
		num == winningTicket.num6
}

func CompareLottoNumbers(userTicket, winningTicket Lotto) int {
	matches := 0

	if condition(userTicket.num1, winningTicket) {
		matches++
	}
	if condition(userTicket.num2, winningTicket) {
		matches++
	}
	if condition(userTicket.num3, winningTicket) {
		matches++
	}
	if condition(userTicket.num4, winningTicket) {
		matches++
	}
	if condition(userTicket.num5, winningTicket) {
		matches++
	}
	if condition(userTicket.num6, winningTicket) {
		matches++
	}

	return matches
}

func PickLottoNumber(i int) []Lotto {
	picNum := make([]Lotto, i)
	for j := 0; j < i; j++ {

		var nums [6]int
		done := false

		for !done {
			// 1부터 35까지의 난수 생성하여 nums 배열에 저장
			for k := 0; k < 6; k++ {
				nums[k] = rand.Intn(35) + 1
			}

			// 배열을 오름차순으로 정렬
			sort.Ints(nums[:]) // 배열을 오름차순으로 정렬

			// 중복을 제거하기 위해 중복 여부 확인
			if !hasDuplicates(nums) {
				done = true
			}
		}
		lotto := Lotto{num1: nums[0], num2: nums[1], num3: nums[2], num4: nums[3], num5: nums[4], num6: nums[5]}

		picNum[j] = lotto
	}

	return picNum
}

// 중복 확인 함수: 배열 내에 중복된 요소가 있는지 확인
func hasDuplicates(arr [6]int) bool {
	encountered := map[int]bool{}

	for _, v := range arr {
		if encountered[v] {
			return true
		}
		encountered[v] = true
	}

	return false
}
