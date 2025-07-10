package dtos

import "time"

type Date struct {
	time.Time
}

const layout = "2006-01-02"

func (d *Date) UnmarshalJSON(b []byte) error {
	s := string(b)
	s = s[1 : len(s)-1]

	t, err := time.Parse(layout, s)
	if err != nil {
		return err
	}

	d.Time = t

	return nil
}
