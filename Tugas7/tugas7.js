// Membuat Array
let tinggiSiswa = [
  165,
  155,
  168,
  167,
  158,
  173,
  164,
  145,
  171,
  153
];

// Menampilkan Array dengan perulangan for
for (let i = 0; i < tinggiSiswa.length; i++) {
  console.log("Tinggi Siswa ke-" + (i+1) + " adalah " + tinggiSiswa[i]);
}

console.log("------------");

// Menampilkan Array dengan perulangan for of
for (dataTinggi of tinggiSiswa){
  console.log(dataTinggi);
}