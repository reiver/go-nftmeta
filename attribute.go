package nftmeta

import (
	"encoding/json"
	"math/big"
	"strconv"

	"sourcecode.social/reiver/go-erorr"
	"sourcecode.social/reiver/go-opt"
)

// Attribute represents an 'attribute' in the "attributes" array of the NFT metadata JSON.
type Attribute struct {
	displayType opt.Optional[string]
	traitType   opt.Optional[string]
	value       interface{}
}

func AttributeBigFloat(traitType string, value *big.Float) Attribute {
	return Attribute{
		traitType: opt.Something(traitType),
		value:     big.NewFloat(0).Set(value),
	}
}
func TypedAttributeBigFloat(traitType string, value *big.Float, displayType string) Attribute {
	return Attribute{
		displayType: opt.Something(displayType),
		traitType:   opt.Something(traitType),
		value:       big.NewFloat(0).Set(value),
	}
}

func AttributeBigInt(traitType string, value *big.Int) Attribute {
	return Attribute{
		traitType: opt.Something(traitType),
		value:     big.NewInt(0).Set(value),
	}
}
func TypedAttributeBigInt(traitType string, value *big.Int, displayType string) Attribute {
	return Attribute{
		displayType: opt.Something(displayType),
		traitType:   opt.Something(traitType),
		value:       big.NewInt(0).Set(value),
	}
}

func AttributeInt64(traitType string, value int64) Attribute {
	return Attribute{
		traitType: opt.Something(traitType),
		value:     value,
	}
}
func TypedAttributeInt64(traitType string, value int64, displayType string) Attribute {
	return Attribute{
		displayType: opt.Something(displayType),
		traitType:   opt.Something(traitType),
		value:       value,
	}
}

func AttributeString(traitType string, value string) Attribute {
	return Attribute{
		traitType: opt.Something(traitType),
		value:     value,
	}
}
func TypedAttributeString(traitType string, value string, displayType string) Attribute {
	return Attribute{
		displayType: opt.Something(displayType),
		traitType: opt.Something(traitType),
		value:     value,
	}
}

func AttributeUint64(traitType string, value uint64) Attribute {
	return Attribute{
		traitType: opt.Something(traitType),
		value:     value,
	}
}
func TypedAttributeUint64(traitType string, value uint64, displayType string) Attribute {
	return Attribute{
		displayType: opt.Something(displayType),
		traitType: opt.Something(traitType),
		value:     value,
	}
}

func (receiver Attribute) MarshalJSON() ([]byte, error) {
	var buffer [256]byte
	var p []byte = buffer[0:0]

	if receiver.traitType.IsNothing() {
		return nil, errTraitTypeNothing
	}

	p = append(p, '{')

	{
		value, something := receiver.displayType.Get()
		if something {
			p = append(p, `"display_type":`...)
			{
				bytes, err := json.Marshal(value)
				if nil != err {
					return nil, erorr.Errorf("nftmeta: problem marshaling %T: %s", value, err)
				}
				p = append(p, bytes...)
			}
			p = append(p, ',')
		}
	}

	{
		p = append(p, `"trait_type":`...)

		value, something := receiver.traitType.Get()
		if something {
			bytes, err := json.Marshal(value)
			if nil != err {
				return nil, erorr.Errorf("nftmeta: problem marshaling %T: %s", value, err)
			}
			p = append(p, bytes...)
		} else {
			p = append(p, `nil`...)
		}
	}

	p = append(p, ',')

	{
		p = append(p, `"value":`...)

		value := receiver.value
		if nil == value {
			p = append(p, `nil`...)
		} else {
			switch casted := value.(type) {
			case json.Marshaler:
				bytes, err := casted.MarshalJSON()
				if nil != err {
					return nil, err
				}
				p = append(p, bytes...)
			case string:
				bytes, err := json.Marshal(casted)
				if nil != err {
					return nil, err
				}
				p = append(p, bytes...)
			case int64:
				p = append(p, strconv.FormatInt(casted, 10)...)
			case uint64:
				p = append(p, strconv.FormatUint(casted, 10)...)
			case float64:
				p = append(p, strconv.FormatFloat(casted, 'f', -1, 64)...)
//			case *big.Int:
//				if nil == casted {
//					p = append(p, `null`...)
//				} else {
//					p = append(p, casted.String()...)
//				}
			case *big.Float:
				if nil == casted {
					p = append(p, `null`...)
				} else {
					p = append(p, casted.Text('f', -1)...)
				}
			default:
				return nil, erorr.Errorf("nftmeta: cannot json-marshal something of type %T", value)
			}
		}
	}

	p = append(p, '}')

	return p, nil
}
