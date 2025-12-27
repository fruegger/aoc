package common

typealias Vector<T> = Array<T>

fun Vector<Real>.sum(): Real {
    var sum = Real(0)
    for (element in this) {
        sum += element
    }
    return sum
}

typealias Matrix<T> = Array<Vector<T>>

inline val <T> Matrix<T>.rows get() = size
inline val <T> Matrix<T>.cols get() = this[0].size

fun <T> Vector<T>.toString(maxLen: Int = 0): String {
    var res = ""
    forEach {
        val st = it.toString().padStart(maxLen, ' ')
        res = "$res$st,"
    }
    return res.dropLast(1)
}

fun <T> Vector<T>.sameAs(other: Vector<T>): Boolean {
    if (size != other.size) {
        return false
    }
    for (i in 0 until size) {
        if (this[i] != other[i]) {
            return false
        }
    }
    return true
}

fun <T> Vector<T>.print(maxLen: Int = 0) {
    print("|${toString(maxLen)}|")
}

fun <T> Matrix<T>.print(maxLen: Int = 0) = forEach {
    println("|${it.toString(maxLen)}|")
}

fun <T> Matrix<T>.leftmostNonZero(row: Int, col: Int, zero: T): Pair<Int, Int> {
    var leftmost = cols
    var res = col
    for (i in row..<rows) {
        var j = 0
        while (j < cols && this[i][j] == zero) {
            j++
        }
        if (j < leftmost) {
            res = i
            leftmost = j
        }
    }
    return Pair(res, leftmost)
}

fun <T> Matrix<T>.findPivot(startRow: Int, col: Int): Int? {
    return (startRow until rows).firstOrNull { this[it][col] != 0 }
}


fun <T> Matrix<T>.swapRows(r1: Int, r2: Int) {
    val r = this[r1]
    this[r1] = this[r2]
    this[r2] = r
}

fun gaussReal(a: Matrix<Real>, b: Vector<Real>): Matrix<Real> =
    gauss(a, b, Real(0.0))

fun gaussRational(a: Matrix<Rational>, b: Vector<Rational>): Matrix<Rational> =
    gauss(a, b, Rational(0, 1))

@OptIn(ExperimentalStdlibApi::class)
fun <T : Num<T>> gauss(a: Matrix<T>, b: Vector<T>, zero: T): Matrix<T> {
    val r = augmentMatrix(a, b)
    var h = 0
    var k = 0
    while (h < r.rows && k < r.cols - 1) {
        val pivot = r.leftmostNonZero(h, k, zero).first
        if (pivot >= r.rows) {
            //done
            return r
        }
        if (r[pivot][k] == zero) {
            k++
        } else {
            r.swapRows(h, pivot)
            for (i in 0..<r.rows) {
                val f = r[i][k] / r[h][k]
                if (i != h) {
                    for (j in k + 1..<r.cols) {
                        r[i][j] = r[i][j] - r[h][j] * f
                    }
                    r[i][k] = zero
                }
            }
            h++
            k++
        }
    }
    return r
}

// int-version of gauss, according to wikipedia. Very rarely used, there are simpler ways.0
fun bareiss(a: Matrix<Int>, b: Vector<Int>): Matrix<Int> {
    val r = augmentMatrix(a, b)

    var previousF = 1

    var h = 0
    var k = 0

    while (h < r.rows && k < r.cols - 1) {
        val pivot = r.findPivot(h, k)

        if (pivot == null) {
            k++
            continue
        }

        if (r[pivot][k] == 0) {
            k++
        } else {

            r.swapRows(h, pivot)

            val f = r[h][k]
            for (i in 0 until r.rows) {
                if (i != h) {
                    val g = r[i][k]
                    for (j in k until r.cols) {
                        r[i][j] = (f * r[i][j] - g * r[h][j]) / previousF
                    }
                }
            }

            previousF = f
            h++
            k++
        }
    }
    return r
}

@OptIn(ExperimentalStdlibApi::class)
fun <T> augmentMatrix(a: Matrix<T>, b: Vector<T>): Matrix<T> {
    val r = a.copyOf()
    for (i in 0..<r.rows) {
        r[i] = a[i].copyOf(a.cols + 1) {
            b[i]
        }
    }
    return r
}

