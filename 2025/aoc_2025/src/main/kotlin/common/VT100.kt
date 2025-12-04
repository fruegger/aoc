package common

object VT100 {
    private const val ESC='\u001B'
    const val CODE_GREEN = "$ESC[32m"
    const val CODE_BLUE = "$ESC[34m"
    const val CODE_WHITE = "$ESC[97m"
    const val CODE_RED = "$ESC[31m"
    const val CODE_YELLOW = "$ESC[33m"
    const val CODE_GRAY = "$ESC[90m"

    const val CODE_SAVE = "$ESC[s"
    const val CODE_RESTORE = "$ESC[u"

    const val CODE_UP = "$ESC[A"
    const val CODE_DOWN = "$ESC[B"
    const val CODE_RIGHT = "$ESC[C"
    const val CODE_LEFT = "$ESC[D"

    const val CODE_NORMAL = "$ESC[0m"
    const val CODE_BRIGHT = "$ESC[1m"
    const val CODE_DIM = "$ESC[2m"
    const val CODE_UNDERSCORE = "$ESC[4m"
    const val CODE_BLINK = "$ESC[5m"
    const val CODE_REVERSE = "$ESC[7m"

    fun highlightIf(condition : Boolean, code : String, s : () -> String) = if (condition) {"$code${s()}$CODE_NORMAL"} else s
    fun green(condition : Boolean = true, s : () -> String) = highlightIf(condition,CODE_GREEN, s)
    fun blue(condition : Boolean = true, s : () -> String) = highlightIf(condition,CODE_BLUE, s)
    fun red(condition : Boolean = true, s : () -> String) = highlightIf(condition,CODE_BLUE, s)
    fun white(condition : Boolean = true, s : () -> String) = highlightIf(condition,CODE_WHITE, s)
    fun yellow(condition : Boolean = true, s : () -> String) = highlightIf(condition,CODE_YELLOW, s)
}