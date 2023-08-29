package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/v6/list"
	"github.com/olekukonko/tablewriter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"cookbook/devto-grpc/repo"
	"cookbook/devto-grpc/service"
)

var addr = flag.String("address", "", "the server addres")

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("dial server error: %v", err)
	}
	client := repo.NewLaptopServiceClient(conn)

	// laptop := service.NewLaptop()
	// req := &repo.CreateLaptopRequest{Laptop: laptop}
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	// defer cancel()
	// res, err := client.CreateLaptop(ctx, req)
	// if err != nil {
	// 	state, ok := status.FromError(err)
	// 	if ok && state.Code() == codes.AlreadyExists {
	// 		log.Printf("laptop [%v] alreay exists", res.Id)
	// 	} else {
	// 		log.Fatalf("create laptop failed: %v", err)
	// 	}
	// }
	// log.Printf("laptop created with id: %s", res.Id)

	for i := 0; i < 50; i++ {
		createLaptop(client)
	}

	filter := &repo.Filter{
		MaxPriceUsed: 3000,
		MinCpuCores:  4,
		MinCpuGhz:    2.5,
		MinRam: &repo.Memory{
			Value: 8,
			Unit:  repo.Memory_GIGABYTE,
		},
	}
	searchLoptop(client, filter)
}

func createLaptop(client repo.LaptopServiceClient) error {
	laptop := service.NewLaptop()
	laptop.Id = ""
	req := &repo.CreateLaptopRequest{Laptop: laptop}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	res, err := client.CreateLaptop(ctx, req)
	if err != nil {
		state, ok := status.FromError(err)
		if ok && state.Code() == codes.AlreadyExists {
			log.Printf("laptop [%v] alreay exists", res.Id)
		} else {
			log.Fatalf("create laptop failed: %v", err)
			return fmt.Errorf("create laptop failed: %w", err)
		}
	}
	log.Printf("laptop created with id: %s", res.Id)
	return nil
}

func searchLoptop(client repo.LaptopServiceClient, filter *repo.Filter) {
	log.Printf("search filter: %v", filter)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req := &repo.SearchLaptopRequest{Filter: filter}
	stream, err := client.SearchLaptop(ctx, req)
	if err != nil {
		log.Fatalf("search laptop failed: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("search stream loop failed: %v", err)
		}

		laptop := res.GetLaptop()
		// tableOutput(laptop)
		prettyOutput(laptop)
	}
}

func tableOutput(l *repo.Laptop) {
	data := [][]string{
		{l.Brand, l.Name},
	}

	table := tablewriter.NewWriter(os.Stdout)
	// table.SetHeader([]string{"Brand", "Name", "Cores", "MinGHz", "RAM", "Price"})
	table.SetHeader([]string{"Brand", "Name"})
	// table.SetAutoWrapText(false)
	// table.SetAutoFormatHeaders(true)
	// table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	// table.SetAlignment(tablewriter.ALIGN_LEFT)
	// table.SetCenterSeparator("")
	// table.SetColumnSeparator("")
	// table.SetRowSeparator("")
	// table.SetHeaderLine(true)
	// table.EnableBorder(true)
	// table.SetTablePadding("\t") // pad with tabs
	// table.SetNoWhiteSpace(true)

	// table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	// table.SetCenterSeparator("|")

	table.AppendBulk(data) // Add Bulk Data
	table.Render()
}

func prettyOutput(l *repo.Laptop) {
	t := list.NewWriter()
	t.SetStyle(list.StyleConnectedRounded)
	// t.AppendItem(laptop.Id)
	// t.AppendItem(laptop.Name)
	// t.AppendItem(laptop.Brand)

	data := []interface{}{
		l.Brand + " " + l.Name,
		l.PriceUsd,
		l.GetCpu().GetCores(),
		l.GetUpdateAt().AsTime().Format(time.RFC3339),
		l.Id,
	}
	t.AppendItems(data)

	for _, line := range strings.Split(t.Render(), "\n") {
		fmt.Printf("%s%s\n", "", line)
	}
	fmt.Println()
}
