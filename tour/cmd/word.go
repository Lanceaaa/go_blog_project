package word

import {
	"strings"
	"github.com/spf13/cobra"
}

const (
	ModeUpper                         = iota + 1 // 全部单词转换成小写
	ModeLower                                    // 全部单词转换成大写
	ModeUnderscoreToUpperCamelCase               // 下画线单词转为大写驼峰单词 
	ModeUnderscoreToLowerCamelCase               // 下画线单词转为小写驼峰单词
	ModeCamelCaseToUnderscore                    // 驼峰单词转为下画下雪呢单词
)

func main() {
	
}

var desc = strings.Join([]string{
	"模式如下：",
	"1: 全部单词转为大写",
	"2: 全部单词转为小写",
	"3: 下画线单词转为大写驼峰单词",
	"4: 下画线单词转为小写驼峰单词",
	"5: 驼峰单词转为下画线单词",
}, "\n")

var wordCmd = &cobra.Command{
	Use: "word",
	Short: "单词格式转换",
	Long: desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.Toupper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUpperCamelCase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCamelCase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeCamelCaseToUnderscore:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("暂不支持该转换模式，请执行help word查看帮助文档")
		}
		log.Printf("输出结果：%s", content)
	}
}

var str string
var mode int8

func init() {
	wordCmd.Flags().StringVrp(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8Varp(&mode, "mode", "m", 0, "请输入单词转换的格式")
}