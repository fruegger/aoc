package day_4

import common.AocInput
import common.AocReader
import common.Cell
import common.Distance
import common.Maze
import common.Node
import common.Position
import common.VT100

class PaperRollMaze(input: AocInput, start: Position, end: Position) : Maze(input, start, end) {
    companion object {
        const val PAPER_ROLL = '@'
    }

    override fun isPathSymbol(c: Cell) = true

    override fun distance(n1: Node<Cell>, n2: Node<Cell>) = 1

    fun isRoll(pos: Position) = symbolAt(pos) == PAPER_ROLL

    fun countAdjecentRolls(pos: Position) : Int {
        var cnt = 0
        for (dy in -1..1) {
            for (dx in -1..1) {
                    val p2 = pos.add(Distance(dx, dy))
                    if ((p2.x in 0..<width) && (p2.y in 0..<heigth)) {
                        if (pos!=p2 && isRoll(p2)) {
                            cnt++
                        }
                    }
            }
        }
        return cnt
    }
}

fun main() {
    val input = AocReader.startDay(4, "input")
    val start = Position(0, 0)
    val maze = PaperRollMaze(input, start, start)
    maze.print()
    var tot = 0
    maze.allNodes.forEach {
        val pos = it.data.pos
        if (maze.isRoll(pos) && maze.countAdjecentRolls(pos)<4) {
            tot++
        }
    }
    println()

    maze.print {
        maze.isRoll(it.data.pos) && maze.countAdjecentRolls(it.data.pos)<4
    }
    print("Part 1: ")
    println(VT100.blue { "$tot" })
}
