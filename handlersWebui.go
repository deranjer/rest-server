package restserver

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/asdine/storm/v3"
	"goji.io/pat"
)

var homeTemplate *template.Template

// WebServer handles the webui
type WebServer struct {
	Listen   string
	Log      string
	Database *storm.DB
	Tset     *template.Template // The webui Templates
	RootPath string             //The repo root path
}

// HomePage contains the variables needed to render the homepage
type HomePage struct {
	Title     string //Page title
	Repos     []Repository
	RepoCount int
}

type RepoPage struct {
	Files     []RepoBlob
	Repos     []Repository
	RepoCount int
}

// ServeHome serves up the homepage
func (web *WebServer) ServeHome(w http.ResponseWriter, r *http.Request) {
	repoList, err := FetchRepos(web.Database)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Repos: ", repoList)
	data := HomePage{
		Title:     "Restic Repo",
		Repos:     repoList,
		RepoCount: len(repoList),
	}
	err = web.Tset.ExecuteTemplate(w, "home.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// ServeRepo displays all the repo information
func (web *WebServer) ServeRepo(w http.ResponseWriter, r *http.Request) {
	repoIDString := pat.Param(r, "repoID")
	repoID, err := strconv.Atoi(repoIDString)
	if err != nil {
		fmt.Println("Unable to convert repoID string to int")
		return
	}
	repo, err := FetchRepoFromID(repoID, web.Database)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fileList, err := FetchBlobsFromRepo(repo.Path, web.Database)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	repoList, err := FetchRepos(web.Database)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := RepoPage{
		Files:     fileList,
		Repos:     repoList,
		RepoCount: len(repoList),
	}
	if err := web.Tset.ExecuteTemplate(w, "repos.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
