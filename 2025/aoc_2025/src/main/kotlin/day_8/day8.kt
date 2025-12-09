package day_8

import common.AocReader
import common.VT100

data class Point(
    val x: Long,
    val y: Long,
    val z: Long,
) {
    fun distanceTo2(other: Point): Long =
        (other.x - x) * (other.x - x) +
                (other.y - y) * (other.y - y) +
                (other.z - z) * (other.z - z)

    fun print() {
        print("$x, $y, $z")
    }
}

data class Box(
    val pos: Point
) {
    val connections = mutableSetOf<Box>()
    fun isConnectedTo(other : Box) = connections.contains(other)
}

typealias Circuit = MutableSet<Box>
typealias CircuitList = MutableList<Circuit>

fun main() {
    val input = AocReader.startDay(8, "input")
    input.print()
    val boxes = input.data.map {
        val values = it.split(",")
        Box(
            pos = Point(values[0].toLong(), values[1].toLong(), values[2].toLong())
        )
    }
    val circuits: CircuitList = mutableListOf()
    var cnt = 1
    while (cnt <= 1000) {
        val nextPair = findClosestBoxes(boxes)
        if (nextPair != null) {
            circuits.addPair(nextPair)
            cnt++
        }
    }
    circuits.sortByDescending { it.size }

    val tot = circuits[0].size * circuits[1].size * circuits[2].size
    print("Part 1: ")
    println(VT100.blue { "$tot" })

    var lastPair : Pair<Box, Box>? = null
    while (circuits.size>1 || boxes.any { it.connections.isEmpty()} ) {
        val nextPair = findClosestBoxes(boxes)
        if (nextPair != null) {
            circuits.addPair(nextPair)
            lastPair = nextPair
        }
        var sum= 0
        boxes.forEach {
            if (it.connections.isEmpty()) sum++
        }
        println("$sum - ${circuits.size}")
    }


    val tot2 = lastPair!!.first.pos.x * lastPair.second.pos.x
    print("Part 2: ")
    println(VT100.blue { "$tot2" })

}

fun findClosestBoxes(
    boxes: List<Box>,
): Pair<Box, Box>? {
    var mindist = Long.MAX_VALUE
    var result: Pair<Box, Box>? = null
    boxes.forEach { b1 ->
        boxes.forEach { b2 ->
            if (b1 != b2 && !b1.isConnectedTo(b2)) {

                val dist = b1.pos.distanceTo2(b2.pos)
                if (dist < mindist) {
                    mindist = dist
                    result = Pair(b1, b2)
                }
            }
        }
    }
    return result
}

fun CircuitList.addPair(nextPair : Pair<Box,Box>) {
    nextPair.first.connections.add(nextPair.second)
    nextPair.second.connections.add(nextPair.first)

    val circuit1 = findCircuitFor(nextPair.first)
    val circuit2 = findCircuitFor(nextPair.second)
    if (circuit1 == null) {
        if (circuit2 == null) {
            val circuit = mutableSetOf(nextPair.first, nextPair.second)
            add(circuit)
        } else {
            circuit2.add(nextPair.first)
        }
    } else {
        if (circuit2 == null) {
            circuit1.add(nextPair.second)
        } else {
            if (circuit1 != circuit2) {
                circuit1.addAll(circuit2)
                remove(circuit2)
                // cannot merge yet
                /*                        circuit1.addAll(circuit2)
                                        circuits.remove(circuit2)
                                        circuit1.add(nextPair)
                */
            }
        }
    }
}

fun CircuitList.findCircuitFor(b: Box): Circuit? = firstOrNull { list ->
    list.contains(b)
}

