package json

type Status struct {
	Attempts int     `json:"attempts"`
	Good     int     `json:"good"`
	Rate     string  `json:"rate"`
	Percent  float32 `json:"percent"`
}

type Pair struct {
	Foreign   string
	Translate string
	Status    Status
}

func (p *Pair) Rate() {
	if p.Status.Attempts < 10 {
		return
	}
	p.Status.Percent = float32(p.Status.Good) / float32(p.Status.Attempts)

	switch {
	case p.Status.Percent >= 0.8:
		p.Status.Rate = "Well-known"
	case p.Status.Percent < 0.8 && p.Status.Percent >= 0.5:
		p.Status.Rate = "Known"
	case p.Status.Percent < 0.5 && p.Status.Percent >= 0.2:
		p.Status.Rate = "Familiar"
	case p.Status.Percent < 0.2:
		p.Status.Rate = "New"
	}
}
