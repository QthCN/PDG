package structs

type AuditRecord struct {
	Id         int64  `json:"id"`
	Username   string `json:"username"`
	Action     string `json:"action"`
	URL        string `json:"url"`
	Args       string `json:"args"`
	ActionTime string `json:"action_time"`
}

type AuditRecords struct {
	TotalCnt int64          `json:"total_cnt"`
	Records  []*AuditRecord `json:"records"`
}

type ListAuditRecordsCondition struct {
	CurrentPage    int64 `json:"current_page"`
	RecordsPerPage int64 `json:"records_per_page"`
}
