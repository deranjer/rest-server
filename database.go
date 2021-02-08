package restserver

import (
	"fmt"
	"log"
	"time"

	"github.com/asdine/storm/v3"
)

// Repository is the struct for all our unique repos
type Repository struct {
	RepoID       int    `storm:"id,increment"` //primary key
	Path         string `storm:"index,unique"` //path to the repo
	DateCreated  time.Time
	LastPushDate time.Time
	Client       string //might need to modify the restic client for this.
}

// RepoBlob is the struct for each indiv file in a repo
type RepoBlob struct {
	FileID int `storm:"id,increment"`
	//Hash string     `storm:"id"` //Id of the file
	Name         string // Name of the file
	Type         string // Type of File
	Path         string `storm:"unique"` // Path of the file
	DateModified time.Time
	RepoPath     string
	//Repo Repository // Repo the file belongs to
}

// OpenOrCreateDatabase creates or opens the database connection
func OpenOrCreateDatabase() (db *storm.DB) {
	db, err := storm.Open("repoDB.db")
	if err != nil {
		fmt.Println("Unable to open or create database... ", err)
		log.Fatal("Unable to open or create database... ", err)
	}
	return db
}

// CreateRepoDB creates a database entry with the repo information
func CreateRepoDB(repo Repository, db *storm.DB) error {
	err := db.Save(&repo)
	if err != nil {
		fmt.Println("Fatal error occured creating repo: ", err)
		return err
	}
	log.Printf("Created repo from client: %s, with ID: %d, at Path: %s", repo.Client, repo.RepoID, repo.Path)
	//fmt.Printf("Created repo from client: %s, with ID: %d, at Path: %s", repo.client, repo.ID, repo.Path)
	return nil
}

// FetchRepos fetches a list of all repos in the db
func FetchRepos(db *storm.DB) (repoList []Repository, err error) {
	var repos []Repository
	err = db.All(&repos)
	if err != nil {
		fmt.Println("Fatal error occured fetching repos: ", err)
		return nil, err
	}
	log.Print("Fetched list of all repos...")
	return repos, nil
}

// FetchRepoFromID will fetch a repo from the ID
func FetchRepoFromID(repoID int, db *storm.DB) (repo Repository, err error) {
	err = db.One("RepoID", repoID, &repo)
	if err != nil {
		fmt.Printf("Unable to find repo from ID: %d err: %s", repoID, err)
		return repo, err
	}
	log.Print("Found Repo: ", repo.Path)
	return repo, nil
}

func CreateRepoBlob(repoBlob RepoBlob, db *storm.DB) error {
	err := db.Save(&repoBlob)
	if err != nil {
		fmt.Println("Fatal error occured creating blob: ", err)
		return err
	}
	return nil
}

// FetchBlobsFromRepo gets all of the files listed in the repo
func FetchBlobsFromRepo(repoPath string, db *storm.DB) (fileList []RepoBlob, err error) {
	//var files []RepoBlob
	// query := db.Select(q.Eq("RepoPath", repoPath))
	// err = query.Find(&files)
	fmt.Println("Looking in repo path: ", repoPath)
	err = db.Find("RepoPath", repoPath, &fileList)
	if err != nil {
		fmt.Println("Fatal error occurred selecting blobs in repo: ", err)
		return nil, err
	}
	return fileList, nil
}
