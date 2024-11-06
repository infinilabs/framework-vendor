package main

import (
	"context"
	"fmt"

	"github.com/libdns/tencentcloud"
)

func main() {

	p := tencentcloud.Provider{
		SecretId:  "YOUR_Secret_ID",
		SecretKey: "YOUR_Secret_Key",
	}

	ret, err := p.GetRecords(context.TODO(), "your-domain")

	fmt.Println("Result:", ret)
	fmt.Println("Error:", err)

}
