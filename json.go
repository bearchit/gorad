package gorad

import "encoding/json"

func PrettyJSON(v any) string {
	j, _ := json.MarshalIndent(v, "", "  ")
	return string(j)
}
