package main

type RawExtractBook struct {
	ASIN       string                  `json:"asin"`
	Title      string                    `json:"title"`
	Authors    string                    `json:"authors"`
	Highlights []RawExtractBookHighlight `json:"highlights"`
}

type RawExtractBookHighlight struct {
	Text     string `json:"text"`
	IsNote   bool   `json:"is_note"`
	Location struct {
		Value int    `json:"value"`
		URL   string `json:"url"`
	} `json:"location"`

	IsNotOnly bool   `json:"is_not_only"`
	Note      string `json:"note"`
}
type Book struct {
	ISBN      string `json:"isbn"`
	Title     string `json:"title"`
	Authors   string `json:"authors"`
	CreatedAt string `json:"created_at"`
}

type Highlight struct {
	ID       int    `json:"id"`
	Text     string `json:"text"`
	Location string `json:"location"`
	Note     string `json:"note"`
	UserID   int    `json:"user_id"`
	BookID   string  `json:"book_id"`
	CreatedAt string `json:"created_at"`
}

type User struct {
	ID       int    `json:"id"`
	Email string `json:"email"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	CreatedAt string `json:"created_at"`
}
