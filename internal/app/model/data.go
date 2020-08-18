package model


type DataCovid struct {
	ID            int      `json:"id"`
	Title         string   `json:"title"`
	Address       string   `json:"address"`
	Description   string   `json:"description"`
	Lat           float64  `json:"lat"`
	Lng           float64  `json:"lng"`
	Logo          string   `json:"logo"`
	LogoSize      int      `json:"logo_size"`
	Images        []string `json:"images"`
	ActionTitle   string   `json:"action_title"`
	ActionType    string   `json:"action_type"`
	ActionPayload string   `json:"action_payload"`
	MarkerStyle   string   `json:"marker_style"`
	MarkerAnchorX int      `json:"marker_anchor_x"`
	MarkerAnchorY int      `json:"marker_anchor_y"`
	StatusID      int      `json:"status_id"`
}