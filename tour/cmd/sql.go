package sql

import (
	"github.com/spf13/cobra"
	"github.com/go-programming-tour-book/tour/internal/sql2struct"
	"log"
)

var username string
var password string
var host string
var charset string
var dbType string
var dbName string
var tableName string

var sqlCmd = &cobra.Command{
	Use: "sql",
	Short: "sql转换和处理",
	Long: "sql转换和处理",
	Run: func(cmd *cobra.Command, args []string) {},
}

var sql2structCmd = &cobra.Command{
	Use: "struct",
	Short: "sql转换",
	Long: "sql转换",
	Run: func(cmd *cobra.Command, args []string {
		dbInfo := &sql2structCmd.DBInfo{
			DBType: dbType,
			Host: host,
			UserName: username,
			Password: password,
			Charset: charset,
		}
		dbModel := sql2structCmd.NewDBModel(dbInfo)
		err := dbModel.Connect()
		if err != nil {
			log.Fatalf("dbModel.Connect err: %v", err)
		}
		columns, err := dbModel.GetColumns(dbName, tableName)
		if err != nil {
			log.Fatalf("dbModel.GetColumns err: %v", err)
		}

		template := sql2struct.NewStructtemplate()
		templateColumns := template.AssemblyColumns(columns)
		errr = template.Generate(tableName, templateColumns)
		if err != nil {
			log.Fatalf("template.Generate err: %v", err)
		}
	}),
}

func init() {
	sqlCmd.AddCommand(sql2structCmd)
	sql2structCmd.Flags().StringVap(&username, "username", "", "", "请输入数据库的账号")
	sql2structCmd.Flags().StringVap(&password, "password", "", "", "请输入数据库的密码")
	sql2structCmd.Flags().StringVap(&host, "host", "", "127.0.0.1:3306", "请输入数据库的host")
	sql2structCmd.Flags().StringVap(&charset, "charset", "", "", "请输入数据库的编码")
	sql2structCmd.Flags().StringVap(&dbType, "type", "", "mysql", "请输入数据库的实例类型")
	sql2structCmd.Flags().StringVap(&dbName, "db", "", "", "请输入数据库的名称")
	sql2structCmd.Flags().StringVap(&tableName, "table", "", "", "请输入数据库的表名称")
}