package valueobject

import "time"

type DateTime struct {
	value time.Time
}

func NewDateTime(t time.Time) DateTime {
	return DateTime{value: t.UTC()}
}

func NowDateTime() DateTime {
	return DateTime{value: time.Now().UTC()}
}

func (dt DateTime) Value() time.Time {
	return dt.value
}

func (dt DateTime) String() string {
	return dt.value.Format(time.RFC3339)
}

func (dt DateTime) IsZero() bool {
	return dt.value.IsZero()
}

func (dt DateTime) Before(other DateTime) bool {
	return dt.value.Before(other.value)
}

func (dt DateTime) After(other DateTime) bool {
	return dt.value.After(other.value)
}

func (dt DateTime) Add(duration time.Duration) DateTime {
	return DateTime{value: dt.value.Add(duration)}
}

func (dt DateTime) Unix() int64 {
	return dt.value.Unix()
}
