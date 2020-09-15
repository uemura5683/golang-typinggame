package main

import (
    "bufio"
    "flag"
    "fmt"
    "io"
    "os"
    "time"
    "math/rand"
)

// 配列をシャッフルする
func shuffle(data []string) {
    n := len(data)
    rand.Seed(time.Now().Unix())
    for i := n - 1; i >= 0; i-- {
        j := rand.Intn(i + 1)
        data[i], data[j] = data[j], data[i]
    }
}

var t int

func init() {
    //オプションで制限時間をできる
    flag.IntVar(&t, "t", 1, "制限時間(分)")
    flag.Parse()
}

func main() {
    var (
        ch_rcv    = myinput(os.Stdin)
        tm        = time.After(time.Duration(t) * time.Minute)
        words     = []string{ "raccoon", "dog", "wild boar", "rabbit", "cow", "horse", "wolf", "hippopotamus", "kangaroo", "fox", "giraffe", "bear", "koala", "bat", "gorilla", "rhinoceros", "monkey", "deer", "zebra", "jaguar", "polar bear", "skunk", "elephant", "raccoon dog", "animal", "reindeer", "rat", "tiger", "cat", "mouse", "buffalo", "hamster", "panda", "sheep", "leopard", "pig", "mole", "goat", "lion", "camel", "squirrel", "donkey", "Crocodile", "lion is very cool!", "this is a dog", "I love donkey", "do you like animal?", "Amphibians", "Butterfly", "excellentswimmer is long language" }
        score     = 0
    )
    fmt.Println()

    shuffle(words);
    fmt.Println("タイピングゲームを始めます。制限時間は", t, "分。1語1点、", len(words), "点満点")
    //送信用チャネル
    num := 1
    for i := true; i && score < len(words); {
        question := words[score]
        fmt.Print("[質問", num ,"]次の単語を入力してください:", question, "\n")
        fmt.Print("[答え]")
        select {
        case x := <-ch_rcv:
            //標準入力に何か入力された時の処理
            // 入力された文字が一致しているかどうかをチェックする
            if question == x {
                fmt.Println("正解です！")
                score++
                num++
            } else {
                fmt.Println("不正解です！")
            }
        case <-tm:
            //制限時間が過ぎた際の処理
            fmt.Println("\n制限時間を過ぎました")
            i = false
        }
        
    }
    fmt.Println("あなたの点数:", score, "点 / ", len(words), " 点")
    if score <= 10 {
      fmt.Println("判定 F")
    } else if score <= 15 && score > 10 {
      fmt.Println("判定 E")
    } else if score <= 20 && score > 15 {
      fmt.Println("判定 D")
    } else if score <= 25 && score > 20 {
      fmt.Println("判定 C")
    } else if score <= 30 && score > 25 {
      fmt.Println("判定 B")
    } else if score <= 35 && score > 30 {
      fmt.Println("判定 A")
    } else if score <= 40 && score > 35 {
      fmt.Println("判定 S")
    } else if score <= 45 && score > 40 {
      fmt.Println("判定 SS")
    } else if score > 45 {
      fmt.Println("判定 SSS")
    }
}

func myinput(r io.Reader) <-chan string {
    // サブgo ルーチン
    // 標準入力から受け取った文字列を標準出力へ出力する
    ch1 := make(chan string)
    go func() {
        s := bufio.NewScanner(r)
        for s.Scan() {
            ch1 <- s.Text()
        }
    }()
    return ch1
}