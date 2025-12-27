package day_10

import common.AocReader
import common.Matrix
import common.Rational
import common.VT100
import common.Vector
import common.VectorCombinations
import common.cols
import common.gaussRational
import common.matMult
import common.nextNumWithIOnes
import common.nullSpaceRational
import common.permutations
import common.print
import common.sameAs
import common.specificSolution
import common.zeroSolutionRational
import kotlin.collections.mutableListOf

object IntCombinations {

    private val combinationsNK = mutableListOf<MutableList<MutableList<Int>>>()

    init {

        for (i in 0..<15) {
            val i2 = i + 1
            val setN = mutableListOf<MutableList<Int>>()
            for (j in 0..<i) {
                val j2 = j + 1
                var combination = (1 shl j2) - 1
                val setNK = mutableListOf<Int>()
                for (k in 0..<permutations(i2.toLong(), j2.toLong())) {
                    setNK.add(combination)
                    combination = nextNumWithIOnes(combination)
                }
                setN.add(setNK)
            }
            combinationsNK.add(setN)
        }
    }

    // integers of n bits wih k ones
    fun combination(n: Int, k: Int, i: Int) = combinationsNK[n - 1][k - 1][i]
}

class Machine(input: String) {
    val lights: Int
    val buttons = mutableListOf<Int>()
    val jolts = mutableListOf<Int>()

    val a: Matrix<Int>

    val b: Vector<Int>
    val u: Vector<Rational>
    val v: Matrix<Rational>

    var done = false
    val freeVars: Int get() = v.cols

    init {
        val lightsAsString = input.split("]")[0].substring(1)
        var bits = 0
        var shift = 1
        var nrLights = 0
        lightsAsString.forEach {
            val bit = if (it == '#') 1 else 0
            if (bit == 1) {
                bits = bits or shift
            }
            shift *= 2
            nrLights++
        }
        lights = bits
        val rest = input.split("]")[1].split("{")

        val joltsAsString = rest[1].substring(0, rest[1].length - 1)
        joltsAsString.split(",").forEach {
            jolts.add(it.toInt())
        }
        b = Array(nrLights) { jolts[it] }

        val buttonsAsString = rest[0].substring(2, rest[0].length - 2).split(") (")

        val nrButtons = buttonsAsString.size
        a = Array(nrLights) {
            Array(nrButtons) { 0 }
        }

        for (row in 0..<nrButtons) {
            val buttonAsString = buttonsAsString[row]
            bits = 0
            buttonAsString.split(",").forEach {
                val value = 1 shl it.toInt()
                a[it.toInt()][row] = 1
                bits = bits xor value
            }
            buttons.add(bits)
        }

        val ra = Array(a.size) { row ->
            Array(a[row].size) { col ->
                Rational(a[row][col])
            }
        }
        val rb = Array(b.size) {
            Rational(b[it])
        }

        val r = gaussRational(ra, rb)
        u = zeroSolutionRational(r)
        v = nullSpaceRational(r, u)
    }


    fun findBestButtonCombination(): Int {
        for (bits in 1..buttons.count()) {
            // an i-bits number with all 1s
            val max = permutations(buttons.count().toLong(), bits.toLong())
            var k = 0
            while (k < max) {
                val c2 = IntCombinations.combination(buttons.count(), bits, k)

                val result = pushButtons(c2)
                if (lights == result) {
                    return bits
                }
                k++
            }
        }
        throw Exception("no combination")
    }

    fun pushButtons(combination: Int): Int {
        var result = 0
        var w = combination
        var cnt = 0
        while (w > 0) {
            val bit = w and 1
            w = w shr 1
            if (bit == 1) {
                result = result xor buttons[buttons.size - 1 - cnt]
            }
            cnt++
        }
        return result
    }

    fun print() {
        print(VT100.white { "[" })
        print(lights.toString(2))
        print(VT100.white { "] " })
        buttons.forEach {
            print(VT100.white { "(" })
            print(it.toString(2))
            print(VT100.white { ") " })
        }
        println()
    }
}

fun main() {
    val input = AocReader.startDay(10, "input")
    input.print()
    val machines = mutableListOf<Machine>()
    input.data.forEach {
        machines.add(Machine(it))
    }


    var tot = 0
    machines.forEach { machine ->
        tot += machine.findBestButtonCombination()
    }
    print("Part1 : ")

    println(VT100.blue { "$tot" })
    var tot2 = 0L

    machines.forEach { machine ->
        if (machine.done == false) {
            print("f:${machine.freeVars}, b:")
            machine.b.print()
            println()

            val max = machine.b.max()
            val c = VectorCombinations(machine.freeVars, max)
            var min = Int.MAX_VALUE
            while (c.hasNext()) {
                val x = c.next()
                val sr = specificSolution(x, machine.u, machine.v)
                val s = Array(sr.size) {
                    val v = sr[it]
                    if (v.den != 1) {
                        continue
                    } else {
                        v.num
                    }
                }
                val sum = s.sum()
                if (s.all { it >= 0 }) {
                    val y = matMult(machine.a, s)
                    if (y.sameAs(machine.b)) {
                        if (sum < min) {
                            min = sum
                        }
                    } else {
                        throw Exception("something wrong: a non-solution qualified.")
                    }
                }
            }
            if (min == Int.MAX_VALUE) {
                throw Exception("something wrong: no min found.")
            }

            tot2 += min
        }
    }

    print("Part2 : ")
    println(VT100.blue { "$tot2" })
    // 15132

}


