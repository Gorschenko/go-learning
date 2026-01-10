<h2>Слайсы</h2>

<h3>Особенности:</h3>
<ul>
    <li>
        Между функциями передается по ссылке.
    </li>
    <li>
        Если добавлять элементы a = append(a, []int{1, 2, 3, 4, 5}...) (обазятельное нечетное кол-во элементов),
        то длина будет всегда на 1 больше
    </li>
</ul>

<h3>Теория:</h3>

<details style="margin-bottom: 12px;">
    <summary>
        Что такое слайс?
    </summary>
    <p>
        Динамический массив
    </p>
</details>

<details style="margin-bottom: 12px;">
    <summary>
        Из чего состоит слайс?
    </summary>
    <p>
        Длина, вместимость, ссылка на массив.
    </p>
    <pre><code>
        type slice sturct {
            array unsafe.Pointer
            len int
            cap int
        }
    </code></pre>
</details>

<details style="margin-bottom: 12px;">
    <summary>
        Как работает append? Как рассчитыватются начальные len и cap?
    </summary>
    <pre><code>
        a := []int{1,2,3} // l=3, c=3
        //
        a := []int{} // l=0, c=0
        a = append(a, []int{1,2,3}...) // l=3, c=3
        //
        a := []int{} // l=0, c=0
        for i := range 3 {
            a = append(a, i)
        }
        //
        // i=0 l=1 c=1
        // i=1 l=2 c=2
        // i=2 l=3 c=4
    </code></pre>
</details>

<details>
    <summary>
        Механика работы ссылочного типа. Что происходи с ссылкой при копирование?
    </summary>
    <p>
        Если копируется слайс, и в него добавляется новый элемент,
        то будет выделена новая память для нового слайса и изменится ссылка.
    </p>
</details>

<h3>Задачи:</h3>

<details>
    <summary>
       Лекция 3.5 Копирование слайсов и добавление новых элементов.
       Вариант №1
    </summary>
    <p>Исходные данные</p>
    <pre><code>
        a := []int{1, 2, 3}
        fmt.Println("cap(a) = ", cap(a))
        //
        b := append(a, 4)
        c := append(a, 5)
        //
        c[1] = 0
        //
        fmt.Println("a =", a)
        fmt.Println("b =", b)
        fmt.Println("c =", c)
    </pre></code>
    <details>
        <summary>
            Решение
        </summary>
        <pre><code>
            a := []int{1, 2, 3} // l= 3 cap =3
            fmt.Println("cap(a) = ", cap(a)) // c=3
            //
            b := append(a, 4) // l=4, cap=6, new [], b=[1,2,3,4],0,0
            c := append(a, 5) // l=4, cap=6, new [], c=[1,2,3,5],0,0
            //
            c[1] = 0 // c=[1,0,3,5],0,0
            //
            fmt.Println("a =", a)
            fmt.Println("b =", b)
            fmt.Println("c =", c)
        </pre></code>
    </details>
</details>

<details>
    <summary>
       Лекция 3.6 Копирование слайсов и добавление новых элементов.
       Вариант №2. Добавляем элементы в исходный слайс исключительным способом.
    </summary>
    <p>Исходные данные</p>
    <pre><code>
        a := []int{}
        a = append(a, []int{1, 2, 3, 4, 5}...)
        fmt.Println("cap(a) = ", cap(a))
        //
        b := append(a, 6)
        c := append(a, 7)
        //
        c[1] = 0
        //
        fmt.Println("a =", a)
        fmt.Println("b =", b)
        fmt.Println("c =", c)
    </pre></code>
    <details>
        <summary>
            Решение
        </summary>
        <pre><code>
            a := []int{} // l=0, c=0
            a= append(a, []int{1, 2, 3, 4, 5}...) // l=5, c=6, a=[1,2,3,4,5],0
            fmt.Println("cap(a) = ", cap(a)) // c=6
            //
            b := append(a, 6) // l=6, c=6, based on a, b=[1,2,3,4,5,6], a=[1,2,3,4,5],6
            c := append(a, 7) // l=6, c=6, based on a, c=[1,2,3,4,5,7], b=[1,2,3,4,5,7], a=[1,2,3,4,5],7
            //
            c[1] = 0 // c=[1,0,3,4,5,7], b=[1,0,3,4,5,7], a=[1,0,3,4,5],7
            //
            fmt.Println("a =", a)
            fmt.Println("b =", b)
            fmt.Println("c =", c)
        </pre></code>
    </details>
</details>

