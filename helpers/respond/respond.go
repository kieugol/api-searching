package respond

type Respond struct {
	Code     int 		 `json:"code,omitempty"`
	Message  string 	 `json:"message"`
	Data	 interface{} `json:"data"`
}
