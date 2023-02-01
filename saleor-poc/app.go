package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/machinebox/graphql"
)

const (
	url   = "https://dev-arise-saleor-api.gonuclei.com/graphql/"
	token = "ehzgrgm6QWqzYB3rwmq5x48XcKXPz7"
)

type ProductCreateInput struct {
	name        string      `json:"name,omitempty"`
	description string      `json:"description,omitempty"`
	productType string      `json:"productType,omitempty"`
	category    string      `json:"category,omitempty"`
	attributes  []attribute `json:"attributes,omitempty"`
}

type attribute struct {
	id        string `json:"id,omitempty"`
	values    string `json:"values,omitempty"`
	plainText string `json:"plainText,omitempty"`
}
type authedTransport struct {
	key     string
	wrapped http.RoundTripper
}

func (t *authedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("authorization", "Bearer "+token)
	return t.wrapped.RoundTrip(req)
}

func Start() error {
	fmt.Println("Hasura Client")
	fmt.Println("**********Query***************")
	// client := graphql.NewClient(url, &http.Client{
	// 	Transport: &authedTransport{
	// 		key:     token,
	// 		wrapped: http.DefaultTransport,
	// 	},
	// })

	// var q struct {
	// 	Products struct {
	// 		Edges []struct {
	// 			Node struct {
	// 				ID string `graphql:"id" json:"id,omitempty"`
	// 			} `json:"node,omitempty"`
	// 		} `json:"edges,omitempty"`
	// 	} `graphql:"products(first: $firstl)"`
	// }

	// variables := map[string]interface{}{
	// 	"firstl": graphql.NewInt(10),
	// }
	// err := client.Query(context.Background(), &q, variables)
	// if err != nil {
	// 	return err
	// }
	//print(q)
	/*
			mutation($input: ProductCreateInput!) {
		  productCreate(input: $input) {
		    product {
		      name
		      id
		    }
		    errors {
		      code
		      message
		      field
		      values
		    }
		  }
		}
	*/

	// var prodcr struct {
	// 	productCreate struct {
	// 		product struct {
	// 			id   string
	// 			name string
	// 		}
	// 	} `graphql:"productCreate(input: $input)"`
	// }

	// variables := map[string]interface{}{
	// 	"input": ProductCreateInput{
	// 		name:        "hubspot-test",
	// 		description: "{\"blocks\":[{\"id\":\"uafR0l_B3B\",\"data\":{\"text\":\"HubSpot offers a full platform of marketing, sales, customer service, and CRM software — plus the methodology, resources, and support — to help businesses grow better. Get started with free tools, and upgrade as you grow.\"},\"type\":\"paragraph\"}]}",
	// 		productType: "UHJvZHVjdFR5cGU6MjU=",
	// 		category:    "Q2F0ZWdvcnk6NDY=",
	// 		attributes: []attribute{
	// 			{
	// 				id:     "QXR0cmlidXRlOjQ0",
	// 				values: "A CRM Tool To connect",
	// 			},
	// 		},
	// 	},
	// }

	// err := client.Mutate(context.Background(), &prodcr, variables)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	/*
		mutation InsertUser($objects: [user_insert_input!]!) {
			insert_user(objects: $objects) {
				id
				name
			}
		}
	*/

	fmt.Println("Machine Box Client")
	fmt.Println("**********Query***************")
	client := graphql.NewClient(url)
	// req := graphql.NewRequest(`
	// 	query fetchProducts($first: Int!,$channel: String!) {
	// 		products(first: $first,channel: $channel) {
	// 			edges {
	// 				node {
	// 				  id
	// 				}
	// 			  }
	// 		}
	// 	}
	// `)

	// req1 := graphql.NewRequest(`
	// mutation($input: ProductCreateInput!) {
	// 	productCreate(input: $input) {
	// 	  product {
	// 		name
	// 		id
	// 	  }

	// 	}
	//   }
	// `)

	req2 := graphql.NewRequest(`
	mutation($ID:ID!,$input: [MetadataInput!]!) {
		updateMetadata(id:$ID,input: $input) {
		  item {
		   metadata{
			key
			value
		  }
		  }
		  errors {
			code
			message
			field
		  }
		}
	  }
	`)
	req2.Header.Set("authorization", "Bearer ehzgrgm6QWqzYB3rwmq5x48XcKXPz7")
	//req.Var("first", 10)
	//req.Var("channel", "nuclei-staging")
	// prodCreatre := ProductCreateInput{
	// 	name:        "hubspot-test",
	// 	description: "{\"blocks\":[{\"id\":\"uafR0l_B3B\",\"data\":{\"text\":\"HubSpot offers a full platform of marketing, sales, customer service, and CRM software — plus the methodology, resources, and support — to help businesses grow better. Get started with free tools, and upgrade as you grow.\"},\"type\":\"paragraph\"}]}",
	// 	productType: "UHJvZHVjdFR5cGU6MjU=",
	// 	category:    "Q2F0ZWdvcnk6NDY=",
	// 	attributes: []attribute{
	// 		{
	// 			id:     "QXR0cmlidXRlOjQ0",
	// 			values: "A CRM Tool To connect",
	// 		},
	// 	},
	// }

	//var reqobj map[string]interface{}

	// reqobj = map[string]interface{}{
	// 	"name":        "hubspot-test",
	// 	"description": "{\"blocks\":[{\"id\":\"uafR0l_B3B\",\"data\":{\"text\":\"HubSpot offers a full platform of marketing, sales, customer service, and CRM software — plus the methodology, resources, and support — to help businesses grow better. Get started with free tools, and upgrade as you grow.\"},\"type\":\"paragraph\"}]}",
	// 	"productType": "UHJvZHVjdFR5cGU6MjU=",
	// 	"category":    "Q2F0ZWdvcnk6NDY=",
	// 	"attributes": []map[string]interface{}{
	// 		{
	// 			"id":        "QXR0cmlidXRlOjQ0",
	// 			"plainText": "A CRM Tool To connect",
	// 		},
	// 		{
	// 			"id":        "QXR0cmlidXRlOjU4",
	// 			"plainText": "DMmd2kzN1-Y",
	// 		},
	// 		{
	// 			"id":     "QXR0cmlidXRlOjYz",
	// 			"values": "18",
	// 		},
	// 		{
	// 			"id":     "QXR0cmlidXRlOjY0",
	// 			"values": "19",
	// 		},
	// 		{
	// 			"id":        "QXR0cmlidXRlOjY1",
	// 			"plainText": "Expect activation within $MIN - $MAX Hrs. We 'll notify you over email. Cancel anytime.",
	// 		},
	// 		{
	// 			"id":        "QXR0cmlidXRlOjc0",
	// 			"plainText": "Quote Delivery Note",
	// 		},
	// 		{
	// 			"id":     "QXR0cmlidXRlOjgx",
	// 			"values": "COMING SOON",
	// 		},
	// 		{
	// 			"id":     "QXR0cmlidXRlOjgz",
	// 			"values": "0087",
	// 		},
	// 		{
	// 			"id":     "QXR0cmlidXRlOjg3",
	// 			"values": "DEFAULT",
	// 		},
	// 	},
	// }

	var reqObjMetada map[string]interface{}
	reqObjMetada = map[string]interface{}{
		"input": []map[string]interface{}{
			{
				"key":   "merchant_description",
				"value": "HubSpot is an American developer and marketer of software products for inbound marketing, sales, and customer service. Hubspot was founded by Brian Halligan and Dharmesh Shah in 2006.",
			},
			{
				"key":   "merchant_id",
				"value": "Hubspot1",
			},
			{
				"key":   "merchant_linkedin_link",
				"value": "https://www.linkedin.com/company/hubspot/",
			},
			{
				"key":   "merchant_name",
				"value": "About Hubspot",
			},
			{
				"key":   "merchant_twitter_link",
				"value": "https://twitter.com/HubSpot",
			},
			{
				"key":   "merchant_website_link",
				"value": "https://www.hubspot.com/products/crm",
			},
			{
				"key":   "tag",
				"value": "Popular CRM",
			},
		},
	}

	//m := structs.Map(&prodCreatre)
	// b, err := json.Marshal(prodCreatre)
	// if err != nil {
	// 	log.Println(err)
	// }
	// var m map[string]interface{}
	// err = json.Unmarshal(b, &m)
	// if err != nil {
	// 	log.Println(err)
	// }
	req2.Var("ID", "UHJvZHVjdDoxODk=")
	req2.Var("input", reqObjMetada["input"])

	var respData interface{}
	if err := client.Run(context.Background(), req2, &respData); err != nil {
		log.Fatalf("failed to fetch todo: %v", err)
	}
	fmt.Println("Products:", respData)
	return nil
}
