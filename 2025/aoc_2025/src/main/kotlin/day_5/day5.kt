package day_5

import common.AocReader
import common.VT100

fun main() {
    val input = AocReader.startDay(5, "input")
    input.print()
    val ranges = mutableListOf<LongRange?>()
    val inventory = mutableListOf<Long>()

    input.data.forEach {
        if (it.contains('-')) {
            val values = it.split('-')
            if (values[0].toLong()> values[1].toLong()) {
                throw Exception("what the input ?")
            }
            ranges.add(LongRange(values[0].toLong(), values[1].toLong()))
        } else {
            if (it.isNotEmpty()) {
                inventory.add(it.toLong())
            }
        }
    }

    val result = selectFresh(inventory, ranges)

    print("Part 1: ")
    println(VT100.blue { "${result.size}" })
    do {
        var changes = false
        println ("iteration")
        for (i in 0..<ranges.size) {
            for (j in 0..<ranges.size) {
                if (i != j) {
                    val newRange = ranges[i]?.handleOverlaps(ranges[j])

                    if (newRange != null && newRange.first > newRange.last) {
                        throw Exception("what the range ?")
                    }

                    if (newRange != ranges[i]) {
                        ranges[i] = newRange
                        changes = true
                    }
                }
            }
        }
    } while (changes)

    var result2 = 0L
    ranges.forEach { range ->
        range?.let {
            result2 += it.last - it.first + 1
        }
    }

    print("Part 2: ")
    println(VT100.blue { "${result2}" })
// 355958587454579 too low
// 358298154172769 too high
// 357674099117260

    for (i in 0..<ranges.size) {
        for (j in 0..<ranges.size) {
            if (i != j) {
                if (ranges[i]!= null && ranges[j]!=null && !ranges[i]!!.isDisjointFrom(ranges[j]!!)) {
                    throw Exception("not disjoint : $i, $j")
                }
            }
        }
    }
}

private fun selectFresh(
    inventory: MutableList<Long>,
    ranges: MutableList<LongRange?>
): MutableSet<Long> {
    val result = mutableSetOf<Long>()

    inventory.forEach { item ->
        ranges.forEach { range ->
            range?.let {
                if (it.contains(item)) {
                    result.add(item)
                    return@forEach
                }
            }
        }
    }
    return result
}

fun LongRange.handleOverlaps(other: LongRange?): LongRange? {
    if (other == null) {
        return this
    }
    if (isDisjointFrom(other)) {
        return this
    }
    if (isIncludedIn(other)) {
        return null
    }
    if (overlapsBefore(other)) {
        return LongRange(this.first, other.first - 1)
    }
    if (overlapsAfter(other)) {
        return LongRange(other.last + 1, this.last)
    }
    if (other.isIncludedIn(this)) {
        // will be taken car of in case 2 (this.isIncludedIn(other)) later in the loop
        return this
    }
    throw Exception("what the hack?")
}

fun LongRange.isDisjointFrom(other: LongRange): Boolean =
    this.first > other.last || this.last < other.first

fun LongRange.isIncludedIn(other: LongRange): Boolean =
    this.first >= other.first && this.last <= other.last

fun LongRange.overlapsBefore(other: LongRange): Boolean =
    this.last >= other.first && this.last < other.last && this.first<other.first

fun LongRange.overlapsAfter(other: LongRange): Boolean =
    this.first <= other.last && this.first > other.first && this.last>other.last
