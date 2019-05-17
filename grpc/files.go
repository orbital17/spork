package grpc_api

import (
	context "context"
	fmt "fmt"
	"io/ioutil"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type Files struct{}

func NewFilesServer() *Files {
	return &Files{}
}

func saveFile(bytes []byte) {
	fmt.Printf("got file %v bytes size\n", len(bytes))
	name := fmt.Sprintf("./im%v.jpg", len(bytes))
	err := ioutil.WriteFile(name, bytes, 0644)
	if err != nil {
		fmt.Print(err)
	}
}

func (*Files) UploadChunk(ctx context.Context, req *UploadChunkRequest) (*Status, error) {
	saveFile(req.Bytes)
	return &Status{Status: true}, nil
}
func (*Files) Save(ctx context.Context, req *InputFile) (*FileLocation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Save not implemented")
}
