// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: outputtypes.go
// DO NOT EDIT!

package server

import (
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (mj *StatsOutput) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *StatsOutput) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"total":`)
	fflib.FormatBits2(buf, uint64(mj.Total), 10, false)
	buf.WriteString(`,"verified":`)
	fflib.FormatBits2(buf, uint64(mj.Verified), 10, false)
	buf.WriteByte('}')
	return nil
}

func (mj *SubmitOutput) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *SubmitOutput) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	var err error
	var obj []byte
	_ = obj
	_ = err
	if mj.Success {
		buf.WriteString(`{ "success":true`)
	} else {
		buf.WriteString(`{ "success":false`)
	}
	buf.WriteByte(',')
	if len(mj.Error) != 0 {
		buf.WriteString(`"error":`)
		fflib.WriteJsonString(buf, string(mj.Error))
		buf.WriteByte(',')
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}
