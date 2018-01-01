package request

type ResendRegisterRequest struct {
  Account              string         `json:"account"`
  TypeAccount          string         `json:"typeAccount"`
}
