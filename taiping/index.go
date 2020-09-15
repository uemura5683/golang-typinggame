package main

import (
    "fmt"
    "os"
    "bufio"
)

func main() {
    totalScore := 0
    // 引数にtotalScoreのポインタを渡してください
    ask(1, "dog", &totalScore)
    ask(2, "cat", &totalScore)
    ask(3, "fish", &totalScore)
    ask(4, "Tiger", &totalScore)
    ask(5, "Elephant", &totalScore)
    ask(6, "Crocodile", &totalScore)
    ask(7, "this is a dog", &totalScore)
    ask(8, "Amphibians", &totalScore)
    ask(9, "Butterfly", &totalScore)
    ask(10, "excellentswimmer", &totalScore)

    fmt.Println("スコア", totalScore)
    if totalScore <= 30 {
      fmt.Println("頑張りましょう！")
    } else if totalScore <= 60 && totalScore >= 31 {
      fmt.Println("もう少しです！")
    } else if totalScore <= 80 && totalScore >= 61 {
      fmt.Println("なかなかいいですね！")
    } else if totalScore >= 81 {
      fmt.Println("素晴らしい！！")
    }
}

// 渡されるtotalScoreのポインタを受け取るように変更してください
func ask(number int, question string, scorePtr*int) {
    var ans string
    fmt.Printf("[質問%d] 次の単語を入力してください: %s\n", number, question)

    sc := bufio.NewScanner(os.Stdin)
    if sc.Scan() {
       ans = sc.Text()
    }
    if question == ans {
        fmt.Println("正解です！")
        // ポインタを使って加算してください
        *scorePtr += 10        
    } else {
        fmt.Println("不正解です!")
    }

}