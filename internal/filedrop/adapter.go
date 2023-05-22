package filedrop

import "dmitysh/your-drop/internal/entity"

type StreamSender struct {
	gRPCFileStream FileDrop_GetFileServer
}

func NewStreamSender(fileStream FileDrop_GetFileServer) *StreamSender {
	return &StreamSender{gRPCFileStream: fileStream}
}

func (s *StreamSender) Send(chunk entity.FileChunk) error {
	return s.gRPCFileStream.Send(&FileRequest{ChunkData: chunk})
}

type StreamReceiver struct {
	gRPCFileStream FileDrop_GetFileClient
}

func NewStreamReceiver(fileStream FileDrop_GetFileClient) *StreamReceiver {
	return &StreamReceiver{gRPCFileStream: fileStream}
}

func (s *StreamReceiver) Receive() (entity.FileChunk, error) {
	fileChunk, recvErr := s.gRPCFileStream.Recv()
	if recvErr != nil {
		return entity.FileChunk{}, recvErr
	}

	return fileChunk.GetChunkData(), nil
}
