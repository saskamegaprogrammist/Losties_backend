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

func easyjson14a1085DecodeGithubComSaskamegaprogrammistLostiesBackendModels(in *jlexer.Lexer, out *Pet) {
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
		case "name":
			out.Name = string(in.String())
		case "animal":
			out.Animal = string(in.String())
		case "breed":
			out.Breed = string(in.String())
		case "color":
			out.Color = string(in.String())
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
func easyjson14a1085EncodeGithubComSaskamegaprogrammistLostiesBackendModels(out *jwriter.Writer, in Pet) {
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
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"animal\":"
		out.RawString(prefix)
		out.String(string(in.Animal))
	}
	{
		const prefix string = ",\"breed\":"
		out.RawString(prefix)
		out.String(string(in.Breed))
	}
	{
		const prefix string = ",\"color\":"
		out.RawString(prefix)
		out.String(string(in.Color))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Pet) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson14a1085EncodeGithubComSaskamegaprogrammistLostiesBackendModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Pet) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson14a1085EncodeGithubComSaskamegaprogrammistLostiesBackendModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Pet) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson14a1085DecodeGithubComSaskamegaprogrammistLostiesBackendModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Pet) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson14a1085DecodeGithubComSaskamegaprogrammistLostiesBackendModels(l, v)
}
