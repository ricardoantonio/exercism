package gigasecond

import "time"

// Gigasecond represents one billion seconds.
const Gigasecond = 1_000_000_000 * time.Second

// AddGigasecond returns the time corresponding to t plus one gigasecond.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(Gigasecond)
}
