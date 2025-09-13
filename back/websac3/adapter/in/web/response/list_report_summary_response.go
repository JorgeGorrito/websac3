package response

import "time"

type ListReportSummaryResponse struct {
	ID        uint      `json:"id"`
	Lang      string    `json:"lang"`
	Score     float32   `json:"score"`
	CreatedAt time.Time `json:"created_at"`
}

