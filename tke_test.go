package v2

import (
	"fmt"
	"os"
	"testing"
)

func TestUpateImage(t *testing.T) {
	tke := New(os.Getenv("TC_SECRET_ID"), os.Getenv("TC_SECRET_KEY"))
	err := tke.UpdateImage("nginx:latest", "test", "test-svc", os.Getenv("TC_REGION"), os.Getenv("TC_CLUSTER_ID"))
	if err != nil {
		t.Error(err)
	}
}

func TestGetServiceInfo(t *testing.T) {
	tke := New(os.Getenv("TC_SECRET_ID"), os.Getenv("TC_SECRET_KEY"))
	info, err := tke.GetServiceInfo("test", "test-svc", os.Getenv("TC_REGION"), os.Getenv("TC_CLUSTER_ID"))
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", info)
}
