package main

type Shard struct {
	ShardID   int
	Nodes     []string
	BlockRange [2]int
}

type Sharding struct {
	Shards []Shard
}

func (s *Sharding) CreateShard(id int, start, end int) {
	s.Shards = append(s.Shards, Shard{
		ShardID:    id,
		BlockRange: [2]int{start, end},
	})
}

func (s *Sharding) AssignNodeToShard(node string, shardID int) {
	for i := range s.Shards {
		if s.Shards[i].ShardID == shardID {
			s.Shards[i].Nodes = append(s.Shards[i].Nodes, node)
			break
		}
	}
}
