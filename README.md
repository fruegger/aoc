# Advent of Code
Every year, at the gates of December, a new adventurous and thrilling coding path towards Christmas appears through the slowly dissolving but yet incumbing layers of the all pervading
developer's everyday.  

https://adventofcode.com/ 

This is my diary of the effort I put in to travel this path and keep up with this nice albeit geeky and challenging tradition:

The curious, odd, mildly amusing or even remotely noteworthy is collected here.

P.S. The appearance of most emojis (such as 😆 or 🙄) warn of ugly puns in the vicinity  

## 2024 edition

For 2024, I wanted to take a short coding vacation and revisit a long forgotten place I had seen before. 
This time I'd travel to a different programming coutry I had more or less fond remembrances of.
As any self respecting tourist would, I surely would not miss any of the oddities, bizzare landmarks and memorable habits on my trip.

And so it was that I made my way through goLand.

Many a difficult piuzzle to solve would soon cross my way; armed with the utmost confidence that I'd always find the stupidest and most naively intricate solution, 
I took that first stride on my quest to find the silliest way possible to use goLand features, constructs and particularities.

This are some entries in my log book:  

---
### Day 1
Here we go again : the usual struggle with more than untotal recall, fortunately I always carry a cheat sheet with me, here is an extract:

* What was the difference between array and slices again ?
```An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.```

* And what is a rune if not a glorified character ? 
``The Go language defines the word rune as an alias for the type int32``
indeed : glorified string characters.

* How did range loop work ?
```go
for i, s := range strings {
	fmt.Println(i, s)
}
```
that's neat i is the index, s is the value at that position ... also works for maps, arrays.

* What are the names of the commonly used packages and functions I need ?
.. come on now; go find them.

---
### Day 2 
I have a strong feeling that I'm moving in circles ... Things are definitely repeating : time to define a module.
goLang does not let you define more than one package in the same folder; 
so I just defined a module, which is nothing else thah a collection of packages.

Also I moved some code I'd use over and over in a common package. 

The odd thing to remember, is that everything Capitalized is public, all smallcase is private : go figure ( :laughing: ).  

---
### Day 3
I will not use regex again, they are evil, i will not ....
... and then again ... but, NO, really no; i will not use regexes if it kills me.

And, what's more: a fabulous opportunity for a trip along memory lane... oh, well, maybe somebody else's memory 
if I'm completely honest.
Be it as it will: it's hard to use up more lines of code to attack this otherwise quite scantily solvable puzzle than 
to write a full fledged character-per-character reading state machine... all hand crafted, beautifully wasteful and awesomely overcomplicated.

Enjoy!

By the way: the odd ball in there is the iota definition for constants: it increases every following value without you declaring it :
exactlty what I thought : " say what ?"
```go
type ScanState uint8
const (
	SCAN_Start ScanState = iota
	SCAN_M
	SCAN_U
	...
)
```

---
### Day 4
This is the first riddle for this year, where I'm sanguine that it is a good idea to exercise a little on the parking lot and try 
my solution on the example data in the story before dashing out on the real data interstate.
All nicely straight forward, linear and repetitive ... good boring stuff.
The only sparkle of worrying semi-intelligence I'm coughing up is the fact that I start the MAS check in the middle: with th A.

Knowing that the AoC chaps and gentleladies are nice people at heart ... I assumed (correctly) that all lines have the same length.

---
### Day 5

Hmmm... I could read and sort the rules, create a list of known pages ... write my own sort ...

I'll have none of that (even though a quick dash into quick sort territory would create sweetly complicated and ugly code); 

I have a much sillier idea: the first line of code says it all:
```go
// order[x][y] >0  -> x follows y
// order[x][y] <0  -> x precedes y
// order[x][y] =0  -> order is unknown
var order [100][100]int
```
I'll build a comparison table that tells me for every pair of pages, which one should come first.
I will even pretend I don't know it's symmetric ... no sir ... you have memory ... waste it !

Have I already uttered my utmost appreciation for the way the AoC girls and blokes never seem to trick you into goTCHAland ? (🙄) Well, without fail; there are indeed no pages in the updates that are not mentioned in the rules... 

