
package lexer

type Lexer struct {
  input string
  cur int
  next int
  ch byte
}

func New(input string) *Lexer {
  l := &Lexer{input: input}
  return l
}
