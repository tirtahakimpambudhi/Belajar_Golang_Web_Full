package controller

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
)

func StaticFiles() http.Handler {
	dir := http.Dir("./resource")
	fileserver := http.FileServer(dir)
	return fileserver
}
var Resource embed.FS
func StaticFilesEmbed() http.Handler {
	workdir , _ := fs.Sub(Resource,"resource")
	fileserver := http.FileServer(http.FS(workdir))
	return fileserver
}

func ServerFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w,r,"./resource/html/about.html")
}
func ServeFileEmbed(w http.ResponseWriter, r *http.Request) {
	content , _ := Resource.ReadFile("resource/html/home.html")
	fmt.Fprint(w, string(content))
}

