// Go (нужно явно указать указатель)
func LoadConfig() *Config {
    return &Config{Env: "dev", Port: 3000} // явно возвращаем указатель
}

config := LoadConfig()
config.Port = 4000 // меняем оригинальную структуру!


func LoadConfig() Config {  // без *
    return Config{}         // возвращается КОПИЯ структуры
}

config := LoadConfig()
config.Port = 4000  // меняем только локальную копию!


func LoadConfig() *Config { // с *
    return &Config{}        // возвращаем указатель на оригинал
}

config := LoadConfig()
config.Port = 4000  // меняем оригинальную структуру!

Итог:
*Config - "функция возвращает указатель на Config"
&Config{} - "создай Config и дай мне на него указатель"