fun <T : Num<T>> zeroSolution(
    a: Matrix<T>,
    result: Vector<T>,
    zero: T
) {
    var k = 0
    for (i in 0..<a.rows) {
        while (k < a.cols && a[i][k] == zero) {
            k++
        }
        if (k < a.cols - 1) {
            result[k] = a[i].last() / a[i][k]
        }
        k++
    }
}

fun zeroSolutionRational(a: Matrix<Rational>): Vector<Rational> {
    val zero = Rational(0)

    val result = Array(a.cols - 1) { zero }
    zeroSolution(a, result, zero)
    return result
}

fun zeroSolutionReal(a: Matrix<Real>): Vector<Real> {
    val zero = Real(0)
    val result = Array(a.cols - 1) { zero }
    zeroSolution(a, result, zero)
    return result
}

fun <T : Num<T>> nullSpace(
    a: Matrix<T>,
    zero: T,
    one: T,
    createResult: (cols: Int, rows: Int, zero: T) -> Matrix<T>
): Matrix<T> {
    val isFreeVar = findFreeVars(a, zero)
    val freeVars = isFreeVar.count { it }

    val result = createResult(a.cols - 1, freeVars, zero)

    var k = 0
    var h = 0
    var l = 0
    while (k < a.cols - 1) {
        if (h >= a.rows || a[h][k] == zero) {
            collectRow(a, result, k, l, isFreeVar, zero, one)
            l++
            k++
        } else {
            k++
            h++
        }
    }
    return result
}

private fun <T : Num<T>> collectRow(
    a: Matrix<T>,
    result: Matrix<T>,
    k: Int,
    l: Int,
    isFreeVar: List<Boolean>,
    zero: T,
    one: T
) {
    result[k][l] = one
    var j = 0
    var m = 0
    while (m < a.cols-1) {
            if (!isFreeVar[m]) {
                val p = a.leftmostNonZero(j, j, zero)
                if (p.second >= a.cols - 1) {
                    result[m][l] = zero - a[j][k]
                    // all zero col !
                } else {
                    result[m][l] = zero - a[j][k] / a[j][p.second]
                }
                j++
                m++
            } else {
              m++
            }
    }
}

fun <T> findFreeVars(a: Matrix<T>, zero: T): List<Boolean> {
    val result = mutableListOf<Boolean>()
    var h = 0
    var k = 0
    while (k < a.cols - 1) {
        if (h >= a.rows || k>=a.cols-1 || a[h][k] == zero) {
            result.add(true)
            k++
        } else {
            result.add(false)
            h++
            k++
        }
    }
    return (result)
}

fun <T> countZeroRows(a: Matrix<T>, zero: T): Int {
    var zeroRows = 0
    a.forEach { row ->
        var allZeroes = true
        for (i in 0..<row.size - 1) {
            if (row[i] != zero) {
                allZeroes = false
                break
            }
        }
        if (allZeroes) {
            zeroRows++
        }
    }
    return zeroRows
}

inline fun <reified T> createNumResult(cols: Int, rows: Int, zero: T) =
    Array(cols) {
        Array(rows) { zero }
    }

fun nullSpaceReal(a: Matrix<Real>): Matrix<Real> =
    nullSpace(a, Real(0.0), Real(1.0), ::createNumResult)

fun nullSpaceRational(a: Matrix<Rational>, b: Vector<Rational>): Matrix<Rational> =
    nullSpace(a, Rational(0), Rational(1), ::createNumResult)

fun specificSolution(freeVals: Vector<Int>, zero: Vector<Rational>, nullSpace: Matrix<Rational>): Vector<Rational> {
    val result = zero.copyOf()
    for (i in 0..<zero.size) {
        for (j in 0..<nullSpace.cols) {
            result[i] = result[i] + Rational(freeVals[j]) * nullSpace[i][j]
        }
    }
    return result
}

fun matMult(a: Matrix<Int>, x: Vector<Int>): Vector<Int> {
    val result = Array(a.rows) { 0 }
    for (i in 0..<a.rows) {
        for (j in 0..<a.cols) {
            result[i] += x[j] * a[i][j]
        }
    }
    return result
}
