package message

import (
	"bytes"
	"testing"
)

func TestAnyNonceMessage_GetTransactions(t *testing.T) {
	// Prepare test data
	nonceBytes := make(map[[2]byte][][]byte)
	nnb := [2]byte{}
	copy(nnb[:], "nn")
	nonceBytes[nnb] = [][]byte{}
	anyNonceMessage := AnyNonceMessage{
		BaseMessage: BaseMessage{
			Chain:   1,
			Head:    []byte("nn"),
			ChainID: 100,
		},
		NonceBytes: nonceBytes,
	}
	// Call GetTransactions method
	transactions := anyNonceMessage.GetTransactions()
	// Check if the result is as expected
	if len(transactions) != 0 {
		t.Errorf("Expected 0 transactions, got %d", len(transactions))
	}
}
func TestAnyNonceMessage_GetChain(t *testing.T) {
	// Prepare test data
	anyNonceMessage := AnyNonceMessage{
		BaseMessage: BaseMessage{
			Chain: 1,
		},
	}
	// Call GetChain method
	chain := anyNonceMessage.GetChain()
	// Check if the result is as expected
	if chain != 1 {
		t.Errorf("Expected chain 1, got %d", chain)
	}
}
func TestAnyNonceMessage_GetHead(t *testing.T) {
	// Prepare test data
	anyNonceMessage := AnyNonceMessage{
		BaseMessage: BaseMessage{
			Head: []byte("nn"),
		},
	}
	// Call GetHead method
	head := anyNonceMessage.GetHead()
	// Check if the result is as expected
	if !bytes.Equal(head, []byte("nn")) {
		t.Errorf("Expected head 'nn', got %s", string(head))
	}
}
func TestAnyNonceMessage_GetChainID(t *testing.T) {
	// Prepare test data
	anyNonceMessage := AnyNonceMessage{
		BaseMessage: BaseMessage{
			ChainID: 100,
		},
	}
	// Call GetChainID method
	chainID := anyNonceMessage.GetChainID()
	// Check if the result is as expected
	if chainID != 100 {
		t.Errorf("Expected chainID 100, got %d", chainID)
	}
}
func TestAnyNonceMessage_GetValidHead(t *testing.T) {
	// Prepare test data
	anyNonceMessage := AnyNonceMessage{}
	// Call GetValidHead method
	validHead := anyNonceMessage.GetValidHead()
	// Check if the result is as expected
	if len(validHead) != 3 {
		t.Errorf("Expected 3 valid heads, got %d", len(validHead))
	}
}
func TestAnyNonceMessage_GetBytes(t *testing.T) {
	// Prepare test data
	anyNonceMessage := AnyNonceMessage{
		BaseMessage: BaseMessage{
			Chain:   1,
			Head:    []byte("nn"),
			ChainID: 100,
		},
		NonceBytes: make(map[[2]byte][][]byte),
	}
	// Call GetBytes method
	bytes := anyNonceMessage.GetBytes()
	// Check if the result is as expected
	if len(bytes) == 0 {
		t.Error("Expected non-empty byte slice, got empty")
	}
}
func TestAnyNonceMessage_GetFromBytes(t *testing.T) {
	// Prepare test data
	anyNonceMessage := AnyNonceMessage{
		BaseMessage: BaseMessage{
			Chain:   1,
			Head:    []byte("nn"),
			ChainID: 100,
		},
		NonceBytes: make(map[[2]byte][][]byte),
	}
	inputBytes := anyNonceMessage.GetBytes()
	// Call GetFromBytes method
	err := anyNonceMessage.GetFromBytes(inputBytes)
	// Check if there is no error
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	// Check if the result is as expected
	if anyNonceMessage.GetChain() != 1 {
		t.Errorf("Expected chain 1, got %d", anyNonceMessage.GetChain())
	}
	if !bytes.Equal(anyNonceMessage.GetHead(), []byte("nn")) {
		t.Errorf("Expected head 'nn', got %s", string(anyNonceMessage.GetHead()))
	}
	if anyNonceMessage.GetChainID() != 100 {
		t.Errorf("Expected chainID 100, got %d", anyNonceMessage.GetChainID())
	}
}
