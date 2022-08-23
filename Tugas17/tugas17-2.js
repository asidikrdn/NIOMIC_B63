let foo = `Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.`

let bar = foo.substr(0,6);
// console.log(bar);
console.log(bar.charCodeAt(0)+'-'+bar.charCodeAt(1)+'-'+bar.charCodeAt(2)+'-'+bar.charCodeAt(3)+'-'+bar.charCodeAt(4)+'-'+bar.charCodeAt(5));