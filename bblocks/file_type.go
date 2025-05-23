package bblocks

import (
	"net/http"
)

/*
La fonction DetermineOutputFileName détermine le nom du fichier de sortie
en se basant sur certaines conditions.
*/
func DetermineOutputFileName(resp *http.Response, url string) string {
	if *Output_name_arg_flag != "" {
		return *Output_name_arg_flag
	}
	// if strings.Contains(resp.Header.Get("Content-Type"), "text/html") {
	// 	return "index.html"
	// }
	return Get_filename(url)
}
