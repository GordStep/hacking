# Hacking
Теперь ты можешь почуствовать себя хакером

![Запись экрана от 2024-02-10 22-19-09](https://github.com/GordStep/hacking/assets/106314109/a48ba2a0-93b5-4257-968b-b89239e81fe3)


## О принципе работы
Не бойтесь программа ничего не далает с вашим компьютером, кроме засорения собой памяти :)

### Способ окрашивания текста
```cpp
cout << "[\033[37m  Load resurs  \033[0m] " << text[t] << endl;
```
Вместо "37" мы пишем число отвечающее за цвет(коды цветов можно найти в интернете или методом подбора)

# Способ вывода надписи Loading...

``` cpp
for (int i = 0; i < 3; i++)
{
  for (int j = 0; j < 3; j++)
  {
    cout << "." << flush;
    this_thread::sleep_for(chrono::milliseconds(300));
  }
  cout << '\b' << '\b' << '\b';
}
cout << endl;
```
Здесь символ '\b' - забор, используется для перемещения курсора на один символ назад.
'flush' здесь используется для обхода ошибки с задержкой, при которой задержка срабатывает зарание

Командой:
```cpp
\033[#;#dH.\n // вместо "#;#" пишем "строку;столбец"
```
Мы переносим курсор вывода в нужные нам координаты в консоли








