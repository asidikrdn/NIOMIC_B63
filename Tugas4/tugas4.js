// Membuat variabel tinggi 3 siswa
siswaSatu = 165;
siswaDua = 175;
siswaTiga = 155;

// Program membandingkan tinggi siswa
if (siswaSatu > siswaDua && siswaSatu > siswaTiga) {
  if (siswaDua > siswaTiga) {
    console.log(siswaSatu);
    console.log(siswaDua);
    console.log(siswaTiga);
  }
  else {
    console.log(siswaSatu);
    console.log(siswaTiga);
    console.log(siswaDua);
  }
}

else if (siswaDua > siswaSatu && siswaDua > siswaTiga) {
  if (siswaSatu > siswaTiga) {
    console.log(siswaDua);
    console.log(siswaSatu);
    console.log(siswaTiga);
  }
  else {
    console.log(siswaDua);
    console.log(siswaTiga);
    console.log(siswaSatu);
  }
}

else if (siswaTiga > siswaSatu && siswaTiga > siswaDua) {
  if (siswaSatu > siswaDua) {
    console.log(siswaTiga);
    console.log(siswaSatu);
    console.log(siswaDua);
  }
  else {
    console.log(siswaTiga);
    console.log(siswaDua);
    console.log(siswaSatu);
  }
}