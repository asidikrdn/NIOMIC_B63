// Cara 1, menggunakan perulangan for
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



// Cara 2, menggunakan map()
let transposeArray = (array) => {
  let arrayBaru = array[0].map((nilaiKolomArrayLama, indexKolomArrayLama) => {
    let barisArrayBaru = array.map((barisArrayLama) => {
      console.log(barisArrayLama[indexKolomArrayLama]);
      return barisArrayLama[indexKolomArrayLama];
    })
    console.log(barisArrayBaru);
    return barisArrayBaru
  });
  console.log(arrayBaru);
  return arrayBaru;
}
/**
Penjelasan Cara 2 

Perulangan ke-1 :
  indexKolomArrayLama = 0;
  didalamnya ada perulangan lagi,
    perulangan ke-1 :
    menghasilkan baris pertama dari array,
    kemudian ambil element ke 0 (sesuai indexKolomArrayLama) dari array baris pertama,
    kirim nilai tersebut ke barisArrayBaru.
    perulangan ke-2 :
    menghasilkan baris kedua dari array,
    kemudian ambil element ke 0 (sesuai indexKolomArrayLama) dari array baris kedua,
    kirim nilai tersebut ke barisArrayBaru.
    perulangan ke-3 :
    menghasilkan baris ketiga dari array,
    kemudian ambil element ke 0 (sesuai indexKolomArrayLama) dari array baris ketiga,
    kirim nilai tersebut ke barisArrayBaru.
  barisArrayBaru kini telah menjadi sebuah array yang berisi 3 element (yang dihasilkan dari tiap perulangan),
  kirim barisArrayBaru sebagai element ke-0 ke arrayBaru.
Perulangan ke-2 :
  indexKolomArrayLama = 1;
  didalamnya ada perulangan lagi,
    perulangan ke-1 :
    menghasilkan baris pertama dari array,
    kemudian ambil element ke 1 (sesuai indexKolomArrayLama) dari array baris pertama,
    kirim nilai tersebut ke barisArrayBaru.
    perulangan ke-2 :
    menghasilkan baris kedua dari array,
    kemudian ambil element ke 1 (sesuai indexKolomArrayLama) dari array baris kedua,
    kirim nilai tersebut ke barisArrayBaru.
    perulangan ke-3 :
    menghasilkan baris ketiga dari array,
    kemudian ambil element ke 1 (sesuai indexKolomArrayLama) dari array baris ketiga,
    kirim nilai tersebut ke barisArrayBaru.
  barisArrayBaru kini telah menjadi sebuah array yang berisi 3 element (yang dihasilkan dari tiap perulangan),
  kirim barisArrayBaru sebagai element ke-1 ke arrayBaru.
Perulangan ke-3 :
  indexKolomArrayLama = 2;
  didalamnya ada perulangan lagi,
    perulangan ke-1 :
    menghasilkan baris pertama dari array,
    kemudian ambil element ke 2 (sesuai indexKolomArrayLama) dari array baris pertama,
    kirim nilai tersebut ke barisArrayBaru.
    perulangan ke-2 :
    menghasilkan baris kedua dari array,
    kemudian ambil element ke 2 (sesuai indexKolomArrayLama) dari array baris kedua,
    kirim nilai tersebut ke barisArrayBaru.
    perulangan ke-3 :
    menghasilkan baris ketiga dari array,
    kemudian ambil element ke 2 (sesuai indexKolomArrayLama) dari array baris ketiga,
    kirim nilai tersebut ke barisArrayBaru.
  barisArrayBaru kini telah menjadi sebuah array yang berisi 3 element (yang dihasilkan dari tiap perulangan),
  kirim barisArrayBaru sebagai element ke-2 ke arrayBaru.
Perulangan ke-4 :
  indexKolomArrayLama = 3;
  didalamnya ada perulangan lagi,
    perulangan ke-1 :
    menghasilkan baris pertama dari array,
    kemudian ambil element ke 3 (sesuai indexKolomArrayLama) dari array baris pertama,
    kirim nilai tersebut ke barisArrayBaru.
    perulangan ke-2 :
    menghasilkan baris kedua dari array,
    kemudian ambil element ke 3 (sesuai indexKolomArrayLama) dari array baris kedua,
    kirim nilai tersebut ke barisArrayBaru.
    perulangan ke-3 :
    menghasilkan baris ketiga dari array,
    kemudian ambil element ke 3 (sesuai indexKolomArrayLama) dari array baris ketiga,
    kirim nilai tersebut ke barisArrayBaru.
  barisArrayBaru kini telah menjadi sebuah array yang berisi 3 element (yang dihasilkan dari tiap perulangan),
  kirim barisArrayBaru sebagai element ke-3 ke arrayBaru.
 */





var nestedArray = [
  [1, 2, 3, 4],
  ['Mark Zuckerberg', 'Elon Musk', 'Bill Gates', 'Steve Jobs'],
  ['Facebook', 'Tesla', 'Microsoft', 'Apple']
];

console.log(transposeArray(nestedArray));
console.log();
panggilNestedArray(nestedArray);




/**
// Cara 2, menggunakan map()
let transposeArray = (array) => {
  // Menjalankan method map() pada array dalam dan menyimpannya dalam variabel outerArray
  let outerArray = array[0].map((valueInnerArray, indexInnerArray) => {
    // Pada setiap perulangan array dalam, jalankan method map() pada array luar dan menyimpannya pada variabel innerArray
    let innerArray = array.map((valueOuterArray) => {
      // Proses ini akan menghasilkan 1 buah array tiap perulangannya

      console.log(valueOuterArray[indexInnerArray]);
      // Kembalikan element array dengan index sesuai perulangan array dalam
      return valueOuterArray[indexInnerArray];
    })

    console.log(innerArray);
    // mengembalikan nilai innerArray yang merupakan kombinasi dari element tiap innerArray sebelumnya
    return innerArray
  });

  console.log(outerArray);
  // mengembalikan nilai outerArray yang merupakan gabungan dari beberapa innerArray yang sudah didapatkan
  return outerArray;
}





var nestedArray = [
  [1, 2, 3, 4],
  ['Mark Zuckerberg', 'Elon Musk', 'Bill Gates', 'Steve Jobs'],
  ['Facebook', 'Tesla', 'Microsoft', 'Apple']
];

console.log(transposeArray(nestedArray));
 */