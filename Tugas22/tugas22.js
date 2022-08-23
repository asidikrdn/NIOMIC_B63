let foo = 'Saya ingin belajar bersama';
let bar = foo.split(' ');
// console.log(bar);

bar.forEach(
  function(value,key){
    console.log('Item : '+value+' Index ke '+key);
  }
);