Printf has over a dozen such conversions, which Go programmers call verb s. This table is far
from a complete specification but illustrates many of the features that are available:
%d decimal integer
%x, %o, %b integer in hexadecimal, octal, binar y
%f, %g, %e floating-point number: 3.141593 3.141592653589793 3.141593e+00
%t boolean: true or false
%c rune (Unico de co de point)
%s string
%q quoted string "abc" or rune 'c'
%v any value in a natural format
%T type of any value
%% literal percent sign (no operand)
