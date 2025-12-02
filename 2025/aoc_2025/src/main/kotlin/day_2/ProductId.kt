package day_2

data class ProductId(
    val first: String,
    val last: String
) {
    fun forAllInvalid(reps : Int, block: (String) -> Unit) {
        var i = first.toLong()
        while (i <= last.toLong()) {
            val id = "$i"
            if (id.repeatsExactly(reps)) {
                block(id)
            }
            i++
        }
    }
}

fun String.toProductId(): ProductId {
    val pair = split('-')
    return ProductId(pair[0], pair[1])
}

fun String.repeatsExactly(reps : Int) : Boolean {
    var result = true
    val repLen = length / reps
    if (length % reps != 0) {
        return false
    }
    var i = 0
    while (result && i < repLen) {
        for ( r in 1..reps-1) {
            result = result && this[i] == this[repLen*r+i]
        }
        i++
    }
    return result
}