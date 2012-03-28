package main

import (
	"log"
	"net/http"
	"path/filepath"
	"thegoods.biz/tmplmgr"
)

var (
	mode          = tmplmgr.Development
	assets_dir    = filepath.Join(env("APPROOT", ""), "assets")
	template_dir  = filepath.Join(env("APPROOT", ""), "templates")
	base_template = tmplmgr.Parse(tmpl_root("base.tmpl"))

	context = d{
		"css": list{
			"bootstrap.min.css",
			"bootstrap-responsive.min.css",
			"main.css",
		},
		"js": list{
			"jquery.min.js",
			"jquery-ui.min.js",
			"bootstrap.js",
		},
	}
)

func init() {
	//set our compiler mode
	tmplmgr.CompileMode(mode)

	//add blocks to base template
	base_template.Blocks(tmpl_root("*.block"))
}

func main() {
	http.HandleFunc("/", handle_index)
	http.HandleFunc("/status/", handle_status)
	serve_static("/assets", asset_root(""))
	if err := http.ListenAndServe(":"+env("PORT", "9080"), nil); err != nil {
		log.Fatal(err)
	}
}
