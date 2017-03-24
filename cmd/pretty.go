package cmd

import (
	"fmt"
	"reflect"
)

type PrettyMeta struct {
	Name   string
	Length int
}

func SliceToMap(v interface{}) (mapValue map[string]interface{}, ok bool) {
	values := reflect.Indirect(reflect.ValueOf(v))
	if values.Kind() != reflect.Slice {
		return
	}

	if values.Len() <= 0 {
		fmt.Print("Empty set")
		fmt.Print("\n\n")
		return
	}
	// cache
	dicts := make([][]string, values.Len())
	metas := []PrettyMeta{}

	// string size counter
	for i := 0; i < values.Len(); i++ {
		value := values.Index(i)
		t := value.Type()
		dicts[i] = make([]string, value.NumField())

		// check length
		for j := 0; j < value.NumField(); j++ {
			fieldName := t.Field(j).Name

			// not found key
			if i == 0 {
				metas = append(metas, PrettyMeta{
					Name:   t.Field(j).Name,
					Length: len(t.Field(j).Name),
				})
			}

			vv := fmt.Sprint(value.FieldByName(fieldName).Interface())
			le := len(vv)

			if metas[j].Length < le {
				metas[j].Length = le
			}

			dicts[i][j] = vv

		}

	}

	// get total line length
	total := 0
	for _, d := range metas {
		// add span
		total = total + d.Length + 3
	}
	total = total - 1

	// print header
	fmt.Print("+")
	for i := 0; i < total; i++ {
		fmt.Print("-")
	}
	fmt.Print("+\n")

	fmt.Print("|")
	for _, dd := range metas {
		fmt.Printf(" %"+fmt.Sprint(-dd.Length)+"s ", dd.Name)
		fmt.Print("|")
	}
	fmt.Print("\n")

	fmt.Print("+")
	for i := 0; i < total; i++ {
		fmt.Print("-")
	}
	fmt.Print("+\n")

	for _, ddd := range dicts {
		fmt.Print("|")
		for kkkk, dddd := range ddd {
			fmt.Printf(" %"+fmt.Sprint(metas[kkkk].Length)+"s |", dddd)
		}
		fmt.Print("\n")
	}

	fmt.Print("+")
	for i := 0; i < total; i++ {
		fmt.Print("-")
	}
	fmt.Print("+\n")
	fmt.Printf("%d rows in set", len(dicts))
	fmt.Print("\n\n")

	return
}
