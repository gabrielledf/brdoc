package brdoc

import (
	"errors"
	"regexp"
	"strconv"
)

var errorValidate = errors.New("Invalid format")
var errorDigit = errors.New("Invalid digit")

var reCpf *regexp.Regexp
var reCnpj *regexp.Regexp
var reAll0, reAll1, reAll2, reAll3, reAll4, reAll5, reAll6, reAll7, reAll8, reAll9 *regexp.Regexp
var reall0, reall1, reall2, reall3, reall4, reall5, reall6, reall7, reall8, reall9 *regexp.Regexp

var invalid []*regexp.Regexp

func init() {
	reCpf = regexp.MustCompile(`^\\d{3}\\.?\\d{3}\\.?\\d{3}-?\\d{2}$`)
	reCnpj = regexp.MustCompile(`^\\d{2}\\.?\\d{3}\\.?\\d{3}\\/?\\d{4}-?\\d{2}$`)

	//all digit equals to CPF
	reAll0 = regexp.MustCompile(`^000\.?000\.?000-?00$`)
	reAll1 = regexp.MustCompile(`^111\.?111\.?111-?11$`)
	reAll2 = regexp.MustCompile(`^222\.?222\.?222-?22$`)
	reAll3 = regexp.MustCompile(`^333\.?333\.?333-?33`)
	reAll4 = regexp.MustCompile(`^444\.?444\.?444-?44`)
	reAll5 = regexp.MustCompile(`^555\.?555\.?555-?55`)
	reAll6 = regexp.MustCompile(`^666\.?666\.?666-?66$`)
	reAll7 = regexp.MustCompile(`^777\.?777\.?777-?77$`)
	reAll8 = regexp.MustCompile(`^888\.?888\.?888-?88$`)
	reAll9 = regexp.MustCompile(`^999\.?999\.?999-?99$`)

	//all digit equals to CNPJ
	reall0 = regexp.MustCompile(`^\d{2}\.?\d{3}\.?\d{3}\/?0000-?\d{2}$`)
	reall1 = regexp.MustCompile(`^11\.?111\.?111\/?1111-?11$`)
	reall2 = regexp.MustCompile(`^22\.?222\.?222\/?2222-?22$`)
	reall3 = regexp.MustCompile(`^33\.?333\.?333\/?3333-?33$`)
	reall4 = regexp.MustCompile(`^44\.?444\.?444\/?4444-?44$`)
	reall5 = regexp.MustCompile(`^55\.?555\.?555\/?5555-?55`)
	reall6 = regexp.MustCompile(`^66\.?666\.?666\/?6666-?66$`)
	reall7 = regexp.MustCompile(`^77\.?777\.?777\/?7777-?77$`)
	reall8 = regexp.MustCompile(`^88\.?888\.?888\/?8888-?88$`)
	reall9 = regexp.MustCompile(`^99\.?999\.?999\/?9999-?99$`)

	invalid = append(invalid, reAll0, reAll1, reAll2, reAll3, reAll4, reAll5, reAll6, reAll7, reAll8, reAll9)
	invalid = append(invalid, reall0, reall1, reall2, reall3, reall4, reall5, reall6, reall7, reall8, reall9)
}

// IsCPF verifies if the string is a valid CPF document.
func IsCPF(doc string) (bool, error) {
	return check(doc, reCpf, 10)
}

// IsCNPJ verifies if the string is a valid CNPJ document.
func IsCNPJ(doc string) (bool, error) {
	return check(doc, reCnpj, 5)
}
func validateFormat(pattern *regexp.Regexp, doc string) bool {
	for _, p := range invalid {
		if p.MatchString(doc) {
			return false
		}
	}
	return !pattern.MatchString(doc)
}

func Clean(doc string) string {
	re, err := regexp.Compile("\\D")
	if err != nil {
		return ""
	}
	return re.ReplaceAllString(doc, "")
}

func calculateDigit(doc string, positions int) string {
	sum := 0

	// Sums all the digits in the document.
	// Ex.
	//   3    4    2    6    1    8    7    1    0
	// x10   x9   x8   x7   x6   x5   x4   x3   x2
	//  30 + 36 + 16 + 42 +  6 + 40 + 28 +  3 +  0 = 201
	for i := 0; i < len(doc); i++ {
		sum += int(doc[i]-'0') * positions
		positions--

		if positions < 2 {
			positions = 9
		}
	}

	sum %= 11
	if sum < 2 {
		return "0"
	}
	return strconv.FormatInt(int64(11-sum), 10)
}

func check(doc string, re *regexp.Regexp, digit int) (bool, error) {
	var pos int

	doc = Clean(doc)
	if doc == "" {
		return false, errorValidate
	}
	if !validateFormat(re, doc) {
		return false, errorValidate
	}

	pos = len(doc) - 2

	// Calculates the first digit.
	if calculateDigit(doc[:pos], digit) != doc[pos:pos+1] {
		return false, errorDigit
	}

	// Calculates the second digit.
	if calculateDigit(doc[:pos+1], digit+1) != doc[pos+1:pos+2] {
		return false, errorDigit
	}

	return true, nil
}
