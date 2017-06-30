package model

type Sender struct {
	ID string `json:"id"`
}

type Recipient struct {
	ID string `json:"id"`
}

type Message struct {
	Mid  string `json:"mid"`
	Seq  int64  `json:"seq"`
	Text string `json:"text"`
}

type Messaging struct {
	Sender    `json:"sender"`
	Recipient `json:"recipient"`
	Message   `json:"message"`
	Timestamp int64 `json:"timestamp"`
}

type Entry struct {
	ID        string      `json:"id"`
	Time      int64       `json:"time"`
	Messaging []Messaging `json:"messaging"`
}

type CommonFormat struct {
	Object string  `json:"object"`
	Entry  []Entry `json:"entry"`
}
