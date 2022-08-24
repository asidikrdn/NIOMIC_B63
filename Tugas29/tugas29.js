function regex() {
  let data = "Belajar menimba ilmu programming bersama Niomic";

  return /bersama/.exec(data);
}

console.log(regex());