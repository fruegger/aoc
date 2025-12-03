package day_3

import common.AocReader
import common.VT100

fun main() {
    val input = AocReader.startDay(3, "input")
    input.print()
    var tot = 0L
    input.data.forEach {
        val max = max2Batteries(it)
        println(max)
        tot += max
    }
    print("Part 1: ")
    println(VT100.blue { "$tot" })

    var tot2 = 0L
    input.data.forEach {
        val max = maxNBatteries(it, 12)
        println(max)
        tot2 += max
    }
    print("Part 2: ")
    println(VT100.blue { "$tot2" })

}

fun max2Batteries(line: String): Long {
    var max = 0L
    for (i in 0..<line.length) {
        for (j in i + 1..<line.length) {
            val newVal = (line[i] - '0') * 10L + (line[j] - '0')
            if (newVal > max) {
                max = newVal
            }
        }
    }
    return max
}

fun maxNBatteries(line: String, n: Int): Long =
    maxN(line,n,0)

fun maxN(line : String, n: Int, subTotal : Long) : Long {
    if (n<=0) {
        return subTotal
    } else {
        val p1 = firstMax(line, n)
        val tot = subTotal * 10 + (line[p1] - '0')
        return maxN(line.substring(p1+1),n-1, tot)
    }
}

private fun firstMax(line: String, n: Int) : Int{
    var p1 = 0

    var max = line[0]
    var i = 1
    while (i<line.length-n+1 && max < '9') {
        if ( line[i]>max) {
            max = line[i]
            p1 = i
        }
        i++
    }
    return p1
}