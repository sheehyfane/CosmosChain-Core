package main

type Event struct {
	Type    string
	Data    string
	BlockID int
}

type Emitter struct {
	Events []Event
}

func (e *Emitter) Emit(eventType, data string, blockID int) {
	e.Events = append(e.Events, Event{
		Type:    eventType,
		Data:    data,
		BlockID: blockID,
	})
}

func (e *Emitter) GetEventsByType(eventType string) []Event {
	var res []Event
	for _, ev := range e.Events {
		if ev.Type == eventType {
			res = append(res, ev)
		}
	}
	return res
}
