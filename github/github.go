package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {

	// curl https://api.github.com/users/ardanlabs > reply.json
	// GoLand: github/reply.json
	file, err := os.Open("reply.json")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer file.Close()
	// data, err := io.ReadAll(file)

	login := "ardanlabs"
	name, numRepos, err := userInfo(login)
	if err != nil {
		log.Fatalf("error: can't get %q info - %s", login, err)
	}
	fmt.Printf("name: %s, num repos: %d\n", name, numRepos)
}

// Solving: Get, Parse, Analyze, Output

// userInfo return name and number of public repos from GitHub API.
func userInfo(login string) (string, int, error) {
	url := "https://api.github.com/users/" + url.PathEscape(login) // also fmt.Sprintf
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("%q: bad status - %s", url, resp.Status)
	}
	/*
		fmt.Println("content-type:", resp.Header.Get("Content-Type"))
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			log.Fatalf("error: can't print body - %s", err)
		}
	*/
	return decodeUserInfo(resp.Body)
}

func decodeUserInfo(r io.Reader) (string, int, error) {
	var userInfo struct {
		Name     string
		NumRepos int `json:"public_repos"` // field tag
	}
	dec := json.NewDecoder(r)
	if err := dec.Decode(&userInfo); err != nil {
		return "", 0, fmt.Errorf("can't decode JSON - %s", err)
	}
	return userInfo.Name, userInfo.NumRepos, nil
}

/* JSON
JSON <-> Go
string <-> string
number <-> float64, int8 ... int64, int, uint8 ... uint64, uint, float32 ..
true/false <-> true/false
null <-> nil
array <-> []T, []any
object <-> struct, map[string]any

encoding/json API
JSON -> io.Reader -> Go: json.NewDecoder
Go -> io.Writer -> JSON: json.NewEncoder
JSON -> []byte -> Go: json.Unmarshal
Go -> []byte -> JSON: json.Marshal
*/
