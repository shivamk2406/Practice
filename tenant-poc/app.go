package tenantpoc

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/shivamk2406/Practice/tenant-poc/model"
)

type authedTransport struct {
	// key     string // set this key while doing e2e test cases
	wrapped http.RoundTripper
}

func (t *authedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	///req.Header.Set(header.Authorization, "Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJMVmJKSnpudHhyOXBnMWdGeUR3Rk4tN2RhaXdDSkxtNU5VbFZGVnIyb0tRIn0.eyJleHAiOjE2OTc2OTcyMzEsImlhdCI6MTY3NTIzMzIzMSwianRpIjoiZjIwYjVhYWMtNTA5MC00ZTA5LThmNTctZDBhN2E4Y2Q0OTE1IiwiaXNzIjoiaHR0cHM6Ly9kZXYtYXJpc2Uta2V5Y2xvYWt4LmdvbnVjbGVpLmNvbS9hdXRoL3JlYWxtcy9hcmlzZS1iMmItaWFtLXRudC0xIiwic3ViIjoiYTk1MzQ0YmEtZTJlOC00ODI3LWE1YzEtNGZjNWEzMzg1NGEzIiwidHlwIjoiQmVhcmVyIiwiYXpwIjoiYWRtaW4tY2xpIiwic2Vzc2lvbl9zdGF0ZSI6ImZjMDYzNjMzLTY0ZWQtNDkzMC1hODJhLTI0MmJhMWY4OGM0ZSIsImFjciI6IjEiLCJzY29wZSI6InByb2ZpbGUgZW1haWwiLCJzaWQiOiJmYzA2MzYzMy02NGVkLTQ5MzAtYTgyYS0yNDJiYTFmODhjNGUiLCJ0ZW5hbnRfaWQiOiIxZjFmN2U4NC0zZDZhLTExZWQtYjg3OC0wMjQyYWMxMjAwMDIiLCJlbWFpbF92ZXJpZmllZCI6ZmFsc2UsInByZWZlcnJlZF91c2VybmFtZSI6Ii9hb2I1d2NhNTlzLzB0bWl6bm1ob3hvNmRud2ZrKzh4ZXNpNXlzanBqZm9zZ3lnZmJ6N215ZGFpZWsvNjFvbHJuazduNXJobXZudWxjY2N1aDJ3OXNxPT0iLCJjdXN0b21lcl9pZCI6ImE5NTM0NGJhLWUyZTgtNDgyNy1hNWMxLTRmYzVhMzM4NTRhMyIsImVtYWlsIjoiL2FvYjV3Y2E1OXMvMHRtaXpubWhveG82ZG53ZmsrOHhlc2k1eXNqcGpmb3NneWdmYno3bXlkYWllay82MW9scm5rN241cmhtdm51bGNjY3VoMnc5c3E9PSIsInRlbmFudF9yZWNvcmRfaWQiOjF9.KwW-EPEfBZRipiu8k6NxxHuW6VEBrF62MwccTKfiazIvcHx_o2HiS3jbUVGccSaN0AgK6n_q9CCmY7dLCzaNx39ZYI7lEbarbvLo1u9yOxlriJPRaXcJmPsRFth8dwReZfExG0I2FxkkpzH11wT23iAdxRlXpEmZt7qiDv9tFeOqWuhU-xA8NnMRM7dN0PscMT6bUzmkpOZ2xEyCqmB6kKwmIWbUwHH0in6jg1lyeQMIlcb5xGwnlblakXshNmqFv_V1n5wNAAv29TPm3MGN3wHlRDZiqN4LmaEt0Z11AobHENvD_3URkMmTlLJdvjZcFWRIJ7w9FVbvGKjMB-g5Ew")
	return t.wrapped.RoundTrip(req)
}

func Start() error {
	// fmt.Println("***Tenant POC *****")
	// transpor := &http.Transport{
	// 	Proxy: http.ProxyFromEnvironment,
	// 	Dial: (&net.Dialer{
	// 		Timeout:   1 * time.Second,
	// 		KeepAlive: 30 * time.Second,
	// 	}).Dial,
	// 	TLSHandshakeTimeout: 10 * time.Second,
	// }
	// client := graphql.NewClient("http://localhost:8080/tenantb2b/v1/query", &http.Client{
	// 	Transport: &authedTransport{
	// 		wrapped: transpor,
	// 	},
	// 	Timeout: 2 * time.Second,
	// })

	// // tenantIds := []string{
	// // 	"1f1f7e84-3d6a-11ed-b878-0242ac120002",
	// // 	"1f1f7e84-3d6a-11ed-b878-0242ac128765",
	// // 	"32b973c3-966b-4ce9-9316-a37329bc73c1",
	// // 	"5640b4a1-a137-4d56-b06a-ece20d88c1f5",
	// // }

	// domains := []string{
	// 	"arise-dev.gonuclei.com",
	// 	"arise-smart-bank-dev.gonuclei.com",
	// 	"rakbank-dev.gonuclei.com",
	// 	"hsbcbank.gonuclei.com",
	// }
	// //ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	// //defer cancel()

	// //for i := 0; i < 10000; i++ {
	// // tenant, err := graphql_schema.GetTenantById(context.Background(), client, val)
	// // if err != nil {
	// // 	fmt.Println(err)
	// // }
	// // fmt.Println(tenant.GetTenant.DisplayName)
	// //rand.Seed(time.Now().Unix())
	// RandomIntegerwithinRange := rand.Intn(3)
	// tenant1, err := graphql_schema.GetConfig(context.Background(), client, domains[RandomIntegerwithinRange])
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(tenant1.GetTenantByDomain.Name)
	// //	}

	config:=model.Tenant{
		ExchangeRates: []model.ExchangeRate{
			{
				ID: "ID",
				FromCurrency: "INR",
				ToCurrency: "USD",
				StartAt: time.Now(),
				ExpireAt: time.Now().Add(time.Hour*2),
				ExchangeRate: 34.45,
			},
		},
		
	}

	bytees,err:=json.Marshal(config)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(string(bytees))

	file, err := os.Create("Tenant.json")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	file.Write(bytees)
	return nil
}