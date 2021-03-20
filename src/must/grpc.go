package must

import (
	"context"
	"google.golang.org/grpc"
)

func main() {
	grpc.NewClientStream(context.TODO(), nil, nil, "GET")
}
