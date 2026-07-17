package model_test

import (
	"testing"

	"github.com/cdvelop/model"
)

// TestItem represents a single generated record.
type TestItem struct{}

func (t *TestItem) Schema() any {
	return "test_schema"
}

// TestItemList represents a generated list of TestItems.
type TestItemList struct {
	items []model.Model
}

func (l *TestItemList) Len() int {
	return len(l.items)
}

func (l *TestItemList) At(i int) model.Model {
	return l.items[i]
}

func (l *TestItemList) Append(m model.Model) {
	l.items = append(l.items, m)
}

func (l *TestItemList) Decode(data []byte) error {
	// Simple mock decoding logic that populates items
	if len(data) > 0 {
		l.items = append(l.items, &TestItem{})
	}
	return nil
}

// Verification with a compile-time assertion that the generated list implements model.ModelSlice:
var _ model.ModelSlice = (*TestItemList)(nil)

var sampleWirePayload = []byte("sample payload")

func decodeInto(t *testing.T, dec model.Decodable, payload []byte) {
	err := dec.Decode(payload)
	if err != nil {
		t.Fatalf("decode failed: %v", err)
	}
}

func TestModelSliceNamesTheListBoundary(t *testing.T) {
	// El uso con forma de consumidor: una función que exige el contrato combinado,
	// como lo hará view.New(..., newList func() model.ModelSlice, ...).
	var newList func() model.ModelSlice = func() model.ModelSlice { return &TestItemList{} }

	list := newList()
	// 1. Llega del wire (lado Decodable) — usar el codec real de los tests del repo.
	decodeInto(t, list, sampleWirePayload)
	// 2. Se itera sin reflexión (lado FielderSlice).
	if list.Len() == 0 {
		t.Fatal("decoded list must be iterable")
	}
	_ = list.At(0).Schema()
}
