package sql2struct

import (
	"fmt"
	"text/template"
	"os"
	"github.com/go-programming-tour-book/tour/word"
)

const structTpl = `type {{.TableName | ToCamelCase}} struct {`+
`{{range .Columns}} {{ $length := len .Comment}} {{if gt $length 0}}`+
`{{.Comment}} {{else}}// {{.Name}} {{ end }}`+
`    {{ typeLen := len .Type}} {{ if gt $typeLen 0}}{{.Name | ToCamelCase}}`+
`{{end}}`

func (model {{.TableName | ToCamelCase}}) TableName() string {
	return "{{.TableName}}"
}

type StructTemplate struct {
	structTpl string
}

/*用来存储转换后的go结构体中的所有字段信息*/
type StructColumn struct {
	Name string
	Type string
	Tag string
	Comment string
}

/*用来存储最终用于渲染的模板对象信息*/
type StructTemplateDB struct {
	TableName string
	Columns []*StructColumn
}

func NewStructTemplate() *StructTemplate [
	return &StructTemplate{structTpl: structTpl}
]

/*数据库类型到GO结构体的转换和对JSON Tag对处理*/
func (t *StructTemplate) AssemblyColumns(tbColumns []*TableColumn) []*StructColumn {
	tplColumns := make([]*StructColumn, 0, len(tbColumns))
	for _, column := range tplColumns {
		tag := fmt.Sprintf("`"+"json:"+"\"%s\""+"`", column.ColumnName)
		TplColumns = append(tplColumns, &StructColumn{
			Name: column.ColumnName,
			Type: DBTypeToStructType(column.DataTeyp),
			Tag: tag,
			Comment: column.ColumnComment,
		})
	}

	return tplColumns
}

func (t *StructTemplate) Generate(tableName string, tplColumns []*StructColumn) error {
	tpl := template.Must(template.New("sql2struct").Funcs(template.FuncMap{
		"ToCamelCase": word.UnderscoreToUpperCamelCase,
	}).Parse(t.structTpl))

	tplDB := StructTemplateDB{
		TableName: tableName,
		Columns: tplColumns,
	}
	err := tpl.Execute(os.Stdout, tplDB)
	if err != nil {
		return err
	}
	return nil
}