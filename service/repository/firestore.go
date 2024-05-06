package repository

import (
	"context"
	"fmt"
	"os"
	"refugio/objects"
	"refugio/utils"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/google/uuid"
	"google.golang.org/api/option"
)

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

		doc := collection.Doc(uuid.NewString())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error adding to collection: %v", err)
		}
		bulkWriter.Create(doc, pessoa)
	}

	bulkWriter.End()
	return nil
}

func FetchFromFirestore(docID string) (*objects.PessoaResult, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, os.Getenv("FIRESTORE_PROJECT_ID"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating client: %v", err)
		return nil, err
	}

	pessoas := client.Collection(os.Getenv("FIRESTORE_COLLECTION"))

	pessoa := pessoas.Doc(docID)

	var pessoaResult map[string]interface{}
	snapshot, _ := pessoa.Get(ctx)

	if err := snapshot.DataTo(&pessoaResult); err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching from collection: %v", err)
		return nil, err
	}

	return &objects.PessoaResult{
		Pessoa: &objects.Pessoa{Nome: pessoaResult["Nome"].(string),
			Abrigo: pessoaResult["Abrigo"].(string),
			Idade:  pessoaResult["Idade"].(string)},
		SheetId:   pessoaResult["SheetId"].(string),
		Timestamp: pessoaResult["Timestamp"].(time.Time),
	}, nil
}
