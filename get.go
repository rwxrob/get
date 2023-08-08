package get

import "regexp"

// For comparison, here's the valid git repo expression:
// /^([A-Za-z0-9]+@|http(|s)\:\/\/)([A-Za-z0-9.]+(:\d+)?)(?::|\/)([\d\/\w.-]+?)(\.git)?$/i

var RegexpSchema = regexp.MustCompile(
	`^(https?|[A-Za-z0-9]+@[A-Za-z0-9.]+|ssh|file|head|tail|env(.(file|head|tail))?|(home|conf|cache)(.(head|tail))?):`,
)

func Schema(url string) string {
	return RegexpSchema.FindString(url)
}

func SchemaAndValue(url string) (string, string) {
	var val string
	schema := Schema(url)
	if len(url) > len(schema) {
		val = url[len(url):]
	}
	return schema, val
}

func String(source string) (string, error) {
	var it string

	return it, nil
}
