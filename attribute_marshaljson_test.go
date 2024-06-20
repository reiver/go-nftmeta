package nftmeta_test

import (
	"testing"

	"bytes"
	"encoding/json"
	"math/big"

	"github.com/reiver/go-nftmeta"
)

func TestAttribute_MarshalJSON(t *testing.T) {

	tests := []struct{
		Attribute nftmeta.Attribute
		Expected []byte
	}{
		{
			Attribute: nftmeta.AttributeString("", ""),
			Expected: []byte(`{"trait_type":"","value":""}`),
		},
		{
			Attribute: nftmeta.AttributeString("key", ""),
			Expected: []byte(`{"trait_type":"key","value":""}`),
		},
		{
			Attribute: nftmeta.AttributeString("", "the-value"),
			Expected: []byte(`{"trait_type":"","value":"the-value"}`),
		},
		{
			Attribute: nftmeta.AttributeString("apple", "ONE"),
			Expected: []byte(`{"trait_type":"apple","value":"ONE"}`),
		},
		{
			Attribute: nftmeta.AttributeString("banana", "TWO"),
			Expected: []byte(`{"trait_type":"banana","value":"TWO"}`),
		},
		{
			Attribute: nftmeta.AttributeString("cherry", "THREE"),
			Expected: []byte(`{"trait_type":"cherry","value":"THREE"}`),
		},



		{
			Attribute: nftmeta.AttributeInt64("", -2),
			Expected: []byte(`{"trait_type":"","value":-2}`),
		},
		{
			Attribute: nftmeta.AttributeInt64("something", -1),
			Expected: []byte(`{"trait_type":"something","value":-1}`),
		},
		{
			Attribute: nftmeta.AttributeInt64("ZERO", 0),
			Expected: []byte(`{"trait_type":"ZERO","value":0}`),
		},
		{
			Attribute: nftmeta.AttributeInt64("ONE", 1),
			Expected: []byte(`{"trait_type":"ONE","value":1}`),
		},
		{
			Attribute: nftmeta.AttributeInt64("TWO", 2),
			Expected: []byte(`{"trait_type":"TWO","value":2}`),
		},



		{
			Attribute: nftmeta.AttributeUint64("ZERO", 0),
			Expected: []byte(`{"trait_type":"ZERO","value":0}`),
		},
		{
			Attribute: nftmeta.AttributeUint64("ONE", 1),
			Expected: []byte(`{"trait_type":"ONE","value":1}`),
		},
		{
			Attribute: nftmeta.AttributeUint64("TWO", 2),
			Expected: []byte(`{"trait_type":"TWO","value":2}`),
		},



		{
			Attribute: nftmeta.AttributeBigInt("", big.NewInt(-2)),
			Expected: []byte(`{"trait_type":"","value":-2}`),
		},
		{
			Attribute: nftmeta.AttributeBigInt("something", big.NewInt(-1)),
			Expected: []byte(`{"trait_type":"something","value":-1}`),
		},
		{
			Attribute: nftmeta.AttributeBigInt("ZERO", big.NewInt(0)),
			Expected: []byte(`{"trait_type":"ZERO","value":0}`),
		},
		{
			Attribute: nftmeta.AttributeBigInt("ONE", big.NewInt(1)),
			Expected: []byte(`{"trait_type":"ONE","value":1}`),
		},
		{
			Attribute: nftmeta.AttributeBigInt("TWO", big.NewInt(2)),
			Expected: []byte(`{"trait_type":"TWO","value":2}`),
		},









		{
			Attribute: nftmeta.TypedAttributeString("", "", "super_string"),
			Expected: []byte(`{"display_type":"super_string","trait_type":"","value":""}`),
		},
		{
			Attribute: nftmeta.TypedAttributeString("key", "", "super_string"),
			Expected: []byte(`{"display_type":"super_string","trait_type":"key","value":""}`),
		},
		{
			Attribute: nftmeta.TypedAttributeString("", "the-value", "super_string"),
			Expected: []byte(`{"display_type":"super_string","trait_type":"","value":"the-value"}`),
		},
		{
			Attribute: nftmeta.TypedAttributeString("apple", "ONE", "super_string"),
			Expected: []byte(`{"display_type":"super_string","trait_type":"apple","value":"ONE"}`),
		},
		{
			Attribute: nftmeta.TypedAttributeString("banana", "TWO", "super_string"),
			Expected: []byte(`{"display_type":"super_string","trait_type":"banana","value":"TWO"}`),
		},
		{
			Attribute: nftmeta.TypedAttributeString("cherry", "THREE", "super_string"),
			Expected: []byte(`{"display_type":"super_string","trait_type":"cherry","value":"THREE"}`),
		},



		{
			Attribute: nftmeta.TypedAttributeInt64("", -2, "boost_percentage"),
			Expected: []byte(`{"display_type":"boost_percentage","trait_type":"","value":-2}`),
		},
		{
			Attribute: nftmeta.TypedAttributeInt64("something", -1, "boost_percentage"),
			Expected: []byte(`{"display_type":"boost_percentage","trait_type":"something","value":-1}`),
		},
		{
			Attribute: nftmeta.TypedAttributeInt64("ZERO", 0, "boost_percentage"),
			Expected: []byte(`{"display_type":"boost_percentage","trait_type":"ZERO","value":0}`),
		},
		{
			Attribute: nftmeta.TypedAttributeInt64("ONE", 1, "boost_percentage"),
			Expected: []byte(`{"display_type":"boost_percentage","trait_type":"ONE","value":1}`),
		},
		{
			Attribute: nftmeta.TypedAttributeInt64("TWO", 2, "boost_percentage"),
			Expected: []byte(`{"display_type":"boost_percentage","trait_type":"TWO","value":2}`),
		},



		{
			Attribute: nftmeta.TypedAttributeUint64("ZERO", 0, "boost_percentage"),
			Expected: []byte(`{"display_type":"boost_percentage","trait_type":"ZERO","value":0}`),
		},
		{
			Attribute: nftmeta.TypedAttributeUint64("ONE", 1, "boost_percentage"),
			Expected: []byte(`{"display_type":"boost_percentage","trait_type":"ONE","value":1}`),
		},
		{
			Attribute: nftmeta.TypedAttributeUint64("TWO", 2, "boost_percentage"),
			Expected: []byte(`{"display_type":"boost_percentage","trait_type":"TWO","value":2}`),
		},



		{
			Attribute: nftmeta.TypedAttributeBigInt("", big.NewInt(-2), "boost_percentage"),
			Expected: []byte(`{"display_type":"boost_percentage","trait_type":"","value":-2}`),
		},
		{
			Attribute: nftmeta.TypedAttributeBigInt("something", big.NewInt(-1), "boost_percentage"),
			Expected: []byte(`{"display_type":"boost_percentage","trait_type":"something","value":-1}`),
		},
		{
			Attribute: nftmeta.TypedAttributeBigInt("ZERO", big.NewInt(0), "boost_percentage"),
			Expected: []byte(`{"display_type":"boost_percentage","trait_type":"ZERO","value":0}`),
		},
		{
			Attribute: nftmeta.TypedAttributeBigInt("ONE", big.NewInt(1), "boost_percentage"),
			Expected: []byte(`{"display_type":"boost_percentage","trait_type":"ONE","value":1}`),
		},
		{
			Attribute: nftmeta.TypedAttributeBigInt("TWO", big.NewInt(2), "boost_percentage"),
			Expected: []byte(`{"display_type":"boost_percentage","trait_type":"TWO","value":2}`),
		},
	}

	for testNumber, test := range tests {

		actual, err := json.Marshal(test.Attribute)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("ATTRIBUTE: %#v", test.Attribute)
			continue
		}

		{
			expected := test.Expected

			if !bytes.Equal(expected, actual) {
				t.Errorf("For test #%d, the actual marshaled-json is not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("EXPECTED:\n%s", expected)
				t.Logf("ACTUAL:\n%s", actual)
				t.Logf("ATTRIBUTE: %#v", test.Attribute)
				continue
			}
		}

		{
			if !json.Valid(actual) {
				t.Errorf("For test #%d, invalid json.", testNumber)
				t.Logf("JSON:\n%s", actual)
				t.Logf("ATTRIBUTE: %#v", test.Attribute)
				continue
			}
		}
	}
}
