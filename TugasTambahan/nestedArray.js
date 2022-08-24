function panggilNestedArray(value) {
  // tulis jawabanmu disini
  // Membuat array transpose
  let arrTransposed = [];

  // Mengambil panjang array luar (baris) dan array dalam (kolom)
  let outerLength = value.length;
  let innerLength = value[0] instanceof Array ? value[0].length : 0;

  // Membuat perulangan sesuai banyaknya kolom (untuk menjadi baris)
  for (let i = 0; i < innerLength; i++) {
    // Membuat variabel untuk menampung array dalam 1 baris
    let arrInner = [];

    // Membuat perulangan sesuai banyaknya baris (untuk menjadi kolom dalam 1 baris)
    for (let j = 0; j < outerLength; j++) {
      // console.log(value[j][i]);
      // arrInner.push(value[j][i]);

      // Mengambil element pertama (kolom pertama) pada setiap baris array untuk dijadikan satu baris array
      arrInner.push(value[j].shift());
    }

    // Menggabungkan array 1 baris ke dalam array induk
    arrTransposed.push(arrInner);
  }

  console.log(arrTransposed);
}

var nestedArray = [
  [1, 2, 3, 4],
  ['Mark Zuckerberg', 'Elon Musk', 'Bill Gates', 'Steve Jobs'],
  ['Facebook', 'Tesla', 'Microsoft', 'Apple']
];

panggilNestedArray(nestedArray);