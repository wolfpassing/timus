package tinymustache

import (
	"regexp"
)

func TinyMustache(formular string, mappi map[string]string) string {
	ss := formular
	for i, v := range mappi {
		re := regexp.MustCompile(i)
		ss = re.ReplaceAllString(ss, v)
	}
	return ss
}
