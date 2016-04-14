// (C) Philip Schlump, 2014-2015.

package templatestrings

import (
	"fmt"
	"os"
	"testing"

	"github.com/pschlump/json" //	"encoding/json"
)

//type TestCase struct {
//	format string
//	expected string
//	value float64
//}
//
//var testCases = []TestCase{
//	{ "##,##0.00",     "  -123.46",       -123.456 },		// round to .47???
//	{ "##,##0.00",     "   123.46",        123.456 },		// 1
//	{ "##,##0.00",     "##,##0.00",     123123.456 },		// 2
//
//	// Escape
//	// &TestCase{"YYYY/MM/DD hh:mm:ss", "2014/01/10 11:31:32"},
//	// &TestCase{"YYYY-MM-DD hh:mm:ss", "2014-01-10 11:31:32"},
//	// In a string
//	// &TestCase{"/aaaa/YYYY/mm/bbbb", "/aaaa/2014/31/bbbb"},
//
//	// No Format - get rid of value
//	{ "",              "",                 123.456 },		// 3
//	{ ".",             ".",                123.456 },		// 4
//	{ "0",             "3",                  3.456 },		// 5
//	{ "#",             "3",                  3.456 },		// 6
//
//	{ "##,##0.00",     "     0.00",          0.0   },		// 7
//	{ "##,##0.00",     "     0.00",         -0.0   },		// 1
//}

func TestPictureFormats(t *testing.T) {
	//if false {
	//	fmt.Printf ( "keep compiler happy when we are not using fmt.\n" )
	//}
	rv := CenterStr(10, "Adm")
	ex := "   Adm    "
	if rv != "   Adm    " {
		t.Fatalf("Error CenterStr(10,\"Adm\") results=[%s] expected=[%s]", rv, ex)
	}
	//for i, v := range testCases {
	//	// fmt.Printf ( "Running %d\n", i )
	//	result := Format(v.format, v.value)
	//	if result != v.expected {
	//		t.Fatalf("Error for %f at [%d] in table: format=[%s]: results=[%s] expected=[%s]", v.value, i, v.format, result, v.expected)
	//	}
	//}
}

