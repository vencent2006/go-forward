setTimeout(() => {
  console.log('timeout callback')
}, 0)

setImmediate(() => {
  console.log('immediate callback')
})

console.log('main thread execution')
