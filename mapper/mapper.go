package mapper

import "unicode"

func CapitalizeEveryThirdAlphanumericChar(in string) string {
	arr := []rune(in)
	var out []rune
	for i, batch := 0, 0; i < len(arr); i++ {

		if IsAlphaNumeric(arr[i]) {
			if batch == 2 {
				out = append(out, unicode.ToUpper(arr[i]))
				batch = 0
				continue
			}
			out = append(out, unicode.ToLower(arr[i]))
			batch += 1
		} else {
			out = append(out, arr[i])
		}
	}

	return string(out)
}

type Interface interface {
	TransformRune(pos int)
	GetValueAsRuneSlice() []rune
}

type SkipString struct {
	BatchSize int
	Text      string
	Step      int
}

func NewSkipString(batchSize int, text string) Interface {
	return &SkipString{Text: text, BatchSize: batchSize, Step: 0}
}

func (st SkipString) String() string {
	return st.Text
}

func (st *SkipString) GetValueAsRuneSlice() []rune {
	arr := []rune(st.Text)
	return arr
}

func (st *SkipString) TransformRune(pos int) {
	ch := rune(st.Text[pos])

	if IsAlphaNumeric(ch) {
		if st.Step == st.BatchSize-1 {
			st.Text = st.Text[:pos] + string(unicode.ToUpper(ch)) + st.Text[pos+1:]
			st.Step = 0
		} else {
			st.Text = st.Text[:pos] + string(unicode.ToLower(ch)) + st.Text[pos+1:]
			st.Step += 1
		}

	} else {
		st.Text = st.Text[:pos] + string(ch) + st.Text[pos+1:]
	}
}

func MapString(i Interface) {
	for pos, _ := range i.GetValueAsRuneSlice() {
		i.TransformRune(pos)
	}
}

func IsAlphaNumeric(c rune) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || ('0' <= c && c <= '9')
}
