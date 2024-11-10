// Code generated by ogen, DO NOT EDIT.

package api

import (
	"math/bits"
	"strconv"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/validate"
)

// Encode implements json.Marshaler.
func (s *CategoriesResp) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *CategoriesResp) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("data")
		e.ArrStart()
		for _, elem := range s.Data {
			elem.Encode(e)
		}
		e.ArrEnd()
	}
	{
		e.FieldStart("page")
		e.Int(s.Page)
	}
	{
		e.FieldStart("pages")
		e.Int(s.Pages)
	}
	{
		e.FieldStart("per_page")
		e.Int(s.PerPage)
	}
}

var jsonFieldsNameOfCategoriesResp = [4]string{
	0: "data",
	1: "page",
	2: "pages",
	3: "per_page",
}

// Decode decodes CategoriesResp from json.
func (s *CategoriesResp) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode CategoriesResp to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "data":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				s.Data = make([]Category, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem Category
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.Data = append(s.Data, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"data\"")
			}
		case "page":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Int()
				s.Page = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"page\"")
			}
		case "pages":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Int()
				s.Pages = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"pages\"")
			}
		case "per_page":
			requiredBitSet[0] |= 1 << 3
			if err := func() error {
				v, err := d.Int()
				s.PerPage = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"per_page\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode CategoriesResp")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00001111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfCategoriesResp) {
					name = jsonFieldsNameOfCategoriesResp[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *CategoriesResp) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *CategoriesResp) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Category) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Category) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("id")
		e.Int(s.ID)
	}
	{
		e.FieldStart("title")
		e.Str(s.Title)
	}
	{
		e.FieldStart("public")
		e.Bool(s.Public)
	}
}

var jsonFieldsNameOfCategory = [3]string{
	0: "id",
	1: "title",
	2: "public",
}

// Decode decodes Category from json.
func (s *Category) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Category to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "id":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Int()
				s.ID = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"id\"")
			}
		case "title":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Str()
				s.Title = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		case "public":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				v, err := d.Bool()
				s.Public = bool(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"public\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Category")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfCategory) {
					name = jsonFieldsNameOfCategory[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Category) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Category) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *CategoryPost) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *CategoryPost) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("title")
		e.Str(s.Title)
	}
	{
		e.FieldStart("public")
		e.Bool(s.Public)
	}
}

var jsonFieldsNameOfCategoryPost = [2]string{
	0: "title",
	1: "public",
}

// Decode decodes CategoryPost from json.
func (s *CategoryPost) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode CategoryPost to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "title":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.Title = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		case "public":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Bool()
				s.Public = bool(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"public\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode CategoryPost")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfCategoryPost) {
					name = jsonFieldsNameOfCategoryPost[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *CategoryPost) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *CategoryPost) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *CategoryPut) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *CategoryPut) encodeFields(e *jx.Encoder) {
	{
		if s.Title.Set {
			e.FieldStart("title")
			s.Title.Encode(e)
		}
	}
	{
		if s.Public.Set {
			e.FieldStart("public")
			s.Public.Encode(e)
		}
	}
}

var jsonFieldsNameOfCategoryPut = [2]string{
	0: "title",
	1: "public",
}

// Decode decodes CategoryPut from json.
func (s *CategoryPut) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode CategoryPut to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "title":
			if err := func() error {
				s.Title.Reset()
				if err := s.Title.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"title\"")
			}
		case "public":
			if err := func() error {
				s.Public.Reset()
				if err := s.Public.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"public\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode CategoryPut")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *CategoryPut) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *CategoryPut) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Error) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Error) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("message")
		e.Str(s.Message)
	}
	{
		e.FieldStart("details")
		s.Details.Encode(e)
	}
}

var jsonFieldsNameOfError = [2]string{
	0: "message",
	1: "details",
}

// Decode decodes Error from json.
func (s *Error) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Error to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "message":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.Message = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"message\"")
			}
		case "details":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				if err := s.Details.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"details\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Error")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfError) {
					name = jsonFieldsNameOfError[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Error) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Error) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s ErrorDetails) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields implements json.Marshaler.
func (s ErrorDetails) encodeFields(e *jx.Encoder) {
	for k, elem := range s {
		e.FieldStart(k)

		if len(elem) != 0 {
			e.Raw(elem)
		}
	}
}

// Decode decodes ErrorDetails from json.
func (s *ErrorDetails) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ErrorDetails to nil")
	}
	m := s.init()
	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		var elem jx.Raw
		if err := func() error {
			v, err := d.RawAppend(nil)
			elem = jx.Raw(v)
			if err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return errors.Wrapf(err, "decode field %q", k)
		}
		m[string(k)] = elem
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ErrorDetails")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s ErrorDetails) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ErrorDetails) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes bool as json.
func (o OptBool) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Bool(bool(o.Value))
}

// Decode decodes bool from json.
func (o *OptBool) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptBool to nil")
	}
	o.Set = true
	v, err := d.Bool()
	if err != nil {
		return err
	}
	o.Value = bool(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptBool) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptBool) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes string as json.
func (o OptString) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes string from json.
func (o *OptString) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptString to nil")
	}
	o.Set = true
	v, err := d.Str()
	if err != nil {
		return err
	}
	o.Value = string(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptString) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptString) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes V1AdminCategoriesGetBadRequest as json.
