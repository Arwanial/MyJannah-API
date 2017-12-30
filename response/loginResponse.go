package response

type LoginResponse struct {
	Status			 string					`json:"status"`
	Complete		 string					`json:"complete"`
	Token        string         `json:"token"`
}
