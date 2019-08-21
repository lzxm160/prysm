package kv

// The schema will define how to store and retrieve data from the db.
// we can prefix or suffix certain values such as `block` with attributes
// for prefix-wide scans across the underlying BoltDB buckets when filtering data.
// For example, we might store attestations as shard + attestation_root -> attestation, making
// it easy to scan for keys that have a certain shard number as a prefix and return those
// corresponding attestations.
var (
	attestationsBucket = []byte("attestations")
	blocksBucket       = []byte("blocks")
	validatorsBucket   = []byte("validators")
	stateBucket        = []byte("state")

	// Key indices buckets.
	blockParentRootIndicesBucket       = []byte("block-parent-root-indices")
	blockSlotIndicesBucket             = []byte("block-slot-indices")
	attestationParentRootIndicesBucket = []byte("attestation-parent-root-indices")
	attestationShardIndicesBucket      = []byte("attestation-shard-indices")
	attestationStartEpochIndicesBucket = []byte("attestation-start-epoch-indices")
	attestationEndEpochIndicesBucket   = []byte("attestation-end-epoch-indices")

	// Block keys.
	headBlockRootKey = []byte("head-root")
)