... and all strings have an odd length so that center pages indexes are always integers.

that's nice.

On the purely syntactic side - the lambda syntax used to pass a comparator to the generic sort is quite odd.
```go
sort.Slice(v,func(i, j int) bool {
	return order[v[i]][v[j]] == -1
})
```
---
### Day 6

Yeeeees !! A Labyrinth - I will certainly find a use for dropping breadcrumbs at the slightest provocation.
A look to the left, one to the right ... no Minotaur in sight (😁), no competing Ariadne, no bird eating up my breadcrumbs ..

... ret's curse and recurse!
```go
func predictPath(p Position, d Direction, lines []string) {
	obj := lookAhead(p, d, lines)
	if obj == OBJ_Nothing {
		replaceSymAtPos(p, lines, SYM_Breadcrumb)
		predictPath(p.move(d), d, lines)
	} else {
		if obj == OBJ_Obstacle {
			predictPath(p, d.turnRight(), lines)
		} else {
			replaceSymAtPos(p, lines, SYM_Breadcrumb)
		}
	}
}
```
The annoying bit was the fact that replacing a character in an existing string in go is a bit of a hassle.
Also making copies of array is a pain in the pointers... mentioning which: 
This time I was in luck - passing arguments to functions by reference or by value is immaterial in this case.

I have done nothing horribly silly today ... yet. I will correct the issue and goof around idly with ANSI / VT100 Terminal 
escape sequences.
```go
const GREEN = "\x1B[32m"
const BLUE = "\x1B[34m"
const WHITE = "\x1B[97m"
const RED = "\x1B[31m"
const YELLOW = "\x1B[33m"
```
I also wanted to use cusror motion sequences. The effect of watching breadcrumbs and obstructions drop on the map as the solution is calculated 
is very nice on the small example map, but unfortunately will scramble up the real map completely.

So I'll b e satisfied with a childishly colorful depiction of the map... glorious waste of time! 

---
### Day 7
The day before yesterday, I advocated profusely that wasting memory space is an art to be trained and perfected. Loud acclamations failed to echo, heads did surprisingly not nod accondiscendently and hands did not clap.

To cut a long story short ... the world was unimpressed, unconvinced uand vastly unaffected, hence my change of heart: maybe memory frugality is not the worst of virtues.
Today I therefore persist in my quest for imbecile coding and will steer in the diametrally opposite direction; I'll pack as much as I can into as little memory as I can .. and since I am at it, I will also try to write some seriously dense code.

I can save a few lines of code by NOT creating permutations of operations or defining those operation in the first place. 

I'll just increase a good old integer counter and look at it as a permutation of bits each denoting an operation(or actually groups of two-bits, for the second part); 

I'll certainly not miss the opportunity to improve the code illegibility in that i'll use deliciously unreadable fully parentesized shifts and bitwise operations to extract the individual operations from the permutation; No need for verbous opeation names or structures; let's just hardcode everything compactly.

```go
func findOpCombination(values []int, expectedResult int) bool { ...
	for combination := 0; combination < 1<<((len(values)-1)<<1) && incorrect; combination++ {
		sum := values[0]
		for i := 0; i < len(values)-1; i++ {
			sum = applyOp((combination>>(i<<1))&3, sum, values[i+1])
		}...}
```
Here applyOp decides what operation to apply depending on the first argument ( 0 for +, 1 for *, 2 for ||).. if it is 3 then just apply +; it will slow things down,  but will not affect the result; the code will be equally hard to understand and no memory will be used. Nice!

The bit I am proudest of, is the way concat is implemented: I originally wanted to roll out the loops and hard-code the ifs ... that would have looked gorgeously crazy, but compactness won gthe argument... and I desisted. 

```go
func concat(o1 int, o2 int) int {
	for decade := 10; ; decade *= 10 {
		if o2 < decade {
			return o1*decade + o2
		}
	}
}
```
Quick, memory saving, ugly as sin; love it.  All in 73 lines of code inluding all boilerplate... a lovely nectar of code stupidity.  

---
### Day 8 
