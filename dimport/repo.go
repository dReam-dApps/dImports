package dimport

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type repository struct {
	Name   string `json:"name"`
	Commit struct {
		Sha string `json:"sha"`
	} `json:"commit"`
}

// Get the commit hash of a repos main branch from github
func GetCommitHash(owner, repo string) (sha string) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/branches/main", owner, repo)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logger.Errorln("[GetCommitHash]", err)
		return
	}

	req.Header.Set("Accept", "application/vnd.github.v3+json")
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		logger.Errorln("[GetCommitHash]", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logger.Errorf("[GetCommitHash] failed to get commit hash. Status: %s\n", resp.Status)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var r repository
	if err := json.Unmarshal(body, &r); err != nil {
		logger.Errorln("[GetCommitHash]", err)
		return
	}

	return r.Commit.Sha
}
