package structs

type IPSet struct {
	UUID    string `json:"uuid"`
	CIDR    string `json:"cidr"`
	Comment string `json:"comment"`
}

type IP struct {
	UUID      string `json:"uuid"`
	IPAddress string `json:"ip_address"`
	Type      string `json:"type"`
	Role      string `json:"role"`
	TargetId  string `json:"target_id"`
	IPSetId   string `json:"ipset_id"`
}
