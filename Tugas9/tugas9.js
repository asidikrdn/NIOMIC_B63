// Membuat Object
let mahasiswa = {
  nama: 'Ahmad Sidik Rudini',
  jurusan: 'Teknik Informatika',
  ipk: 3.95,
  semester: 4
}

// Menampilkan Object dengan perulangan for in
for (key in mahasiswa) {
  console.log(key + ' mahasiswa = ' + mahasiswa[key]);
}