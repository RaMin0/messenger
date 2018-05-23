package messenger

// Profile type
type Profile struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Gender     string
	PictureURL string `json:"profile_pic"`
	Locale     string
	Timezone   float64
}

// Response type
type response struct {
	Object string
	Entry  []*entry
}

type entry struct {
	ID        string
	RawTime   float64 `json:"time"`
	Messaging []*Messaging
}

// Messaging type
type Messaging struct {
	Sender       *User
	Recipient    *User
	RawTimestamp float64 `json:"timestamp"`
	Message      *Message
	Postback     *Postback
}

// User type
type User struct {
	ID string
}
