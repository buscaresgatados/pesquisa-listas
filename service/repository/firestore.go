package repository

import (
	"context"
	"fmt"
	"os"
	"refugio/objects"
	"refugio/utils"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

/* Firestore collections */
const (
	PessoasAbrigos = "PessoasAbrigos"
	Sources        = "Sources"
	Filters        = "Filters"
)

var (
	err    error
	client *firestore.Client
)

func createClient(ctx context.Context) (*firestore.Client, error) {
	if os.Getenv("ENVIRONMENT") == "local" {
		serviceAccJSON := utils.GetServiceAccountJSON(os.Getenv("APP_SERVICE_ACCOUNT_JSON"))
		client, err = firestore.NewClient(ctx, os.Getenv("FIRESTORE_PROJECT_ID"), option.WithCredentialsJSON(serviceAccJSON))
	} else {
		client, err = firestore.NewClient(ctx, os.Getenv("FIRESTORE_PROJECT_ID"))
	}
	return client, err
}

func AddPessoasToFirestore(pessoas []*objects.PessoaResult) error {
	ctx := context.Background()
	client, err = createClient(ctx)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating client: %v\n", err)
		client.Close()
		return err
	}
	defer client.Close()

	bulkWriter := client.BulkWriter(ctx)

	collection := client.Collection(PessoasAbrigos)
	fmt.Fprintf(os.Stdout, "Adding %d documents to Firestore collection %v\n", len(pessoas), collection.Path)
	jobs := make([]*firestore.BulkWriterJob, 0, len(pessoas))
	for _, pessoa := range pessoas {
		doc := collection.Doc(pessoa.AggregateKey())
		job, err := bulkWriter.Set(doc, &pessoa)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create job: %v\n", err)
			return err
		}
		jobs = append(jobs, job)
	}

	bulkWriter.End()
	for _, i := range jobs {
		_, err := i.Results()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get job results: %v\n", err)
		}
	}
	return nil
}

func FetchPessoaFromFirestore(docIDs []string) ([]*objects.PessoaResult, error) {
	ctx := context.Background()
	client, err = createClient(ctx)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating client: %v\n", err)
		client.Close()
		return nil, err
	}
	defer client.Close()

	pessoas := client.Collection(PessoasAbrigos)
	refs := make([]*firestore.DocumentRef, 0, len(docIDs))

	for _, id := range docIDs {
		refs = append(refs, pessoas.Doc(id))
	}

	docs, err := client.GetAll(ctx, refs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to retrieve documents: %v\n", err)
	}

	var results []*objects.PessoaResult
	for _, doc := range docs {
		if doc.Exists() {
			var data map[string]interface{}
			if err := doc.DataTo(&data); err != nil {
				fmt.Fprintf(os.Stderr, "Failed to read document: %v\n", err)
			}
			sheetId, ok := data["SheetId"].(string)
			if !ok {
				sheetId = ""
			}
			url, ok := data["URL"].(string)
			if !ok {
				url = ""
			}
			results = append(results, &objects.PessoaResult{
				Pessoa: &objects.Pessoa{
					Nome:   data["Nome"].(string),
					Abrigo: data["Abrigo"].(string),
					Idade:  data["Idade"].(string),
				},
				SheetId:   &sheetId,
				URL:       &url,
				Timestamp: data["Timestamp"].(time.Time),
			})
		} else {
			fmt.Fprintln(os.Stderr, "Document does not exist")
		}
	}

	return results, nil
}

func AddSourcesToFirestore(sources []*objects.Source) error {
	ctx := context.Background()
	client, err = createClient(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating client: %v\n", err)
		client.Close()
		return err
	}
	defer client.Close()

	bulkWriter := client.BulkWriter(ctx)

	collection := client.Collection(Sources)
	fmt.Fprintf(os.Stdout, "Adding %d documents to Firestore collection %v\n", len(sources), collection.Path)
	for _, source := range sources {
		doc := collection.Doc(source.URL + source.SheetId)
		bulkWriter.Set(doc, &source)
	}

	bulkWriter.End()
	return nil
}

func FetchSourcesFromFirestore() ([]*objects.Source, error) {
	ctx := context.Background()
	client, err = createClient(ctx)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating client: %v", err)
		client.Close()
		return nil, err
	}
	defer client.Close()

	sources := client.Collection(os.Getenv("FIRESTORE_SOURCES_COLLECTION"))
	docs, err := sources.Documents(ctx).GetAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to retrieve documents: %v", err)
	}

	var results []*objects.Source
	for _, doc := range docs {
		if doc.Exists() {
			var data map[string]interface{}
			if err := doc.DataTo(&data); err != nil {
				fmt.Fprintf(os.Stderr, "Failed to read document: %v", err)
			}
			results = append(results, &objects.Source{
				Nome:    data["Nome"].(string),
				URL:     data["URL"].(string),
				SheetId: data["SheetId"].(string),
			})
		} else {
			fmt.Fprintln(os.Stderr, "Document does not exist")
		}
	}
	return results, nil
}

func FetchFilterFromFirestore(key string) ([]byte, error) {
	ctx := context.Background()
	client, err = createClient(ctx)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating client: %v\n", err)
		client.Close()
		return nil, err
	}
	defer client.Close()

	filterCollection := client.Collection(Filters)
	doc := filterCollection.Doc(key)
	docSnap, err := doc.Get(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to retrieve document: %v\n", err)
		return nil, err
	}
	var data map[string][]byte
	if err := docSnap.DataTo(&data); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read document: %v\n", err)
	}

	if filter, ok := data["filter"]; !ok {
		fmt.Fprintf(os.Stderr, "Filter not found in document\n")
		return nil, fmt.Errorf("filter not found in document")
	} else {
		return filter, nil
	}
}

func UpdateFilterOnFirestore(key string, data []byte) error {
	ctx := context.Background()
	client, err = createClient(ctx)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating client: %v\n", err)
		client.Close()
		return err
	}
	defer client.Close()

	filterCollection := client.Collection(Filters)
	doc := filterCollection.Doc(key)
	_, err := doc.Set(ctx, map[string][]byte{"filter": data})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to update document: %v\n", err)
		return err
	}
	fmt.Fprintf(os.Stdout, "Filter %s updated successfully\n", key)
	return nil
}

func FetchMostRecent(key string) (*time.Time, error) {
	ctx := context.Background()
	client, err = createClient(ctx)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating client: %v\n", err)
		client.Close()
	}
	defer client.Close()

	collection := client.Collection(PessoasAbrigos)
	query := collection.Query.OrderBy("Timestamp", firestore.Desc).Limit(1)
	docs, err := query.Documents(ctx).GetAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to retrieve documents: %v\n", err)
		return nil, err
	}

	for _, doc := range docs {
		if doc.Exists() {
			var data map[string]interface{}
			if err := doc.DataTo(&data); err != nil {
				fmt.Fprintf(os.Stderr, "Failed to read document: %v\n", err)
			}
			if timestamp, ok := data["Timestamp"].(time.Time); !ok {
				fmt.Fprintf(os.Stderr, "Timestamp not found in document\n")
				return nil, fmt.Errorf("timestamp not found in document")
			} else {
				return &timestamp, nil
			}
		}
	}
	return nil, err
}
