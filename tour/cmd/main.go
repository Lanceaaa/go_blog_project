package main

import (
	"github.com/spf13/cobra"
	"log"
)

func main() {
	var wordCmd = &cobra.Command{
		Use: "word",
		Short: "单词格式转换",
		Log: "支持多种单词格式转换",
		Run: func(cmd *cobra.Commad, ags []string) {},
	}

	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}