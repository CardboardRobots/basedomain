package domain

type EventName string

func (e EventName) String() string { return string(e) }
