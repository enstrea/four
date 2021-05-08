package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	v1 "four/api/book/v1"
	kraerr "github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	transgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	"net/http"
	"testing"
)

func TestBook(t *testing.T) {
	//createBook()
	createBookGrpc()
}

func createBook() {
	in := &v1.CreateBookRequest{}
	data, err := json.Marshal(in)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", "http://127.0.0.1:8000/api.book.v1.Book/CreateBook", bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer rsp.Body.Close()

	var all []byte
	var rspData = make([]byte, 100)
	for {
		n, err := rsp.Body.Read(rspData)
		if n > 0 && err != nil {
			all = append(all, rspData[:n]...)
		} else {
			break
		}
	}
	fmt.Println("status: ", rsp.StatusCode)
	//_, err = rsp.Body.Read(rspData)
	//if err != nil {
	//	panic(err)
	//}
	//out := &v1.CreateBookReply{}
	//e := &errors.StatusError{}
	//err = json.Unmarshal(rspData, e)
	//if err != nil {
	//	panic(err)
	//}

	fmt.Println(fmt.Sprintf("%+v", string(rspData)))
}

func createBookGrpc() {
	conn, err := transgrpc.DialInsecure(
		context.Background(),
		transgrpc.WithEndpoint("127.0.0.1:9000"),
		transgrpc.WithMiddleware(
			middleware.Chain(
				recovery.Recovery(),
			),
		),
	)
	if err != nil {
		panic(err)
	}

	bookClient := v1.NewBookClient(conn)
	reply, err := bookClient.CreateBook(context.Background(), &v1.CreateBookRequest{})
	if err != nil {
		if kraerr.IsInternal(err) {
			fmt.Println("err is internal")
		}
		fmt.Println(fmt.Sprintf("%+v", err))
		return
	}
	fmt.Println(reply)
}
