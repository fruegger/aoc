package common

interface Num<T> {
    operator fun plus(other: T): T
    operator fun minus(other: T): T
    operator fun times(other: T): T
    operator fun div(other: T): T
}

data class Rational(val num: Int, val den: Int = 1) : Num<Rational> {
    init {
        require(den != 0)
    }

    fun reduced(): Rational {
        val g = gcd(kotlin.math.abs(num), kotlin.math.abs(den))
        val sign = if (den < 0) -1 else 1
        return Rational(sign * num / g, sign * den / g)
    }

    override operator fun plus(other: Rational) = Rational(num * other.den + other.num * den, den * other.den).reduced()
    override operator fun minus(other: Rational) =
        Rational(num * other.den - other.num * den, den * other.den).reduced()

    override operator fun times(other: Rational) = Rational(num * other.num, den * other.den).reduced()
    override operator fun div(other: Rational) = Rational(num * other.den, den * other.num).reduced()
}

data class Real(val r: Double) : Num<Real> {

    constructor( i : Int) : this(i * 1.0)
    override operator fun plus(other: Real) = Real(r + other.r)
    override operator fun minus(other: Real) = Real(r - other.r)
    override operator fun times(other: Real) = Real(r * other.r)
    override operator fun div(other: Real) = Real(r / other.r)
}

fun gcd(a: Int, b: Int): Int = if (b == 0) a else gcd(b, a % b)
