package language

import (
	"bytes"
	"strings"
	"unicode"
)

var reservedKeywords = [...]string{
	"__halt_compiler",
	"abstract",
	"and",
	"array",
	"as",
	"break",
	"callable",
	"case",
	"catch",
	"class",
	"clone",
	"const",
	"continue",
	"declare",
	"default",
	"die",
	"do",
	"echo",
	"else",
	"elseif",
	"empty",
	"enddeclare",
	"endfor",
	"endforeach",
	"endif",
	"endswitch",
	"endwhile",
	"eval",
	"exit",
	"extends",
	"final",
	"for",
	"foreach",
	"function",
	"global",
	"goto",
	"if",
	"implements",
	"include",
	"include_once",
	"instanceof",
	"insteadof",
	"interface",
	"isset",
	"list",
	"namespace",
	"new",
	"or",
	"print",
	"private",
	"protected",
	"public",
	"require",
	"require_once",
	"return",
	"static",
	"switch",
	"throw",
	"trait",
	"try",
	"unset",
	"use",
	"var",
	"while",
	"xor",
	"int",
	"float",
	"bool",
	"string",
	"true",
	"false",
	"null",
	"void",
	"iterable",
	"yield",
	"match",
}

type PHP struct {
}

// isReserved check if the name is a reserved keyword
func (_ PHP) isReserved(name string) bool {
	name = strings.ToLower(name)
	for _, k := range reservedKeywords {
		if name == k {
			return true
		}
	}
	return false
}

// Identifier snake_case to CamelCase
func (p PHP) Identifier(name string, suffix string) string {
	name = p.Camelize(name)
	if suffix != "" {
		return name + p.Camelize(suffix)
	}
	return name
}

func (p PHP) resolveReserved(identifier string, pkg string) string {
	if p.isReserved(strings.ToLower(identifier)) {
		if pkg == ".google.protobuf" {
			return "GPB" + identifier
		}
		return "PB" + identifier
	}

	return identifier
}

// Camelize "dino_party" -> "DinoParty"
func (p PHP) Camelize(word string) string {
	words := p.splitAtCaseChangeWithTitlecase(word)
	return strings.Join(words, "")
}

func (_ PHP) splitAtCaseChangeWithTitlecase(s string) []string {
	words := make([]string, 0)
	word := make([]rune, 0)
	for _, c := range s {
		spacer := isSpacerChar(c)
		if len(word) > 0 {
			if unicode.IsUpper(c) || spacer {
				words = append(words, string(word))
				word = make([]rune, 0)
			}
		}
		if !spacer {
			if len(word) > 0 {
				word = append(word, unicode.ToLower(c))
			} else {
				word = append(word, unicode.ToUpper(c))
			}
		}
	}
	words = append(words, string(word))
	return words
}

func isSpacerChar(c rune) bool {
	switch {
	case c == rune("_"[0]):
		return true
	case c == rune(" "[0]):
		return true
	case c == rune(":"[0]):
		return true
	case c == rune("-"[0]):
		return true
	}
	return false
}

func (p PHP) Namespace(pkg *string, sep string) string {
	if pkg == nil {
		return ""
	}
	result := bytes.NewBuffer(nil)
	for _, pk := range strings.Split(*pkg, ".") {
		result.WriteString(p.Identifier(pk, ""))
		result.WriteString(sep)
	}

	return strings.Trim(result.String(), sep)
}
