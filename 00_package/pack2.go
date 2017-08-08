package pack1

import "strings"

func Revwolf() string {

	local := make([]string,0)

	for i:=len(Wolf)-1; i>=0; i-- {
		local = append(local,string(Wolf[i]))
	}
	return strings.Join(local,"")
}
