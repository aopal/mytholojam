package gameplay

type Move struct { // the swap move is unique
	Name           string `json:"name"`
	Power          int    `json:"power"`
	Type           string `json:"type"`
	Priority       int    `json:"priority"`
	MultiTarget    bool   `json:"multiTarget"`
	TeamTargetable string `json:"teamTargetable"`
}
