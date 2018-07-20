package params

// AuthRep is the struct that to response to the client
type AuthRep struct {
	RetCode string `json:"ret_code,omitempty"`
	RetMsg  string `json:"ret_msg,omitempty"`
	CodeUrl string `json:"code_url,omitempty"`
}
