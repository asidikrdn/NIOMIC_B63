let cekBil = (bil) => {
  if(isFinite(bil)) {
    console.log('Angka '+bil+' NOT Infinity.');
  }
  else {
    console.log('Angka '+bil+' Infinity.');
  }
}; 

let deretBil = [2,39,76,50,9,7,41,2,24,1,16];

deretBil.forEach(cekBil);