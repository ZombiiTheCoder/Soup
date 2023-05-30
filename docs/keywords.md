# [<](readme.md) Keywords

Not Reassignable Variable [**val**]()
<br/>
Reassignable Variable [**var**]()
<br/>
Imports [**use**](#function) or [**Use**](#function)
<br/>
Functions [**fn**](#function), [**func**](#function), or [**function**](#function)
<br/>
Returns [**ret**](#return) or [**return**](#return)
<br/>
[**If**](#If_Else)
<br/>
[**Else**](#If_Else)
<br/>
[**While**](#While_Loop)

# val
```soup
val x = 1
```

# var
```soup
var x = 1
```

# function
```soup
?? using @ in your import declarations points it to the "package"s folder and without it it points to your current file's folder 

import `@std`
fn Hello_World() {
    std.io.println(`Hello World`)
}

Hello_World()
```

# return

```soup
?? Using Return In Your Function ends it and outputs your desiried value
import `@std`

fn Hello_World() {
    ret `Hello World`
}

std.io.println(Hello_World())
```

# If_Else
```soup
?? Using The If and Else together you can make some nice conditional logic

val v = `A`
if (v == `A`) {
    std.io.println(`The Value is A`)
}else {
    std.io.println(`The Value Does Not Equal Any Preset Values`)
}

?? You can chain these together to make a tree of conditional logic

val value = `A`
if (value == `A`) {
    std.io.println(`The Value is A`)
} else if (value == `B`){
    std.io.println(`The Value is B`)
} else if (value == `C`) {
    std.io.println(`The Value is C`)
} else {
    std.io.println(`The Value Does Not Equal Any Preset Values`)
}
```

# While_Loop
```soup
?? using the while-loop you can make an infinite loop. Or you can make a loop to certain numbers

var i = 0
while ()

while (true) {
    std.io.println(`Hello World`)
}

```