package common

data class Node<T>(
    val data : T,
    var distanceFromStart : Int,
    var predecessor : Node<T>? =null,
    var visited : Boolean = false
)

const val INFINITY = Int.MAX_VALUE

interface Graph<T> {
    val startNode : Node<T>
    val allNodes : List<Node<T>>
    fun neighboursOf(n : Node<T>) : List<Node<T>>
    fun distance(n1 : Node<T>, n2 : Node<T>) : Int
}

fun <T>djikstra(g : Graph<T>) {
    val queue = ArrayDeque<Node<T>>()
    // init
    g.allNodes.forEach {
        if (it==g.startNode) {
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
        x.visited =  true
        g.neighboursOf(x).forEach { y ->
            if (!y.visited) {
                val newDist = x.distanceFromStart + g.distance(x,y)
                if (newDist < y.distanceFromStart) {
                    y.distanceFromStart = newDist
                    y.predecessor = x
                }
            }
        }
    }
}

