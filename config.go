package main

import "time"


const (
	APIBaseURL = "https://spla3.yuu26.com/api"
	UserAgent  = "Spla3ApiGo/1.0(contact)"
)


type NormalStageInfo struct {
	Results []struct {
		StartTime     time.Time `json:"start_time"`
		EndTime       time.Time `json:"end_time"`
		Rule          *struct {
			Key  string `json:"key"`
			Name string `json:"name"`
		} `json:"rule"`
		Stages        []struct {
			ID    int    `json:"id"`
			Name  string `json:"name"`
			Image string `json:"image"`
		} `json:"stages"`
		Event         *struct {
			ID   string `json:"id"`
			Name string `json:"name"`
			Desc string `json:"desc"`
		} `json:"event"`
		IsFest        bool `json:"is_fest"`
		IsTricolor    bool `json:"is_tricolor"`
		TricolorStage *struct {
			Name  string `json:"name"`
			Image string `json:"image"`
		} `json:"tricolor_stage"`
	} `json:"results"`
}

