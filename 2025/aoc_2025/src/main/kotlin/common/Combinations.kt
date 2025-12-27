package common


class VectorCombinations(val size: Int, val maxValue: Int) : Iterator<Vector<Int>> {

    var current = Array(size) { 0 }
    var next: Vector<Int>? = nextPermutationByMaxValue(current)
    var lastRead = false

    fun nextPermutationByMaxValue(current: Vector<Int>): Vector<Int>? {
        if (size <= 0) {
            return null
        }
        val next = current.copyOf()
        val currentMax = next.maxOrNull() ?: 0

        // Try to find the next permutation within the current max constraint
        var pos = next.size - 1
        while (pos >= 0) {
            // Can we increment this position while staying <= current_max?
            if (next[pos] < currentMax) {
                // Increment this position
                next[pos]++

                // Reset all positions to the right to 0
                for (i in pos + 1 until next.size) {
                    next[i] = 0
                }

                if (next.maxOrNull() == currentMax) {
                    return next
                }

                next[next.size - 1] = currentMax
                return next
            }

            pos--
        }

        // All positions are at current_max, move to next max value group
        val nextMax = currentMax + 1
        if (nextMax > maxValue) return null  // No more permutations

        // Start new group with smallest permutation having max = nextMax
        // That's all 0s except last position set to nextMax
        next.fill(0)
        next[next.size - 1] = nextMax

        return next
    }

    override fun next(): Vector<Int> {
        val result = current.copyOf()
        next?.let {
            current = it
            next = nextPermutationByMaxValue(current)
        }
        return result
    }

    override fun hasNext(): Boolean {
        if (next != null) {
            return true
        } else {
            if (lastRead) {
                return false
            } else {
                lastRead = true
                return true
            }
        }
    }
}

fun nextNumWithIOnes(combination: Int): Int {
    val t = (combination or (combination - 1)) + 1
    val t2 = t and -t
    val c2 = combination and -combination
    val w2 = t2 / c2
    val s2 = w2 shr 1
    return t or (s2 - 1)
}

fun permutations(n: Long, k: Long): Long = factorial(n) / (factorial(k) * factorial(n - k))
fun factorial(n: Long): Long = if (n <= 1) 1 else n * factorial(n - 1)
