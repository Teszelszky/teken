// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: handtekening.pb.go
// DO NOT EDIT!

package data

import (
	"bytes"
	"encoding/base64"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
	"reflect"
)

const (
	ffj_t_Handtekeningbase = iota
	ffj_t_Handtekeningno_such_key

	ffj_t_Handtekening_Voornaam

	ffj_t_Handtekening_Tussenvoegsel

	ffj_t_Handtekening_Achternaam

	ffj_t_Handtekening_Geboortedatum

	ffj_t_Handtekening_Geboorteplaats

	ffj_t_Handtekening_Straat

	ffj_t_Handtekening_Huisnummer

	ffj_t_Handtekening_Postcode

	ffj_t_Handtekening_Woonplaats

	ffj_t_Handtekening_Handtekening

	ffj_t_Handtekening_Email

	ffj_t_Handtekening_CaptchaResponse
)

var ffj_key_Handtekening_Voornaam = []byte("voornaam")

var ffj_key_Handtekening_Tussenvoegsel = []byte("tussenvoegsel")

var ffj_key_Handtekening_Achternaam = []byte("achternaam")

var ffj_key_Handtekening_Geboortedatum = []byte("geboortedatum")

var ffj_key_Handtekening_Geboorteplaats = []byte("geboorteplaats")

var ffj_key_Handtekening_Straat = []byte("straat")

var ffj_key_Handtekening_Huisnummer = []byte("huisnummer")

var ffj_key_Handtekening_Postcode = []byte("postcode")

var ffj_key_Handtekening_Woonplaats = []byte("woonplaats")

var ffj_key_Handtekening_Handtekening = []byte("handtekening")

var ffj_key_Handtekening_Email = []byte("email")

var ffj_key_Handtekening_CaptchaResponse = []byte("captchaResponse")

func (uj *Handtekening) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *Handtekening) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_Handtekeningbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffj_t_Handtekeningno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'a':

					if bytes.Equal(ffj_key_Handtekening_Achternaam, kn) {
						currentKey = ffj_t_Handtekening_Achternaam
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'c':

					if bytes.Equal(ffj_key_Handtekening_CaptchaResponse, kn) {
						currentKey = ffj_t_Handtekening_CaptchaResponse
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'e':

					if bytes.Equal(ffj_key_Handtekening_Email, kn) {
						currentKey = ffj_t_Handtekening_Email
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'g':

					if bytes.Equal(ffj_key_Handtekening_Geboortedatum, kn) {
						currentKey = ffj_t_Handtekening_Geboortedatum
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Handtekening_Geboorteplaats, kn) {
						currentKey = ffj_t_Handtekening_Geboorteplaats
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'h':

					if bytes.Equal(ffj_key_Handtekening_Huisnummer, kn) {
						currentKey = ffj_t_Handtekening_Huisnummer
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Handtekening_Handtekening, kn) {
						currentKey = ffj_t_Handtekening_Handtekening
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'p':

					if bytes.Equal(ffj_key_Handtekening_Postcode, kn) {
						currentKey = ffj_t_Handtekening_Postcode
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 's':

					if bytes.Equal(ffj_key_Handtekening_Straat, kn) {
						currentKey = ffj_t_Handtekening_Straat
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 't':

					if bytes.Equal(ffj_key_Handtekening_Tussenvoegsel, kn) {
						currentKey = ffj_t_Handtekening_Tussenvoegsel
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'v':

					if bytes.Equal(ffj_key_Handtekening_Voornaam, kn) {
						currentKey = ffj_t_Handtekening_Voornaam
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'w':

					if bytes.Equal(ffj_key_Handtekening_Woonplaats, kn) {
						currentKey = ffj_t_Handtekening_Woonplaats
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffj_key_Handtekening_CaptchaResponse, kn) {
					currentKey = ffj_t_Handtekening_CaptchaResponse
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Handtekening_Email, kn) {
					currentKey = ffj_t_Handtekening_Email
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Handtekening_Handtekening, kn) {
					currentKey = ffj_t_Handtekening_Handtekening
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Handtekening_Woonplaats, kn) {
					currentKey = ffj_t_Handtekening_Woonplaats
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Handtekening_Postcode, kn) {
					currentKey = ffj_t_Handtekening_Postcode
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Handtekening_Huisnummer, kn) {
					currentKey = ffj_t_Handtekening_Huisnummer
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Handtekening_Straat, kn) {
					currentKey = ffj_t_Handtekening_Straat
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Handtekening_Geboorteplaats, kn) {
					currentKey = ffj_t_Handtekening_Geboorteplaats
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Handtekening_Geboortedatum, kn) {
					currentKey = ffj_t_Handtekening_Geboortedatum
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Handtekening_Achternaam, kn) {
					currentKey = ffj_t_Handtekening_Achternaam
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Handtekening_Tussenvoegsel, kn) {
					currentKey = ffj_t_Handtekening_Tussenvoegsel
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Handtekening_Voornaam, kn) {
					currentKey = ffj_t_Handtekening_Voornaam
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_Handtekeningno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffj_t_Handtekening_Voornaam:
					goto handle_Voornaam

				case ffj_t_Handtekening_Tussenvoegsel:
					goto handle_Tussenvoegsel

				case ffj_t_Handtekening_Achternaam:
					goto handle_Achternaam

				case ffj_t_Handtekening_Geboortedatum:
					goto handle_Geboortedatum

				case ffj_t_Handtekening_Geboorteplaats:
					goto handle_Geboorteplaats

				case ffj_t_Handtekening_Straat:
					goto handle_Straat

				case ffj_t_Handtekening_Huisnummer:
					goto handle_Huisnummer

				case ffj_t_Handtekening_Postcode:
					goto handle_Postcode

				case ffj_t_Handtekening_Woonplaats:
					goto handle_Woonplaats

				case ffj_t_Handtekening_Handtekening:
					goto handle_Handtekening

				case ffj_t_Handtekening_Email:
					goto handle_Email

				case ffj_t_Handtekening_CaptchaResponse:
					goto handle_CaptchaResponse

				case ffj_t_Handtekeningno_such_key:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_Voornaam:

	/* handler: uj.Voornaam type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Voornaam = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Tussenvoegsel:

	/* handler: uj.Tussenvoegsel type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Tussenvoegsel = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Achternaam:

	/* handler: uj.Achternaam type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Achternaam = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Geboortedatum:

	/* handler: uj.Geboortedatum type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Geboortedatum = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Geboorteplaats:

	/* handler: uj.Geboorteplaats type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Geboorteplaats = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Straat:

	/* handler: uj.Straat type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Straat = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Huisnummer:

	/* handler: uj.Huisnummer type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Huisnummer = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Postcode:

	/* handler: uj.Postcode type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Postcode = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Woonplaats:

	/* handler: uj.Woonplaats type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Woonplaats = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Handtekening:

	/* handler: uj.Handtekening type=[]uint8 kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			uj.Handtekening = nil
		} else {
			b := make([]byte, base64.StdEncoding.DecodedLen(fs.Output.Len()))
			n, err := base64.StdEncoding.Decode(b, fs.Output.Bytes())
			if err != nil {
				return fs.WrapErr(err)
			}

			v := reflect.ValueOf(&uj.Handtekening).Elem()
			v.SetBytes(b[0:n])

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Email:

	/* handler: uj.Email type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.Email = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_CaptchaResponse:

	/* handler: uj.CaptchaResponse type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			uj.CaptchaResponse = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:
	return nil
}
