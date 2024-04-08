package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Lotto struct {
	num1 int
	num2 int
	num3 int
	num4 int
	num5 int
	num6 int
}

type myForm struct {
	Ticket int `form:"ticket"`
}

type PageData struct {
	Spend  string
	Prize  string
	Profit string
}

var data PageData

func main() {
	// Seed를 설정하여 매번 다른 난수가 생성되도록 함

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title":  "Main website",
			"spend":  data.Spend,
			"prize":  data.Prize,
			"profit": data.Profit,
		})
	})

	router.POST("/", func(c *gin.Context) {
		var fakeForm myForm
		c.ShouldBind(&fakeForm)
		// c.JSON(200, gin.H{"ticket": fakeForm.Ticket})
		runLotti(fakeForm.Ticket)
		c.Redirect(http.StatusFound, "/")

	})

	router.Run(":8080")

}

func runLotti(ticket int) {

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
	formatSpend := formatNumber(totalSpend)
	fmt.Println("Total Spend: ", formatSpend)
	formattPrize := formatNumber(totalPrice)
	fmt.Println("Your total Price is : ", formattPrize)

	total := totalPrice - totalSpend
	formattedNum := formatNumber(total)

	fmt.Println("your Profit is: ", formattedNum)
	data = PageData{Prize: formattPrize, Spend: formatSpend, Profit: formattedNum}

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
