# gocifromslovom
Изписване на число с думи (Цифром -> Словом) на Go

## Употреба
Импортирайте пакета gocifromslovom

```import github.com/dekiland/gocifromslovom```

Конвертиране цифром -> словом
```go
  str := gocifromslovom.ConvertMale(21) // "двадесет и един"
  ...
  str := gocifromslovom.ConvertFemale(-31) // "минус тридесет и една"
  ...
  str := gocifromslovom.ConvertNeutral(1000001) // "един милион и едно"
  ...
  str := gocifromslovom.ConvertMale(101) + " лева " + gocifromslovom.ConvertFemale(51) + " стотинки" // "сто и един лева и петдесет и една стотинки"
  ...
```

