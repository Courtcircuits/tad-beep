package quickwit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func CreateIndex(endpoint string, payload CreateIndexPayload) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	r, err := http.Post(endpoint, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	readbody, err := io.ReadAll(r.Body)
	if r.StatusCode != 200 {
		return fmt.Errorf("%s", readbody)
	}
	defer r.Body.Close()
	return nil
}

type CreateIndexPayload struct {
	Version          string            `json:"version,omitempty"`
	IndexID          string            `json:"index_id,omitempty"`
	IndexURI         string            `json:"index_uri,omitempty"`
	DocMapping       *DocMapping       `json:"doc_mapping,omitempty"`
	IndexingSettings *IndexingSettings `json:"indexing_settings,omitempty"`
	SearchSettings   *SearchSettings   `json:"search_settings,omitempty"`
	Retention        *Retention        `json:"retention,omitempty"`
}

type DocMapping struct {
	// Collection of field mapping, each having its own data type (text, binary, datetime, bool, i64, u64, f64, ip, json).
	FieldMappings  []map[string]any `json:"field_mappings,omitempty"`
	Mode           Mode             `json:"mode,omitempty"`
	DynamicMapping *DynamicMapping  `json:"dynamic_mapping,omitempty"`
	TagFields      []any            `json:"tag_fields,omitempty"`
	StoreSource    bool             `json:"store_source,omitempty"`
	// rfc3339
	TimestampField     string `json:"timestamp_field,omitempty"`
	PartitionKey       string `json:"partition_key,omitempty"`
	MaxNumPartitions   int    `json:"max_num_partitions,omitempty"`
	IndexFieldPresence bool   `json:"index_field_presence,omitempty"`
}

type IndexingSettings struct {
	CommitTimeoutSecs        int          `json:"commit_timeout_secs,omitempty"`
	SplitNumDocsTarget       int          `json:"split_num_docs_target,omitempty"`
	MergePolicy              *MergePolicy `json:"merge_policy,omitempty"`
	ResourcesHeapSize        int          `json:"resources.heap_size,omitempty"`
	DocstoreCompressionLevel int          `json:"docstore_compression_level,omitempty"`
	DocstoreBlockSize        int          `json:"docstore_block_size,omitempty"`
}

type Retention struct {
	Period   string `json:"period,omitempty"`
	Schedule string `json:"schedule,omitempty"`
}

type SearchSettings struct {
	DefaultSearchFields []string `json:"default_search_fields,omitempty"`
}

type Mode string

const (
	DYNAMIC Mode = "dynamic"
	LENIENT Mode = "lenient"
	STRICT  Mode = "strict"
)

type FieldMapping struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
	Fast bool   `json:"fast,omitempty"`
}

type DynamicMapping struct {
	Indexed    bool   `json:"indexed,omitempty"`
	Stored     bool   `json:"stored,omitempty"`
	Tokenizer  string `json:"tokenizer,omitempty"`
	Record     string `json:"record,omitempty"`
	ExpandDots bool   `json:"expand_dots,omitempty"`
	Fast       bool   `json:"fast,omitempty"`
}

type MergePolicy struct {
	MergeFactor      int    `json:"merge_factor,omitempty"`
	MaxMergeFactor   int    `json:"max_merge_factor,omitempty"`
	MinLevelNumDocs  int    `json:"min_level_num_docs,omitempty"`
	MaturationPeriod string `json:"maturation_period,omitempty"`
}

func CheckIndexExists(indexID string, endpoint string) error {
	r, err := http.Get(fmt.Sprintf("%s/api/v1/indexes/%s", endpoint, indexID))
	if err != nil {
		return err
	}
	readbody, err := io.ReadAll(r.Body)
	if r.StatusCode != 200 {
		return fmt.Errorf("%s", readbody)
	}
	defer r.Body.Close()
	return nil
}
