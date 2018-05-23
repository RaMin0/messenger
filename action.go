package messenger

import (
	rjson "github.com/ramin0/json"
)

func (ms *Messenger) send(userID string, element interface{}) error {
	req := map[string]interface{}{
		"recipient": map[string]interface{}{
			"id": userID,
		},
	}

	if a, ok := element.(*ElmSenderAction); ok {
		req["sender_action"] = a.Action
	} else {
		req["message"] = element
		req["messaging_type"] = "RESPONSE"
	}

	if _, err := ms.httpClient.Post("me/messages", nil, rjson.NewJSON(req), nil); err != nil {
		return err
	}

	return nil
}

// ElmSenderAction type
type ElmSenderAction struct {
	Action string
}

// ElmText type
type ElmText struct {
	Text string `json:"text"`
	ElmTemplate
	ElmQuickReplies
}

// ElmButton type
type ElmButton struct {
	Type                string `json:"type"`
	Title               string `json:"title"`
	Payload             string `json:"payload,omitempty"`
	URL                 string `json:"url,omitempty"`
	WebviewHeightRatio  string `json:"webview_height_ratio,omitempty"`
	WebviewShareButton  string `json:"webview_share_button,omitempty"`
	MessengerExtensions string `json:"messenger_extensions,omitempty"`
}

// ElmButtons type
type ElmButtons struct {
	Buttons []*ElmButton `json:"buttons,omitempty"`
}

// ElmQuickReply type
type ElmQuickReply struct {
	Type     string `json:"content_type"`
	Title    string `json:"title,omitempty"`
	Payload  string `json:"payload,omitempty"`
	ImageURL string `json:"image_url,omitempty"`
}

// ElmQuickReplies type
type ElmQuickReplies struct {
	QuickReplies []*ElmQuickReply `json:"quick_replies,omitempty"`
}

// ElmTemplate type
type ElmTemplate struct {
	Attachment *ElmAttachment `json:"attachment,omitempty"`
}

// ElmAttachment type
type ElmAttachment struct {
	Type    string      `json:"type"`
	Payload *ElmPayload `json:"payload"`
}

// ElmPayload type
type ElmPayload struct {
	Type     string `json:"template_type"`
	Sharable bool   `json:"sharable"`
	Text     string `json:"text,omitempty"`
	ElmButtons
	ElmElements
}

// ElmElement type
type ElmElement struct {
	Title    string `json:"title,omitempty"`
	Subtitle string `json:"subtitle,omitempty"`
	ImageURL string `json:"image_url,omitempty"`
	ElmButtons
}

// ElmElements type
type ElmElements struct {
	Elements []*ElmElement `json:"elements,omitempty"`
}

// SendSenderAction func
func (ms *Messenger) SendSenderAction(userID, action string) error {
	elm := &ElmSenderAction{
		Action: action,
	}
	return ms.send(userID, elm)
}

// SendText func
func (ms *Messenger) SendText(userID, text string) error {
	elm := &ElmText{
		Text: text,
	}
	return ms.send(userID, elm)
}

// SendButtons func
func (ms *Messenger) SendButtons(userID, text string, buttons []*ElmButton) error {
	elm := &ElmTemplate{
		Attachment: &ElmAttachment{
			Type: "template",
			Payload: &ElmPayload{
				Type: "button",
				Text: text,
			},
		},
	}
	elm.Attachment.Payload.Buttons = buttons
	return ms.send(userID, elm)
}

// SendQuickReplies func
func (ms *Messenger) SendQuickReplies(userID, text string, quickReplies []*ElmQuickReply) error {
	elm := &ElmText{
		Text: text,
	}
	elm.QuickReplies = quickReplies
	return ms.send(userID, elm)
}

// SendGenericTemplate func
func (ms *Messenger) SendGenericTemplate(userID string, elements []*ElmElement) error {
	elm := &ElmTemplate{
		Attachment: &ElmAttachment{
			Type: "template",
			Payload: &ElmPayload{
				Type: "generic",
			},
		},
	}
	elm.Attachment.Payload.Elements = elements
	return ms.send(userID, elm)
}
