package enum

type Status int

const (
	Public Status = iota
	Limited
	Private
)

func (s Status) Ja() string     { return s.Ja() }
func (s Status) Value() string  { return s.get().Str }
func (s Status) String() string { return s.Value() }

type state struct {
	V   Status
	Str string
	Ja  string
}

var statuses = [cnt]state{
	{Public, "public", "公開中"},
	{Limited, "limited", "限定公開"},
	{Private, "private", "非公開"},
}

const cnt = 3

func (s Status) get() state {
	if s < 0 || s >= cnt {
		return statuses[0]
	}
	return statuses[s]
}
