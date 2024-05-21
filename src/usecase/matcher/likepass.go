package matcher

type LikePass struct {
	UserEmail   string
	EmailTarget string `json:"target"`
	Action      string `json:"action"`
}
