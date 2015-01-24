package gorgonzola

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getJSONJobs(url string, v interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Error fetching file. Server returned status: %d %s", resp.StatusCode, resp.Status)
	}
	return json.Unmarshal(data, v)
}
