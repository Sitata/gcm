package ccs

// OutMsg is a message to be sent to GCM CCS.
// If ID field is not set, it will be generated automatically using crypto/rand.
// Google recommends Data field to be strings key/value pairs and keys cannot be
// reserved words described in GCM server documentation.
// https://firebase.google.com/docs/cloud-messaging/xmpp-server-ref
type OutMsg struct {
	To                       string            `json:"to"`
	Condition                string            `json:"condition,omitempty"`
	ID                       string            `json:"message_id"`
	CollapseKey              string            `json:"collapse_key,omitempty"`
	MessageType              string            `json:"message_type"`
	Priority                 string            `json:"priority,omitempty"`
	ContentAvail             bool              `json:"content_available,omitempty"`
	DelayWhileIdle           bool              `json:"delay_while_idle,omitempty"`
	TimeToLive               int               `json:"time_to_live,omitempty"`               //default:2419200 (in seconds = 4 weeks)
	DeliveryReceiptRequested bool              `json:"delivery_receipt_requested,omitempty"` //default:false
	DryRun                   bool              `json:"dry_run,omitempty"`                    //default:false
	Data                     map[string]string `json:"data,omitempty"`
	Notification             map[string]string `json:"notification,omitempty"`
}

// InMsg is an incoming GCM CCS message.
type InMsg struct {
	From           string            `json:"from"`
	ID             string            `json:"message_id"`
	Category       string            `json:"category"`
	Priority       string            `json:"priority"`
	Data           map[string]string `json:"data"`
	Notification   map[string]string `json:"notification"`
	MessageType    string            `json:"message_type"`
	RegistrationID string            `json:"registration_id"`
	ControlType    string            `json:"control_type"`
	Err            string            `json:"error"`
	ErrDesc        string            `json:"error_description"`
}
