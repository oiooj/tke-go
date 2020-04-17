# Tencent Kubernetes Engine v2 API library

[![GoDoc](https://godoc.org/github.com/oiooj/tke-go?status.svg)](https://godoc.org/github.com/oiooj/tke-go)

Detail doc: https://cloud.tencent.com/document/product/457/41100

Status: Beta


```
package main

import (
        "fmt"
		"os"

        "github.com/oiooj/tke-go/v2"
)

func main() {
	tke := v2.New(os.Getenv("TC_SECRET_ID"), os.Getenv("TC_SECRET_KEY"))
	info, err := tke.GetServiceInfo("test", "test-svc", os.Getenv("TC_REGION"), os.Getenv("TC_CLUSTER_ID"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", info)
}
```

