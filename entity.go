package domain

type Entity[I id] struct {
	id      I
	version Version
	events  []Event
}

type id interface {
	String() string
}

func (e Entity[I]) ID() I            { return e.id }
func (e Entity[I]) Version() Version { return e.version }

func NewEntity[I id](
	id I,
	version Version,
) Entity[I] {
	return Entity[I]{
		id:      id,
		version: version,
	}
}

func (ee *Entity[I]) Events() []Event { return ee.events }

func (ee *Entity[I]) ClearEvents() {
	ee.events = nil
}

func (ee *Entity[I]) PushEvent(e Event) Event {
	if ee.events == nil {
		ee.events = make([]Event, 0)
	}

	ee.events = append(ee.events, e)
	return e
}
