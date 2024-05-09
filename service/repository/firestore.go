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

var err error

func createClient(ctx context.Context) (*firestore.Client, error) {
	var client *firestore.Client
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
	var client *firestore.Client
	client, err = createClient(ctx)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating client: %v", err)
		client.Close()
		return err
	}
	defer client.Close()

	bulkWriter := client.BulkWriter(ctx)

	collection := client.Collection(os.Getenv("FIRESTORE_COLLECTION"))
	fmt.Fprintf(os.Stdout, "Adding %d documents to Firestore collection %v\n", len(pessoas), collection.Path)
	jobs := make([]*firestore.BulkWriterJob, 0, len(pessoas))
	for _, pessoa := range pessoas {
		doc := collection.Doc(pessoa.AggregateKey())
		job, err := bulkWriter.Set(doc, &pessoa)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create job: %v", err)
			return err
		}
		jobs = append(jobs, job)
	}

	bulkWriter.End()
	for _, i := range jobs {
		_, err := i.Results()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get job results: %v)", err)
		}
	}
	return nil
}

func FetchPessoaFromFirestore(docIDs []string) ([]*objects.PessoaResult, error) {
	ctx := context.Background()
	var client *firestore.Client
	client, err = createClient(ctx)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating client: %v", err)
		client.Close()
		return nil, err
	}
	defer client.Close()

	pessoas := client.Collection(os.Getenv("FIRESTORE_COLLECTION"))
	refs := make([]*firestore.DocumentRef, 0, len(docIDs))

	for _, id := range docIDs {
		refs = append(refs, pessoas.Doc(id))
	}

	docs, err := client.GetAll(ctx, refs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to retrieve documents: %v", err)
	}

	var results []*objects.PessoaResult
	for _, doc := range docs {
		if doc.Exists() {
			var data map[string]interface{}
			if err := doc.DataTo(&data); err != nil {
				fmt.Fprintf(os.Stderr, "Failed to read document: %v", err)
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
	var client *firestore.Client
	client, err = createClient(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating client: %v", err)
		client.Close()
		return err
	}
	defer client.Close()

	bulkWriter := client.BulkWriter(ctx)

	collection := client.Collection(os.Getenv("FIRESTORE_SOURCES_COLLECTION"))
	fmt.Fprintf(os.Stdout, "Adding %d documents to Firestore collection %v\n", len(sources), collection.Path)
	for _, source := range sources {
		doc := collection.Doc(source.URL + source.SheetId)
		bulkWriter.Set(doc, &source)
	}

	bulkWriter.End()
	return nil
}
