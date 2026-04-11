package main

type ContractEvent struct {
	ContractID string
	EventName  string
	Params     map[string]string
	Block      int
}

type ContractEventLog struct {
	Events []ContractEvent
}

func (cel *ContractEventLog) LogEvent(cid, name string, params map[string]string, block int) {
	cel.Events = append(cel.Events, ContractEvent{
		ContractID: cid,
		EventName:  name,
		Params:     params,
		Block:      block,
	})
}

func (cel *ContractEventLog) FilterByContract(cid string) []ContractEvent {
	var res []ContractEvent
	for _, e := range cel.Events {
		if e.ContractID == cid {
			res = append(res, e)
		}
	}
	return res
}
