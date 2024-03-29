// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package gate

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

func easyjson6891ccbbDecodeQntRobotMarketGate(in *jlexer.Lexer, out *Depth) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "time":
			out.Time = int(in.Int())
		case "time_ms":
			out.TimeMs = int64(in.Int64())
		case "channel":
			out.Channel = string(in.String())
		case "event":
			out.Event = string(in.String())
		case "result":
			easyjson6891ccbbDecode(in, &out.Result)
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
func easyjson6891ccbbEncodeQntRobotMarketGate(out *jwriter.Writer, in Depth) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"time\":"
		out.RawString(prefix[1:])
		out.Int(int(in.Time))
	}
	{
		const prefix string = ",\"time_ms\":"
		out.RawString(prefix)
		out.Int64(int64(in.TimeMs))
	}
	{
		const prefix string = ",\"channel\":"
		out.RawString(prefix)
		out.String(string(in.Channel))
	}
	{
		const prefix string = ",\"event\":"
		out.RawString(prefix)
		out.String(string(in.Event))
	}
	{
		const prefix string = ",\"result\":"
		out.RawString(prefix)
		easyjson6891ccbbEncode(out, in.Result)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Depth) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6891ccbbEncodeQntRobotMarketGate(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Depth) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6891ccbbEncodeQntRobotMarketGate(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Depth) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6891ccbbDecodeQntRobotMarketGate(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Depth) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6891ccbbDecodeQntRobotMarketGate(l, v)
}
func easyjson6891ccbbDecode(in *jlexer.Lexer, out *struct {
	T            int64      `json:"t"`
	LastUpdateId int        `json:"lastUpdateId"`
	S            string     `json:"s"`
	Bids         [][]string `json:"bids"`
	Asks         [][]string `json:"asks"`
}) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "t":
			out.T = int64(in.Int64())
		case "lastUpdateId":
			out.LastUpdateId = int(in.Int())
		case "s":
			out.S = string(in.String())
		case "bids":
			if in.IsNull() {
				in.Skip()
				out.Bids = nil
			} else {
				in.Delim('[')
				if out.Bids == nil {
					if !in.IsDelim(']') {
						out.Bids = make([][]string, 0, 2)
					} else {
						out.Bids = [][]string{}
					}
				} else {
					out.Bids = (out.Bids)[:0]
				}
				for !in.IsDelim(']') {
					var v1 []string
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						in.Delim('[')
						if v1 == nil {
							if !in.IsDelim(']') {
								v1 = make([]string, 0, 4)
							} else {
								v1 = []string{}
							}
						} else {
							v1 = (v1)[:0]
						}
						for !in.IsDelim(']') {
							var v2 string
							v2 = string(in.String())
							v1 = append(v1, v2)
							in.WantComma()
						}
						in.Delim(']')
					}
					out.Bids = append(out.Bids, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "asks":
			if in.IsNull() {
				in.Skip()
				out.Asks = nil
			} else {
				in.Delim('[')
				if out.Asks == nil {
					if !in.IsDelim(']') {
						out.Asks = make([][]string, 0, 2)
					} else {
						out.Asks = [][]string{}
					}
				} else {
					out.Asks = (out.Asks)[:0]
				}
				for !in.IsDelim(']') {
					var v3 []string
					if in.IsNull() {
						in.Skip()
						v3 = nil
					} else {
						in.Delim('[')
						if v3 == nil {
							if !in.IsDelim(']') {
								v3 = make([]string, 0, 4)
							} else {
								v3 = []string{}
							}
						} else {
							v3 = (v3)[:0]
						}
						for !in.IsDelim(']') {
							var v4 string
							v4 = string(in.String())
							v3 = append(v3, v4)
							in.WantComma()
						}
						in.Delim(']')
					}
					out.Asks = append(out.Asks, v3)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjson6891ccbbEncode(out *jwriter.Writer, in struct {
	T            int64      `json:"t"`
	LastUpdateId int        `json:"lastUpdateId"`
	S            string     `json:"s"`
	Bids         [][]string `json:"bids"`
	Asks         [][]string `json:"asks"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"t\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.T))
	}
	{
		const prefix string = ",\"lastUpdateId\":"
		out.RawString(prefix)
		out.Int(int(in.LastUpdateId))
	}
	{
		const prefix string = ",\"s\":"
		out.RawString(prefix)
		out.String(string(in.S))
	}
	{
		const prefix string = ",\"bids\":"
		out.RawString(prefix)
		if in.Bids == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Bids {
				if v5 > 0 {
					out.RawByte(',')
				}
				if v6 == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
					out.RawString("null")
				} else {
					out.RawByte('[')
					for v7, v8 := range v6 {
						if v7 > 0 {
							out.RawByte(',')
						}
						out.String(string(v8))
					}
					out.RawByte(']')
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"asks\":"
		out.RawString(prefix)
		if in.Asks == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v9, v10 := range in.Asks {
				if v9 > 0 {
					out.RawByte(',')
				}
				if v10 == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
					out.RawString("null")
				} else {
					out.RawByte('[')
					for v11, v12 := range v10 {
						if v11 > 0 {
							out.RawByte(',')
						}
						out.String(string(v12))
					}
					out.RawByte(']')
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
