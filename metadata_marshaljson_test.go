package nftmeta_test

import (
	"testing"

	"bytes"
	"encoding/json"
	"math/big"

	"github.com/reiver/go-nftmeta"
)

func TestMetaData_MarshalJSON(t *testing.T) {

	tests := []struct{
		MetaData nftmeta.MetaData
		Expected []byte
	}{
		{
			MetaData: nftmeta.MetaData{},
			Expected: []byte(`{}`),
		},



		{
			MetaData: func()nftmeta.MetaData{
				const value string = "http://example.com/video.mp4"

				var metadata nftmeta.MetaData
				metadata.SetAnimationURL(value) // <-------------

				return metadata
			}(),
			Expected: []byte(`{"animation_url":"http://example.com/video.mp4"}`),
		},
		{
			MetaData: func()nftmeta.MetaData{
				const value string = "#0055BF"

				var metadata nftmeta.MetaData
				metadata.SetBackgroundColor(value) // <-------------

				return metadata
			}(),
			Expected: []byte(`{"background_color":"#0055BF"}`),
		},
		{
			MetaData: func()nftmeta.MetaData{
				const value string = "To explore!"

				var metadata nftmeta.MetaData
				metadata.SetDescription(value) // <-------------

				return metadata
			}(),
			Expected: []byte(`{"description":"To explore!"}`),
		},
		{
			MetaData: func()nftmeta.MetaData{
				const value string = "http://example.com/token/123"

				var metadata nftmeta.MetaData
				metadata.SetExternalLink(value) // <-------------

				return metadata
			}(),
			Expected: []byte(`{"external_link":"http://example.com/token/123"}`),
		},
		{
			MetaData: func()nftmeta.MetaData{
				const value string = "http://example.com/token/123/img.png"

				var metadata nftmeta.MetaData
				metadata.SetImage(value) // <-------------

				return metadata
			}(),
			Expected: []byte(`{"image":"http://example.com/token/123/img.png"}`),
		},
		{
			MetaData: func()nftmeta.MetaData{
				const value string = `<svg width="300" height="130" xmlns="http://www.w3.org/2000/svg"></svg>`

				var metadata nftmeta.MetaData
				metadata.SetImageData(value) // <-------------

				return metadata
			}(),
			Expected: []byte(`{"image_data":"\u003csvg width=\"300\" height=\"130\" xmlns=\"http://www.w3.org/2000/svg\"\u003e\u003c/svg\u003e"}`),
		},
		{
			MetaData: func()nftmeta.MetaData{
				const value string = `peanut-butter-jelly-time`

				var metadata nftmeta.MetaData
				metadata.SetName(value) // <--------------

				return metadata
			}(),
			Expected: []byte(`{"name":"peanut-butter-jelly-time"}`),
		},
		{
			MetaData: func()nftmeta.MetaData{
				const value string = `https://youtu.be/eRBOgtp0Hac`

				var metadata nftmeta.MetaData
				metadata.SetYouTubeURL(value) // <--------------

				return metadata
			}(),
			Expected: []byte(`{"youtube_url":"https://youtu.be/eRBOgtp0Hac"}`),
		},



		{
			MetaData: func()nftmeta.MetaData{
				var metadata nftmeta.MetaData
				metadata.AppendAttribute(nftmeta.AttributeString("Base", "Starfish"))

				return metadata
			}(),
			Expected: []byte(`{"attributes":[{"trait_type":"Base","value":"Starfish"}]}`),
		},
		{
			MetaData: func()nftmeta.MetaData{
				var metadata nftmeta.MetaData
				metadata.AppendAttribute(nftmeta.AttributeString("Base", "Starfish"))
				metadata.AppendAttribute(nftmeta.AttributeString("Big", "Eyes"))

				return metadata
			}(),
			Expected: []byte(`{"attributes":[{"trait_type":"Base","value":"Starfish"},{"trait_type":"Big","value":"Eyes"}]}`),
		},
		{
			MetaData: func()nftmeta.MetaData{
				var metadata nftmeta.MetaData
				metadata.AppendAttribute(nftmeta.AttributeString("Base", "Starfish"))
				metadata.AppendAttribute(nftmeta.AttributeString("Big", "Eyes"))
				metadata.AppendAttribute(nftmeta.AttributeString("Mouth", "Surprised"))

				return metadata
			}(),
			Expected: []byte(`{"attributes":[{"trait_type":"Base","value":"Starfish"},{"trait_type":"Big","value":"Eyes"},{"trait_type":"Mouth","value":"Surprised"}]}`),
		},
		{
			MetaData: func()nftmeta.MetaData{
				var metadata nftmeta.MetaData
				metadata.AppendAttribute(nftmeta.AttributeString("Base", "Starfish"))
				metadata.AppendAttribute(nftmeta.AttributeString("Big", "Eyes"))
				metadata.AppendAttribute(nftmeta.AttributeString("Mouth", "Surprised"))
				metadata.AppendAttribute(nftmeta.AttributeUint64("Level", 5))

				return metadata
			}(),
			Expected: []byte(`{"attributes":[{"trait_type":"Base","value":"Starfish"},{"trait_type":"Big","value":"Eyes"},{"trait_type":"Mouth","value":"Surprised"},{"trait_type":"Level","value":5}]}`),
		},
		{
			MetaData: func()nftmeta.MetaData{
				var metadata nftmeta.MetaData
				metadata.AppendAttribute(nftmeta.AttributeString("Base", "Starfish"))
				metadata.AppendAttribute(nftmeta.AttributeString("Big", "Eyes"))
				metadata.AppendAttribute(nftmeta.AttributeString("Mouth", "Surprised"))
				metadata.AppendAttribute(nftmeta.AttributeUint64("Level", 5))
				metadata.AppendAttribute(nftmeta.AttributeBigFloat("Stamina", big.NewFloat(0).SetRat(big.NewRat(14,10))))

				return metadata
			}(),
			Expected: []byte(`{"attributes":[{"trait_type":"Base","value":"Starfish"},{"trait_type":"Big","value":"Eyes"},{"trait_type":"Mouth","value":"Surprised"},{"trait_type":"Level","value":5},{"trait_type":"Stamina","value":1.4}]}`),
		},
		{
			MetaData: func()nftmeta.MetaData{
				var metadata nftmeta.MetaData
				metadata.AppendAttribute(nftmeta.AttributeString("Base", "Starfish"))
				metadata.AppendAttribute(nftmeta.AttributeString("Big", "Eyes"))
				metadata.AppendAttribute(nftmeta.AttributeString("Mouth", "Surprised"))
				metadata.AppendAttribute(nftmeta.AttributeUint64("Level", 5))
				metadata.AppendAttribute(nftmeta.AttributeBigFloat("Stamina", big.NewFloat(0).SetRat(big.NewRat(14,10))))
				metadata.AppendAttribute(nftmeta.AttributeString("Personality", "Sad"))

				return metadata
			}(),
			Expected: []byte(`{"attributes":[{"trait_type":"Base","value":"Starfish"},{"trait_type":"Big","value":"Eyes"},{"trait_type":"Mouth","value":"Surprised"},{"trait_type":"Level","value":5},{"trait_type":"Stamina","value":1.4},{"trait_type":"Personality","value":"Sad"}]}`),
		},
		{
			MetaData: func()nftmeta.MetaData{
				var metadata nftmeta.MetaData
				metadata.AppendAttribute(nftmeta.AttributeString("Base", "Starfish"))
				metadata.AppendAttribute(nftmeta.AttributeString("Big", "Eyes"))
				metadata.AppendAttribute(nftmeta.AttributeString("Mouth", "Surprised"))
				metadata.AppendAttribute(nftmeta.AttributeUint64("Level", 5))
				metadata.AppendAttribute(nftmeta.AttributeBigFloat("Stamina", big.NewFloat(0).SetRat(big.NewRat(14,10))))
				metadata.AppendAttribute(nftmeta.AttributeString("Personality", "Sad"))
				metadata.AppendAttribute(nftmeta.TypedAttributeInt64("Aqua Power", 40, "boost_number"))

				return metadata
			}(),
			Expected: []byte(`{"attributes":[{"trait_type":"Base","value":"Starfish"},{"trait_type":"Big","value":"Eyes"},{"trait_type":"Mouth","value":"Surprised"},{"trait_type":"Level","value":5},{"trait_type":"Stamina","value":1.4},{"trait_type":"Personality","value":"Sad"},{"display_type":"boost_number","trait_type":"Aqua Power","value":40}]}`),
		},
		{
			MetaData: func()nftmeta.MetaData{
				var metadata nftmeta.MetaData
				metadata.AppendAttribute(nftmeta.AttributeString("Base", "Starfish"))
				metadata.AppendAttribute(nftmeta.AttributeString("Big", "Eyes"))
				metadata.AppendAttribute(nftmeta.AttributeString("Mouth", "Surprised"))
				metadata.AppendAttribute(nftmeta.AttributeUint64("Level", 5))
				metadata.AppendAttribute(nftmeta.AttributeBigFloat("Stamina", big.NewFloat(0).SetRat(big.NewRat(14,10))))
				metadata.AppendAttribute(nftmeta.AttributeString("Personality", "Sad"))
				metadata.AppendAttribute(nftmeta.TypedAttributeInt64("Aqua Power", 40, "boost_number"))
				metadata.AppendAttribute(nftmeta.AttributeInt64("Shift", -3))

				return metadata
			}(),
			Expected: []byte(`{"attributes":[{"trait_type":"Base","value":"Starfish"},{"trait_type":"Big","value":"Eyes"},{"trait_type":"Mouth","value":"Surprised"},{"trait_type":"Level","value":5},{"trait_type":"Stamina","value":1.4},{"trait_type":"Personality","value":"Sad"},{"display_type":"boost_number","trait_type":"Aqua Power","value":40},{"trait_type":"Shift","value":-3}]}`),
		},
		{
			MetaData: func()nftmeta.MetaData{
				var metadata nftmeta.MetaData
				metadata.AppendAttribute(nftmeta.AttributeString("Base", "Starfish"))
				metadata.AppendAttribute(nftmeta.AttributeString("Big", "Eyes"))
				metadata.AppendAttribute(nftmeta.AttributeString("Mouth", "Surprised"))
				metadata.AppendAttribute(nftmeta.AttributeUint64("Level", 5))
				metadata.AppendAttribute(nftmeta.AttributeBigFloat("Stamina", big.NewFloat(0).SetRat(big.NewRat(14,10))))
				metadata.AppendAttribute(nftmeta.AttributeString("Personality", "Sad"))
				metadata.AppendAttribute(nftmeta.TypedAttributeInt64("Aqua Power", 40, "boost_number"))
				metadata.AppendAttribute(nftmeta.AttributeInt64("Shift", -3))
				metadata.AppendAttribute(nftmeta.AttributeBigInt("Max", big.NewInt(0).SetBytes([]byte{0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF})))

				return metadata
			}(),
			Expected: []byte(`{"attributes":[{"trait_type":"Base","value":"Starfish"},{"trait_type":"Big","value":"Eyes"},{"trait_type":"Mouth","value":"Surprised"},{"trait_type":"Level","value":5},{"trait_type":"Stamina","value":1.4},{"trait_type":"Personality","value":"Sad"},{"display_type":"boost_number","trait_type":"Aqua Power","value":40},{"trait_type":"Shift","value":-3},{"trait_type":"Max","value":6277101735386680763835789423207666416102355444464034512895}]}`),
		},
		{
			MetaData: func()nftmeta.MetaData{
				var metadata nftmeta.MetaData
				metadata.AppendAttribute(nftmeta.AttributeString("Base", "Starfish"))
				metadata.AppendAttribute(nftmeta.AttributeString("Big", "Eyes"))
				metadata.AppendAttribute(nftmeta.AttributeString("Mouth", "Surprised"))
				metadata.AppendAttribute(nftmeta.AttributeUint64("Level", 5))
				metadata.AppendAttribute(nftmeta.AttributeBigFloat("Stamina", big.NewFloat(0).SetRat(big.NewRat(14,10))))
				metadata.AppendAttribute(nftmeta.AttributeString("Personality", "Sad"))
				metadata.AppendAttribute(nftmeta.TypedAttributeInt64("Aqua Power", 40, "boost_number"))
				metadata.AppendAttribute(nftmeta.AttributeInt64("Shift", -3))
				metadata.AppendAttribute(nftmeta.AttributeBigInt("Max", big.NewInt(0).SetBytes([]byte{0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF})))
				metadata.AppendAttribute(nftmeta.TypedAttributeUint64("Super Shift", 1000000, "boost_number"))

				return metadata
			}(),
			Expected: []byte(`{"attributes":[{"trait_type":"Base","value":"Starfish"},{"trait_type":"Big","value":"Eyes"},{"trait_type":"Mouth","value":"Surprised"},{"trait_type":"Level","value":5},{"trait_type":"Stamina","value":1.4},{"trait_type":"Personality","value":"Sad"},{"display_type":"boost_number","trait_type":"Aqua Power","value":40},{"trait_type":"Shift","value":-3},{"trait_type":"Max","value":6277101735386680763835789423207666416102355444464034512895},{"display_type":"boost_number","trait_type":"Super Shift","value":1000000}]}`),
		},
		{
			MetaData: func()nftmeta.MetaData{
				var metadata nftmeta.MetaData
				metadata.AppendAttribute(nftmeta.AttributeString("Base", "Starfish"))
				metadata.AppendAttribute(nftmeta.AttributeString("Big", "Eyes"))
				metadata.AppendAttribute(nftmeta.AttributeString("Mouth", "Surprised"))
				metadata.AppendAttribute(nftmeta.AttributeUint64("Level", 5))
				metadata.AppendAttribute(nftmeta.AttributeBigFloat("Stamina", big.NewFloat(0).SetRat(big.NewRat(14,10))))
				metadata.AppendAttribute(nftmeta.AttributeString("Personality", "Sad"))
				metadata.AppendAttribute(nftmeta.TypedAttributeInt64("Aqua Power", 40, "boost_number"))
				metadata.AppendAttribute(nftmeta.AttributeInt64("Shift", -3))
				metadata.AppendAttribute(nftmeta.AttributeBigInt("Max", big.NewInt(0).SetBytes([]byte{0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF})))
				metadata.AppendAttribute(nftmeta.TypedAttributeUint64("Super Shift", 1000000, "boost_number"))
				metadata.AppendAttribute(nftmeta.TypedAttributeString("Pie Color", "#FFFFED", "pie_style"))

				return metadata
			}(),
			Expected: []byte(`{"attributes":[{"trait_type":"Base","value":"Starfish"},{"trait_type":"Big","value":"Eyes"},{"trait_type":"Mouth","value":"Surprised"},{"trait_type":"Level","value":5},{"trait_type":"Stamina","value":1.4},{"trait_type":"Personality","value":"Sad"},{"display_type":"boost_number","trait_type":"Aqua Power","value":40},{"trait_type":"Shift","value":-3},{"trait_type":"Max","value":6277101735386680763835789423207666416102355444464034512895},{"display_type":"boost_number","trait_type":"Super Shift","value":1000000},{"display_type":"pie_style","trait_type":"Pie Color","value":"#FFFFED"}]}`),
		},
		{
			MetaData: func()nftmeta.MetaData{
				var metadata nftmeta.MetaData
				metadata.AppendAttribute(nftmeta.AttributeString("Base", "Starfish"))
				metadata.AppendAttribute(nftmeta.AttributeString("Big", "Eyes"))
				metadata.AppendAttribute(nftmeta.AttributeString("Mouth", "Surprised"))
				metadata.AppendAttribute(nftmeta.AttributeUint64("Level", 5))
				metadata.AppendAttribute(nftmeta.AttributeBigFloat("Stamina", big.NewFloat(0).SetRat(big.NewRat(14,10))))
				metadata.AppendAttribute(nftmeta.AttributeString("Personality", "Sad"))
				metadata.AppendAttribute(nftmeta.TypedAttributeInt64("Aqua Power", 40, "boost_number"))
				metadata.AppendAttribute(nftmeta.AttributeInt64("Shift", -3))
				metadata.AppendAttribute(nftmeta.AttributeBigInt("Max", big.NewInt(0).SetBytes([]byte{0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF})))
				metadata.AppendAttribute(nftmeta.TypedAttributeUint64("Super Shift", 1000000, "boost_number"))
				metadata.AppendAttribute(nftmeta.TypedAttributeString("Pie Color", "#FFFFED", "pie_style"))
				metadata.AppendAttribute(nftmeta.TypedAttributeBigInt("Super Min", big.NewInt(-9090909), "ultimate_combo"))

				return metadata
			}(),
			Expected: []byte(`{"attributes":[{"trait_type":"Base","value":"Starfish"},{"trait_type":"Big","value":"Eyes"},{"trait_type":"Mouth","value":"Surprised"},{"trait_type":"Level","value":5},{"trait_type":"Stamina","value":1.4},{"trait_type":"Personality","value":"Sad"},{"display_type":"boost_number","trait_type":"Aqua Power","value":40},{"trait_type":"Shift","value":-3},{"trait_type":"Max","value":6277101735386680763835789423207666416102355444464034512895},{"display_type":"boost_number","trait_type":"Super Shift","value":1000000},{"display_type":"pie_style","trait_type":"Pie Color","value":"#FFFFED"},{"display_type":"ultimate_combo","trait_type":"Super Min","value":-9090909}]}`),
		},
		{
			MetaData: func()nftmeta.MetaData{
				var metadata nftmeta.MetaData
				metadata.AppendAttribute(nftmeta.AttributeString("Base", "Starfish"))
				metadata.AppendAttribute(nftmeta.AttributeString("Big", "Eyes"))
				metadata.AppendAttribute(nftmeta.AttributeString("Mouth", "Surprised"))
				metadata.AppendAttribute(nftmeta.AttributeUint64("Level", 5))
				metadata.AppendAttribute(nftmeta.AttributeBigFloat("Stamina", big.NewFloat(0).SetRat(big.NewRat(14,10))))
				metadata.AppendAttribute(nftmeta.AttributeString("Personality", "Sad"))
				metadata.AppendAttribute(nftmeta.TypedAttributeInt64("Aqua Power", 40, "boost_number"))
				metadata.AppendAttribute(nftmeta.AttributeInt64("Shift", -3))
				metadata.AppendAttribute(nftmeta.AttributeBigInt("Max", big.NewInt(0).SetBytes([]byte{0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF,0xFF})))
				metadata.AppendAttribute(nftmeta.TypedAttributeUint64("Super Shift", 1000000, "boost_number"))
				metadata.AppendAttribute(nftmeta.TypedAttributeString("Pie Color", "#FFFFED", "pie_style"))
				metadata.AppendAttribute(nftmeta.TypedAttributeBigInt("Super Min", big.NewInt(-9090909), "ultimate_combo"))

				bigfloat, _, err := new(big.Float).SetPrec(62).SetMode(big.ToZero).Parse("3.141592653589793", 10)
				if nil != err {
					panic("could not set string on big-float: "+err.Error())
				}
				metadata.AppendAttribute(nftmeta.TypedAttributeBigFloat("Pie", bigfloat, "magic_number"))

				return metadata
			}(),
			Expected: []byte(`{"attributes":[{"trait_type":"Base","value":"Starfish"},{"trait_type":"Big","value":"Eyes"},{"trait_type":"Mouth","value":"Surprised"},{"trait_type":"Level","value":5},{"trait_type":"Stamina","value":1.4},{"trait_type":"Personality","value":"Sad"},{"display_type":"boost_number","trait_type":"Aqua Power","value":40},{"trait_type":"Shift","value":-3},{"trait_type":"Max","value":6277101735386680763835789423207666416102355444464034512895},{"display_type":"boost_number","trait_type":"Super Shift","value":1000000},{"display_type":"pie_style","trait_type":"Pie Color","value":"#FFFFED"},{"display_type":"ultimate_combo","trait_type":"Super Min","value":-9090909},{"display_type":"magic_number","trait_type":"Pie","value":3.141592653589793}]}`),
		},



		{
			MetaData: func()nftmeta.MetaData{
				var metadata nftmeta.MetaData

				metadata.SetName("super-nft-0000001-holesky")
				metadata.SetDescription("super-nft-token on holesky")

				metadata.AppendAttribute(nftmeta.AttributeString("Maturity", "2024-06-20T18:03:14.636Z"))

				return metadata
			}(),
			Expected: []byte(`{"description":"super-nft-token on holesky","name":"super-nft-0000001-holesky","attributes":[{"trait_type":"Maturity","value":"2024-06-20T18:03:14.636Z"}]}`),
		},
	}

	for testNumber, test := range tests {

		actual, err := json.Marshal(test.MetaData)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("METADATA: %#v", test.MetaData)
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
				t.Logf("METADATA: %#v", test.MetaData)
				continue
			}
		}

		{
			if !json.Valid(actual) {
				t.Errorf("For test #%d, invalid json.", testNumber)
				t.Logf("JSON:\n%s", actual)
				t.Logf("METADATA: %#v", test.MetaData)
				continue
			}
		}
	}
}
