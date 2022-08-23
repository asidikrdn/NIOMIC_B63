let foo = [32,33,16,40,9,7,11,2,44,1,66];

const cekBilDiatas15 = (bil) => {
  return bil > 15;
};

console.log(foo.filter(cekBilDiatas15));