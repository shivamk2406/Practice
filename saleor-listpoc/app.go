package saleorlistpoc

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/Khan/genqlient/graphql"
	"github.com/shivamk2406/Practice/saleor-listpoc/graphql_schema"
)

const (
	url   = "https://dev-arise-saleor-api.gonuclei.com/graphql/"
	token = "ehzgrgm6QWqzYB3rwmq5x48XcKXPz7"
)

type authedTransport struct {
	// key     string // set this key while doing e2e test cases
	wrapped http.RoundTripper
}

func (t *authedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	///req.Header.Set(header.Authorization, "Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJMVmJKSnpudHhyOXBnMWdGeUR3Rk4tN2RhaXdDSkxtNU5VbFZGVnIyb0tRIn0.eyJleHAiOjE2OTc2OTcyMzEsImlhdCI6MTY3NTIzMzIzMSwianRpIjoiZjIwYjVhYWMtNTA5MC00ZTA5LThmNTctZDBhN2E4Y2Q0OTE1IiwiaXNzIjoiaHR0cHM6Ly9kZXYtYXJpc2Uta2V5Y2xvYWt4LmdvbnVjbGVpLmNvbS9hdXRoL3JlYWxtcy9hcmlzZS1iMmItaWFtLXRudC0xIiwic3ViIjoiYTk1MzQ0YmEtZTJlOC00ODI3LWE1YzEtNGZjNWEzMzg1NGEzIiwidHlwIjoiQmVhcmVyIiwiYXpwIjoiYWRtaW4tY2xpIiwic2Vzc2lvbl9zdGF0ZSI6ImZjMDYzNjMzLTY0ZWQtNDkzMC1hODJhLTI0MmJhMWY4OGM0ZSIsImFjciI6IjEiLCJzY29wZSI6InByb2ZpbGUgZW1haWwiLCJzaWQiOiJmYzA2MzYzMy02NGVkLTQ5MzAtYTgyYS0yNDJiYTFmODhjNGUiLCJ0ZW5hbnRfaWQiOiIxZjFmN2U4NC0zZDZhLTExZWQtYjg3OC0wMjQyYWMxMjAwMDIiLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsInByZWZlcnJlZF91c2VybmFtZSI6Ii9hb2I1d2NhNTlzLzB0bWl6bm1ob3hvNmRud2ZrKzh4ZXNpNXlzanBqZm9zZ3lnZmJ6N215ZGFpZWsvNjFvbHJuazduNXJobXZudWxjY2N1aDJ3OXNxPT0iLCJjdXN0b21lcl9pZCI6ImE5NTM0NGJhLWUyZTgtNDgyNy1hNWMxLTRmYzVhMzM4NTRhMyIsImVtYWlsIjoiL2FvYjV3Y2E1OXMvMHRtaXpubWhveG82ZG53ZmsrOHhlc2k1eXNqcGpmb3NneWdmYno3bXlkYWllay82MW9scm5rN241cmhtdm51bGNjY3VoMnc5c3E9PSIsInRlbmFudF9yZWNvcmRfaWQiOjF9.KwW-EPEfBZRipiu8k6NxxHuW6VEBrF62MwccTKfiazIvcHx_o2HiS3jbUVGccSaN0AgK6n_q9CCmY7dLCzaNx39ZYI7lEbarbvLo1u9yOxlriJPRaXcJmPsRFth8dwReZfExG0I2FxkkpzH11wT23iAdxRlXpEmZt7qiDv9tFeOqWuhU-xA8NnMRM7dN0PscMT6bUzmkpOZ2xEyCqmB6kKwmIWbUwHH0in6jg1lyeQMIlcb5xGwnlblakXshNmqFv_V1n5wNAAv29TPm3MGN3wHlRDZiqN4LmaEt0Z11AobHENvD_3URkMmTlLJdvjZcFWRIJ7w9FVbvGKjMB-g5Ew")
	req.Header.Set("Authorization", "bearer "+token)
	return t.wrapped.RoundTrip(req)
}

type Entity struct {
	ID   string
	Name string
}

func Start() error {
	fmt.Println("Saleor POC")
	productMap := make(map[string]string, 0)
	variantMap := make(map[string]string, 0)
	transpor := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   1 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 10 * time.Second,
	}
	client := graphql.NewClient(url, &http.Client{
		Transport: &authedTransport{
			wrapped: transpor,
		},
		Timeout: 2 * time.Second,
	})
	tenant1, err := graphql_schema.ListCategories(context.Background(), client, 100)
	if err != nil {
		fmt.Println(err)
	}

	for _, val := range tenant1.Products.Edges {
		productMap[val.Node.Id] = val.Node.Name
		if (len(val.Node.Variants)) != 0 {
			for _, varia := range val.Node.Variants {
				variantMap[varia.Id] = varia.Name
			}
		}

	}
	var prodStruct []Entity
	for key, val := range productMap {
		prodStruct = append(prodStruct, Entity{
			ID:   key,
			Name: val,
		})
	}
	jsonData, err := json.Marshal(prodStruct)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Create a file and write the JSON data to it
	file, err := os.Create("productMap.json")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	file.Write(jsonData)

	var varStruct []Entity
	for key, val := range variantMap {
		varStruct = append(varStruct, Entity{
			ID:   key,
			Name: val,
		})
	}

	jsonData, err = json.Marshal(varStruct)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// Create a file and write the JSON data to it
	file, err = os.Create("variantMap.json")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	file.Write(jsonData)
	fmt.Println("JSON data written to file:", "person.json")
	return nil
}