func TestHomeDir(t *testing.T) {
	s := HomeDir()
	h := os.Getenv("HOME")
	if string(os.PathSeparator) == `\` {
		fmt.Printf("Windows: s= ->%s<- h= ->%s<-\n", s, h)
		if !(s[2:3] == string(os.PathSeparator)) {
			t.Fatalf("Error did not get back a home directory. Got %s", s)
		}
	} else {
		if !(s[0:1] == string(os.PathSeparator)) {
			t.Fatalf("Error did not get back a home directory. Got %s", s)
		}
	}
	// fmt.Printf ( "h=%s\n", h )
	if h != "" {
		if !(s == h) {
			t.Fatalf("(Unix/Linux/Mac/Specific): Error did not get back a home directory. Got %s, expecting %s", s, h)
		}
	} else {
		fmt.Printf("Warning: HOME environment variable not set\n")
	}
}

// func SplitOnWords ( s string ) ( record []string ) {
func TestSplitOnWords(t *testing.T) {

	rv := SplitOnWords("a b c")
	if len(rv) != 3 {
		t.Fatalf("Wrong number of items returned")
	}
	if rv[0] != "a" {
		t.Fatalf("Wrong [0] items returned")
	}

	rv = SplitOnWords(`a "b c d" e`)
	if len(rv) != 3 {
		t.Fatalf("Wrong number of items returned")
	}
	if rv[1] != "b c d" {
		t.Fatalf("Wrong [0] items returned")
	}

}

func TestNvl(t *testing.T) {
	s := Nvl("x", "")
	if s != "x" {
		t.Fatalf("Nvl failed to see empty string")
	}
	s = Nvl("x", "y")
	if s != "y" {
		t.Fatalf("Nvl failed to see non empty string")
	}
}

// func PicFloat ( format string, flt float64 ) ( r string ) {
func TestPicFloat(t *testing.T) {
	s := PicFloat("##,##0.00", 123.321)
	if s != "   123.32" {
		t.Fatalf("PicFloat failed to format string, expected \"   123.32\", got ->%s<-", s)
	}
}

//func PadOnRight ( n int, s string ) ( r string ) {
func TestPadOnRight(t *testing.T) {
	s := PadOnRight(5, "abc")
	ex := "abc  "
	if s != ex {
		t.Fatalf("PadOnRight failed to get expected ->%s<-, got ->%s<-", ex, s)
	}
	s = PadOnRight(5, 12)
	ex = "12   "
	if s != ex {
		t.Fatalf("PadOnRight failed to get expected ->%s<-, got ->%s<-", ex, s)
	}
}

func TestPadOnLeft(t *testing.T) {
	s := PadOnLeft(5, "abc")
	ex := "  abc"
	if s != ex {
		t.Fatalf("PadOnRight failed to get expected ->%s<-, got ->%s<-", ex, s)
	}
	s = PadOnLeft(5, 12)
	ex = "   12"
	if s != ex {
		t.Fatalf("PadOnRight failed to get expected ->%s<-, got ->%s<-", ex, s)
	}
}

// func FixBindParams ( qry string, data ...interface{} ) ( qryFixed string, retData []interface{}, err error ) {
func SVar(v interface{}) string {
	s, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprintf("Error:%s", err)
	} else {
		return string(s)
	}
}

type TBindParams struct {
	q  string // the query
	D  int    // the data [ index to correct data in func ]
	lD int    // length of returnd data
	hE bool   // error found
	rv string
	ds string
}

var testCasesTBindParams = []TBindParams{
	{"select * from test1", 0, 0, false, "select * from test1", "null"},
	{`select * from "test1"`, 0, 0, false, "select * from [test1]", "null"},
	{`select * from "test1" where "n" = $1`, 1, 1, false, "select * from [test1] where [n] = ?", "[12]"},
	{`select * from "test1" where "n" = $1 and M = $2`, 2, 2, false, "select * from [test1] where [n] = ? and M = ?", "[12,\"a\"]"},
	{`select * from "test1" where "n" = $2 and M = $1`, 2, 2, false, "select * from [test1] where [n] = ? and M = ?", "[\"a\",12]"},
	{`select * from "test1" where "n" = $1 or "n" = $1 or "n" = $3`,
		3, 3, false, "select * from [test1] where [n] = ? or [n] = ? or [n] = ?", "[12,12,98.7]"},
	{`select * from "test1" where "n" = $3 or "n" = $2 or "n" = $1`,
		3, 3, false, "select * from [test1] where [n] = ? or [n] = ? or [n] = ?", "[98.7,\"a\",12]"},
	// test with $ in quoted string ' quote, also that you can reduce # of bind values returned.
	{`select * from "test1" where "n" = '$3' or "n" = $2 or "n" = $1`,
		3, 2, false, "select * from [test1] where [n] = '$3' or [n] = ? or [n] = ?", "[\"a\",12]"},
	// test with $ in quoted string " quote, weird column name "$2"
	{`select * from "test1" where "n" = $3 or "n" = "$2" or "n" = $1`,
		3, 2, false, "select * from [test1] where [n] = ? or [n] = [$2] or [n] = ?", "[98.7,12]"},
	// test with " in ' string
	{`select * from "test1" where "n" = 'ab"cd' or "n" = $2 or "n" = $1`,
		3, 2, false, "select * from [test1] where [n] = 'ab\"cd' or [n] = ? or [n] = ?", "[\"a\",12]"},
	// test (2) with " in ' string
	{`select * from "test1" where "n" = 'ab"' or "n" = $2 or "n" = $1`,
		3, 2, false, "select * from [test1] where [n] = 'ab\"' or [n] = ? or [n] = ?", "[\"a\",12]"},
	// test (3) with " in ' string
	{`select * from "test1" where "n" = '"' or "n" = $2 or "n" = $1`,
		3, 2, false, "select * from [test1] where [n] = '\"' or [n] = ? or [n] = ?", "[\"a\",12]"},
	// test with ' in " string
	{`select * from "test1" where "n" = $3 or "v'" = $2 or "n" = $1`,
		3, 3, false, "select * from [test1] where [n] = ? or [v'] = ? or [n] = ?", "[98.7,\"a\",12]"},
	// [13] test with improperly terminated '
	{`select * from "test1" where "n" = $3 or "n" = $2 or "n" = 'a`,
		3, 2, false, "select * from [test1] where [n] = ? or [n] = ? or [n] = 'a", "[98.7,\"a\"]"},
	// [14] test with improperly terminated "
	{`select * from "test1" where "n" = $3 or "n" = $2 or "n`,
		3, 2, false, "select * from [test1] where [n] = ? or [n] = ? or [n", "[98.7,\"a\"]"},
	// [15] test with bad $ bind, "$" at end of string
	{`select * from "test1" where "n" = $3 or "n" = $2 or "n" = $`,
		3, 2, true, "select * from [test1] where [n] = ? or [n] = ? or [n] = ?", "[98.7,\"a\"]"},
	// test 3 input bind and 5 output bind
	{`select * from "test1" where "n" = $3 or "n" = $2 or "n" = $1 or "n" = $1 or "n" = $1`,
		3, 5, false, "select * from [test1] where [n] = ? or [n] = ? or [n] = ? or [n] = ? or [n] = ?", "[98.7,\"a\",12,12,12]"},
}

func TestFixBindParams01(t *testing.T) {
	// qf, dt, err := FixBindParams ( "select * from test1" )
	var qf string
	var dt []interface{}
	var err error
	for i, v := range testCasesTBindParams {
		_ = i
		// fmt.Printf ( "Test %d Start\n", i )
		switch v.D {
		case 0:
			qf, dt, err = FixBindParams(v.q)
		case 1:
			qf, dt, err = FixBindParams(v.q, 12)
		case 2:
			qf, dt, err = FixBindParams(v.q, 12, "a")
		case 3:
			qf, dt, err = FixBindParams(v.q, 12, "a", 98.7)
		}
		sd := SVar(dt)
		if qf != v.rv {
			t.Fatalf("Failed to return simple query")
		}
		if len(dt) != v.lD {
			t.Fatalf("Failed to retrun bind set of correct length, expected len=%d/%v, got len=%d/%v", v.lD, v.ds, len(dt), sd)
		}
		if !v.hE {
			if err != nil {
				t.Fatalf("Generated error where none is found, err=%s", err)
			}
		} else {
			if err == nil {
				t.Fatalf("Generated NO error where one should be found, at test %d", i)
			}
		}
		if sd != v.ds {
			t.Fatalf("Incorrect bind params returned, expectd=%s got=%s\n", v.ds, sd)
		}
		// fmt.Printf ( "Test %d Done\n", i )
	}
}

/* vim: set noai ts=4 sw=4: */
