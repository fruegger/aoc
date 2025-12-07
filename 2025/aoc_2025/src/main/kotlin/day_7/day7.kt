package day_7

import common.AocInput
import common.AocReader
import common.Cell
import common.Direction
import common.Maze
import common.Node
import common.Position
import common.VT100

class Manifold(input: AocInput, start: Position) : Maze(input, start, start) {
    override fun isPathSymbol(c: Cell): Boolean = false
    override fun distance(n1: Node<Cell>, n2: Node<Cell>) = 0

    var root: ManifoldNode? = null
    var visited = mutableMapOf<Position, ManifoldNode>()

    fun buildDag() {
        root = buildDag(findSplit((start)))
    }

    private fun buildDag(pos: Position): ManifoldNode? {
        if (pos.y < heigth && symbolAt(pos) == '^') {
            val node = ManifoldNode(pos)
            visited[pos] = node

            var leftPos = pos.move(Direction.LEFT)
            if (leftPos.x >= 0) {
                leftPos = findSplit(leftPos)
                if (visited[leftPos] == null) {
                    node.left = buildDag(leftPos)
                } else {
                    node.left = visited[leftPos]
                }
                if (node.left!= null) {
                    node.count = node.count + (node.left!!.count) - 1
                }
            }

            var rightPos = pos.move(Direction.RIGHT)
            if (rightPos.x < width) {
                rightPos = findSplit(rightPos)
                if (visited[rightPos] == null) {
                    node.right = buildDag(rightPos)
                } else {
                    node.right = visited[rightPos]
                }
                if (node.right!= null) {
                    node.count = node.count + (node.right!!.count) - 1
                }
            }
            return node
        }
        return null
    }

    private fun findSplit(pos: Position): Position {
        var pos2 = pos.move(Direction.DOWN)
        while (pos2.y < this.heigth && symbolAt(pos2) == '.') {
            pos2 = pos2.move(Direction.DOWN)
        }
        return pos2
    }
}

data class ManifoldNode(
    val pos: Position
) {
    var left: ManifoldNode? = null
    var right: ManifoldNode? = null
    var count = 2L
}

fun main() {
    val input = AocReader.startDay(7, "input")
    input.print()
    val startx = input.data[0].indexOf("S")
    val start = Position(startx, 0)
    val manifold = Manifold(input, start)

    manifold.buildDag()

    print("Part 1:")
    println(VT100.blue { "${manifold.visited.size}" })


    print("Part 2:")
    println(VT100.blue { "${manifold.root!!.count}" })
    // total ^ <- 1754
    // 1752 too low
}

