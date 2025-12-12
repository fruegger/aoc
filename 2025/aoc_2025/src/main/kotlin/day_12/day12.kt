package day_12

import common.AocInput
import common.AocReader


class Present(
    val nr: Int,
    val original: Array<Long>
) {
    // presents are always 3x3.
    val rotations: Array<Array<Long>> = Array(8) {
        when (it) {
            0 -> original
            1 -> original.rotate90()
            2 -> original.rotate90().rotate90()
            3 -> original.rotate90().rotate90().rotate90()
            4 -> original.flip()
            5 -> original.flip().rotate90()
            6 -> original.flip().rotate90()
            7 -> original.flip().rotate90()
            else -> original
        }
    }

    private fun Array<Long>.rotate90(): Array<Long> {
        val result = Array(3) { 0L }
        result[0] = (this[2] and 4L) or ((this[1] and 4L) shr 1) or ((this[0] and 4L) shr 2)
        result[1] = ((this[2] and 2L) shl 1) or (this[1] and 2L) or ((this[0] and 2L) shr 1)
        result[2] = ((this[2] and 1L) shl 2) or ((this[1] and 1L) shl 1) or (this[0] and 1L)
        return result
    }

    private fun Array<Long>.flip(): Array<Long> = Array(3) { i -> this[2 - i] }

    fun print(rotation: Int) {
        val ch = (nr + 'A'.code).toChar()
        rotations[rotation].forEach { w ->
            print(if (w and 4L == 0L) "." else ch)
            print(if (w and 2L == 0L) "." else ch)
            println(if (w and 1L == 0L) "." else ch)
        }
        println()
    }
}

class Region(
    val w: Int,
    val h: Int,
    val accepts: List<Int>
) {
    val field = Array(h) { 0L }

    fun reset() {
        for (i in 0..<field.size) {
            field[i] = 0
        }
    }

    fun fit(shape: Array<Long>, x: Int, y: Int): Boolean {
        val mask = 7L shl (w - x - 3)
        for (i in 0..2) {
            val fBits = field[y + i] and mask
            val pBits = shape[i] shl (w - x - 3)
            if ((fBits and pBits) != 0L) {
                return false
            }
            field[y + i] = field[y + i] or pBits
        }
        return true
    }

    fun createBag(): List<State> {
        val bag = mutableListOf<State>()
        accepts.forEach {
            for (ignore in 0..<it) {
                bag.add(State(this, it))
            }
        }
        return bag
    }

    fun print() {
        print("${w}x${h}: ")
        accepts.forEach {
            print("$it ")
        }
        println()
    }
}

fun List<State>.nextState() {
    forEach {
        if (it.hasNext()) {
            it.next()
            return
        } else {
            it.reset()
        }
    }
}

fun List<State>.hasNextState(): Boolean = last().hasNext()

val presents = mutableListOf<Present>()
val regions = mutableListOf<Region>()

data class State(
    val region: Region,
    val presentNr: Int,
    var rotation: Int = 0,
    var posX: Int = 0,
    var posY: Int = 0
) : Iterator<State> {

    override fun next(): State {
        if (posX + 3 < region.w) {
            posX = posX + 3
        } else {
            posX = 0
            if (posY + 3 < region.h) {
                posY+= posY+3
            } else {
                if (rotation < 7) {
                    rotation++
                }
            }
        }
        return this
    }

    override fun hasNext(): Boolean = rotation < 7 || posX + 3 < region.w || posY + 3 < region.h

    fun reset() {
        posX = 0
        posY = 0
        rotation = 0
    }
}

fun main() {
    val input = AocReader.startDay(12, "input")
    input.print()
    read(input)

    presents.forEach { it.print(0) }
    regions.forEach { it.print() }

    // for the hack of it..
    var tot = 0
    regions.forEach { region ->
        val area1 = region.w * region.h
        var area2 = 0
        region.accepts.forEach {
            area2+=it*9
        }
        println("$area1,$area2")
        if (area1>= area2) {
            tot++
        }
    }
    println(tot)
    // what the serendipitous f*ck ?

// all this is no longer needed, ...
// kinda sucks.
    var canFit = false
    regions.forEach { region ->
        val bag = region.createBag()

        var cnt = 0
        //pretest
        while ( bag.hasNextState()) {
            bag.nextState()
            cnt++
        //print(".")
        }
        println(cnt)
    }

    regions.forEach { region ->
        val bag = region.createBag()

        while (!canFit && bag.hasNextState()) {
            // try one state
            bag.forEach { state ->
                val shape = presents[state.presentNr].rotations[state.rotation]
                val px = state.posX
                val py = state.posY
                canFit = region.fit(shape, px, py)
                if (!canFit) {
                    return@forEach
                }
            }
            region.reset()
            bag.nextState()
        }
        println(canFit)
    }
}


fun read(input: AocInput) {
    var presentNr = 0
    var presentData = Array(3) { 0L }
    var presentWord = 0


    for (i in 0..<input.data.size) {
        val line = input.data[i]
        if (line.contains("x")) {
            val parts = line.split(":")
            val dims = parts[0].split("x")
            val accepts = parts[1].substring(1).split(" ")
            regions.add(
                Region(
                    w = dims[0].toInt(),
                    h = dims[1].toInt(),
                    accepts = accepts.map { it.toInt() }
                )
            )
        } else {
            // present
            if (line.contains(":")) {
                presentNr = line.substring(0, line.length - 1).toInt()
                presentData = Array(3) { 0L }
                presentWord = 0
            } else {
                if (line.isNotEmpty()) {
                    var l = 0L
                    for (j in 0..<line.length) {
                        if (line[line.length - 1 - j] == '#') {
                            l = l or (1L shl j)
                        }
                    }
                    presentData[presentWord] = l
                    presentWord++
                } else {
                    presents.add(
                        Present(presentNr, presentData)
                    )
                }
            }
        }
    }
}