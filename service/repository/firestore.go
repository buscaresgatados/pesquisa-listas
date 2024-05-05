package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"refugio/objects"
	"refugio/utils"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func AddToFirestore(pessoa *objects.PessoaResult) (interface{}, error) {
	ctx := context.Background()
	serviceAccJSON := utils.GetServiceAccountJSON(os.Getenv("APP_SERVICE_ACCOUNT_JSON"))
	client, err := firestore.NewClient(ctx, os.Getenv("FIRESTORE_PROJECT_ID"), option.WithCredentialsJSON(serviceAccJSON))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating client: %v", err)
		return nil, err
	}

	pessoas := client.Collection(os.Getenv("FIRESTORE_COLLECTION"))

	ref, result, err := pessoas.Add(ctx, pessoa)
	fmt.Fprintln(os.Stdout, ref)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error adding to collection: %v", err)
		return nil, err
	}
	return result, err
}

func FetchFromFirestore(name string) ([]*objects.PessoaResult, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, os.Getenv("FIRESTORE_PROJECT_ID"))
	if err != nil {
		// TODO: Handle error.
	}

	pessoas := client.Collection(os.Getenv("FIRESTORE_COLLECTION"))
	it := pessoas.Documents(ctx)

	var pessoaResults []*objects.PessoaResult
	log.Panic()
	for {
		docSnap, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return pessoaResults, err
		}

		var exp objects.PessoaResult
		if err := docSnap.DataTo(&exp); err != nil {
			return pessoaResults, err
		}

		pessoaResults = append(pessoaResults, &exp)
	}

	return pessoaResults, nil
}
