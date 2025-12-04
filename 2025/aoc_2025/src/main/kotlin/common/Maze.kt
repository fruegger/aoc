package common

data class Cell(
    var symbol: Char,
    val pos: Position
)

abstract class Maze(input: AocInput, val start: Position, val end: Position) : Graph<Cell> {

    val nodes: Array<Array<Node<Cell>>>

    companion object {
        const val START = 'S'
        const val END = 'E'
        const val WALL = '#'
        const val NOTHING = '.'
        const val OBSTACLE = 'O'

        val neighborLeft = Distance(dx = -1, dy = 0)
        val neighborRight = Distance(dx = 1, dy = 0)
        val neighborUp = Distance(dx = 0, dy = -1)
        val neighborDown = Distance(dx = 0, dy = 1)
    }


    init {
        val rows = input.data.size
        val cols = input.data[0].length
        nodes = Array(rows) { row ->
            Array(cols) { col ->
                Node<Cell>(
                    data = Cell(
                        symbol = input.data[row][col],
                        pos = Position(x = col, y = row)
                    )
                )
            }
        }
    }

    override val startNode: Node<Cell>
        get() = nodes[start.y][start.x]

    override val allNodes: Sequence<Node<Cell>>
        get() = sequence {
            nodes.forEach { row ->
                row.forEach {
                    yield(it)
                }
            }
        }

    abstract fun isPathSymbol(c: Cell): Boolean
    override fun neighboursOf(n: Node<Cell>): Sequence<Node<Cell>> = sequence {
        var pos = n.data.pos.add(neighborUp)
        if (pos.y > 0 && isPathSymbol(nodes[pos.y][pos.x].data)) {
            yield(nodes[pos.y][pos.x])
        }
        pos = n.data.pos.add(neighborDown)
        if (pos.y + 1 < heigth && isPathSymbol(nodes[pos.y][pos.x].data)) {
            yield(nodes[pos.y][pos.x])
        }
        pos = n.data.pos.add(neighborLeft)
        if (pos.x > 0 && isPathSymbol(nodes[pos.y][pos.x].data)) {
            yield(nodes[pos.y][pos.x])
        }
        pos = n.data.pos.add(neighborRight)
        if (pos.x + 1 < width && isPathSymbol(nodes[pos.y][pos.x].data)) {
            yield(nodes[pos.y][pos.x])
        }
    }


    fun symbolAt(p : Position) = nodes[p.y][p.x].data.symbol
    fun changeSymbol(p : Position, sym : Char) {
        nodes[p.y][p.x].data.symbol = sym
    }

    val width : Int get() = nodes[0].size
    val heigth : Int get() = nodes.size

    fun print(highlight : (Node<Cell>) -> Boolean = { false }) {
        nodes.forEach { row ->
            row.forEach {
                val c = it.data.symbol
                if (highlight(it)) {
                    print(VT100.CODE_REVERSE)
                }
                val st = when (c) {
                    START,
                    END -> {
                        VT100.red { "$c" }
                    }

                    WALL -> {
                        VT100.blue { "$c" }
                    }

                    NOTHING -> {
                        "$c"
                    }

                    OBSTACLE -> {
                        VT100.yellow { "$c" }
                    }

                    else -> VT100.green { "$c" }

                }
                print(st)
            }
            println()
        }
    }

}