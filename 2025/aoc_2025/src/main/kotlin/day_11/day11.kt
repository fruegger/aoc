package day_11

import common.AocInput
import common.AocReader
import common.Graph
import common.Node
import common.VT100
import common.bfs

fun main() {
    val input = AocReader.startDay(11, "test")
    input.print()
    val circuit = Circuit(input, "you")
    var tot = 0
    bfs(circuit) {
        tot++
    }

    println()
    circuit.print()

    print("Part1: ")
    println(VT100.blue { "$tot" })

//    SVR->FFT, FFT->DAC, and DAC->out
    val input2 = AocReader.startDay(11, "input")
    val circuit2 = Circuit(input2, "svr")

    circuit2.startNode = circuit2.allNodes.first { it.data == "dac"}
    circuit2.endSym = "out"
    val tot2C = dfs(circuit2)
    print("Part2C: ")
    println(VT100.blue { "$tot2C" })
    //2C : 4860

    circuit2.startNode = circuit2.allNodes.first { it.data == "fft"}
    circuit2.endSym = "dac"

    val tot2B = dfs(circuit2)
    print("Part2B: ")
    println(VT100.blue { "$tot2B" })

    circuit2.startNode = circuit2.allNodes.first { it.data == "svr"}
    circuit2.endSym = "fft"
    val tot2A = dfs(circuit2)
    print("Part2A: ")
    println(VT100.blue { "$tot2A" })

    print("Part2: ")
    println(VT100.blue { "${tot2A*tot2B*tot2C}" })
    // 5959310 too low
    // 32003206 too low
    // 511378159390560

}

fun dfs(g: Graph<String>): Long {
    val visited = mutableSetOf<Server>()
    cache.clear()
    dfs(g, g.startNode, visited)
    return cache[g.startNode.data]!!.count
}

data class CacheElement(
    var count: Long = 0L,
    var done: Boolean = false
)

val cache = mutableMapOf<String, CacheElement>()

private fun dfs(
    g: Graph<String>,
    node: Server,
    visited: MutableSet<Server>
) {
    if (g.isGoal(node)) {
        val p = node.predecessor
        collectCacheInfo(p)
    } else {
        visited.add(node)
        if (cache[node.data] == null) {
            cache[node.data] = CacheElement()
        }
        val neighbors = g.neighboursOf(node)
        neighbors.forEach { w ->
            w.predecessor = node
            if (!visited.any { it.data == w.data }) {
                if (cache[w.data] != null && cache[w.data]!!.done) {
                    var p : Server? = node
                    while (p != null) {
                        cache[p.data]!!.count+= cache[w.data]!!.count
                        p = p.predecessor
                    }
                } else {
                    dfs(g, w, visited)
                }
            }
        }
        cache[node.data]?.done = true
        visited.remove(node)
    }
}

private fun collectCacheInfo(p: Node<String>?): Node<String>? {
    var p1 = p
    while (p1 != null) {
        cache[p1.data]!!.count++
        p1 = p1.predecessor
    }
    return p1
}

typealias Server = Node<String>

class Circuit(
    input: AocInput,
    startSym: String,
    var endSym : String="out"
) : Graph<String> {
    val nodes = mutableMapOf<String, Server>()
    val edges = mutableMapOf<String, MutableList<Server>>()

    override lateinit var startNode: Server

    override val allNodes: Sequence<Server> get() = nodes.values.asSequence()

    override fun neighboursOf(n: Server): Sequence<Server> =
        if (n.data==endSym) {
            emptySequence()
        } else {
            edges[n.data]?.asSequence() ?: emptySequence()
        }

    override fun distance(n1: Node<String>, n2: Node<String>): Int = 1
    override fun isGoal(n1: Node<String>): Boolean = n1.data == endSym

    init {
        input.data.forEach { line ->
            val pair = line.split(":")
            val nodeName = pair[0]
            val newNode = Server(nodeName)
            if (nodeName == startSym) {
                startNode = newNode
            }
            if (nodes[nodeName] == null) {
                nodes[nodeName] = newNode
                edges[nodeName] = mutableListOf()
            }
            pair[1].substring(1).split(" ").forEach { name ->
                val node = Server(name)
                if (nodes[name] == null) {
                    nodes[name] = node
                    edges[name] = mutableListOf()
                }
                edges[nodeName]!!.add(node)
            }
        }
    }

    fun print() {
        allNodes.forEach {
            print("${it.data}: ")
            neighboursOf(it).forEach {
                print("${it.data}, ")
            }
            println()
        }
    }
}

