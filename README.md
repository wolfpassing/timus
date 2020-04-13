# Tiny Mustache & Eval
**Lightweight Mustache logic-less template implementation including a evaluation parser**
Actual version v1.3

For some projects it is useful to do calculations with formulas coming from a string or input.
Nothing high sophisticated, but in detail it may be tricky...

##### Examples Tiny Mustache:
`<H1>{{header}}</H1>`,
`USER {{username}} last login {{date_and_time}}`

##### Examples Calculator:
`5*2+(77.77/66.66+(34*35*36)*8)-800*6`

##### Examples Evaluator:
`Multiply ({{price}} by {{quantity}} equals <% {{price}} * {{quantity}} %>`

## Getting tiny mustache

`go get github.com/wolfpassing/timus` on command line   
or  
`import "github.com/wolfpassing/timus"`  in your go source

## Initialize tiny mustache
First we need an TinyMustache object, that holds all the information
```go
myTiMus := timus.NewMustache()
```

## Simple use of mustache replacement
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
Value of replacesString : `The constant 'e' = 2.7182818284`
>As long as you are only replacing one value a simple fmt.Sprintf would do the same job of course. The power lays in multiple occurrences of multiple replacements

## Using the calculator
Now a step further we use the calculator for a simple task
```go
result := myTiMus.Calculate("100+10+1")
```
Value of result: `111`  

## Using calculator, evaluator and mustache
combined we can calculate the circumference

```go
evaluateMe := "( {{radius}} * 2) * {{pi}} "
result := myTiMus.Calculate(myTiMus.Mustache(evaluateMe))
``` 
Value of result: `314.159265359`
## Merge tiny mustache 
```go
firstTiMus.Merge(secondTiMus)
```
## Extracting from a structure
First we need a structure
```go   
type Car struct {
    Brand               string      `mustache:"brand"`
    Wheels              string      `mustache:"-"`
    HorsePower          int         `mustache:"ps"`
    Doors               int         `mustache:"doors"`
    AutomaticGearBox    bool
}
```
then some code

```go
myCar := Car{
	Brand:            "BMW",
	Wheels:           4,
	HorsePower:       280,
	Doors:            4,
	AutomaticGearBox: true,
}

myMus := NewMustache()
myMus.Extract(myCar)
s := "The car is a {{brand}} with {{ps}}PS, {{doors}} doors, thats <% {{ps}}/{{wheels}}PS per wheel %>"
fmt.Println(myMus.Evluate(myMus.Mustache(s)))
```
The result:  
`The car is a BMW with 280PS, 4 doors, thats 70PS per wheel`

### Have fun...