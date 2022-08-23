let foo = [2, 39, 76, 50, 9, 7, 41, 2, 24, 1, 16];

// Fungsi callback sorting ascending
const asc= (x,y) => {
  return x-y;
};
// Fungsi callback sorting descending
const desc= (x,y) => {
  return y-x;
};
// Fungsi callback filter bilangan diatas 10
const cekBilDiatas10 = (bil) => {
  return bil>10;
};

// Versi sesuai soal
console.log('Sebelumnya : '+foo);
console.log('Ascending : '+foo.sort());
console.log('Descending : '+foo.sort().reverse());
console.log('Filter > 10 : '+foo.filter(cekBilDiatas10));

console.log();

// Versi pengurutan berdasarkan nilai angka
console.log('Sebelumnya : '+foo);
console.log('Ascending : '+foo.sort(asc));
console.log('Descending : '+foo.sort(desc));
console.log('Filter > 10 : '+foo.filter(cekBilDiatas10));
