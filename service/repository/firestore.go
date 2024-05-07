package repository

import (
	"context"
	"fmt"
	"os"
	"refugio/objects"
	"refugio/utils"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
	"google.golang.org/api/option"
)

var err error

func AddToFirestore(pessoas []*objects.PessoaResult) error {
	ctx := context.Background()
	serviceAccJSON := utils.GetServiceAccountJSON(os.Getenv("APP_SERVICE_ACCOUNT_JSON"))
	client, err := firestore.NewClient(ctx, os.Getenv("FIRESTORE_PROJECT_ID"), option.WithCredentialsJSON(serviceAccJSON))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating client: %v", err)
		return err
	}
	defer client.Close()

	bulkWriter := client.BulkWriter(ctx)

	collection := client.Collection(os.Getenv("FIRESTORE_COLLECTION"))
	for _, pessoa := range pessoas {
		if pessoa.Nome == "" || pessoa.Abrigo == "" || len(strings.Split(pessoa.Nome, " ")) == 1 {
			continue
		}
		doc := collection.Doc(uuid.NewString())
		bulkWriter.Create(doc, pessoa)
	}

	bulkWriter.End()
	return nil
}

func FetchFromFirestore(docIDs []string) ([]*objects.PessoaResult, error) {
	ctx := context.Background()
	var client *firestore.Client
	if os.Getenv("ENVIRONMENT") == "local" {
		serviceAccJSON := utils.GetServiceAccountJSON(os.Getenv("APP_SERVICE_ACCOUNT_JSON"))
		client, err = firestore.NewClient(ctx, os.Getenv("FIRESTORE_PROJECT_ID"), option.WithCredentialsJSON(serviceAccJSON))
	} else {
		client, err = firestore.NewClient(ctx, os.Getenv("FIRESTORE_PROJECT_ID"))
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating client: %v", err)
		return nil, err
	}

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
			results = append(results, &objects.PessoaResult{
				Pessoa: &objects.Pessoa{Nome: data["Nome"].(string),
					Abrigo: data["Abrigo"].(string),
					Idade:  data["Idade"].(string)},
				SheetId:   data["SheetId"].(*string),
				URL:       data["URL"].(*string),
				Timestamp: data["Timestamp"].(time.Time),
			})
		} else {
			fmt.Fprintln(os.Stderr, "Document does not exist")
		}
	}

	return results, nil
}
