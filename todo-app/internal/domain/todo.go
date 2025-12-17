package domain

type Todo struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Completed  bool   `json:"completed"`
	AICategory string `json:"aiCategory"`
	Category   string `json:"category"`
	Refined    bool   `json:"refined"`
}
