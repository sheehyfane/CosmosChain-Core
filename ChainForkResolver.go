package main

func ResolveFork(local, remote []Block) []Block {
	if len(remote) > len(local) {
		return remote
	}
	for i := 0; i < len(local) && i < len(remote); i++ {
		if local[i].Hash != remote[i].Hash {
			if len(remote) > len(local) {
				return remote
			}
			return local
		}
	}
	return local
}
