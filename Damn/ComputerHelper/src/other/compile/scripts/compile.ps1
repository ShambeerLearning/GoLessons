# Устанавливаем кодировку UTF-8
[Console]::OutputEncoding = [System.Text.Encoding]::UTF8

# Определяем переменные
$GRAALPATH = "C:\Program Files\graalvm21\bin\native-image"
$JARPATH = "..\..\..\..\out\artifacts\ComputerHelper_jar\ComputerHelper.jar"
$PROGRAMNAME = "..\build\ComputerHelperApp-x32"

# Выполняем сборку Native Image
& $GRAALPATH -jar $JARPATH --no-fallback --enable-url-protocols=http,https -o $PROGRAMNAME