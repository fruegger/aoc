package common

object AocReader {

    fun startDay(dayNumber: Int, prefix: String): AocInput {
        val filename = "day_$dayNumber/$prefix.txt"
        val result = mutableListOf<String>()
        this::class.java.classLoader.getResourceAsStream(filename)?.bufferedReader()?.useLines { lines ->
            lines.forEach {
                result.add(it)
            }
        }
        print(VT100.green { "--- Day $dayNumber" })
        println(" [$prefix] ---")
        return AocInput(result)
    }


}