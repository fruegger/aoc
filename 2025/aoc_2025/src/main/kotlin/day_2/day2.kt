package day_2

import common.AocReader
import common.VT100

fun main() {
    val input = AocReader.startDay(2, "input")
    val ids = input.data[0].split(',')
    val total1 = part1(ids)
    print("Part 1: ")
    println(VT100.blue { "$total1" })
    val total2 = part2(ids)
    print("Part 2: ")
    println(VT100.blue { "$total2" })
}

fun part1(ids: List<String>): Long {
    var total = 0L
    ids.forEach { idRange ->
        val producTId = idRange.toProductId()
        print(producTId.first)
        print("-")
        println(producTId.last)
        producTId.forAllInvalid(2) { id ->
            total += id.toLong()
            print(" ")
            print(VT100.white { id })
        }
        println()
    }
    return total
}

fun part2(ids: List<String>): Long {
    var total = 0L
    ids.forEach { idRange ->
        val matches = mutableSetOf<String>()
        val productId = idRange.toProductId()
        print(productId.first)
        print("-")
        println(productId.last)
        for (reps in 2..productId.last.length) {
            productId.forAllInvalid(reps) { id ->
                matches.add(id)
            }
        }
        matches.forEach {
            total += it.toLong()
            print(" ")
            print(VT100.white { it })
        }
        println()
    }
    return total
}
