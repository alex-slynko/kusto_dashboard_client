package kusto_dashboard_client

import "time"

type Dashboard struct {
	Id        string `json:"id"`
	Title     string `json:"title"`
	CreatedAt int64  `json:"createdAt"`
	CreatedBy string `json:"createdBy"`
	OpenedAt  int64  `json:"openedAt"`
}

func (d *Dashboard) GetCreatedAt() time.Time {
	return time.Unix(d.CreatedAt, 0)
}
