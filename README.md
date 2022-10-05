# mini-sort

$ minisort опции путь_к_файлу

## Опции
-k int  
&emsp;&emsp;&emsp;номер столбца для сортировки, не может быть меньше 1 (default 1)  
        
-b    игнорировать начальные пробелы  
-c    проверить отсортированы ли данные  
-n    cортировать по числовому значению  
-r    сортировать в обратном порядке  
-u    не выводить повторяющиеся строки  

## Примеры
  go run .\cmd\main.go data/2.txt  
  go run .\cmd\main.go -k 2 data/2.txt  
  go run .\cmd\main.go -c -k 2 data/2.txt  