package common

data class Node<T>(
    val data: T,
    var distanceFromStart: Int = INFINITY,
    var predecessor: Node<T>? = null,
    var visited: Boolean = false
)

typealias Path<T> = MutableList<Node<T>>

const val INFINITY = Int.MAX_VALUE

interface Graph<T> {
    val startNode: Node<T>
    val allNodes: Sequence<Node<T>>
    fun neighboursOf(n: Node<T>): Sequence<Node<T>>
    fun distance(n1: Node<T>, n2: Node<T>): Int
    fun isGoal(n1: Node<T>): Boolean = false
}

//todo .. add a visited list and remove the visited bool from the node ?
fun <T> djikstra(g: Graph<T>) {
    val queue = ArrayDeque<Node<T>>()
    // init
    g.allNodes.forEach {
        if (it == g.startNode) {
            it.distanceFromStart = 0
            queue.addFirst(it)
        } else {
            it.distanceFromStart = INFINITY
            queue.addLast(it)
        }
    }
    while (queue.isNotEmpty()) {
        // find the unvisited node that has the smallest distance to the start node
        var x = queue.first()
        var minDist = x.distanceFromStart
        queue.forEach {
            if (it.distanceFromStart < minDist) {
                x = it
                minDist = it.distanceFromStart
            }
        }

        queue.remove(x)
        x.visited = true
        g.neighboursOf(x).forEach { y ->
            if (!y.visited) {
                val newDist = x.distanceFromStart + g.distance(x, y)
                if (newDist < y.distanceFromStart) {
                    y.distanceFromStart = newDist
                    y.predecessor = x
                }
            }
        }
    }
}

fun <T> bfs(g: Graph<T>, match: (p: Path<T>) -> Unit) {
    val queue = ArrayDeque<Path<T>>()
    queue.add(mutableListOf(g.startNode))
    while (queue.isNotEmpty()) {
        val v = queue.removeFirst()
        if (g.isGoal(v.last())) {
            match(v)

        } else {

            val neighbours = g.neighboursOf(v.last())
            neighbours.forEach { w ->
                if (!v.any { it.data == w.data }) {

                    val v2 = mutableListOf<Node<T>>()
                    v2.addAll(v)
                    v2.add(w)
                    w.predecessor = v.last()
                    queue.addLast(v2)
                }
            }
        }
    }
}

fun <T> dfs(g: Graph<T>, match: (p: Node<T>) -> Unit) {
    val visited = mutableSetOf<Node<T>>()
    dfs(g, g.startNode, visited, match)
}

private fun <T> dfs(
    g: Graph<T>,
    node: Node<T>,
    visited: MutableSet<Node<T>>,
    match: (p: Node<T>) -> Unit
) {
    visited.add(node)
    val neighbors = g.neighboursOf(node)
    if (g.isGoal(node)) {
        match(node)
    } else {
        neighbors.forEach { w ->
            w.predecessor = node
            if (!visited.any { it.data == w.data }) {
                dfs(g, w, visited, match)
            }
        }
    }
    visited.remove(node)
}

