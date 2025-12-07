package day_6

import common.AocReader
import common.VT100

fun main() {
    val input = AocReader.startDay(6, "input")
    input.print()
    val operations = parseLine(input.data.last())
    val data = mutableListOf<List<Long>>()

    input.data.filter { it != input.data.last() }.forEach {
        val elements = parseLine(it)
        data.add(elements.map { el -> el.toLong() })
    }

    val total = doCephalopodMath(operations, data)
    print("Part 1: ")
    println(VT100.blue { "$total" })

    val data2asStrings = mutableListOf<List<String>>()
    val data2 = mutableListOf<MutableList<Long>>()
    val startingPositions = findStartingPositions(input.data.last())
    input.data.filter { it != input.data.last() }.forEach {
        val elements = parseLine2(it, startingPositions)
        data2asStrings.add(elements.map { el -> el })
        data2.add(elements.map { 0L }.toMutableList())
    }

    for (i in 0..<data2asStrings[0].size) {
        for (j in 0..<data2asStrings.size) {
            val st = data2asStrings[j][i]
//            for (n in 0..<data2asStrings.size) {
            for (n in 0..<st.length) {
                val digit = st[st.length-n-1]
                if (digit!=' ') {
                    data2[n][i] = data2[n][i] * 10L + (digit - '0')
                }
            }
        }
    }
    val total2 = doCephalopodMath(operations, data2)
    print("Part 2: ")
    println(VT100.blue { "$total2" })
//5129287 too low

}

private fun doCephalopodMath(
    operations: List<String>,
    data: List<List<Long>>
): Long {
    var total = 0L
    for (i in 0..<operations.size) {
        var result = 0L
        when (operations[i]) {
            "+" -> {
                data.forEach {
                    result += it[i]
                }
            }

            "*" -> {
                result = data[0][i]
                data.filter { it != data.first() }.forEach {
                    if (it[i]>0) {
                        result *= it[i]
                    }
                }
            }

            else -> throw Exception("what the Op ?")
        }
        total += result
    }
    return total
}

fun parseLine(st: String): List<String> {
    val elements = st.split(" ")
    val result = mutableListOf<String>()
    elements.forEach {
        if (it.isNotEmpty()) {
            result.add(it)
        }
    }
    return result
}

fun parseLine2(st: String, startingPos: List<Int>): MutableList<String> {
    val result = mutableListOf<String>()
    for (i in 0..<startingPos.size - 1) {
        result.add(st.substring(startingPos[i], startingPos[i + 1]-1))
    }
    result.add(st.substring(startingPos.last()))
    return result
}

fun findStartingPositions(st: String): List<Int> {
    val result = mutableListOf<Int>()
    for (i: Int in 0..<st.length) {
        if (st[i] != ' ') {
            result.add(i)
        }
    }
    return result
}