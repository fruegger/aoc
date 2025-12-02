package day_1

import common.AocReader
import common.VT100

fun main() {
    val input = AocReader.startDay(1, "input")
//    input.print()
    input.data.forEach {
        print("$position:")
        turn(it)
        println(" $it - -> $position ...cross $turns")
    }
    print("Part 1 : ")
    println(VT100.blue { "$zeroes" })
    print("Part 2 : ")
    println(VT100.blue { "${zeroCrosses}" })

}

var position = 50
var zeroes = 0
var zeroCrosses = 0
var turns = 0

private fun amount(instruction: String) = instruction.substring(1).toInt()

fun turn(instruction: String) {
    var shift = amount(instruction)
    turns = 0
    when (instruction[0]) {
        'L' -> {
            while (shift>0) {
                position--
                if (position<0) {position=99}
                shift--
                if (position==0) {
                    turns++
                }
            }
        }

        'R' -> {
            while (shift>0) {
                position++
                if (position>=100) {position=0}
                shift--
                if (position==0) {
                    turns++
                }
            }
        }
        else -> {
            println(VT100.red { "ERROR" })
        }
    }
    if (position == 0) {
        zeroes++
    }
    zeroCrosses+=turns
}

