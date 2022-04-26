Http basic auth adalah salah satu teknik otektikasi http request
method ini membutuhkan 2 parameter yaitu username dan password yang akan dimasukan ke dalam header request
method ini tidak memerlukan cookies maupun session

Namun username dan password yang akan dimasukan ke header tidak se"mentah"itu, harus kita rahasiakan dengan cara meng-encode terlebih dahulu.Dengan syarat username dan password dipisahkan dengan :

Contoh encode
`bade64encode("username:password")`

Pada kasus kali ini kita akan membuat web service sederhana yang hanya menggunakan 1 endpoint dengan detail :
1. /student -> menampilkan semua student
2. /student?id=2 -> menampilkan data id 2

Kita butuh 3 file yaitu middleware.go, main.go, student.go