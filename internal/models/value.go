package models

import "fmt"

type Value struct {
	Content string `json:"content"`
	Sign    bool   `json:"sign"`
}

func (v Value) Multiply(other Value) Value {
	v.Content = fmt.Sprintf("((%s)*(%s))", v.Content, other.Content)
	v.Sign = other.Sign && v.Sign
	return v
}

func (v Value) Divide(other Value) Value {
	v.Content = fmt.Sprintf("((%s)/(%s))", v.Content, other.Content)
	v.Sign = other.Sign && v.Sign
	return v
}

// add
func (v Value) Plus(other Value) Value {
	var (
		localSignOtherStr   string
		localSignCurrentStr string
	)

	localSignOtherStr = GetSign(true == other.Sign)
	localSignCurrentStr = GetSign(v.Sign)

	v.Content = fmt.Sprintf("(%s1*(%s)*%s1*(%s))", localSignCurrentStr, v.Content, localSignOtherStr, other.Content)
	return v
}

func (v Value) Minus(other Value) Value {
	var (
		localSignOtherStr   string
		localSignCurrentStr string
	)

	localSignOtherStr = GetSign(false == other.Sign)
	localSignCurrentStr = GetSign(v.Sign)

	v.Content = fmt.Sprintf("(%s1*(%s)*%s1*(%s))", localSignCurrentStr, v.Content, localSignOtherStr, other.Content)
	return v
}

func (v Value) String() string {
	return fmt.Sprintf("%s1*(%s)", GetSign(v.Sign), v.Content)
}

func (v Value) StringFinal() string {
	return fmt.Sprintf("%s", v.Content)
}

func GetSign(sign bool) string {
	if sign {
		return "+"
	} else {
		return "-"
	}
}
