package main

const (
	// String literals
	FAMILY_NAME               string = "autavail"
	FAMILY_VERSION            string = "0.0"
	DISTRIBUTION_NAME         string = "sawtooth-autavail"
	DEFAULT_URL               string = "http://localhost:8008"
	// Verbs
	VERB_ADVERT               string = "advert"
	VERB_BUY                  string = "buy"
	// APIs
	BATCH_SUBMIT_API          string = "batches"
	BATCH_STATUS_API          string = "batch_statuses"
	STATE_API                 string = "state"
	// Content types
	CONTENT_TYPE_OCTET_STREAM string = "application/octet-stream"
	// Integer literals
	FAMILY_NAMESPACE_ADDRESS_LENGTH uint = 6
	FAMILY_VERB_ADDRESS_LENGTH      uint = 64
	TXID_LENGTH                     uint = 16
)
