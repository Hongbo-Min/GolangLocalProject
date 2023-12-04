package clause

import "strings"

type generator func(values ...interface{}) (string, []interface{})

var generators map[Type]generator

func genBindVars(num int) string {
	var vars []string
	for i := 0; i < num; i++ {
		vars = append(vars, "?")
	}
	return strings.Join(vars, ",")
}
