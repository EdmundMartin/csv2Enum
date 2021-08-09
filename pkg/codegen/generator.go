package codegen

import (
	"fmt"
	"os"
	"text/template"
)

type EnumEntry struct {
	Id   string
	Repr string
	Done bool
}

type EnumInfo struct {
	Package  string
	EnumName string
	Entries  []EnumEntry
}

func GenerateEnum(info EnumInfo) error {
	f, err := os.OpenFile(fmt.Sprintf("%s.lang", info.EnumName), os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	if err != nil {
		return err
	}
	temp := template.Must(template.New("enum_java.lang.tmpl").ParseFiles("pkg/codegen/enum_java.lang.tmpl"))
	err = temp.Execute(f, info)
	return err
}
