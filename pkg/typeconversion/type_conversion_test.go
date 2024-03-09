package typeconversion

import (
	"testing"
)

type FooDB struct {
	ID    uint64   `db:"id"`
	Name  string   `db:"name"`
	Items []string `db:"items"`
}

type FooJSON struct {
	ID    uint64   `json:"id"`
	Name  string   `json:"name"`
	Items []string `json:"items"`
}

func TestPlainTypeConversion(t *testing.T) {
	fooDB := &FooDB{
		ID:   1,
		Name: "albert43",
		Items: []string{
			"item1",
			"item2",
		},
	}
	fooJSON := *(*FooJSON)(fooDB) // fooDB and fooJSON share the same memory address
	t.Logf("db=%p json=%p\n", fooDB, &fooJSON)
}

// FooJSONSwapFieldsOrder is invalid because the field names are in different order
type FooJSONSwapFieldsOrder struct {
	Name  string   `json:"name"`
	ID    uint64   `json:"id"`
	Items []string `json:"items"`
}

func TestPlainTypeConversionWithSwapOrder(t *testing.T) {
	fooDB := &FooDB{
		ID:   1,
		Name: "albert43",
		Items: []string{
			"item1",
			"item2",
		},
	}
	_ = fooDB
	// fooJSON := *(*FooJSONInvalid)(fooDB) // This will cause compile error
	// t.Logf("db=%p json=%p\n", fooDB, &fooJSON)
}

type FooJSONWithExtraFields struct {
	ID      uint64   `json:"id"`
	Name    string   `json:"name"`
	Items   []string `json:"items"`
	CurTime string   `json:"cur_time"`
}

func TestWithExtraData(t *testing.T) {
	fooDB := &FooDB{
		ID:   1,
		Name: "albert43",
		Items: []string{
			"item1",
			"item2",
		},
	}
	_ = fooDB
	// fooJSON := *(*FooJSONWithExtraFields)(fooDB) // This will cause compile error due to extra fields in target struct
	// t.Logf("db=%p json=%p\n", fooDB, &fooJSON)
}

type FooDBWithExtraFields struct {
	ID           uint64   `db:"id"`
	Name         string   `db:"name"`
	Items        []string `db:"items"`
	SaveUnixTime int64    `db:"save_unix_time"`
}

func TestWithExtraData2(t *testing.T) {
	fooDB := &FooDBWithExtraFields{
		ID:   1,
		Name: "albert43",
		Items: []string{
			"item1",
			"item2",
		},
		SaveUnixTime: 123456,
	}
	_ = fooDB
	// fooJSON := *(*FooJSON)(fooDB) // This will cause compile error due to extra fields in source struct
	// t.Logf("db=%p json=%p\n", fooDB, &fooJSON)
}

type FooJSONWithAlterNames struct {
	ID        uint64   `json:"id"`
	AlterName string   `json:"alter_name"`
	Items     []string `json:"items"`
}

func TestWithAlterNames(t *testing.T) {
	fooDB := &FooDB{
		ID:   1,
		Name: "albert43",
		Items: []string{
			"item1",
			"item2",
		},
	}
	_ = fooDB
	// fooJSON := *(*FooJSONWithAlterNames)(fooDB) // This will cause compile error due to the field name mismatch
	// t.Logf("db=%p json=%p\n", fooDB, &fooJSON)
}

type FooDBWithExtra struct {
	FooDB
	SharedData  []byte `db:"shared_data"`
	PrivateData []byte `db:"private_data"`
}

func TestUsingComposition(t *testing.T) {
	fooDBWithExtra := &FooDBWithExtra{
		FooDB: FooDB{
			ID:   1,
			Name: "albert43",
			Items: []string{
				"item1",
				"item2",
			},
		},
		SharedData:  []byte("shared data"),
		PrivateData: []byte("private data"),
	}
	fooJSON := *(*FooJSON)(&fooDBWithExtra.FooDB) // valid because FooDBWithExtra has a FooDB field using composition
	t.Logf("db=%p json=%p\n", fooDBWithExtra, &fooJSON)
}
