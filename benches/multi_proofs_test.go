package benches

import (
	"encoding/hex"
	"encoding/json"
	"math"
	"os"
	"testing"

	ssz "github.com/ferranbt/fastssz"
	"github.com/ferranbt/fastssz/spectests"
)

// Test blocks just above and below 256, a power of 2.
// 21315748.json contains 247 transactions.
// 21327802.json contains 261 transactions.
var TransactionsJsonPaths = []string{"./test_data/blocks_transactions/21315748.json", "./test_data/blocks_transactions/21327802.json"}

func loadTransactions(path string) (*spectests.ExecutionPayloadTransactions, error) {
	// Open the JSON file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Decode the JSON into a map
	var txs []string
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&txs); err != nil {
		return nil, err
	}

	byteTxs := make([][]byte, len(txs))
	for i := range byteTxs {
		byteTxs[i], _ = hex.DecodeString(txs[i][2:])
	}

	return &spectests.ExecutionPayloadTransactions{Transactions: byteTxs}, nil
}

func BenchmarkMultiProof(b *testing.B) {
	for _, path := range TransactionsJsonPaths {
		// create a sub-benchmark for each path
		b.Run(path, func(b *testing.B) {
			bellatrixPayloadTxs, err := loadTransactions(path)
			if err != nil {
				b.Fatalf("Failed to load transactions for path %s: %v\n", path, err)
			}

			// Setup generalized indexes
			baseGeneralizedIndex := int(math.Pow(float64(2), float64(21)))
			generalizedIndexes := []int{
				baseGeneralizedIndex,
				baseGeneralizedIndex + 3,
			}

			// Reset timer to exclude setup time
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				rootNode, err := bellatrixPayloadTxs.GetTree()
				if err != nil {
					b.Fatalf("Failed to construct tree for transactions: %v\n", err)
				}

				rootNode.Hash()

				multiProof, err := rootNode.ProveMulti(generalizedIndexes)
				if err != nil {
					b.Fatalf("Failed to generate multiproof: %v\n", err)
				}

				// Verify the proof
				if ok, err := ssz.VerifyMultiproof(rootNode.Hash(), multiProof.Hashes, multiProof.Leaves, multiProof.Indices); !ok || err != nil {
					b.Fatalf("NOT OK while verifying multiproof: %v\n", err)
				}
			}
		})
	}
}
