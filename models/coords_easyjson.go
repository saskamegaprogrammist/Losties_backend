// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson7d4fb8f6DecodeGithubComSaskamegaprogrammistLostiesBackendModels(in *jlexer.Lexer, out *CoordsAll) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(CoordsAll, 0, 2)
			} else {
				*out = CoordsAll{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 Coords
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson7d4fb8f6EncodeGithubComSaskamegaprogrammistLostiesBackendModels(out *jwriter.Writer, in CoordsAll) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v CoordsAll) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7d4fb8f6EncodeGithubComSaskamegaprogrammistLostiesBackendModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CoordsAll) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7d4fb8f6EncodeGithubComSaskamegaprogrammistLostiesBackendModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CoordsAll) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7d4fb8f6DecodeGithubComSaskamegaprogrammistLostiesBackendModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CoordsAll) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7d4fb8f6DecodeGithubComSaskamegaprogrammistLostiesBackendModels(l, v)
}
func easyjson7d4fb8f6DecodeGithubComSaskamegaprogrammistLostiesBackendModels1(in *jlexer.Lexer, out *Coords) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.Id = int(in.Int())
		case "adid":
			out.AdId = int(in.Int())
		case "x":
			out.X = float64(in.Float64())
		case "y":
			out.Y = float64(in.Float64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson7d4fb8f6EncodeGithubComSaskamegaprogrammistLostiesBackendModels1(out *jwriter.Writer, in Coords) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Id))
	}
	{
		const prefix string = ",\"adid\":"
		out.RawString(prefix)
		out.Int(int(in.AdId))
	}
	{
		const prefix string = ",\"x\":"
		out.RawString(prefix)
		out.Float64(float64(in.X))
	}
	{
		const prefix string = ",\"y\":"
		out.RawString(prefix)
		out.Float64(float64(in.Y))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Coords) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson7d4fb8f6EncodeGithubComSaskamegaprogrammistLostiesBackendModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Coords) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson7d4fb8f6EncodeGithubComSaskamegaprogrammistLostiesBackendModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Coords) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson7d4fb8f6DecodeGithubComSaskamegaprogrammistLostiesBackendModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Coords) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson7d4fb8f6DecodeGithubComSaskamegaprogrammistLostiesBackendModels1(l, v)
}