<details>
    <summary>
       Лекция 3.7 Копирование слайсов и добавление новых элементов.
       Вариант №3. Добавляем элементы в исходный слайс исключительным способом, копируем другой массив
    </summary>
    <p>Исходные данные</p>
    <pre><code>
        a := []int{}
        a = append(a, []int{1, 2, 3, 4, 5}...)
        fmt.Println("cap(a) = ", cap(a))
        //
        b := append(a, 6)
        c := append(b, 7)
        //
        c[1] = 0
        //
        fmt.Println("a =", a)
        fmt.Println("b =", b)
        fmt.Println("c =", c)
        //
        d := c[0:12]
        fmt.Println("d =", d)
    </pre></code>
    <details>
        <summary>
            Решение
        </summary>
        <pre><code>
            a := []int{} // l=0, c=0
            a = append(a, []int{1, 2, 3, 4, 5}...) // l=5, c=6, a=[1,2,3,4,5],0
            fmt.Println("cap(a) = ", cap(a)) // c=6
            //
            b := append(a, 6) // l=6, c=6, based on a, b=[1,2,3,4,5,6], a=[1,2,3,4,5],6
            c := append(b, 7) // l=7, c=12, new arr, c=[1,2,3,4,5,6,7],0,0,0,0,0
            //
            c[1] = 0 // c=[1,0,3,4,5,6,7],0,0,0,0,0
            //
            fmt.Println("a =", a)
            fmt.Println("b =", b)
            fmt.Println("c =", c)
            //
            d := c[0:12]
            fmt.Println("d =", d) // d=[1,0,3,4,5,6,7,0,0,0,0,0]
        </pre></code>
    </details>
</details>

<details>
    <summary>
       Лекция 3.8 Копирование слайсов и добавление новых элементов.
       Вариант №4. Заполнение слайса через цикл
    </summary>
    <p>Исходные данные</p>
    <pre><code>
        a := []int{}
        for i := range 3 {
            a = append(a, i+1)
        }
        fmt.Println("cap(a) = ", cap(a))
        //
        b := append(a, 4)
        c := append(b, 5)
        //
        c[1] = 0
        //
        fmt.Println("a =", a)
        fmt.Println("b =", b)
        fmt.Println("c =", c)
        //
        d := a[0:4]
        fmt.Prinln("d =", d)
    </pre></code>
    <details>
        <summary>
            Решение
        </summary>
        <pre><code>
            a := []int{} // l=0, c=0
            for i := range 3 {
                a = append(a, i+1)
            }
            // i=0, l=1, c=1
            // i=1, l=2, c=2
            // i=2, l=3, c=4, a=[1,2,3],0
            fmt.Println("cap(a) = ", cap(a)) // c=4
            //
            b := append(a, 4) // l=4, c=4, based on a, b=[1,2,3,4]
            c := append(b, 5) // l=5, c=8, new array, c=[1,2,3,4,5],0,0,0
            //
            c[1] = 0 // c=[1,0,3,4,5],0,0,0
            //
            fmt.Println("a =", a)
            fmt.Println("b =", b)
            fmt.Println("c =", c)
            //
            d := a[0:4]
            fmt.Prinln("d =", d) // a=[1,2,3,4]
        </pre></code>
    </details>
</details>

<details>
    <summary>
       Лекция 3.9 Копирование слайсов и добавление новых элементов.
       Вариант №4. Заполнение слайса через цикл
    </summary>
    <p>Исходные данные</p>
    <pre><code>
        a := []int{1, 2, 3, 4, 5}
        b := a[2:4]
        c := append(b, 10)
        c[1] = 55
        //
        fmt.Println("a =", a)
        fmt.Println("b =", b)
        fmt.Println("c =", c)
        //
        d := b[:3]
        fmt.Println("d =", d)
    </pre></code>
    <details>
        <summary>
            Решение
        </summary>
        <pre><code>
            a := []int{1, 2, 3, 4, 5} // l=5, c=5
            b := a[2:4] // l=2, c=3, based on a, b=[3,4],5
            c := append(b, 10) // l=3, c=3, based on b, c=[3,4,10], b=[3,4],10 , a=[1,2,3,4,10]
            c[1] = 55 // c=[3,55,10], b=[3,55],10, a=[1,2,3,55,10]
            //
            fmt.Println("a =", a)
            fmt.Println("b =", b)
            fmt.Println("c =", c)
            //
            d := b[:3]
            fmt.Println("d =", d) // d=[3,55,10]
        </pre></code>
    </details>
</details>