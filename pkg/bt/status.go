package bt

type Status int

const (
	Success Status = iota
	Failure
	Running
)

func (s Status) String() string {
	switch s {
	case Success:
		return "success"
	case Failure:
		return "failure"
	case Running:
		return "running"
	default:
		return "unknown"
	}
}
