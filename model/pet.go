package model

type Pet struct {
	Name           string   `json:"name"`
	About          string   `json:"about"`
	Age            string   `json:"age"`
	Size           string   `json:"size"`
	EnergyLevel    string   `json:"energy_level"`
	DependeceLevel string   `json:"dependece_level"`
	Ambience       string   `json:"ambience"`
	Photos         []string `json:"photos"`
	Requirements   string   `json:"requirements"`
}
