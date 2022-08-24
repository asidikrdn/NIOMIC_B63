let sambungKata = (...kata) => {
  return kata.join(' ');
};

let foo = sambungKata('No', 'rumah saya', 30, 'A');
let bar = sambungKata('Jalan pangeran Tirtayasa No', 70);

console.log(foo);
console.log(bar);