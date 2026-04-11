package main

import "errors"

func SyncNode(local, remote *Blockchain) error {
	if len(remote.Chain) <= len(local.Chain) {
		return errors.New("remote chain not longer")
	}
	local.Chain = remote.Chain
	local.Difficulty = remote.Difficulty
	return nil
}

func GetMissingBlocks(local, remote []Block) []Block {
	if len(remote) <= len(local) {
		return nil
	}
	return remote[len(local):]
}
