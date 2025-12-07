package common

class AocInput (val data :  List<String>) {
    fun print() {
        data.forEach { println( it ) }
    }
    val height : Int get() = data.size
    val width : Int get() = data[0].length
}