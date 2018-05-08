package thehive

const (
	taskMain   = caseMain + "/task"
	taskSearch = taskMain + "/_search" // POST
	taskStats  = taskMain + "/_stats"  // POST
	taskRoute  = taskMain + "/%s"      // GET, PATCH
	taskCreate = caseRoute + "/task"   // POST
)

// Task represents TheHive case's task
type Task struct {
	Entity
	CreatedAt int64  `json:"createdAt"`
	Status    string `json:"status"`
	CreatedBy string `json:"createdBy"`
	Title     string `json:"title"`
	Order     int    `json:"order"`
	User      string `json:"user"`
	Flag      bool   `json:"flag"`
	ID        string `json:"id"`
}

// SendableTask is task that is able to be sent
type SendableTask struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Owner       string `json:"owner,omitempty"`
	Status      string `json:"status,omitempty"`
	Flag        bool   `json:"flag,omitempty"`
	StartDate   int64  `json:"startDate,omitempty"`
	EndDate     int64  `json:"endDate,omitempty"`
}
