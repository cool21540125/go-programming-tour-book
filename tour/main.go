package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

var (
	operator    string
	showProcess bool
)

const (
	plus       string = "+"
	minus      string = "-"
	multiplied string = "*"
	dividedBy  string = "/"
)

// 設定要從 command line 讀取的參數
// 這邊所設定的會在 -h 或者輸入錯誤時出現提示哦！
func init() {
	// 使用方法 -o *
	flag.StringVar(&operator, "o", "+", "the math `operator`")
	// bool 要使用 -p=true
	flag.BoolVar(&showProcess, "p", false, "`show` process")
	flag.Usage = usage
}

// 印出預設的說明
func usage() {
	fmt.Fprintf(os.Stderr, "Usage: math [options] [root]\n")
	fmt.Fprintf(os.Stderr, "  Currently, there are four URI routes could be used:\n")
	flag.PrintDefaults()
}

func main() {
	// 將 command line 上的 arguments 進行處理
	flag.Parse()

	// 將 argument 轉成 int, flag.Arg 可以拿取非指定的 argument
	number1, _ := strconv.Atoi(flag.Arg(0))
	number2, _ := strconv.Atoi(flag.Arg(1))

	if showProcess {
		fmt.Printf("%d %s %d = ", number1, operator, number2)
	}

	// 處理運算
	switch operator {
	case plus:
		fmt.Printf("%d \n", number1+number2)
	case minus:
		fmt.Printf("%d \n", number1-number2)
	case multiplied:
		fmt.Printf("%d \n", number1*number2)
	case dividedBy:
		fmt.Printf("%0.2f \n", float64(number1)/float64(number2))
	}
}
