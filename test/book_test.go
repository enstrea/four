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
	"io/ioutil"
	"net/http"
	"testing"
)

func TestBook(t *testing.T) {
	bookId1 := createBook()
	bookId2 := createBookGrpc()

	getBook(bookId1)
	getBookGrpc(bookId2)
}

func createBook() string {
	fmt.Println("------------------------------ createBook --------------------------------------------")
	book := &v1.CreateBookRequest{
		ISBN:   "isbn-isbn-isbn-http",
		Name:   "name-name-name-http",
		Author: "author-author-http",
	}
	data, err := json.Marshal(book)
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

	fmt.Println("createBook status: ", rsp.StatusCode)
	if rsp.StatusCode == 200 {
		rpsData, err := ioutil.ReadAll(rsp.Body)
		if err != nil {
			fmt.Println(fmt.Sprintf("createBook read body failed: %+v", err))
			return ""
		}

		out := &v1.CreateBookReply{}
		if err = json.Unmarshal(rpsData, out); err != nil {
			fmt.Println(fmt.Sprintf("createBook json.Unmarshal failed: %+v", err))
			return ""
		}
		fmt.Println(fmt.Sprintf("createBook book: %+v", out.Book))
		return out.Book.Id
	}

	return ""
}

func createBookGrpc() string {
	fmt.Println("------------------------------ createBookGrpc --------------------------------------------")
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
	book := &v1.CreateBookRequest{
		ISBN:   "isbn-isbn-isbn-grpc",
		Name:   "name-name-name-grpc",
		Author: "author-author-grpc",
	}

	reply, err := bookClient.CreateBook(context.Background(), book)
	if err != nil {
		if kraerr.IsInternal(err) {
			fmt.Println("err is internal")
		}
		fmt.Println(fmt.Sprintf("createBookGrpc: %+v", err))
		return ""
	}

	fmt.Println("createBookGrpc: ", reply)
	return reply.Book.Id
}

func getBook(id string) {
	fmt.Println("------------------------------ getBook --------------------------------------------")
	data, err := json.Marshal(&v1.GetBookRequest{
		Id: id,
	})
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", "http://127.0.0.1:8000/api.book.v1.Book/GetBook", bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer rsp.Body.Close()

	fmt.Println("getBook status: ", rsp.StatusCode)
	if rsp.StatusCode == 200 {
		rspData, err := ioutil.ReadAll(rsp.Body)
		if err != nil {
			fmt.Println(fmt.Sprintf("getBook read body failed: %+v", err))
			return
		}
		out := &v1.GetBookReply{}
		if err = json.Unmarshal(rspData, out); err != nil {
			fmt.Println(fmt.Sprintf("getBook json.Unmarshal failed: %+v", err))
		}
		fmt.Println(fmt.Sprintf("getBook book: %+v", out.Book))
	}
}

func getBookGrpc(id string) {
	fmt.Println("------------------------------ getBookGrpc --------------------------------------------")
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

	reply, err := bookClient.GetBook(context.Background(), &v1.GetBookRequest{Id: id})
	if err != nil {
		if kraerr.IsInternal(err) {
			fmt.Println("err is internal")
		}
		fmt.Println(fmt.Sprintf("getBookGrpc: %+v", err))
		return
	}

	fmt.Println("getBookGrpc: ", reply)
}
