# Golang TypeWriter
This app allows you to type content of the files to terminal with the fixed time for every line

## Example
>file text.txt
```txt 
竜田川
神代も聞
竜田川も聞
hello
how are you?
how are you?how are you?
川u?how arもe yo
```
>You shold specify the input text file as the `first` parameter and typing delay in <u>milliseconds</u> as `second`
```bash
$ go run cmd/main.go text.txt 100
```
>Command outpul will be looking something like this
```bash
100.117ms  ->  竜田川
101.3499ms ->  神代も聞
101.7448ms ->  竜田川も聞
102.0256ms ->  hello
100.1498ms ->  how are you?
103.676ms  ->  how are you?how are you?
102.9189ms ->  川u?how arもe yo
```