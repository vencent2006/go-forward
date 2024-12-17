const fs = require('fs').promises

async function readFiles() {
  try {
    const data1 = fs.readFile('file1.txt', 'utf8')
    const data2 = fs.readFile('file2.txt', 'utf8')
    const data3 = fs.readFile('file3.txt', 'utf8')

    console.log('data from all files:', data1, data2, data3)
  } catch (err) {
    console.error('Error reading files:', err)
  }
}

readFiles()
