// Licensed to Elasticsearch B.V under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.
//
// Code generated, DO NOT EDIT

package elasticsearch_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	_ = fmt.Printf
	_ = os.Stdout
	_ = elasticsearch.NewDefaultClient
)

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/getting-started.asciidoc#L169>
//
// --------------------------------------------------------------------------------
// GET /_cat/health?v
// --------------------------------------------------------------------------------

func Test_getting_started_f8cc4b331a19ff4df8e4a490f906ee69(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:f8cc4b331a19ff4df8e4a490f906ee69[]
	res, err := es.Cat.Health(es.Cat.Health.WithV(true))
	fmt.Println(res, err)
	// end:f8cc4b331a19ff4df8e4a490f906ee69[]
	if err != nil {
		t.Fatalf("Error getting the response: %s", err)
	}
	defer res.Body.Close()
}