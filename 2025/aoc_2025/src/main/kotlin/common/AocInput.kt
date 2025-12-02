package common

class AocInput (val data :  List<String>) {
    fun print() {
        data.forEach { println( it ) }
    }
}