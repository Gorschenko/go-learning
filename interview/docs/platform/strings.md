<h2>Интерфейсы</h2>

<h3>Теория:</h3>

<ul>
    <li>
        Каждый символ строки - 1-4 байта.
    </li>
    <li>
        h = 0_1101000 - 1 байт, 0_ говорит о том, что символ состоит из 
        Для записи 2-х байтов или более заполняются high order bits.
        1 байта.
        п = 00000_10000111111 - так могло бы быть, но непонятно как разделять символы.
        110_10000 - 1 байт, 110_ говорит о том, что символ состоит из 2 байтов; 10_111111 - 2 байт.
        3 байта - маска состоит из 1110_, 10_, 10_.
        4 байта - маска состоит из 11110_, 10_, 10_, 10_.
    </li>
    <li>
        <pre><code>
            s:= "привет"
            fmt.Println(s[1]) // выведет 2 байт символа П.
        </code></pre>
    </li>
</ul>

<h3>Задачи:</h3>

<details>
    <summary>
        Лекция 4.2. Вариант №1. Что будет выведено?
    </summary>
    <p>Исходные данные</p>
    <pre><code>
        func main() {
            greetings := "привет как дела" // 15 символов
            fmt.Println(len(greetings))
            fmt.Println("%v %b %c \n", greetings[1], greetings[1], greetings[1])
            //
            runes := []rune(greetings)
            fmt.Println("%c/n", runes[1])
            //
            greetings[0] = "s"
            fmt.Println(greetings)
        }
    </pre></code>
    <details>
        <summary>
            Решение
        </summary>
        <pre><code>
            func main() {
                greetings := "привет как дела" // 15 символов
                fmt.Println(len(greetings)) // 28
                fmt.Prinln("%v %b %c \n", greetings[1], greetings[1], greetings[1]) // число, байт, символ
                //
                runes := []rune(greetings)
                fmt.Println("%c/n", runes[1]) // р
                //
                greetings[0] = "s"
                fmt.Println(greetings) // ошибка компиляции
            }
        </pre></code>
    </details>
</details>

<details>
    <summary>
        Лекция 4.2. Вариант №2. Как изменить строку?
    </summary>
    <p>Исходные данные</p>
    <pre><code>
        func main() {
            greetings := "привет как дела"
            fmt.Println(greetings)
        }
    </pre></code>
    <details>
        <summary>
            Решение
        </summary>
        <pre><code>
            func main() {
                greetings := "привет как дела"
                fmt.Println(greetings)
                runes := []rune(greetings)
                runes[1] = 's'
                result := string(runes)
                fmt.Println(result)
            }
        </pre></code>
    </details>
</details>