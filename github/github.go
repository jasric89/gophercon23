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
	// JSON File Reply
	file, err := os.Open("reply.json")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer file.Close()
	//data, err := io.ReadAll(file) to read a whole file

	var userInfo struct {
		Name         string
		Public_Repos int
		NumRepos     int `json:"public_repos"` /*This is called a field tag if you dont like the field names in Json
		  This is what the Azure DevOps SDK is using. */
	}

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&userInfo); err != nil {
		log.Fatalf("error: cant decode Json - %s", err)
	}
	fmt.Printf("%v\n", userInfo)
	fmt.Println("Finshed Reading Json File")

	// URL Reply
	url := "https://httpbin.org/get"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("error: bad status - %s", resp.Status)
	}
	fmt.Println("content-type:", resp.Header.Get("Content-Type"))
	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		log.Fatalf("error: cant print body - %s", err)
	}

	login := "ardanlabs"
	name, numRepos, err := userInfo(login)
	if err != nil {
		log.Fatal("Error: Cant get %q info -s %s", login, err)
	}
	fmt.Printf("name: %s, num repos: %d\n", name, numRepos)
}

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

// Solving: Get, Paese, Analyze, Out // LOOK AT LATTEERRRRRRR!!!!! FUCK SAKE!!!! WHY DID YOU NOT GET THIS SHIT!!
/*
func userInfo(login string) (string, int, error) {
	url := "https://api.github.com/orgs/users/" + url.PathEscape(login)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("error: bad status - %s", url, resp.Status)
	}

	//	fmt.Println("content-type:", resp.Header.Get("Content-Type"))
	//	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
	//		log.Fatalf("error: cant print body - %s", err)
	//	}
	var userInfo struct {
		Name         string
		Public_Repos int
		NumRepos     int `json:"public_repos"`
	}

	decoder := json.NewDecoder(io.read)
	if err := decoder.Decode(&userInfo); err != nil {
		log.Fatalf("error: cant decode Json - %s", err)
	}
	fmt.Printf("%v\n", userInfo)
	fmt.Println("Finshed Reading Json File")

	return decoder, err, int
}
*/

/*
encoding /json API
JSON -> io.Reader -> Go: json.NewDecoder
go -> io.Writer -> JSON: json.NewEncoder
JSON -> []byte -> go:json.Marshal
[]byte -> JSON -> go:json.Unmarshal

*/
