function deretAngka(n) {
  // tulis code jawabanmu di sini
  let deret=[];
  for(let i=1; i<=n; i++) {
    if(i%3===0 && i%5===0){
      deret.push('NIOMIC');
    }
    else if(i%3===0){
      deret.push('NIO');
    }
    else if(i%5===0){
      deret.push('MIC');
    }
    else {
      deret.push(i);
    }
  }
  return deret.join(' ');
}


console.log(deretAngka(10))
console.log(deretAngka(20))
console.log(deretAngka(30))