# TinyMustache
Lightweight Mustache logic-less template implementation including a evaluation parser

For some projects it is usesful to do calculations with formulas coming from a configuration.
Nothing high sophisticated, but in detail it may be tricky

Examples:  
`({{price}}*{{win_factor}})*{{tax_factor}}`  
`{{radius}}*2*{{pi}}`  
`5*2+(77.77/66.66+(34*35*36)*8)-800*6`  
`<H1>{{header}}</H1>`  
`USER {{username}} last login {{date_and_time}}`      
will be 

##Getting tiny mustache

`go get github.com/wolfpassing/timus` on command line   
or  
`import "github.com/wolfpassing/timus"`  in your go source

##Initialize tiny mustache
First we need an TinyMustache object, that holds all the information
```go
myTiMus := timus.NewMustache()
```

##Simple use of mustache replacement
Now we can feed this object with simple key and value information
```go
myTiMus.Add("pi",math.Pi)                   //float, {{ and }} will be added to key
myTiMus.Add("{{eulers}}","2.7182818284")    //string
myTiMus.Add("radius",50)                    //integer
myTiMus.Add("valid}}}}}","this works too")  //key will be corrected to "{{valid}}"
```

For the simple Mustache replace use the following function
```go
replacedString := myTiMus.Mustache("The constant 'e' = {{eulers}}")
``` 
>As long as you are only replacing one value a simple fmt.Sprintf would do the same job of course. The power lays in multiple occurrences of multiple replacements

##Using the evaluator
Now a step further we use the evaluator for a simple task
```go
result := myTiMus.Evaluate("100+10+1")
```

combined with mustache replacement to calculate the circumference

```go
evaluateMe := "( {{radius}} * 2) * {{pi}} "
result := myTiMus.Evaluate(myTiMus.Mustache(evaluateMe))
``` 

#
