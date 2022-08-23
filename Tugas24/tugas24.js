let foo = [32,33,16,40,9,7,11,2,44,1,66];

const asc= (x,y) => {
  return x-y;
};
const desc= (x,y) => {
  return y-x;
};

// Versi sesuai soal
console.log('Sebelumnya : '+foo);
console.log('Ascending : '+foo.sort());
console.log('Descending : '+foo.sort().reverse());

console.log();

// Versi pengurutan berdasarkan nilai angka
console.log('Sebelumnya : '+foo);
console.log('Ascending : '+foo.sort(asc));
console.log('Descending : '+foo.sort(desc));