package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/google/uuid"
)

// In-memory Search engine

// Your organization has started a new tech blog with interesting tech stories and you’re responsible for designing and implementing an in-memory search engine, supporting the search functionality on the blog content.

// Requirements:
// -----
// It should be possible to create a dataset in the search engine.
// It should be possible to insert and delete documents in a given dataset. Each document is simply a piece of text.
// It should be possible to search through documents for a search pattern in a given dataset.
// It should be possible to order the search results
// Recency (Most recent -> Past)
// Frequency (Highest -> Low)

// Dataset(Grouping) | Document (Unique identifiers)
// D1 Id1 -> Apple is a fruit
// D1 Id2 -> Razorpay is a Software

// D2 id1 -> Mechanical engineering data

// Tech blog
// Software (datasets)
// Mechanical Engineering (datasets)

// Query: <search_text> in Dataset
// Signature:
// insert(document_string: string, created_time: time, dataset_id: dataset)
// query(seach_text: string, dataset: string) -> []document

// Tech blog ; []dataset
// Dataset : []document
// Document: input_string, timestamp

// SearchQuery : abc
// Print all the dataset which is having this query text


type document struct {
	id string
	text string
	created_at time.Time
}

type dataset struct {
	name string
	documents map[string]*document
}

// Tech blog
// Software (datasets)
// Mechanical Engineering (datasets)

type techBlog struct {
	name string
	datasetMap map[string]*dataset
}

type createDatasetRequestDto struct {
	name string
	datasets *dataset
}

type InsertDocumentRequestDto struct {
	datasetName string
	documents []*document
}

type DeleteDocumentRequestDto struct {
	datasetName string
	documents []*document
}

type QueryRequestDto struct {
	searchString string
	datasetName string
}

type OrderDocumentsRequestDto struct {
	searchString string
	datasetName string
	orderByType string
}

func(t *techBlog) CreateDataset( requestDto createDatasetRequestDto) (*dataset, error){
	_, ok := t.datasetMap[requestDto.name]
	if ok{
		return nil, errors.New(fmt.Sprintf("dataset with name %s already exists",requestDto.name))
	}

	t.datasetMap[requestDto.name]=requestDto.datasets

	return t.datasetMap[requestDto.name], nil
}

func (t *techBlog) InsertDocuments( requestDto InsertDocumentRequestDto) (error){
	//getDataset
	_, ok := t.datasetMap[requestDto.datasetName]
	if !ok {
		return errors.New(fmt.Sprintf("no such dataset exists with namre %s", requestDto.datasetName))
	}

	for _, doc := range requestDto.documents{
		t.datasetMap[requestDto.datasetName].documents[uuid.NewString()]=doc
	}
	return nil
}

func (t *techBlog) DeleteDocuments( requestDto DeleteDocumentRequestDto) (error){
	//getDataset
	_, ok := t.datasetMap[requestDto.datasetName]
	if !ok {
		return errors.New(fmt.Sprintf("no such dataset exists with namre %s", requestDto.datasetName))
	}

	for _, doc := range requestDto.documents{
		delete(t.datasetMap[requestDto.datasetName].documents,doc.id)
	}
	return nil
}

//query(seach_text: string, dataset: string) -> []document
func(t *techBlog) Query (requestDto QueryRequestDto)([]*document, error){
	val, ok := t.datasetMap[requestDto.datasetName]
	if !ok {
		return nil,  errors.New(fmt.Sprintf("no such dataset exists with namre %s", requestDto.datasetName))
	}

	response := make ([]*document, 0)

	for _, val := range val.documents{
		if strings.Contains(val.text, requestDto.searchString) {
			response = append(response, val)
		}
	}

	return response, nil
}


func (t *techBlog) OrderDocument(requestDto OrderDocumentsRequestDto) ([]*document, error){
	_, ok := t.datasetMap[requestDto.datasetName]
	if !ok {
		return nil,  errors.New(fmt.Sprintf("no such dataset exists with namre %s", requestDto.datasetName))
	}

	doc, err:=t.Query(QueryRequestDto{
		searchString: requestDto.searchString,
		datasetName: requestDto.datasetName,
	})

	if err!=nil{
		return nil, err
	}

	if requestDto.orderByType=="RECENCY"{
		return OrderByRecency(doc), nil
	}else if requestDto.orderByType=="FREQUENCY"{
		return OrderByFrequency(doc), nil
	} else {
		return nil, errors.New(fmt.Sprintf("order type undefined", requestDto.orderByType))
	}

}

func OrderByRecency(documentMap []*document) ([]*document){
	timeDocList:=make([]*document, 0)

	for _, val := range documentMap{
		timeDocList=append(timeDocList, val)
	}

	sort.Slice(timeDocList,func(i, j int) bool { return timeDocList[i].created_at.Before(timeDocList[j].created_at)})
	return timeDocList

}

func OrderByFrequency(documentMap []*document) ([]*document){
	docFreqencyMap:= make(map[string]int, 0)
	documentIdMap := make(map[string]*document,0)

	for _, val := range documentMap{
		docFreqencyMap[val.id]=docFreqencyMap[val.id] + 1
		documentIdMap[val.id]=val
	}

	type Pair struct {
		key string
		val int
	}

	pairs:= make([]*Pair, 0)

	for key, val := range docFreqencyMap{
		pairs = append(pairs, &Pair{
			key: key,
			val: val,
		})
	}

	sort.Slice(pairs,func(i, j int) bool { return pairs[i].val>pairs[j].val})

	frequencyDocList:=make([]*document, 0)

	for _, val := range pairs{
		frequencyDocList = append(frequencyDocList, documentIdMap[val.key])
	}

	return frequencyDocList
}	

func main() {

	techBlogInit := techBlog{
		name: "Technical",
		datasetMap: map[string]*dataset{},
	}

	id1:= uuid.NewString()
	id2:= uuid.NewString()
	id3:= uuid.NewString()


	techBlogInit.CreateDataset(createDatasetRequestDto{
		name: "Software",
		datasets:&dataset{
			name: "Mechanical",
			documents: map[string]*document{
				id1:&document{
					id: id1,
					text: "This is a sample text",
					created_at: time.Now(),
				},
				id2:&document{
					id: id2,
					text: "Text is a document",
					created_at: time.Now(),
				},
			},
		},
	})

	techBlogInit.CreateDataset(createDatasetRequestDto{
		name: "Software",
		datasets:&dataset{
			name: "Technical",
			documents: map[string]*document{
				id1:&document{
					id: id1,
					text: "This is a sample text",
					created_at: time.Now(),
				},
				id3:&document{
					id: id3,
					text: "apple",
					created_at: time.Now(),
				},
			},
		},
	})

	techBlogInit.InsertDocuments(InsertDocumentRequestDto{
		datasetName: "Technical",
		documents: []*document{
			&document{
				id: uuid.NewString(),
				text: "this is norhing",
				created_at: time.Now(),
			},
		},
	})

	docs, err:=techBlogInit.Query(QueryRequestDto{
		searchString: "apple",
		datasetName: "Technical",
	})
	if err!=nil{
		fmt.Printf(err.Error())
	}
	fmt.Println(docs)

	docs, err=techBlogInit.Query(QueryRequestDto{
		searchString: "is",
		datasetName: "Mechanical",
	})
	
	
}
