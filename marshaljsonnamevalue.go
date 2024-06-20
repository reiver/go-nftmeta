package nftmeta

import (
	"encoding/json"

	"sourcecode.social/reiver/go-erorr"
)

func appendJSONNameValue(p []byte, name string, value string) ([]byte, error) {

	{
		var str string = name

		bytes, err := json.Marshal(str)
		if nil != err {
			return nil, erorr.Errorf("nftmeta: problem json-marshaling %T: %w", str, err)
		}

		p = append(p, bytes...)
	}

	p = append(p, ':')

	{
		var str string = value

		bytes, err := json.Marshal(str)
		if nil != err {
			return nil, erorr.Errorf("nftmeta: problem json-marshaling %T: %w", str, err)
		}

		p = append(p, bytes...)
	}

	return p, nil
}
