function balikKata(kata) {
  // tulis jawabanmu di sini
  let i = kata.length - 1;
  let kataTerbalik='';
  for (let x = 0; x < kata.length; x++) {
    kataTerbalik += kata[i];
    i--;
  }
  return kataTerbalik;
}

// testCase
console.log(balikKata("Niomic!"))
console.log(balikKata("JavaScript"))
console.log(balikKata("alohahola"))
console.log(balikKata("Jawa_Barat"))