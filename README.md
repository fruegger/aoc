# Advent of Code
Every year, at the gates of December, a new adventurous and thrilling coding path towards Christmas appears through the slowly dissolving but yet incumbing layers of the all pervading
developer's everyday.  

https://adventofcode.com/ 

This is my diary of the effort I put in to travel this path and keep up with this nice albeit geeky and challenging tradition:

The curious, odd, mildly amusing or even remotely noteworthy is collected here.

## 2004 edition

For 2004, I wanted to take a short coding vacation and revisit a long forgotten place I had seen before. 
This time I'd travel to a different programming coutry I had more or less fond remembrances of.
As any self respecting tourist would, I surely would not miss any of the oddities, bizzare landmarks and memorable habits on my trip.

And so it was that I made my way through goLand.

Many a difficult piuzzle to solve would soon cross my way; armed with the utmost confidence that I'd always find the stupidest and most straightforward solution, 
I took that first stride on my quest to find the silliest way possible to use goLand features, constructs and particularities.

This are some entries in my log book:  

#### Day 1
Here we go again : the usual struggle with more than untotal recall:

What was the difference between array and slices again ?

And what is a rune if not a glorified character ? 

What are the names of the commonly used packages and functions I need ?


#### Day 2 
I have a strong feeling that I'm moving in circles ... Things are definitely repeating : tme to define a module.
goLang does not let you define more than one package in the same folder; 
so I just defined a module, which is nothing else thah a collection of packages.

Also I moved some code I'd use over and over in a common package. 

The odd thing to remember, is that everything Capitalized is public, all smallcase is private : go figure ( :laughing: ).  

#### Day 3
I will not use regex again, they are evil, i will not ....
... and then again ... but, NO, really no; i will not use regexes if it kills me.

And, what's more: a fabulous opportunity for a trip along memory lane... oh, well, maybe somebody else's memory 
if I'm completely honest.
Be it as it will: it's hard to use up more lines of code to attack this otherwise quite scantily solvable puzzle than 
to write a full fledged character-per-character reading state machine... all hand crafted, beautifully wasteful and awesomely overcomplicated.

Enjoy!

By the way: the odd ball in there is the iota definition for constants: it increases every following value without you declaring it :
exactlty what I thought : " say what ?".iota

#### Day 4
This is the first riddle for this year, where I'm sanguine that it is a good idea to exercise a little on the parking lot and try 
my solution on the example data in the story before dashing out on the real data interstate.
All nicely straight forward, linear and repetitive ... good boring stuff.
The only sparkle of worrying semi-intelligence I'm coughing up is the fact that I start the MAS check in the middle: with th A.

Knowing that the AoC chaps and gentleladies are nice people at heart ... I assumed (correctly) that all lines have the same length.

#### Day 5

Hmmm... I could read and sort the rules, create a list of known pages ... write my own sort ...

I'll have none of that (even though a quick dash into quick sort territory would create sweetly complicated and ugly code); 

I have a much sillier idea: the first line of code says it all:
`
// order[x][y] >0  -> x follows y
// order[x][y] <0  -> x precedes y
// order[x][y] =0  -> order is unknown
var order [100][100]int
`
I'll build a comparison table that tells me for every pair of pages, which one should come first.
I will even pretend I don't know it's symmetric ... no sir ... you have memory ... waste it !

Have I already uttered my utmost appreciation for the way the AoC girls and blokes never seem to trick you into goTCHAland ?

Well, without fail; there are indeed no pages in the updates that are not mentioned in the rules... 

... and all strings have an odd length so that center pages indexes are always integers.

that's nice.

On the purely syntactic side - the lambda syntax used to pass a comparator to the generic sort is quite odd.

#### Day 6

Yeeeees !! A Labyrinth - I will certainly find a use for dropping breadcrumbs at the slightest provocation.
A look to the left, one to the right ... no Minotaur in sight, no competing Ariadne, no bird eating up my breadcrumbs ..

... ret's curse and recurse!

The annoying bit was the fact that replacing a character in an existing string in go is a bit of a hassle.
Also making copies of array is a pain in the pointers... mentioning which: 
This time I was in luck - passing arguments to functions by reference or by value is immaterial in this case.


I have done nothing horribly silly today ... I will correct the issue and goof around idly with ANSI / VT100 Terminal 
escape sequences.
The effect of watching breadcrumbs and obstructions drop on the map as the solution is calculated 
is very nice on the small example map, but unfortunately will scramble up the real map completely.

So I'll b e satisfied with a childishly colorful depiction of the map... glorious waste of time! 
