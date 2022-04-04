package parser

type QuizData struct {
	Purpose      string `json:"purpose"`
	BirthdayYear int    `json:"birthdayYear"`
}

type Wardrobe struct {
	Pants     []Pants     `json:"pants"`
	Shirts    []Shirt     `json:"shirts"`
	Outerwear []Outerwear `json:"outerwear"`
}

type ClothingItem struct {
	ID       int64  `json:"-"`
	PageUrl  string `json:"page"`
	ImageUrl string `json:"image"`
	Color    string `json:"-"`
	Price    string `json:"-"`
}

type Outerwear struct {
	ClothingItem
	Warmth int32 `json:"-"`
}

type Pants struct {
	ClothingItem
}

type Shirt struct {
	ClothingItem
}
