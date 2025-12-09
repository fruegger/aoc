package day_9

import common.AocInput
import common.AocReader
import common.Position
import common.VT100
import kotlin.math.abs
import kotlin.math.max
import kotlin.math.min

fun main() {
    val input = AocReader.startDay(9, "input")
    input.print()
    val part1Floor = Part1Floor(input)
    val area = part1Floor.maxArea
    print("Part 1:")
    println(VT100.blue { "$area" })

    val part2Floor = Part2Floor(input)
    val area2 = part2Floor.maxArea
    print("Part 2:")
    println(VT100.blue { "$area2" })
    // 4754788350 too high
    // 2341519739 too high

}

class Part1Floor(input: AocInput) : Floor(input) {
    override fun select(p1: Position, p2: Position): Boolean = true
}

class Part2Floor(input: AocInput) : Floor(input) {
    override fun select(p1: Position, p2: Position): Boolean {
        if (p1.x == p2.x || p1.y == p2.y) {
            return true
        }

        val leftUp = innerLeftUpCorner(p1, p2)
        val rightDown = innerRightDownCorner(p1, p2)

        val candidatePoly = listOf(
            leftUp,
            Position(rightDown.x, leftUp.y),
            rightDown,
            Position(leftUp.x, rightDown.y)
        )

        for (ic in 0..3) {
            val cand1 = candidatePoly[ic]
            val cand2 = candidatePoly[if (ic + 1 == 4) 0 else ic + 1]
            for (ip in 0..<points.size) {
                val peri1 = points[ip]
                val peri2 = points[if (ip + 1 == points.size) 0 else ip + 1]
                var crosses = false
                if (peri1.x == peri2.x) {
                    // perimeter edge is vertical
                    if (cand1.x != cand2.x) {
                        // candidate edge is vertical
                        var p1y = peri1.y
                        var p2y = peri2.y
                        if (p1y > p2y) {
                            p1y = p2y
                            p2y = peri1.y
                        }
                        var c1x = cand1.x
                        var c2x = cand2.x
                        if (c1x > c2x) {
                            c1x = c2x
                            c2x = cand1.x
                        }

                        crosses =
                            (cand1.y > p1y) && (cand1.y < p2y) && (c1x < peri1.x) && (c2x > peri1.x)
                    }
                } else {
                    // perimeter edge is horizontal
                    if (cand1.x == cand2.x) {
                        // candidate edge is vertical
                        var p1x = peri1.x
                        var p2x = peri2.x
                        if (p1x > p2x) {
                            p1x = p2x
                            p2x = peri1.x
                        }
                        var c1y = cand1.y
                        var c2y = cand2.y
                        if (c1y > c2y) {
                            c1y = c2y
                            c2y = cand1.y
                        }
                        crosses =
                            (cand1.x > p1x) && (cand1.x < p2x) && (c1y < peri1.y) && (c2y > peri1.y)
                    }
                }
                if (crosses) {
                    // polygons intersect
                    return false
                }
            }
        }

        // check if any point of the candidate is outside the perimeter
        candidatePoly.forEach { cand ->

            var cnt = 0
            for (ip in 0..points.size - 1) {
                val peri1 = points[ip]
                val peri2 = points[if (ip + 1 == points.size) 0 else ip + 1]
                if (peri1.x == peri2.x && peri1.x <= cand.x) {
                    // trace an horizontaly ray and count the intersections, that means only check vertical edges to the left
                    val yp1 = min(peri1.y, peri2.y)
                    val yp2 = max(peri1.y, peri2.y)
                    if (cand.y in (yp1 + 1)..<yp2) {
                        cnt++
                    }
                }
            }
            if (cnt % 2 == 0) {
                return false
            }
        }
        return true

    }

    private fun innerLeftUpCorner(p1: Position, p2: Position): Position =
        Position(
            min(p1.x, p2.x) + 1,
            min(p1.y, p2.y) + 1
        )

    private fun innerRightDownCorner(p1: Position, p2: Position): Position =
        Position(
            max(p1.x, p2.x) - 1,
            max(p1.y, p2.y) - 1
        )
}

abstract class Floor(input: AocInput) {
    val points = input.data.map {
        val pair = it.split(",")
        Position(pair[0].toInt(), pair[1].toInt())
    }

    abstract fun select(p1: Position, p2: Position): Boolean

    val maxArea: Long
        get() {
            var cnt = 0
            var area = 0L
            for (i in 0..<points.size) {
                val p1 = points[i]
                for (j in i + 1..<points.size) {
                    val p2 = points[j]
                    if (select(p1, p2)) {
                        val newArea = calcArea(p1, p2)
                        if (newArea > area) {
                            area = newArea
                        }
                    }
                }
                cnt++
            }
            return area
        }

    fun calcArea(p1: Position, p2: Position): Long {
        val d = p1.distanceTo(p2)
        var newArea = (abs(d.dx).toLong() + 1) * (abs(d.dy.toLong()) + 1)
        if (newArea < 0) {
            newArea = -newArea
        }
        return newArea
    }
}