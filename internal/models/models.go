package models

import ()

type Recipe struct {
	Id           string        `json:"id"`
	Name         string        `json:"name"`
	Description  string        `json:"description,omitempty"`
	PrepTime     string        `json:"prep_time,omitempty"`
	Notes        []string      `json:"notes,omitempty"`
	Tags         []string      `json:"tags,omitempty"`
	Ingredients  []Ingredient  `json:"ingredients,omitempty"`
	Instructions []Instruction `json:"instructions,omitempty"`
}

type Ingredient struct {
	Name     string `json:"name"`
	Quantity string `json:"quantity,omitempty"`
}

type Instruction struct {
	Name  string `json:"name"`
	Timer string `json:"timer_minutes,omitempty"`
	Tip   string `json:"tip,omitempty"`
}
