package main

import "fmt"

func main() {
  //戻り値を格納する変数totalScoreを定義。"+="はさらに足し合わせるという意
  totalScore := ask(1, "html")
  totalScore += ask(2, "css")
  totalScore += ask(3, "ruby")

  //結果を表示
  fmt.Println("スコア", totalScore)
}

//戻り値のデータ型を指定
func ask(number int, question string) int {
  var input string

  fmt.Printf("[質問%d] 次の単語を入力してください: %s\n", number, question)

  fmt.Scan(&input)

  if question == input {
    fmt.Println("正解!")
  　 //正解したら10を戻り値として返す
    return 10

  } else {
    fmt.Println("不正解!")
    //不正解なら0を戻り値として返す
    return 0

  }

}