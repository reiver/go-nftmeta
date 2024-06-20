package nftmeta

import (
	"sourcecode.social/reiver/go-opt"
)

// MetaData represents an NFT metadata.
type MetaData struct {
	animationURL    opt.Optional[string]
	backgroundColor opt.Optional[string]
	description     opt.Optional[string]
	externalLink    opt.Optional[string]
	image           opt.Optional[string]
	imageData       opt.Optional[string]
	name            opt.Optional[string]
	youtubeURL      opt.Optional[string]
	attributes    []Attribute
}

func (receiver MetaData) MarshalJSON() ([]byte, error) {

	var after bool

	var buffer [512]byte
	var p []byte = buffer[0:0]

	p = append(p, '{')

	{
		const name string =         "animation_url"
		value, something := receiver.animationURL.Get()
		if something {
			after = true

			var err error
			p, err = appendJSONNameValue(p, name, value )

			if nil != err {
				return nil, err
			}
		}
	}

	{
		const name string =         "background_color"
		value, something := receiver.backgroundColor.Get()
		if something {
			if after {
				p = append(p, ',')
			}

			after = true

			var err error
			p, err = appendJSONNameValue(p, name, value )

			if nil != err {
				return nil, err
			}
		}

	}

	{
		const name string =         "description"
		value, something := receiver.description.Get()
		if something {
			if after {
				p = append(p, ',')
			}

			after = true

			var err error
			p, err = appendJSONNameValue(p, name, value )

			if nil != err {
				return nil, err
			}
		}

	}

	{
		const name string =         "external_link"
		value, something := receiver.externalLink.Get()
		if something {
			if after {
				p = append(p, ',')
			}

			after = true

			var err error
			p, err = appendJSONNameValue(p, name, value )

			if nil != err {
				return nil, err
			}
		}

	}

	{
		const name string =         "image"
		value, something := receiver.image.Get()
		if something {
			if after {
				p = append(p, ',')
			}

			after = true

			var err error
			p, err = appendJSONNameValue(p, name, value )

			if nil != err {
				return nil, err
			}
		}

	}

	{
		const name string =         "image_data"
		value, something := receiver.imageData.Get()
		if something {
			if after {
				p = append(p, ',')
			}

			after = true

			var err error
			p, err = appendJSONNameValue(p, name, value )

			if nil != err {
				return nil, err
			}
		}

	}

	{
		const name string =         "name"
		value, something := receiver.name.Get()
		if something {
			if after {
				p = append(p, ',')
			}

			after = true

			var err error
			p, err = appendJSONNameValue(p, name, value )

			if nil != err {
				return nil, err
			}
		}

	}

	{
		const name string =         "youtube_url"
		value, something := receiver.youtubeURL.Get()
		if something {
			if after {
				p = append(p, ',')
			}

			after = true

			var err error
			p, err = appendJSONNameValue(p, name, value )

			if nil != err {
				return nil, err
			}
		}

	}

	if 0 < len(receiver.attributes) {
		p = append(p , `"attributes":[`...)
		for index, attribute := range receiver.attributes {
			if 0 < index {
				p = append(p, ',')
			}

			bytes, err := attribute.MarshalJSON()
			if nil != err {
				return nil, err
			}

			p = append(p, bytes...)
		}
		p = append(p, ']')
	}

	p = append(p, '}')

	return p, nil
}

func (receiver *MetaData) SetAnimationURL(value string) {
	receiver.animationURL = opt.Something(value)
}

func (receiver *MetaData) SetBackgroundColor(value string) {
	receiver.backgroundColor = opt.Something(value)
}

func (receiver *MetaData) SetDescription(value string) {
	receiver.description = opt.Something(value)
}

func (receiver *MetaData) SetExternalLink(value string) {
	receiver.externalLink = opt.Something(value)
}

func (receiver *MetaData) SetImage(value string) {
	receiver.image = opt.Something(value)
}

func (receiver *MetaData) SetImageData(value string) {
	receiver.imageData = opt.Something(value)
}

func (receiver *MetaData) SetName(value string) {
	receiver.name = opt.Something(value)
}

func (receiver *MetaData) SetYouTubeURL(value string) {
	receiver.youtubeURL = opt.Something(value)
}

func (receiver *MetaData) AppendAttribute(attribute Attribute) {
	receiver.attributes = append(receiver.attributes, attribute)
}
