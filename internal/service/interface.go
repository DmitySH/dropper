package service

import "dmitysh/your-drop/internal/entity"

type SendFile interface {
	GenerateAndGetDropCode() string
	CheckSecretCode(code int) (bool, bool)
	SendFileByChunks(filepath string, fileSender ChunkSender) error
}

type GetFile interface {
	ParseDropCode(dropCode string) (entity.DropCode, error)
	ReceiveFileByChunks(fileReceiver ChunkReceiver) ([]byte, error)
}

type ChunkSender interface {
	Send(chunk entity.FileChunk) error
}

type ChunkReceiver interface {
	Receive() (entity.FileChunk, error)
}
