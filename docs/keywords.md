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
<pre class="language-js">
val x = 1
</pre>

# var
<pre class="language-js">
var x = 1
</pre>

# function
<pre class="language-js">
?? using @ in your import declarations points it to the "package"s folder and without it it points to your current file's folder 

import `@std`
fn Hello_World() {
    std.io.println(`Hello World`)
}

Hello_World()
</pre>

# return

<pre class="language-js">
?? Using Return In Your Function ends it and outputs your desiried value
import `@std`

fn Hello_World() {
    ret `Hello World`
}

std.io.println(Hello_World())
</pre>

# If_Else
<pre class="language-js">
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
</pre>

# While_Loop
<pre class="language-js">
?? using the while-loop you can make an infinite loop. Or you can make a loop to certain numbers

var i = 0
while ()

while (true) {
    std.io.println(`Hello World`)
}

</pre>

<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/prism/1.23.0/themes/prism.min.css" integrity="sha512-tN7Ec6zAFaVSG3TpNAKtk4DOHNpSwKHxxrsiw4GHKESGPs5njn/0sMCUMl2svV4wo4BK/rCP7juYz+zx+l6oeQ==" crossorigin="anonymous" />
<img src="https://img.shields.io/tokei/lines/github/ZombiiTheCoder/SoupLang?style=plastic"/>

<script>
window.Prism = window.Prism || {};
window.Prism.manual = true;
</script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/prism/1.23.0/prism.min.js"></script>

<script>
    let pre = document.querySelectorAll("pre");
    for (let i = 0; i < pre.length; i++){
        let code = pre[i].innerText;
        pre[i].innerHTML = Prism.highlight(code, Prism.languages.javascript, 'javascript');

  }
</script>
