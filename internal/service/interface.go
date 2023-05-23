package service

import "dmitysh/dropper/internal/entity"

type SendFile interface {
	GenerateAndGetDropCode() string
	CheckSecretCode(code int) (bool, bool)
	SendFileByChunks(filepath string, fileSender ChunkSender) error
}

type GetFile interface {
	ParseDropCode(dropCode string) (entity.DropCode, error)
	ReceiveAndSaveFileByChunks(fileReceiver ChunkReceiver, filepath string) error
}

type Archive interface {
	FolderToTempZIPArchive(folderPath string) (string, error)
}

type ChunkSender interface {
	Send(chunk entity.FileChunk) error
}

type ChunkReceiver interface {
	Receive() (entity.FileChunk, error)
	Meta() (map[string]string, error)
}
