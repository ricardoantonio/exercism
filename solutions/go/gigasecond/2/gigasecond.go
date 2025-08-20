package gigasecond

// import path for the time package from the standard library
import "time"

const GIGASECOND = 1_000_000_000 * time.Second // 1 gigasecond = 1 billion seconds

// AddGigasecond takes a time.Time value and returns a new time.Time value
// after adding one gigasecond (1 billion seconds) to it.
func AddGigasecond(t time.Time) time.Time {
	t = t.Add(GIGASECOND)
	return t
}
