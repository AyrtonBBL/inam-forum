package models

// like
type Reaction struct {
	ID        string `json:"id"`
	Type      string `json:"type"` // like ou dislike
	UserID    string `json:"user_id"`
	MessageID string `json:"message_id"`
}

// structure reçue en JSON
type ReactionRequest struct {
	Type      string `json:"type"` //like ou dislike
	MessageID string `json:"message_id"`
}