func (s *V1AdminCategoriesGetBadRequest) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes V1AdminCategoriesGetBadRequest from json.
func (s *V1AdminCategoriesGetBadRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode V1AdminCategoriesGetBadRequest to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = V1AdminCategoriesGetBadRequest(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *V1AdminCategoriesGetBadRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *V1AdminCategoriesGetBadRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes V1AdminCategoriesGetNotFound as json.
func (s *V1AdminCategoriesGetNotFound) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes V1AdminCategoriesGetNotFound from json.
func (s *V1AdminCategoriesGetNotFound) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode V1AdminCategoriesGetNotFound to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = V1AdminCategoriesGetNotFound(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *V1AdminCategoriesGetNotFound) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *V1AdminCategoriesGetNotFound) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes V1AdminCategoriesGetUnauthorized as json.
func (s *V1AdminCategoriesGetUnauthorized) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes V1AdminCategoriesGetUnauthorized from json.
func (s *V1AdminCategoriesGetUnauthorized) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode V1AdminCategoriesGetUnauthorized to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = V1AdminCategoriesGetUnauthorized(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *V1AdminCategoriesGetUnauthorized) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *V1AdminCategoriesGetUnauthorized) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes V1AdminCategoriesIDDeleteConflict as json.
func (s *V1AdminCategoriesIDDeleteConflict) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes V1AdminCategoriesIDDeleteConflict from json.
func (s *V1AdminCategoriesIDDeleteConflict) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode V1AdminCategoriesIDDeleteConflict to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = V1AdminCategoriesIDDeleteConflict(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *V1AdminCategoriesIDDeleteConflict) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *V1AdminCategoriesIDDeleteConflict) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes V1AdminCategoriesIDDeleteNotFound as json.
func (s *V1AdminCategoriesIDDeleteNotFound) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes V1AdminCategoriesIDDeleteNotFound from json.
func (s *V1AdminCategoriesIDDeleteNotFound) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode V1AdminCategoriesIDDeleteNotFound to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = V1AdminCategoriesIDDeleteNotFound(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *V1AdminCategoriesIDDeleteNotFound) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *V1AdminCategoriesIDDeleteNotFound) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes V1AdminCategoriesIDDeleteUnauthorized as json.
func (s *V1AdminCategoriesIDDeleteUnauthorized) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes V1AdminCategoriesIDDeleteUnauthorized from json.
func (s *V1AdminCategoriesIDDeleteUnauthorized) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode V1AdminCategoriesIDDeleteUnauthorized to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = V1AdminCategoriesIDDeleteUnauthorized(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *V1AdminCategoriesIDDeleteUnauthorized) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *V1AdminCategoriesIDDeleteUnauthorized) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes V1AdminCategoriesIDPutBadRequest as json.
func (s *V1AdminCategoriesIDPutBadRequest) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes V1AdminCategoriesIDPutBadRequest from json.
func (s *V1AdminCategoriesIDPutBadRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode V1AdminCategoriesIDPutBadRequest to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = V1AdminCategoriesIDPutBadRequest(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *V1AdminCategoriesIDPutBadRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *V1AdminCategoriesIDPutBadRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes V1AdminCategoriesIDPutConflict as json.
func (s *V1AdminCategoriesIDPutConflict) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes V1AdminCategoriesIDPutConflict from json.
func (s *V1AdminCategoriesIDPutConflict) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode V1AdminCategoriesIDPutConflict to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = V1AdminCategoriesIDPutConflict(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *V1AdminCategoriesIDPutConflict) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *V1AdminCategoriesIDPutConflict) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes V1AdminCategoriesIDPutNotFound as json.
func (s *V1AdminCategoriesIDPutNotFound) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes V1AdminCategoriesIDPutNotFound from json.
func (s *V1AdminCategoriesIDPutNotFound) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode V1AdminCategoriesIDPutNotFound to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = V1AdminCategoriesIDPutNotFound(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *V1AdminCategoriesIDPutNotFound) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *V1AdminCategoriesIDPutNotFound) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes V1AdminCategoriesIDPutUnauthorized as json.
func (s *V1AdminCategoriesIDPutUnauthorized) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes V1AdminCategoriesIDPutUnauthorized from json.
func (s *V1AdminCategoriesIDPutUnauthorized) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode V1AdminCategoriesIDPutUnauthorized to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = V1AdminCategoriesIDPutUnauthorized(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *V1AdminCategoriesIDPutUnauthorized) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *V1AdminCategoriesIDPutUnauthorized) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes V1AdminCategoriesPostBadRequest as json.
func (s *V1AdminCategoriesPostBadRequest) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes V1AdminCategoriesPostBadRequest from json.
func (s *V1AdminCategoriesPostBadRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode V1AdminCategoriesPostBadRequest to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = V1AdminCategoriesPostBadRequest(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *V1AdminCategoriesPostBadRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *V1AdminCategoriesPostBadRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes V1AdminCategoriesPostConflict as json.
func (s *V1AdminCategoriesPostConflict) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes V1AdminCategoriesPostConflict from json.
func (s *V1AdminCategoriesPostConflict) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode V1AdminCategoriesPostConflict to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = V1AdminCategoriesPostConflict(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *V1AdminCategoriesPostConflict) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *V1AdminCategoriesPostConflict) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes V1AdminCategoriesPostUnauthorized as json.
func (s *V1AdminCategoriesPostUnauthorized) Encode(e *jx.Encoder) {
	unwrapped := (*Error)(s)

	unwrapped.Encode(e)
}

// Decode decodes V1AdminCategoriesPostUnauthorized from json.
func (s *V1AdminCategoriesPostUnauthorized) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode V1AdminCategoriesPostUnauthorized to nil")
	}
	var unwrapped Error
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = V1AdminCategoriesPostUnauthorized(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *V1AdminCategoriesPostUnauthorized) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *V1AdminCategoriesPostUnauthorized) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}
