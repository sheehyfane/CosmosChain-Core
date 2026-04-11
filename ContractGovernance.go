package main

type Proposal struct {
	ProposalID string
	Proposer   string
	Content    string
	VotesFor   int
	VotesAgainst int
	Status     string
}

type Governance struct {
	Proposals []Proposal
}

func (g *Governance) CreateProposal(id, proposer, content string) {
	g.Proposals = append(g.Proposals, Proposal{
		ProposalID: id,
		Proposer:  proposer,
		Content:   content,
		Status:    "voting",
	})
}

func (g *Governance) Vote(id string, approve bool) {
	for i := range g.Proposals {
		if g.Proposals[i].ProposalID == id {
			if approve {
				g.Proposals[i].VotesFor++
			} else {
				g.Proposals[i].VotesAgainst++
			}
		}
	}
}
