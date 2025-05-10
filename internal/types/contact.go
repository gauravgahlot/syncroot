package types

// Contact represents a contact object in syncroot.
type Contact struct {
	ID        string `json:"id"`
	FullName  string `json:"fullName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func (c *Contact) GetID() string {
	return c.ID
}

func (c *Contact) GetType() string {
	return "Contact"
}
