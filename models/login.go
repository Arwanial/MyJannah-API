package models

type Login struct{
  Username            string           `json:"username"`
  Password            string           `json:"password"`
  TypeLogin           string           `json:"type"`
}
