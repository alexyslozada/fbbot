package model

type Buttons struct {
	Type    string `json:"type"`
	Url     string `json:"url,omitempty"`
	Title   string `json:"title"`
	Payload string `json:"payload"`
}

type Elements struct {
	Title    string    `json:"title"`
	Subtitle string    `json:"subtitle"`
	ItemUrl  string    `json:"item_url"`
	ImageUrl string    `json:"image_url"`
	Buttons  []Buttons `json:"buttons"`
}

type Payload struct {
	Url          string     `json:"url,omitempty"`
	TemplateType string     `json:"template_type"`
	Elements     []Elements `json:"elements"`
}

type Attachment struct {
	Type    string `json:"type"`
	Payload `json:"payload"`
}

type MessageContent struct {
	Text       string      `json:"text,omitempty"`
	Attachment *Attachment `json:"attachment,omitempty"`
}

type ResponseMessage struct {
	Recipient      `json:"recipient"`
	MessageContent `json:"message"`
}
