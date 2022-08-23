// Membuat Object
let mahasiswa = {
  nama: 'Ahmad Sidik Rudini',
  jurusan: 'Teknik Informatika',
  ipk: 3.95,
  semester: 4
}

// Menampilkan nilai awal object mahasiswa
console.log(mahasiswa);

console.log('--------------');

// Menampilkan nilai awal tiap property object mahasiswa
console.log(mahasiswa.nama);
console.log(mahasiswa.jurusan);
console.log(mahasiswa.semester);
console.log(mahasiswa.ipk);

console.log('============================');

// Merubah nilai  object mahasiswa
mahasiswa.ipk = 3.99;   // merubah nilai property ipk 
mahasiswa.semester = 5; // merubah nilai property semester
mahasiswa.fakultas = 'Fakultas Teknik dan Sains'  // menambah property fakultas

// Menampilkan nilai awal object mahasiswa
console.log(mahasiswa);

console.log('--------------');

// Menampilkan nilai tiap property object mahasiswa setelah diupdate
console.log(mahasiswa.nama);
console.log(mahasiswa.jurusan);
console.log(mahasiswa.semester);
console.log(mahasiswa.ipk);
console.log(mahasiswa.fakultas);