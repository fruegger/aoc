package common

class Position(val x: Int, val y: Int) {
    fun equals(p2: Position): Boolean = x == p2.x && y == p2.y
    fun move(d: Direction): Position = Position(x = x + d.dx, y = y + d.dy)
    fun add(d: Distance): Position = Position(x = x + d.dx, y = y + d.dy)
    fun distanceTo(p2: Position): Distance = Distance(dx = p2.x - x, dy = p2.y - y)
}

typealias Dimension = Position

class Distance(
    val dx: Int,
    val dy: Int
)

enum class Direction(
    dx: Int,
    dy: Int,
    val symbol: String
) {
    RIGHT(dx = 1, dy = 0, symbol = ">"),
    DOWN(dx = 0, dy = 1, symbol = "v"),
    LEFT(dx = -1, dy = 0, symbol = "<"),
    UP(dx = 0, dy = -1, symbol = "^");

    private val dist: Distance = Distance(dx, dy)

    val dx: Int
        get() = dist.dx

    val dy: Int
        get() = dist.dy

    fun turnRight(): Direction =
        when (this) {
            RIGHT -> DOWN
            DOWN -> LEFT
            LEFT -> UP
            UP -> RIGHT
        }

    fun TurnLeft(): Direction =
        when (this) {
            RIGHT -> UP
            UP -> LEFT
            LEFT -> DOWN
            DOWN -> RIGHT
        }

    fun Opposite(): Direction =
        when (this) {
            RIGHT -> LEFT
            DOWN -> UP
            LEFT -> RIGHT
            UP -> DOWN
        }
}
