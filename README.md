# Baku - Convert Indonesian Slang Word to Formal Word

## Bahasa Indonesia
Baku adalah library yang dapat digunakan untuk mengubah kata gaul atau kata sehari-hari menjadi kata yang baku
1. Dengan menggunakan konsep yang sama dengan kamus, saya mengubah pengalaman penggunaan kamus menjadi code
2. Proyek ini berfokus pada kontribusi github tanpa ada batasan. Itulah yang membuat saya menggunakan JSON sebagai media penulisan data kamus agar orang lain dengan kemampuan komputer dapat berkontribusi dengan mudah. Improvisasi dari berbagai hal sangat diterima.

## English
Baku is a library that can be used to translate or convert any indonesian slang word or texting word to get the formal word
1. Using the basic concept as dictionary book with thumb index, I transalated the basic flow of using dictionary book to code
2. This project focus on github contribution with no restricted experince. That's why I use JSON as data storing to allow people from any computer background to work with because its simplicity. Any improvement is welcomed!

## The Art of Using Thumb Index Search in Code
```
1. Get the first character of the inputed value
2. Access the Index of first character inside the dictionary array
3. Get the Formal Word by Accessing the index of inputed value on selected thumb index array
```

## How to Run Example
You can use both `go build` then run the binary or use `go run main.go` on example directory

## TODO 
This needs more data to be written on dictionary