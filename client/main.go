package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	calculatorpb "example.com/cloud-native-grpc-calculator/gen/example.com/cloud-native-grpc-calculator/gen/calculator"
)

func main() {
	addr := getenv("TARGET_ADDR", "localhost:50051")
	if len(os.Args) < 4 {
		fmt.Println("usage: client <add|sub> <a> <b>")
		os.Exit(1)
	}
	op := os.Args[1]
	a := parseInt(os.Args[2])
	b := parseInt(os.Args[3])

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to %s: %v", addr, err)
	}
	defer conn.Close()

	c := calculatorpb.NewCalculatorClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	switch op := op; op {
	case "add":
		resp, err := c.Add(ctx, &calculatorpb.AddRequest{A: int32(a), B: int32(b)})
		if err != nil { log.Fatalf("Add error: %v", err) }
		fmt.Println(resp.Result)
	case "sub":
		resp, err := c.Sub(ctx, &calculatorpb.SubRequest{A: int32(a), B: int32(b)})
		if err != nil { log.Fatalf("Sub error: %v", err) }
		fmt.Println(resp.Result)
	case "mul":
		resp, err := c.Mul(ctx, &calculatorpb.MulRequest{A: int32(a), B: int32(b)})
		if err != nil { log.Fatalf("Mul error: %v", err) }
		fmt.Println(resp.Result)
	case "div":
		resp, err := c.Div(ctx, &calculatorpb.DivRequest{A: int32(a), B: int32(b)})
		if err != nil { log.Fatalf("Div error: %v", err) }
		fmt.Println(resp.Result)
	default:
		log.Fatalf("unknown op: %s", op)
	}
}

func parseInt(s string) int {
	var v int
	_, err := fmt.Sscanf(s, "%d", &v)
	if err != nil { log.Fatalf("invalid int: %s", s) }
	return v
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" { return v }
	return def
}
