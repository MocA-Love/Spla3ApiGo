package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)


func main() {
	matchType := "regular"
	limit     := 3

	matchTypes := map[string]string{
		"regular":           "レギュラーマッチ",
		"bankara-open":      "バンカラマッチ(オープン)",
		"bankara-challenge": "バンカラマッチ(チャレンジ)",
		"fest":              "フェスマッチ(オープン)",
		"fest-challenge":    "フェスマッチ(チャレンジ)",
		"x":                 "Xマッチ",
		"event":             "イベントマッチ",
	}

	endpoint  := fmt.Sprintf("%s/%s/schedule", APIBaseURL, matchType)
	data, err := fetchSchedule(endpoint)
	if err != nil {
		fmt.Println("Error fetching schedule:", err)
		return
	}

	message := createMessage(data, matchType, matchTypes, limit)
	printMessage(message)
}


func fetchSchedule(endpoint string) (NormalStageInfo, error) {
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return NormalStageInfo{}, fmt.Errorf("creating request: %w", err)
	}

	req.Header.Set("User-Agent", UserAgent)

	client    := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return NormalStageInfo{}, fmt.Errorf("sending request: %w", err)
	}

	defer resp.Body.Close()

	var data NormalStageInfo
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return NormalStageInfo{}, fmt.Errorf("decoding response: %w", err)
	}

	return data, nil
}


func createMessage(data NormalStageInfo, matchType string, matchTypes map[string]string, limit int) []string {
	var messages []string
	messages = append(messages, fmt.Sprintf(">> %s\n\n", matchTypes[matchType]))

	count := 0
	for _, result := range data.Results {
		if !strings.Contains(matchType, "fest") && result.IsFest {
			messages = append(messages, "今はフェス開催中なのでみれないよ\n")
			break
		} else if strings.Contains(matchType, "fest") && !result.IsFest {
			messages = append(messages, "今はフェス開催してないよ\n")
			break
		}

		count++
		if result.Event != nil && count == 1 {
			messages = append(messages, fmt.Sprintf("[%s]\n", result.Event.Name))
			messages = append(messages, fmt.Sprintf("%s\n\n", result.Event.Desc))
		}

		messages = append(messages, fmt.Sprintf("・%s～%s\n", result.StartTime.Format("01/02 15:04"), result.EndTime.Format("15:04")))

		if result.Rule != nil {
			messages = append(messages, fmt.Sprintf(" [%s]\n", result.Rule.Name))
		}

		if result.Stages != nil {
			for _, stage := range result.Stages {
				messages = append(messages, fmt.Sprintf("　- %s\n", stage.Name))
			}
		}

		if result.TricolorStage != nil {
			messages = append(messages, "\n>> トリカラバトル\n")
			messages = append(messages, fmt.Sprintf("　・%s,\n", result.TricolorStage.Name))
			messages = append(messages, "\n")
		}

		if count == limit {
			break
		}

		messages = append(messages, "---------------------\n")
	}

	return messages
}


func printMessage(message []string) {
	for _, m := range message {
		fmt.Print(m)
	}
}
