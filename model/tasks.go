package model

type FileTasks struct {
	Tasks []*TaskItem `json:"tasks"`
}

type TaskItem struct {
	Name    string `json:"name"`
	Url     string `json:"url"`
	Headers struct {
		UserAgent string `json:"user-agent"`
	} `json:"headers"`
}
